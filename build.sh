#!/bin/bash
set -e

IMAGE=registry.corp.lazycat.cloud/lzc/lzc-api-protoc

if [ "$1" = "" ]; then
    echo "usage $0 <command>"
    echo "      $0 build-image   # build the builder docker image"
    echo "      $0 push-image    # push the builder docker image to registry.corp.lazycat.cloud"
    echo "      $0 gen           # run generation"
    exit
fi

if [ "$1" = "build-image" ]; then
    cd tools
    rm -rf cache
    mkdir -p cache/usr/bin
    cd cache/usr
    CGO_ENABLED=0 GOBIN=$(pwd)/bin go install -ldflags "-w -s" github.com/golang/protobuf/protoc-gen-go@v1.5.3
    CGO_ENABLED=0 GOBIN=$(pwd)/bin go install -ldflags "-w -s" google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
    curl -o protoc.zip -L https://github.com/protocolbuffers/protobuf/releases/download/v23.4/protoc-23.4-linux-x86_64.zip &
    curl -o maven.tar.gz -L https://dlcdn.apache.org/maven/maven-3/3.9.6/binaries/apache-maven-3.9.6-bin.tar.gz &
    curl -o bin/jre.deb -L https://mirrors.ustc.edu.cn/adoptium/deb/pool/main/t/temurin-17/temurin-17-jre_17.0.10.0.0%2B7_amd64.deb & # or full jdk?
    curl -o bin/adoptium-ca.deb -L https://mirrors.ustc.edu.cn/adoptium/deb/pool/main/a/adoptium-ca-certificates/adoptium-ca-certificates_1.0.2-1_all.deb &
    curl -o bin/protoc-gen-grpc-java -L https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/1.47.0/protoc-gen-grpc-java-1.47.0-linux-x86_64.exe &
    curl -o nodejs.tar.xz -L https://registry.npmmirror.com/-/binary/node/v20.11.1/node-v20.11.1-linux-x64.tar.xz &
    wait # wait all parallel download jobs
    unzip protoc.zip
    tar -xf maven.tar.gz --strip-components 1
    chmod a+x bin/protoc-gen-grpc-java
    tar -xf nodejs.tar.xz --strip-components 1
    mv lib/node_modules bin/node_modules # fix path error
    cd ../..
    if [ ! -e ~/.docker/cli-plugins/docker-buildx ]; then
        mkdir -p ~/.docker/cli-plugins/
        curl -o ~/.docker/cli-plugins/docker-buildx -L https://github.com/docker/buildx/releases/download/v0.13.1/buildx-v0.13.1.linux-amd64
        chmod +x ~/.docker/cli-plugins/docker-buildx
    fi
    DOCKER_BUILDKIT=1 docker build -t $IMAGE .
    rm -rf cache
    exit
fi

if [ "$1" = "push-image" ]; then
    docker push $IMAGE
    exit
fi

# generate grpc ptotos, run inner docker
if [ "$1" = "gen-protos" ]; then
    PROTOC_ARGS=()
    if [[ "$PROTOC_GEN_GO" -eq 1 ]]; then
        OUT_GO=$(pwd)/lang/go/
        rm -rf $OUT_GO/{common,localdevice} || true
        PROTOC_ARGS+=(
            --go_out=$OUT_GO
            --go-grpc_out=$OUT_GO
            --go_opt=paths=source_relative
            --go-grpc_opt=paths=source_relative
        )
    fi
    if [[ "$PROTOC_GEN_JS" -eq 1 ]]; then
        OUT_JS=$(pwd)/lang/js/
        rm -rf $OUT_JS/{common,localdevice,sys} || true
        PROTOC_ARGS+=(
            --ts_proto_out=$OUT_JS
            --ts_proto_opt=outputClientImpl=grpc-web
            --ts_proto_opt=exportCommonSymbols=false
            --ts_proto_opt=esModuleInterop=true
            --ts_proto_opt=useAbortSignal=true
        )
    fi
    if [[ "$PROTOC_GEN_JAVA" -eq 1 ]]; then
        OUT_JAVA=$(pwd)/lang/java/lib/src/main/java/
        rm -rf $OUT_JAVA/ || true
        mkdir -p $OUT_JAVA || true
        PROTOC_ARGS+=(
            --java_out=$OUT_JAVA
            --grpc-java_out=$OUT_JAVA
        )
    fi
    pushd protos
    PROTOC_ARGS+=(
        ./sys/*/*.proto
        ./sys/*.proto
        ./localdevice/*.proto
        ./dlna/*.proto
        ./common/*.proto
    )
    protoc -I . "${PROTOC_ARGS[@]}"
    popd
    if [[ "$PROTOC_GEN_JS" -eq 1 ]]; then
        pushd lang/js/
        npm run build
        popd
    fi
    if [[ "$PROTOC_GEN_JAVA" -eq 1 ]]; then
        cp -r lang/java/lib/src java_api_workspace/
        if [[ "$ENABLE_JAR" == "" ]]; then
            echo "Skip generate JAR, You can enable generating jar by set env 'ENABLE_JAR'"
        else
            pushd java_api_workspace
            echo "java package start... "
            mvn -Dmaven.test.skip=true -Dmaven.repo.local=/tmp/maven clean package
            echo "java package complete... "
            popd
        fi
    fi
    exit
fi

# generate security content rules
if [ "$1" = "gen-security-content-rules" ]; then
    mv $(pwd)/protoc-gen-lzc /usr/bin/
    OUT_LZC=$(pwd)/scontext/
    mkdir -p $OUT_LZC || true
    pushd protos
    protoc -I . \
        --lzc_out=$OUT_LZC \
        ./sys/*/*.proto \
        ./sys/*.proto \
        ./localdevice/*.proto \
        ./common/*.proto
    popd
    # 上面直接用 -u $UID 会导致 npm 无法执行
    chown -R $ORIGIN_UID .
    exit
fi

# generate all
if [ "$1" = "gen" ]; then
    inner_builder="docker run -it --network host --rm -w $(pwd) -v $(pwd):$(pwd) $IMAGE"
    $inner_builder env NPM_CONFIG_UPDATE_NOTIFIER=false PROTOC_GEN_GO=1 PROTOC_GEN_JS=1 PROTOC_GEN_JAVA=0 ENABLE_JAR=0 ./build.sh gen-protos
    # 必须先生成 lang/go 的内容后才能生成 security content rules
    go build -ldflags "-w -s" -o protoc-gen-lzc tools/protoc-gen-lzc/permgen.go
    $inner_builder env ORIGIN_UID=$UID ./build.sh gen-security-content-rules
    pushd lang/go/
    go build
    popd
    exit
fi

# 将 java grpc api 文件传送到 ftp 里面
if [ "$1" = "push-java-lib" ]; then
    #PUTFILE=cloud-lazycat-apis-1.0.jar
    #ftp -v -n 192.168.1.189<<EOF
    #user anonymous lnks
    #binary
    #cd ./sdk
    #lcd ./java_api_workspace/target/
    #prompt
    #del $PUTFILE
    #put $PUTFILE
    #bye
    ##here document
    #EOF
    #echo "push java-lib complete"
    cd java_api_workspace/target
    md5sum cloud-lazycat-apis-1.0.jar | cut -c-32 > cloud-lazycat-apis-1.0.jar.md5
    scp -O -P 2022  cloud-lazycat-apis-1.0.jar cloud-lazycat-apis-1.0.jar.md5 root@192.168.1.189:/mnt/dl/sdk/
    exit
fi

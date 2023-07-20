#!/bin/bash
set -e

IMAGE=registry.lazycat.cloud/lzc/lzc-api-protoc

docker pull "$IMAGE"

docker run -ti -w $(pwd) -v $(pwd):$(pwd) -v /tmp/maven:/tmp/maven -e ENABLE_JAR=1 --rm "$IMAGE" env NPM_CONFIG_UPDATE_NOTIFIER=false PROTOC_GEN_GO=1 PROTOC_GEN_JS=1 PROTOC_GEN_JAVA=0 ./tools/gen.sh

#必须先生成lang/go的内容后才能生成security content rules
pushd tools/protoc-gen-lzc/
go build
popd
docker run -ti -w $(pwd) -v $(pwd):$(pwd) --rm registry.lazycat.cloud/lzc/lzc-api-protoc ./tools/gen_scontext_rules.sh
# 上面直接用-u 1000会导致npm无法执行
docker run -ti -w $(pwd) -v $(pwd):$(pwd) --rm registry.lazycat.cloud/lzc/lzc-api-protoc chown -R $UID:$UID ./

pushd lang/go/
go build
popd

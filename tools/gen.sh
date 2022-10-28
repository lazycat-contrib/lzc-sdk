#!/bin/bash
set -e

OUT_GO=$(pwd)/lang/go/
OUT_JS=$(pwd)/lang/js/
OUT_JAVA=$(pwd)/lang/java/lib/src/main/java/

JS_PLUGIN_PATH=$(pwd)/lang/js/node_modules/.bin/protoc-gen-ts_proto

if [[ ! -e $JS_PLUGIN_PATH ]]; then
    pushd lang/js/
    npm install
    popd
fi


rm -rf $OUT_GO/{common,localdevice} || true
rm -rf $OUT_JS/{common,localdevice,sys} || true
rm -rf $OUT_JAVA/ || true
mkdir -p $OUT_JAVA || true


pushd protos

protoc -I . \
       --go_out=$OUT_GO\
       --go-grpc_out=$OUT_GO\
       --ts_proto_out=$OUT_JS\
       --java_out=$OUT_JAVA \
       --grpc-java_out=$OUT_JAVA \
       --go_opt=paths=source_relative \
       --go-grpc_opt=paths=source_relative \
       --ts_proto_opt=outputClientImpl=grpc-web \
       --ts_proto_opt=exportCommonSymbols=false \
       --ts_proto_opt=esModuleInterop=true \
       --plugin=$JS_PLUGIN_PATH \
       --plugin=/usr/bin/protoc-gen-grpc-java \
       ./sys/*/*.proto \
       ./sys/*.proto \
       ./localdevice/*.proto \
       ./ssdp/*.proto \
       ./lanforward/*.proto \
       ./common/*.proto
popd


cp -r lang/java/lib/src java_api_workspace/


if [[ "$ENABLE_JAR" == "" ]]; then
    echo "Skip generate JAR, You can enable generating jar by set env 'ENABLE_JAR'"
else
    pushd java_api_workspace
    echo "java package start... "
    mvn -Dmaven.test.skip=true clean package
    echo "java package complete... "
    popd
fi



pushd lang/js/
npm run build
popd

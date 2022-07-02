#!/bin/bash

set -e

OUT_GO=$(pwd)/lang/go/
OUT_JS=$(pwd)/lang/js/
OUT_JAVA=$(pwd)/lang/java/

JS_PLUGIN_PATH=$(pwd)/lang/js/node_modules/.bin/protoc-gen-ts_proto

mkdir -p $OUT_GO $OUT_JS $OUT_JAVA

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
       ./localdevice/*/*.proto \
       ./common/*/*.proto
popd

# docker run -ti -w $(pwd) -v $(pwd):$(pwd) --rm registry.lazycat.cloud/lzc/lzc-api-protoc ./tools/gen.sh

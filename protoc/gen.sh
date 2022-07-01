#!/bin/bash

set -e

mkdir -p ./.cache/java >& /dev/null || true

protoc -I . \
       --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       --ts_proto_opt=outputClientImpl=grpc-web --ts_proto_opt=exportCommonSymbols=false --ts_proto_opt=esModuleInterop=true \
       --plugin=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=.\
       --plugin=/usr/bin/protoc-gen-grpc-java \
       --java_out=./.cache/java \
       --grpc-java_out=./.cache/java \
       */*.proto localdevice/*/*.proto

# jar -c -f cloud.lazycat.apis.jar -C ./.cache/java/ .
# pushd gohelper/
# go build
# popd

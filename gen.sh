#!/bin/bash

set -e


mkdir ./.cache >& /dev/null || true
mkdir ./java >& /dev/null || true

if [[ ! -f ./.cache/protoc-gen-grpc-java ]]; then
    wget https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-java/1.47.0/protoc-gen-grpc-java-1.47.0-linux-x86_64.exe -O ./.cache/protoc-gen-grpc-java
    chmod a+x ./.cache/protoc-gen-grpc-java
fi

protoc -I . \
       --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       --ts_proto_opt=outputClientImpl=grpc-web --ts_proto_opt=exportCommonSymbols=false --ts_proto_opt=esModuleInterop=true \
       --plugin=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=.\
       --plugin=./.cache/protoc-gen-grpc-java \
       --java_out=./java \
       --grpc-java_out=./java \
       */*.proto localdevice/*/*.proto

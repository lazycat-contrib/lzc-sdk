#!/bin/bash

PLUGIN_PATH=$(npm config get prefix)/bin/protoc-gen-ts_proto

protoc --ts_proto_out=js \
       --ts_proto_opt=outputClientImpl=grpc-web \
       --ts_proto_opt=exportCommonSymbols=false \
       --ts_proto_opt=esModuleInterop=true \
       --plugin=$PLUGIN_PATH \
       ./network_manager.proto 

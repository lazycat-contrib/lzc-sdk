#!/bin/bash
set -e

LZC_PLUGIN_PATH=$(pwd)/tools/protoc-gen-lzc/protoc-gen-lzc
OUT_LZC=$(pwd)/scontext/

mkdir -p $OUT_LZC || true

pushd protos

protoc -I . \
       --plugin=$LZC_PLUGIN_PATH \
       --lzc_out=$OUT_LZC \
       ./sys/*/*.proto \
       ./sys/*.proto \
       ./localdevice/*.proto \
       ./common/*.proto
popd

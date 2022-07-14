#!/bin/bash
set -e

docker run -ti -w $(pwd) -v $(pwd):$(pwd) --rm registry.lazycat.cloud/lzc/lzc-api-protoc ./tools/gen.sh

# 上面直接用-u 1000会导致npm无法执行
docker run -ti -w $(pwd) -v $(pwd):$(pwd) --rm registry.lazycat.cloud/lzc/lzc-api-protoc chown -R 1000:1000 ./

pushd lang/go/
go build
popd

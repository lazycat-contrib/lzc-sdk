#!/bin/bash

docker run -ti -w $(pwd) -v $(pwd):$(pwd) --rm registry.lazycat.cloud/lzc/lzc-api-protoc ./tools/gen.sh

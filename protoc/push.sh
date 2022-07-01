#!/bin/bash

IMAGE=registry.lazycat.cloud/lzc/lzc-api-protoc
docker build -t $IMAGE .
docker push $IMAGE

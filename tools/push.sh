#!/bin/bash

IMAGE=registry.corp.lazycat.cloud/lzc/lzc-api-protoc
docker build -t $IMAGE .
docker push $IMAGE

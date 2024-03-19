#!/bin/bash
set -e

# 作为不更新 docker image 的处理措施
npm config set registry https://registry.npmmirror.com/

PROTOC_ARGS=()

if [[ "$PROTOC_GEN_GO" -eq 1 ]]; then
	OUT_GO=$(pwd)/lang/go/
	rm -rf $OUT_GO/{common,localdevice} || true

	PROTOC_ARGS+=(
		--go_out=$OUT_GO
		--go-grpc_out=$OUT_GO
		--go_opt=paths=source_relative
		--go-grpc_opt=paths=source_relative
	)
fi

if [[ "$PROTOC_GEN_JS" -eq 1 ]]; then
	JS_PLUGIN_PATH=$(pwd)/lang/js/node_modules/.bin/protoc-gen-ts_proto
	if [[ ! -e $JS_PLUGIN_PATH ]]; then
		pushd lang/js/
		npm install
		popd
	fi

	OUT_JS=$(pwd)/lang/js/
	rm -rf $OUT_JS/{common,localdevice,sys} || true

	PROTOC_ARGS+=(
		--ts_proto_out=$OUT_JS
		--ts_proto_opt=outputClientImpl=grpc-web
		--ts_proto_opt=exportCommonSymbols=false
		--ts_proto_opt=esModuleInterop=true
		--ts_proto_opt=useAbortSignal=true
		--plugin=$JS_PLUGIN_PATH
	)
fi

if [[ "$PROTOC_GEN_JAVA" -eq 1 ]]; then
	OUT_JAVA=$(pwd)/lang/java/lib/src/main/java/
	rm -rf $OUT_JAVA/ || true
	mkdir -p $OUT_JAVA || true
	PROTOC_ARGS+=(
		--java_out=$OUT_JAVA
		--grpc-java_out=$OUT_JAVA
		--plugin=/usr/bin/protoc-gen-grpc-java
	)
fi


pushd protos

PROTOC_ARGS+=(
	./sys/*/*.proto
	./sys/*.proto
	./localdevice/*.proto
	./dlna/*.proto
	./common/*.proto
)

protoc -I . "${PROTOC_ARGS[@]}"

popd


if [[ "$PROTOC_GEN_JS" -eq 1 ]]; then
	pushd lang/js/
	npm run build
	popd
fi

if [[ "$PROTOC_GEN_JAVA" -eq 1 ]]; then
	cp -r lang/java/lib/src java_api_workspace/

	if [[ "$ENABLE_JAR" == "" ]]; then
		echo "Skip generate JAR, You can enable generating jar by set env 'ENABLE_JAR'"
	else
		pushd java_api_workspace
		echo "java package start... "
		mvn -Dmaven.test.skip=true -Dmaven.repo.local=/tmp/maven clean package
		echo "java package complete... "
		popd
	fi
fi

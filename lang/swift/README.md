<!--
 * @Author: Bin
 * @Date: 2022-08-09
 * @FilePath: /lzc-sdk/lang/swift/README.md
-->

# Cloud_Lazycat_Apis

> Swift Library for Cloud Lazycat Apis

## development 💽

```
# install Protocol (if you have installed, please skip)
wget -O protoc.zip  https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-osx-x86_64.zip  && unzip -d ./protoc/ -o protoc.zip && cp ./protoc/bin/protoc /usr/local/bin/

# use brew install swift-protobuf cil tools
brew install swift-protobuf
```

Protocol Compiler Installation: <https://github.com/protocolbuffers/protobuf#protocol-compiler-installation>

help docs: <https://github.com/apple/swift-protobuf/blob/main/Documentation/PLUGIN.md>

## build 🗜

```
# go to protos dir
cd protos/

# build
protoc -I . --swift_opt=Visibility=Public --swift_opt=ProtoPathModuleMappings=./swift_mappings.asciipb --swift_out=../lang/swift/Sources/ ./localdevice/*.proto ./sys/*.proto ./sys/*/*.proto ./common/*.proto
```

## resource 💾

Swift Protobuf: <https://github.com/apple/swift-protobuf>

Swift 下使用 Protobuf: <http://bartontang.github.io/2018/01/01/Swift%E4%B8%8B%E4%BD%BF%E7%94%A8Protobuf/>

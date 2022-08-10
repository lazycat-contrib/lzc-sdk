<!--
 * @Author: Bin
 * @Date: 2022-08-09
 * @FilePath: /lzc-apis-protos/lang/swift/README.md
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
protoc -I . --swift_opt=Visibility=Public --swift_out=/Volumes/tzmaxData/data/users/linakesi/lzc-apis-protos/lang/swift/Sources/ ./sys/*/*.proto ./sys/*.proto ./localdevice/*.proto ./common/*.proto
```

## resource 💾
Swift Protobuf: <https://github.com/apple/swift-protobuf>

Swift下使用Protobuf: <http://bartontang.github.io/2018/01/01/Swift%E4%B8%8B%E4%BD%BF%E7%94%A8Protobuf/>
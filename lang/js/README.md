<!--
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /lzc-sdk/lang/js/README.md
-->
# 懒猫云应用开发

> 懒猫云应用平台提供大量系统以及应用层的接口，用户或开发者可以基于开放的接口创造出有趣的应用程序，可以通过阅读本接口文档来帮助开发。

## 开始开发

@lazycatcloud/sdk 通过 npm 和 yarn 进行分发，你可以通过相应的包管理工具进行安装。

```
npm i @lazycatcloud/sdk

# or

yarn add @lazycatcloud/sdk
```

@lazycatcloud/sdk 目前由 `grpc api` 和 `lzc sdk` 两个部分组成。

- `grpc api` 是基于与懒猫云盒子以及其他组网设备数据以及交互所使用的接口。

- `lzc sdk` 是针对懒猫云客户端提供的云应用客户端程序、客户端原生应用系列交互 SDK。

我将通过下列两个示例帮助你分别理解 `grpc api` 和 `lzc sdk`

通过 `grpc api` 获取当前设备信息示例代码：
```
import { lzcApiGateway } from '@lazycatcloud/sdk/dist/localdevice/device';

const lzcApiGateway = new lzcAPIGateway();
const currentDevice = lzcAPI.current?.currentDevice;
console.log('currentDevice', currentDevice);
```

通过 `lzc sdk` 打开懒猫云相册应用（更多接口信息请参考 [懒猫云应用通用能力开发文档](./docs/extentions/AppCommon.md)）：
```
import { AppCommon } from '@lazycatcloud/sdk/dist/extentions';

AppCommon.LaunchApp("https://pbkbin.zmide.com", "com.zmide.bin");
```


## 更多文档

[懒猫云应用通用能力开发文档](./docs/extentions/AppCommon.md)

[实践指南（面向内部开发者）](./docs/contribute.md)

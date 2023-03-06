<!--
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /lzc-sdk/lang/js/docs/extentions/AppCommon.md
-->
# 懒猫云应用通用能力开发文档


## 打开应用能力接口文档

在部分场景下，用户在使用懒猫云应用的时候需要启动其他应用来完成任务。为此我们提供为懒猫云应用了打开应用能力以满足应用打开其他应用处理数据或完成任务的需求。打开应用能力面向开发者提供懒猫云应用唤起功能，通过打开应用能力开发者可安全便捷地使用懒猫云客户端或系统的能力，为懒猫云用户或设备使用者提供良好的用户体验。

### 使用步骤

1. 引入 SDK, 支持使用 CMD 标准模块加载方法加载
```
import { AppCommon } from '@lazycatcloud/sdk/dist/extentions';
```

2. 调用 `AppCommon.LaunchApp` 函数
```
AppCommon.LaunchApp("http://settings.pbkbin.heiyu.space/?path=/users/bin/download/", "pbkbin.heiyu.space");
```

### 调用参数

| 名称 | 必填 | 默认值 | 备注 |
| --- | --- | --- | --- |
| url | 是 | | 懒猫云应用 URL，可通过 URL Parameter 方式携带参数
| appid | 是 | | 懒猫云应用 appid 
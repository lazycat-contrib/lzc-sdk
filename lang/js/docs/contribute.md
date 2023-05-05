<!--
 * @Author: Bin
 * @Date: 2023-03-06
 * @FilePath: /lzc-sdk/lang/js/docs/contribute.md
-->

# 懒猫云客户端 SDK 部分贡献指南

> 该部分文档主要针对 SDK 开发者（内部）对接沟通使用。

## 准备开发

1. `lzc sdk` 开发目录 `./extentions`，目前 SDK 主要对象为 `LzcAppSdk`，该对象为一个单例。调用 `LzcAppSdk.getInstance()` 函数即可获得对象实例。

2. `LzcAppSdk` 内置工具函数如下 [更多函数参考 base.ts 代码](../extentions/base.ts):

   - isWebShell() 判断是否属于 webshell 环境
   - isAndroidWebShell() 是否是 android webshell 环境
   - isPCWebShell() 是否是 pc webshell 环境
   - isIosWebShell() 是否是 ios webshell 环境
   - isControlView() 是否属于 Control webshell 环境
   - isContentView() 是否属于 Content webshell 环境
   - isInApplication() 是否在客户端内

3. 通过 `LzcAppSdk` 获取原生（Android、iOS、PC）JSBridge 对象引用，由于部分平台（对，没错就是 iOS）特性获取 JSBridge 方式提供了异步和同步获取两种方式。（注: IOS 在调用 `useNative()` 获取到的 JSBridge 引用对象是缓存对象，如果未初始化完就调用可能会出现调用失败情况，所以类似 `GetValue()` 这种需要返回值的实现请一定一定尽量使用 `useNativeAsync()` 来获取 JSBridge 引用对象。）

   - useNativeAsync(namespaces?: any)
   - useNative(namespaces?: any)

4. `@native` 装饰器，为了避免跨平台调用函数导致页面异常问题，请实现函数时尽量使用该装饰器声明其支持的平台，对于不支持的平台该装饰器会自动处理替换为空函数实现。平台类型参数如下

   - LzcAppPlatformType.Android
   - LzcAppPlatformType.IOS
   - LzcAppPlatformType.PC
   - LzcAppPlatformType.Browser

5. 关于如何实现一个支持全平台的 SDK API 最佳实践请参考 [app_common.ts](../extentions/app_common.ts)

6. `LzcAppSdk` 提供回调函数处理实现请参考 [关于如何实现回调函数](#关于如何实现回调函数)

## 关于如何实现回调函数

文档整理中（🕊️）…

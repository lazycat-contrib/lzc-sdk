import LzcAppSdk, { LzcAppPlatformType, LzcAppSdkManage, native } from "./base"
import { IntentAction, intentActionFromJSON } from "../common/file_handler"
import { ClientAuthorizationBaseStatus, ClientAuthorizationType } from './client_authorization';

class AppCommon extends LzcAppSdkManage {
    /**
     * @description: 打开链接工具函数为了兼容 Safari 浏览器默认打开阻止弹出式窗口问题
     * @param {string} url
     * @return {*}
     */
    private static _linkOpen(url: string): void {
        // const item = document.createElement('a');
        // item.href = url;
        // item.target = '_blank';
        // item.click();
        // item.remove();
        setTimeout(() => window.open(url, "_blank"))
    }

    /**
     * @description: IOS 测试函数
     * @deprecated: 该函数仅用于 IOS 测试
     * @return {*}
    */
    @native(LzcAppPlatformType.IOS)
    public static async ScriptHandlers(): Promise<string[]> {
        // const callback = await LzcAppSdk.useNative.ScriptHandlers()
        // console.log("LzcAppSdk.useNative", LzcAppSdk.useNative);
        const jsBridge = await LzcAppSdk.useNativeAsync()
        const call = await jsBridge.ScriptHandlers()
        // jsBridge.Test()
        return call
    }

    /**
     * @description: 打开应用
     * @param {string} url
     * @param {string} appid
     * @return {Promise<void>}
     */
    @native(LzcAppPlatformType.IOS, LzcAppPlatformType.Android, LzcAppPlatformType.PC, LzcAppPlatformType.Browser, LzcAppPlatformType.TvOs)
    public static async LaunchApp(url: string, appid: string): Promise<void> {
        // 判断是否在浏览器中
        if (!LzcAppSdk.isInApplication()) {
            // window.open(url, '_blank')
            AppCommon._linkOpen(url)
            return
        }

        // 在 Android 环境中
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
            const openApprequest = {
                appid: appid,
                url: url,
            }
            var openAppRequestJsonStr = JSON.stringify(openApprequest)
            jsBridge.LaunchApp(openAppRequestJsonStr)
            return
        }

        if (LzcAppSdk.isIosWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync()
            jsBridge.LaunchApp(url, appid)
            return
        }
        // Electron launch app
        if (LzcAppSdk.isPCWebShell()) {
            window.ipcRenderer.invoke("launchApp", url, appid)
            return
        }

        // tvos打开应用
        if (LzcAppSdk.isTvOsWebShell()) {
            // @ts-ignore
            // mipcr支持的类型: https://gitee.com/linakesi/lzc-tvos/blob/master/client/src/preload.ts#L19
            window.mipcr.LaunchApp(url, appid)
        }
    }

    /**
     * @description: 打开原生应用
     * @param {string} filepath
     * @param {string} appid
     * @param {IntentAction} intentAction
     * @return {*}
     */
    @native(LzcAppPlatformType.Android)
    public static async LaunchNativeApp(filepath: string, appid: string, intentAction: IntentAction = IntentAction.UN_KNOWN): Promise<void> {
        // 在浏览器环境中
        if (!LzcAppSdk.isInApplication()) {
            console.error("LaunchNativeApp 方法暂未实现。")
            return
        }
        // 在 Android 环境
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
            const openApprequest = {
                appid: appid,
                filepath: filepath,
                intentAction: intentActionFromJSON(intentAction),
            }
            var requestStr = JSON.stringify(openApprequest)
            jsBridge.LaunchNativeApp(requestStr)
            return
        }
        // TODO: 其他环境中待实现
        console.error("LaunchNativeApp 方法暂未实现。")
    }

    /**
     * @description: 使用指定的原生应用打开在线文件
     * @param boxName 盒子名称
     * @param {string} webdav_file_url webdav 文件地址
     * @param {IntentAction} intentAction
     * @param {string} appid 原生应用 id
     * @return {*}
     */
    @native(LzcAppPlatformType.Android)
    public static async ShareWith(boxName: string, webdav_url: string, appid: string): Promise<void> {
        // 在浏览器环境中
        if (!LzcAppSdk.isInApplication()) {
            console.error("LaunchNativeApp 方法暂未实现。")
            return
        }
        // 在 Android 环境
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
            const openApprequest = {
                webdavUrl: webdav_url,
                boxName: boxName,
                appid: appid,
            }
            var requestStr = JSON.stringify(openApprequest)
            jsBridge.ShareWith(requestStr)
            return
        }
        // TODO: 其他环境中待实现
        console.error("LaunchNativeApp 方法暂未实现。")
    }

    @native(LzcAppPlatformType.Android)
    public static async OpenWith(boxName: string, webdav_url: string, appid: string): Promise<void> {
        // 在浏览器环境中
        if (!LzcAppSdk.isInApplication()) {
            console.error("LaunchNativeApp 方法暂未实现。")
            return
        }
        // 在 Android 环境
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
            const openApprequest = {
                webdavUrl: webdav_url,
                boxName: boxName,
                appid: appid,
            }
            var requestStr = JSON.stringify(openApprequest)
            jsBridge.OpenWith(requestStr)
            return
        }
        // TODO: 其他环境中待实现
        console.error("LaunchNativeApp 方法暂未实现。")
    }

    /**
     *
     * @param options
     * actionName: 安卓必须填写 actionName 由packageName和classname组成, ios可以忽略
     * ids/id 安卓或者ios 提供的文件id
     * @constructor
     */
    @native(LzcAppPlatformType.Android, LzcAppPlatformType.IOS)
    public static async ShareMedia(options: {
        actionName?: string,
        ids?: string[], id?: string
    }): Promise<void> {
        // 在浏览器环境中
        if (!LzcAppSdk.isInApplication()) {
            console.error("LaunchNativeApp 方法暂未实现。")
            return
        }
        const idArray = []
        if (options.ids && options.ids.length > 0) {
            idArray.push(...options.ids)
        }
        if (options.id && options.id.length > 0) {
            idArray.push(options.id)
        }
        if (idArray.length == 0) {
            throw Error("至少有一个id参数")
        }
        if (LzcAppSdk.isAndroidWebShell()) {
            if (options.actionName.length == 0 || options.actionName.indexOf(":") < 0) {
                throw Error("必须填写actionName")
            }
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
            const openApprequest = {
                appid: options.actionName,
                ids: idArray.join(",")
            }
            var requestStr = JSON.stringify(openApprequest)
            jsBridge.ShareMedia(requestStr)
            return
        }
    }

    /**
     * @description: 获取客户端授权状态
     * @return {*}
     */
    @native(LzcAppPlatformType.IOS)
    public static async GetClientAuthorizationStatus(type: ClientAuthorizationType | string): Promise<ClientAuthorizationBaseStatus> {
        // 判断当前权限是否受支持
        switch (type) {
            case ClientAuthorizationType.PhotoLibrary:
                break;
            default:
                console.warn(`the current lzc-sdk version does not support the ${type} permission type`);
                return ClientAuthorizationBaseStatus.Unknown
        }

        // 调用各客户端具体实现
        if (LzcAppSdk.isIosWebShell()) {
            // 调用客户端 GetClientAuthorizationStatus 函数
            const jsBridge = await LzcAppSdk.useNativeAsync()
            return jsBridge.GetClientAuthorizationStatus(type)
        }
        return ClientAuthorizationBaseStatus.Unknown
    }

    /**
     * @description: 申请客户端授权
     * @return {*}
     */
    @native(LzcAppPlatformType.IOS)
    public static async RequestClientAuthorization(type: ClientAuthorizationType | string): Promise<void> {
        // 判断当前权限是否受支持
        switch (type) {
            case ClientAuthorizationType.PhotoLibrary:
                break;
            default:
                console.warn(`the current lzc-sdk version does not support the ${type} permission type`);
        }

        // 调用各客户端具体实现
        if (LzcAppSdk.isIosWebShell()) {
            // 调用客户端 RequestClientAuthorization 函数
            const jsBridge = await LzcAppSdk.useNativeAsync()
            jsBridge.RequestClientAuthorization(type)
            return
        }

        console.warn('the current client does not support');
        
    }
}

export * from './client_authorization'
export { AppCommon }

import LzcAppSdk, { LzcAppPlatformType, LzcAppSdkManage, native } from "./base"
import { IntentAction, intentActionFromJSON } from "../common/file_handler"
import { ClientAuthorizationBaseStatus, ClientAuthorizationType } from './client_authorization';
import { VibrateType } from './vibrate_type';

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
            // mipcr支持的类型: https://gitee.com/linakesi/lzc-tvos/blob/master/client/src/preload.ts#L19
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
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
     * @deprecated 该接口已弃用，预计后续版本可能会被删除，请使用 `ShareWithFiles` 接口替代
     * @description: 使用指定的原生应用打开在线文件
     * @param boxName 盒子名称
     * @param {string} path 文件路径
     * @param {string} appid 原生应用 id (在 IOS 中此参数传空字符串)
     * @return {*}
     */
    @native(LzcAppPlatformType.Android, LzcAppPlatformType.IOS)
    public static async ShareWith(boxName: string, path: string, appid: string): Promise<void> {
        // 在浏览器环境中
        if (!LzcAppSdk.isInApplication()) {
            console.error("ShareWith 方法暂未实现。")
            return
        }
        // 在 Android 环境
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
            const openApprequest = {
                webdavUrl: path,
                boxName: boxName,
                appid: appid,
            }
            var requestStr = JSON.stringify(openApprequest)
            jsBridge.ShareWith(requestStr)
            return
        }

        // 在 IOS 中仅需要 path 参数
        if (LzcAppSdk.isIosWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync()
            await jsBridge.ShareWith(path)
            return
        }

        // TODO: 其他环境中待实现
        console.error("ShareWith 方法暂未实现。")
    }

    /**
     * @description: 分享本地离线文件或媒体资源
     * @param {string | undefined} path 文件路径
     * @param {string[] | undefined} paths 分享多文件路径
     * @param {string[] | undefined} mediaIds 分享多文件路径
     * @return {boolean}
     */
    @native(LzcAppPlatformType.Android, LzcAppPlatformType.IOS)
    public static async ShareWithFiles(path: string = undefined, paths: string[] = undefined, mediaIds: string[] = undefined): Promise<boolean> {
        // 在浏览器环境中
        if (!LzcAppSdk.isInApplication()) {
            console.error("ShareWithFiles 方法暂未实现。")
            return false
        }

        // 判断参数
        if (path == undefined && paths == undefined && mediaIds == undefined) {
            throw Error("path、paths 和 mediaIds 参数不能同时为空")
        }

        // 在 Android 环境
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
            const content = {
                "filePath": [],
                "mediaIds": [],
            }
            if (path != undefined) {
                content.filePath.push(path)
            }
            if (paths != undefined) {
                content.filePath = content.filePath.concat(paths)
            }
            if (mediaIds != undefined) {
                content.mediaIds = content.mediaIds.concat(mediaIds)
            }
            jsBridge.ShareWithFiles(JSON.stringify(content));
            return true
        }

        // 在 IOS 环境
        if (LzcAppSdk.isIosWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync()
            return await jsBridge.ShareWith(path, paths, mediaIds)
        }

        // TODO: 其他环境中待实现
        console.error("ShareWithFiles 方法暂未实现。")
        return false
    }

    /**
     * @description: 
     * @param {string} boxName 盒子名称
     * @param {string} path 文件路径
     * @param {string} appid 原生应用 id (在 IOS 中此参数传空字符串)
     * @return {*}
     */
    @native(LzcAppPlatformType.Android, LzcAppPlatformType.IOS)
    public static async OpenWith(boxName: string, path: string, appid: string): Promise<void> {
        // 在浏览器环境中
        if (!LzcAppSdk.isInApplication()) {
            console.error("OpenWith 方法暂未实现。")
            return
        }
        // 在 Android 环境
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
            const openApprequest = {
                webdavUrl: path,
                boxName: boxName,
                appid: appid,
            }
            var requestStr = JSON.stringify(openApprequest)
            jsBridge.OpenWith(requestStr)
            return
        }

        // 在 IOS 中仅需要 path 参数
        if (LzcAppSdk.isIosWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync()
            await jsBridge.OpenWith(path)
            return
        }

        // TODO: 其他环境中待实现
        console.error("OpenWith 方法暂未实现。")
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
        ids?: string[],
        id?: string
    }): Promise<void> {
        // 在浏览器环境中
        if (!LzcAppSdk.isInApplication()) {
            console.error("LaunchNativeApp 方法暂未实现。")
            return
        }

        if (LzcAppSdk.isAndroidWebShell()) {
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
            return
        }

        if (LzcAppSdk.isIosWebShell()) {
            const { ids: _ids = [] } = options
            let ids = [..._ids];
            if (options.id) {
                ids.push(options.id)
            }
            const jsBridge = await LzcAppSdk.useNative()
            jsBridge.ShareMedia(ids)
            return
        }

        // TODO: 其他环境中待实现
        console.error("Not supported ShareMedia on this platform")
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
    /**
     * @description: 使手机发起振动
     * @return {*}
     */
    @native(LzcAppPlatformType.IOS,LzcAppPlatformType.Android)
    public static async Vibrate(type: VibrateType): Promise<void> {
        // 调用各客户端具体实现
        if (LzcAppSdk.isIosWebShell()) {
            // TODO
            return
        }
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(lzc_vibrate)
            jsBridge.Vibrate(type)
            return
        }
        console.warn('the current client does not support');
    }
    /**
 * @description: 发起扫码请求
 * @return {Promise<string>} - 扫码结果的 Promise（字符串）
 */
@native(LzcAppPlatformType.IOS, LzcAppPlatformType.Android)
public static async ScanQrCode(): Promise<string> {
    if (LzcAppSdk.isIosWebShell()) {
        try {
            const jsBridge = await LzcAppSdk.useNativeAsync();
            const result = await new Promise<string>((resolve, reject) => {
                jsBridge.ScanQrCode((scanResult: string) => {
                    resolve(scanResult); // 将扫码结果作为 Promise 的解析值
                });
            });
            return result;
        } catch (error) {
            console.error('Error scanning QR code on iOS:', error);
            return 'Failed to scan QR code'; // 返回错误信息
        }
    }

    if (LzcAppSdk.isAndroidWebShell()) {
        try {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service);
            const result = await jsBridge.ScanQrCode(); // 获取结果
            return result;
        } catch (error) {
            console.error('Error scanning QR code on Android:', error);
            return 'Failed to scan QR code'; // 返回错误信息
        }
    }

    console.warn('The current client does not support');
    return 'The current client does not support'; // 返回不支持的错误信息
}

}

export * from './client_authorization'
export { AppCommon }

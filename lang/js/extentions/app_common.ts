import LzcAppSdk, { LzcAppPlatformType, LzcAppSdkManage, native } from './base';

class AppCommon extends LzcAppSdkManage {

    @native(LzcAppPlatformType.IOS)
    /**
     * @description: IOS 测试函数
     * @deprecated: 该函数仅用于 IOS 测试
     * @return {*}
     */
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
    @native(LzcAppPlatformType.IOS,
        LzcAppPlatformType.Android,
        LzcAppPlatformType.PC,
        LzcAppPlatformType.Browser
    )
    public static async LaunchApp(url: string, appid: string): Promise<void> {
        // 判断是否在浏览器中
        if (!LzcAppSdk.isInApplication()) {
            window.open(url, '_blank')
            return
        }

        // 在 Android 环境中
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
            const openApprequest = {
                "appid": appid,
                "url": url,
            }
            var openAppRequestJsonStr = JSON.stringify(openApprequest);
            jsBridge.LaunchApp(openAppRequestJsonStr)
            return
        }

        if (LzcAppSdk.isIosWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync()
            jsBridge.LaunchApp(url, appid)
            return
        }
        // Electron launch app
        if(LzcAppSdk.isPCWebShell()){
	    window.ipcRenderer.invoke("launchApp", url, appid)
	    return
        }
    }

    @native(
        LzcAppPlatformType.Android,
    )
    /**
     * @description: 打开原生应用
     * @param {string} filepath
     * @param {string} appid
     * @return {*}
     */
    public static async LaunchNativeApp(filepath: string, appid: string): Promise<void> {
        // 在浏览器环境中
        if (!LzcAppSdk.isInApplication()) {
            console.error("LaunchNativeApp 方法暂未实现。")
            return
        }

        // 在 Android 环境
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
            const openApprequest = {
                "appid": appid,
                "filepath": filepath,
            }
            var requestStr = JSON.stringify(openApprequest);
            jsBridge.LaunchNativeApp(requestStr)
        }

        // TODO: 其他环境中待实现
        console.error("LaunchNativeApp 方法暂未实现。")
    }

}

export { AppCommon }

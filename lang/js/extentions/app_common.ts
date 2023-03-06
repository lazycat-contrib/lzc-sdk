import LzcAppSdk, { LzcAppPlatformType, LzcAppSdkManage, native } from './base';

class AppCommon extends LzcAppSdkManage {

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
    @native(LzcAppPlatformType.IOS,
        LzcAppPlatformType.Android,
        LzcAppPlatformType.PC,
        LzcAppPlatformType.Browser
    )
    public static async LaunchApp(url: string, appid: string): Promise<void> {
        let browserOnly = false // 读取用户是否设置为在浏览器打开

        if (!LzcAppSdk.isInApplication() || browserOnly) {
            // 判断是否在浏览器中，在浏览器中时或者用户设置为在浏览器打开时，将直接使用 window.open 打开页面
            window.open(url, '_blank')
        } else {
            // 判断是在客户端中
            const jsBridge = await LzcAppSdk.useNativeAsync()
            jsBridge.LaunchApp(url, appid)
        }
    }

    public static async LaunchNativeApp() {
        // TODO: 待实现
        const jsBridge = LzcAppSdk.useNative()
        console.error("LaunchNativeApp 方法暂未实现。")
    }

}

export { AppCommon }

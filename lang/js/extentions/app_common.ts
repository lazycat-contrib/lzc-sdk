import LzcAppSdk, { LzcAppPlatformType, LzcAppSdkManage, native } from './base';


async function launchNativeApp(appid: string, filepath: string) {
    const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
    const openApprequest = {
        "appid": appid,
        "filepath": filepath,
    }
    var requestStr = JSON.stringify(openApprequest);
    jsBridge.LaunchNativeApp(requestStr)
}

async function launchApp(appid: string, url: string) {
    const jsBridge = await LzcAppSdk.useNativeAsync(android_launch_service)
    const openApprequest = {
        "appid": appid,
        "url": url,
    }
    var openAppRequestJsonStr = JSON.stringify(openApprequest);
    jsBridge.LaunchApp(openAppRequestJsonStr)
}

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
        if (LzcAppSdk.isInApplication()) {
            // 判断是否在浏览器中，在浏览器中时或者用户设置为在浏览器打开时，将直接使用 window.open 打开页面
            if (LzcAppSdk.isAndroidWebShell()) {
                await launchApp(appid, url);
            } else {
                window.open(url, '_blank')
            }
        } else {
            if(LzcAppSdk.isAndroidWebShell()){
                await launchApp(appid, url);
            }else{
                const jsBridge = await LzcAppSdk.useNativeAsync()
                jsBridge.LaunchApp(url, appid)
            }
        }
    }

    public static async LaunchNativeApp(filepath: string,appid: string) {
        // TODO: 待实现
        if(LzcAppSdk.isInApplication()){
            if(LzcAppSdk.isAndroidWebShell()){
                await launchNativeApp(appid, filepath);
            }else{
                console.error("LaunchNativeApp 方法暂未实现。")
            }
        }else{
            if(LzcAppSdk.isAndroidWebShell()){
                await launchNativeApp(appid, filepath);
            }else{
                console.error("LaunchNativeApp 方法暂未实现。")
            }
        }
    }

}

export { AppCommon }

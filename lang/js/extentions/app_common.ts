import LzcAppSdk, { LzcAppPlatformType, LzcAppSdkManage, native } from "./base"

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
  @native(LzcAppPlatformType.IOS, LzcAppPlatformType.Android, LzcAppPlatformType.PC, LzcAppPlatformType.Browser)
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
  }

  /**
   * @description: 打开原生应用
   * @param {string} filepath
   * @param {string} appid
   * @return {*}
   */
  @native(LzcAppPlatformType.Android)
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
        appid: appid,
        filepath: filepath,
      }
      var requestStr = JSON.stringify(openApprequest)
      jsBridge.LaunchNativeApp(requestStr)
    }

    // TODO: 其他环境中待实现
    console.error("LaunchNativeApp 方法暂未实现。")
  }
}

export { AppCommon }

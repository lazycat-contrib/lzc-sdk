import App, { LzcClient } from "./base"
import LzcAppSdk from "./base"

// android懒猫云商城的第三方应用在打包后，再次激活该窗口的时候，会发送事件。 页面可以通过事件传来的地址，

// 跟随系统默认
const SCREEN_ORIENTATION_UNSPECIFIED = 1
// 强制横屏
const SCREEN_ORIENTATION_LANDSCAPE = 2
// 强制竖屏
const SCREEN_ORIENTATION_PORTRAIT = 3

export const SAFE = {
  // h5调用webview 打开App
  OpenApp: function (url, isFullScreen) {
    if (App.isWebShell()) {
      return
    }
    LzcClient.OpenApp(url, isFullScreen)
  },
  // 目前只支持android
  // 设置指定key 对应的value
  SetValue: function (key, value) {
    LzcClient.SetValue(key, value)
  },
  // 获取指定key的value
  // 目前只支持android
  GetValue: function (key) {
    return LzcClient.GetValue(key)
  },
  // 目前只支持android
  // 方向请使用常量
  SetScreenOrientation: function (ori) {
    LzcClient.SetScreenOrientation(ori)
  },
  // 最大255的亮度
  SetScreentBrightness(number) {
    LzcClient.SetScreentBrightness(number)
  },
  // 获取当前屏幕亮度百分比(0~1)
  GetScreenBrightNess() {
    return LzcClient.GetScreenBrightNess()
  },
  // 获取当前音量
  GetCurrentVolume() {
    return LzcClient.GetCurrentVolume()
  },
  // 设置指定音量
  SetVoice(num) {
    LzcClient.SetVoice(num)
  },
  // 增加音量
  AddVoice() {
    LzcClient.AddVoice()
  },
  SubVoice() {
    LzcClient.SubVoice()
  },
  CloseWindow() {
    LzcClient.CloseWindow
  },
  Minimize() {
    LzcClient.Minimize()
  },
  /**
   * // 仅仅android 实现
   * @param dimEnable 是否开启dialog遮罩透明
   * @param backgroundDimAmount 透明度 0~1
   * @param url dialog显示的地址
   * @param callback 回调的callback(),接受结果为CloseDialog中的result 参数
   * @constructor
   */
  OpenDialog(dimEnable: boolean, backgroundDimAmount: number, url: string, callback: string) {
    const openApprequest = {
      dimEnable: dimEnable,
      backgroundDimAmount: backgroundDimAmount,
      url: url,
    }
    var openAppRequestJsonStr = JSON.stringify(openApprequest)
    const jsBridge = LzcAppSdk.useNative(android_dialog)
    jsBridge.OpenDialog(openAppRequestJsonStr, callback)
  },
  /**
   * 仅仅在打开的dialog中可以使用
   * @param result 关闭dialog 的回调结果
   * @constructor
   */
  CloseDialog(result: string) {
    const jsBridge = LzcAppSdk.useNative(android_dialog)
    jsBridge.Close(result)
  },
  InstallApk(url: string) {
    const jsBridge = LzcAppSdk.useNative(lzc_app_manager)
    jsBridge.InstallClient(url)
  }
}

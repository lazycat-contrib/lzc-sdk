import App, { LzcClient } from './base';

// 横屏
const SCREEN_ORIENTATION_SENSOR_PORTRAIT = 1
// 横向屏幕
const SCREEN_ORIENTATION_SENSOR_LANDSCAPE = 2

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
        LzcClient.SetScreentBrightness(number);
    },
    // 获取当前屏幕亮度百分比(0~1)
    GetScreenBrightNess() {
        return LzcClient.GetScreenBrightNess()
    },
    // 获取当前音量
    GetCurrentVolume() {
        return LzcClient.GetCurrentVolume();
    },
    // 设置指定音量
    SetVoice(num) {
        LzcClient.SetVoice(num);
    },
    // 增加音量
    AddVoice() {
        LzcClient.AddVoice()
    },
    SubVoice() {
        LzcClient.SubVoice();
    },
    CloseWindow() {
        LzcClient.CloseWindow
    },
    Minimize() {
        LzcClient.Minimize();
    }
}

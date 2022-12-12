// 横屏
const SCREEN_ORIENTATION_SENSOR_PORTRAIT = 1
// 横向屏幕
const SCREEN_ORIENTATION_SENSOR_LANDSCAPE = 2

const SAFE = {
// h5调用webview 打开App
    OpenApp: function (url, isFullScreen) {
        if (isWebShell()) {
            return
        }
        view.OpenApp(url, isFullScreen)
    },
    // 目前只支持android
    // 设置指定key 对应的value
    SetValue: function (key, value) {
        view.SetValue(key, value)
    },
    // 获取指定key的value
    // 目前只支持android
    GetValue: function (key) {
        return view.GetValue(key)
    },
    // 目前只支持android
    // 方向请使用常量
    SetScreenOrientation: function (ori) {
        view.SetScreenOrientation(ori)
    },
    // 最大255的亮度
    SetScreentBrightness(number){
        view.SetScreentBrightness(number);
    },
    // 获取当前屏幕亮度百分比(0~1)
    GetScreenBrightNess(callBackMethod: String) {
        view.GetScreenBrightNess(callBackMethod)
    },
    // 获取当前音量
    GetCurrentVolume(callBackMethod: String) {
        view.GetCurrentVolume(callBackMethod);
    },
    // 设置指定音量
    SetVoice(num) {
        view.SetVoice(num);
    },
    // 增加音量
    AddVoice() {
        view.AddVoice()
    },
    SubVoice() {
        view.SubVoice();
    }
}

// 横屏
const SCREEN_ORIENTATION_SENSOR_PORTRAIT = 1
// 横向屏幕
const SCREEN_ORIENTATION_SENSOR_LANDSCAPE  = 2

const SAFE = {
// h5调用webview 打开App
    OpenApp: function (url,isFullScreen) {
        if (isWebShell()) {
            return
        }
        view.OpenApp(url,isFullScreen)
    },
    // 目前只支持android
    // 设置指定key 对应的value
    SetValue: function(key,value){
        view.SetValue(key,value)
    },
    // 获取指定key的value
    // 目前只支持android
    GetValue: function(key){
        return view.GetValue(key)
    },
    // 目前只支持android
    // 方向请使用常量
    SetScreenOrientation: function(ori){
        view.SetScreenOrientation(ori)
    }
}

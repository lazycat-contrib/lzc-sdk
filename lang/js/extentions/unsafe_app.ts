import App, { LzcClient } from './base';
import Extentions from "./index";
import LzcAppSdk from "./base";

declare global {
    const android_ad_blocker: any
}

const UNSAFE = {
    // h5调用webview 打开App
    OpenApp: function (url, isFullScreen) {
        if (!App.isWebShell()) {
            return
        }
        LzcClient.OpenApp(url, isFullScreen)
    },
    // contentView 打开指定url
    SetContentURL: function (url) {
        if (!App.isControlView()) {
            return
        }
        LzcClient.SetContentURL(url)
    },
    // 获取contentView 的当前url地址
    GetContentURL:
        function (callback) {
            if (!App.isControlView()) {
                return
            }
            return LzcClient.GetContentURL(callback)
        },
    // contentView 回退
    BackContentView:
        function () {
            if (!App.isControlView()) {
                return
            }
            LzcClient.BackContentView()
        },
    // contentView 前进
    ForwardContentView: function () {
        if (!App.isControlView()) {
            return
        }
        LzcClient.ForwardContentView();
    },
    // 在contentView 注入jsContent, 并且提供controlView 页面的js的 callbackMethod方法
    InjectJS: function (jsContent, callbackMethod) {
        if (!App.isControlView()) {
            return
        }
        LzcClient.InjectJS(jsContent, callbackMethod);
    },
    // 获取contentView 下面指定domain的cookie
    ReadCookie(domain) {
        if (!App.isControlView()) {
            return
        }
        return LzcClient.ReadCookie(domain)
    },
    // 在controlView 中执行jsContent(用于contentView通知 controlView)
    NotifyControView: function (jsContent) {
        if (!App.isContentView()) {
            return
        }
        LzcClient.NotifyControView(jsContent);
    },
    // 修改contentView 的useragent
    // 目前只支持android
    UpdateContentViewUserAgent: function (ua) {
        LzcClient.UpdateContentViewUserAgent(ua)
    },
    // contentView 执行
    // 目前只支持android
    XmlHttpRequest4ContentView: function (reqJsonStr) {
        return LzcClient.XmlHttpRequest4ContentView(reqJsonStr)
    },
    //切换controlView 显示状态
    // 目前只支持android
    ToggleControlView: function () {
        return LzcClient.ToggleControlView();
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
    // callbackMethod: 结果回调方法
    CanGoback: function (callbackMethod) {
        return LzcClient.CanGoback(callbackMethod);
    },
    // controlViewHeight 两种使用方式,第一种是使用pixel: 100px ,另外一种使用百分比: 20%,
    SetControlViewHeight: function (controlViewHeight) {
        LzcClient.SetControlViewHeight(controlViewHeight);
    },
    ImportAdHosts: async function (hostUrl) {
        if (LzcAppSdk.isAndroidWebShell()) {
            const jsBridge = await LzcAppSdk.useNativeAsync(android_ad_blocker)
            jsBridge.ImportAdHosts(hostUrl)
        }
    },
    /**
     * // 仅仅android 实现
     * @param dimEnable 是否开启dialog遮罩透明
     * @param backgroundDimAmount 透明度 0~1
     * @param url dialog显示的地址
     * @param callback 回调的callback(),接受结果为CloseDialog中的result 参数
     * @constructor
     */
    OpenDialog(dimEnable: boolean,
               backgroundDimAmount: number,
               url :string,
               callback:string,
    ) {
        const openApprequest = {
            dimEnable: dimEnable,
            backgroundDimAmount: backgroundDimAmount,
            url: url
        }
        var openAppRequestJsonStr = JSON.stringify(openApprequest);
        const jsBridge = LzcAppSdk.useNative(android_dialog)
        jsBridge.OpenDialog(openAppRequestJsonStr,callback)
    },
    /**
     * 仅仅在打开的dialog中可以使用
     * @param result 关闭dialog 的回调结果
     * @constructor
     */
    CloseDialog(result: string) {
        const jsBridge = LzcAppSdk.useNative(android_dialog)
        jsBridge.Close(result);
    }

};

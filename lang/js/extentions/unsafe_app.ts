var view;
function isWebShell() {
    return navigator.userAgent.indexOf("Lazycat") != -1 && !isControlView() && !isContentView();
}
// 是否是android webshell 环境
function isAndroidWebShell() {
    return navigator.userAgent.indexOf("Lazycat_101") != -1;
}

// 是否是pc webshell 环境
function isPCWebShell() {
    return navigator.userAgent.indexOf("Lazycat_102") != -1;
}

// 是否是ios webshell 环境
function isIosWebShell() {
    return navigator.userAgent.indexOf("Lazycat_103") != -1;
}


function isControlView(){
    return navigator.userAgent.indexOf("Lazycat_ControlView") != -1;
}

function isContentView(){
    return navigator.userAgent.indexOf("Lazycat_ContentView") != -1;
}


if (isAndroidWebShell()) {
    // @ts-ignore
    view = android;
} else if (isPCWebShell()) {
    // @ts-ignore
    view = window.electronAPI
}


const UNSAFE = {
// h5调用webview 打开App
    OpenApp: function (url) {
        if (!isWebShell()) {
            return
        }
        view.OpenApp(url)
    },
    // contentView 打开指定url
    SetContentURL: function (url) {
        if (!isControlView()) {
            return
        }
        view.SetContentURL(url)
    },
    // 获取contentView 的当前url地址
    GetContentURL:
        function (callback) {
            if (!isControlView()) {
                return
            }
            return view.GetContentURL(callback)
        },
    // contentView 回退
    BackContentView:
        function () {
            if (!isControlView()) {
                return
            }
            view.BackContentView()
        },
    // contentView 前进
    ForwardContentView: function () {
        if (!isControlView()) {
            return
        }
        view.ForwardContentView();
    },
    // 在contentView 注入jsContent, 并且提供controlView 页面的js的 callbackMethod方法
    InjectJS: function (jsContent,callbackMethod){
        if (!isControlView()) {
            return
        }
        view.InjectJS(jsContent, callbackMethod);
    },
    // 获取contentView 下面指定domain的cookie
    ReadCookie(domain){
        if(!isControlView()){
            return
        }
        return  view.ReadCookie(domain)
    },
    // 在controlView 中执行jsContent(用于contentView通知 controlView)
    NotifyControView: function (jsContent) {
        if (!isContentView()) {
            return
        }
        view.NotifyControView(jsContent);
    },
    // 修改contentView 的useragent
    UpdateContentViewUserAgent: function (ua) {
        view.UpdateContentViewUserAgent(ua)
    },
    // contentView 执行
    XmlHttpRequest4ContentView: function (reqJsonStr){
        return view.XmlHttpRequest4ContentView(reqJsonStr)
    },
    //切换controlView 显示状态
    ToggleControlView: function (){
        return view.ToggleControlView();
    }
};


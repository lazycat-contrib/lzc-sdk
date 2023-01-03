const UNSAFE = {
// h5调用webview 打开App
    OpenApp: function (url,isFullScreen) {
        if (!isWebShell()) {
            return
        }
        view.OpenApp(url,isFullScreen)
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
    // 目前只支持android
    UpdateContentViewUserAgent: function (ua) {
        view.UpdateContentViewUserAgent(ua)
    },
    // contentView 执行
    // 目前只支持android
    XmlHttpRequest4ContentView: function (reqJsonStr){
        return view.XmlHttpRequest4ContentView(reqJsonStr)
    },
    //切换controlView 显示状态
    // 目前只支持android
    ToggleControlView: function (){
        return view.ToggleControlView();
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
    // callbackMethod: 结果回调方法
    CanGoback: function(callbackMethod){
       return view.CanGoback(callbackMethod);
    },
    // controlViewHeight 两种使用方式,第一种是使用pixel: 100px ,另外一种使用百分比: 20%,
    SetControlViewHeight:function (controlViewHeight){
        view.SetControlViewHeight(controlViewHeight);
    },
};


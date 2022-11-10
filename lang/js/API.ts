// 是否是webshell 环境
export function isWebShell() {
    return navigator.userAgent.indexOf("Lazycat") != -1;
}
// 是否是pc webshell 环境
export function isPCWebShell() {
    return navigator.userAgent.indexOf("Lazycat_PC") != -1;
}
// 是否是android webshell 环境
export function isAndroidWebShell() {
    return navigator.userAgent.indexOf("Lazycat_Android") != -1;
}
// 是否是ios webshell环境
export function isIosWebShell() {
    return navigator.userAgent.indexOf("Lazycat_Ios") != -1;
}
// h5调用webview 打开App
export function OpenApp(url) {
    if (!isWebShell()) {
        window.location.href = url;
        return
    }
    if (isAndroidWebShell()) {
        // @ts-ignore
        android.OpenApp(url);
    }
}
export function SetContentURL(url){
    if (!isWebShell()) {
        return
    }
    if (isAndroidWebShell()) {
        // @ts-ignore
        android.OpenApp(url);
    }else if(isPCWebShell()){
        // @ts-ignore
        windows.electronAPI.OpenApp();
    }
}
export function GetContentURL(callback){
    if (!isWebShell()) {
        return
    }
    if (isAndroidWebShell()) {
        // @ts-ignore
        android.GetContentURL(callback);
    }else if(isPCWebShell()){
        // @ts-ignore
        windows.electronAPI.GetContentURL(callback);
    }
}
export function BackContentView(){
    if(!isWebShell()){
        return
    }
    if (isAndroidWebShell()) {
        // @ts-ignore
        android.BackContentView();
    }else if(isPCWebShell()){
        // @ts-ignore
        windows.electronAPI.BackContentView();
    }
}

export function ForwardContentView(){
    if(!isWebShell()){
        return
    }
    if (isAndroidWebShell()) {
        // @ts-ignore
        android.ForwardContentView();
    }else if(isPCWebShell()){
        // @ts-ignore
        windows.electronAPI.ForwardContentView();
    }
}

function NotifyControView(jscontent){
    if(!isWebShell()){
        return
    }
    if (isAndroidWebShell()) {
        // @ts-ignore
        android.NotifyControView(jscontent);
    }else if(isPCWebShell()){
        // @ts-ignore
        windows.electronAPI.NotifyControView(jscontent);
    }
}

export default {
    NotifyControView,
    ForwardContentView,
    BackContentView,
    GetContentURL,
    SetContentURL,
    OpenApp,
    isIosWebShell,
    isAndroidWebShell,
    isWebShell,
    isPCWebShell
}
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

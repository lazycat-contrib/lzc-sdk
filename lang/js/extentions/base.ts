/*
 * @Author: Bin
 * @Date: 2023-03-03
 * @FilePath: /lzc-sdk/lang/js/extentions/base.ts
 */

enum PlatformType {
    Android,
    IOS,
    PC,
    Browser
}

declare global {
  interface Window {
    ipcRenderer: any
  }
}


// 不支持的平台 notSupportPlatform
function disabled(...platforms: PlatformType[]) {
    // console.log("PlatformType", list);
    return function (target: any, propertyKey: string, descriptor: PropertyDescriptor) {
        if (!(target.constructor === LzcAppSdk)) {
            console.error("LzcAppSdk:", `notSupportPlatform error ${target.constructor}`)
            return
        }

        let notSupport = false;
        platforms.forEach((item: PlatformType) => {
            switch (item) {
                case PlatformType.Android:
                    notSupport = target.isAndroidWebShell();
                    return
                case PlatformType.IOS:
                    notSupport = target.isIosWebShell();
                    return
                case PlatformType.PC:
                    notSupport = target.isPCWebShell();
                    return
                case PlatformType.Browser:
                    notSupport = !(target.isWebShell());
                    return
            }
        });

        if (notSupport) {
            descriptor.value = ((...args: any) => {
                console.error("LzcAppSdk:", `The current platform does not support the ${propertyKey} method`)
            })
        }

        // console.log("target", target);
        // console.log("propertyKey", propertyKey);
        // console.log("descriptor", descriptor);

        return descriptor;
    };
}

/**
 * @description: 指定函数支持平台函数装饰器
 * @param {array} platforms
 * @return {Function}
 */
function native(...platforms: PlatformType[]): (...args: any[]) => any {
    return function (target: object, propertyKey: string, descriptor: PropertyDescriptor) {
        const lzcSdk = LzcAppSdk.getInstance()
        if (!lzcSdk) {
            console.error("LzcAppSdk:", `instance is null`)
            return
        }

        let support = false;
        platforms.forEach((item: PlatformType) => {
            if (support) {
                return
            }
            switch (item) {
                case PlatformType.Android:
                    support = lzcSdk.isAndroidWebShell();
                    return
                case PlatformType.IOS:
                    support = lzcSdk.isIosWebShell();
                    return
                case PlatformType.PC:
                    support = lzcSdk.isPCWebShell();
                    return
                case PlatformType.Browser:
                    support = !(lzcSdk.isInApplication());
                    return
            }
        });

        const method = descriptor.value;
        // console.log("!!method ", method);
        descriptor.value = new Proxy(descriptor.value, {
            apply(target, thisArg, argArray) {
                // console.log("!!method apply", Object.prototype.toString.call(target));
                if (support) {
                    // 当前函数支持该平台，调用函数实现
                    return target.apply(thisArg, argArray)
                }

                // 当前函数不支持该平台，替换执行空函数体
                console.warn("LzcAppSdk:", `The current platform does not support the ${propertyKey} method`)
                if (Object.prototype.toString.call(target) === "[object AsyncFunction]") {
                    // AsyncFunction
                    return (async function () { return undefined })()
                } else {
                    // Function
                    return (() => { return undefined })()
                }
            },
        })

        // if (!support) {
        //     console.error("LzcAppSdk:", `The current platform does not support the ${propertyKey} method`)
        //     let fun = method

        //     if (!!method && typeof method.then === 'function') {
        //         descriptor.value = (...args: any) => { }
        //         fun = method.then
        //     } else {
        //         descriptor.value = async (...args: any) => { }
        //     }

        //     return descriptor
        // }

        // TODO: 处理 Native 函数注入拉平
        // let useNative = {}
        // useNative = new Proxy(useNative, {
        //     get(target, prop, receiver) {

        //         let item = target[prop]
        //         let nativeClient = LzcAppSdk.getInstance().client
        //         // let nativeClient = {
        //         //     open: async () => {
        //         //         console.log("useNative open.....");
        //         //     },
        //         //     getValue: async (name: String) => {
        //         //         console.log("useNative open.....");
        //         //         return name + ", emmmm"
        //         //     }
        //         // }

        //         // console.log("target", target);
        //         // console.log("prop", prop);
        //         // console.log("receiver", receiver);

        //         if (nativeClient) {
        //             // 原生注入对象存在
        //             let nativeAttribute = nativeClient[prop]
        //             if (!item && nativeAttribute) {
        //                 // 原生注入对象包含该属性/方法
        //                 item = nativeAttribute
        //                 if (typeof nativeAttribute == "function") {
        //                     // 原生注入的是函数，准备拦截
        //                     item = new Proxy<any>(item, {
        //                         apply(target, thisArg, argArray) {
        //                             return target(...argArray);
        //                         },
        //                     })
        //                 }
        //             }
        //         }

        //         // 拦截不存在的原生函数调用并提示警告
        //         if (!item && (!nativeClient || !nativeClient[prop])) {
        //             item = function (...params: any[]) { }
        //             console.warn("LzcAppSdk:", `The ${String(prop)} method is not implemented on the current platform`)
        //         }

        //         return item
        //     },
        // })

        // lzcSdk.useNative = useNative

        // descriptor.value = Promise.resolve(function (...args: any) {
        //     // console.log("应用", target);
        //     const context = Object.assign(target, this instanceof LzcAppSdk ? {} : { LzcApp: lzcSdk })
        //     method.apply(context, ...args);
        //     // method(...args)
        // })

        // console.log("target", target);
        // console.log("propertyKey", propertyKey);
        // console.log("descriptor", descriptor);
        return descriptor
    };
}

class LzcAppSdk {
    private static instance: LzcAppSdk;
    private static cacheJSBridgeIOS: any; // 缓存 IOS JSBridge

    protected constructor() {
        // 初始化 JSBridge 回调事件队列
        this.initCallbackQueue();
        (async () => {
            if (this.isIosWebShell()) {
                // 由于 IOS JSBridge 需要异步获取，构造对象时触发缓存 JSBridgeIOS 对象
                LzcAppSdk.cacheJSBridgeIOS = await LzcAppSdk.initJSBridgeIOS()
            }
        })();

        // 缓存 LzcAppSdk 对象应用到 window 对象
        if (!window.lzcAppSdk_client || window.lzcAppSdk_client !== this) {
            window.lzcAppSdk_client = this;
        }

    }

    /**
     * @description: 获取 LzcAppSdk 对象实例
     * @return {LzcAppSdk}
     */
    public static getInstance(): LzcAppSdk {
        if (!LzcAppSdk.instance) {
            LzcAppSdk.instance = new LzcAppSdk()
        }
        return LzcAppSdk.instance
    }

    /**
     * @description: 初始化 JSBridge 全局回调队列
     * @return {void}
     */
    private initCallbackQueue(): void {
        if (!window.lzcAppSdk_responseCallBackFuncDict || !window.lzcAppSdk_responseCallBackFuncUniqueID) {
            // 初始化回调方法存储器和索引ID
            window.lzcAppSdk_responseCallBackFuncDict = {}
            window.lzcAppSdk_responseCallBackFuncUniqueID = 1;
        }

        if (
            !window.lzcAppSdk_sendCallBackFunc ||
            window.lzcAppSdk_sendCallBackFunc != LzcAppSdk.sendCallBackFunc
        ) {
            window.lzcAppSdk_sendCallBackFunc = LzcAppSdk.sendCallBackFunc
        }
    }

    /**
     * @description: JS 调用原生函数时，添加回调方法到回调方法存储器，并返回索引ID
     * @param {Function} responseCallBackFunc
     * @return {string}
     */
    public static addToCallBackFuncDictWith(responseCallBackFunc: (...args: any[]) => any): string {
        if (!responseCallBackFunc) return 'unValid_funcUniqueID'
        const funcUniqueID = `lzc_${window.lzcAppSdk_responseCallBackFuncUniqueID++}_${new Date().getTime()}`
        window.lzcAppSdk_responseCallBackFuncDict[funcUniqueID] =
            responseCallBackFunc
        return funcUniqueID
    }


    /**
     * @description: 原生接受 JS 调用并处理相关操作后，发送回调给js, js 根据索引ID寻找回调方法来处理数据，然后移除js的回调方法
     * @param {string} funcUniqueID
     * @param {any} responseData
     * @return {void}
     */
    public static sendCallBackFunc(
        funcUniqueID: string,
        responseData: any
    ): void {
        if (!funcUniqueID) return
        const responseCallBackFunc =
            window.lzcAppSdk_responseCallBackFuncDict[funcUniqueID]
        if (!responseCallBackFunc) return
        responseCallBackFunc(responseData)
        delete window.lzcAppSdk_responseCallBackFuncDict[funcUniqueID]
    }

    /**
     * @description: 注册一个 JS 函数
     * @param {string} name
     * @return {Promise<any> | undefined}
     */
    private static async _registerCallBackFunc(name: string): Promise<any> | undefined {
        if (!name) return null
        return async function (...args: any[]) {
            let returnData = undefined
            const funcUniqueID = LzcAppSdk.addToCallBackFuncDictWith(function (
                data: any
            ) {
                returnData = data
            })

            const func = window.webkit.messageHandlers[name]
            if (!func) {
                console.error(
                    'LzcAppSdk:',
                    `ios webshell ${name} injection failed.`
                )
                return returnData
            }
            const value = await func.postMessage({
                funcUniqueID,
                params: [...args],
            })
            return value || returnData
        }
    }

    /**
     * @description: 获取 IOS 原生 JSBridge 对象
     * @return {Promise<any>}
     */
    // @native(PlatformType.IOS)
    private static async initJSBridgeIOS(): Promise<any> {
        let view: any = {}
        if (LzcAppSdk.cacheJSBridgeIOS) {
            view = LzcAppSdk.cacheJSBridgeIOS
        }
        view['ScriptHandlers'] = await LzcAppSdk._registerCallBackFunc('ScriptHandlers')
        const handlers = await view.ScriptHandlers()
        try {
            const data = typeof handlers == 'string' ? JSON.parse(handlers) : handlers
            if (!data || data.length < 1) {
                throw new Error('ios webshell ScriptHandlers handlers error.')
            }
        } catch (error) {
            console.error('LzcAppSdk: ', error)
            return
        }

        // 注册全部可用函数
        for (const key in handlers) {
            if (!handlers[key]) continue
            view[handlers[key]] = await LzcAppSdk._registerCallBackFunc(handlers[key])
        }

        // 缓存 IOS JSBridge 对象
        if (!LzcAppSdk.cacheJSBridgeIOS || (
            view &&
            Object.keys(view).length > 0 &&
            LzcAppSdk.cacheJSBridgeIOS !== view
        )) {
            LzcAppSdk.cacheJSBridgeIOS = view
        }

        return view
    }

    /**
     * @description: 根据原生注入命名空间 异步获取 JSBridge
     * @param {any} namespaces
     * @return {Promise<any>}
     */
    public async useNativeAsync(namespaces?: any): Promise<any> {
        namespaces = this.useNative(namespaces)
        if (this.isIosWebShell()) {
            // ios JSBridge 是异步注入的，强制使用最新的
            namespaces = await LzcAppSdk.initJSBridgeIOS();
        }
        return namespaces
    }

    /**
     * @description: 根据原生注入命名空间 同步获取 JSBridge (如调用原生方法不需要回调值的可以使用该函数获取原生对象否则最好使用 useNativeAsync 函数)
     * @param {any} namespaces
     * @return {any}
     */
    public useNative(namespaces?: any): any {
        if (!namespaces) {
            let defaultNamespaces = {}
            if (this.isIosWebShell() && LzcAppSdk.cacheJSBridgeIOS) {
                LzcAppSdk.initJSBridgeIOS(); // 触发刷新缓存 ios JSBridge
                defaultNamespaces = LzcAppSdk.cacheJSBridgeIOS
            }
            if (this.isAndroidWebShell()) {
                defaultNamespaces = android
            }
            if (this.isPCWebShell()) {
                defaultNamespaces = window.electronAPI
            }
            namespaces = defaultNamespaces
        }

        // TODO: 应该代理 JSBridge 对象函数调用，避免出现异常
        if (!namespaces) {
            console.warn("LzcAppSdk", "Does not exist JSBridge namespaces.")
            namespaces = {}
        }

        return namespaces
    }

    /**
     * @description: 判断是否属于 webshell 环境
     * @return {boolean}
     */
    public isWebShell(): boolean {
        return navigator.userAgent.indexOf("Lazycat") != -1 && !this.isControlView() && !this.isContentView();
    }

    /**
     * @description: 是否是android webshell 环境
     * @return {boolean}
     */
    public isAndroidWebShell(): boolean {
        return navigator.userAgent.indexOf("Lazycat_101") != -1;
    }

    /**
     * @description: 是否是pc webshell 环境
     * @return {boolean}
     */
    public isPCWebShell(): boolean {
        return !!window.ipcRenderer;
    }

    /**
     * @description: 是否是ios webshell 环境
     * @return {boolean}
     */
    public isIosWebShell(): boolean {
        return navigator.userAgent.indexOf("Lazycat_103") != -1;
    }

    /**
     * @description: 是否属于 Control webshell 环境
     * @return {boolean}
     */
    public isControlView(): boolean {
        return navigator.userAgent.indexOf("Lazycat_ControlView") != -1;
    }

    /**
     * @description: 是否属于 Content webshell 环境
     * @return {boolean}
     */
    public isContentView(): boolean {
        return navigator.userAgent.indexOf("Lazycat_ContentView") != -1;
    }

    /**
     * @description: 是否在客户端内
     * @return {boolean}
     */
    public isInApplication(): boolean {
        return navigator.userAgent.indexOf('Lazycat_Client') != -1 || this.isPCWebShell()
    }
}

class LzcAppSdkManage {
    protected static LzcApp: LzcAppSdk
    protected LzcApp: LzcAppSdk
}

const LzcApp = LzcAppSdk.getInstance()

/**
 * @description: 兼容旧的 android 注入的 js 调用（待迁移后删除改引用）
 * @deprecated: 该 JS 注入引用即将弃用，请尽快适配最新的 LzcAppSdk.useNative() 调用方式
 * @return {any}
 */
const LzcClient: any = LzcApp.useNative()

export default LzcApp
export { PlatformType as LzcAppPlatformType, LzcAppSdkManage, LzcClient, native }

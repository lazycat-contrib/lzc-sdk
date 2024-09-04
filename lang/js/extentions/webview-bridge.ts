// @ts-nocheck
if (window.lzc_bridge == undefined) {
  window.lzc_bridge = {}
}

export async function warpIosBridge(func: string, argv = "") {
  if (window.lzc_bridge === undefined) {
    window.lzc_bridge = {}
  }
  const key = Number(Math.random() * 10000000).toFixed()
  const handlerFuncName = `handler_${func}_${key}`

  window.lzc_bridge[handlerFuncName] = (v: any) => v

  const proxyHandler = (targetFunc: Function) => {
    return new Promise(resolve => {
      const proxyFunc = new Proxy(targetFunc, {
        apply: (target, thisArg, argvlist) => {
          const resp = Reflect.apply(target, thisArg, argvlist)
          console.log("resp:", resp)
          resolve(resp)
          return resp // 确保原函数的返回值不会丢失
        },
      })
      // 直接替换原始函数为代理后的函数
      window.lzc_bridge[handlerFuncName] = proxyFunc
    })
  }

  // 发送请求给 Swift 代码
  window.webkit.messageHandlers[func].postMessage({ handler: handlerFuncName, argv: argv })

  // 监听代理函数的调用
  return proxyHandler(window.lzc_bridge[handlerFuncName]).then(resp => {
    delete window.lzc_bridge[handlerFuncName]
    return resp
  })
}

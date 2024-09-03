if (window.lzc_bridge == undefined) {
  window.lzc_bridge = {}
}

async function warpIosBridge(func: string, argv = {}) {
  const key = Number(Math.random() * 10000000).toFixed()
  const handlerFuncName = `handler_${func}_${key}`

  window.lzc_bridge[handlerFuncName] = function (v) {
    console.log("v:", v)
    return v
  }

  const proxyHandler = async func => {
    return new Promise(res => {
      const proxyFunc = new Proxy(func, {
        apply: (target, thisArg, argvlist) => {
          const resp = target.apply(thisArg, argvlist)
          return res(resp)
        },
      })
      func = proxyFunc
    })
  }

  // requst swift code
  window.webkit.messageHandlers[func].postMessage({ handler: handlerFuncName, argv: argv })
  return await proxyHandler(window.lzc_bridge[handlerFuncName]).then(resp => {
    delete window.lzc_bridge[handlerFuncName]
    return resp
  })
}

export { warpIosBridge }

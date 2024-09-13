/*
 * @Author: Bin
 * @Date: 2023-03-03
 * @FilePath: /lzc-app-ext/src/lzc-sdk/lang/js/extentions/index.ts
 */

import LzcApp from "./base"

declare global {
  interface Window {
    electronAPI?: any
    webkit?: any
    lzcAppSdk_responseCallBackFuncDict?: any
    lzcAppSdk_responseCallBackFuncUniqueID?: any
    lzcAppSdk_sendCallBackFunc?: any
    lzcAppSdk_client?: any
    lzc_bridge: {}
  }
  const android: any
  const android_launch_service: any
  const android_dialog: any
  const lzc_app_manager: any
  const lzc_status_bar: any
  const lzc_vibrate: any
  const lzc_permission: any
}
export default LzcApp
export { LzcAppPlatformType, LzcClient } from "./base"
export { AppCommon } from "./app_common"

export { LzcApp }

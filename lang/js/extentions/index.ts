import { lzcAPIGateway } from './../index';
/*
 * @Author: Bin
 * @Date: 2023-03-03
 * @FilePath: /lzc-sdk/lang/js/extentions/index.ts
 */

import LzcApp from './base';

declare global {
    interface Window {
        electronAPI?: any
        webkit?: any
        lzcAppSdk_responseCallBackFuncDict?: any
        lzcAppSdk_responseCallBackFuncUniqueID?: any
        lzcAppSdk_sendCallBackFunc?: any
        lzcAppSdk_client?: any
    }
    const android: any
    const android_launch_service: any
}
export default LzcApp
export { LzcAppPlatformType, LzcClient } from './base'
export { AppCommon } from './app_common'

export { LzcApp }


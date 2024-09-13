/*
 * @Author: Bin
 * @Date: 2023-02-28
 * @FilePath: /lzc-sdk/examples/lzc-react-demo/src/LzcApp.tsx
 */
import React from 'react'
import {AppCommon} from '../../../lang/js/extentions';
import {SAFE} from '../../../lang/js/extentions/safe_app';
import './App.css';
import {ClientAuthorizationType} from "../../../lang/js/extentions/client_authorization";

export default function LzcApp() {
    const runtime = React.useRef(false)
    React.useEffect(() => {
        if (!runtime.current) {
            runtime.current = true
            // console.log(`app ok!!! load lzc sdk ${LzcAppSdk.isWebShell()}`);

            // AppCommon.ScriptHandlers().then((data) => {
            //     console.log(data);
            // })


        }
    }, []);

    const testLaunchApp = () => {
        AppCommon.LaunchApp("https://bin.zmide.com", "com.zmide.bin")
    }

    const testContactsPermission = () => {
        AppCommon.GetClientAuthorizationStatus(ClientAuthorizationType.Contacts)
    }

    const testInstall = () => {
        SAFE.InstallApk("http://192.168.100.81:8081/")
    }

    return (
        <div className="App"
             style={{flex: 1, width: '100%', display: 'flex', justifyContent: 'center', flexDirection: 'column'}}>
            <div style={{padding: "5px 10px", background: "#ff5161"}} onClick={testLaunchApp}>LzcApp</div>
            <div style={{padding: "5px 10px", marginTop: 10, background: "#ff5161"}} onClick={testInstall}>Install</div>
            <div style={{padding: "5px 10px", marginTop: 10, background: "#ff5161"}} onClick={testContactsPermission}>测试联系人权限</div>
        </div>
    )
}

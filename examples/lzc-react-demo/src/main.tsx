import React from 'react'
import ReactDOM from 'react-dom/client'
// import App from './App'
import './index.css'

import {
    AppCommon
} from "@lazycatcloud/sdk/extentions/index";

const AppView = () => {

    React.useEffect(() => {
    console.log("emmmm")

        AppCommon.LaunchApp("https://m.baidu.com","cloud.lazycat.app.re")
        // AppCommon.LaunchNativeApp("/sdcard/懒猫云离线/321.docx","cn.wps.moffice_eng:cn.wps.moffice.documentmanager.PreStartActivity2")

    }, [])

    return (
     <div>Hello</div>
    )
}

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <AppView />
  </React.StrictMode>
)

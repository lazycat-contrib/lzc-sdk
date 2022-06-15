import { GrpcWebImpl } from "./internal/grpcweb"

import { DevicesClientImpl, Devices } from "./devices/devices"
import { BrowserOnlyClientImpl, UserInfo, AppInfo } from "./browseronly/browseronly"
import { PermissionManager, PermissionManagerClientImpl, PermissionDesc, PermissionToken } from "./permissions/permissions"

import { DialogManagerClientImpl, DialogManager } from "./localdevice/dialog/dialog"
import { ClipboardManagerClientImpl, ClipboardManager } from "./localdevice/clipboard/clipboard"

import { grpc } from "@improbable-eng/grpc-web";


const browserTransport = grpc.CrossBrowserHttpTransport({ withCredentials: true });

export class lzcAPIGateway {
    constructor(host: string = window.origin) {
        const rpc = new GrpcWebImpl(host, {
            debug: true,
            transport: browserTransport,
        })
        this.devices = new DevicesClientImpl(rpc);

        let b = new BrowserOnlyClientImpl(rpc);
        this.userinfo = b.QueryUserInfo({})
        this.appinfo = b.QueryAppInfo({})

        this.pm = new PermissionManagerClientImpl(rpc);

    }

    private pm : PermissionManager;
    public userinfo: Promise<UserInfo>;
    public appinfo: Promise<AppInfo>;
    public devices: Devices;
}


export class DeviceProxy {
    constructor(apiurl :string)  {
        const rpc = new GrpcWebImpl(apiurl, {
            debug: true,
            transport: browserTransport,
        })

        this.dialog = new DialogManagerClientImpl(rpc)
        this.clipboard = new ClipboardManagerClientImpl(rpc)
    }

    public dialog : DialogManager;
    public clipboard: ClipboardManager;
}

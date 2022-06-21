import { GrpcWebImpl } from "./internal/grpcweb"

import { Empty } from "./google/protobuf/empty";

import { Observable, Subscriber } from "rxjs";

import { DevicesClientImpl, Devices, Device } from "./devices/devices"
import { BrowserOnly, BrowserOnlyClientImpl, UserInfo, AppInfo } from "./browseronly/browseronly"
import { PermissionManager, PermissionManagerClientImpl, PermissionDesc, PermissionToken } from "./permissions/permissions"

import { DialogManagerClientImpl, DialogManager } from "./localdevice/dialog/dialog"
import { ClipboardManagerClientImpl, ClipboardManager } from "./localdevice/clipboard/clipboard"

import { grpc } from "@improbable-eng/grpc-web";


const opt = {
    debug: true,
    transport: grpc.CrossBrowserHttpTransport({ withCredentials: true }),
    // streamingTransport: grpc.WebsocketTransport(),
}

export class lzcAPIGateway {
    constructor(host: string = window.origin) {
        const rpc = new GrpcWebImpl(host, opt)
        this.devices = new DevicesClientImpl(rpc);

        this.bo = new BrowserOnlyClientImpl(rpc);
        this.userinfo = this.bo.QueryUserInfo({})
        this.appinfo = this.bo.QueryAppInfo({})

        this.pm = new PermissionManagerClientImpl(rpc);
    }
    private bo : BrowserOnly;
    private pm : PermissionManager;
    public async openDevices() {
        return new Promise<void>((resolve, reject) => {
            this.bo.PairAllDevices({}).subscribe({
                error: (err) => reject(err),
                complete: () => resolve(),
            })
        })
    }

    public userinfo: Promise<UserInfo>;
    public appinfo: Promise<AppInfo>;
    public devices: Devices;
}

async function test() {
    let cc = new lzcAPIGateway()
    cc.openDevices()
}

export class DeviceProxy {
    constructor(apiurl :string)  {
        const rpc = new GrpcWebImpl(apiurl, opt)

        this.dialog = new DialogManagerClientImpl(rpc)
        this.clipboard = new ClipboardManagerClientImpl(rpc)
    }

    public dialog : DialogManager;
    public clipboard: ClipboardManager;
}

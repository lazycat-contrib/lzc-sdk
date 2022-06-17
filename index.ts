import { GrpcWebImpl } from "./internal/grpcweb"

import { Empty } from "./google/protobuf/empty";

import { Observable, Subscriber } from "rxjs";

import { DevicesClientImpl, Devices, Device, KeepConnectRequest } from "./devices/devices"
import { BrowserOnlyClientImpl, UserInfo, AppInfo } from "./browseronly/browseronly"
import { PermissionManager, PermissionManagerClientImpl, PermissionDesc, PermissionToken } from "./permissions/permissions"

import { DialogManagerClientImpl, DialogManager } from "./localdevice/dialog/dialog"
import { ClipboardManagerClientImpl, ClipboardManager } from "./localdevice/clipboard/clipboard"

import { grpc } from "@improbable-eng/grpc-web";


const opt = {
    debug: true,
    transport: grpc.CrossBrowserHttpTransport({ withCredentials: true }),
    streamingTransport: grpc.WebsocketTransport(),
}

export class lzcAPIGateway {
    constructor(host: string = window.origin) {
        const rpc = new GrpcWebImpl(host, opt)
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

    private keepDeviceSub: Subscriber<KeepConnectRequest>;

    async keepConnect(d: string) {
        if (!this.keepDeviceSub) {
            let req = new Observable<KeepConnectRequest>(
                subscriber => this.keepDeviceSub = subscriber
            )
            this.devices.KeepConnect(req)
        }
        this.keepDeviceSub.next({
            sourceDevice: (await this.userinfo).deviceId,
            open: true,
            devices: [d],
        })
    }
}

function test() {
    let cc = new lzcAPIGateway()

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

import { GrpcWebImpl } from "./grpcweb"

import { Empty } from "./google/protobuf/empty";

import { Observable, Subscriber } from "rxjs";

import { DevicesClientImpl, Devices, Device } from "./common/devices"
import { UserManagerClientImpl, UserManager } from "./common/users"
import { BrowserOnlyProxy, BrowserOnlyProxyClientImpl, SessionInfo, AppInfo } from "./common/browseronly"
import { PermissionManager, PermissionManagerClientImpl, PermissionDesc, PermissionToken } from "./common/permissions"
import { APIGateway, APIGatewayClientImpl } from "./common/gateway"

import { DialogManagerClientImpl, DialogManager } from "./localdevice/dialog"
import { ClipboardManagerClientImpl, ClipboardManager } from "./localdevice/clipboard"
import { PhotoLibraryClientImpl, PhotoLibrary } from "./localdevice/photo"
import { NetworkManagerClientImpl, NetworkManager } from "./localdevice/network"

import { grpc } from "@improbable-eng/grpc-web";


const opt = {
    debug: true,
    transport: grpc.CrossBrowserHttpTransport({ withCredentials: true }),
    // streamingTransport: grpc.WebsocketTransport(),
}



async function buildCurrentDevice(cc: lzcAPIGateway) : Promise<DeviceProxy>{
    let s = await cc.session
    let uid = s.uid
    let ds = await cc.devices.ListDevices({uid})
    let d = ds.devices.find((d) => d.peerId == s.deviceId)
    return new DeviceProxy(d.deviceApiUrl)
}


export class lzcAPIGateway {
    constructor(host: string = window.origin) {
        host = host.replace(/\/+$/, '')
        const rpc = new GrpcWebImpl(host, opt)
        this.devices = new DevicesClientImpl(rpc);
        this.users = new UserManagerClientImpl(rpc);

        this.bo = new BrowserOnlyProxyClientImpl(rpc);
        this.session = this.bo.QuerySessionInfo({})
        this.appinfo = this.bo.QueryAppInfo({})
        this.gw = new APIGatewayClientImpl(rpc);

        this.pm = new PermissionManagerClientImpl(rpc);

        this.currentDevice = buildCurrentDevice(this)
    }
    private bo : BrowserOnlyProxy;
    private gw : APIGateway;
    private pm : PermissionManager;
    public async openDevices() {
        return new Promise<void>((resolve, reject) => {
            this.bo.PairAllDevices({}).subscribe({
                error: (err) => reject(err),
                complete: () => resolve(),
            })
        })
    }

    public users: UserManager;

    public session: Promise<SessionInfo>;

    public currentDevice: Promise<DeviceProxy>;

    public appinfo: Promise<AppInfo>;
    public devices: Devices;
}

async function test() {
    let cc = new lzcAPIGateway()
    cc.openDevices()
}

export class DeviceProxy {
    constructor(apiurl :string) {
        const rpc = new GrpcWebImpl(apiurl, opt)

        this.dialog = new DialogManagerClientImpl(rpc)
        this.clipboard = new ClipboardManagerClientImpl(rpc)
        this.photolibrary = new PhotoLibraryClientImpl(rpc)
        this.network = new NetworkManagerClientImpl(rpc)
    }

    public dialog : DialogManager;
    public clipboard: ClipboardManager;
    public photolibrary: PhotoLibrary;
    public network: NetworkManager;
}

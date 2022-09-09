import { GrpcWebImpl } from "./grpcweb"

import { Empty } from "./google/protobuf/empty";

import { Observable, Subscriber } from "rxjs";

import { EndDeviceServiceClientImpl, EndDeviceService, EndDevice } from "./common/end_device"
import { UserManagerClientImpl, UserManager } from "./common/users"
import { BoxService, BoxServiceClientImpl } from "./common/box"
import { BrowserOnlyProxy, BrowserOnlyProxyClientImpl, SessionInfo, AppInfo } from "./common/browseronly"
import { PermissionManager, PermissionManagerClientImpl, PermissionDesc, PermissionToken } from "./common/security_context"
import { APIGateway, APIGatewayClientImpl } from "./common/gateway"
import { PeripheralDeviceService, PeripheralDeviceServiceClientImpl } from "./common/peripheral_device"
import { PackageManager, PackageManagerClientImpl } from "./sys/package_manager"
import { NetworkManager as NM, NetworkManagerClientImpl as NMClientImpl } from "./sys/network_manager"

import { OSSnapshotService, OSSnapshotServiceClientImpl } from "./sys/OS_snapshot"
import { OSUpgradeService, OSUpgradeServiceClientImpl } from "./sys/OS_upgrader"
import { IngressService, IngressServiceClientImpl } from "./sys/ingress"

import { DialogManagerClientImpl, DialogManager } from "./localdevice/dialog"
import { ClipboardManagerClientImpl, ClipboardManager } from "./localdevice/clipboard"
import { PhotoLibraryClientImpl, PhotoLibrary } from "./localdevice/photo"
import { NetworkManagerClientImpl, NetworkManager } from "./localdevice/network"
import { DeviceServiceClientImpl, DeviceService } from "./localdevice/device"

import { grpc } from "@improbable-eng/grpc-web";


const opt = {
    debug: true,
    transport: grpc.CrossBrowserHttpTransport({ withCredentials: true }),
    // streamingTransport: grpc.WebsocketTransport(),
}



async function buildCurrentDevice(cc: lzcAPIGateway) : Promise<EndDeviceProxy>{
    let s = await cc.session
    let uid = s.uid
    let ds = await cc.devices.ListEndDevices({uid})
    let d = ds.devices.find((d) => d.peerId == s.deviceId)
    let u = new URL(d.deviceApiUrl)

    return new EndDeviceProxy("https://" + u.host.replace(".d.", ".local."))
}


export class lzcAPIGateway {
    constructor(host: string = window.origin) {
        host = host.replace(/\/+$/, '')
        const rpc = new GrpcWebImpl(host, opt)
        this.devices = new EndDeviceServiceClientImpl(rpc);
        this.users = new UserManagerClientImpl(rpc);

        this.bo = new BrowserOnlyProxyClientImpl(rpc);
        this.session = this.bo.QuerySessionInfo({})
        this.appinfo = this.bo.QueryAppInfo({})
        this.gw = new APIGatewayClientImpl(rpc);

        this.nm = new NMClientImpl(rpc);
        this.pm = new PermissionManagerClientImpl(rpc);
        this.pkgm = new PackageManagerClientImpl(rpc);

        this.pd = new PeripheralDeviceServiceClientImpl(rpc);

        this.ingress = new IngressServiceClientImpl(rpc);
        this.box = new BoxServiceClientImpl(rpc);

        this.osUpgrader = new OSUpgradeServiceClientImpl(rpc);
        this.osSnapshot = new OSSnapshotServiceClientImpl(rpc);

        this.currentDevice = buildCurrentDevice(this)

        dumpInfo(this.bo)
    }
    private bo : BrowserOnlyProxy;
    private gw : APIGateway;
    private pm : PermissionManager;
    private pd : PeripheralDeviceService;

    public async openDevices() {
        return new Promise<void>((resolve, reject) => {
            this.bo.PairAllDevices({}).subscribe({
                error: (err) => reject(err),
                complete: () => resolve(),
            })
        })
    }

    public nm: NM; // 盒子内部的NetworkManager 
    public pkgm: PackageManager;
    public users: UserManager;
    public box: BoxService;
    public ingress: IngressService

    public session: Promise<SessionInfo>;

    public currentDevice: Promise<EndDeviceProxy>;

    public osUpgrader: OSUpgradeService
    public osSnapshot: OSSnapshotService

    public appinfo: Promise<AppInfo>;
    public devices: EndDeviceService;
}

async function test() {
    let cc = new lzcAPIGateway()
    cc.openDevices()
}

export class EndDeviceProxy {
    constructor(apiurl :string) {
        const rpc = new GrpcWebImpl(apiurl, opt)

        this.dialog = new DialogManagerClientImpl(rpc)
        this.clipboard = new ClipboardManagerClientImpl(rpc)
        this.photolibrary = new PhotoLibraryClientImpl(rpc)
        this.network = new NetworkManagerClientImpl(rpc)
        this.device = new DeviceServiceClientImpl(rpc)
    }

    public device : DeviceService;
    public dialog : DialogManager;
    public clipboard: ClipboardManager;
    public photolibrary: PhotoLibrary;
    public network: NetworkManager;
}


import pkg from './package.json';

async function dumpInfo(bo :BrowserOnlyProxy) {
    function capsule(title, info) {
        console.log(
            `%c ${title} %c ${info} %c`,
            'background:#35495E; padding: 1px; border-radius: 3px 0 0 3px; color: #fff;',
            `background:#3488ff; padding: 1px; border-radius: 0 3px 3px 0;  color: #fff;`,
            'background:transparent'
        )
    }
    capsule(`The ${pkg.name} version is`, `${pkg.version}`)

    let info = await bo.QueryAPIServerInfo({})
    capsule(`LZC SDK Version is`, `${info.frontendVersion}`)
}

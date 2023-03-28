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
import { OsDaemonService, OsDaemonServiceClientImpl } from "./sys/OS_daemon"

import { DialogManagerClientImpl, DialogManager } from "./localdevice/dialog"
import { UserConfig, UserConfigClientImpl } from "./localdevice/config";
import { ClipboardManagerClientImpl, ClipboardManager } from "./localdevice/clipboard"
import { PhotoLibraryClientImpl, PhotoLibrary } from "./localdevice/photo"
import { NetworkManagerClientImpl, NetworkManager } from "./localdevice/network"
import { DeviceServiceClientImpl, DeviceService } from "./localdevice/device"
import { PermissionManager as DevicePermissionManager, PermissionManagerClientImpl as DevicePermissionManagerClientImpl } from "./localdevice/permission"
import { FileHandlerClientImpl, FileHandler } from "./common/file_handler"
import { FileTransferServiceClientImpl, FileTransferService } from "./common/filetrans"
import { LocalLaunchService, LocalLaunchServiceClientImpl } from "./localdevice/local-launch"
import { RemoteMediaPlayerService, RemoteMediaPlayerServiceClientImpl } from "./dlna/dlna"
import { grpc } from "@improbable-eng/grpc-web";


const opt = {
    debug: true,
    transport: grpc.CrossBrowserHttpTransport({ withCredentials: true }),
    // streamingTransport: grpc.WebsocketTransport(),
}


async function getApiUrl(cc: lzcAPIGateway): Promise<URL> {
    let s = await cc.session
    let uid = s.uid
    let ds = await cc.devices.ListEndDevices({ uid })
    let d = ds.devices.find((d) => d.uniqueDeivceId == s.deviceId)
    return new URL(d.deviceApiUrl)
}

async function getAuthToken(host: string, apiurl: string): Promise<string> {
    const resp = (await fetch(host + "/_lzc/auth_token", {
        method: "POST",
        body: apiurl,
    }))
    if (!resp.ok) {
        throw new Error(`${resp.status}: ${resp.statusText}`)
    }
    return resp.json()["Token"]
}

async function buildCurrentDevice(cc: lzcAPIGateway): Promise<EndDeviceProxy> {
    const url = (await getApiUrl(cc)).toString().replace(/\/+$/, '')

    const authToken = await getAuthToken(cc.host, url)
    cc.authToken = authToken

    const metadata = new grpc.Metadata()
    try {
        metadata.set("lzc_dapi_auth_token", authToken)
    } catch (e) {
        console.log(e)
    }

    const rpc = new GrpcWebImpl(url, {...opt, ...{metadata: metadata}})
    return new EndDeviceProxy(rpc)
}


export class lzcAPIGateway {
    constructor(host: string = window.origin) {
        host = host.replace(/\/+$/, '')
        this.host = host

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
        this.osDaemon = new OsDaemonServiceClientImpl(rpc);

        this.rmp = new RemoteMediaPlayerServiceClientImpl(rpc);

        this.fileTransfer = new FileTransferServiceClientImpl(rpc)

        this.devopt = new DevOptServiceClientImpl(rpc);

        this.currentDevice = buildCurrentDevice(this)
        dumpInfo(this.bo)
    }
    public host: string;

    private bo: BrowserOnlyProxy;
    private gw: APIGateway;
    private pm: PermissionManager;
    public pd: PeripheralDeviceService;

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
    public osDaemon: OsDaemonService

    public appinfo: Promise<AppInfo>;
    public fileTransfer: FileTransferService;
    public devopt: DevOptService;
    public rmp: RemoteMediaPlayerService;
    public devices: EndDeviceService;

    public authToken: string;
}

async function test() {
    let cc = new lzcAPIGateway()
    cc.openDevices()
}

export class EndDeviceProxy {
    constructor(rpc: GrpcWebImpl) {
        this.dialog = new DialogManagerClientImpl(rpc)
        this.config = new UserConfigClientImpl(rpc);
        this.clipboard = new ClipboardManagerClientImpl(rpc)
        this.photolibrary = new PhotoLibraryClientImpl(rpc)
        this.network = new NetworkManagerClientImpl(rpc)
        this.device = new DeviceServiceClientImpl(rpc)
        this.fileHandler = new FileHandlerClientImpl(rpc)
        this.permission = new DevicePermissionManagerClientImpl(rpc)
        this.localLaunch = new LocalLaunchServiceClientImpl(rpc);
    }

    public device: DeviceService;
    public dialog: DialogManager;
    public config: UserConfig;
    public clipboard: ClipboardManager;
    public photolibrary: PhotoLibrary;
    public network: NetworkManager;
    public fileHandler: FileHandler;
    public permission: DevicePermissionManager;
    public localLaunch: LocalLaunchService;
}


import pkg from './package.json';
import { DevOptService, DevOptServiceClientImpl } from "./sys/devopt";

async function dumpInfo(bo: BrowserOnlyProxy) {
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

/**
 * 是否是webshell环境
 */
export function isWebShell() {
    return navigator.userAgent.indexOf("Lazycat") != -1;
}

// Local Variables:
// typescript-ts-mode-indent-offset: 4
// End:

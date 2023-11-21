import { GrpcWebImpl } from "./grpcweb"

import { EndDeviceServiceClientImpl, EndDeviceService } from "./common/end_device"
import { UserManagerClientImpl, UserManager } from "./common/users"
import { BoxService, BoxServiceClientImpl } from "./common/box"
import { BrowserOnlyProxy, BrowserOnlyProxyClientImpl, SessionInfo, AppInfo } from "./common/browseronly"
import { PeripheralDeviceService, PeripheralDeviceServiceClientImpl } from "./common/peripheral_device"
import { PackageManager, PackageManagerClientImpl } from "./sys/package_manager"
import { NetworkManager as NM, NetworkManagerClientImpl as NMClientImpl } from "./sys/network_manager"

import { OSSnapshotService, OSSnapshotServiceClientImpl } from "./sys/OS_snapshot"
import { OSUpgradeService, OSUpgradeServiceClientImpl } from "./sys/OS_upgrader"

import { AccessControlerService, AccessControlerServiceClientImpl } from "./sys/ingress"

import { DialogManagerClientImpl, DialogManager } from "./localdevice/dialog"
import { UserConfig, UserConfigClientImpl } from "./localdevice/config"
import { ClipboardManagerClientImpl, ClipboardManager } from "./localdevice/clipboard"
import { PhotoLibraryClientImpl, PhotoLibrary } from "./localdevice/photo"
import { NetworkManagerClientImpl, NetworkManager } from "./localdevice/network"
import { DeviceServiceClientImpl, DeviceService } from "./localdevice/device"
import { PermissionManager as DevicePermissionManager, PermissionManagerClientImpl as DevicePermissionManagerClientImpl } from "./localdevice/permission"
import { FileHandlerClientImpl, FileHandler } from "./common/file_handler"
import { FileTransferServiceClientImpl, FileTransferService } from "./common/filetrans"
import { LocalLaunchService, LocalLaunchServiceClientImpl } from "./localdevice/local-launch"
import { Client, ClientClientImpl } from "./localdevice/client"
import { RemoteMediaPlayerService, RemoteMediaPlayerServiceClientImpl } from "./dlna/dlna"
import { grpc } from "@improbable-eng/grpc-web"
const opt = {
  transport: grpc.CrossBrowserHttpTransport({ withCredentials: true }),
  debug: false,
}

export class lzcAPIGateway {
  constructor(host: string = window.origin, debug = false) {
    host = host.replace(/\/+$/, "")
    this.host = host

    opt.debug = debug
    const rpc = new GrpcWebImpl(host, opt)
    this.devices = new EndDeviceServiceClientImpl(rpc)
    this.users = new UserManagerClientImpl(rpc)

    this.bo = new BrowserOnlyProxyClientImpl(rpc)
    this._session = this.bo.QuerySessionInfo({})
    this.appinfo = this.bo.QueryAppInfo({})

    this.pkgm = new PackageManagerClientImpl(rpc)

    this.pd = new PeripheralDeviceServiceClientImpl(rpc)
    this.ac = new AccessControlerServiceClientImpl(rpc)
    this.box = new BoxServiceClientImpl(rpc)
    this.nm = new NMClientImpl(rpc)
    this.osSnapshot = new OSSnapshotServiceClientImpl(rpc)
    this.osUpgrade = new OSUpgradeServiceClientImpl(rpc)

    this.rmp = new RemoteMediaPlayerServiceClientImpl(rpc)

    this.fileTransfer = new FileTransferServiceClientImpl(rpc)

    this.devopt = new DevOptServiceClientImpl(rpc)

    this.message = new MessageServiceClientImpl(rpc)
    dumpInfo(this.bo)
  }

  private bo: BrowserOnlyProxy
  private _currentDevice: EndDeviceProxy
  private deviceApiTokenDeadline: number
  private _session: Promise<SessionInfo>
  public host: string
  public pd: PeripheralDeviceService

  public nm: NM // 盒子内部的NetworkManager
  public pkgm: PackageManager
  public users: UserManager
  public box: BoxService

  public ac: AccessControlerService
  public osSnapshot: OSSnapshotService
  public osUpgrade: OSUpgradeService

  public appinfo: Promise<AppInfo>
  public fileTransfer: FileTransferService
  public devopt: DevOptService
  public rmp: RemoteMediaPlayerService
  public devices: EndDeviceService

  public message: MessageServiceClientImpl

  public async openDevices() {
    return new Promise<void>((resolve, reject) => {
      this.bo.PairAllDevices({}).subscribe({
        error: err => reject(err),
        complete: () => resolve(),
      })
    })
  }

  // 获取当前设备 url
  public async currentDeviceURL(gateway: lzcAPIGateway = this): Promise<URL | undefined> {
    let session = await gateway.session
    let uid = session.uid
    let endDevices = await gateway.devices.ListEndDevices({ uid })
    let endDevice = endDevices.devices.find(d => d.uniqueDeivceId == session.deviceId)
    return new URL(endDevice.deviceApiUrl)
  }

  public authToken: string
  public get session(): Promise<SessionInfo> {
    return new Promise<SessionInfo>(async res => {
      this._session = this.bo.QuerySessionInfo({})
      res(this._session)
    })
  }
  public get currentDevice(): Promise<EndDeviceProxy> {
    if (this._currentDevice && Date.now() > this.deviceApiTokenDeadline) {
      return new Promise(res => res(this._currentDevice))
    }

    async function currentDeviceApiHost(cc: lzcAPIGateway): Promise<string> {
      const url = await cc.currentDeviceURL(cc)
      return `${url.protocol}//${url.hostname}`
    }
    async function requestAuthToken(cc: lzcAPIGateway, deviceApiUrl: string): Promise<string> {
      const resp = await fetch(cc.host + "/_lzc/deviceapi_auth_token", {
        method: "POST",
        body: deviceApiUrl,
      })
      if (!resp.ok) {
        throw new Error(`${resp.status}: ${resp.statusText}`)
      }
      const respJson = await resp.json()
      const token: string = respJson["Token"]
      const deadline: string = respJson["Deadline"]

      cc.deviceApiTokenDeadline = Date.parse(deadline)

      if (token === undefined) {
        throw new Error(`Token not set: ${respJson}`)
      }
      return token
    }
    return new Promise(async res => {
      const deviceApiUrl = await currentDeviceApiHost(this)
      const authToken = await requestAuthToken(this, deviceApiUrl)
      // save authToken for other uses(eg. webdav auth)
      this.authToken = authToken
      // build grpc metadata for credentials
      const metadata = new grpc.Metadata()
      metadata.set("lzc_dapi_auth_token", authToken)

      const rpc = new GrpcWebImpl(deviceApiUrl, { ...opt, ...{ metadata: metadata } })
      this._currentDevice = new EndDeviceProxy(rpc)
      res(this._currentDevice)
    })
  }
}

async function test() {
  let cc = new lzcAPIGateway()
  cc.openDevices()
}

export class EndDeviceProxy {
  constructor(rpc: GrpcWebImpl) {
    this.dialog = new DialogManagerClientImpl(rpc)
    this.config = new UserConfigClientImpl(rpc)
    this.clipboard = new ClipboardManagerClientImpl(rpc)
    this.photolibrary = new PhotoLibraryClientImpl(rpc)
    this.network = new NetworkManagerClientImpl(rpc)
    this.device = new DeviceServiceClientImpl(rpc)
    this.fileHandler = new FileHandlerClientImpl(rpc)
    this.permission = new DevicePermissionManagerClientImpl(rpc)
    this.localLaunch = new LocalLaunchServiceClientImpl(rpc)
    this.client = new ClientClientImpl(rpc)
  }
  public device: DeviceService
  public dialog: DialogManager
  public config: UserConfig
  public clipboard: ClipboardManager
  public photolibrary: PhotoLibrary
  public network: NetworkManager
  public fileHandler: FileHandler
  public permission: DevicePermissionManager
  public localLaunch: LocalLaunchService
  public client: Client
}

import pkg from "./package.json"
import { DevOptService, DevOptServiceClientImpl } from "./sys/devopt"
import { MessageServiceClientImpl } from "./common/message"

async function dumpInfo(bo: BrowserOnlyProxy) {
  function capsule(title, info) {
    console.log(
      `%c ${title} %c ${info} %c`,
      "background:#35495E; padding: 1px; border-radius: 3px 0 0 3px; color: #fff;",
      `background:#3488ff; padding: 1px; border-radius: 0 3px 3px 0;  color: #fff;`,
      "background:transparent"
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
  return navigator.userAgent.indexOf("Lazycat") != -1
}

import { GrpcWebImpl } from "./internal/grpcweb"

import { DevicesClientImpl, Devices } from "./devices/devices"
import { BrowserOnlyClientImpl, UserInfo } from "./browseronly/browseronly"

import { grpc } from "@improbable-eng/grpc-web";

export class lzcAPIGateway {
    constructor(host: string) {
        const myt = grpc.CrossBrowserHttpTransport({ withCredentials: true });

        const rpc = new GrpcWebImpl(host, {
            debug: true,
            transport: myt,
        })
        this.devices = new DevicesClientImpl(rpc);

        let b = new BrowserOnlyClientImpl(rpc);
        this.userinfo = b.QueryUserInfo({})
    }

    public userinfo: Promise<UserInfo>;
    public devices: Devices;
}

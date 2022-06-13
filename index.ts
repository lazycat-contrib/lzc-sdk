import { GrpcWebImpl } from "./internal/grpcweb"

import { DevicesClientImpl, Devices } from "./devices/devices"

import { grpc } from "@improbable-eng/grpc-web";

export class lzcAPIGateway {
    constructor(host: string) {
        const myt = grpc.CrossBrowserHttpTransport({ withCredentials: true });

        const rpc = new GrpcWebImpl(host, {
            debug: true,
            transport: myt,
        })
        this.devices = new DevicesClientImpl(rpc);
    }
    public devices: Devices;
}

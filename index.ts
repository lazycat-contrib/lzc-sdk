import { GrpcWebImpl } from "./internal/grpcweb"

import { DevicesClientImpl, Devices } from "./devices/devices"

export class lzcAPIGateway {
    constructor(host: string) {
        const rpc = new GrpcWebImpl(host, {})
        this.devices = new DevicesClientImpl(rpc);
    }

    public devices: Devices;
}

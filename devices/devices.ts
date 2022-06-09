/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { BrowserHeaders } from "browser-headers";

export const protobufPackage = "cloud.lazycat.apis";

export interface Device {
  peerId: string;
  isOnline: boolean;
  /** 因为device api的监听端口可能会变化，所以此url有效性不会太长 */
  deviceApiUrl: string;
}

export interface ListDeviceRequest {
  uid: string;
}

export interface ListDeviceReply {
  devices: Device[];
}

function createBaseDevice(): Device {
  return { peerId: "", isOnline: false, deviceApiUrl: "" };
}

export const Device = {
  encode(
    message: Device,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.peerId !== "") {
      writer.uint32(10).string(message.peerId);
    }
    if (message.isOnline === true) {
      writer.uint32(16).bool(message.isOnline);
    }
    if (message.deviceApiUrl !== "") {
      writer.uint32(26).string(message.deviceApiUrl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Device {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDevice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.peerId = reader.string();
          break;
        case 2:
          message.isOnline = reader.bool();
          break;
        case 3:
          message.deviceApiUrl = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Device {
    return {
      peerId: isSet(object.peerId) ? String(object.peerId) : "",
      isOnline: isSet(object.isOnline) ? Boolean(object.isOnline) : false,
      deviceApiUrl: isSet(object.deviceApiUrl)
        ? String(object.deviceApiUrl)
        : "",
    };
  },

  toJSON(message: Device): unknown {
    const obj: any = {};
    message.peerId !== undefined && (obj.peerId = message.peerId);
    message.isOnline !== undefined && (obj.isOnline = message.isOnline);
    message.deviceApiUrl !== undefined &&
      (obj.deviceApiUrl = message.deviceApiUrl);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Device>, I>>(object: I): Device {
    const message = createBaseDevice();
    message.peerId = object.peerId ?? "";
    message.isOnline = object.isOnline ?? false;
    message.deviceApiUrl = object.deviceApiUrl ?? "";
    return message;
  },
};

function createBaseListDeviceRequest(): ListDeviceRequest {
  return { uid: "" };
}

export const ListDeviceRequest = {
  encode(
    message: ListDeviceRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListDeviceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListDeviceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListDeviceRequest {
    return {
      uid: isSet(object.uid) ? String(object.uid) : "",
    };
  },

  toJSON(message: ListDeviceRequest): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListDeviceRequest>, I>>(
    object: I
  ): ListDeviceRequest {
    const message = createBaseListDeviceRequest();
    message.uid = object.uid ?? "";
    return message;
  },
};

function createBaseListDeviceReply(): ListDeviceReply {
  return { devices: [] };
}

export const ListDeviceReply = {
  encode(
    message: ListDeviceReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.devices) {
      Device.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListDeviceReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListDeviceReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.devices.push(Device.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListDeviceReply {
    return {
      devices: Array.isArray(object?.devices)
        ? object.devices.map((e: any) => Device.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ListDeviceReply): unknown {
    const obj: any = {};
    if (message.devices) {
      obj.devices = message.devices.map((e) =>
        e ? Device.toJSON(e) : undefined
      );
    } else {
      obj.devices = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListDeviceReply>, I>>(
    object: I
  ): ListDeviceReply {
    const message = createBaseListDeviceReply();
    message.devices = object.devices?.map((e) => Device.fromPartial(e)) || [];
    return message;
  },
};

export interface Devices {
  ListDevices(
    request: DeepPartial<ListDeviceRequest>,
    metadata?: grpc.Metadata
  ): Promise<ListDeviceReply>;
}

export class DevicesClientImpl implements Devices {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ListDevices = this.ListDevices.bind(this);
  }

  ListDevices(
    request: DeepPartial<ListDeviceRequest>,
    metadata?: grpc.Metadata
  ): Promise<ListDeviceReply> {
    return this.rpc.unary(
      DevicesListDevicesDesc,
      ListDeviceRequest.fromPartial(request),
      metadata
    );
  }
}

export const DevicesDesc = {
  serviceName: "cloud.lazycat.apis.Devices",
};

export const DevicesListDevicesDesc: UnaryMethodDefinitionish = {
  methodName: "ListDevices",
  service: DevicesDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ListDeviceRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...ListDeviceReply.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

interface UnaryMethodDefinitionishR
  extends grpc.UnaryMethodDefinition<any, any> {
  requestStream: any;
  responseStream: any;
}

type UnaryMethodDefinitionish = UnaryMethodDefinitionishR;

interface Rpc {
  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined
  ): Promise<any>;
}

export class GrpcWebImpl {
  private host: string;
  private options: {
    transport?: grpc.TransportFactory;

    debug?: boolean;
    metadata?: grpc.Metadata;
  };

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory;

      debug?: boolean;
      metadata?: grpc.Metadata;
    }
  ) {
    this.host = host;
    this.options = options;
  }

  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined
  ): Promise<any> {
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata =
      metadata && this.options.metadata
        ? new BrowserHeaders({
            ...this.options?.metadata.headersMap,
            ...metadata?.headersMap,
          })
        : metadata || this.options.metadata;
    return new Promise((resolve, reject) => {
      grpc.unary(methodDesc, {
        request,
        host: this.host,
        metadata: maybeCombinedMetadata,
        transport: this.options.transport,
        debug: this.options.debug,
        onEnd: function (response) {
          if (response.status === grpc.Code.OK) {
            resolve(response.message);
          } else {
            const err = new Error(response.statusMessage) as any;
            err.code = response.status;
            err.metadata = response.trailers;
            reject(err);
          }
        },
      });
    });
  }
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

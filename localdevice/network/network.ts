/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../../google/protobuf/empty";
import { BrowserHeaders } from "browser-headers";

export enum ConnectionType {
  Unknown = 0,
  CELL_NONE = 1,
  ETHERNET = 2,
  WIFI = 3,
  CELL = 4,
  CELL_2G = 5,
  CELL_3G = 6,
  CELL_4G = 7,
  CELL_5G = 8,
  UNRECOGNIZED = -1,
}

export function connectionTypeFromJSON(object: any): ConnectionType {
  switch (object) {
    case 0:
    case "Unknown":
      return ConnectionType.Unknown;
    case 1:
    case "CELL_NONE":
      return ConnectionType.CELL_NONE;
    case 2:
    case "ETHERNET":
      return ConnectionType.ETHERNET;
    case 3:
    case "WIFI":
      return ConnectionType.WIFI;
    case 4:
    case "CELL":
      return ConnectionType.CELL;
    case 5:
    case "CELL_2G":
      return ConnectionType.CELL_2G;
    case 6:
    case "CELL_3G":
      return ConnectionType.CELL_3G;
    case 7:
    case "CELL_4G":
      return ConnectionType.CELL_4G;
    case 8:
    case "CELL_5G":
      return ConnectionType.CELL_5G;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ConnectionType.UNRECOGNIZED;
  }
}

export function connectionTypeToJSON(object: ConnectionType): string {
  switch (object) {
    case ConnectionType.Unknown:
      return "Unknown";
    case ConnectionType.CELL_NONE:
      return "CELL_NONE";
    case ConnectionType.ETHERNET:
      return "ETHERNET";
    case ConnectionType.WIFI:
      return "WIFI";
    case ConnectionType.CELL:
      return "CELL";
    case ConnectionType.CELL_2G:
      return "CELL_2G";
    case ConnectionType.CELL_3G:
      return "CELL_3G";
    case ConnectionType.CELL_4G:
      return "CELL_4G";
    case ConnectionType.CELL_5G:
      return "CELL_5G";
    case ConnectionType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface NetworkInformation {
  ctype: ConnectionType;
  IsOnline: boolean;
}

function createBaseNetworkInformation(): NetworkInformation {
  return { ctype: 0, IsOnline: false };
}

export const NetworkInformation = {
  encode(
    message: NetworkInformation,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.ctype !== 0) {
      writer.uint32(8).int32(message.ctype);
    }
    if (message.IsOnline === true) {
      writer.uint32(16).bool(message.IsOnline);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NetworkInformation {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNetworkInformation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ctype = reader.int32() as any;
          break;
        case 2:
          message.IsOnline = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NetworkInformation {
    return {
      ctype: isSet(object.ctype) ? connectionTypeFromJSON(object.ctype) : 0,
      IsOnline: isSet(object.IsOnline) ? Boolean(object.IsOnline) : false,
    };
  },

  toJSON(message: NetworkInformation): unknown {
    const obj: any = {};
    message.ctype !== undefined &&
      (obj.ctype = connectionTypeToJSON(message.ctype));
    message.IsOnline !== undefined && (obj.IsOnline = message.IsOnline);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NetworkInformation>, I>>(
    object: I
  ): NetworkInformation {
    const message = createBaseNetworkInformation();
    message.ctype = object.ctype ?? 0;
    message.IsOnline = object.IsOnline ?? false;
    return message;
  },
};

export interface NetworkManager {
  Query(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<NetworkInformation>;
}

export class NetworkManagerClientImpl implements NetworkManager {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Query = this.Query.bind(this);
  }

  Query(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<NetworkInformation> {
    return this.rpc.unary(
      NetworkManagerQueryDesc,
      Empty.fromPartial(request),
      metadata
    );
  }
}

export const NetworkManagerDesc = {
  serviceName: "cloud.lazycat.localdevice.apis.NetworkManager",
};

export const NetworkManagerQueryDesc: UnaryMethodDefinitionish = {
  methodName: "Query",
  service: NetworkManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...NetworkInformation.decode(data),
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

type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
type Exact<P, I extends P> = P extends Builtin
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

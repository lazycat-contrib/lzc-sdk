/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { Empty } from "../google/protobuf/empty";
import { BrowserHeaders } from "browser-headers";
import { share } from "rxjs/operators";
import { Timestamp } from "../google/protobuf/timestamp";

export interface SessionInfo {
  uid: string;
  deviceId: string;
  when: Date | undefined;
}

export interface AppInfo {
  boxId: string;
  appId: string;
  appDomain: string;
}

function createBaseSessionInfo(): SessionInfo {
  return { uid: "", deviceId: "", when: undefined };
}

export const SessionInfo = {
  encode(
    message: SessionInfo,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.deviceId !== "") {
      writer.uint32(18).string(message.deviceId);
    }
    if (message.when !== undefined) {
      Timestamp.encode(
        toTimestamp(message.when),
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SessionInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSessionInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        case 2:
          message.deviceId = reader.string();
          break;
        case 3:
          message.when = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SessionInfo {
    return {
      uid: isSet(object.uid) ? String(object.uid) : "",
      deviceId: isSet(object.deviceId) ? String(object.deviceId) : "",
      when: isSet(object.when) ? fromJsonTimestamp(object.when) : undefined,
    };
  },

  toJSON(message: SessionInfo): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.deviceId !== undefined && (obj.deviceId = message.deviceId);
    message.when !== undefined && (obj.when = message.when.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SessionInfo>, I>>(
    object: I
  ): SessionInfo {
    const message = createBaseSessionInfo();
    message.uid = object.uid ?? "";
    message.deviceId = object.deviceId ?? "";
    message.when = object.when ?? undefined;
    return message;
  },
};

function createBaseAppInfo(): AppInfo {
  return { boxId: "", appId: "", appDomain: "" };
}

export const AppInfo = {
  encode(
    message: AppInfo,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.boxId !== "") {
      writer.uint32(10).string(message.boxId);
    }
    if (message.appId !== "") {
      writer.uint32(18).string(message.appId);
    }
    if (message.appDomain !== "") {
      writer.uint32(26).string(message.appDomain);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.boxId = reader.string();
          break;
        case 2:
          message.appId = reader.string();
          break;
        case 3:
          message.appDomain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppInfo {
    return {
      boxId: isSet(object.boxId) ? String(object.boxId) : "",
      appId: isSet(object.appId) ? String(object.appId) : "",
      appDomain: isSet(object.appDomain) ? String(object.appDomain) : "",
    };
  },

  toJSON(message: AppInfo): unknown {
    const obj: any = {};
    message.boxId !== undefined && (obj.boxId = message.boxId);
    message.appId !== undefined && (obj.appId = message.appId);
    message.appDomain !== undefined && (obj.appDomain = message.appDomain);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AppInfo>, I>>(object: I): AppInfo {
    const message = createBaseAppInfo();
    message.boxId = object.boxId ?? "";
    message.appId = object.appId ?? "";
    message.appDomain = object.appDomain ?? "";
    return message;
  },
};

/** this service is provied by frontend server, backend code shouldn't use it. */
export interface BrowserOnlyProxy {
  /** 查询当前登陆用户对应信息 */
  QuerySessionInfo(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<SessionInfo>;
  /** 查询当前访问的lzcapp对应信息 */
  QueryAppInfo(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<AppInfo>;
  /** 对devices.proto:_PairAllDeivces的自动封装 */
  PairAllDevices(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Observable<Empty>;
}

export class BrowserOnlyProxyClientImpl implements BrowserOnlyProxy {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.QuerySessionInfo = this.QuerySessionInfo.bind(this);
    this.QueryAppInfo = this.QueryAppInfo.bind(this);
    this.PairAllDevices = this.PairAllDevices.bind(this);
  }

  QuerySessionInfo(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<SessionInfo> {
    return this.rpc.unary(
      BrowserOnlyProxyQuerySessionInfoDesc,
      Empty.fromPartial(request),
      metadata
    );
  }

  QueryAppInfo(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<AppInfo> {
    return this.rpc.unary(
      BrowserOnlyProxyQueryAppInfoDesc,
      Empty.fromPartial(request),
      metadata
    );
  }

  PairAllDevices(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Observable<Empty> {
    return this.rpc.invoke(
      BrowserOnlyProxyPairAllDevicesDesc,
      Empty.fromPartial(request),
      metadata
    );
  }
}

export const BrowserOnlyProxyDesc = {
  serviceName: "cloud.lazycat.apis.BrowserOnlyProxy",
};

export const BrowserOnlyProxyQuerySessionInfoDesc: UnaryMethodDefinitionish = {
  methodName: "QuerySessionInfo",
  service: BrowserOnlyProxyDesc,
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
        ...SessionInfo.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const BrowserOnlyProxyQueryAppInfoDesc: UnaryMethodDefinitionish = {
  methodName: "QueryAppInfo",
  service: BrowserOnlyProxyDesc,
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
        ...AppInfo.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const BrowserOnlyProxyPairAllDevicesDesc: UnaryMethodDefinitionish = {
  methodName: "PairAllDevices",
  service: BrowserOnlyProxyDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...Empty.decode(data),
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
  invoke<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined
  ): Observable<any>;
}

export class GrpcWebImpl {
  private host: string;
  private options: {
    transport?: grpc.TransportFactory;
    streamingTransport?: grpc.TransportFactory;
    debug?: boolean;
    metadata?: grpc.Metadata;
  };

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory;
      streamingTransport?: grpc.TransportFactory;
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

  invoke<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined
  ): Observable<any> {
    // Status Response Codes (https://developers.google.com/maps-booking/reference/grpc-api/status_codes)
    const upStreamCodes = [2, 4, 8, 9, 10, 13, 14, 15];
    const DEFAULT_TIMEOUT_TIME: number = 3_000;
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata =
      metadata && this.options.metadata
        ? new BrowserHeaders({
            ...this.options?.metadata.headersMap,
            ...metadata?.headersMap,
          })
        : metadata || this.options.metadata;
    return new Observable((observer) => {
      const upStream = () => {
        const client = grpc.invoke(methodDesc, {
          host: this.host,
          request,
          transport: this.options.streamingTransport || this.options.transport,
          metadata: maybeCombinedMetadata,
          debug: this.options.debug,
          onMessage: (next) => observer.next(next),
          onEnd: (code: grpc.Code, message: string) => {
            if (code === 0) {
              observer.complete();
            } else if (upStreamCodes.includes(code)) {
              setTimeout(upStream, DEFAULT_TIMEOUT_TIME);
            } else {
              observer.error(new Error(`Error ${code} ${message}`));
            }
          },
        });
        observer.add(() => client.close());
      };
      upStream();
    }).pipe(share());
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

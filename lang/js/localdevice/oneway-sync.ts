/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";
import { Empty } from "../google/protobuf/empty";
import { Timestamp } from "../google/protobuf/timestamp";

export enum SyncStatus {
  /** Ready - 就绪 */
  Ready = 0,
  /** Syncing - 同步中 */
  Syncing = 1,
  /** Fail - 失败 */
  Fail = 2,
  /** UnRef - 失联 */
  UnRef = 3,
  /** Finish - 完成 */
  Finish = 4,
  UNRECOGNIZED = -1,
}

export function syncStatusFromJSON(object: any): SyncStatus {
  switch (object) {
    case 0:
    case "Ready":
      return SyncStatus.Ready;
    case 1:
    case "Syncing":
      return SyncStatus.Syncing;
    case 2:
    case "Fail":
      return SyncStatus.Fail;
    case 3:
    case "UnRef":
      return SyncStatus.UnRef;
    case 4:
    case "Finish":
      return SyncStatus.Finish;
    case -1:
    case "UNRECOGNIZED":
    default:
      return SyncStatus.UNRECOGNIZED;
  }
}

export function syncStatusToJSON(object: SyncStatus): string {
  switch (object) {
    case SyncStatus.Ready:
      return "Ready";
    case SyncStatus.Syncing:
      return "Syncing";
    case SyncStatus.Fail:
      return "Fail";
    case SyncStatus.UnRef:
      return "UnRef";
    case SyncStatus.Finish:
      return "Finish";
    case SyncStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface RequestSyncDir {
  dir: string;
}

export interface SyncDirOwner {
  boxId: string;
  uid: string;
}

export interface SyncDir {
  /** 需要同步的目录 */
  dir: string;
  /** 当前同步状态 */
  status: SyncStatus;
  /** 最后变化时间 */
  lastModifiedTime:
    | Date
    | undefined;
  /** 所属者(属于哪个盒子和用户) */
  owner: SyncDirOwner | undefined;
}

export interface ResponseListSyncDir {
  dirs: SyncDir[];
}

function createBaseRequestSyncDir(): RequestSyncDir {
  return { dir: "" };
}

export const RequestSyncDir = {
  encode(message: RequestSyncDir, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.dir !== "") {
      writer.uint32(10).string(message.dir);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RequestSyncDir {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRequestSyncDir();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.dir = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RequestSyncDir {
    return { dir: isSet(object.dir) ? String(object.dir) : "" };
  },

  toJSON(message: RequestSyncDir): unknown {
    const obj: any = {};
    if (message.dir !== "") {
      obj.dir = message.dir;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RequestSyncDir>, I>>(base?: I): RequestSyncDir {
    return RequestSyncDir.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RequestSyncDir>, I>>(object: I): RequestSyncDir {
    const message = createBaseRequestSyncDir();
    message.dir = object.dir ?? "";
    return message;
  },
};

function createBaseSyncDirOwner(): SyncDirOwner {
  return { boxId: "", uid: "" };
}

export const SyncDirOwner = {
  encode(message: SyncDirOwner, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.boxId !== "") {
      writer.uint32(10).string(message.boxId);
    }
    if (message.uid !== "") {
      writer.uint32(18).string(message.uid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SyncDirOwner {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSyncDirOwner();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.boxId = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.uid = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SyncDirOwner {
    return { boxId: isSet(object.boxId) ? String(object.boxId) : "", uid: isSet(object.uid) ? String(object.uid) : "" };
  },

  toJSON(message: SyncDirOwner): unknown {
    const obj: any = {};
    if (message.boxId !== "") {
      obj.boxId = message.boxId;
    }
    if (message.uid !== "") {
      obj.uid = message.uid;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SyncDirOwner>, I>>(base?: I): SyncDirOwner {
    return SyncDirOwner.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SyncDirOwner>, I>>(object: I): SyncDirOwner {
    const message = createBaseSyncDirOwner();
    message.boxId = object.boxId ?? "";
    message.uid = object.uid ?? "";
    return message;
  },
};

function createBaseSyncDir(): SyncDir {
  return { dir: "", status: 0, lastModifiedTime: undefined, owner: undefined };
}

export const SyncDir = {
  encode(message: SyncDir, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.dir !== "") {
      writer.uint32(10).string(message.dir);
    }
    if (message.status !== 0) {
      writer.uint32(16).int32(message.status);
    }
    if (message.lastModifiedTime !== undefined) {
      Timestamp.encode(toTimestamp(message.lastModifiedTime), writer.uint32(26).fork()).ldelim();
    }
    if (message.owner !== undefined) {
      SyncDirOwner.encode(message.owner, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SyncDir {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSyncDir();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.dir = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.lastModifiedTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.owner = SyncDirOwner.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SyncDir {
    return {
      dir: isSet(object.dir) ? String(object.dir) : "",
      status: isSet(object.status) ? syncStatusFromJSON(object.status) : 0,
      lastModifiedTime: isSet(object.lastModifiedTime) ? fromJsonTimestamp(object.lastModifiedTime) : undefined,
      owner: isSet(object.owner) ? SyncDirOwner.fromJSON(object.owner) : undefined,
    };
  },

  toJSON(message: SyncDir): unknown {
    const obj: any = {};
    if (message.dir !== "") {
      obj.dir = message.dir;
    }
    if (message.status !== 0) {
      obj.status = syncStatusToJSON(message.status);
    }
    if (message.lastModifiedTime !== undefined) {
      obj.lastModifiedTime = message.lastModifiedTime.toISOString();
    }
    if (message.owner !== undefined) {
      obj.owner = SyncDirOwner.toJSON(message.owner);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SyncDir>, I>>(base?: I): SyncDir {
    return SyncDir.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SyncDir>, I>>(object: I): SyncDir {
    const message = createBaseSyncDir();
    message.dir = object.dir ?? "";
    message.status = object.status ?? 0;
    message.lastModifiedTime = object.lastModifiedTime ?? undefined;
    message.owner = (object.owner !== undefined && object.owner !== null)
      ? SyncDirOwner.fromPartial(object.owner)
      : undefined;
    return message;
  },
};

function createBaseResponseListSyncDir(): ResponseListSyncDir {
  return { dirs: [] };
}

export const ResponseListSyncDir = {
  encode(message: ResponseListSyncDir, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.dirs) {
      SyncDir.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ResponseListSyncDir {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResponseListSyncDir();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.dirs.push(SyncDir.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ResponseListSyncDir {
    return { dirs: Array.isArray(object?.dirs) ? object.dirs.map((e: any) => SyncDir.fromJSON(e)) : [] };
  },

  toJSON(message: ResponseListSyncDir): unknown {
    const obj: any = {};
    if (message.dirs?.length) {
      obj.dirs = message.dirs.map((e) => SyncDir.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ResponseListSyncDir>, I>>(base?: I): ResponseListSyncDir {
    return ResponseListSyncDir.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ResponseListSyncDir>, I>>(object: I): ResponseListSyncDir {
    const message = createBaseResponseListSyncDir();
    message.dirs = object.dirs?.map((e) => SyncDir.fromPartial(e)) || [];
    return message;
  },
};

export interface OnewaySync {
  /** 列出所有符合来源微服和用户的同步目录j */
  ListSyncDir(
    request: DeepPartial<RequestSyncDir>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ResponseListSyncDir>;
  /** 新建一个同步目录 */
  NewSyncDir(
    request: DeepPartial<RequestSyncDir>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<SyncDir>;
  /** 主动做一次同步，目标是来源微服和来源用户的同步目录 */
  DoSync(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<SyncDir>;
}

export class OnewaySyncClientImpl implements OnewaySync {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ListSyncDir = this.ListSyncDir.bind(this);
    this.NewSyncDir = this.NewSyncDir.bind(this);
    this.DoSync = this.DoSync.bind(this);
  }

  ListSyncDir(
    request: DeepPartial<RequestSyncDir>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ResponseListSyncDir> {
    return this.rpc.unary(OnewaySyncListSyncDirDesc, RequestSyncDir.fromPartial(request), metadata, abortSignal);
  }

  NewSyncDir(
    request: DeepPartial<RequestSyncDir>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<SyncDir> {
    return this.rpc.unary(OnewaySyncNewSyncDirDesc, RequestSyncDir.fromPartial(request), metadata, abortSignal);
  }

  DoSync(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<SyncDir> {
    return this.rpc.unary(OnewaySyncDoSyncDesc, Empty.fromPartial(request), metadata, abortSignal);
  }
}

export const OnewaySyncDesc = { serviceName: "cloud.lazycat.apis.localdevice.OnewaySync" };

export const OnewaySyncListSyncDirDesc: UnaryMethodDefinitionish = {
  methodName: "ListSyncDir",
  service: OnewaySyncDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return RequestSyncDir.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ResponseListSyncDir.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const OnewaySyncNewSyncDirDesc: UnaryMethodDefinitionish = {
  methodName: "NewSyncDir",
  service: OnewaySyncDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return RequestSyncDir.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = SyncDir.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const OnewaySyncDoSyncDesc: UnaryMethodDefinitionish = {
  methodName: "DoSync",
  service: OnewaySyncDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = SyncDir.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

interface UnaryMethodDefinitionishR extends grpc.UnaryMethodDefinition<any, any> {
  requestStream: any;
  responseStream: any;
}

type UnaryMethodDefinitionish = UnaryMethodDefinitionishR;

interface Rpc {
  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined,
    abortSignal?: AbortSignal,
  ): Promise<any>;
}

export class GrpcWebImpl {
  private host: string;
  private options: {
    transport?: grpc.TransportFactory;

    debug?: boolean;
    metadata?: grpc.Metadata;
    upStreamRetryCodes?: number[];
  };

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory;

      debug?: boolean;
      metadata?: grpc.Metadata;
      upStreamRetryCodes?: number[];
    },
  ) {
    this.host = host;
    this.options = options;
  }

  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined,
    abortSignal?: AbortSignal,
  ): Promise<any> {
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata = metadata && this.options.metadata
      ? new BrowserHeaders({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
      : metadata ?? this.options.metadata;
    return new Promise((resolve, reject) => {
      const client = grpc.unary(methodDesc, {
        request,
        host: this.host,
        metadata: maybeCombinedMetadata ?? {},
        ...(this.options.transport !== undefined ? { transport: this.options.transport } : {}),
        debug: this.options.debug ?? false,
        onEnd: function (response) {
          if (response.status === grpc.Code.OK) {
            resolve(response.message!.toObject());
          } else {
            const err = new GrpcWebError(response.statusMessage, response.status, response.trailers);
            reject(err);
          }
        },
      });

      if (abortSignal) {
        abortSignal.addEventListener("abort", () => {
          client.close();
          reject(abortSignal.reason);
        });
      }
    });
  }
}

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export class GrpcWebError extends tsProtoGlobalThis.Error {
  constructor(message: string, public code: grpc.Code, public metadata: grpc.Metadata) {
    super(message);
  }
}

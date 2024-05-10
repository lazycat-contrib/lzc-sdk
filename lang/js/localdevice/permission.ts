/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";
import { Empty } from "../google/protobuf/empty";
import { Timestamp } from "../google/protobuf/timestamp";

export enum Permission {
  /** CLIPBOARD - 剪贴板 */
  CLIPBOARD = 0,
  /** DEVICE_INFO - 设备信息 */
  DEVICE_INFO = 1,
  /** OPEN_DIALOG - 弹出对话框 */
  OPEN_DIALOG = 2,
  /** OPEN_THIRD_PARTY_APP - 使用第三方app打开文件 */
  OPEN_THIRD_PARTY_APP = 3,
  /** PIN_APP - 发送应用到桌面快捷方式 */
  PIN_APP = 4,
  /** NETWORK_INFO - 网络信息 */
  NETWORK_INFO = 5,
  /** PHOTO_LIBRARY - 相册 */
  PHOTO_LIBRARY = 6,
  /** DOCUMENT - 文件 */
  DOCUMENT = 7,
  /** USER_CONFIG - 用户配置 */
  USER_CONFIG = 8,
  UNRECOGNIZED = -1,
}

export function permissionFromJSON(object: any): Permission {
  switch (object) {
    case 0:
    case "CLIPBOARD":
      return Permission.CLIPBOARD;
    case 1:
    case "DEVICE_INFO":
      return Permission.DEVICE_INFO;
    case 2:
    case "OPEN_DIALOG":
      return Permission.OPEN_DIALOG;
    case 3:
    case "OPEN_THIRD_PARTY_APP":
      return Permission.OPEN_THIRD_PARTY_APP;
    case 4:
    case "PIN_APP":
      return Permission.PIN_APP;
    case 5:
    case "NETWORK_INFO":
      return Permission.NETWORK_INFO;
    case 6:
    case "PHOTO_LIBRARY":
      return Permission.PHOTO_LIBRARY;
    case 7:
    case "DOCUMENT":
      return Permission.DOCUMENT;
    case 8:
    case "USER_CONFIG":
      return Permission.USER_CONFIG;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Permission.UNRECOGNIZED;
  }
}

export function permissionToJSON(object: Permission): string {
  switch (object) {
    case Permission.CLIPBOARD:
      return "CLIPBOARD";
    case Permission.DEVICE_INFO:
      return "DEVICE_INFO";
    case Permission.OPEN_DIALOG:
      return "OPEN_DIALOG";
    case Permission.OPEN_THIRD_PARTY_APP:
      return "OPEN_THIRD_PARTY_APP";
    case Permission.PIN_APP:
      return "PIN_APP";
    case Permission.NETWORK_INFO:
      return "NETWORK_INFO";
    case Permission.PHOTO_LIBRARY:
      return "PHOTO_LIBRARY";
    case Permission.DOCUMENT:
      return "DOCUMENT";
    case Permission.USER_CONFIG:
      return "USER_CONFIG";
    case Permission.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface PermissionRequest {
  permission: Permission;
}

export interface PermissionReply {
  result: boolean;
}

export interface ListPermissionsReply {
  /** map<Permission, bool> */
  result: { [key: number]: boolean };
}

export interface ListPermissionsReply_ResultEntry {
  key: number;
  value: boolean;
}

export interface RequestAuthTokenRequest {
  boxCert: Uint8Array;
  appCert: Uint8Array;
  /** signature of $LAZYCAT_APPID.$LAZYCAT_BOX_DOMAIN */
  signature: Uint8Array;
}

export interface RequestAuthTokenResponse {
  token: string;
  deadline: Date | undefined;
}

function createBasePermissionRequest(): PermissionRequest {
  return { permission: 0 };
}

export const PermissionRequest = {
  encode(message: PermissionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.permission !== 0) {
      writer.uint32(8).int32(message.permission);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PermissionRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePermissionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.permission = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PermissionRequest {
    return { permission: isSet(object.permission) ? permissionFromJSON(object.permission) : 0 };
  },

  toJSON(message: PermissionRequest): unknown {
    const obj: any = {};
    if (message.permission !== 0) {
      obj.permission = permissionToJSON(message.permission);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PermissionRequest>, I>>(base?: I): PermissionRequest {
    return PermissionRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PermissionRequest>, I>>(object: I): PermissionRequest {
    const message = createBasePermissionRequest();
    message.permission = object.permission ?? 0;
    return message;
  },
};

function createBasePermissionReply(): PermissionReply {
  return { result: false };
}

export const PermissionReply = {
  encode(message: PermissionReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.result === true) {
      writer.uint32(8).bool(message.result);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PermissionReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePermissionReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.result = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PermissionReply {
    return { result: isSet(object.result) ? Boolean(object.result) : false };
  },

  toJSON(message: PermissionReply): unknown {
    const obj: any = {};
    if (message.result === true) {
      obj.result = message.result;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PermissionReply>, I>>(base?: I): PermissionReply {
    return PermissionReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PermissionReply>, I>>(object: I): PermissionReply {
    const message = createBasePermissionReply();
    message.result = object.result ?? false;
    return message;
  },
};

function createBaseListPermissionsReply(): ListPermissionsReply {
  return { result: {} };
}

export const ListPermissionsReply = {
  encode(message: ListPermissionsReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    Object.entries(message.result).forEach(([key, value]) => {
      ListPermissionsReply_ResultEntry.encode({ key: key as any, value }, writer.uint32(10).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPermissionsReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPermissionsReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          const entry1 = ListPermissionsReply_ResultEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.result[entry1.key] = entry1.value;
          }
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListPermissionsReply {
    return {
      result: isObject(object.result)
        ? Object.entries(object.result).reduce<{ [key: number]: boolean }>((acc, [key, value]) => {
          acc[Number(key)] = Boolean(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: ListPermissionsReply): unknown {
    const obj: any = {};
    if (message.result) {
      const entries = Object.entries(message.result);
      if (entries.length > 0) {
        obj.result = {};
        entries.forEach(([k, v]) => {
          obj.result[k] = v;
        });
      }
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListPermissionsReply>, I>>(base?: I): ListPermissionsReply {
    return ListPermissionsReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListPermissionsReply>, I>>(object: I): ListPermissionsReply {
    const message = createBaseListPermissionsReply();
    message.result = Object.entries(object.result ?? {}).reduce<{ [key: number]: boolean }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[Number(key)] = Boolean(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseListPermissionsReply_ResultEntry(): ListPermissionsReply_ResultEntry {
  return { key: 0, value: false };
}

export const ListPermissionsReply_ResultEntry = {
  encode(message: ListPermissionsReply_ResultEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== 0) {
      writer.uint32(8).int32(message.key);
    }
    if (message.value === true) {
      writer.uint32(16).bool(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPermissionsReply_ResultEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPermissionsReply_ResultEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.key = reader.int32();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.value = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListPermissionsReply_ResultEntry {
    return {
      key: isSet(object.key) ? Number(object.key) : 0,
      value: isSet(object.value) ? Boolean(object.value) : false,
    };
  },

  toJSON(message: ListPermissionsReply_ResultEntry): unknown {
    const obj: any = {};
    if (message.key !== 0) {
      obj.key = Math.round(message.key);
    }
    if (message.value === true) {
      obj.value = message.value;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListPermissionsReply_ResultEntry>, I>>(
    base?: I,
  ): ListPermissionsReply_ResultEntry {
    return ListPermissionsReply_ResultEntry.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListPermissionsReply_ResultEntry>, I>>(
    object: I,
  ): ListPermissionsReply_ResultEntry {
    const message = createBaseListPermissionsReply_ResultEntry();
    message.key = object.key ?? 0;
    message.value = object.value ?? false;
    return message;
  },
};

function createBaseRequestAuthTokenRequest(): RequestAuthTokenRequest {
  return { boxCert: new Uint8Array(0), appCert: new Uint8Array(0), signature: new Uint8Array(0) };
}

export const RequestAuthTokenRequest = {
  encode(message: RequestAuthTokenRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.boxCert.length !== 0) {
      writer.uint32(10).bytes(message.boxCert);
    }
    if (message.appCert.length !== 0) {
      writer.uint32(18).bytes(message.appCert);
    }
    if (message.signature.length !== 0) {
      writer.uint32(26).bytes(message.signature);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RequestAuthTokenRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRequestAuthTokenRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.boxCert = reader.bytes();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.appCert = reader.bytes();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.signature = reader.bytes();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RequestAuthTokenRequest {
    return {
      boxCert: isSet(object.boxCert) ? bytesFromBase64(object.boxCert) : new Uint8Array(0),
      appCert: isSet(object.appCert) ? bytesFromBase64(object.appCert) : new Uint8Array(0),
      signature: isSet(object.signature) ? bytesFromBase64(object.signature) : new Uint8Array(0),
    };
  },

  toJSON(message: RequestAuthTokenRequest): unknown {
    const obj: any = {};
    if (message.boxCert.length !== 0) {
      obj.boxCert = base64FromBytes(message.boxCert);
    }
    if (message.appCert.length !== 0) {
      obj.appCert = base64FromBytes(message.appCert);
    }
    if (message.signature.length !== 0) {
      obj.signature = base64FromBytes(message.signature);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RequestAuthTokenRequest>, I>>(base?: I): RequestAuthTokenRequest {
    return RequestAuthTokenRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RequestAuthTokenRequest>, I>>(object: I): RequestAuthTokenRequest {
    const message = createBaseRequestAuthTokenRequest();
    message.boxCert = object.boxCert ?? new Uint8Array(0);
    message.appCert = object.appCert ?? new Uint8Array(0);
    message.signature = object.signature ?? new Uint8Array(0);
    return message;
  },
};

function createBaseRequestAuthTokenResponse(): RequestAuthTokenResponse {
  return { token: "", deadline: undefined };
}

export const RequestAuthTokenResponse = {
  encode(message: RequestAuthTokenResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.deadline !== undefined) {
      Timestamp.encode(toTimestamp(message.deadline), writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RequestAuthTokenResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRequestAuthTokenResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.token = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.deadline = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RequestAuthTokenResponse {
    return {
      token: isSet(object.token) ? String(object.token) : "",
      deadline: isSet(object.deadline) ? fromJsonTimestamp(object.deadline) : undefined,
    };
  },

  toJSON(message: RequestAuthTokenResponse): unknown {
    const obj: any = {};
    if (message.token !== "") {
      obj.token = message.token;
    }
    if (message.deadline !== undefined) {
      obj.deadline = message.deadline.toISOString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RequestAuthTokenResponse>, I>>(base?: I): RequestAuthTokenResponse {
    return RequestAuthTokenResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RequestAuthTokenResponse>, I>>(object: I): RequestAuthTokenResponse {
    const message = createBaseRequestAuthTokenResponse();
    message.token = object.token ?? "";
    message.deadline = object.deadline ?? undefined;
    return message;
  },
};

/** 权限管理 */
export interface PermissionManager {
  /** 检测权限 */
  GetPermission(
    request: DeepPartial<PermissionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<PermissionReply>;
  /** 申请权限（会弹出对话框让用户决定是否同意） */
  RequestPermission(
    request: DeepPartial<PermissionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<PermissionReply>;
  /** 列举权限 */
  ListPermissions(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListPermissionsReply>;
  /** 申请Token */
  RequestAuthToken(
    request: DeepPartial<RequestAuthTokenRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<RequestAuthTokenResponse>;
}

export class PermissionManagerClientImpl implements PermissionManager {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.GetPermission = this.GetPermission.bind(this);
    this.RequestPermission = this.RequestPermission.bind(this);
    this.ListPermissions = this.ListPermissions.bind(this);
    this.RequestAuthToken = this.RequestAuthToken.bind(this);
  }

  GetPermission(
    request: DeepPartial<PermissionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<PermissionReply> {
    return this.rpc.unary(
      PermissionManagerGetPermissionDesc,
      PermissionRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  RequestPermission(
    request: DeepPartial<PermissionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<PermissionReply> {
    return this.rpc.unary(
      PermissionManagerRequestPermissionDesc,
      PermissionRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  ListPermissions(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListPermissionsReply> {
    return this.rpc.unary(PermissionManagerListPermissionsDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  RequestAuthToken(
    request: DeepPartial<RequestAuthTokenRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<RequestAuthTokenResponse> {
    return this.rpc.unary(
      PermissionManagerRequestAuthTokenDesc,
      RequestAuthTokenRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }
}

export const PermissionManagerDesc = { serviceName: "cloud.lazycat.apis.localdevice.PermissionManager" };

export const PermissionManagerGetPermissionDesc: UnaryMethodDefinitionish = {
  methodName: "GetPermission",
  service: PermissionManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return PermissionRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = PermissionReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const PermissionManagerRequestPermissionDesc: UnaryMethodDefinitionish = {
  methodName: "RequestPermission",
  service: PermissionManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return PermissionRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = PermissionReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const PermissionManagerListPermissionsDesc: UnaryMethodDefinitionish = {
  methodName: "ListPermissions",
  service: PermissionManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ListPermissionsReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const PermissionManagerRequestAuthTokenDesc: UnaryMethodDefinitionish = {
  methodName: "RequestAuthToken",
  service: PermissionManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return RequestAuthTokenRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = RequestAuthTokenResponse.decode(data);
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

function bytesFromBase64(b64: string): Uint8Array {
  if (tsProtoGlobalThis.Buffer) {
    return Uint8Array.from(tsProtoGlobalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = tsProtoGlobalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (tsProtoGlobalThis.Buffer) {
    return tsProtoGlobalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte));
    });
    return tsProtoGlobalThis.btoa(bin.join(""));
  }
}

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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export class GrpcWebError extends tsProtoGlobalThis.Error {
  constructor(message: string, public code: grpc.Code, public metadata: grpc.Metadata) {
    super(message);
  }
}

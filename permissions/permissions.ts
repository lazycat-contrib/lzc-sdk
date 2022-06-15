/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { BrowserHeaders } from "browser-headers";

export enum TokenType {
  HttpCookie = 0,
  HttpHeader = 1,
  UNRECOGNIZED = -1,
}

export function tokenTypeFromJSON(object: any): TokenType {
  switch (object) {
    case 0:
    case "HttpCookie":
      return TokenType.HttpCookie;
    case 1:
    case "HttpHeader":
      return TokenType.HttpHeader;
    case -1:
    case "UNRECOGNIZED":
    default:
      return TokenType.UNRECOGNIZED;
  }
}

export function tokenTypeToJSON(object: TokenType): string {
  switch (object) {
    case TokenType.HttpCookie:
      return "HttpCookie";
    case TokenType.HttpHeader:
      return "HttpHeader";
    case TokenType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface PermissionDesc {
  /** default to current box. */
  targetId: string;
  /** should empty or equal userinfo.uid when request from browser */
  userId: string;
  /** cloud.lazycat.apis.XX || cloud.lazycat.devices.apis.YY */
  serviceName: string;
  /**  */
  methodName: string;
}

export interface PermissionToken {
  keyType: TokenType;
  keyName: string;
  keyValue: string;
}

function createBasePermissionDesc(): PermissionDesc {
  return { targetId: "", userId: "", serviceName: "", methodName: "" };
}

export const PermissionDesc = {
  encode(
    message: PermissionDesc,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.targetId !== "") {
      writer.uint32(10).string(message.targetId);
    }
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    if (message.serviceName !== "") {
      writer.uint32(26).string(message.serviceName);
    }
    if (message.methodName !== "") {
      writer.uint32(34).string(message.methodName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PermissionDesc {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePermissionDesc();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.targetId = reader.string();
          break;
        case 2:
          message.userId = reader.string();
          break;
        case 3:
          message.serviceName = reader.string();
          break;
        case 4:
          message.methodName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PermissionDesc {
    return {
      targetId: isSet(object.targetId) ? String(object.targetId) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      serviceName: isSet(object.serviceName) ? String(object.serviceName) : "",
      methodName: isSet(object.methodName) ? String(object.methodName) : "",
    };
  },

  toJSON(message: PermissionDesc): unknown {
    const obj: any = {};
    message.targetId !== undefined && (obj.targetId = message.targetId);
    message.userId !== undefined && (obj.userId = message.userId);
    message.serviceName !== undefined &&
      (obj.serviceName = message.serviceName);
    message.methodName !== undefined && (obj.methodName = message.methodName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PermissionDesc>, I>>(
    object: I
  ): PermissionDesc {
    const message = createBasePermissionDesc();
    message.targetId = object.targetId ?? "";
    message.userId = object.userId ?? "";
    message.serviceName = object.serviceName ?? "";
    message.methodName = object.methodName ?? "";
    return message;
  },
};

function createBasePermissionToken(): PermissionToken {
  return { keyType: 0, keyName: "", keyValue: "" };
}

export const PermissionToken = {
  encode(
    message: PermissionToken,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.keyType !== 0) {
      writer.uint32(8).int32(message.keyType);
    }
    if (message.keyName !== "") {
      writer.uint32(18).string(message.keyName);
    }
    if (message.keyValue !== "") {
      writer.uint32(26).string(message.keyValue);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PermissionToken {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePermissionToken();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.keyType = reader.int32() as any;
          break;
        case 2:
          message.keyName = reader.string();
          break;
        case 3:
          message.keyValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PermissionToken {
    return {
      keyType: isSet(object.keyType) ? tokenTypeFromJSON(object.keyType) : 0,
      keyName: isSet(object.keyName) ? String(object.keyName) : "",
      keyValue: isSet(object.keyValue) ? String(object.keyValue) : "",
    };
  },

  toJSON(message: PermissionToken): unknown {
    const obj: any = {};
    message.keyType !== undefined &&
      (obj.keyType = tokenTypeToJSON(message.keyType));
    message.keyName !== undefined && (obj.keyName = message.keyName);
    message.keyValue !== undefined && (obj.keyValue = message.keyValue);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PermissionToken>, I>>(
    object: I
  ): PermissionToken {
    const message = createBasePermissionToken();
    message.keyType = object.keyType ?? 0;
    message.keyName = object.keyName ?? "";
    message.keyValue = object.keyValue ?? "";
    return message;
  },
};

/** this service is provied by frontend server, backend code shouldn't use it. */
export interface PermissionManager {
  RequestPermission(
    request: DeepPartial<PermissionDesc>,
    metadata?: grpc.Metadata
  ): Promise<PermissionToken>;
}

export class PermissionManagerClientImpl implements PermissionManager {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RequestPermission = this.RequestPermission.bind(this);
  }

  RequestPermission(
    request: DeepPartial<PermissionDesc>,
    metadata?: grpc.Metadata
  ): Promise<PermissionToken> {
    return this.rpc.unary(
      PermissionManagerRequestPermissionDesc,
      PermissionDesc.fromPartial(request),
      metadata
    );
  }
}

export const PermissionManagerDesc = {
  serviceName: "cloud.lazycat.apis.PermissionManager",
};

export const PermissionManagerRequestPermissionDesc: UnaryMethodDefinitionish =
  {
    methodName: "RequestPermission",
    service: PermissionManagerDesc,
    requestStream: false,
    responseStream: false,
    requestType: {
      serializeBinary() {
        return PermissionDesc.encode(this).finish();
      },
    } as any,
    responseType: {
      deserializeBinary(data: Uint8Array) {
        return {
          ...PermissionToken.decode(data),
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

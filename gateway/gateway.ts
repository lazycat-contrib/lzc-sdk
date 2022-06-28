/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../google/protobuf/empty";
import { BrowserHeaders } from "browser-headers";

export interface GatewayInfo {
  /** 域名信息 */
  baseDomain: string;
  /** 服务端口 */
  port: string;
  /** 是否使用TLS */
  isTls: boolean;
}

export interface GRPCServiceInfo {
  serviceName: string;
  backendAddr: string;
}

export interface RegisterReply {
  isOk: boolean;
  errorMessage: string;
}

function createBaseGatewayInfo(): GatewayInfo {
  return { baseDomain: "", port: "", isTls: false };
}

export const GatewayInfo = {
  encode(
    message: GatewayInfo,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.baseDomain !== "") {
      writer.uint32(10).string(message.baseDomain);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.isTls === true) {
      writer.uint32(24).bool(message.isTls);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GatewayInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGatewayInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.baseDomain = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.isTls = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GatewayInfo {
    return {
      baseDomain: isSet(object.baseDomain) ? String(object.baseDomain) : "",
      port: isSet(object.port) ? String(object.port) : "",
      isTls: isSet(object.isTls) ? Boolean(object.isTls) : false,
    };
  },

  toJSON(message: GatewayInfo): unknown {
    const obj: any = {};
    message.baseDomain !== undefined && (obj.baseDomain = message.baseDomain);
    message.port !== undefined && (obj.port = message.port);
    message.isTls !== undefined && (obj.isTls = message.isTls);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GatewayInfo>, I>>(
    object: I
  ): GatewayInfo {
    const message = createBaseGatewayInfo();
    message.baseDomain = object.baseDomain ?? "";
    message.port = object.port ?? "";
    message.isTls = object.isTls ?? false;
    return message;
  },
};

function createBaseGRPCServiceInfo(): GRPCServiceInfo {
  return { serviceName: "", backendAddr: "" };
}

export const GRPCServiceInfo = {
  encode(
    message: GRPCServiceInfo,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.serviceName !== "") {
      writer.uint32(10).string(message.serviceName);
    }
    if (message.backendAddr !== "") {
      writer.uint32(18).string(message.backendAddr);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GRPCServiceInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGRPCServiceInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.serviceName = reader.string();
          break;
        case 2:
          message.backendAddr = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GRPCServiceInfo {
    return {
      serviceName: isSet(object.serviceName) ? String(object.serviceName) : "",
      backendAddr: isSet(object.backendAddr) ? String(object.backendAddr) : "",
    };
  },

  toJSON(message: GRPCServiceInfo): unknown {
    const obj: any = {};
    message.serviceName !== undefined &&
      (obj.serviceName = message.serviceName);
    message.backendAddr !== undefined &&
      (obj.backendAddr = message.backendAddr);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GRPCServiceInfo>, I>>(
    object: I
  ): GRPCServiceInfo {
    const message = createBaseGRPCServiceInfo();
    message.serviceName = object.serviceName ?? "";
    message.backendAddr = object.backendAddr ?? "";
    return message;
  },
};

function createBaseRegisterReply(): RegisterReply {
  return { isOk: false, errorMessage: "" };
}

export const RegisterReply = {
  encode(
    message: RegisterReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.isOk === true) {
      writer.uint32(8).bool(message.isOk);
    }
    if (message.errorMessage !== "") {
      writer.uint32(18).string(message.errorMessage);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RegisterReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRegisterReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.isOk = reader.bool();
          break;
        case 2:
          message.errorMessage = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RegisterReply {
    return {
      isOk: isSet(object.isOk) ? Boolean(object.isOk) : false,
      errorMessage: isSet(object.errorMessage)
        ? String(object.errorMessage)
        : "",
    };
  },

  toJSON(message: RegisterReply): unknown {
    const obj: any = {};
    message.isOk !== undefined && (obj.isOk = message.isOk);
    message.errorMessage !== undefined &&
      (obj.errorMessage = message.errorMessage);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RegisterReply>, I>>(
    object: I
  ): RegisterReply {
    const message = createBaseRegisterReply();
    message.isOk = object.isOk ?? false;
    message.errorMessage = object.errorMessage ?? "";
    return message;
  },
};

export interface APIGateway {
  QueryGatewayInfo(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<GatewayInfo>;
  RegisterGRPCService(
    request: DeepPartial<GRPCServiceInfo>,
    metadata?: grpc.Metadata
  ): Promise<RegisterReply>;
}

export class APIGatewayClientImpl implements APIGateway {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.QueryGatewayInfo = this.QueryGatewayInfo.bind(this);
    this.RegisterGRPCService = this.RegisterGRPCService.bind(this);
  }

  QueryGatewayInfo(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<GatewayInfo> {
    return this.rpc.unary(
      APIGatewayQueryGatewayInfoDesc,
      Empty.fromPartial(request),
      metadata
    );
  }

  RegisterGRPCService(
    request: DeepPartial<GRPCServiceInfo>,
    metadata?: grpc.Metadata
  ): Promise<RegisterReply> {
    return this.rpc.unary(
      APIGatewayRegisterGRPCServiceDesc,
      GRPCServiceInfo.fromPartial(request),
      metadata
    );
  }
}

export const APIGatewayDesc = {
  serviceName: "cloud.lazycat.apis.APIGateway",
};

export const APIGatewayQueryGatewayInfoDesc: UnaryMethodDefinitionish = {
  methodName: "QueryGatewayInfo",
  service: APIGatewayDesc,
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
        ...GatewayInfo.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const APIGatewayRegisterGRPCServiceDesc: UnaryMethodDefinitionish = {
  methodName: "RegisterGRPCService",
  service: APIGatewayDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return GRPCServiceInfo.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...RegisterReply.decode(data),
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

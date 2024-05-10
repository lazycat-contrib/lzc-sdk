/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";
import { Empty } from "../google/protobuf/empty";

export interface DeviceInfo {
  OS: string;
  CPU: string;
  name: string;
  documentRootDir: string;
  /** 用于扫码或辅助登录标记操作端 */
  identityProvider?: string | undefined;
}

function createBaseDeviceInfo(): DeviceInfo {
  return { OS: "", CPU: "", name: "", documentRootDir: "", identityProvider: undefined };
}

export const DeviceInfo = {
  encode(message: DeviceInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.OS !== "") {
      writer.uint32(10).string(message.OS);
    }
    if (message.CPU !== "") {
      writer.uint32(18).string(message.CPU);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.documentRootDir !== "") {
      writer.uint32(34).string(message.documentRootDir);
    }
    if (message.identityProvider !== undefined) {
      writer.uint32(42).string(message.identityProvider);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeviceInfo {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeviceInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.OS = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.CPU = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.name = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.documentRootDir = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.identityProvider = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeviceInfo {
    return {
      OS: isSet(object.OS) ? String(object.OS) : "",
      CPU: isSet(object.CPU) ? String(object.CPU) : "",
      name: isSet(object.name) ? String(object.name) : "",
      documentRootDir: isSet(object.documentRootDir) ? String(object.documentRootDir) : "",
      identityProvider: isSet(object.identityProvider) ? String(object.identityProvider) : undefined,
    };
  },

  toJSON(message: DeviceInfo): unknown {
    const obj: any = {};
    if (message.OS !== "") {
      obj.OS = message.OS;
    }
    if (message.CPU !== "") {
      obj.CPU = message.CPU;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.documentRootDir !== "") {
      obj.documentRootDir = message.documentRootDir;
    }
    if (message.identityProvider !== undefined) {
      obj.identityProvider = message.identityProvider;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeviceInfo>, I>>(base?: I): DeviceInfo {
    return DeviceInfo.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeviceInfo>, I>>(object: I): DeviceInfo {
    const message = createBaseDeviceInfo();
    message.OS = object.OS ?? "";
    message.CPU = object.CPU ?? "";
    message.name = object.name ?? "";
    message.documentRootDir = object.documentRootDir ?? "";
    message.identityProvider = object.identityProvider ?? undefined;
    return message;
  },
};

export interface DeviceService {
  Query(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<DeviceInfo>;
}

export class DeviceServiceClientImpl implements DeviceService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Query = this.Query.bind(this);
  }

  Query(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<DeviceInfo> {
    return this.rpc.unary(DeviceServiceQueryDesc, Empty.fromPartial(request), metadata, abortSignal);
  }
}

export const DeviceServiceDesc = { serviceName: "cloud.lazycat.apis.localdevice.DeviceService" };

export const DeviceServiceQueryDesc: UnaryMethodDefinitionish = {
  methodName: "Query",
  service: DeviceServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = DeviceInfo.decode(data);
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export class GrpcWebError extends tsProtoGlobalThis.Error {
  constructor(message: string, public code: grpc.Code, public metadata: grpc.Metadata) {
    super(message);
  }
}

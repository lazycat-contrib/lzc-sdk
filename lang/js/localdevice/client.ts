/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Empty } from "../google/protobuf/empty";

export interface GetClientInfoResponse {
  /** 设备客户端版本名字 */
  versionName: string;
  /** 设备版本编码 */
  clientCode: number;
}

function createBaseGetClientInfoResponse(): GetClientInfoResponse {
  return { versionName: "", clientCode: 0 };
}

export const GetClientInfoResponse = {
  encode(message: GetClientInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.versionName !== "") {
      writer.uint32(10).string(message.versionName);
    }
    if (message.clientCode !== 0) {
      writer.uint32(16).int64(message.clientCode);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetClientInfoResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetClientInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.versionName = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.clientCode = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetClientInfoResponse {
    return {
      versionName: isSet(object.versionName) ? String(object.versionName) : "",
      clientCode: isSet(object.clientCode) ? Number(object.clientCode) : 0,
    };
  },

  toJSON(message: GetClientInfoResponse): unknown {
    const obj: any = {};
    if (message.versionName !== "") {
      obj.versionName = message.versionName;
    }
    if (message.clientCode !== 0) {
      obj.clientCode = Math.round(message.clientCode);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetClientInfoResponse>, I>>(base?: I): GetClientInfoResponse {
    return GetClientInfoResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetClientInfoResponse>, I>>(object: I): GetClientInfoResponse {
    const message = createBaseGetClientInfoResponse();
    message.versionName = object.versionName ?? "";
    message.clientCode = object.clientCode ?? 0;
    return message;
  },
};

export interface Client {
  GetClientInfo(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<GetClientInfoResponse>;
}

export class ClientClientImpl implements Client {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.GetClientInfo = this.GetClientInfo.bind(this);
  }

  GetClientInfo(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<GetClientInfoResponse> {
    return this.rpc.unary(ClientGetClientInfoDesc, Empty.fromPartial(request), metadata, abortSignal);
  }
}

export const ClientDesc = { serviceName: "cloud.lazycat.apis.localdevice.Client" };

export const ClientGetClientInfoDesc: UnaryMethodDefinitionish = {
  methodName: "GetClientInfo",
  service: ClientDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = GetClientInfoResponse.decode(data);
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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export class GrpcWebError extends tsProtoGlobalThis.Error {
  constructor(message: string, public code: grpc.Code, public metadata: grpc.Metadata) {
    super(message);
  }
}

/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";

export interface HashInfo {
  id: string;
  hash: string;
}

export interface DiffResponse {
  isDiff: boolean;
}

function createBaseHashInfo(): HashInfo {
  return { id: "", hash: "" };
}

export const HashInfo = {
  encode(message: HashInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.hash !== "") {
      writer.uint32(18).string(message.hash);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): HashInfo {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseHashInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.hash = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): HashInfo {
    return { id: isSet(object.id) ? String(object.id) : "", hash: isSet(object.hash) ? String(object.hash) : "" };
  },

  toJSON(message: HashInfo): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.hash !== "") {
      obj.hash = message.hash;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<HashInfo>, I>>(base?: I): HashInfo {
    return HashInfo.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<HashInfo>, I>>(object: I): HashInfo {
    const message = createBaseHashInfo();
    message.id = object.id ?? "";
    message.hash = object.hash ?? "";
    return message;
  },
};

function createBaseDiffResponse(): DiffResponse {
  return { isDiff: false };
}

export const DiffResponse = {
  encode(message: DiffResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.isDiff === true) {
      writer.uint32(8).bool(message.isDiff);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DiffResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDiffResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.isDiff = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DiffResponse {
    return { isDiff: isSet(object.isDiff) ? Boolean(object.isDiff) : false };
  },

  toJSON(message: DiffResponse): unknown {
    const obj: any = {};
    if (message.isDiff === true) {
      obj.isDiff = message.isDiff;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DiffResponse>, I>>(base?: I): DiffResponse {
    return DiffResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DiffResponse>, I>>(object: I): DiffResponse {
    const message = createBaseDiffResponse();
    message.isDiff = object.isDiff ?? false;
    return message;
  },
};

export interface Helper {
  Diff(request: DeepPartial<HashInfo>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<DiffResponse>;
}

export class HelperClientImpl implements Helper {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Diff = this.Diff.bind(this);
  }

  Diff(request: DeepPartial<HashInfo>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<DiffResponse> {
    return this.rpc.unary(HelperDiffDesc, HashInfo.fromPartial(request), metadata, abortSignal);
  }
}

export const HelperDesc = { serviceName: "cloud.lazycat.apis.localdevice.Helper" };

export const HelperDiffDesc: UnaryMethodDefinitionish = {
  methodName: "Diff",
  service: HelperDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return HashInfo.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = DiffResponse.decode(data);
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

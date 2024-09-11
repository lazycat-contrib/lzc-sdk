/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";

export interface OpenFileRequest {
  boxName: string;
  filePath: string;
}

export interface OpenFileReply {
}

export interface EnableAutoMountRequest {
  boxName: string;
}

export interface EnableAutoMountReply {
}

export interface IsEnableAutoMountRequest {
  boxName: string;
}

export interface IsEnableAutoMountReply {
  enable: boolean;
}

function createBaseOpenFileRequest(): OpenFileRequest {
  return { boxName: "", filePath: "" };
}

export const OpenFileRequest = {
  encode(message: OpenFileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.boxName !== "") {
      writer.uint32(10).string(message.boxName);
    }
    if (message.filePath !== "") {
      writer.uint32(18).string(message.filePath);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenFileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenFileRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.boxName = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.filePath = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OpenFileRequest {
    return {
      boxName: isSet(object.boxName) ? String(object.boxName) : "",
      filePath: isSet(object.filePath) ? String(object.filePath) : "",
    };
  },

  toJSON(message: OpenFileRequest): unknown {
    const obj: any = {};
    if (message.boxName !== "") {
      obj.boxName = message.boxName;
    }
    if (message.filePath !== "") {
      obj.filePath = message.filePath;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenFileRequest>, I>>(base?: I): OpenFileRequest {
    return OpenFileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OpenFileRequest>, I>>(object: I): OpenFileRequest {
    const message = createBaseOpenFileRequest();
    message.boxName = object.boxName ?? "";
    message.filePath = object.filePath ?? "";
    return message;
  },
};

function createBaseOpenFileReply(): OpenFileReply {
  return {};
}

export const OpenFileReply = {
  encode(_: OpenFileReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenFileReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenFileReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): OpenFileReply {
    return {};
  },

  toJSON(_: OpenFileReply): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenFileReply>, I>>(base?: I): OpenFileReply {
    return OpenFileReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OpenFileReply>, I>>(_: I): OpenFileReply {
    const message = createBaseOpenFileReply();
    return message;
  },
};

function createBaseEnableAutoMountRequest(): EnableAutoMountRequest {
  return { boxName: "" };
}

export const EnableAutoMountRequest = {
  encode(message: EnableAutoMountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.boxName !== "") {
      writer.uint32(10).string(message.boxName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EnableAutoMountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEnableAutoMountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.boxName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EnableAutoMountRequest {
    return { boxName: isSet(object.boxName) ? String(object.boxName) : "" };
  },

  toJSON(message: EnableAutoMountRequest): unknown {
    const obj: any = {};
    if (message.boxName !== "") {
      obj.boxName = message.boxName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<EnableAutoMountRequest>, I>>(base?: I): EnableAutoMountRequest {
    return EnableAutoMountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EnableAutoMountRequest>, I>>(object: I): EnableAutoMountRequest {
    const message = createBaseEnableAutoMountRequest();
    message.boxName = object.boxName ?? "";
    return message;
  },
};

function createBaseEnableAutoMountReply(): EnableAutoMountReply {
  return {};
}

export const EnableAutoMountReply = {
  encode(_: EnableAutoMountReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EnableAutoMountReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEnableAutoMountReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): EnableAutoMountReply {
    return {};
  },

  toJSON(_: EnableAutoMountReply): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<EnableAutoMountReply>, I>>(base?: I): EnableAutoMountReply {
    return EnableAutoMountReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<EnableAutoMountReply>, I>>(_: I): EnableAutoMountReply {
    const message = createBaseEnableAutoMountReply();
    return message;
  },
};

function createBaseIsEnableAutoMountRequest(): IsEnableAutoMountRequest {
  return { boxName: "" };
}

export const IsEnableAutoMountRequest = {
  encode(message: IsEnableAutoMountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.boxName !== "") {
      writer.uint32(10).string(message.boxName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IsEnableAutoMountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIsEnableAutoMountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.boxName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IsEnableAutoMountRequest {
    return { boxName: isSet(object.boxName) ? String(object.boxName) : "" };
  },

  toJSON(message: IsEnableAutoMountRequest): unknown {
    const obj: any = {};
    if (message.boxName !== "") {
      obj.boxName = message.boxName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<IsEnableAutoMountRequest>, I>>(base?: I): IsEnableAutoMountRequest {
    return IsEnableAutoMountRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<IsEnableAutoMountRequest>, I>>(object: I): IsEnableAutoMountRequest {
    const message = createBaseIsEnableAutoMountRequest();
    message.boxName = object.boxName ?? "";
    return message;
  },
};

function createBaseIsEnableAutoMountReply(): IsEnableAutoMountReply {
  return { enable: false };
}

export const IsEnableAutoMountReply = {
  encode(message: IsEnableAutoMountReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.enable === true) {
      writer.uint32(8).bool(message.enable);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IsEnableAutoMountReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIsEnableAutoMountReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.enable = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IsEnableAutoMountReply {
    return { enable: isSet(object.enable) ? Boolean(object.enable) : false };
  },

  toJSON(message: IsEnableAutoMountReply): unknown {
    const obj: any = {};
    if (message.enable === true) {
      obj.enable = message.enable;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<IsEnableAutoMountReply>, I>>(base?: I): IsEnableAutoMountReply {
    return IsEnableAutoMountReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<IsEnableAutoMountReply>, I>>(object: I): IsEnableAutoMountReply {
    const message = createBaseIsEnableAutoMountReply();
    message.enable = object.enable ?? false;
    return message;
  },
};

export interface CloudDriver {
  EnableAutoMount(
    request: DeepPartial<EnableAutoMountRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<EnableAutoMountReply>;
  IsEnableAutoMount(
    request: DeepPartial<IsEnableAutoMountRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<IsEnableAutoMountReply>;
  OpenFile(
    request: DeepPartial<OpenFileRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<OpenFileReply>;
}

export class CloudDriverClientImpl implements CloudDriver {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.EnableAutoMount = this.EnableAutoMount.bind(this);
    this.IsEnableAutoMount = this.IsEnableAutoMount.bind(this);
    this.OpenFile = this.OpenFile.bind(this);
  }

  EnableAutoMount(
    request: DeepPartial<EnableAutoMountRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<EnableAutoMountReply> {
    return this.rpc.unary(
      CloudDriverEnableAutoMountDesc,
      EnableAutoMountRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  IsEnableAutoMount(
    request: DeepPartial<IsEnableAutoMountRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<IsEnableAutoMountReply> {
    return this.rpc.unary(
      CloudDriverIsEnableAutoMountDesc,
      IsEnableAutoMountRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  OpenFile(
    request: DeepPartial<OpenFileRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<OpenFileReply> {
    return this.rpc.unary(CloudDriverOpenFileDesc, OpenFileRequest.fromPartial(request), metadata, abortSignal);
  }
}

export const CloudDriverDesc = { serviceName: "cloud.lazycat.apis.localdevice.CloudDriver" };

export const CloudDriverEnableAutoMountDesc: UnaryMethodDefinitionish = {
  methodName: "EnableAutoMount",
  service: CloudDriverDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return EnableAutoMountRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = EnableAutoMountReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const CloudDriverIsEnableAutoMountDesc: UnaryMethodDefinitionish = {
  methodName: "IsEnableAutoMount",
  service: CloudDriverDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return IsEnableAutoMountRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = IsEnableAutoMountReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const CloudDriverOpenFileDesc: UnaryMethodDefinitionish = {
  methodName: "OpenFile",
  service: CloudDriverDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return OpenFileRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = OpenFileReply.decode(data);
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

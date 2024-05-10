/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { share } from "rxjs/operators";

export interface ReadClipRequest {
  /** 目前只支持text/plain和image/png */
  mime: string;
}

export interface ReadClipResponse {
  content: Uint8Array;
}

export interface WriteClipRequest {
  mime: string;
  content: Uint8Array;
}

export interface WriteClipResponse {
  success: boolean;
}

function createBaseReadClipRequest(): ReadClipRequest {
  return { mime: "" };
}

export const ReadClipRequest = {
  encode(message: ReadClipRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.mime !== "") {
      writer.uint32(10).string(message.mime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadClipRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadClipRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.mime = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadClipRequest {
    return { mime: isSet(object.mime) ? String(object.mime) : "" };
  },

  toJSON(message: ReadClipRequest): unknown {
    const obj: any = {};
    if (message.mime !== "") {
      obj.mime = message.mime;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadClipRequest>, I>>(base?: I): ReadClipRequest {
    return ReadClipRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadClipRequest>, I>>(object: I): ReadClipRequest {
    const message = createBaseReadClipRequest();
    message.mime = object.mime ?? "";
    return message;
  },
};

function createBaseReadClipResponse(): ReadClipResponse {
  return { content: new Uint8Array(0) };
}

export const ReadClipResponse = {
  encode(message: ReadClipResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.content.length !== 0) {
      writer.uint32(10).bytes(message.content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadClipResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadClipResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.content = reader.bytes();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadClipResponse {
    return { content: isSet(object.content) ? bytesFromBase64(object.content) : new Uint8Array(0) };
  },

  toJSON(message: ReadClipResponse): unknown {
    const obj: any = {};
    if (message.content.length !== 0) {
      obj.content = base64FromBytes(message.content);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadClipResponse>, I>>(base?: I): ReadClipResponse {
    return ReadClipResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadClipResponse>, I>>(object: I): ReadClipResponse {
    const message = createBaseReadClipResponse();
    message.content = object.content ?? new Uint8Array(0);
    return message;
  },
};

function createBaseWriteClipRequest(): WriteClipRequest {
  return { mime: "", content: new Uint8Array(0) };
}

export const WriteClipRequest = {
  encode(message: WriteClipRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.mime !== "") {
      writer.uint32(10).string(message.mime);
    }
    if (message.content.length !== 0) {
      writer.uint32(18).bytes(message.content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WriteClipRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWriteClipRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.mime = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.content = reader.bytes();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): WriteClipRequest {
    return {
      mime: isSet(object.mime) ? String(object.mime) : "",
      content: isSet(object.content) ? bytesFromBase64(object.content) : new Uint8Array(0),
    };
  },

  toJSON(message: WriteClipRequest): unknown {
    const obj: any = {};
    if (message.mime !== "") {
      obj.mime = message.mime;
    }
    if (message.content.length !== 0) {
      obj.content = base64FromBytes(message.content);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<WriteClipRequest>, I>>(base?: I): WriteClipRequest {
    return WriteClipRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WriteClipRequest>, I>>(object: I): WriteClipRequest {
    const message = createBaseWriteClipRequest();
    message.mime = object.mime ?? "";
    message.content = object.content ?? new Uint8Array(0);
    return message;
  },
};

function createBaseWriteClipResponse(): WriteClipResponse {
  return { success: false };
}

export const WriteClipResponse = {
  encode(message: WriteClipResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WriteClipResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWriteClipResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.success = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): WriteClipResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: WriteClipResponse): unknown {
    const obj: any = {};
    if (message.success === true) {
      obj.success = message.success;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<WriteClipResponse>, I>>(base?: I): WriteClipResponse {
    return WriteClipResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WriteClipResponse>, I>>(object: I): WriteClipResponse {
    const message = createBaseWriteClipResponse();
    message.success = object.success ?? false;
    return message;
  },
};

export interface ClipboardManager {
  Read(
    request: DeepPartial<ReadClipRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ReadClipResponse>;
  Write(
    request: DeepPartial<WriteClipRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<WriteClipResponse>;
  Watch(
    request: DeepPartial<ReadClipRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<ReadClipResponse>;
}

export class ClipboardManagerClientImpl implements ClipboardManager {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Read = this.Read.bind(this);
    this.Write = this.Write.bind(this);
    this.Watch = this.Watch.bind(this);
  }

  Read(
    request: DeepPartial<ReadClipRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ReadClipResponse> {
    return this.rpc.unary(ClipboardManagerReadDesc, ReadClipRequest.fromPartial(request), metadata, abortSignal);
  }

  Write(
    request: DeepPartial<WriteClipRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<WriteClipResponse> {
    return this.rpc.unary(ClipboardManagerWriteDesc, WriteClipRequest.fromPartial(request), metadata, abortSignal);
  }

  Watch(
    request: DeepPartial<ReadClipRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<ReadClipResponse> {
    return this.rpc.invoke(ClipboardManagerWatchDesc, ReadClipRequest.fromPartial(request), metadata, abortSignal);
  }
}

export const ClipboardManagerDesc = { serviceName: "cloud.lazycat.apis.localdevice.ClipboardManager" };

export const ClipboardManagerReadDesc: UnaryMethodDefinitionish = {
  methodName: "Read",
  service: ClipboardManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ReadClipRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ReadClipResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const ClipboardManagerWriteDesc: UnaryMethodDefinitionish = {
  methodName: "Write",
  service: ClipboardManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return WriteClipRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = WriteClipResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const ClipboardManagerWatchDesc: UnaryMethodDefinitionish = {
  methodName: "Watch",
  service: ClipboardManagerDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return ReadClipRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ReadClipResponse.decode(data);
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
  invoke<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined,
    abortSignal?: AbortSignal,
  ): Observable<any>;
}

export class GrpcWebImpl {
  private host: string;
  private options: {
    transport?: grpc.TransportFactory;
    streamingTransport?: grpc.TransportFactory;
    debug?: boolean;
    metadata?: grpc.Metadata;
    upStreamRetryCodes?: number[];
  };

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory;
      streamingTransport?: grpc.TransportFactory;
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

  invoke<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined,
    abortSignal?: AbortSignal,
  ): Observable<any> {
    const upStreamCodes = this.options.upStreamRetryCodes ?? [];
    const DEFAULT_TIMEOUT_TIME: number = 3_000;
    const request = { ..._request, ...methodDesc.requestType };
    const transport = this.options.streamingTransport ?? this.options.transport;
    const maybeCombinedMetadata = metadata && this.options.metadata
      ? new BrowserHeaders({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
      : metadata ?? this.options.metadata;
    return new Observable((observer) => {
      const upStream = () => {
        const client = grpc.invoke(methodDesc, {
          host: this.host,
          request,
          ...(transport !== undefined ? { transport } : {}),
          metadata: maybeCombinedMetadata ?? {},
          debug: this.options.debug ?? false,
          onMessage: (next) => observer.next(next),
          onEnd: (code: grpc.Code, message: string, trailers: grpc.Metadata) => {
            if (code === 0) {
              observer.complete();
            } else if (upStreamCodes.includes(code)) {
              setTimeout(upStream, DEFAULT_TIMEOUT_TIME);
            } else {
              const err = new Error(message) as any;
              err.code = code;
              err.metadata = trailers;
              observer.error(err);
            }
          },
        });
        observer.add(() => {
          if (!abortSignal || !abortSignal.aborted) {
            return client.close();
          }
        });

        if (abortSignal) {
          abortSignal.addEventListener("abort", () => {
            observer.error(abortSignal.reason);
            client.close();
          });
        }
      };
      upStream();
    }).pipe(share());
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export class GrpcWebError extends tsProtoGlobalThis.Error {
  constructor(message: string, public code: grpc.Code, public metadata: grpc.Metadata) {
    super(message);
  }
}

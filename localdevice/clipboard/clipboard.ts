/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { BrowserHeaders } from "browser-headers";
import { share } from "rxjs/operators";

export interface ReadClipRequest {
  /** 目前只支持text/plain和image/png */
  mime: string;
}

export interface ReadClipReply {
  content: Uint8Array;
}

export interface WriteClipRequest {
  mimie: string;
  content: Uint8Array;
}

export interface WriteClipReply {}

function createBaseReadClipRequest(): ReadClipRequest {
  return { mime: "" };
}

export const ReadClipRequest = {
  encode(
    message: ReadClipRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.mime !== "") {
      writer.uint32(10).string(message.mime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadClipRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadClipRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mime = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ReadClipRequest {
    return {
      mime: isSet(object.mime) ? String(object.mime) : "",
    };
  },

  toJSON(message: ReadClipRequest): unknown {
    const obj: any = {};
    message.mime !== undefined && (obj.mime = message.mime);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ReadClipRequest>, I>>(
    object: I
  ): ReadClipRequest {
    const message = createBaseReadClipRequest();
    message.mime = object.mime ?? "";
    return message;
  },
};

function createBaseReadClipReply(): ReadClipReply {
  return { content: new Uint8Array() };
}

export const ReadClipReply = {
  encode(
    message: ReadClipReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.content.length !== 0) {
      writer.uint32(10).bytes(message.content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadClipReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadClipReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.content = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ReadClipReply {
    return {
      content: isSet(object.content)
        ? bytesFromBase64(object.content)
        : new Uint8Array(),
    };
  },

  toJSON(message: ReadClipReply): unknown {
    const obj: any = {};
    message.content !== undefined &&
      (obj.content = base64FromBytes(
        message.content !== undefined ? message.content : new Uint8Array()
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ReadClipReply>, I>>(
    object: I
  ): ReadClipReply {
    const message = createBaseReadClipReply();
    message.content = object.content ?? new Uint8Array();
    return message;
  },
};

function createBaseWriteClipRequest(): WriteClipRequest {
  return { mimie: "", content: new Uint8Array() };
}

export const WriteClipRequest = {
  encode(
    message: WriteClipRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.mimie !== "") {
      writer.uint32(10).string(message.mimie);
    }
    if (message.content.length !== 0) {
      writer.uint32(18).bytes(message.content);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WriteClipRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWriteClipRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mimie = reader.string();
          break;
        case 2:
          message.content = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WriteClipRequest {
    return {
      mimie: isSet(object.mimie) ? String(object.mimie) : "",
      content: isSet(object.content)
        ? bytesFromBase64(object.content)
        : new Uint8Array(),
    };
  },

  toJSON(message: WriteClipRequest): unknown {
    const obj: any = {};
    message.mimie !== undefined && (obj.mimie = message.mimie);
    message.content !== undefined &&
      (obj.content = base64FromBytes(
        message.content !== undefined ? message.content : new Uint8Array()
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<WriteClipRequest>, I>>(
    object: I
  ): WriteClipRequest {
    const message = createBaseWriteClipRequest();
    message.mimie = object.mimie ?? "";
    message.content = object.content ?? new Uint8Array();
    return message;
  },
};

function createBaseWriteClipReply(): WriteClipReply {
  return {};
}

export const WriteClipReply = {
  encode(
    _: WriteClipReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WriteClipReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWriteClipReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): WriteClipReply {
    return {};
  },

  toJSON(_: WriteClipReply): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<WriteClipReply>, I>>(
    _: I
  ): WriteClipReply {
    const message = createBaseWriteClipReply();
    return message;
  },
};

export interface ClipboardManager {
  Read(
    request: DeepPartial<ReadClipRequest>,
    metadata?: grpc.Metadata
  ): Promise<ReadClipReply>;
  Write(
    request: DeepPartial<WriteClipRequest>,
    metadata?: grpc.Metadata
  ): Promise<WriteClipReply>;
  Watch(
    request: DeepPartial<ReadClipRequest>,
    metadata?: grpc.Metadata
  ): Observable<ReadClipReply>;
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
    metadata?: grpc.Metadata
  ): Promise<ReadClipReply> {
    return this.rpc.unary(
      ClipboardManagerReadDesc,
      ReadClipRequest.fromPartial(request),
      metadata
    );
  }

  Write(
    request: DeepPartial<WriteClipRequest>,
    metadata?: grpc.Metadata
  ): Promise<WriteClipReply> {
    return this.rpc.unary(
      ClipboardManagerWriteDesc,
      WriteClipRequest.fromPartial(request),
      metadata
    );
  }

  Watch(
    request: DeepPartial<ReadClipRequest>,
    metadata?: grpc.Metadata
  ): Observable<ReadClipReply> {
    return this.rpc.invoke(
      ClipboardManagerWatchDesc,
      ReadClipRequest.fromPartial(request),
      metadata
    );
  }
}

export const ClipboardManagerDesc = {
  serviceName: "cloud.lazycat.localdevice.apis.ClipboardManager",
};

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
      return {
        ...ReadClipReply.decode(data),
        toObject() {
          return this;
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
      return {
        ...WriteClipReply.decode(data),
        toObject() {
          return this;
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
      return {
        ...ReadClipReply.decode(data),
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

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

const atob: (b64: string) => string =
  globalThis.atob ||
  ((b64) => globalThis.Buffer.from(b64, "base64").toString("binary"));
function bytesFromBase64(b64: string): Uint8Array {
  const bin = atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) {
    arr[i] = bin.charCodeAt(i);
  }
  return arr;
}

const btoa: (bin: string) => string =
  globalThis.btoa ||
  ((bin) => globalThis.Buffer.from(bin, "binary").toString("base64"));
function base64FromBytes(arr: Uint8Array): string {
  const bin: string[] = [];
  arr.forEach((byte) => {
    bin.push(String.fromCharCode(byte));
  });
  return btoa(bin.join(""));
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

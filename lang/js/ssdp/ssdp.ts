/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { share } from "rxjs/operators";

export interface SearchRequest {
  /** "urn:schemas-upnp-org:service:AVTransport:1", "ssdp:all" 之类的 */
  type: string;
  /** 0-120之间 */
  waitSeconds: number;
}

export interface Service {
  /** Type is a property of "ST" */
  type: string;
  /** USN is a property of "USN" */
  USN: string;
  /** Location is a property of "LOCATION" */
  location: string;
  /** Server is a property of "SERVER" */
  server: string;
}

function createBaseSearchRequest(): SearchRequest {
  return { type: "", waitSeconds: 0 };
}

export const SearchRequest = {
  encode(message: SearchRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== "") {
      writer.uint32(10).string(message.type);
    }
    if (message.waitSeconds !== 0) {
      writer.uint32(16).int32(message.waitSeconds);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSearchRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.string();
          break;
        case 2:
          message.waitSeconds = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchRequest {
    return {
      type: isSet(object.type) ? String(object.type) : "",
      waitSeconds: isSet(object.waitSeconds) ? Number(object.waitSeconds) : 0,
    };
  },

  toJSON(message: SearchRequest): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = message.type);
    message.waitSeconds !== undefined && (obj.waitSeconds = Math.round(message.waitSeconds));
    return obj;
  },

  create<I extends Exact<DeepPartial<SearchRequest>, I>>(base?: I): SearchRequest {
    return SearchRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SearchRequest>, I>>(object: I): SearchRequest {
    const message = createBaseSearchRequest();
    message.type = object.type ?? "";
    message.waitSeconds = object.waitSeconds ?? 0;
    return message;
  },
};

function createBaseService(): Service {
  return { type: "", USN: "", location: "", server: "" };
}

export const Service = {
  encode(message: Service, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== "") {
      writer.uint32(10).string(message.type);
    }
    if (message.USN !== "") {
      writer.uint32(18).string(message.USN);
    }
    if (message.location !== "") {
      writer.uint32(26).string(message.location);
    }
    if (message.server !== "") {
      writer.uint32(34).string(message.server);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Service {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseService();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.string();
          break;
        case 2:
          message.USN = reader.string();
          break;
        case 3:
          message.location = reader.string();
          break;
        case 4:
          message.server = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Service {
    return {
      type: isSet(object.type) ? String(object.type) : "",
      USN: isSet(object.USN) ? String(object.USN) : "",
      location: isSet(object.location) ? String(object.location) : "",
      server: isSet(object.server) ? String(object.server) : "",
    };
  },

  toJSON(message: Service): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = message.type);
    message.USN !== undefined && (obj.USN = message.USN);
    message.location !== undefined && (obj.location = message.location);
    message.server !== undefined && (obj.server = message.server);
    return obj;
  },

  create<I extends Exact<DeepPartial<Service>, I>>(base?: I): Service {
    return Service.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Service>, I>>(object: I): Service {
    const message = createBaseService();
    message.type = object.type ?? "";
    message.USN = object.USN ?? "";
    message.location = object.location ?? "";
    message.server = object.server ?? "";
    return message;
  },
};

export interface SSDPService {
  Search(request: DeepPartial<SearchRequest>, metadata?: grpc.Metadata): Observable<Service>;
}

export class SSDPServiceClientImpl implements SSDPService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Search = this.Search.bind(this);
  }

  Search(request: DeepPartial<SearchRequest>, metadata?: grpc.Metadata): Observable<Service> {
    return this.rpc.invoke(SSDPServiceSearchDesc, SearchRequest.fromPartial(request), metadata);
  }
}

export const SSDPServiceDesc = { serviceName: "lzc.ssdp.SSDPService" };

export const SSDPServiceSearchDesc: UnaryMethodDefinitionish = {
  methodName: "Search",
  service: SSDPServiceDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return SearchRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = Service.decode(data);
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
  ): Promise<any>;
  invoke<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined,
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
  ): Promise<any> {
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata = metadata && this.options.metadata
      ? new BrowserHeaders({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
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
            resolve(response.message!.toObject());
          } else {
            const err = new GrpcWebError(response.statusMessage, response.status, response.trailers);
            reject(err);
          }
        },
      });
    });
  }

  invoke<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined,
  ): Observable<any> {
    const upStreamCodes = this.options.upStreamRetryCodes || [];
    const DEFAULT_TIMEOUT_TIME: number = 3_000;
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata = metadata && this.options.metadata
      ? new BrowserHeaders({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
      : metadata || this.options.metadata;
    return new Observable((observer) => {
      const upStream = (() => {
        const client = grpc.invoke(methodDesc, {
          host: this.host,
          request,
          transport: this.options.streamingTransport || this.options.transport,
          metadata: maybeCombinedMetadata,
          debug: this.options.debug,
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
        observer.add(() => client.close());
      });
      upStream();
    }).pipe(share());
  }
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var tsProtoGlobalThis: any = (() => {
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

/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { share } from "rxjs/operators";
import { Empty } from "../google/protobuf/empty";

export interface InputContentReply {
  text: string;
}

export interface IsInputFocusResponse {
  focus: boolean;
}

export interface SetInputTextRequest {
  text: string;
}

function createBaseInputContentReply(): InputContentReply {
  return { text: "" };
}

export const InputContentReply = {
  encode(message: InputContentReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.text !== "") {
      writer.uint32(10).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InputContentReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInputContentReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.text = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): InputContentReply {
    return { text: isSet(object.text) ? String(object.text) : "" };
  },

  toJSON(message: InputContentReply): unknown {
    const obj: any = {};
    if (message.text !== "") {
      obj.text = message.text;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<InputContentReply>, I>>(base?: I): InputContentReply {
    return InputContentReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<InputContentReply>, I>>(object: I): InputContentReply {
    const message = createBaseInputContentReply();
    message.text = object.text ?? "";
    return message;
  },
};

function createBaseIsInputFocusResponse(): IsInputFocusResponse {
  return { focus: false };
}

export const IsInputFocusResponse = {
  encode(message: IsInputFocusResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.focus === true) {
      writer.uint32(8).bool(message.focus);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IsInputFocusResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIsInputFocusResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.focus = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): IsInputFocusResponse {
    return { focus: isSet(object.focus) ? Boolean(object.focus) : false };
  },

  toJSON(message: IsInputFocusResponse): unknown {
    const obj: any = {};
    if (message.focus === true) {
      obj.focus = message.focus;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<IsInputFocusResponse>, I>>(base?: I): IsInputFocusResponse {
    return IsInputFocusResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<IsInputFocusResponse>, I>>(object: I): IsInputFocusResponse {
    const message = createBaseIsInputFocusResponse();
    message.focus = object.focus ?? false;
    return message;
  },
};

function createBaseSetInputTextRequest(): SetInputTextRequest {
  return { text: "" };
}

export const SetInputTextRequest = {
  encode(message: SetInputTextRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.text !== "") {
      writer.uint32(18).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SetInputTextRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetInputTextRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          if (tag !== 18) {
            break;
          }

          message.text = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SetInputTextRequest {
    return { text: isSet(object.text) ? String(object.text) : "" };
  },

  toJSON(message: SetInputTextRequest): unknown {
    const obj: any = {};
    if (message.text !== "") {
      obj.text = message.text;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SetInputTextRequest>, I>>(base?: I): SetInputTextRequest {
    return SetInputTextRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SetInputTextRequest>, I>>(object: I): SetInputTextRequest {
    const message = createBaseSetInputTextRequest();
    message.text = object.text ?? "";
    return message;
  },
};

export interface Rim {
  /** 监听输入框聚焦 */
  ListenInputFocus(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<InputContentReply>;
  /** 监听输入框失焦 */
  ListenInputBlur(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<InputContentReply>;
  /** 获取当前聚焦输入框文本 */
  GetInputText(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<InputContentReply>;
  /** 设置当前聚焦输入框文本 */
  SetInputText(
    request: DeepPartial<SetInputTextRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 获取当前是否有聚焦输入框 */
  IsInputFocus(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<IsInputFocusResponse>;
  /** 给聚焦窗口发送回车输入事件 */
  SendInputEventEnter(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<Empty>;
}

export class RimClientImpl implements Rim {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ListenInputFocus = this.ListenInputFocus.bind(this);
    this.ListenInputBlur = this.ListenInputBlur.bind(this);
    this.GetInputText = this.GetInputText.bind(this);
    this.SetInputText = this.SetInputText.bind(this);
    this.IsInputFocus = this.IsInputFocus.bind(this);
    this.SendInputEventEnter = this.SendInputEventEnter.bind(this);
  }

  ListenInputFocus(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<InputContentReply> {
    return this.rpc.invoke(RimListenInputFocusDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  ListenInputBlur(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<InputContentReply> {
    return this.rpc.invoke(RimListenInputBlurDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  GetInputText(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<InputContentReply> {
    return this.rpc.unary(RimGetInputTextDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  SetInputText(
    request: DeepPartial<SetInputTextRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(RimSetInputTextDesc, SetInputTextRequest.fromPartial(request), metadata, abortSignal);
  }

  IsInputFocus(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<IsInputFocusResponse> {
    return this.rpc.unary(RimIsInputFocusDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  SendInputEventEnter(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(RimSendInputEventEnterDesc, Empty.fromPartial(request), metadata, abortSignal);
  }
}

export const RimDesc = { serviceName: "cloud.lazycat.apis.localdevice.Rim" };

export const RimListenInputFocusDesc: UnaryMethodDefinitionish = {
  methodName: "ListenInputFocus",
  service: RimDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = InputContentReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RimListenInputBlurDesc: UnaryMethodDefinitionish = {
  methodName: "ListenInputBlur",
  service: RimDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = InputContentReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RimGetInputTextDesc: UnaryMethodDefinitionish = {
  methodName: "GetInputText",
  service: RimDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = InputContentReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RimSetInputTextDesc: UnaryMethodDefinitionish = {
  methodName: "SetInputText",
  service: RimDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SetInputTextRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = Empty.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RimIsInputFocusDesc: UnaryMethodDefinitionish = {
  methodName: "IsInputFocus",
  service: RimDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = IsInputFocusResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RimSendInputEventEnterDesc: UnaryMethodDefinitionish = {
  methodName: "SendInputEventEnter",
  service: RimDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = Empty.decode(data);
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

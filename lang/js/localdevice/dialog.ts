/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";
import { Empty } from "../google/protobuf/empty";

export interface QuestionRequest {
  title: string;
  text: string;
  defaultCancel: boolean;
}

export interface QuestionResult {
  yes: boolean;
}

export interface MessageBoxRequest {
  title: string;
  text: string;
}

export interface PasswordRequest {
  title: string;
  text: string;
}

export interface PasswordResult {
  password: string;
  ok: boolean;
}

export interface OpenFileSeletorRequest {
  /** 选择类型 */
  type: OpenFileSeletorRequest_SelectType;
  /** 是否多选（多选在SelectType 为Dir时无效） */
  isSingle: boolean;
  /**
   * 文件后缀过滤器
   * 使用2种类型
   * 第一种 jpb,png,mp3，指定文件的后缀名,使用,分割
   * 支持的大类型有必须 image/* , 或者image/*,audio/*,video/*,document/*, 四种大类可以写在一起，也可以分开写,使用,分割。 如何该字段为空，则不过滤
   */
  filter: string;
}

export enum OpenFileSeletorRequest_SelectType {
  Dir = 0,
  File = 1,
  UNRECOGNIZED = -1,
}

export function openFileSeletorRequest_SelectTypeFromJSON(object: any): OpenFileSeletorRequest_SelectType {
  switch (object) {
    case 0:
    case "Dir":
      return OpenFileSeletorRequest_SelectType.Dir;
    case 1:
    case "File":
      return OpenFileSeletorRequest_SelectType.File;
    case -1:
    case "UNRECOGNIZED":
    default:
      return OpenFileSeletorRequest_SelectType.UNRECOGNIZED;
  }
}

export function openFileSeletorRequest_SelectTypeToJSON(object: OpenFileSeletorRequest_SelectType): string {
  switch (object) {
    case OpenFileSeletorRequest_SelectType.Dir:
      return "Dir";
    case OpenFileSeletorRequest_SelectType.File:
      return "File";
    case OpenFileSeletorRequest_SelectType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface OpenFileSeletorResult {
  /** 文件或者目录的路径, 如果是文件就是文件的全路径 比如 /home/lnks/1.jpg ,如果是目录就是目录的全路径,并且末尾以/结尾, 比如/home/lnks/ */
  path: string[];
}

function createBaseQuestionRequest(): QuestionRequest {
  return { title: "", text: "", defaultCancel: false };
}

export const QuestionRequest = {
  encode(message: QuestionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.text !== "") {
      writer.uint32(18).string(message.text);
    }
    if (message.defaultCancel === true) {
      writer.uint32(24).bool(message.defaultCancel);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuestionRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuestionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.title = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.text = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.defaultCancel = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QuestionRequest {
    return {
      title: isSet(object.title) ? String(object.title) : "",
      text: isSet(object.text) ? String(object.text) : "",
      defaultCancel: isSet(object.defaultCancel) ? Boolean(object.defaultCancel) : false,
    };
  },

  toJSON(message: QuestionRequest): unknown {
    const obj: any = {};
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.text !== "") {
      obj.text = message.text;
    }
    if (message.defaultCancel === true) {
      obj.defaultCancel = message.defaultCancel;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QuestionRequest>, I>>(base?: I): QuestionRequest {
    return QuestionRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<QuestionRequest>, I>>(object: I): QuestionRequest {
    const message = createBaseQuestionRequest();
    message.title = object.title ?? "";
    message.text = object.text ?? "";
    message.defaultCancel = object.defaultCancel ?? false;
    return message;
  },
};

function createBaseQuestionResult(): QuestionResult {
  return { yes: false };
}

export const QuestionResult = {
  encode(message: QuestionResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.yes === true) {
      writer.uint32(8).bool(message.yes);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuestionResult {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuestionResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.yes = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QuestionResult {
    return { yes: isSet(object.yes) ? Boolean(object.yes) : false };
  },

  toJSON(message: QuestionResult): unknown {
    const obj: any = {};
    if (message.yes === true) {
      obj.yes = message.yes;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QuestionResult>, I>>(base?: I): QuestionResult {
    return QuestionResult.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<QuestionResult>, I>>(object: I): QuestionResult {
    const message = createBaseQuestionResult();
    message.yes = object.yes ?? false;
    return message;
  },
};

function createBaseMessageBoxRequest(): MessageBoxRequest {
  return { title: "", text: "" };
}

export const MessageBoxRequest = {
  encode(message: MessageBoxRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.text !== "") {
      writer.uint32(18).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MessageBoxRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMessageBoxRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.title = reader.string();
          continue;
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

  fromJSON(object: any): MessageBoxRequest {
    return {
      title: isSet(object.title) ? String(object.title) : "",
      text: isSet(object.text) ? String(object.text) : "",
    };
  },

  toJSON(message: MessageBoxRequest): unknown {
    const obj: any = {};
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.text !== "") {
      obj.text = message.text;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MessageBoxRequest>, I>>(base?: I): MessageBoxRequest {
    return MessageBoxRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<MessageBoxRequest>, I>>(object: I): MessageBoxRequest {
    const message = createBaseMessageBoxRequest();
    message.title = object.title ?? "";
    message.text = object.text ?? "";
    return message;
  },
};

function createBasePasswordRequest(): PasswordRequest {
  return { title: "", text: "" };
}

export const PasswordRequest = {
  encode(message: PasswordRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.text !== "") {
      writer.uint32(18).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PasswordRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePasswordRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.title = reader.string();
          continue;
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

  fromJSON(object: any): PasswordRequest {
    return {
      title: isSet(object.title) ? String(object.title) : "",
      text: isSet(object.text) ? String(object.text) : "",
    };
  },

  toJSON(message: PasswordRequest): unknown {
    const obj: any = {};
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.text !== "") {
      obj.text = message.text;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PasswordRequest>, I>>(base?: I): PasswordRequest {
    return PasswordRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PasswordRequest>, I>>(object: I): PasswordRequest {
    const message = createBasePasswordRequest();
    message.title = object.title ?? "";
    message.text = object.text ?? "";
    return message;
  },
};

function createBasePasswordResult(): PasswordResult {
  return { password: "", ok: false };
}

export const PasswordResult = {
  encode(message: PasswordResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.password !== "") {
      writer.uint32(10).string(message.password);
    }
    if (message.ok === true) {
      writer.uint32(16).bool(message.ok);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PasswordResult {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePasswordResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.password = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.ok = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PasswordResult {
    return {
      password: isSet(object.password) ? String(object.password) : "",
      ok: isSet(object.ok) ? Boolean(object.ok) : false,
    };
  },

  toJSON(message: PasswordResult): unknown {
    const obj: any = {};
    if (message.password !== "") {
      obj.password = message.password;
    }
    if (message.ok === true) {
      obj.ok = message.ok;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PasswordResult>, I>>(base?: I): PasswordResult {
    return PasswordResult.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PasswordResult>, I>>(object: I): PasswordResult {
    const message = createBasePasswordResult();
    message.password = object.password ?? "";
    message.ok = object.ok ?? false;
    return message;
  },
};

function createBaseOpenFileSeletorRequest(): OpenFileSeletorRequest {
  return { type: 0, isSingle: false, filter: "" };
}

export const OpenFileSeletorRequest = {
  encode(message: OpenFileSeletorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.isSingle === true) {
      writer.uint32(16).bool(message.isSingle);
    }
    if (message.filter !== "") {
      writer.uint32(26).string(message.filter);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenFileSeletorRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenFileSeletorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.isSingle = reader.bool();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.filter = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OpenFileSeletorRequest {
    return {
      type: isSet(object.type) ? openFileSeletorRequest_SelectTypeFromJSON(object.type) : 0,
      isSingle: isSet(object.isSingle) ? Boolean(object.isSingle) : false,
      filter: isSet(object.filter) ? String(object.filter) : "",
    };
  },

  toJSON(message: OpenFileSeletorRequest): unknown {
    const obj: any = {};
    if (message.type !== 0) {
      obj.type = openFileSeletorRequest_SelectTypeToJSON(message.type);
    }
    if (message.isSingle === true) {
      obj.isSingle = message.isSingle;
    }
    if (message.filter !== "") {
      obj.filter = message.filter;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenFileSeletorRequest>, I>>(base?: I): OpenFileSeletorRequest {
    return OpenFileSeletorRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OpenFileSeletorRequest>, I>>(object: I): OpenFileSeletorRequest {
    const message = createBaseOpenFileSeletorRequest();
    message.type = object.type ?? 0;
    message.isSingle = object.isSingle ?? false;
    message.filter = object.filter ?? "";
    return message;
  },
};

function createBaseOpenFileSeletorResult(): OpenFileSeletorResult {
  return { path: [] };
}

export const OpenFileSeletorResult = {
  encode(message: OpenFileSeletorResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.path) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenFileSeletorResult {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenFileSeletorResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.path.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OpenFileSeletorResult {
    return { path: Array.isArray(object?.path) ? object.path.map((e: any) => String(e)) : [] };
  },

  toJSON(message: OpenFileSeletorResult): unknown {
    const obj: any = {};
    if (message.path?.length) {
      obj.path = message.path;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenFileSeletorResult>, I>>(base?: I): OpenFileSeletorResult {
    return OpenFileSeletorResult.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OpenFileSeletorResult>, I>>(object: I): OpenFileSeletorResult {
    const message = createBaseOpenFileSeletorResult();
    message.path = object.path?.map((e) => e) || [];
    return message;
  },
};

export interface DialogManager {
  Question(
    request: DeepPartial<QuestionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<QuestionResult>;
  MessageBox(
    request: DeepPartial<MessageBoxRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  Password(
    request: DeepPartial<PasswordRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<PasswordResult>;
  OpenFileSeletor(
    request: DeepPartial<OpenFileSeletorRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<OpenFileSeletorResult>;
}

export class DialogManagerClientImpl implements DialogManager {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Question = this.Question.bind(this);
    this.MessageBox = this.MessageBox.bind(this);
    this.Password = this.Password.bind(this);
    this.OpenFileSeletor = this.OpenFileSeletor.bind(this);
  }

  Question(
    request: DeepPartial<QuestionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<QuestionResult> {
    return this.rpc.unary(DialogManagerQuestionDesc, QuestionRequest.fromPartial(request), metadata, abortSignal);
  }

  MessageBox(
    request: DeepPartial<MessageBoxRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(DialogManagerMessageBoxDesc, MessageBoxRequest.fromPartial(request), metadata, abortSignal);
  }

  Password(
    request: DeepPartial<PasswordRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<PasswordResult> {
    return this.rpc.unary(DialogManagerPasswordDesc, PasswordRequest.fromPartial(request), metadata, abortSignal);
  }

  OpenFileSeletor(
    request: DeepPartial<OpenFileSeletorRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<OpenFileSeletorResult> {
    return this.rpc.unary(
      DialogManagerOpenFileSeletorDesc,
      OpenFileSeletorRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }
}

export const DialogManagerDesc = { serviceName: "cloud.lazycat.apis.localdevice.DialogManager" };

export const DialogManagerQuestionDesc: UnaryMethodDefinitionish = {
  methodName: "Question",
  service: DialogManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return QuestionRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = QuestionResult.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const DialogManagerMessageBoxDesc: UnaryMethodDefinitionish = {
  methodName: "MessageBox",
  service: DialogManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return MessageBoxRequest.encode(this).finish();
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

export const DialogManagerPasswordDesc: UnaryMethodDefinitionish = {
  methodName: "Password",
  service: DialogManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return PasswordRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = PasswordResult.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const DialogManagerOpenFileSeletorDesc: UnaryMethodDefinitionish = {
  methodName: "OpenFileSeletor",
  service: DialogManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return OpenFileSeletorRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = OpenFileSeletorResult.decode(data);
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

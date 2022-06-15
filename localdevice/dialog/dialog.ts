/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../../google/protobuf/empty";
import { BrowserHeaders } from "browser-headers";

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

function createBaseQuestionRequest(): QuestionRequest {
  return { title: "", text: "", defaultCancel: false };
}

export const QuestionRequest = {
  encode(
    message: QuestionRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
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
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuestionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.text = reader.string();
          break;
        case 3:
          message.defaultCancel = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuestionRequest {
    return {
      title: isSet(object.title) ? String(object.title) : "",
      text: isSet(object.text) ? String(object.text) : "",
      defaultCancel: isSet(object.defaultCancel)
        ? Boolean(object.defaultCancel)
        : false,
    };
  },

  toJSON(message: QuestionRequest): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.text !== undefined && (obj.text = message.text);
    message.defaultCancel !== undefined &&
      (obj.defaultCancel = message.defaultCancel);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QuestionRequest>, I>>(
    object: I
  ): QuestionRequest {
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
  encode(
    message: QuestionResult,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.yes === true) {
      writer.uint32(8).bool(message.yes);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuestionResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuestionResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.yes = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuestionResult {
    return {
      yes: isSet(object.yes) ? Boolean(object.yes) : false,
    };
  },

  toJSON(message: QuestionResult): unknown {
    const obj: any = {};
    message.yes !== undefined && (obj.yes = message.yes);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QuestionResult>, I>>(
    object: I
  ): QuestionResult {
    const message = createBaseQuestionResult();
    message.yes = object.yes ?? false;
    return message;
  },
};

function createBaseMessageBoxRequest(): MessageBoxRequest {
  return { title: "", text: "" };
}

export const MessageBoxRequest = {
  encode(
    message: MessageBoxRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.text !== "") {
      writer.uint32(18).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MessageBoxRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMessageBoxRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.text = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
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
    message.title !== undefined && (obj.title = message.title);
    message.text !== undefined && (obj.text = message.text);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MessageBoxRequest>, I>>(
    object: I
  ): MessageBoxRequest {
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
  encode(
    message: PasswordRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.text !== "") {
      writer.uint32(18).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PasswordRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePasswordRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.text = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
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
    message.title !== undefined && (obj.title = message.title);
    message.text !== undefined && (obj.text = message.text);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PasswordRequest>, I>>(
    object: I
  ): PasswordRequest {
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
  encode(
    message: PasswordResult,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.password !== "") {
      writer.uint32(10).string(message.password);
    }
    if (message.ok === true) {
      writer.uint32(16).bool(message.ok);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PasswordResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePasswordResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.password = reader.string();
          break;
        case 2:
          message.ok = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
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
    message.password !== undefined && (obj.password = message.password);
    message.ok !== undefined && (obj.ok = message.ok);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PasswordResult>, I>>(
    object: I
  ): PasswordResult {
    const message = createBasePasswordResult();
    message.password = object.password ?? "";
    message.ok = object.ok ?? false;
    return message;
  },
};

export interface DialogManager {
  Question(
    request: DeepPartial<QuestionRequest>,
    metadata?: grpc.Metadata
  ): Promise<QuestionResult>;
  MessageBox(
    request: DeepPartial<MessageBoxRequest>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
  Password(
    request: DeepPartial<PasswordRequest>,
    metadata?: grpc.Metadata
  ): Promise<PasswordResult>;
}

export class DialogManagerClientImpl implements DialogManager {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Question = this.Question.bind(this);
    this.MessageBox = this.MessageBox.bind(this);
    this.Password = this.Password.bind(this);
  }

  Question(
    request: DeepPartial<QuestionRequest>,
    metadata?: grpc.Metadata
  ): Promise<QuestionResult> {
    return this.rpc.unary(
      DialogManagerQuestionDesc,
      QuestionRequest.fromPartial(request),
      metadata
    );
  }

  MessageBox(
    request: DeepPartial<MessageBoxRequest>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      DialogManagerMessageBoxDesc,
      MessageBoxRequest.fromPartial(request),
      metadata
    );
  }

  Password(
    request: DeepPartial<PasswordRequest>,
    metadata?: grpc.Metadata
  ): Promise<PasswordResult> {
    return this.rpc.unary(
      DialogManagerPasswordDesc,
      PasswordRequest.fromPartial(request),
      metadata
    );
  }
}

export const DialogManagerDesc = {
  serviceName: "cloud.lazycat.localdevice.apis.DialogManager",
};

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
      return {
        ...QuestionResult.decode(data),
        toObject() {
          return this;
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
      return {
        ...Empty.decode(data),
        toObject() {
          return this;
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
      return {
        ...PasswordResult.decode(data),
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

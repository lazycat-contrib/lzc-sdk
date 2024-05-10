/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";

export interface SetUserConfigRequest {
  boxId: string;
  userId: string;
  appId: string;
  configKey: string;
  configValue: string;
}

export interface SetUserConfigResponse {
}

export interface GetUserConfigRequest {
  boxId: string;
  userId: string;
  appId: string;
  configKey: string;
}

export interface GetUserConfigResponse {
  configValue: string;
}

function createBaseSetUserConfigRequest(): SetUserConfigRequest {
  return { boxId: "", userId: "", appId: "", configKey: "", configValue: "" };
}

export const SetUserConfigRequest = {
  encode(message: SetUserConfigRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.boxId !== "") {
      writer.uint32(10).string(message.boxId);
    }
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    if (message.appId !== "") {
      writer.uint32(42).string(message.appId);
    }
    if (message.configKey !== "") {
      writer.uint32(26).string(message.configKey);
    }
    if (message.configValue !== "") {
      writer.uint32(34).string(message.configValue);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SetUserConfigRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetUserConfigRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.boxId = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.userId = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.appId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.configKey = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.configValue = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SetUserConfigRequest {
    return {
      boxId: isSet(object.boxId) ? String(object.boxId) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      appId: isSet(object.appId) ? String(object.appId) : "",
      configKey: isSet(object.configKey) ? String(object.configKey) : "",
      configValue: isSet(object.configValue) ? String(object.configValue) : "",
    };
  },

  toJSON(message: SetUserConfigRequest): unknown {
    const obj: any = {};
    if (message.boxId !== "") {
      obj.boxId = message.boxId;
    }
    if (message.userId !== "") {
      obj.userId = message.userId;
    }
    if (message.appId !== "") {
      obj.appId = message.appId;
    }
    if (message.configKey !== "") {
      obj.configKey = message.configKey;
    }
    if (message.configValue !== "") {
      obj.configValue = message.configValue;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SetUserConfigRequest>, I>>(base?: I): SetUserConfigRequest {
    return SetUserConfigRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SetUserConfigRequest>, I>>(object: I): SetUserConfigRequest {
    const message = createBaseSetUserConfigRequest();
    message.boxId = object.boxId ?? "";
    message.userId = object.userId ?? "";
    message.appId = object.appId ?? "";
    message.configKey = object.configKey ?? "";
    message.configValue = object.configValue ?? "";
    return message;
  },
};

function createBaseSetUserConfigResponse(): SetUserConfigResponse {
  return {};
}

export const SetUserConfigResponse = {
  encode(_: SetUserConfigResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SetUserConfigResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetUserConfigResponse();
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

  fromJSON(_: any): SetUserConfigResponse {
    return {};
  },

  toJSON(_: SetUserConfigResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<SetUserConfigResponse>, I>>(base?: I): SetUserConfigResponse {
    return SetUserConfigResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SetUserConfigResponse>, I>>(_: I): SetUserConfigResponse {
    const message = createBaseSetUserConfigResponse();
    return message;
  },
};

function createBaseGetUserConfigRequest(): GetUserConfigRequest {
  return { boxId: "", userId: "", appId: "", configKey: "" };
}

export const GetUserConfigRequest = {
  encode(message: GetUserConfigRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.boxId !== "") {
      writer.uint32(10).string(message.boxId);
    }
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    if (message.appId !== "") {
      writer.uint32(34).string(message.appId);
    }
    if (message.configKey !== "") {
      writer.uint32(26).string(message.configKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserConfigRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserConfigRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.boxId = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.userId = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.appId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.configKey = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserConfigRequest {
    return {
      boxId: isSet(object.boxId) ? String(object.boxId) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      appId: isSet(object.appId) ? String(object.appId) : "",
      configKey: isSet(object.configKey) ? String(object.configKey) : "",
    };
  },

  toJSON(message: GetUserConfigRequest): unknown {
    const obj: any = {};
    if (message.boxId !== "") {
      obj.boxId = message.boxId;
    }
    if (message.userId !== "") {
      obj.userId = message.userId;
    }
    if (message.appId !== "") {
      obj.appId = message.appId;
    }
    if (message.configKey !== "") {
      obj.configKey = message.configKey;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserConfigRequest>, I>>(base?: I): GetUserConfigRequest {
    return GetUserConfigRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetUserConfigRequest>, I>>(object: I): GetUserConfigRequest {
    const message = createBaseGetUserConfigRequest();
    message.boxId = object.boxId ?? "";
    message.userId = object.userId ?? "";
    message.appId = object.appId ?? "";
    message.configKey = object.configKey ?? "";
    return message;
  },
};

function createBaseGetUserConfigResponse(): GetUserConfigResponse {
  return { configValue: "" };
}

export const GetUserConfigResponse = {
  encode(message: GetUserConfigResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.configValue !== "") {
      writer.uint32(10).string(message.configValue);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetUserConfigResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserConfigResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.configValue = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserConfigResponse {
    return { configValue: isSet(object.configValue) ? String(object.configValue) : "" };
  },

  toJSON(message: GetUserConfigResponse): unknown {
    const obj: any = {};
    if (message.configValue !== "") {
      obj.configValue = message.configValue;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserConfigResponse>, I>>(base?: I): GetUserConfigResponse {
    return GetUserConfigResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetUserConfigResponse>, I>>(object: I): GetUserConfigResponse {
    const message = createBaseGetUserConfigResponse();
    message.configValue = object.configValue ?? "";
    return message;
  },
};

export interface UserConfig {
  GetUserConfig(
    request: DeepPartial<GetUserConfigRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<GetUserConfigResponse>;
  SetUserConfig(
    request: DeepPartial<SetUserConfigRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<SetUserConfigResponse>;
}

export class UserConfigClientImpl implements UserConfig {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.GetUserConfig = this.GetUserConfig.bind(this);
    this.SetUserConfig = this.SetUserConfig.bind(this);
  }

  GetUserConfig(
    request: DeepPartial<GetUserConfigRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<GetUserConfigResponse> {
    return this.rpc.unary(
      UserConfigGetUserConfigDesc,
      GetUserConfigRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  SetUserConfig(
    request: DeepPartial<SetUserConfigRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<SetUserConfigResponse> {
    return this.rpc.unary(
      UserConfigSetUserConfigDesc,
      SetUserConfigRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }
}

export const UserConfigDesc = { serviceName: "cloud.lazycat.apis.localdevice.UserConfig" };

export const UserConfigGetUserConfigDesc: UnaryMethodDefinitionish = {
  methodName: "GetUserConfig",
  service: UserConfigDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return GetUserConfigRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = GetUserConfigResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const UserConfigSetUserConfigDesc: UnaryMethodDefinitionish = {
  methodName: "SetUserConfig",
  service: UserConfigDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SetUserConfigRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = SetUserConfigResponse.decode(data);
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

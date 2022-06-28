/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../google/protobuf/empty";
import { BrowserHeaders } from "browser-headers";

export interface UserID {
  uid: string;
}

export interface UserInfo {
  uid: string;
  nickname: string;
  avatar: string;
}

export interface ListUsersRequest {}

export interface ListUsersReply {
  uids: string[];
}

function createBaseUserID(): UserID {
  return { uid: "" };
}

export const UserID = {
  encode(
    message: UserID,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserID {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserID();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserID {
    return {
      uid: isSet(object.uid) ? String(object.uid) : "",
    };
  },

  toJSON(message: UserID): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserID>, I>>(object: I): UserID {
    const message = createBaseUserID();
    message.uid = object.uid ?? "";
    return message;
  },
};

function createBaseUserInfo(): UserInfo {
  return { uid: "", nickname: "", avatar: "" };
}

export const UserInfo = {
  encode(
    message: UserInfo,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.nickname !== "") {
      writer.uint32(18).string(message.nickname);
    }
    if (message.avatar !== "") {
      writer.uint32(26).string(message.avatar);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        case 2:
          message.nickname = reader.string();
          break;
        case 3:
          message.avatar = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserInfo {
    return {
      uid: isSet(object.uid) ? String(object.uid) : "",
      nickname: isSet(object.nickname) ? String(object.nickname) : "",
      avatar: isSet(object.avatar) ? String(object.avatar) : "",
    };
  },

  toJSON(message: UserInfo): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.nickname !== undefined && (obj.nickname = message.nickname);
    message.avatar !== undefined && (obj.avatar = message.avatar);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserInfo>, I>>(object: I): UserInfo {
    const message = createBaseUserInfo();
    message.uid = object.uid ?? "";
    message.nickname = object.nickname ?? "";
    message.avatar = object.avatar ?? "";
    return message;
  },
};

function createBaseListUsersRequest(): ListUsersRequest {
  return {};
}

export const ListUsersRequest = {
  encode(
    _: ListUsersRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListUsersRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUsersRequest();
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

  fromJSON(_: any): ListUsersRequest {
    return {};
  },

  toJSON(_: ListUsersRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListUsersRequest>, I>>(
    _: I
  ): ListUsersRequest {
    const message = createBaseListUsersRequest();
    return message;
  },
};

function createBaseListUsersReply(): ListUsersReply {
  return { uids: [] };
}

export const ListUsersReply = {
  encode(
    message: ListUsersReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.uids) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListUsersReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUsersReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uids.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListUsersReply {
    return {
      uids: Array.isArray(object?.uids)
        ? object.uids.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: ListUsersReply): unknown {
    const obj: any = {};
    if (message.uids) {
      obj.uids = message.uids.map((e) => e);
    } else {
      obj.uids = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListUsersReply>, I>>(
    object: I
  ): ListUsersReply {
    const message = createBaseListUsersReply();
    message.uids = object.uids?.map((e) => e) || [];
    return message;
  },
};

export interface UserManager {
  ListUsers(
    request: DeepPartial<ListUsersRequest>,
    metadata?: grpc.Metadata
  ): Promise<ListUsersReply>;
  QueryUserInfo(
    request: DeepPartial<UserID>,
    metadata?: grpc.Metadata
  ): Promise<UserInfo>;
  /** 更新nickname和avator */
  UpdateUserInfo(
    request: DeepPartial<UserInfo>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
}

export class UserManagerClientImpl implements UserManager {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ListUsers = this.ListUsers.bind(this);
    this.QueryUserInfo = this.QueryUserInfo.bind(this);
    this.UpdateUserInfo = this.UpdateUserInfo.bind(this);
  }

  ListUsers(
    request: DeepPartial<ListUsersRequest>,
    metadata?: grpc.Metadata
  ): Promise<ListUsersReply> {
    return this.rpc.unary(
      UserManagerListUsersDesc,
      ListUsersRequest.fromPartial(request),
      metadata
    );
  }

  QueryUserInfo(
    request: DeepPartial<UserID>,
    metadata?: grpc.Metadata
  ): Promise<UserInfo> {
    return this.rpc.unary(
      UserManagerQueryUserInfoDesc,
      UserID.fromPartial(request),
      metadata
    );
  }

  UpdateUserInfo(
    request: DeepPartial<UserInfo>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      UserManagerUpdateUserInfoDesc,
      UserInfo.fromPartial(request),
      metadata
    );
  }
}

export const UserManagerDesc = {
  serviceName: "cloud.lazycat.apis.UserManager",
};

export const UserManagerListUsersDesc: UnaryMethodDefinitionish = {
  methodName: "ListUsers",
  service: UserManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ListUsersRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...ListUsersReply.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const UserManagerQueryUserInfoDesc: UnaryMethodDefinitionish = {
  methodName: "QueryUserInfo",
  service: UserManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return UserID.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...UserInfo.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const UserManagerUpdateUserInfoDesc: UnaryMethodDefinitionish = {
  methodName: "UpdateUserInfo",
  service: UserManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return UserInfo.encode(this).finish();
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

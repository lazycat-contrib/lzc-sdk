/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";
import { Empty } from "../google/protobuf/empty";

export interface Contact {
  id: string;
  name: string;
  phones: string[];
}

export interface ListContactsReply {
  contacts: Contact[];
}

export interface AddContactsRequest {
  ids: string[];
}

export interface AddContactsReply {
}

export interface DeleteContactsRequest {
  ids: string[];
}

export interface DeleteContactsReply {
  /** 删除失败的 id */
  failedIds: string[];
}

export interface UpdateContactsRequest {
  ids: string[];
}

export interface UpdateContactsReply {
  /** 修改失败的 id */
  failedIds: string[];
}

function createBaseContact(): Contact {
  return { id: "", name: "", phones: [] };
}

export const Contact = {
  encode(message: Contact, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    for (const v of message.phones) {
      writer.uint32(26).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Contact {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseContact();
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

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.phones.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Contact {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : "",
      phones: Array.isArray(object?.phones) ? object.phones.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: Contact): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.phones?.length) {
      obj.phones = message.phones;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Contact>, I>>(base?: I): Contact {
    return Contact.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Contact>, I>>(object: I): Contact {
    const message = createBaseContact();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.phones = object.phones?.map((e) => e) || [];
    return message;
  },
};

function createBaseListContactsReply(): ListContactsReply {
  return { contacts: [] };
}

export const ListContactsReply = {
  encode(message: ListContactsReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.contacts) {
      Contact.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListContactsReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListContactsReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.contacts.push(Contact.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListContactsReply {
    return { contacts: Array.isArray(object?.contacts) ? object.contacts.map((e: any) => Contact.fromJSON(e)) : [] };
  },

  toJSON(message: ListContactsReply): unknown {
    const obj: any = {};
    if (message.contacts?.length) {
      obj.contacts = message.contacts.map((e) => Contact.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListContactsReply>, I>>(base?: I): ListContactsReply {
    return ListContactsReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListContactsReply>, I>>(object: I): ListContactsReply {
    const message = createBaseListContactsReply();
    message.contacts = object.contacts?.map((e) => Contact.fromPartial(e)) || [];
    return message;
  },
};

function createBaseAddContactsRequest(): AddContactsRequest {
  return { ids: [] };
}

export const AddContactsRequest = {
  encode(message: AddContactsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.ids) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddContactsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddContactsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.ids.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AddContactsRequest {
    return { ids: Array.isArray(object?.ids) ? object.ids.map((e: any) => String(e)) : [] };
  },

  toJSON(message: AddContactsRequest): unknown {
    const obj: any = {};
    if (message.ids?.length) {
      obj.ids = message.ids;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AddContactsRequest>, I>>(base?: I): AddContactsRequest {
    return AddContactsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AddContactsRequest>, I>>(object: I): AddContactsRequest {
    const message = createBaseAddContactsRequest();
    message.ids = object.ids?.map((e) => e) || [];
    return message;
  },
};

function createBaseAddContactsReply(): AddContactsReply {
  return {};
}

export const AddContactsReply = {
  encode(_: AddContactsReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddContactsReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddContactsReply();
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

  fromJSON(_: any): AddContactsReply {
    return {};
  },

  toJSON(_: AddContactsReply): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<AddContactsReply>, I>>(base?: I): AddContactsReply {
    return AddContactsReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<AddContactsReply>, I>>(_: I): AddContactsReply {
    const message = createBaseAddContactsReply();
    return message;
  },
};

function createBaseDeleteContactsRequest(): DeleteContactsRequest {
  return { ids: [] };
}

export const DeleteContactsRequest = {
  encode(message: DeleteContactsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.ids) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteContactsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteContactsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.ids.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteContactsRequest {
    return { ids: Array.isArray(object?.ids) ? object.ids.map((e: any) => String(e)) : [] };
  },

  toJSON(message: DeleteContactsRequest): unknown {
    const obj: any = {};
    if (message.ids?.length) {
      obj.ids = message.ids;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteContactsRequest>, I>>(base?: I): DeleteContactsRequest {
    return DeleteContactsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteContactsRequest>, I>>(object: I): DeleteContactsRequest {
    const message = createBaseDeleteContactsRequest();
    message.ids = object.ids?.map((e) => e) || [];
    return message;
  },
};

function createBaseDeleteContactsReply(): DeleteContactsReply {
  return { failedIds: [] };
}

export const DeleteContactsReply = {
  encode(message: DeleteContactsReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.failedIds) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteContactsReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteContactsReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.failedIds.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteContactsReply {
    return { failedIds: Array.isArray(object?.failedIds) ? object.failedIds.map((e: any) => String(e)) : [] };
  },

  toJSON(message: DeleteContactsReply): unknown {
    const obj: any = {};
    if (message.failedIds?.length) {
      obj.failedIds = message.failedIds;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteContactsReply>, I>>(base?: I): DeleteContactsReply {
    return DeleteContactsReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeleteContactsReply>, I>>(object: I): DeleteContactsReply {
    const message = createBaseDeleteContactsReply();
    message.failedIds = object.failedIds?.map((e) => e) || [];
    return message;
  },
};

function createBaseUpdateContactsRequest(): UpdateContactsRequest {
  return { ids: [] };
}

export const UpdateContactsRequest = {
  encode(message: UpdateContactsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.ids) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateContactsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateContactsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.ids.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateContactsRequest {
    return { ids: Array.isArray(object?.ids) ? object.ids.map((e: any) => String(e)) : [] };
  },

  toJSON(message: UpdateContactsRequest): unknown {
    const obj: any = {};
    if (message.ids?.length) {
      obj.ids = message.ids;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateContactsRequest>, I>>(base?: I): UpdateContactsRequest {
    return UpdateContactsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateContactsRequest>, I>>(object: I): UpdateContactsRequest {
    const message = createBaseUpdateContactsRequest();
    message.ids = object.ids?.map((e) => e) || [];
    return message;
  },
};

function createBaseUpdateContactsReply(): UpdateContactsReply {
  return { failedIds: [] };
}

export const UpdateContactsReply = {
  encode(message: UpdateContactsReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.failedIds) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateContactsReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateContactsReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.failedIds.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateContactsReply {
    return { failedIds: Array.isArray(object?.failedIds) ? object.failedIds.map((e: any) => String(e)) : [] };
  },

  toJSON(message: UpdateContactsReply): unknown {
    const obj: any = {};
    if (message.failedIds?.length) {
      obj.failedIds = message.failedIds;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateContactsReply>, I>>(base?: I): UpdateContactsReply {
    return UpdateContactsReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UpdateContactsReply>, I>>(object: I): UpdateContactsReply {
    const message = createBaseUpdateContactsReply();
    message.failedIds = object.failedIds?.map((e) => e) || [];
    return message;
  },
};

export interface ContactsManager {
  /** 获取全部联系人列表 */
  ListContacts(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListContactsReply>;
  /** 添加联系人到通讯录 */
  AddContacts(
    request: DeepPartial<AddContactsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<AddContactsReply>;
  /** 从通讯录删除联系人 */
  DeleteContacts(
    request: DeepPartial<DeleteContactsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<DeleteContactsReply>;
  /** 修改通讯录的联系人 */
  UpdateContacts(
    request: DeepPartial<UpdateContactsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<UpdateContactsReply>;
}

export class ContactsManagerClientImpl implements ContactsManager {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ListContacts = this.ListContacts.bind(this);
    this.AddContacts = this.AddContacts.bind(this);
    this.DeleteContacts = this.DeleteContacts.bind(this);
    this.UpdateContacts = this.UpdateContacts.bind(this);
  }

  ListContacts(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListContactsReply> {
    return this.rpc.unary(ContactsManagerListContactsDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  AddContacts(
    request: DeepPartial<AddContactsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<AddContactsReply> {
    return this.rpc.unary(
      ContactsManagerAddContactsDesc,
      AddContactsRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  DeleteContacts(
    request: DeepPartial<DeleteContactsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<DeleteContactsReply> {
    return this.rpc.unary(
      ContactsManagerDeleteContactsDesc,
      DeleteContactsRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  UpdateContacts(
    request: DeepPartial<UpdateContactsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<UpdateContactsReply> {
    return this.rpc.unary(
      ContactsManagerUpdateContactsDesc,
      UpdateContactsRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }
}

export const ContactsManagerDesc = { serviceName: "cloud.lazycat.apis.localdevice.ContactsManager" };

export const ContactsManagerListContactsDesc: UnaryMethodDefinitionish = {
  methodName: "ListContacts",
  service: ContactsManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ListContactsReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const ContactsManagerAddContactsDesc: UnaryMethodDefinitionish = {
  methodName: "AddContacts",
  service: ContactsManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return AddContactsRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = AddContactsReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const ContactsManagerDeleteContactsDesc: UnaryMethodDefinitionish = {
  methodName: "DeleteContacts",
  service: ContactsManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return DeleteContactsRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = DeleteContactsReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const ContactsManagerUpdateContactsDesc: UnaryMethodDefinitionish = {
  methodName: "UpdateContacts",
  service: ContactsManagerDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return UpdateContactsRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = UpdateContactsReply.decode(data);
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

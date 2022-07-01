/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { BrowserHeaders } from "browser-headers";
import { share } from "rxjs/operators";
import { Timestamp } from "../../google/protobuf/timestamp";

export interface PutPhotoReply {
  /** 新添加图片的id */
  id: string;
}

export interface DeletePhotoReply {
  /** 删除失败的图片id */
  failedId: string[];
}

export interface DeletePhotoRequest {
  id: string[];
}

export interface ListPhotoMetasRequest {
  /** 相册ID, 若为空，则表示返回所有相册中的图片 */
  albumIds: string[];
  thumbnailWidth: number;
  thumbnailHeight: number;
  needFileName: boolean;
  needAlbumIds: boolean;
}

export interface PhotoMeta {
  id: string;
  /** 浏览器直接可以使用的url, 可能是device domain下的一个文件或是一个data url */
  photoUrl: string;
  thumbnailUrl: string;
  width: string;
  height: string;
  albumIds: string[];
  fileName?: string | undefined;
  creationDate?: Date | undefined;
  latitude?: number | undefined;
  longitude?: number | undefined;
}

export interface Album {
  id: string;
  title: string;
}

export interface ListAlbumsRequest {}

export interface ListAlbumsReply {
  albums: Album[];
}

export interface PutPhotoRequest {
  albumId: string;
  /** 图片路径, 支持dataurl */
  url: string;
  fileName?: string | undefined;
}

function createBasePutPhotoReply(): PutPhotoReply {
  return { id: "" };
}

export const PutPhotoReply = {
  encode(
    message: PutPhotoReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PutPhotoReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePutPhotoReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PutPhotoReply {
    return {
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: PutPhotoReply): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PutPhotoReply>, I>>(
    object: I
  ): PutPhotoReply {
    const message = createBasePutPhotoReply();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseDeletePhotoReply(): DeletePhotoReply {
  return { failedId: [] };
}

export const DeletePhotoReply = {
  encode(
    message: DeletePhotoReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.failedId) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletePhotoReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePhotoReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.failedId.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeletePhotoReply {
    return {
      failedId: Array.isArray(object?.failedId)
        ? object.failedId.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: DeletePhotoReply): unknown {
    const obj: any = {};
    if (message.failedId) {
      obj.failedId = message.failedId.map((e) => e);
    } else {
      obj.failedId = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeletePhotoReply>, I>>(
    object: I
  ): DeletePhotoReply {
    const message = createBaseDeletePhotoReply();
    message.failedId = object.failedId?.map((e) => e) || [];
    return message;
  },
};

function createBaseDeletePhotoRequest(): DeletePhotoRequest {
  return { id: [] };
}

export const DeletePhotoRequest = {
  encode(
    message: DeletePhotoRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.id) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletePhotoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePhotoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeletePhotoRequest {
    return {
      id: Array.isArray(object?.id) ? object.id.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: DeletePhotoRequest): unknown {
    const obj: any = {};
    if (message.id) {
      obj.id = message.id.map((e) => e);
    } else {
      obj.id = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeletePhotoRequest>, I>>(
    object: I
  ): DeletePhotoRequest {
    const message = createBaseDeletePhotoRequest();
    message.id = object.id?.map((e) => e) || [];
    return message;
  },
};

function createBaseListPhotoMetasRequest(): ListPhotoMetasRequest {
  return {
    albumIds: [],
    thumbnailWidth: 0,
    thumbnailHeight: 0,
    needFileName: false,
    needAlbumIds: false,
  };
}

export const ListPhotoMetasRequest = {
  encode(
    message: ListPhotoMetasRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.albumIds) {
      writer.uint32(10).string(v!);
    }
    if (message.thumbnailWidth !== 0) {
      writer.uint32(16).int32(message.thumbnailWidth);
    }
    if (message.thumbnailHeight !== 0) {
      writer.uint32(24).int32(message.thumbnailHeight);
    }
    if (message.needFileName === true) {
      writer.uint32(32).bool(message.needFileName);
    }
    if (message.needAlbumIds === true) {
      writer.uint32(40).bool(message.needAlbumIds);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ListPhotoMetasRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPhotoMetasRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.albumIds.push(reader.string());
          break;
        case 2:
          message.thumbnailWidth = reader.int32();
          break;
        case 3:
          message.thumbnailHeight = reader.int32();
          break;
        case 4:
          message.needFileName = reader.bool();
          break;
        case 5:
          message.needAlbumIds = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListPhotoMetasRequest {
    return {
      albumIds: Array.isArray(object?.albumIds)
        ? object.albumIds.map((e: any) => String(e))
        : [],
      thumbnailWidth: isSet(object.thumbnailWidth)
        ? Number(object.thumbnailWidth)
        : 0,
      thumbnailHeight: isSet(object.thumbnailHeight)
        ? Number(object.thumbnailHeight)
        : 0,
      needFileName: isSet(object.needFileName)
        ? Boolean(object.needFileName)
        : false,
      needAlbumIds: isSet(object.needAlbumIds)
        ? Boolean(object.needAlbumIds)
        : false,
    };
  },

  toJSON(message: ListPhotoMetasRequest): unknown {
    const obj: any = {};
    if (message.albumIds) {
      obj.albumIds = message.albumIds.map((e) => e);
    } else {
      obj.albumIds = [];
    }
    message.thumbnailWidth !== undefined &&
      (obj.thumbnailWidth = Math.round(message.thumbnailWidth));
    message.thumbnailHeight !== undefined &&
      (obj.thumbnailHeight = Math.round(message.thumbnailHeight));
    message.needFileName !== undefined &&
      (obj.needFileName = message.needFileName);
    message.needAlbumIds !== undefined &&
      (obj.needAlbumIds = message.needAlbumIds);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListPhotoMetasRequest>, I>>(
    object: I
  ): ListPhotoMetasRequest {
    const message = createBaseListPhotoMetasRequest();
    message.albumIds = object.albumIds?.map((e) => e) || [];
    message.thumbnailWidth = object.thumbnailWidth ?? 0;
    message.thumbnailHeight = object.thumbnailHeight ?? 0;
    message.needFileName = object.needFileName ?? false;
    message.needAlbumIds = object.needAlbumIds ?? false;
    return message;
  },
};

function createBasePhotoMeta(): PhotoMeta {
  return {
    id: "",
    photoUrl: "",
    thumbnailUrl: "",
    width: "",
    height: "",
    albumIds: [],
    fileName: undefined,
    creationDate: undefined,
    latitude: undefined,
    longitude: undefined,
  };
}

export const PhotoMeta = {
  encode(
    message: PhotoMeta,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.photoUrl !== "") {
      writer.uint32(18).string(message.photoUrl);
    }
    if (message.thumbnailUrl !== "") {
      writer.uint32(26).string(message.thumbnailUrl);
    }
    if (message.width !== "") {
      writer.uint32(34).string(message.width);
    }
    if (message.height !== "") {
      writer.uint32(42).string(message.height);
    }
    for (const v of message.albumIds) {
      writer.uint32(50).string(v!);
    }
    if (message.fileName !== undefined) {
      writer.uint32(58).string(message.fileName);
    }
    if (message.creationDate !== undefined) {
      Timestamp.encode(
        toTimestamp(message.creationDate),
        writer.uint32(66).fork()
      ).ldelim();
    }
    if (message.latitude !== undefined) {
      writer.uint32(77).float(message.latitude);
    }
    if (message.longitude !== undefined) {
      writer.uint32(85).float(message.longitude);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PhotoMeta {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhotoMeta();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.photoUrl = reader.string();
          break;
        case 3:
          message.thumbnailUrl = reader.string();
          break;
        case 4:
          message.width = reader.string();
          break;
        case 5:
          message.height = reader.string();
          break;
        case 6:
          message.albumIds.push(reader.string());
          break;
        case 7:
          message.fileName = reader.string();
          break;
        case 8:
          message.creationDate = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 9:
          message.latitude = reader.float();
          break;
        case 10:
          message.longitude = reader.float();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PhotoMeta {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      photoUrl: isSet(object.photoUrl) ? String(object.photoUrl) : "",
      thumbnailUrl: isSet(object.thumbnailUrl)
        ? String(object.thumbnailUrl)
        : "",
      width: isSet(object.width) ? String(object.width) : "",
      height: isSet(object.height) ? String(object.height) : "",
      albumIds: Array.isArray(object?.albumIds)
        ? object.albumIds.map((e: any) => String(e))
        : [],
      fileName: isSet(object.fileName) ? String(object.fileName) : undefined,
      creationDate: isSet(object.creationDate)
        ? fromJsonTimestamp(object.creationDate)
        : undefined,
      latitude: isSet(object.latitude) ? Number(object.latitude) : undefined,
      longitude: isSet(object.longitude) ? Number(object.longitude) : undefined,
    };
  },

  toJSON(message: PhotoMeta): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.photoUrl !== undefined && (obj.photoUrl = message.photoUrl);
    message.thumbnailUrl !== undefined &&
      (obj.thumbnailUrl = message.thumbnailUrl);
    message.width !== undefined && (obj.width = message.width);
    message.height !== undefined && (obj.height = message.height);
    if (message.albumIds) {
      obj.albumIds = message.albumIds.map((e) => e);
    } else {
      obj.albumIds = [];
    }
    message.fileName !== undefined && (obj.fileName = message.fileName);
    message.creationDate !== undefined &&
      (obj.creationDate = message.creationDate.toISOString());
    message.latitude !== undefined && (obj.latitude = message.latitude);
    message.longitude !== undefined && (obj.longitude = message.longitude);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PhotoMeta>, I>>(
    object: I
  ): PhotoMeta {
    const message = createBasePhotoMeta();
    message.id = object.id ?? "";
    message.photoUrl = object.photoUrl ?? "";
    message.thumbnailUrl = object.thumbnailUrl ?? "";
    message.width = object.width ?? "";
    message.height = object.height ?? "";
    message.albumIds = object.albumIds?.map((e) => e) || [];
    message.fileName = object.fileName ?? undefined;
    message.creationDate = object.creationDate ?? undefined;
    message.latitude = object.latitude ?? undefined;
    message.longitude = object.longitude ?? undefined;
    return message;
  },
};

function createBaseAlbum(): Album {
  return { id: "", title: "" };
}

export const Album = {
  encode(message: Album, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Album {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAlbum();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Album {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      title: isSet(object.title) ? String(object.title) : "",
    };
  },

  toJSON(message: Album): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.title !== undefined && (obj.title = message.title);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Album>, I>>(object: I): Album {
    const message = createBaseAlbum();
    message.id = object.id ?? "";
    message.title = object.title ?? "";
    return message;
  },
};

function createBaseListAlbumsRequest(): ListAlbumsRequest {
  return {};
}

export const ListAlbumsRequest = {
  encode(
    _: ListAlbumsRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAlbumsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAlbumsRequest();
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

  fromJSON(_: any): ListAlbumsRequest {
    return {};
  },

  toJSON(_: ListAlbumsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListAlbumsRequest>, I>>(
    _: I
  ): ListAlbumsRequest {
    const message = createBaseListAlbumsRequest();
    return message;
  },
};

function createBaseListAlbumsReply(): ListAlbumsReply {
  return { albums: [] };
}

export const ListAlbumsReply = {
  encode(
    message: ListAlbumsReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.albums) {
      Album.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAlbumsReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAlbumsReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.albums.push(Album.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListAlbumsReply {
    return {
      albums: Array.isArray(object?.albums)
        ? object.albums.map((e: any) => Album.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ListAlbumsReply): unknown {
    const obj: any = {};
    if (message.albums) {
      obj.albums = message.albums.map((e) => (e ? Album.toJSON(e) : undefined));
    } else {
      obj.albums = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListAlbumsReply>, I>>(
    object: I
  ): ListAlbumsReply {
    const message = createBaseListAlbumsReply();
    message.albums = object.albums?.map((e) => Album.fromPartial(e)) || [];
    return message;
  },
};

function createBasePutPhotoRequest(): PutPhotoRequest {
  return { albumId: "", url: "", fileName: undefined };
}

export const PutPhotoRequest = {
  encode(
    message: PutPhotoRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.albumId !== "") {
      writer.uint32(10).string(message.albumId);
    }
    if (message.url !== "") {
      writer.uint32(18).string(message.url);
    }
    if (message.fileName !== undefined) {
      writer.uint32(26).string(message.fileName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PutPhotoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePutPhotoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.albumId = reader.string();
          break;
        case 2:
          message.url = reader.string();
          break;
        case 3:
          message.fileName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PutPhotoRequest {
    return {
      albumId: isSet(object.albumId) ? String(object.albumId) : "",
      url: isSet(object.url) ? String(object.url) : "",
      fileName: isSet(object.fileName) ? String(object.fileName) : undefined,
    };
  },

  toJSON(message: PutPhotoRequest): unknown {
    const obj: any = {};
    message.albumId !== undefined && (obj.albumId = message.albumId);
    message.url !== undefined && (obj.url = message.url);
    message.fileName !== undefined && (obj.fileName = message.fileName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PutPhotoRequest>, I>>(
    object: I
  ): PutPhotoRequest {
    const message = createBasePutPhotoRequest();
    message.albumId = object.albumId ?? "";
    message.url = object.url ?? "";
    message.fileName = object.fileName ?? undefined;
    return message;
  },
};

export interface PhotoLibrary {
  /** 列举所有的系统相册 */
  ListAlbums(
    request: DeepPartial<ListAlbumsRequest>,
    metadata?: grpc.Metadata
  ): Promise<ListAlbumsReply>;
  /** 存储一张图片到某个相册中 */
  PutPhoto(
    request: DeepPartial<PutPhotoRequest>,
    metadata?: grpc.Metadata
  ): Promise<PutPhotoReply>;
  DeletePhoto(
    request: DeepPartial<DeletePhotoRequest>,
    metadata?: grpc.Metadata
  ): Promise<DeletePhotoReply>;
  /** 枚举具体相册中的图片元信息 */
  ListPhotoMetas(
    request: DeepPartial<ListPhotoMetasRequest>,
    metadata?: grpc.Metadata
  ): Observable<PhotoMeta>;
}

export class PhotoLibraryClientImpl implements PhotoLibrary {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ListAlbums = this.ListAlbums.bind(this);
    this.PutPhoto = this.PutPhoto.bind(this);
    this.DeletePhoto = this.DeletePhoto.bind(this);
    this.ListPhotoMetas = this.ListPhotoMetas.bind(this);
  }

  ListAlbums(
    request: DeepPartial<ListAlbumsRequest>,
    metadata?: grpc.Metadata
  ): Promise<ListAlbumsReply> {
    return this.rpc.unary(
      PhotoLibraryListAlbumsDesc,
      ListAlbumsRequest.fromPartial(request),
      metadata
    );
  }

  PutPhoto(
    request: DeepPartial<PutPhotoRequest>,
    metadata?: grpc.Metadata
  ): Promise<PutPhotoReply> {
    return this.rpc.unary(
      PhotoLibraryPutPhotoDesc,
      PutPhotoRequest.fromPartial(request),
      metadata
    );
  }

  DeletePhoto(
    request: DeepPartial<DeletePhotoRequest>,
    metadata?: grpc.Metadata
  ): Promise<DeletePhotoReply> {
    return this.rpc.unary(
      PhotoLibraryDeletePhotoDesc,
      DeletePhotoRequest.fromPartial(request),
      metadata
    );
  }

  ListPhotoMetas(
    request: DeepPartial<ListPhotoMetasRequest>,
    metadata?: grpc.Metadata
  ): Observable<PhotoMeta> {
    return this.rpc.invoke(
      PhotoLibraryListPhotoMetasDesc,
      ListPhotoMetasRequest.fromPartial(request),
      metadata
    );
  }
}

export const PhotoLibraryDesc = {
  serviceName: "cloud.lazycat.localdevice.apis.PhotoLibrary",
};

export const PhotoLibraryListAlbumsDesc: UnaryMethodDefinitionish = {
  methodName: "ListAlbums",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ListAlbumsRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...ListAlbumsReply.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const PhotoLibraryPutPhotoDesc: UnaryMethodDefinitionish = {
  methodName: "PutPhoto",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return PutPhotoRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...PutPhotoReply.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const PhotoLibraryDeletePhotoDesc: UnaryMethodDefinitionish = {
  methodName: "DeletePhoto",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return DeletePhotoRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...DeletePhotoReply.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const PhotoLibraryListPhotoMetasDesc: UnaryMethodDefinitionish = {
  methodName: "ListPhotoMetas",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return ListPhotoMetasRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...PhotoMeta.decode(data),
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

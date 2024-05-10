/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { share } from "rxjs/operators";
import { Empty } from "../google/protobuf/empty";
import { Timestamp } from "../google/protobuf/timestamp";

/** 查询照片信息的排序方式 */
export enum PhotoMetasSortType {
  /** CREATE_DATE_DESC - 创建时间降序 */
  CREATE_DATE_DESC = 0,
  /** TAKEN_DATE_DESC - 拍照时间降序 */
  TAKEN_DATE_DESC = 1,
  /** CREATE_DATE_ASC - 创建时间升序 */
  CREATE_DATE_ASC = 2,
  /** TAKEN_DATE_ASC - 拍摄时间升序 */
  TAKEN_DATE_ASC = 3,
  /** CREATE_TAKEN_DESC - 创建时间,拍摄时间降序 */
  CREATE_TAKEN_DESC = 4,
  /** CREATE_TAKEN_ASC - 创建时间,拍摄时间升序 */
  CREATE_TAKEN_ASC = 5,
  /** TOKEN_CREATE_DESC - 拍摄时间，创建时间 降序 */
  TOKEN_CREATE_DESC = 6,
  /** TOKEN_CREATE_ASC - 拍摄时间，创建时间升序 */
  TOKEN_CREATE_ASC = 7,
  UNRECOGNIZED = -1,
}

export function photoMetasSortTypeFromJSON(object: any): PhotoMetasSortType {
  switch (object) {
    case 0:
    case "CREATE_DATE_DESC":
      return PhotoMetasSortType.CREATE_DATE_DESC;
    case 1:
    case "TAKEN_DATE_DESC":
      return PhotoMetasSortType.TAKEN_DATE_DESC;
    case 2:
    case "CREATE_DATE_ASC":
      return PhotoMetasSortType.CREATE_DATE_ASC;
    case 3:
    case "TAKEN_DATE_ASC":
      return PhotoMetasSortType.TAKEN_DATE_ASC;
    case 4:
    case "CREATE_TAKEN_DESC":
      return PhotoMetasSortType.CREATE_TAKEN_DESC;
    case 5:
    case "CREATE_TAKEN_ASC":
      return PhotoMetasSortType.CREATE_TAKEN_ASC;
    case 6:
    case "TOKEN_CREATE_DESC":
      return PhotoMetasSortType.TOKEN_CREATE_DESC;
    case 7:
    case "TOKEN_CREATE_ASC":
      return PhotoMetasSortType.TOKEN_CREATE_ASC;
    case -1:
    case "UNRECOGNIZED":
    default:
      return PhotoMetasSortType.UNRECOGNIZED;
  }
}

export function photoMetasSortTypeToJSON(object: PhotoMetasSortType): string {
  switch (object) {
    case PhotoMetasSortType.CREATE_DATE_DESC:
      return "CREATE_DATE_DESC";
    case PhotoMetasSortType.TAKEN_DATE_DESC:
      return "TAKEN_DATE_DESC";
    case PhotoMetasSortType.CREATE_DATE_ASC:
      return "CREATE_DATE_ASC";
    case PhotoMetasSortType.TAKEN_DATE_ASC:
      return "TAKEN_DATE_ASC";
    case PhotoMetasSortType.CREATE_TAKEN_DESC:
      return "CREATE_TAKEN_DESC";
    case PhotoMetasSortType.CREATE_TAKEN_ASC:
      return "CREATE_TAKEN_ASC";
    case PhotoMetasSortType.TOKEN_CREATE_DESC:
      return "TOKEN_CREATE_DESC";
    case PhotoMetasSortType.TOKEN_CREATE_ASC:
      return "TOKEN_CREATE_ASC";
    case PhotoMetasSortType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

/** 查询照片信息的排序方式 */
export enum ListAssetsSortType {
  /** ASSETS_CREATE_DATE_DESC - 创建时间降序 */
  ASSETS_CREATE_DATE_DESC = 0,
  /** ASSETS_CREATE_DATE_ASC - 创建时间升序 */
  ASSETS_CREATE_DATE_ASC = 1,
  /** ASSETS_TIMELINE_DATE_DESC - 时间线降序 */
  ASSETS_TIMELINE_DATE_DESC = 2,
  /** ASSETS_TIMELINE_DATE_ASC - 时间线降序 */
  ASSETS_TIMELINE_DATE_ASC = 3,
  UNRECOGNIZED = -1,
}

export function listAssetsSortTypeFromJSON(object: any): ListAssetsSortType {
  switch (object) {
    case 0:
    case "ASSETS_CREATE_DATE_DESC":
      return ListAssetsSortType.ASSETS_CREATE_DATE_DESC;
    case 1:
    case "ASSETS_CREATE_DATE_ASC":
      return ListAssetsSortType.ASSETS_CREATE_DATE_ASC;
    case 2:
    case "ASSETS_TIMELINE_DATE_DESC":
      return ListAssetsSortType.ASSETS_TIMELINE_DATE_DESC;
    case 3:
    case "ASSETS_TIMELINE_DATE_ASC":
      return ListAssetsSortType.ASSETS_TIMELINE_DATE_ASC;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ListAssetsSortType.UNRECOGNIZED;
  }
}

export function listAssetsSortTypeToJSON(object: ListAssetsSortType): string {
  switch (object) {
    case ListAssetsSortType.ASSETS_CREATE_DATE_DESC:
      return "ASSETS_CREATE_DATE_DESC";
    case ListAssetsSortType.ASSETS_CREATE_DATE_ASC:
      return "ASSETS_CREATE_DATE_ASC";
    case ListAssetsSortType.ASSETS_TIMELINE_DATE_DESC:
      return "ASSETS_TIMELINE_DATE_DESC";
    case ListAssetsSortType.ASSETS_TIMELINE_DATE_ASC:
      return "ASSETS_TIMELINE_DATE_ASC";
    case ListAssetsSortType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface QueryPhotoHashRequest {
  /**
   * changed_after语义为: changed_after之后的新增文件都会加入返回的集合里,
   * 与文件本身的modify time无关。
   */
  changedAfter: Date | undefined;
}

export interface QueryPhotoRequest {
  id: string;
}

export interface MakeAlbumRequest {
  title: string;
}

export interface PutPhotoRequest {
  albumId: string;
  /** 图片路径, 支持dataurl */
  url: string;
  fileName?: string | undefined;
}

export interface PutPhotoReply {
  /** 是否完成. 若已完成则photo_id字段生效 */
  done: boolean;
  /** 若无total_size表示无法获取文件大小，此时无法计算出准确进度 */
  totalSize?: number | undefined;
  fetchedSize?:
    | number
    | undefined;
  /** 新添加图片的id, 仅在completed=true后才有意义 */
  photoId?: string | undefined;
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
  /** 不填写默认按照创建时间降序排序 */
  stype?:
    | PhotoMetasSortType
    | undefined;
  /** 查询第几页(不填写默认为第一页) */
  pageNum?:
    | number
    | undefined;
  /** （每页返回的数据条数，不填写将返回所有数据） */
  pageSize?:
    | number
    | undefined;
  /** 是否过滤视频 （默认false,照片和视频信息一起返回) */
  isFilterVideo?:
    | boolean
    | undefined;
  /** 给定文件的修改时间过滤meta信息返回 */
  filterByModifiedDate?: Date | undefined;
}

export interface ListAssetsRequest {
  /** 相册ID， 若为空，则表示返回所有相册中的图片 */
  albumIds: string[];
  /** 是否过滤视频 （默认false,照片和视频信息一起返回) */
  isFilterVideo?:
    | boolean
    | undefined;
  /** 排序方法 */
  sortBy?:
    | ListAssetsSortType
    | undefined;
  /** 给定指定日期后的资源，为空表示获取全部 */
  startDate?: Date | undefined;
}

export interface PhotoMeta {
  id: string;
  /** 浏览器直接可以使用的url, 可能是device domain下的一个文件或是一个data url */
  photoUrl: string;
  thumbnailUrl: string;
  width: string;
  height: string;
  albumIds: string[];
  size: number;
  mime: string;
  fileName?: string | undefined;
  creationDate?: Date | undefined;
  latitude?: number | undefined;
  longitude?: number | undefined;
  takenDate?: Date | undefined;
  modifiedDate?: Date | undefined;
}

export interface Album {
  id: string;
  title: string;
  imageCount: number;
  videoCount: number;
  coverImageUrl: string;
  /** 封面文件(图）id */
  coverImageId: string;
}

export interface ListAlbumsRequest {
  thumbnailWidth?: number | undefined;
  thumbnailHeight?: number | undefined;
  thumbnailChooseWay?: number | undefined;
}

export interface ListAlbumsReply {
  albums: Album[];
}

export interface ListAssetStatsRequest {
  /** 是否过滤视频 （默认false,照片和视频信息一起返回) */
  isFilterVideo: boolean;
  /** 排序方法 */
  sortBy: ListAssetsSortType;
  startDate?: Date | undefined;
}

export interface ListAssetStatsReply {
  id: string;
  date: Date | undefined;
  mime?: string | undefined;
}

export interface ListAssetsByIdsRequest {
  ids: string[];
  sortBy: ListAssetsSortType;
}

export interface QueryAssetUrlPathReply {
  assetUrl: string;
  thumbnailUrl: string;
}

function createBaseQueryPhotoHashRequest(): QueryPhotoHashRequest {
  return { changedAfter: undefined };
}

export const QueryPhotoHashRequest = {
  encode(message: QueryPhotoHashRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.changedAfter !== undefined) {
      Timestamp.encode(toTimestamp(message.changedAfter), writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryPhotoHashRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryPhotoHashRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.changedAfter = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryPhotoHashRequest {
    return { changedAfter: isSet(object.changedAfter) ? fromJsonTimestamp(object.changedAfter) : undefined };
  },

  toJSON(message: QueryPhotoHashRequest): unknown {
    const obj: any = {};
    if (message.changedAfter !== undefined) {
      obj.changedAfter = message.changedAfter.toISOString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryPhotoHashRequest>, I>>(base?: I): QueryPhotoHashRequest {
    return QueryPhotoHashRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<QueryPhotoHashRequest>, I>>(object: I): QueryPhotoHashRequest {
    const message = createBaseQueryPhotoHashRequest();
    message.changedAfter = object.changedAfter ?? undefined;
    return message;
  },
};

function createBaseQueryPhotoRequest(): QueryPhotoRequest {
  return { id: "" };
}

export const QueryPhotoRequest = {
  encode(message: QueryPhotoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryPhotoRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryPhotoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryPhotoRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: QueryPhotoRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryPhotoRequest>, I>>(base?: I): QueryPhotoRequest {
    return QueryPhotoRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<QueryPhotoRequest>, I>>(object: I): QueryPhotoRequest {
    const message = createBaseQueryPhotoRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseMakeAlbumRequest(): MakeAlbumRequest {
  return { title: "" };
}

export const MakeAlbumRequest = {
  encode(message: MakeAlbumRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MakeAlbumRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMakeAlbumRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.title = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MakeAlbumRequest {
    return { title: isSet(object.title) ? String(object.title) : "" };
  },

  toJSON(message: MakeAlbumRequest): unknown {
    const obj: any = {};
    if (message.title !== "") {
      obj.title = message.title;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MakeAlbumRequest>, I>>(base?: I): MakeAlbumRequest {
    return MakeAlbumRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<MakeAlbumRequest>, I>>(object: I): MakeAlbumRequest {
    const message = createBaseMakeAlbumRequest();
    message.title = object.title ?? "";
    return message;
  },
};

function createBasePutPhotoRequest(): PutPhotoRequest {
  return { albumId: "", url: "", fileName: undefined };
}

export const PutPhotoRequest = {
  encode(message: PutPhotoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePutPhotoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.albumId = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.url = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.fileName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
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
    if (message.albumId !== "") {
      obj.albumId = message.albumId;
    }
    if (message.url !== "") {
      obj.url = message.url;
    }
    if (message.fileName !== undefined) {
      obj.fileName = message.fileName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PutPhotoRequest>, I>>(base?: I): PutPhotoRequest {
    return PutPhotoRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PutPhotoRequest>, I>>(object: I): PutPhotoRequest {
    const message = createBasePutPhotoRequest();
    message.albumId = object.albumId ?? "";
    message.url = object.url ?? "";
    message.fileName = object.fileName ?? undefined;
    return message;
  },
};

function createBasePutPhotoReply(): PutPhotoReply {
  return { done: false, totalSize: undefined, fetchedSize: undefined, photoId: undefined };
}

export const PutPhotoReply = {
  encode(message: PutPhotoReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.done === true) {
      writer.uint32(8).bool(message.done);
    }
    if (message.totalSize !== undefined) {
      writer.uint32(16).int64(message.totalSize);
    }
    if (message.fetchedSize !== undefined) {
      writer.uint32(24).int64(message.fetchedSize);
    }
    if (message.photoId !== undefined) {
      writer.uint32(34).string(message.photoId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PutPhotoReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePutPhotoReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.done = reader.bool();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.totalSize = longToNumber(reader.int64() as Long);
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.fetchedSize = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.photoId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PutPhotoReply {
    return {
      done: isSet(object.done) ? Boolean(object.done) : false,
      totalSize: isSet(object.totalSize) ? Number(object.totalSize) : undefined,
      fetchedSize: isSet(object.fetchedSize) ? Number(object.fetchedSize) : undefined,
      photoId: isSet(object.photoId) ? String(object.photoId) : undefined,
    };
  },

  toJSON(message: PutPhotoReply): unknown {
    const obj: any = {};
    if (message.done === true) {
      obj.done = message.done;
    }
    if (message.totalSize !== undefined) {
      obj.totalSize = Math.round(message.totalSize);
    }
    if (message.fetchedSize !== undefined) {
      obj.fetchedSize = Math.round(message.fetchedSize);
    }
    if (message.photoId !== undefined) {
      obj.photoId = message.photoId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PutPhotoReply>, I>>(base?: I): PutPhotoReply {
    return PutPhotoReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PutPhotoReply>, I>>(object: I): PutPhotoReply {
    const message = createBasePutPhotoReply();
    message.done = object.done ?? false;
    message.totalSize = object.totalSize ?? undefined;
    message.fetchedSize = object.fetchedSize ?? undefined;
    message.photoId = object.photoId ?? undefined;
    return message;
  },
};

function createBaseDeletePhotoReply(): DeletePhotoReply {
  return { failedId: [] };
}

export const DeletePhotoReply = {
  encode(message: DeletePhotoReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.failedId) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletePhotoReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePhotoReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.failedId.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeletePhotoReply {
    return { failedId: Array.isArray(object?.failedId) ? object.failedId.map((e: any) => String(e)) : [] };
  },

  toJSON(message: DeletePhotoReply): unknown {
    const obj: any = {};
    if (message.failedId?.length) {
      obj.failedId = message.failedId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeletePhotoReply>, I>>(base?: I): DeletePhotoReply {
    return DeletePhotoReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeletePhotoReply>, I>>(object: I): DeletePhotoReply {
    const message = createBaseDeletePhotoReply();
    message.failedId = object.failedId?.map((e) => e) || [];
    return message;
  },
};

function createBaseDeletePhotoRequest(): DeletePhotoRequest {
  return { id: [] };
}

export const DeletePhotoRequest = {
  encode(message: DeletePhotoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.id) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeletePhotoRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeletePhotoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeletePhotoRequest {
    return { id: Array.isArray(object?.id) ? object.id.map((e: any) => String(e)) : [] };
  },

  toJSON(message: DeletePhotoRequest): unknown {
    const obj: any = {};
    if (message.id?.length) {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeletePhotoRequest>, I>>(base?: I): DeletePhotoRequest {
    return DeletePhotoRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DeletePhotoRequest>, I>>(object: I): DeletePhotoRequest {
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
    stype: undefined,
    pageNum: undefined,
    pageSize: undefined,
    isFilterVideo: undefined,
    filterByModifiedDate: undefined,
  };
}

export const ListPhotoMetasRequest = {
  encode(message: ListPhotoMetasRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.stype !== undefined) {
      writer.uint32(48).int32(message.stype);
    }
    if (message.pageNum !== undefined) {
      writer.uint32(56).uint32(message.pageNum);
    }
    if (message.pageSize !== undefined) {
      writer.uint32(64).uint32(message.pageSize);
    }
    if (message.isFilterVideo !== undefined) {
      writer.uint32(72).bool(message.isFilterVideo);
    }
    if (message.filterByModifiedDate !== undefined) {
      Timestamp.encode(toTimestamp(message.filterByModifiedDate), writer.uint32(82).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListPhotoMetasRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListPhotoMetasRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.albumIds.push(reader.string());
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.thumbnailWidth = reader.int32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.thumbnailHeight = reader.int32();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.needFileName = reader.bool();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.needAlbumIds = reader.bool();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.stype = reader.int32() as any;
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.pageNum = reader.uint32();
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.pageSize = reader.uint32();
          continue;
        case 9:
          if (tag !== 72) {
            break;
          }

          message.isFilterVideo = reader.bool();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.filterByModifiedDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListPhotoMetasRequest {
    return {
      albumIds: Array.isArray(object?.albumIds) ? object.albumIds.map((e: any) => String(e)) : [],
      thumbnailWidth: isSet(object.thumbnailWidth) ? Number(object.thumbnailWidth) : 0,
      thumbnailHeight: isSet(object.thumbnailHeight) ? Number(object.thumbnailHeight) : 0,
      needFileName: isSet(object.needFileName) ? Boolean(object.needFileName) : false,
      needAlbumIds: isSet(object.needAlbumIds) ? Boolean(object.needAlbumIds) : false,
      stype: isSet(object.stype) ? photoMetasSortTypeFromJSON(object.stype) : undefined,
      pageNum: isSet(object.pageNum) ? Number(object.pageNum) : undefined,
      pageSize: isSet(object.pageSize) ? Number(object.pageSize) : undefined,
      isFilterVideo: isSet(object.isFilterVideo) ? Boolean(object.isFilterVideo) : undefined,
      filterByModifiedDate: isSet(object.filterByModifiedDate)
        ? fromJsonTimestamp(object.filterByModifiedDate)
        : undefined,
    };
  },

  toJSON(message: ListPhotoMetasRequest): unknown {
    const obj: any = {};
    if (message.albumIds?.length) {
      obj.albumIds = message.albumIds;
    }
    if (message.thumbnailWidth !== 0) {
      obj.thumbnailWidth = Math.round(message.thumbnailWidth);
    }
    if (message.thumbnailHeight !== 0) {
      obj.thumbnailHeight = Math.round(message.thumbnailHeight);
    }
    if (message.needFileName === true) {
      obj.needFileName = message.needFileName;
    }
    if (message.needAlbumIds === true) {
      obj.needAlbumIds = message.needAlbumIds;
    }
    if (message.stype !== undefined) {
      obj.stype = photoMetasSortTypeToJSON(message.stype);
    }
    if (message.pageNum !== undefined) {
      obj.pageNum = Math.round(message.pageNum);
    }
    if (message.pageSize !== undefined) {
      obj.pageSize = Math.round(message.pageSize);
    }
    if (message.isFilterVideo !== undefined) {
      obj.isFilterVideo = message.isFilterVideo;
    }
    if (message.filterByModifiedDate !== undefined) {
      obj.filterByModifiedDate = message.filterByModifiedDate.toISOString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListPhotoMetasRequest>, I>>(base?: I): ListPhotoMetasRequest {
    return ListPhotoMetasRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListPhotoMetasRequest>, I>>(object: I): ListPhotoMetasRequest {
    const message = createBaseListPhotoMetasRequest();
    message.albumIds = object.albumIds?.map((e) => e) || [];
    message.thumbnailWidth = object.thumbnailWidth ?? 0;
    message.thumbnailHeight = object.thumbnailHeight ?? 0;
    message.needFileName = object.needFileName ?? false;
    message.needAlbumIds = object.needAlbumIds ?? false;
    message.stype = object.stype ?? undefined;
    message.pageNum = object.pageNum ?? undefined;
    message.pageSize = object.pageSize ?? undefined;
    message.isFilterVideo = object.isFilterVideo ?? undefined;
    message.filterByModifiedDate = object.filterByModifiedDate ?? undefined;
    return message;
  },
};

function createBaseListAssetsRequest(): ListAssetsRequest {
  return { albumIds: [], isFilterVideo: undefined, sortBy: undefined, startDate: undefined };
}

export const ListAssetsRequest = {
  encode(message: ListAssetsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.albumIds) {
      writer.uint32(10).string(v!);
    }
    if (message.isFilterVideo !== undefined) {
      writer.uint32(16).bool(message.isFilterVideo);
    }
    if (message.sortBy !== undefined) {
      writer.uint32(24).int32(message.sortBy);
    }
    if (message.startDate !== undefined) {
      Timestamp.encode(toTimestamp(message.startDate), writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAssetsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAssetsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.albumIds.push(reader.string());
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.isFilterVideo = reader.bool();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.sortBy = reader.int32() as any;
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.startDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListAssetsRequest {
    return {
      albumIds: Array.isArray(object?.albumIds) ? object.albumIds.map((e: any) => String(e)) : [],
      isFilterVideo: isSet(object.isFilterVideo) ? Boolean(object.isFilterVideo) : undefined,
      sortBy: isSet(object.sortBy) ? listAssetsSortTypeFromJSON(object.sortBy) : undefined,
      startDate: isSet(object.startDate) ? fromJsonTimestamp(object.startDate) : undefined,
    };
  },

  toJSON(message: ListAssetsRequest): unknown {
    const obj: any = {};
    if (message.albumIds?.length) {
      obj.albumIds = message.albumIds;
    }
    if (message.isFilterVideo !== undefined) {
      obj.isFilterVideo = message.isFilterVideo;
    }
    if (message.sortBy !== undefined) {
      obj.sortBy = listAssetsSortTypeToJSON(message.sortBy);
    }
    if (message.startDate !== undefined) {
      obj.startDate = message.startDate.toISOString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListAssetsRequest>, I>>(base?: I): ListAssetsRequest {
    return ListAssetsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListAssetsRequest>, I>>(object: I): ListAssetsRequest {
    const message = createBaseListAssetsRequest();
    message.albumIds = object.albumIds?.map((e) => e) || [];
    message.isFilterVideo = object.isFilterVideo ?? undefined;
    message.sortBy = object.sortBy ?? undefined;
    message.startDate = object.startDate ?? undefined;
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
    size: 0,
    mime: "",
    fileName: undefined,
    creationDate: undefined,
    latitude: undefined,
    longitude: undefined,
    takenDate: undefined,
    modifiedDate: undefined,
  };
}

export const PhotoMeta = {
  encode(message: PhotoMeta, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.size !== 0) {
      writer.uint32(56).int64(message.size);
    }
    if (message.mime !== "") {
      writer.uint32(66).string(message.mime);
    }
    if (message.fileName !== undefined) {
      writer.uint32(82).string(message.fileName);
    }
    if (message.creationDate !== undefined) {
      Timestamp.encode(toTimestamp(message.creationDate), writer.uint32(90).fork()).ldelim();
    }
    if (message.latitude !== undefined) {
      writer.uint32(101).float(message.latitude);
    }
    if (message.longitude !== undefined) {
      writer.uint32(109).float(message.longitude);
    }
    if (message.takenDate !== undefined) {
      Timestamp.encode(toTimestamp(message.takenDate), writer.uint32(114).fork()).ldelim();
    }
    if (message.modifiedDate !== undefined) {
      Timestamp.encode(toTimestamp(message.modifiedDate), writer.uint32(122).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PhotoMeta {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePhotoMeta();
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

          message.photoUrl = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.thumbnailUrl = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.width = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.height = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.albumIds.push(reader.string());
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.size = longToNumber(reader.int64() as Long);
          continue;
        case 8:
          if (tag !== 66) {
            break;
          }

          message.mime = reader.string();
          continue;
        case 10:
          if (tag !== 82) {
            break;
          }

          message.fileName = reader.string();
          continue;
        case 11:
          if (tag !== 90) {
            break;
          }

          message.creationDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 12:
          if (tag !== 101) {
            break;
          }

          message.latitude = reader.float();
          continue;
        case 13:
          if (tag !== 109) {
            break;
          }

          message.longitude = reader.float();
          continue;
        case 14:
          if (tag !== 114) {
            break;
          }

          message.takenDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 15:
          if (tag !== 122) {
            break;
          }

          message.modifiedDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PhotoMeta {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      photoUrl: isSet(object.photoUrl) ? String(object.photoUrl) : "",
      thumbnailUrl: isSet(object.thumbnailUrl) ? String(object.thumbnailUrl) : "",
      width: isSet(object.width) ? String(object.width) : "",
      height: isSet(object.height) ? String(object.height) : "",
      albumIds: Array.isArray(object?.albumIds) ? object.albumIds.map((e: any) => String(e)) : [],
      size: isSet(object.size) ? Number(object.size) : 0,
      mime: isSet(object.mime) ? String(object.mime) : "",
      fileName: isSet(object.fileName) ? String(object.fileName) : undefined,
      creationDate: isSet(object.creationDate) ? fromJsonTimestamp(object.creationDate) : undefined,
      latitude: isSet(object.latitude) ? Number(object.latitude) : undefined,
      longitude: isSet(object.longitude) ? Number(object.longitude) : undefined,
      takenDate: isSet(object.takenDate) ? fromJsonTimestamp(object.takenDate) : undefined,
      modifiedDate: isSet(object.modifiedDate) ? fromJsonTimestamp(object.modifiedDate) : undefined,
    };
  },

  toJSON(message: PhotoMeta): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.photoUrl !== "") {
      obj.photoUrl = message.photoUrl;
    }
    if (message.thumbnailUrl !== "") {
      obj.thumbnailUrl = message.thumbnailUrl;
    }
    if (message.width !== "") {
      obj.width = message.width;
    }
    if (message.height !== "") {
      obj.height = message.height;
    }
    if (message.albumIds?.length) {
      obj.albumIds = message.albumIds;
    }
    if (message.size !== 0) {
      obj.size = Math.round(message.size);
    }
    if (message.mime !== "") {
      obj.mime = message.mime;
    }
    if (message.fileName !== undefined) {
      obj.fileName = message.fileName;
    }
    if (message.creationDate !== undefined) {
      obj.creationDate = message.creationDate.toISOString();
    }
    if (message.latitude !== undefined) {
      obj.latitude = message.latitude;
    }
    if (message.longitude !== undefined) {
      obj.longitude = message.longitude;
    }
    if (message.takenDate !== undefined) {
      obj.takenDate = message.takenDate.toISOString();
    }
    if (message.modifiedDate !== undefined) {
      obj.modifiedDate = message.modifiedDate.toISOString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PhotoMeta>, I>>(base?: I): PhotoMeta {
    return PhotoMeta.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PhotoMeta>, I>>(object: I): PhotoMeta {
    const message = createBasePhotoMeta();
    message.id = object.id ?? "";
    message.photoUrl = object.photoUrl ?? "";
    message.thumbnailUrl = object.thumbnailUrl ?? "";
    message.width = object.width ?? "";
    message.height = object.height ?? "";
    message.albumIds = object.albumIds?.map((e) => e) || [];
    message.size = object.size ?? 0;
    message.mime = object.mime ?? "";
    message.fileName = object.fileName ?? undefined;
    message.creationDate = object.creationDate ?? undefined;
    message.latitude = object.latitude ?? undefined;
    message.longitude = object.longitude ?? undefined;
    message.takenDate = object.takenDate ?? undefined;
    message.modifiedDate = object.modifiedDate ?? undefined;
    return message;
  },
};

function createBaseAlbum(): Album {
  return { id: "", title: "", imageCount: 0, videoCount: 0, coverImageUrl: "", coverImageId: "" };
}

export const Album = {
  encode(message: Album, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.imageCount !== 0) {
      writer.uint32(24).int32(message.imageCount);
    }
    if (message.videoCount !== 0) {
      writer.uint32(32).int32(message.videoCount);
    }
    if (message.coverImageUrl !== "") {
      writer.uint32(42).string(message.coverImageUrl);
    }
    if (message.coverImageId !== "") {
      writer.uint32(50).string(message.coverImageId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Album {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAlbum();
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

          message.title = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.imageCount = reader.int32();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.videoCount = reader.int32();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.coverImageUrl = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.coverImageId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Album {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      title: isSet(object.title) ? String(object.title) : "",
      imageCount: isSet(object.imageCount) ? Number(object.imageCount) : 0,
      videoCount: isSet(object.videoCount) ? Number(object.videoCount) : 0,
      coverImageUrl: isSet(object.coverImageUrl) ? String(object.coverImageUrl) : "",
      coverImageId: isSet(object.coverImageId) ? String(object.coverImageId) : "",
    };
  },

  toJSON(message: Album): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.imageCount !== 0) {
      obj.imageCount = Math.round(message.imageCount);
    }
    if (message.videoCount !== 0) {
      obj.videoCount = Math.round(message.videoCount);
    }
    if (message.coverImageUrl !== "") {
      obj.coverImageUrl = message.coverImageUrl;
    }
    if (message.coverImageId !== "") {
      obj.coverImageId = message.coverImageId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Album>, I>>(base?: I): Album {
    return Album.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Album>, I>>(object: I): Album {
    const message = createBaseAlbum();
    message.id = object.id ?? "";
    message.title = object.title ?? "";
    message.imageCount = object.imageCount ?? 0;
    message.videoCount = object.videoCount ?? 0;
    message.coverImageUrl = object.coverImageUrl ?? "";
    message.coverImageId = object.coverImageId ?? "";
    return message;
  },
};

function createBaseListAlbumsRequest(): ListAlbumsRequest {
  return { thumbnailWidth: undefined, thumbnailHeight: undefined, thumbnailChooseWay: undefined };
}

export const ListAlbumsRequest = {
  encode(message: ListAlbumsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.thumbnailWidth !== undefined) {
      writer.uint32(8).int32(message.thumbnailWidth);
    }
    if (message.thumbnailHeight !== undefined) {
      writer.uint32(16).int32(message.thumbnailHeight);
    }
    if (message.thumbnailChooseWay !== undefined) {
      writer.uint32(24).int32(message.thumbnailChooseWay);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAlbumsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAlbumsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.thumbnailWidth = reader.int32();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.thumbnailHeight = reader.int32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.thumbnailChooseWay = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListAlbumsRequest {
    return {
      thumbnailWidth: isSet(object.thumbnailWidth) ? Number(object.thumbnailWidth) : undefined,
      thumbnailHeight: isSet(object.thumbnailHeight) ? Number(object.thumbnailHeight) : undefined,
      thumbnailChooseWay: isSet(object.thumbnailChooseWay) ? Number(object.thumbnailChooseWay) : undefined,
    };
  },

  toJSON(message: ListAlbumsRequest): unknown {
    const obj: any = {};
    if (message.thumbnailWidth !== undefined) {
      obj.thumbnailWidth = Math.round(message.thumbnailWidth);
    }
    if (message.thumbnailHeight !== undefined) {
      obj.thumbnailHeight = Math.round(message.thumbnailHeight);
    }
    if (message.thumbnailChooseWay !== undefined) {
      obj.thumbnailChooseWay = Math.round(message.thumbnailChooseWay);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListAlbumsRequest>, I>>(base?: I): ListAlbumsRequest {
    return ListAlbumsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListAlbumsRequest>, I>>(object: I): ListAlbumsRequest {
    const message = createBaseListAlbumsRequest();
    message.thumbnailWidth = object.thumbnailWidth ?? undefined;
    message.thumbnailHeight = object.thumbnailHeight ?? undefined;
    message.thumbnailChooseWay = object.thumbnailChooseWay ?? undefined;
    return message;
  },
};

function createBaseListAlbumsReply(): ListAlbumsReply {
  return { albums: [] };
}

export const ListAlbumsReply = {
  encode(message: ListAlbumsReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.albums) {
      Album.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAlbumsReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAlbumsReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.albums.push(Album.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListAlbumsReply {
    return { albums: Array.isArray(object?.albums) ? object.albums.map((e: any) => Album.fromJSON(e)) : [] };
  },

  toJSON(message: ListAlbumsReply): unknown {
    const obj: any = {};
    if (message.albums?.length) {
      obj.albums = message.albums.map((e) => Album.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListAlbumsReply>, I>>(base?: I): ListAlbumsReply {
    return ListAlbumsReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListAlbumsReply>, I>>(object: I): ListAlbumsReply {
    const message = createBaseListAlbumsReply();
    message.albums = object.albums?.map((e) => Album.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListAssetStatsRequest(): ListAssetStatsRequest {
  return { isFilterVideo: false, sortBy: 0, startDate: undefined };
}

export const ListAssetStatsRequest = {
  encode(message: ListAssetStatsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.isFilterVideo === true) {
      writer.uint32(8).bool(message.isFilterVideo);
    }
    if (message.sortBy !== 0) {
      writer.uint32(16).int32(message.sortBy);
    }
    if (message.startDate !== undefined) {
      Timestamp.encode(toTimestamp(message.startDate), writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAssetStatsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAssetStatsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.isFilterVideo = reader.bool();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.sortBy = reader.int32() as any;
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.startDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListAssetStatsRequest {
    return {
      isFilterVideo: isSet(object.isFilterVideo) ? Boolean(object.isFilterVideo) : false,
      sortBy: isSet(object.sortBy) ? listAssetsSortTypeFromJSON(object.sortBy) : 0,
      startDate: isSet(object.startDate) ? fromJsonTimestamp(object.startDate) : undefined,
    };
  },

  toJSON(message: ListAssetStatsRequest): unknown {
    const obj: any = {};
    if (message.isFilterVideo === true) {
      obj.isFilterVideo = message.isFilterVideo;
    }
    if (message.sortBy !== 0) {
      obj.sortBy = listAssetsSortTypeToJSON(message.sortBy);
    }
    if (message.startDate !== undefined) {
      obj.startDate = message.startDate.toISOString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListAssetStatsRequest>, I>>(base?: I): ListAssetStatsRequest {
    return ListAssetStatsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListAssetStatsRequest>, I>>(object: I): ListAssetStatsRequest {
    const message = createBaseListAssetStatsRequest();
    message.isFilterVideo = object.isFilterVideo ?? false;
    message.sortBy = object.sortBy ?? 0;
    message.startDate = object.startDate ?? undefined;
    return message;
  },
};

function createBaseListAssetStatsReply(): ListAssetStatsReply {
  return { id: "", date: undefined, mime: undefined };
}

export const ListAssetStatsReply = {
  encode(message: ListAssetStatsReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.date !== undefined) {
      Timestamp.encode(toTimestamp(message.date), writer.uint32(18).fork()).ldelim();
    }
    if (message.mime !== undefined) {
      writer.uint32(26).string(message.mime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAssetStatsReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAssetStatsReply();
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

          message.date = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          continue;
        case 3:
          if (tag !== 26) {
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

  fromJSON(object: any): ListAssetStatsReply {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      date: isSet(object.date) ? fromJsonTimestamp(object.date) : undefined,
      mime: isSet(object.mime) ? String(object.mime) : undefined,
    };
  },

  toJSON(message: ListAssetStatsReply): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.date !== undefined) {
      obj.date = message.date.toISOString();
    }
    if (message.mime !== undefined) {
      obj.mime = message.mime;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListAssetStatsReply>, I>>(base?: I): ListAssetStatsReply {
    return ListAssetStatsReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListAssetStatsReply>, I>>(object: I): ListAssetStatsReply {
    const message = createBaseListAssetStatsReply();
    message.id = object.id ?? "";
    message.date = object.date ?? undefined;
    message.mime = object.mime ?? undefined;
    return message;
  },
};

function createBaseListAssetsByIdsRequest(): ListAssetsByIdsRequest {
  return { ids: [], sortBy: 0 };
}

export const ListAssetsByIdsRequest = {
  encode(message: ListAssetsByIdsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.ids) {
      writer.uint32(10).string(v!);
    }
    if (message.sortBy !== 0) {
      writer.uint32(16).int32(message.sortBy);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListAssetsByIdsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListAssetsByIdsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.ids.push(reader.string());
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.sortBy = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListAssetsByIdsRequest {
    return {
      ids: Array.isArray(object?.ids) ? object.ids.map((e: any) => String(e)) : [],
      sortBy: isSet(object.sortBy) ? listAssetsSortTypeFromJSON(object.sortBy) : 0,
    };
  },

  toJSON(message: ListAssetsByIdsRequest): unknown {
    const obj: any = {};
    if (message.ids?.length) {
      obj.ids = message.ids;
    }
    if (message.sortBy !== 0) {
      obj.sortBy = listAssetsSortTypeToJSON(message.sortBy);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListAssetsByIdsRequest>, I>>(base?: I): ListAssetsByIdsRequest {
    return ListAssetsByIdsRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListAssetsByIdsRequest>, I>>(object: I): ListAssetsByIdsRequest {
    const message = createBaseListAssetsByIdsRequest();
    message.ids = object.ids?.map((e) => e) || [];
    message.sortBy = object.sortBy ?? 0;
    return message;
  },
};

function createBaseQueryAssetUrlPathReply(): QueryAssetUrlPathReply {
  return { assetUrl: "", thumbnailUrl: "" };
}

export const QueryAssetUrlPathReply = {
  encode(message: QueryAssetUrlPathReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.assetUrl !== "") {
      writer.uint32(10).string(message.assetUrl);
    }
    if (message.thumbnailUrl !== "") {
      writer.uint32(18).string(message.thumbnailUrl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAssetUrlPathReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAssetUrlPathReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.assetUrl = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.thumbnailUrl = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAssetUrlPathReply {
    return {
      assetUrl: isSet(object.assetUrl) ? String(object.assetUrl) : "",
      thumbnailUrl: isSet(object.thumbnailUrl) ? String(object.thumbnailUrl) : "",
    };
  },

  toJSON(message: QueryAssetUrlPathReply): unknown {
    const obj: any = {};
    if (message.assetUrl !== "") {
      obj.assetUrl = message.assetUrl;
    }
    if (message.thumbnailUrl !== "") {
      obj.thumbnailUrl = message.thumbnailUrl;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAssetUrlPathReply>, I>>(base?: I): QueryAssetUrlPathReply {
    return QueryAssetUrlPathReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<QueryAssetUrlPathReply>, I>>(object: I): QueryAssetUrlPathReply {
    const message = createBaseQueryAssetUrlPathReply();
    message.assetUrl = object.assetUrl ?? "";
    message.thumbnailUrl = object.thumbnailUrl ?? "";
    return message;
  },
};

export interface PhotoLibrary {
  MakeAlbum(
    request: DeepPartial<MakeAlbumRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Album>;
  /** 列举所有的系统相册 */
  ListAlbums(
    request: DeepPartial<ListAlbumsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListAlbumsReply>;
  /** 存储一张图片到某个相册中 */
  PutPhoto(
    request: DeepPartial<PutPhotoRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<PutPhotoReply>;
  DeletePhoto(
    request: DeepPartial<DeletePhotoRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<DeletePhotoReply>;
  /**
   * 枚举具体相册中的图片元信息
   *
   * @deprecated
   */
  ListPhotoMetas(
    request: DeepPartial<ListPhotoMetasRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<PhotoMeta>;
  /** 枚举相册中的资源（视频 & 图片）列表 */
  ListAssets(
    request: DeepPartial<ListAssetsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<PhotoMeta>;
  QueryPhoto(
    request: DeepPartial<QueryPhotoRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<PhotoMeta>;
  /** 查询指定条件的的图片id信息 */
  ListAssetStats(
    request: DeepPartial<ListAssetStatsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<ListAssetStatsReply>;
  /** 查询指定条件的的图片 */
  ListAssetsByIds(
    request: DeepPartial<ListAssetsByIdsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<PhotoMeta>;
  /** 查询媒体资产 Url 路径 */
  QueryAssetUrlPath(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<QueryAssetUrlPathReply>;
}

export class PhotoLibraryClientImpl implements PhotoLibrary {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.MakeAlbum = this.MakeAlbum.bind(this);
    this.ListAlbums = this.ListAlbums.bind(this);
    this.PutPhoto = this.PutPhoto.bind(this);
    this.DeletePhoto = this.DeletePhoto.bind(this);
    this.ListPhotoMetas = this.ListPhotoMetas.bind(this);
    this.ListAssets = this.ListAssets.bind(this);
    this.QueryPhoto = this.QueryPhoto.bind(this);
    this.ListAssetStats = this.ListAssetStats.bind(this);
    this.ListAssetsByIds = this.ListAssetsByIds.bind(this);
    this.QueryAssetUrlPath = this.QueryAssetUrlPath.bind(this);
  }

  MakeAlbum(
    request: DeepPartial<MakeAlbumRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Album> {
    return this.rpc.unary(PhotoLibraryMakeAlbumDesc, MakeAlbumRequest.fromPartial(request), metadata, abortSignal);
  }

  ListAlbums(
    request: DeepPartial<ListAlbumsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListAlbumsReply> {
    return this.rpc.unary(PhotoLibraryListAlbumsDesc, ListAlbumsRequest.fromPartial(request), metadata, abortSignal);
  }

  PutPhoto(
    request: DeepPartial<PutPhotoRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<PutPhotoReply> {
    return this.rpc.invoke(PhotoLibraryPutPhotoDesc, PutPhotoRequest.fromPartial(request), metadata, abortSignal);
  }

  DeletePhoto(
    request: DeepPartial<DeletePhotoRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<DeletePhotoReply> {
    return this.rpc.unary(PhotoLibraryDeletePhotoDesc, DeletePhotoRequest.fromPartial(request), metadata, abortSignal);
  }

  ListPhotoMetas(
    request: DeepPartial<ListPhotoMetasRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<PhotoMeta> {
    return this.rpc.invoke(
      PhotoLibraryListPhotoMetasDesc,
      ListPhotoMetasRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  ListAssets(
    request: DeepPartial<ListAssetsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<PhotoMeta> {
    return this.rpc.invoke(PhotoLibraryListAssetsDesc, ListAssetsRequest.fromPartial(request), metadata, abortSignal);
  }

  QueryPhoto(
    request: DeepPartial<QueryPhotoRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<PhotoMeta> {
    return this.rpc.unary(PhotoLibraryQueryPhotoDesc, QueryPhotoRequest.fromPartial(request), metadata, abortSignal);
  }

  ListAssetStats(
    request: DeepPartial<ListAssetStatsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<ListAssetStatsReply> {
    return this.rpc.invoke(
      PhotoLibraryListAssetStatsDesc,
      ListAssetStatsRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  ListAssetsByIds(
    request: DeepPartial<ListAssetsByIdsRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<PhotoMeta> {
    return this.rpc.invoke(
      PhotoLibraryListAssetsByIdsDesc,
      ListAssetsByIdsRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  QueryAssetUrlPath(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<QueryAssetUrlPathReply> {
    return this.rpc.unary(PhotoLibraryQueryAssetUrlPathDesc, Empty.fromPartial(request), metadata, abortSignal);
  }
}

export const PhotoLibraryDesc = { serviceName: "cloud.lazycat.apis.localdevice.PhotoLibrary" };

export const PhotoLibraryMakeAlbumDesc: UnaryMethodDefinitionish = {
  methodName: "MakeAlbum",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return MakeAlbumRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = Album.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
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
      const value = ListAlbumsReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const PhotoLibraryPutPhotoDesc: UnaryMethodDefinitionish = {
  methodName: "PutPhoto",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return PutPhotoRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = PutPhotoReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
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
      const value = DeletePhotoReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
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
      const value = PhotoMeta.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const PhotoLibraryListAssetsDesc: UnaryMethodDefinitionish = {
  methodName: "ListAssets",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return ListAssetsRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = PhotoMeta.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const PhotoLibraryQueryPhotoDesc: UnaryMethodDefinitionish = {
  methodName: "QueryPhoto",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return QueryPhotoRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = PhotoMeta.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const PhotoLibraryListAssetStatsDesc: UnaryMethodDefinitionish = {
  methodName: "ListAssetStats",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return ListAssetStatsRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ListAssetStatsReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const PhotoLibraryListAssetsByIdsDesc: UnaryMethodDefinitionish = {
  methodName: "ListAssetsByIds",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return ListAssetsByIdsRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = PhotoMeta.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const PhotoLibraryQueryAssetUrlPathDesc: UnaryMethodDefinitionish = {
  methodName: "QueryAssetUrlPath",
  service: PhotoLibraryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = QueryAssetUrlPathReply.decode(data);
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export class GrpcWebError extends tsProtoGlobalThis.Error {
  constructor(message: string, public code: grpc.Code, public metadata: grpc.Metadata) {
    super(message);
  }
}

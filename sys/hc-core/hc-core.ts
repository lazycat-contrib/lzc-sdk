/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { Empty } from "../../google/protobuf/empty";
import { BrowserHeaders } from "browser-headers";

export enum ListOptType {
  /** ALL - 列出所有 */
  ALL = 0,
  /** SYS - 仅列出系统应用 */
  SYS = 1,
  /** APP - 仅列出用户应用 */
  APP = 2,
  UNRECOGNIZED = -1,
}

export function listOptTypeFromJSON(object: any): ListOptType {
  switch (object) {
    case 0:
    case "ALL":
      return ListOptType.ALL;
    case 1:
    case "SYS":
      return ListOptType.SYS;
    case 2:
    case "APP":
      return ListOptType.APP;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ListOptType.UNRECOGNIZED;
  }
}

export function listOptTypeToJSON(object: ListOptType): string {
  switch (object) {
    case ListOptType.ALL:
      return "ALL";
    case ListOptType.SYS:
      return "SYS";
    case ListOptType.APP:
      return "APP";
    case ListOptType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface ListOptRequest {
  opt: ListOptType;
}

export interface AppId {
  AppId: string;
}

export interface InstanceInfo {
  AppId: string;
  UserId: string;
}

export interface DeleteRequest {
  AppId: string;
  DeleteAllData: boolean;
}

export interface AppIdList {
  appIdList: string[];
}

export interface QueryRequest {
  appIdList: string[];
  opt: ListOptType;
}

export interface AppMetadata {
  id: string;
  system: boolean;
  userApp: boolean;
  version: string;
}

export interface AppMetadataArray {
  metadata: AppMetadata[];
}

export interface StatusInfo {
  status: StatusInfo_StatusType;
  errMsg: string;
}

export enum StatusInfo_StatusType {
  STOPPED = 0,
  STOPPING = 1,
  BUILDING = 2,
  DOWNLOADING = 3,
  HOOKING = 4,
  STARTING = 5,
  RUNNING = 6,
  INSUFFICIENT = 7,
  PAUSED = 8,
  DEAD = 9,
  ERROR = 10,
  DISABLED = 11,
  UNRECOGNIZED = -1,
}

export function statusInfo_StatusTypeFromJSON(
  object: any
): StatusInfo_StatusType {
  switch (object) {
    case 0:
    case "STOPPED":
      return StatusInfo_StatusType.STOPPED;
    case 1:
    case "STOPPING":
      return StatusInfo_StatusType.STOPPING;
    case 2:
    case "BUILDING":
      return StatusInfo_StatusType.BUILDING;
    case 3:
    case "DOWNLOADING":
      return StatusInfo_StatusType.DOWNLOADING;
    case 4:
    case "HOOKING":
      return StatusInfo_StatusType.HOOKING;
    case 5:
    case "STARTING":
      return StatusInfo_StatusType.STARTING;
    case 6:
    case "RUNNING":
      return StatusInfo_StatusType.RUNNING;
    case 7:
    case "INSUFFICIENT":
      return StatusInfo_StatusType.INSUFFICIENT;
    case 8:
    case "PAUSED":
      return StatusInfo_StatusType.PAUSED;
    case 9:
    case "DEAD":
      return StatusInfo_StatusType.DEAD;
    case 10:
    case "ERROR":
      return StatusInfo_StatusType.ERROR;
    case 11:
    case "DISABLED":
      return StatusInfo_StatusType.DISABLED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return StatusInfo_StatusType.UNRECOGNIZED;
  }
}

export function statusInfo_StatusTypeToJSON(
  object: StatusInfo_StatusType
): string {
  switch (object) {
    case StatusInfo_StatusType.STOPPED:
      return "STOPPED";
    case StatusInfo_StatusType.STOPPING:
      return "STOPPING";
    case StatusInfo_StatusType.BUILDING:
      return "BUILDING";
    case StatusInfo_StatusType.DOWNLOADING:
      return "DOWNLOADING";
    case StatusInfo_StatusType.HOOKING:
      return "HOOKING";
    case StatusInfo_StatusType.STARTING:
      return "STARTING";
    case StatusInfo_StatusType.RUNNING:
      return "RUNNING";
    case StatusInfo_StatusType.INSUFFICIENT:
      return "INSUFFICIENT";
    case StatusInfo_StatusType.PAUSED:
      return "PAUSED";
    case StatusInfo_StatusType.DEAD:
      return "DEAD";
    case StatusInfo_StatusType.ERROR:
      return "ERROR";
    case StatusInfo_StatusType.DISABLED:
      return "DISABLED";
    case StatusInfo_StatusType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface InstancesStatus {
  instancesStatus: { [key: string]: StatusInfo };
}

export interface InstancesStatus_InstancesStatusEntry {
  key: string;
  value: StatusInfo | undefined;
}

export interface AppsStatus {
  appsStatus: { [key: string]: InstancesStatus };
}

export interface AppsStatus_AppsStatusEntry {
  key: string;
  value: InstancesStatus | undefined;
}

export interface RawData {
  data: Uint8Array;
}

function createBaseListOptRequest(): ListOptRequest {
  return { opt: 0 };
}

export const ListOptRequest = {
  encode(
    message: ListOptRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.opt !== 0) {
      writer.uint32(8).int32(message.opt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListOptRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListOptRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.opt = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListOptRequest {
    return {
      opt: isSet(object.opt) ? listOptTypeFromJSON(object.opt) : 0,
    };
  },

  toJSON(message: ListOptRequest): unknown {
    const obj: any = {};
    message.opt !== undefined && (obj.opt = listOptTypeToJSON(message.opt));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListOptRequest>, I>>(
    object: I
  ): ListOptRequest {
    const message = createBaseListOptRequest();
    message.opt = object.opt ?? 0;
    return message;
  },
};

function createBaseAppId(): AppId {
  return { AppId: "" };
}

export const AppId = {
  encode(message: AppId, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.AppId !== "") {
      writer.uint32(10).string(message.AppId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppId {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppId();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.AppId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppId {
    return {
      AppId: isSet(object.AppId) ? String(object.AppId) : "",
    };
  },

  toJSON(message: AppId): unknown {
    const obj: any = {};
    message.AppId !== undefined && (obj.AppId = message.AppId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AppId>, I>>(object: I): AppId {
    const message = createBaseAppId();
    message.AppId = object.AppId ?? "";
    return message;
  },
};

function createBaseInstanceInfo(): InstanceInfo {
  return { AppId: "", UserId: "" };
}

export const InstanceInfo = {
  encode(
    message: InstanceInfo,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.AppId !== "") {
      writer.uint32(10).string(message.AppId);
    }
    if (message.UserId !== "") {
      writer.uint32(18).string(message.UserId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InstanceInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInstanceInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.AppId = reader.string();
          break;
        case 2:
          message.UserId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InstanceInfo {
    return {
      AppId: isSet(object.AppId) ? String(object.AppId) : "",
      UserId: isSet(object.UserId) ? String(object.UserId) : "",
    };
  },

  toJSON(message: InstanceInfo): unknown {
    const obj: any = {};
    message.AppId !== undefined && (obj.AppId = message.AppId);
    message.UserId !== undefined && (obj.UserId = message.UserId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InstanceInfo>, I>>(
    object: I
  ): InstanceInfo {
    const message = createBaseInstanceInfo();
    message.AppId = object.AppId ?? "";
    message.UserId = object.UserId ?? "";
    return message;
  },
};

function createBaseDeleteRequest(): DeleteRequest {
  return { AppId: "", DeleteAllData: false };
}

export const DeleteRequest = {
  encode(
    message: DeleteRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.AppId !== "") {
      writer.uint32(10).string(message.AppId);
    }
    if (message.DeleteAllData === true) {
      writer.uint32(16).bool(message.DeleteAllData);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.AppId = reader.string();
          break;
        case 2:
          message.DeleteAllData = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteRequest {
    return {
      AppId: isSet(object.AppId) ? String(object.AppId) : "",
      DeleteAllData: isSet(object.DeleteAllData)
        ? Boolean(object.DeleteAllData)
        : false,
    };
  },

  toJSON(message: DeleteRequest): unknown {
    const obj: any = {};
    message.AppId !== undefined && (obj.AppId = message.AppId);
    message.DeleteAllData !== undefined &&
      (obj.DeleteAllData = message.DeleteAllData);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteRequest>, I>>(
    object: I
  ): DeleteRequest {
    const message = createBaseDeleteRequest();
    message.AppId = object.AppId ?? "";
    message.DeleteAllData = object.DeleteAllData ?? false;
    return message;
  },
};

function createBaseAppIdList(): AppIdList {
  return { appIdList: [] };
}

export const AppIdList = {
  encode(
    message: AppIdList,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.appIdList) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppIdList {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppIdList();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.appIdList.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppIdList {
    return {
      appIdList: Array.isArray(object?.appIdList)
        ? object.appIdList.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: AppIdList): unknown {
    const obj: any = {};
    if (message.appIdList) {
      obj.appIdList = message.appIdList.map((e) => e);
    } else {
      obj.appIdList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AppIdList>, I>>(
    object: I
  ): AppIdList {
    const message = createBaseAppIdList();
    message.appIdList = object.appIdList?.map((e) => e) || [];
    return message;
  },
};

function createBaseQueryRequest(): QueryRequest {
  return { appIdList: [], opt: 0 };
}

export const QueryRequest = {
  encode(
    message: QueryRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.appIdList) {
      writer.uint32(10).string(v!);
    }
    if (message.opt !== 0) {
      writer.uint32(16).int32(message.opt);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.appIdList.push(reader.string());
          break;
        case 2:
          message.opt = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryRequest {
    return {
      appIdList: Array.isArray(object?.appIdList)
        ? object.appIdList.map((e: any) => String(e))
        : [],
      opt: isSet(object.opt) ? listOptTypeFromJSON(object.opt) : 0,
    };
  },

  toJSON(message: QueryRequest): unknown {
    const obj: any = {};
    if (message.appIdList) {
      obj.appIdList = message.appIdList.map((e) => e);
    } else {
      obj.appIdList = [];
    }
    message.opt !== undefined && (obj.opt = listOptTypeToJSON(message.opt));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryRequest>, I>>(
    object: I
  ): QueryRequest {
    const message = createBaseQueryRequest();
    message.appIdList = object.appIdList?.map((e) => e) || [];
    message.opt = object.opt ?? 0;
    return message;
  },
};

function createBaseAppMetadata(): AppMetadata {
  return { id: "", system: false, userApp: false, version: "" };
}

export const AppMetadata = {
  encode(
    message: AppMetadata,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.system === true) {
      writer.uint32(16).bool(message.system);
    }
    if (message.userApp === true) {
      writer.uint32(24).bool(message.userApp);
    }
    if (message.version !== "") {
      writer.uint32(34).string(message.version);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppMetadata {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppMetadata();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.system = reader.bool();
          break;
        case 3:
          message.userApp = reader.bool();
          break;
        case 4:
          message.version = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppMetadata {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      system: isSet(object.system) ? Boolean(object.system) : false,
      userApp: isSet(object.userApp) ? Boolean(object.userApp) : false,
      version: isSet(object.version) ? String(object.version) : "",
    };
  },

  toJSON(message: AppMetadata): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.system !== undefined && (obj.system = message.system);
    message.userApp !== undefined && (obj.userApp = message.userApp);
    message.version !== undefined && (obj.version = message.version);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AppMetadata>, I>>(
    object: I
  ): AppMetadata {
    const message = createBaseAppMetadata();
    message.id = object.id ?? "";
    message.system = object.system ?? false;
    message.userApp = object.userApp ?? false;
    message.version = object.version ?? "";
    return message;
  },
};

function createBaseAppMetadataArray(): AppMetadataArray {
  return { metadata: [] };
}

export const AppMetadataArray = {
  encode(
    message: AppMetadataArray,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.metadata) {
      AppMetadata.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppMetadataArray {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppMetadataArray();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.metadata.push(AppMetadata.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppMetadataArray {
    return {
      metadata: Array.isArray(object?.metadata)
        ? object.metadata.map((e: any) => AppMetadata.fromJSON(e))
        : [],
    };
  },

  toJSON(message: AppMetadataArray): unknown {
    const obj: any = {};
    if (message.metadata) {
      obj.metadata = message.metadata.map((e) =>
        e ? AppMetadata.toJSON(e) : undefined
      );
    } else {
      obj.metadata = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AppMetadataArray>, I>>(
    object: I
  ): AppMetadataArray {
    const message = createBaseAppMetadataArray();
    message.metadata =
      object.metadata?.map((e) => AppMetadata.fromPartial(e)) || [];
    return message;
  },
};

function createBaseStatusInfo(): StatusInfo {
  return { status: 0, errMsg: "" };
}

export const StatusInfo = {
  encode(
    message: StatusInfo,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    if (message.errMsg !== "") {
      writer.uint32(18).string(message.errMsg);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): StatusInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseStatusInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32() as any;
          break;
        case 2:
          message.errMsg = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StatusInfo {
    return {
      status: isSet(object.status)
        ? statusInfo_StatusTypeFromJSON(object.status)
        : 0,
      errMsg: isSet(object.errMsg) ? String(object.errMsg) : "",
    };
  },

  toJSON(message: StatusInfo): unknown {
    const obj: any = {};
    message.status !== undefined &&
      (obj.status = statusInfo_StatusTypeToJSON(message.status));
    message.errMsg !== undefined && (obj.errMsg = message.errMsg);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<StatusInfo>, I>>(
    object: I
  ): StatusInfo {
    const message = createBaseStatusInfo();
    message.status = object.status ?? 0;
    message.errMsg = object.errMsg ?? "";
    return message;
  },
};

function createBaseInstancesStatus(): InstancesStatus {
  return { instancesStatus: {} };
}

export const InstancesStatus = {
  encode(
    message: InstancesStatus,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    Object.entries(message.instancesStatus).forEach(([key, value]) => {
      InstancesStatus_InstancesStatusEntry.encode(
        { key: key as any, value },
        writer.uint32(10).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InstancesStatus {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInstancesStatus();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = InstancesStatus_InstancesStatusEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry1.value !== undefined) {
            message.instancesStatus[entry1.key] = entry1.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InstancesStatus {
    return {
      instancesStatus: isObject(object.instancesStatus)
        ? Object.entries(object.instancesStatus).reduce<{
            [key: string]: StatusInfo;
          }>((acc, [key, value]) => {
            acc[key] = StatusInfo.fromJSON(value);
            return acc;
          }, {})
        : {},
    };
  },

  toJSON(message: InstancesStatus): unknown {
    const obj: any = {};
    obj.instancesStatus = {};
    if (message.instancesStatus) {
      Object.entries(message.instancesStatus).forEach(([k, v]) => {
        obj.instancesStatus[k] = StatusInfo.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InstancesStatus>, I>>(
    object: I
  ): InstancesStatus {
    const message = createBaseInstancesStatus();
    message.instancesStatus = Object.entries(
      object.instancesStatus ?? {}
    ).reduce<{ [key: string]: StatusInfo }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = StatusInfo.fromPartial(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseInstancesStatus_InstancesStatusEntry(): InstancesStatus_InstancesStatusEntry {
  return { key: "", value: undefined };
}

export const InstancesStatus_InstancesStatusEntry = {
  encode(
    message: InstancesStatus_InstancesStatusEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      StatusInfo.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): InstancesStatus_InstancesStatusEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInstancesStatus_InstancesStatusEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = StatusInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InstancesStatus_InstancesStatusEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value)
        ? StatusInfo.fromJSON(object.value)
        : undefined,
    };
  },

  toJSON(message: InstancesStatus_InstancesStatusEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value
        ? StatusInfo.toJSON(message.value)
        : undefined);
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<InstancesStatus_InstancesStatusEntry>, I>
  >(object: I): InstancesStatus_InstancesStatusEntry {
    const message = createBaseInstancesStatus_InstancesStatusEntry();
    message.key = object.key ?? "";
    message.value =
      object.value !== undefined && object.value !== null
        ? StatusInfo.fromPartial(object.value)
        : undefined;
    return message;
  },
};

function createBaseAppsStatus(): AppsStatus {
  return { appsStatus: {} };
}

export const AppsStatus = {
  encode(
    message: AppsStatus,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    Object.entries(message.appsStatus).forEach(([key, value]) => {
      AppsStatus_AppsStatusEntry.encode(
        { key: key as any, value },
        writer.uint32(10).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppsStatus {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppsStatus();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = AppsStatus_AppsStatusEntry.decode(
            reader,
            reader.uint32()
          );
          if (entry1.value !== undefined) {
            message.appsStatus[entry1.key] = entry1.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppsStatus {
    return {
      appsStatus: isObject(object.appsStatus)
        ? Object.entries(object.appsStatus).reduce<{
            [key: string]: InstancesStatus;
          }>((acc, [key, value]) => {
            acc[key] = InstancesStatus.fromJSON(value);
            return acc;
          }, {})
        : {},
    };
  },

  toJSON(message: AppsStatus): unknown {
    const obj: any = {};
    obj.appsStatus = {};
    if (message.appsStatus) {
      Object.entries(message.appsStatus).forEach(([k, v]) => {
        obj.appsStatus[k] = InstancesStatus.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AppsStatus>, I>>(
    object: I
  ): AppsStatus {
    const message = createBaseAppsStatus();
    message.appsStatus = Object.entries(object.appsStatus ?? {}).reduce<{
      [key: string]: InstancesStatus;
    }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = InstancesStatus.fromPartial(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseAppsStatus_AppsStatusEntry(): AppsStatus_AppsStatusEntry {
  return { key: "", value: undefined };
}

export const AppsStatus_AppsStatusEntry = {
  encode(
    message: AppsStatus_AppsStatusEntry,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      InstancesStatus.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): AppsStatus_AppsStatusEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppsStatus_AppsStatusEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = InstancesStatus.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppsStatus_AppsStatusEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value)
        ? InstancesStatus.fromJSON(object.value)
        : undefined,
    };
  },

  toJSON(message: AppsStatus_AppsStatusEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value
        ? InstancesStatus.toJSON(message.value)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AppsStatus_AppsStatusEntry>, I>>(
    object: I
  ): AppsStatus_AppsStatusEntry {
    const message = createBaseAppsStatus_AppsStatusEntry();
    message.key = object.key ?? "";
    message.value =
      object.value !== undefined && object.value !== null
        ? InstancesStatus.fromPartial(object.value)
        : undefined;
    return message;
  },
};

function createBaseRawData(): RawData {
  return { data: new Uint8Array() };
}

export const RawData = {
  encode(
    message: RawData,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.data.length !== 0) {
      writer.uint32(10).bytes(message.data);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RawData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRawData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.data = reader.bytes();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RawData {
    return {
      data: isSet(object.data)
        ? bytesFromBase64(object.data)
        : new Uint8Array(),
    };
  },

  toJSON(message: RawData): unknown {
    const obj: any = {};
    message.data !== undefined &&
      (obj.data = base64FromBytes(
        message.data !== undefined ? message.data : new Uint8Array()
      ));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RawData>, I>>(object: I): RawData {
    const message = createBaseRawData();
    message.data = object.data ?? new Uint8Array();
    return message;
  },
};

/** 应用管理 */
export interface CoreSystem {
  /** List 列出所有已安装的应用 */
  List(
    request: DeepPartial<ListOptRequest>,
    metadata?: grpc.Metadata
  ): Promise<AppIdList>;
  /** Apply 安装或更新一个 app */
  Apply(
    request: DeepPartial<RawData>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
  /** Remove 移除一个 app */
  Remove(
    request: DeepPartial<DeleteRequest>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
  /** Disable 禁用一个 app */
  Disable(
    request: DeepPartial<QueryRequest>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
  /** Enable 启用一个 app */
  Enable(
    request: DeepPartial<QueryRequest>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
  /** Query 查询 app 的元信息 */
  Query(
    request: DeepPartial<QueryRequest>,
    metadata?: grpc.Metadata
  ): Promise<AppMetadataArray>;
  /** 查询所有 app 的所有实例状态 */
  Status(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<AppsStatus>;
  /** 查询 app 的所有实例状态 */
  AppStatus(
    request: DeepPartial<AppId>,
    metadata?: grpc.Metadata
  ): Promise<InstancesStatus>;
  /** 查询实例状态 */
  InstanceStatus(
    request: DeepPartial<InstanceInfo>,
    metadata?: grpc.Metadata
  ): Promise<StatusInfo>;
  /** 添加一个实例 */
  InstanceAdd(
    request: DeepPartial<InstanceInfo>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
  /** 删除一个实例 */
  InstanceRemove(
    request: DeepPartial<InstanceInfo>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
  /** 暂停一个实例 */
  InstancePause(
    request: DeepPartial<InstanceInfo>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
  /** 继续一个实例 */
  InstanceResume(
    request: DeepPartial<InstanceInfo>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
}

export class CoreSystemClientImpl implements CoreSystem {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.List = this.List.bind(this);
    this.Apply = this.Apply.bind(this);
    this.Remove = this.Remove.bind(this);
    this.Disable = this.Disable.bind(this);
    this.Enable = this.Enable.bind(this);
    this.Query = this.Query.bind(this);
    this.Status = this.Status.bind(this);
    this.AppStatus = this.AppStatus.bind(this);
    this.InstanceStatus = this.InstanceStatus.bind(this);
    this.InstanceAdd = this.InstanceAdd.bind(this);
    this.InstanceRemove = this.InstanceRemove.bind(this);
    this.InstancePause = this.InstancePause.bind(this);
    this.InstanceResume = this.InstanceResume.bind(this);
  }

  List(
    request: DeepPartial<ListOptRequest>,
    metadata?: grpc.Metadata
  ): Promise<AppIdList> {
    return this.rpc.unary(
      CoreSystemListDesc,
      ListOptRequest.fromPartial(request),
      metadata
    );
  }

  Apply(
    request: DeepPartial<RawData>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      CoreSystemApplyDesc,
      RawData.fromPartial(request),
      metadata
    );
  }

  Remove(
    request: DeepPartial<DeleteRequest>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      CoreSystemRemoveDesc,
      DeleteRequest.fromPartial(request),
      metadata
    );
  }

  Disable(
    request: DeepPartial<QueryRequest>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      CoreSystemDisableDesc,
      QueryRequest.fromPartial(request),
      metadata
    );
  }

  Enable(
    request: DeepPartial<QueryRequest>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      CoreSystemEnableDesc,
      QueryRequest.fromPartial(request),
      metadata
    );
  }

  Query(
    request: DeepPartial<QueryRequest>,
    metadata?: grpc.Metadata
  ): Promise<AppMetadataArray> {
    return this.rpc.unary(
      CoreSystemQueryDesc,
      QueryRequest.fromPartial(request),
      metadata
    );
  }

  Status(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<AppsStatus> {
    return this.rpc.unary(
      CoreSystemStatusDesc,
      Empty.fromPartial(request),
      metadata
    );
  }

  AppStatus(
    request: DeepPartial<AppId>,
    metadata?: grpc.Metadata
  ): Promise<InstancesStatus> {
    return this.rpc.unary(
      CoreSystemAppStatusDesc,
      AppId.fromPartial(request),
      metadata
    );
  }

  InstanceStatus(
    request: DeepPartial<InstanceInfo>,
    metadata?: grpc.Metadata
  ): Promise<StatusInfo> {
    return this.rpc.unary(
      CoreSystemInstanceStatusDesc,
      InstanceInfo.fromPartial(request),
      metadata
    );
  }

  InstanceAdd(
    request: DeepPartial<InstanceInfo>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      CoreSystemInstanceAddDesc,
      InstanceInfo.fromPartial(request),
      metadata
    );
  }

  InstanceRemove(
    request: DeepPartial<InstanceInfo>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      CoreSystemInstanceRemoveDesc,
      InstanceInfo.fromPartial(request),
      metadata
    );
  }

  InstancePause(
    request: DeepPartial<InstanceInfo>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      CoreSystemInstancePauseDesc,
      InstanceInfo.fromPartial(request),
      metadata
    );
  }

  InstanceResume(
    request: DeepPartial<InstanceInfo>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      CoreSystemInstanceResumeDesc,
      InstanceInfo.fromPartial(request),
      metadata
    );
  }
}

export const CoreSystemDesc = {
  serviceName: "cloud.lazycat.sys.CoreSystem",
};

export const CoreSystemListDesc: UnaryMethodDefinitionish = {
  methodName: "List",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ListOptRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...AppIdList.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const CoreSystemApplyDesc: UnaryMethodDefinitionish = {
  methodName: "Apply",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return RawData.encode(this).finish();
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

export const CoreSystemRemoveDesc: UnaryMethodDefinitionish = {
  methodName: "Remove",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return DeleteRequest.encode(this).finish();
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

export const CoreSystemDisableDesc: UnaryMethodDefinitionish = {
  methodName: "Disable",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return QueryRequest.encode(this).finish();
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

export const CoreSystemEnableDesc: UnaryMethodDefinitionish = {
  methodName: "Enable",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return QueryRequest.encode(this).finish();
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

export const CoreSystemQueryDesc: UnaryMethodDefinitionish = {
  methodName: "Query",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return QueryRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...AppMetadataArray.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const CoreSystemStatusDesc: UnaryMethodDefinitionish = {
  methodName: "Status",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...AppsStatus.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const CoreSystemAppStatusDesc: UnaryMethodDefinitionish = {
  methodName: "AppStatus",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return AppId.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...InstancesStatus.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const CoreSystemInstanceStatusDesc: UnaryMethodDefinitionish = {
  methodName: "InstanceStatus",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return InstanceInfo.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...StatusInfo.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const CoreSystemInstanceAddDesc: UnaryMethodDefinitionish = {
  methodName: "InstanceAdd",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return InstanceInfo.encode(this).finish();
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

export const CoreSystemInstanceRemoveDesc: UnaryMethodDefinitionish = {
  methodName: "InstanceRemove",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return InstanceInfo.encode(this).finish();
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

export const CoreSystemInstancePauseDesc: UnaryMethodDefinitionish = {
  methodName: "InstancePause",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return InstanceInfo.encode(this).finish();
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

export const CoreSystemInstanceResumeDesc: UnaryMethodDefinitionish = {
  methodName: "InstanceResume",
  service: CoreSystemDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return InstanceInfo.encode(this).finish();
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

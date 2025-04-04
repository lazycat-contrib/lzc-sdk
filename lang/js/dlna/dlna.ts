/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { share } from "rxjs/operators";
import { Empty } from "../google/protobuf/empty";

export interface GetPositionInfoRequest {
  playerUuid: string;
}

export interface GetPositionInfoResponse {
  track: string;
  trackDuration: string;
  trackMetadata: string;
  trackUri: string;
  relTime: string;
  absTime: string;
  relCount: number;
  absCount: number;
}

export interface DoActionRequest {
  playerUuid: string;
  action: DoActionRequest_Action;
  mediaFile?: string | undefined;
  mediaSubtitle?: string | undefined;
  seekTarget?: string | undefined;
}

export enum DoActionRequest_Action {
  Unknown = 0,
  /**
   * Play - 需要在mediaFile中设置对应的媒体路径，
   * 支持http://类型或/开头的本地资源路径(相当于当前用户目录)
   */
  Play = 1,
  Pause = 2,
  /** Continue - 若当前是paused状态，则恢复播放，否则不做任何处理，返回成功。 */
  Continue = 3,
  /** Stop - 停止对此instance上的任何操作，后续操作需要重新调用play开始 */
  Stop = 4,
  Seek = 5,
  UNRECOGNIZED = -1,
}

export function doActionRequest_ActionFromJSON(object: any): DoActionRequest_Action {
  switch (object) {
    case 0:
    case "Unknown":
      return DoActionRequest_Action.Unknown;
    case 1:
    case "Play":
      return DoActionRequest_Action.Play;
    case 2:
    case "Pause":
      return DoActionRequest_Action.Pause;
    case 3:
    case "Continue":
      return DoActionRequest_Action.Continue;
    case 4:
    case "Stop":
      return DoActionRequest_Action.Stop;
    case 5:
    case "Seek":
      return DoActionRequest_Action.Seek;
    case -1:
    case "UNRECOGNIZED":
    default:
      return DoActionRequest_Action.UNRECOGNIZED;
  }
}

export function doActionRequest_ActionToJSON(object: DoActionRequest_Action): string {
  switch (object) {
    case DoActionRequest_Action.Unknown:
      return "Unknown";
    case DoActionRequest_Action.Play:
      return "Play";
    case DoActionRequest_Action.Pause:
      return "Pause";
    case DoActionRequest_Action.Continue:
      return "Continue";
    case DoActionRequest_Action.Stop:
      return "Stop";
    case DoActionRequest_Action.Seek:
      return "Seek";
    case DoActionRequest_Action.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface RMPStatus {
  status: RMPStatus_Status;
}

export enum RMPStatus_Status {
  Unknown = 0,
  Playing = 1,
  Paused = 2,
  Stopped = 3,
  UNRECOGNIZED = -1,
}

export function rMPStatus_StatusFromJSON(object: any): RMPStatus_Status {
  switch (object) {
    case 0:
    case "Unknown":
      return RMPStatus_Status.Unknown;
    case 1:
    case "Playing":
      return RMPStatus_Status.Playing;
    case 2:
    case "Paused":
      return RMPStatus_Status.Paused;
    case 3:
    case "Stopped":
      return RMPStatus_Status.Stopped;
    case -1:
    case "UNRECOGNIZED":
    default:
      return RMPStatus_Status.UNRECOGNIZED;
  }
}

export function rMPStatus_StatusToJSON(object: RMPStatus_Status): string {
  switch (object) {
    case RMPStatus_Status.Unknown:
      return "Unknown";
    case RMPStatus_Status.Playing:
      return "Playing";
    case RMPStatus_Status.Paused:
      return "Paused";
    case RMPStatus_Status.Stopped:
      return "Stopped";
    case RMPStatus_Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface ScanRMPResponse {
  remoteMediaPlayers: RemoteMediaPlayer[];
}

export interface SubscribeRequest {
  playerUuid: string;
}

export interface RemoteMediaPlayer {
  uuid: string;
  name: string;
  iconData: string;
  /**
   * 所处局域网信息，若为空表示处于盒子的局域网，
   * 否则为具体客户端peer id的局域网
   */
  lanRegion: string;
}

function createBaseGetPositionInfoRequest(): GetPositionInfoRequest {
  return { playerUuid: "" };
}

export const GetPositionInfoRequest = {
  encode(message: GetPositionInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.playerUuid !== "") {
      writer.uint32(10).string(message.playerUuid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPositionInfoRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPositionInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.playerUuid = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPositionInfoRequest {
    return { playerUuid: isSet(object.playerUuid) ? String(object.playerUuid) : "" };
  },

  toJSON(message: GetPositionInfoRequest): unknown {
    const obj: any = {};
    if (message.playerUuid !== "") {
      obj.playerUuid = message.playerUuid;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPositionInfoRequest>, I>>(base?: I): GetPositionInfoRequest {
    return GetPositionInfoRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPositionInfoRequest>, I>>(object: I): GetPositionInfoRequest {
    const message = createBaseGetPositionInfoRequest();
    message.playerUuid = object.playerUuid ?? "";
    return message;
  },
};

function createBaseGetPositionInfoResponse(): GetPositionInfoResponse {
  return {
    track: "",
    trackDuration: "",
    trackMetadata: "",
    trackUri: "",
    relTime: "",
    absTime: "",
    relCount: 0,
    absCount: 0,
  };
}

export const GetPositionInfoResponse = {
  encode(message: GetPositionInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.track !== "") {
      writer.uint32(10).string(message.track);
    }
    if (message.trackDuration !== "") {
      writer.uint32(18).string(message.trackDuration);
    }
    if (message.trackMetadata !== "") {
      writer.uint32(26).string(message.trackMetadata);
    }
    if (message.trackUri !== "") {
      writer.uint32(34).string(message.trackUri);
    }
    if (message.relTime !== "") {
      writer.uint32(42).string(message.relTime);
    }
    if (message.absTime !== "") {
      writer.uint32(50).string(message.absTime);
    }
    if (message.relCount !== 0) {
      writer.uint32(56).int32(message.relCount);
    }
    if (message.absCount !== 0) {
      writer.uint32(64).int32(message.absCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetPositionInfoResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPositionInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.track = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.trackDuration = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.trackMetadata = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.trackUri = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.relTime = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.absTime = reader.string();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.relCount = reader.int32();
          continue;
        case 8:
          if (tag !== 64) {
            break;
          }

          message.absCount = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPositionInfoResponse {
    return {
      track: isSet(object.track) ? String(object.track) : "",
      trackDuration: isSet(object.trackDuration) ? String(object.trackDuration) : "",
      trackMetadata: isSet(object.trackMetadata) ? String(object.trackMetadata) : "",
      trackUri: isSet(object.trackUri) ? String(object.trackUri) : "",
      relTime: isSet(object.relTime) ? String(object.relTime) : "",
      absTime: isSet(object.absTime) ? String(object.absTime) : "",
      relCount: isSet(object.relCount) ? Number(object.relCount) : 0,
      absCount: isSet(object.absCount) ? Number(object.absCount) : 0,
    };
  },

  toJSON(message: GetPositionInfoResponse): unknown {
    const obj: any = {};
    if (message.track !== "") {
      obj.track = message.track;
    }
    if (message.trackDuration !== "") {
      obj.trackDuration = message.trackDuration;
    }
    if (message.trackMetadata !== "") {
      obj.trackMetadata = message.trackMetadata;
    }
    if (message.trackUri !== "") {
      obj.trackUri = message.trackUri;
    }
    if (message.relTime !== "") {
      obj.relTime = message.relTime;
    }
    if (message.absTime !== "") {
      obj.absTime = message.absTime;
    }
    if (message.relCount !== 0) {
      obj.relCount = Math.round(message.relCount);
    }
    if (message.absCount !== 0) {
      obj.absCount = Math.round(message.absCount);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPositionInfoResponse>, I>>(base?: I): GetPositionInfoResponse {
    return GetPositionInfoResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetPositionInfoResponse>, I>>(object: I): GetPositionInfoResponse {
    const message = createBaseGetPositionInfoResponse();
    message.track = object.track ?? "";
    message.trackDuration = object.trackDuration ?? "";
    message.trackMetadata = object.trackMetadata ?? "";
    message.trackUri = object.trackUri ?? "";
    message.relTime = object.relTime ?? "";
    message.absTime = object.absTime ?? "";
    message.relCount = object.relCount ?? 0;
    message.absCount = object.absCount ?? 0;
    return message;
  },
};

function createBaseDoActionRequest(): DoActionRequest {
  return { playerUuid: "", action: 0, mediaFile: undefined, mediaSubtitle: undefined, seekTarget: undefined };
}

export const DoActionRequest = {
  encode(message: DoActionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.playerUuid !== "") {
      writer.uint32(10).string(message.playerUuid);
    }
    if (message.action !== 0) {
      writer.uint32(16).int32(message.action);
    }
    if (message.mediaFile !== undefined) {
      writer.uint32(26).string(message.mediaFile);
    }
    if (message.mediaSubtitle !== undefined) {
      writer.uint32(34).string(message.mediaSubtitle);
    }
    if (message.seekTarget !== undefined) {
      writer.uint32(42).string(message.seekTarget);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DoActionRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDoActionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.playerUuid = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.action = reader.int32() as any;
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.mediaFile = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.mediaSubtitle = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.seekTarget = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DoActionRequest {
    return {
      playerUuid: isSet(object.playerUuid) ? String(object.playerUuid) : "",
      action: isSet(object.action) ? doActionRequest_ActionFromJSON(object.action) : 0,
      mediaFile: isSet(object.mediaFile) ? String(object.mediaFile) : undefined,
      mediaSubtitle: isSet(object.mediaSubtitle) ? String(object.mediaSubtitle) : undefined,
      seekTarget: isSet(object.seekTarget) ? String(object.seekTarget) : undefined,
    };
  },

  toJSON(message: DoActionRequest): unknown {
    const obj: any = {};
    if (message.playerUuid !== "") {
      obj.playerUuid = message.playerUuid;
    }
    if (message.action !== 0) {
      obj.action = doActionRequest_ActionToJSON(message.action);
    }
    if (message.mediaFile !== undefined) {
      obj.mediaFile = message.mediaFile;
    }
    if (message.mediaSubtitle !== undefined) {
      obj.mediaSubtitle = message.mediaSubtitle;
    }
    if (message.seekTarget !== undefined) {
      obj.seekTarget = message.seekTarget;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DoActionRequest>, I>>(base?: I): DoActionRequest {
    return DoActionRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<DoActionRequest>, I>>(object: I): DoActionRequest {
    const message = createBaseDoActionRequest();
    message.playerUuid = object.playerUuid ?? "";
    message.action = object.action ?? 0;
    message.mediaFile = object.mediaFile ?? undefined;
    message.mediaSubtitle = object.mediaSubtitle ?? undefined;
    message.seekTarget = object.seekTarget ?? undefined;
    return message;
  },
};

function createBaseRMPStatus(): RMPStatus {
  return { status: 0 };
}

export const RMPStatus = {
  encode(message: RMPStatus, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RMPStatus {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRMPStatus();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RMPStatus {
    return { status: isSet(object.status) ? rMPStatus_StatusFromJSON(object.status) : 0 };
  },

  toJSON(message: RMPStatus): unknown {
    const obj: any = {};
    if (message.status !== 0) {
      obj.status = rMPStatus_StatusToJSON(message.status);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RMPStatus>, I>>(base?: I): RMPStatus {
    return RMPStatus.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RMPStatus>, I>>(object: I): RMPStatus {
    const message = createBaseRMPStatus();
    message.status = object.status ?? 0;
    return message;
  },
};

function createBaseScanRMPResponse(): ScanRMPResponse {
  return { remoteMediaPlayers: [] };
}

export const ScanRMPResponse = {
  encode(message: ScanRMPResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.remoteMediaPlayers) {
      RemoteMediaPlayer.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ScanRMPResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseScanRMPResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.remoteMediaPlayers.push(RemoteMediaPlayer.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ScanRMPResponse {
    return {
      remoteMediaPlayers: Array.isArray(object?.remoteMediaPlayers)
        ? object.remoteMediaPlayers.map((e: any) => RemoteMediaPlayer.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ScanRMPResponse): unknown {
    const obj: any = {};
    if (message.remoteMediaPlayers?.length) {
      obj.remoteMediaPlayers = message.remoteMediaPlayers.map((e) => RemoteMediaPlayer.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ScanRMPResponse>, I>>(base?: I): ScanRMPResponse {
    return ScanRMPResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ScanRMPResponse>, I>>(object: I): ScanRMPResponse {
    const message = createBaseScanRMPResponse();
    message.remoteMediaPlayers = object.remoteMediaPlayers?.map((e) => RemoteMediaPlayer.fromPartial(e)) || [];
    return message;
  },
};

function createBaseSubscribeRequest(): SubscribeRequest {
  return { playerUuid: "" };
}

export const SubscribeRequest = {
  encode(message: SubscribeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.playerUuid !== "") {
      writer.uint32(10).string(message.playerUuid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SubscribeRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSubscribeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.playerUuid = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SubscribeRequest {
    return { playerUuid: isSet(object.playerUuid) ? String(object.playerUuid) : "" };
  },

  toJSON(message: SubscribeRequest): unknown {
    const obj: any = {};
    if (message.playerUuid !== "") {
      obj.playerUuid = message.playerUuid;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SubscribeRequest>, I>>(base?: I): SubscribeRequest {
    return SubscribeRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SubscribeRequest>, I>>(object: I): SubscribeRequest {
    const message = createBaseSubscribeRequest();
    message.playerUuid = object.playerUuid ?? "";
    return message;
  },
};

function createBaseRemoteMediaPlayer(): RemoteMediaPlayer {
  return { uuid: "", name: "", iconData: "", lanRegion: "" };
}

export const RemoteMediaPlayer = {
  encode(message: RemoteMediaPlayer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.iconData !== "") {
      writer.uint32(26).string(message.iconData);
    }
    if (message.lanRegion !== "") {
      writer.uint32(34).string(message.lanRegion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RemoteMediaPlayer {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRemoteMediaPlayer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.uuid = reader.string();
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

          message.iconData = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.lanRegion = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RemoteMediaPlayer {
    return {
      uuid: isSet(object.uuid) ? String(object.uuid) : "",
      name: isSet(object.name) ? String(object.name) : "",
      iconData: isSet(object.iconData) ? String(object.iconData) : "",
      lanRegion: isSet(object.lanRegion) ? String(object.lanRegion) : "",
    };
  },

  toJSON(message: RemoteMediaPlayer): unknown {
    const obj: any = {};
    if (message.uuid !== "") {
      obj.uuid = message.uuid;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.iconData !== "") {
      obj.iconData = message.iconData;
    }
    if (message.lanRegion !== "") {
      obj.lanRegion = message.lanRegion;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RemoteMediaPlayer>, I>>(base?: I): RemoteMediaPlayer {
    return RemoteMediaPlayer.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<RemoteMediaPlayer>, I>>(object: I): RemoteMediaPlayer {
    const message = createBaseRemoteMediaPlayer();
    message.uuid = object.uuid ?? "";
    message.name = object.name ?? "";
    message.iconData = object.iconData ?? "";
    message.lanRegion = object.lanRegion ?? "";
    return message;
  },
};

/** 目前支持搜索DLNA的Render设备，并投送媒体文件 */
export interface RemoteMediaPlayerService {
  ScanRMP(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<ScanRMPResponse>;
  /** 立刻返回选择的当前RMP状态，并在状态变动时重新发送 */
  Subscribe(
    request: DeepPartial<SubscribeRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<RMPStatus>;
  DoAction(request: DeepPartial<DoActionRequest>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<Empty>;
  GetPositionInfo(
    request: DeepPartial<GetPositionInfoRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<GetPositionInfoResponse>;
}

export class RemoteMediaPlayerServiceClientImpl implements RemoteMediaPlayerService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ScanRMP = this.ScanRMP.bind(this);
    this.Subscribe = this.Subscribe.bind(this);
    this.DoAction = this.DoAction.bind(this);
    this.GetPositionInfo = this.GetPositionInfo.bind(this);
  }

  ScanRMP(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<ScanRMPResponse> {
    return this.rpc.unary(RemoteMediaPlayerServiceScanRMPDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  Subscribe(
    request: DeepPartial<SubscribeRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Observable<RMPStatus> {
    return this.rpc.invoke(
      RemoteMediaPlayerServiceSubscribeDesc,
      SubscribeRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  DoAction(request: DeepPartial<DoActionRequest>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<Empty> {
    return this.rpc.unary(
      RemoteMediaPlayerServiceDoActionDesc,
      DoActionRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  GetPositionInfo(
    request: DeepPartial<GetPositionInfoRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<GetPositionInfoResponse> {
    return this.rpc.unary(
      RemoteMediaPlayerServiceGetPositionInfoDesc,
      GetPositionInfoRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }
}

export const RemoteMediaPlayerServiceDesc = { serviceName: "lzc.dlna.RemoteMediaPlayerService" };

export const RemoteMediaPlayerServiceScanRMPDesc: UnaryMethodDefinitionish = {
  methodName: "ScanRMP",
  service: RemoteMediaPlayerServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ScanRMPResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RemoteMediaPlayerServiceSubscribeDesc: UnaryMethodDefinitionish = {
  methodName: "Subscribe",
  service: RemoteMediaPlayerServiceDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return SubscribeRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = RMPStatus.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RemoteMediaPlayerServiceDoActionDesc: UnaryMethodDefinitionish = {
  methodName: "DoAction",
  service: RemoteMediaPlayerServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return DoActionRequest.encode(this).finish();
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

export const RemoteMediaPlayerServiceGetPositionInfoDesc: UnaryMethodDefinitionish = {
  methodName: "GetPositionInfo",
  service: RemoteMediaPlayerServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return GetPositionInfoRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = GetPositionInfoResponse.decode(data);
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

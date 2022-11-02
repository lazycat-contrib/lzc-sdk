/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { Empty } from "../google/protobuf/empty";
import { BrowserHeaders } from "browser-headers";
import { share } from "rxjs/operators";
import { Observable } from "rxjs";
import _m0 from "protobufjs/minimal";

export interface DoActionRequest {
  playerUuid: string;
  action: DoActionRequest_Action;
  mediaFile?: string | undefined;
  mediaSubtitle?: string | undefined;
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
  UNRECOGNIZED = -1,
}

export function doActionRequest_ActionFromJSON(
  object: any
): DoActionRequest_Action {
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
    case -1:
    case "UNRECOGNIZED":
    default:
      return DoActionRequest_Action.UNRECOGNIZED;
  }
}

export function doActionRequest_ActionToJSON(
  object: DoActionRequest_Action
): string {
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
}

function createBaseDoActionRequest(): DoActionRequest {
  return {
    playerUuid: "",
    action: 0,
    mediaFile: undefined,
    mediaSubtitle: undefined,
  };
}

export const DoActionRequest = {
  encode(
    message: DoActionRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
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
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DoActionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDoActionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.playerUuid = reader.string();
          break;
        case 2:
          message.action = reader.int32() as any;
          break;
        case 3:
          message.mediaFile = reader.string();
          break;
        case 4:
          message.mediaSubtitle = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DoActionRequest {
    return {
      playerUuid: isSet(object.playerUuid) ? String(object.playerUuid) : "",
      action: isSet(object.action)
        ? doActionRequest_ActionFromJSON(object.action)
        : 0,
      mediaFile: isSet(object.mediaFile) ? String(object.mediaFile) : undefined,
      mediaSubtitle: isSet(object.mediaSubtitle)
        ? String(object.mediaSubtitle)
        : undefined,
    };
  },

  toJSON(message: DoActionRequest): unknown {
    const obj: any = {};
    message.playerUuid !== undefined && (obj.playerUuid = message.playerUuid);
    message.action !== undefined &&
      (obj.action = doActionRequest_ActionToJSON(message.action));
    message.mediaFile !== undefined && (obj.mediaFile = message.mediaFile);
    message.mediaSubtitle !== undefined &&
      (obj.mediaSubtitle = message.mediaSubtitle);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DoActionRequest>, I>>(
    object: I
  ): DoActionRequest {
    const message = createBaseDoActionRequest();
    message.playerUuid = object.playerUuid ?? "";
    message.action = object.action ?? 0;
    message.mediaFile = object.mediaFile ?? undefined;
    message.mediaSubtitle = object.mediaSubtitle ?? undefined;
    return message;
  },
};

function createBaseRMPStatus(): RMPStatus {
  return { status: 0 };
}

export const RMPStatus = {
  encode(
    message: RMPStatus,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RMPStatus {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRMPStatus();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.status = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RMPStatus {
    return {
      status: isSet(object.status)
        ? rMPStatus_StatusFromJSON(object.status)
        : 0,
    };
  },

  toJSON(message: RMPStatus): unknown {
    const obj: any = {};
    message.status !== undefined &&
      (obj.status = rMPStatus_StatusToJSON(message.status));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RMPStatus>, I>>(
    object: I
  ): RMPStatus {
    const message = createBaseRMPStatus();
    message.status = object.status ?? 0;
    return message;
  },
};

function createBaseScanRMPResponse(): ScanRMPResponse {
  return { remoteMediaPlayers: [] };
}

export const ScanRMPResponse = {
  encode(
    message: ScanRMPResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.remoteMediaPlayers) {
      RemoteMediaPlayer.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ScanRMPResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseScanRMPResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.remoteMediaPlayers.push(
            RemoteMediaPlayer.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ScanRMPResponse {
    return {
      remoteMediaPlayers: Array.isArray(object?.remoteMediaPlayers)
        ? object.remoteMediaPlayers.map((e: any) =>
            RemoteMediaPlayer.fromJSON(e)
          )
        : [],
    };
  },

  toJSON(message: ScanRMPResponse): unknown {
    const obj: any = {};
    if (message.remoteMediaPlayers) {
      obj.remoteMediaPlayers = message.remoteMediaPlayers.map((e) =>
        e ? RemoteMediaPlayer.toJSON(e) : undefined
      );
    } else {
      obj.remoteMediaPlayers = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ScanRMPResponse>, I>>(
    object: I
  ): ScanRMPResponse {
    const message = createBaseScanRMPResponse();
    message.remoteMediaPlayers =
      object.remoteMediaPlayers?.map((e) => RemoteMediaPlayer.fromPartial(e)) ||
      [];
    return message;
  },
};

function createBaseSubscribeRequest(): SubscribeRequest {
  return { playerUuid: "" };
}

export const SubscribeRequest = {
  encode(
    message: SubscribeRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.playerUuid !== "") {
      writer.uint32(10).string(message.playerUuid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SubscribeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSubscribeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.playerUuid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SubscribeRequest {
    return {
      playerUuid: isSet(object.playerUuid) ? String(object.playerUuid) : "",
    };
  },

  toJSON(message: SubscribeRequest): unknown {
    const obj: any = {};
    message.playerUuid !== undefined && (obj.playerUuid = message.playerUuid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<SubscribeRequest>, I>>(
    object: I
  ): SubscribeRequest {
    const message = createBaseSubscribeRequest();
    message.playerUuid = object.playerUuid ?? "";
    return message;
  },
};

function createBaseRemoteMediaPlayer(): RemoteMediaPlayer {
  return { uuid: "", name: "", iconData: "" };
}

export const RemoteMediaPlayer = {
  encode(
    message: RemoteMediaPlayer,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.iconData !== "") {
      writer.uint32(26).string(message.iconData);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RemoteMediaPlayer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRemoteMediaPlayer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uuid = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.iconData = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RemoteMediaPlayer {
    return {
      uuid: isSet(object.uuid) ? String(object.uuid) : "",
      name: isSet(object.name) ? String(object.name) : "",
      iconData: isSet(object.iconData) ? String(object.iconData) : "",
    };
  },

  toJSON(message: RemoteMediaPlayer): unknown {
    const obj: any = {};
    message.uuid !== undefined && (obj.uuid = message.uuid);
    message.name !== undefined && (obj.name = message.name);
    message.iconData !== undefined && (obj.iconData = message.iconData);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RemoteMediaPlayer>, I>>(
    object: I
  ): RemoteMediaPlayer {
    const message = createBaseRemoteMediaPlayer();
    message.uuid = object.uuid ?? "";
    message.name = object.name ?? "";
    message.iconData = object.iconData ?? "";
    return message;
  },
};

/** 目前支持搜索DLNA的Render设备，并投送媒体文件 */
export interface RemoteMediaPlayerService {
  ScanRMP(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<ScanRMPResponse>;
  /** 立刻返回选择的当前RMP状态，并在状态变动时重新发送 */
  Subscribe(
    request: DeepPartial<SubscribeRequest>,
    metadata?: grpc.Metadata
  ): Observable<RMPStatus>;
  DoAction(
    request: DeepPartial<DoActionRequest>,
    metadata?: grpc.Metadata
  ): Promise<Empty>;
}

export class RemoteMediaPlayerServiceClientImpl
  implements RemoteMediaPlayerService
{
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ScanRMP = this.ScanRMP.bind(this);
    this.Subscribe = this.Subscribe.bind(this);
    this.DoAction = this.DoAction.bind(this);
  }

  ScanRMP(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<ScanRMPResponse> {
    return this.rpc.unary(
      RemoteMediaPlayerServiceScanRMPDesc,
      Empty.fromPartial(request),
      metadata
    );
  }

  Subscribe(
    request: DeepPartial<SubscribeRequest>,
    metadata?: grpc.Metadata
  ): Observable<RMPStatus> {
    return this.rpc.invoke(
      RemoteMediaPlayerServiceSubscribeDesc,
      SubscribeRequest.fromPartial(request),
      metadata
    );
  }

  DoAction(
    request: DeepPartial<DoActionRequest>,
    metadata?: grpc.Metadata
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteMediaPlayerServiceDoActionDesc,
      DoActionRequest.fromPartial(request),
      metadata
    );
  }
}

export const RemoteMediaPlayerServiceDesc = {
  serviceName: "lzc.dlna.RemoteMediaPlayerService",
};

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
      return {
        ...ScanRMPResponse.decode(data),
        toObject() {
          return this;
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
      return {
        ...RMPStatus.decode(data),
        toObject() {
          return this;
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
    const upStreamCodes = this.options.upStreamRetryCodes || [];
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
          onEnd: (
            code: grpc.Code,
            message: string,
            trailers: grpc.Metadata
          ) => {
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

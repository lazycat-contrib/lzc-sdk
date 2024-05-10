/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";

export interface PinAppRequest {
  url: string;
  appName: string;
  /** icon地址允许是http类型或者base64类型,如果为空将尝试获取网站的favicon.ico */
  iconUrl: string;
}

export interface PinAppReply {
}

export interface UnPinAppRequest {
  url: string;
}

export interface UnPinAppReply {
}

export interface OpenAppRequest {
  url: string;
  isFullScreen: boolean;
  appid: string;
  version: string;
  title: string;
  icon: string;
  /**
   * 当已经存在当前地址的窗口的时候，是否要求使用该地址强制打开。
   * 使用场景，比如播放器存在已经存在，这个时候需要使用过新的链接打开，然后覆盖上一次的播放界面，然后指定forceOpen=true即可.
   */
  forceOpen: boolean;
}

export interface OpenAppReply {
}

export interface OpenUnsafeAppRequest {
  /** control-view的地址 */
  url: string;
  /** control-view的位置 */
  position: OpenUnsafeAppRequest_Position;
  /** control-view的高度(当control-view 在左右的时候表示宽度， 上下表示高度） */
  height: string;
  /**
   * 此外control-view会接收以下特殊事件, 所有事件均统一以post-message形式发送， msg格式为`{ type: string, msg: string }`
   * control-view需要自行调用addEventListener('message'), 并过滤type==lzc_control_api类型的message.
   *
   * 1. OnURLChange(new_url)
   * 2. OnNewLinkClick(url)
   * 3. OnDownloadLinkRequest(url)
   * 4. OnNewResourceLink(url) 将所有的content-view请求的url通知给control-view
   * 5. OnContentMessage(msg string) //在content-view中调用post-message发送的任何消息
   */
  features: OpenUnsafeAppRequest_Feature[];
}

export enum OpenUnsafeAppRequest_Position {
  Left = 0,
  Right = 1,
  Top = 2,
  Bottom = 3,
  UNRECOGNIZED = -1,
}

export function openUnsafeAppRequest_PositionFromJSON(object: any): OpenUnsafeAppRequest_Position {
  switch (object) {
    case 0:
    case "Left":
      return OpenUnsafeAppRequest_Position.Left;
    case 1:
    case "Right":
      return OpenUnsafeAppRequest_Position.Right;
    case 2:
    case "Top":
      return OpenUnsafeAppRequest_Position.Top;
    case 3:
    case "Bottom":
      return OpenUnsafeAppRequest_Position.Bottom;
    case -1:
    case "UNRECOGNIZED":
    default:
      return OpenUnsafeAppRequest_Position.UNRECOGNIZED;
  }
}

export function openUnsafeAppRequest_PositionToJSON(object: OpenUnsafeAppRequest_Position): string {
  switch (object) {
    case OpenUnsafeAppRequest_Position.Left:
      return "Left";
    case OpenUnsafeAppRequest_Position.Right:
      return "Right";
    case OpenUnsafeAppRequest_Position.Top:
      return "Top";
    case OpenUnsafeAppRequest_Position.Bottom:
      return "Bottom";
    case OpenUnsafeAppRequest_Position.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

/** 所有的功能都通过 window.lzc_control_api对象注入. (后面以$API代指) */
export enum OpenUnsafeAppRequest_Feature {
  /** InjectJS - $API.InjectContentJS(jscontent,callback)  content-view 执行的jscontent, 如果有结果，会将结果通过callback(result:Boolean): 通知给control-view */
  InjectJS = 0,
  /** ReadCookie - $API.ReadCookie(domain) -> string  读取content-view cookie */
  ReadCookie = 1,
  UNRECOGNIZED = -1,
}

export function openUnsafeAppRequest_FeatureFromJSON(object: any): OpenUnsafeAppRequest_Feature {
  switch (object) {
    case 0:
    case "InjectJS":
      return OpenUnsafeAppRequest_Feature.InjectJS;
    case 1:
    case "ReadCookie":
      return OpenUnsafeAppRequest_Feature.ReadCookie;
    case -1:
    case "UNRECOGNIZED":
    default:
      return OpenUnsafeAppRequest_Feature.UNRECOGNIZED;
  }
}

export function openUnsafeAppRequest_FeatureToJSON(object: OpenUnsafeAppRequest_Feature): string {
  switch (object) {
    case OpenUnsafeAppRequest_Feature.InjectJS:
      return "InjectJS";
    case OpenUnsafeAppRequest_Feature.ReadCookie:
      return "ReadCookie";
    case OpenUnsafeAppRequest_Feature.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface OpenAppMethodRequest {
}

export interface OpenAppMethodReply {
  /** 当前设备支持的应用打开方式 */
  support: OpenAppMethodReply_Mode;
}

export enum OpenAppMethodReply_Mode {
  All = 0,
  AllNot = 1,
  Browser = 2,
  Client = 3,
  UNRECOGNIZED = -1,
}

export function openAppMethodReply_ModeFromJSON(object: any): OpenAppMethodReply_Mode {
  switch (object) {
    case 0:
    case "All":
      return OpenAppMethodReply_Mode.All;
    case 1:
    case "AllNot":
      return OpenAppMethodReply_Mode.AllNot;
    case 2:
    case "Browser":
      return OpenAppMethodReply_Mode.Browser;
    case 3:
    case "Client":
      return OpenAppMethodReply_Mode.Client;
    case -1:
    case "UNRECOGNIZED":
    default:
      return OpenAppMethodReply_Mode.UNRECOGNIZED;
  }
}

export function openAppMethodReply_ModeToJSON(object: OpenAppMethodReply_Mode): string {
  switch (object) {
    case OpenAppMethodReply_Mode.All:
      return "All";
    case OpenAppMethodReply_Mode.AllNot:
      return "AllNot";
    case OpenAppMethodReply_Mode.Browser:
      return "Browser";
    case OpenAppMethodReply_Mode.Client:
      return "Client";
    case OpenAppMethodReply_Mode.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

function createBasePinAppRequest(): PinAppRequest {
  return { url: "", appName: "", iconUrl: "" };
}

export const PinAppRequest = {
  encode(message: PinAppRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.appName !== "") {
      writer.uint32(18).string(message.appName);
    }
    if (message.iconUrl !== "") {
      writer.uint32(26).string(message.iconUrl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PinAppRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePinAppRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.url = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.appName = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.iconUrl = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): PinAppRequest {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      appName: isSet(object.appName) ? String(object.appName) : "",
      iconUrl: isSet(object.iconUrl) ? String(object.iconUrl) : "",
    };
  },

  toJSON(message: PinAppRequest): unknown {
    const obj: any = {};
    if (message.url !== "") {
      obj.url = message.url;
    }
    if (message.appName !== "") {
      obj.appName = message.appName;
    }
    if (message.iconUrl !== "") {
      obj.iconUrl = message.iconUrl;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<PinAppRequest>, I>>(base?: I): PinAppRequest {
    return PinAppRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PinAppRequest>, I>>(object: I): PinAppRequest {
    const message = createBasePinAppRequest();
    message.url = object.url ?? "";
    message.appName = object.appName ?? "";
    message.iconUrl = object.iconUrl ?? "";
    return message;
  },
};

function createBasePinAppReply(): PinAppReply {
  return {};
}

export const PinAppReply = {
  encode(_: PinAppReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PinAppReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePinAppReply();
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

  fromJSON(_: any): PinAppReply {
    return {};
  },

  toJSON(_: PinAppReply): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<PinAppReply>, I>>(base?: I): PinAppReply {
    return PinAppReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<PinAppReply>, I>>(_: I): PinAppReply {
    const message = createBasePinAppReply();
    return message;
  },
};

function createBaseUnPinAppRequest(): UnPinAppRequest {
  return { url: "" };
}

export const UnPinAppRequest = {
  encode(message: UnPinAppRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnPinAppRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnPinAppRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.url = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UnPinAppRequest {
    return { url: isSet(object.url) ? String(object.url) : "" };
  },

  toJSON(message: UnPinAppRequest): unknown {
    const obj: any = {};
    if (message.url !== "") {
      obj.url = message.url;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UnPinAppRequest>, I>>(base?: I): UnPinAppRequest {
    return UnPinAppRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UnPinAppRequest>, I>>(object: I): UnPinAppRequest {
    const message = createBaseUnPinAppRequest();
    message.url = object.url ?? "";
    return message;
  },
};

function createBaseUnPinAppReply(): UnPinAppReply {
  return {};
}

export const UnPinAppReply = {
  encode(_: UnPinAppReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnPinAppReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnPinAppReply();
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

  fromJSON(_: any): UnPinAppReply {
    return {};
  },

  toJSON(_: UnPinAppReply): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<UnPinAppReply>, I>>(base?: I): UnPinAppReply {
    return UnPinAppReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<UnPinAppReply>, I>>(_: I): UnPinAppReply {
    const message = createBaseUnPinAppReply();
    return message;
  },
};

function createBaseOpenAppRequest(): OpenAppRequest {
  return { url: "", isFullScreen: false, appid: "", version: "", title: "", icon: "", forceOpen: false };
}

export const OpenAppRequest = {
  encode(message: OpenAppRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.isFullScreen === true) {
      writer.uint32(16).bool(message.isFullScreen);
    }
    if (message.appid !== "") {
      writer.uint32(26).string(message.appid);
    }
    if (message.version !== "") {
      writer.uint32(34).string(message.version);
    }
    if (message.title !== "") {
      writer.uint32(42).string(message.title);
    }
    if (message.icon !== "") {
      writer.uint32(50).string(message.icon);
    }
    if (message.forceOpen === true) {
      writer.uint32(56).bool(message.forceOpen);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenAppRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenAppRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.url = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.isFullScreen = reader.bool();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.appid = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.version = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.title = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.icon = reader.string();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.forceOpen = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OpenAppRequest {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      isFullScreen: isSet(object.isFullScreen) ? Boolean(object.isFullScreen) : false,
      appid: isSet(object.appid) ? String(object.appid) : "",
      version: isSet(object.version) ? String(object.version) : "",
      title: isSet(object.title) ? String(object.title) : "",
      icon: isSet(object.icon) ? String(object.icon) : "",
      forceOpen: isSet(object.forceOpen) ? Boolean(object.forceOpen) : false,
    };
  },

  toJSON(message: OpenAppRequest): unknown {
    const obj: any = {};
    if (message.url !== "") {
      obj.url = message.url;
    }
    if (message.isFullScreen === true) {
      obj.isFullScreen = message.isFullScreen;
    }
    if (message.appid !== "") {
      obj.appid = message.appid;
    }
    if (message.version !== "") {
      obj.version = message.version;
    }
    if (message.title !== "") {
      obj.title = message.title;
    }
    if (message.icon !== "") {
      obj.icon = message.icon;
    }
    if (message.forceOpen === true) {
      obj.forceOpen = message.forceOpen;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenAppRequest>, I>>(base?: I): OpenAppRequest {
    return OpenAppRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OpenAppRequest>, I>>(object: I): OpenAppRequest {
    const message = createBaseOpenAppRequest();
    message.url = object.url ?? "";
    message.isFullScreen = object.isFullScreen ?? false;
    message.appid = object.appid ?? "";
    message.version = object.version ?? "";
    message.title = object.title ?? "";
    message.icon = object.icon ?? "";
    message.forceOpen = object.forceOpen ?? false;
    return message;
  },
};

function createBaseOpenAppReply(): OpenAppReply {
  return {};
}

export const OpenAppReply = {
  encode(_: OpenAppReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenAppReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenAppReply();
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

  fromJSON(_: any): OpenAppReply {
    return {};
  },

  toJSON(_: OpenAppReply): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenAppReply>, I>>(base?: I): OpenAppReply {
    return OpenAppReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OpenAppReply>, I>>(_: I): OpenAppReply {
    const message = createBaseOpenAppReply();
    return message;
  },
};

function createBaseOpenUnsafeAppRequest(): OpenUnsafeAppRequest {
  return { url: "", position: 0, height: "", features: [] };
}

export const OpenUnsafeAppRequest = {
  encode(message: OpenUnsafeAppRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.position !== 0) {
      writer.uint32(16).int32(message.position);
    }
    if (message.height !== "") {
      writer.uint32(26).string(message.height);
    }
    writer.uint32(34).fork();
    for (const v of message.features) {
      writer.int32(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenUnsafeAppRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenUnsafeAppRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.url = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.position = reader.int32() as any;
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.height = reader.string();
          continue;
        case 4:
          if (tag === 32) {
            message.features.push(reader.int32() as any);

            continue;
          }

          if (tag === 34) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.features.push(reader.int32() as any);
            }

            continue;
          }

          break;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OpenUnsafeAppRequest {
    return {
      url: isSet(object.url) ? String(object.url) : "",
      position: isSet(object.position) ? openUnsafeAppRequest_PositionFromJSON(object.position) : 0,
      height: isSet(object.height) ? String(object.height) : "",
      features: Array.isArray(object?.features)
        ? object.features.map((e: any) => openUnsafeAppRequest_FeatureFromJSON(e))
        : [],
    };
  },

  toJSON(message: OpenUnsafeAppRequest): unknown {
    const obj: any = {};
    if (message.url !== "") {
      obj.url = message.url;
    }
    if (message.position !== 0) {
      obj.position = openUnsafeAppRequest_PositionToJSON(message.position);
    }
    if (message.height !== "") {
      obj.height = message.height;
    }
    if (message.features?.length) {
      obj.features = message.features.map((e) => openUnsafeAppRequest_FeatureToJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenUnsafeAppRequest>, I>>(base?: I): OpenUnsafeAppRequest {
    return OpenUnsafeAppRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OpenUnsafeAppRequest>, I>>(object: I): OpenUnsafeAppRequest {
    const message = createBaseOpenUnsafeAppRequest();
    message.url = object.url ?? "";
    message.position = object.position ?? 0;
    message.height = object.height ?? "";
    message.features = object.features?.map((e) => e) || [];
    return message;
  },
};

function createBaseOpenAppMethodRequest(): OpenAppMethodRequest {
  return {};
}

export const OpenAppMethodRequest = {
  encode(_: OpenAppMethodRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenAppMethodRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenAppMethodRequest();
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

  fromJSON(_: any): OpenAppMethodRequest {
    return {};
  },

  toJSON(_: OpenAppMethodRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenAppMethodRequest>, I>>(base?: I): OpenAppMethodRequest {
    return OpenAppMethodRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OpenAppMethodRequest>, I>>(_: I): OpenAppMethodRequest {
    const message = createBaseOpenAppMethodRequest();
    return message;
  },
};

function createBaseOpenAppMethodReply(): OpenAppMethodReply {
  return { support: 0 };
}

export const OpenAppMethodReply = {
  encode(message: OpenAppMethodReply, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.support !== 0) {
      writer.uint32(8).int32(message.support);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OpenAppMethodReply {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOpenAppMethodReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.support = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): OpenAppMethodReply {
    return { support: isSet(object.support) ? openAppMethodReply_ModeFromJSON(object.support) : 0 };
  },

  toJSON(message: OpenAppMethodReply): unknown {
    const obj: any = {};
    if (message.support !== 0) {
      obj.support = openAppMethodReply_ModeToJSON(message.support);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<OpenAppMethodReply>, I>>(base?: I): OpenAppMethodReply {
    return OpenAppMethodReply.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<OpenAppMethodReply>, I>>(object: I): OpenAppMethodReply {
    const message = createBaseOpenAppMethodReply();
    message.support = object.support ?? 0;
    return message;
  },
};

export interface LocalLaunchService {
  /** 创建快捷方式 */
  PinApp(
    request: DeepPartial<PinAppRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<PinAppReply>;
  /** 删除快捷方式 */
  UnPinApp(
    request: DeepPartial<UnPinAppRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<UnPinAppReply>;
  /** 打开指定的懒猫云应用 */
  OpenApp(
    request: DeepPartial<OpenAppRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<OpenAppReply>;
  /**
   * 使用不安全模式打开指定APP
   * App页面由两个webview组成：一个control-view,一个content-view
   * 其中app渲染在control-view中，content-view的内容由App代码控制
   */
  OpenUnsafeApp(
    request: DeepPartial<OpenUnsafeAppRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<OpenAppReply>;
  /** 当前设备支持的应用打开方式 */
  OpenAppMethod(
    request: DeepPartial<OpenAppMethodRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<OpenAppMethodReply>;
}

export class LocalLaunchServiceClientImpl implements LocalLaunchService {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.PinApp = this.PinApp.bind(this);
    this.UnPinApp = this.UnPinApp.bind(this);
    this.OpenApp = this.OpenApp.bind(this);
    this.OpenUnsafeApp = this.OpenUnsafeApp.bind(this);
    this.OpenAppMethod = this.OpenAppMethod.bind(this);
  }

  PinApp(
    request: DeepPartial<PinAppRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<PinAppReply> {
    return this.rpc.unary(LocalLaunchServicePinAppDesc, PinAppRequest.fromPartial(request), metadata, abortSignal);
  }

  UnPinApp(
    request: DeepPartial<UnPinAppRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<UnPinAppReply> {
    return this.rpc.unary(LocalLaunchServiceUnPinAppDesc, UnPinAppRequest.fromPartial(request), metadata, abortSignal);
  }

  OpenApp(
    request: DeepPartial<OpenAppRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<OpenAppReply> {
    return this.rpc.unary(LocalLaunchServiceOpenAppDesc, OpenAppRequest.fromPartial(request), metadata, abortSignal);
  }

  OpenUnsafeApp(
    request: DeepPartial<OpenUnsafeAppRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<OpenAppReply> {
    return this.rpc.unary(
      LocalLaunchServiceOpenUnsafeAppDesc,
      OpenUnsafeAppRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  OpenAppMethod(
    request: DeepPartial<OpenAppMethodRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<OpenAppMethodReply> {
    return this.rpc.unary(
      LocalLaunchServiceOpenAppMethodDesc,
      OpenAppMethodRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }
}

export const LocalLaunchServiceDesc = { serviceName: "cloud.lazycat.apis.localdevice.LocalLaunchService" };

export const LocalLaunchServicePinAppDesc: UnaryMethodDefinitionish = {
  methodName: "PinApp",
  service: LocalLaunchServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return PinAppRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = PinAppReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const LocalLaunchServiceUnPinAppDesc: UnaryMethodDefinitionish = {
  methodName: "UnPinApp",
  service: LocalLaunchServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return UnPinAppRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = UnPinAppReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const LocalLaunchServiceOpenAppDesc: UnaryMethodDefinitionish = {
  methodName: "OpenApp",
  service: LocalLaunchServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return OpenAppRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = OpenAppReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const LocalLaunchServiceOpenUnsafeAppDesc: UnaryMethodDefinitionish = {
  methodName: "OpenUnsafeApp",
  service: LocalLaunchServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return OpenUnsafeAppRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = OpenAppReply.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const LocalLaunchServiceOpenAppMethodDesc: UnaryMethodDefinitionish = {
  methodName: "OpenAppMethod",
  service: LocalLaunchServiceDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return OpenAppMethodRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = OpenAppMethodReply.decode(data);
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

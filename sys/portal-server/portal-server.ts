/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import * as _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { Empty } from "../../google/protobuf/empty";
import { BrowserHeaders } from "browser-headers";
import { share } from "rxjs/operators";
import { Timestamp } from "../../google/protobuf/timestamp";

export interface ListUsersReply {
  uids: string[];
}

export interface PairDevicesRequest {
  srcId: string;
  otherIds: string[];
}

export interface AllocVEIPRequest {
  /** 简短说明使用原因，方便后续管理 */
  usage: string;
  /**
   * 关联的子域名。hserver会自动将$subdomain.$boxdomain的AAAA/A记录关联到对应的IP上
   * 1. 外部系统需要自行确保subdomain不会冲突
   */
  subdomain: string;
}

export interface AllocVEIPReply {
  ip: string;
}

export interface FreeVEIPRequest {
  ip: string;
  /**
   * 释放VEIP时为了避免DNS cache的问题，会在删除对应DNS record后等待一定时间后再真实的释放对应VEIP
   * 不指定则默认为10分钟
   */
  waitSecond: number;
}

export interface FreeVEIPReply {}

export interface AppCertRequest {
  appid: string;
  /** 单实例则为空，多实例为APP所部署时的UID */
  uid: string;
  /** 证书有效时间，单位s */
  validSeconds: number;
}

export interface AppCertReply {
  /** PEM格式的证书 */
  cert: string;
  /** PEM格式的私钥 */
  privateKey: string;
}

export interface AuthToken {
  /** 登陆时服务器返回的token */
  token: string;
  /**
   * 若设置了此地址，则在未检测到auth-token时，
   * 可以使用返回信息中的autologin_token_request_url进行自动查找auth-token
   * 此地址一般时ingress等组件上提供的服务。
   * 会通过html post form形式调用，附带token和redirect两个字段
   * token为自动搜索到的token，redirect为AuthToken.autologin_redirect_url的值。
   * 若token为空表示自动搜索token失败。
   */
  autologinTokenPostUrl: string;
  accessIp: string;
  autologinRedirectUrl: string;
}

export interface LoginInfo {
  uid: string;
  deviceId: string;
  when: Date | undefined;
  /**
   * 若uid为空, ingress应该返回此html内容给浏览器进行自动登陆。
   * 此html片段会尝试与hclient通讯获取auth-token后post到AuthToken.autologin_token_post_url上
   */
  autologinTokenRequestContent: string;
}

export interface PeerID {
  id: string;
}

export interface BoxInfo {
  /** 中心化服务器地址, 默认为origin.lazycat.cloud */
  originServer: string;
  /** fc03:1136:38/40地址 */
  virtualIp: string;
  /** p2p节点id */
  boxId: string;
  /** 注册到origin-server内的名称 */
  boxName: string;
  /** 从origin-server获取到的域名 */
  boxDomain: string;
  /** 证书管理器已经就续 */
  certReady: boolean;
  /** PEM格式的盒子证书, 作为盒子系统其他app cert的CA证书。 */
  boxCert: string;
}

export interface DomainCertRequest {
  domain: string;
}

export interface DomainCertReply {
  cert: string;
  key: string;
}

export interface Device {
  peerId: string;
  isOnline: boolean;
  /** 因为device api的监听端口可能会变化，所以此url有效性不会太长 */
  deviceApiUrl: string;
}

export interface ListDeviceRequest {
  uid: string;
}

export interface ListDeviceReply {
  devices: Device[];
}

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

function createBasePairDevicesRequest(): PairDevicesRequest {
  return { srcId: "", otherIds: [] };
}

export const PairDevicesRequest = {
  encode(
    message: PairDevicesRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.srcId !== "") {
      writer.uint32(10).string(message.srcId);
    }
    for (const v of message.otherIds) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PairDevicesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePairDevicesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.srcId = reader.string();
          break;
        case 2:
          message.otherIds.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PairDevicesRequest {
    return {
      srcId: isSet(object.srcId) ? String(object.srcId) : "",
      otherIds: Array.isArray(object?.otherIds)
        ? object.otherIds.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: PairDevicesRequest): unknown {
    const obj: any = {};
    message.srcId !== undefined && (obj.srcId = message.srcId);
    if (message.otherIds) {
      obj.otherIds = message.otherIds.map((e) => e);
    } else {
      obj.otherIds = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PairDevicesRequest>, I>>(
    object: I
  ): PairDevicesRequest {
    const message = createBasePairDevicesRequest();
    message.srcId = object.srcId ?? "";
    message.otherIds = object.otherIds?.map((e) => e) || [];
    return message;
  },
};

function createBaseAllocVEIPRequest(): AllocVEIPRequest {
  return { usage: "", subdomain: "" };
}

export const AllocVEIPRequest = {
  encode(
    message: AllocVEIPRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.usage !== "") {
      writer.uint32(10).string(message.usage);
    }
    if (message.subdomain !== "") {
      writer.uint32(18).string(message.subdomain);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AllocVEIPRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAllocVEIPRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.usage = reader.string();
          break;
        case 2:
          message.subdomain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AllocVEIPRequest {
    return {
      usage: isSet(object.usage) ? String(object.usage) : "",
      subdomain: isSet(object.subdomain) ? String(object.subdomain) : "",
    };
  },

  toJSON(message: AllocVEIPRequest): unknown {
    const obj: any = {};
    message.usage !== undefined && (obj.usage = message.usage);
    message.subdomain !== undefined && (obj.subdomain = message.subdomain);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AllocVEIPRequest>, I>>(
    object: I
  ): AllocVEIPRequest {
    const message = createBaseAllocVEIPRequest();
    message.usage = object.usage ?? "";
    message.subdomain = object.subdomain ?? "";
    return message;
  },
};

function createBaseAllocVEIPReply(): AllocVEIPReply {
  return { ip: "" };
}

export const AllocVEIPReply = {
  encode(
    message: AllocVEIPReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.ip !== "") {
      writer.uint32(10).string(message.ip);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AllocVEIPReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAllocVEIPReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ip = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AllocVEIPReply {
    return {
      ip: isSet(object.ip) ? String(object.ip) : "",
    };
  },

  toJSON(message: AllocVEIPReply): unknown {
    const obj: any = {};
    message.ip !== undefined && (obj.ip = message.ip);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AllocVEIPReply>, I>>(
    object: I
  ): AllocVEIPReply {
    const message = createBaseAllocVEIPReply();
    message.ip = object.ip ?? "";
    return message;
  },
};

function createBaseFreeVEIPRequest(): FreeVEIPRequest {
  return { ip: "", waitSecond: 0 };
}

export const FreeVEIPRequest = {
  encode(
    message: FreeVEIPRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.ip !== "") {
      writer.uint32(10).string(message.ip);
    }
    if (message.waitSecond !== 0) {
      writer.uint32(16).int32(message.waitSecond);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FreeVEIPRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFreeVEIPRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ip = reader.string();
          break;
        case 2:
          message.waitSecond = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FreeVEIPRequest {
    return {
      ip: isSet(object.ip) ? String(object.ip) : "",
      waitSecond: isSet(object.waitSecond) ? Number(object.waitSecond) : 0,
    };
  },

  toJSON(message: FreeVEIPRequest): unknown {
    const obj: any = {};
    message.ip !== undefined && (obj.ip = message.ip);
    message.waitSecond !== undefined &&
      (obj.waitSecond = Math.round(message.waitSecond));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FreeVEIPRequest>, I>>(
    object: I
  ): FreeVEIPRequest {
    const message = createBaseFreeVEIPRequest();
    message.ip = object.ip ?? "";
    message.waitSecond = object.waitSecond ?? 0;
    return message;
  },
};

function createBaseFreeVEIPReply(): FreeVEIPReply {
  return {};
}

export const FreeVEIPReply = {
  encode(
    _: FreeVEIPReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FreeVEIPReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFreeVEIPReply();
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

  fromJSON(_: any): FreeVEIPReply {
    return {};
  },

  toJSON(_: FreeVEIPReply): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FreeVEIPReply>, I>>(
    _: I
  ): FreeVEIPReply {
    const message = createBaseFreeVEIPReply();
    return message;
  },
};

function createBaseAppCertRequest(): AppCertRequest {
  return { appid: "", uid: "", validSeconds: 0 };
}

export const AppCertRequest = {
  encode(
    message: AppCertRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.appid !== "") {
      writer.uint32(10).string(message.appid);
    }
    if (message.uid !== "") {
      writer.uint32(18).string(message.uid);
    }
    if (message.validSeconds !== 0) {
      writer.uint32(24).int32(message.validSeconds);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppCertRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppCertRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.appid = reader.string();
          break;
        case 2:
          message.uid = reader.string();
          break;
        case 3:
          message.validSeconds = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppCertRequest {
    return {
      appid: isSet(object.appid) ? String(object.appid) : "",
      uid: isSet(object.uid) ? String(object.uid) : "",
      validSeconds: isSet(object.validSeconds)
        ? Number(object.validSeconds)
        : 0,
    };
  },

  toJSON(message: AppCertRequest): unknown {
    const obj: any = {};
    message.appid !== undefined && (obj.appid = message.appid);
    message.uid !== undefined && (obj.uid = message.uid);
    message.validSeconds !== undefined &&
      (obj.validSeconds = Math.round(message.validSeconds));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AppCertRequest>, I>>(
    object: I
  ): AppCertRequest {
    const message = createBaseAppCertRequest();
    message.appid = object.appid ?? "";
    message.uid = object.uid ?? "";
    message.validSeconds = object.validSeconds ?? 0;
    return message;
  },
};

function createBaseAppCertReply(): AppCertReply {
  return { cert: "", privateKey: "" };
}

export const AppCertReply = {
  encode(
    message: AppCertReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.cert !== "") {
      writer.uint32(10).string(message.cert);
    }
    if (message.privateKey !== "") {
      writer.uint32(18).string(message.privateKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppCertReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppCertReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cert = reader.string();
          break;
        case 2:
          message.privateKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AppCertReply {
    return {
      cert: isSet(object.cert) ? String(object.cert) : "",
      privateKey: isSet(object.privateKey) ? String(object.privateKey) : "",
    };
  },

  toJSON(message: AppCertReply): unknown {
    const obj: any = {};
    message.cert !== undefined && (obj.cert = message.cert);
    message.privateKey !== undefined && (obj.privateKey = message.privateKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AppCertReply>, I>>(
    object: I
  ): AppCertReply {
    const message = createBaseAppCertReply();
    message.cert = object.cert ?? "";
    message.privateKey = object.privateKey ?? "";
    return message;
  },
};

function createBaseAuthToken(): AuthToken {
  return {
    token: "",
    autologinTokenPostUrl: "",
    accessIp: "",
    autologinRedirectUrl: "",
  };
}

export const AuthToken = {
  encode(
    message: AuthToken,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.token !== "") {
      writer.uint32(10).string(message.token);
    }
    if (message.autologinTokenPostUrl !== "") {
      writer.uint32(18).string(message.autologinTokenPostUrl);
    }
    if (message.accessIp !== "") {
      writer.uint32(26).string(message.accessIp);
    }
    if (message.autologinRedirectUrl !== "") {
      writer.uint32(34).string(message.autologinRedirectUrl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AuthToken {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAuthToken();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token = reader.string();
          break;
        case 2:
          message.autologinTokenPostUrl = reader.string();
          break;
        case 3:
          message.accessIp = reader.string();
          break;
        case 4:
          message.autologinRedirectUrl = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AuthToken {
    return {
      token: isSet(object.token) ? String(object.token) : "",
      autologinTokenPostUrl: isSet(object.autologinTokenPostUrl)
        ? String(object.autologinTokenPostUrl)
        : "",
      accessIp: isSet(object.accessIp) ? String(object.accessIp) : "",
      autologinRedirectUrl: isSet(object.autologinRedirectUrl)
        ? String(object.autologinRedirectUrl)
        : "",
    };
  },

  toJSON(message: AuthToken): unknown {
    const obj: any = {};
    message.token !== undefined && (obj.token = message.token);
    message.autologinTokenPostUrl !== undefined &&
      (obj.autologinTokenPostUrl = message.autologinTokenPostUrl);
    message.accessIp !== undefined && (obj.accessIp = message.accessIp);
    message.autologinRedirectUrl !== undefined &&
      (obj.autologinRedirectUrl = message.autologinRedirectUrl);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AuthToken>, I>>(
    object: I
  ): AuthToken {
    const message = createBaseAuthToken();
    message.token = object.token ?? "";
    message.autologinTokenPostUrl = object.autologinTokenPostUrl ?? "";
    message.accessIp = object.accessIp ?? "";
    message.autologinRedirectUrl = object.autologinRedirectUrl ?? "";
    return message;
  },
};

function createBaseLoginInfo(): LoginInfo {
  return {
    uid: "",
    deviceId: "",
    when: undefined,
    autologinTokenRequestContent: "",
  };
}

export const LoginInfo = {
  encode(
    message: LoginInfo,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.deviceId !== "") {
      writer.uint32(18).string(message.deviceId);
    }
    if (message.when !== undefined) {
      Timestamp.encode(
        toTimestamp(message.when),
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.autologinTokenRequestContent !== "") {
      writer.uint32(34).string(message.autologinTokenRequestContent);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LoginInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoginInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        case 2:
          message.deviceId = reader.string();
          break;
        case 3:
          message.when = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.autologinTokenRequestContent = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LoginInfo {
    return {
      uid: isSet(object.uid) ? String(object.uid) : "",
      deviceId: isSet(object.deviceId) ? String(object.deviceId) : "",
      when: isSet(object.when) ? fromJsonTimestamp(object.when) : undefined,
      autologinTokenRequestContent: isSet(object.autologinTokenRequestContent)
        ? String(object.autologinTokenRequestContent)
        : "",
    };
  },

  toJSON(message: LoginInfo): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.deviceId !== undefined && (obj.deviceId = message.deviceId);
    message.when !== undefined && (obj.when = message.when.toISOString());
    message.autologinTokenRequestContent !== undefined &&
      (obj.autologinTokenRequestContent = message.autologinTokenRequestContent);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LoginInfo>, I>>(
    object: I
  ): LoginInfo {
    const message = createBaseLoginInfo();
    message.uid = object.uid ?? "";
    message.deviceId = object.deviceId ?? "";
    message.when = object.when ?? undefined;
    message.autologinTokenRequestContent =
      object.autologinTokenRequestContent ?? "";
    return message;
  },
};

function createBasePeerID(): PeerID {
  return { id: "" };
}

export const PeerID = {
  encode(
    message: PeerID,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PeerID {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePeerID();
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

  fromJSON(object: any): PeerID {
    return {
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: PeerID): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PeerID>, I>>(object: I): PeerID {
    const message = createBasePeerID();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseBoxInfo(): BoxInfo {
  return {
    originServer: "",
    virtualIp: "",
    boxId: "",
    boxName: "",
    boxDomain: "",
    certReady: false,
    boxCert: "",
  };
}

export const BoxInfo = {
  encode(
    message: BoxInfo,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.originServer !== "") {
      writer.uint32(10).string(message.originServer);
    }
    if (message.virtualIp !== "") {
      writer.uint32(18).string(message.virtualIp);
    }
    if (message.boxId !== "") {
      writer.uint32(26).string(message.boxId);
    }
    if (message.boxName !== "") {
      writer.uint32(34).string(message.boxName);
    }
    if (message.boxDomain !== "") {
      writer.uint32(42).string(message.boxDomain);
    }
    if (message.certReady === true) {
      writer.uint32(48).bool(message.certReady);
    }
    if (message.boxCert !== "") {
      writer.uint32(58).string(message.boxCert);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BoxInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBoxInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.originServer = reader.string();
          break;
        case 2:
          message.virtualIp = reader.string();
          break;
        case 3:
          message.boxId = reader.string();
          break;
        case 4:
          message.boxName = reader.string();
          break;
        case 5:
          message.boxDomain = reader.string();
          break;
        case 6:
          message.certReady = reader.bool();
          break;
        case 7:
          message.boxCert = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BoxInfo {
    return {
      originServer: isSet(object.originServer)
        ? String(object.originServer)
        : "",
      virtualIp: isSet(object.virtualIp) ? String(object.virtualIp) : "",
      boxId: isSet(object.boxId) ? String(object.boxId) : "",
      boxName: isSet(object.boxName) ? String(object.boxName) : "",
      boxDomain: isSet(object.boxDomain) ? String(object.boxDomain) : "",
      certReady: isSet(object.certReady) ? Boolean(object.certReady) : false,
      boxCert: isSet(object.boxCert) ? String(object.boxCert) : "",
    };
  },

  toJSON(message: BoxInfo): unknown {
    const obj: any = {};
    message.originServer !== undefined &&
      (obj.originServer = message.originServer);
    message.virtualIp !== undefined && (obj.virtualIp = message.virtualIp);
    message.boxId !== undefined && (obj.boxId = message.boxId);
    message.boxName !== undefined && (obj.boxName = message.boxName);
    message.boxDomain !== undefined && (obj.boxDomain = message.boxDomain);
    message.certReady !== undefined && (obj.certReady = message.certReady);
    message.boxCert !== undefined && (obj.boxCert = message.boxCert);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BoxInfo>, I>>(object: I): BoxInfo {
    const message = createBaseBoxInfo();
    message.originServer = object.originServer ?? "";
    message.virtualIp = object.virtualIp ?? "";
    message.boxId = object.boxId ?? "";
    message.boxName = object.boxName ?? "";
    message.boxDomain = object.boxDomain ?? "";
    message.certReady = object.certReady ?? false;
    message.boxCert = object.boxCert ?? "";
    return message;
  },
};

function createBaseDomainCertRequest(): DomainCertRequest {
  return { domain: "" };
}

export const DomainCertRequest = {
  encode(
    message: DomainCertRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.domain !== "") {
      writer.uint32(10).string(message.domain);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DomainCertRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDomainCertRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.domain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DomainCertRequest {
    return {
      domain: isSet(object.domain) ? String(object.domain) : "",
    };
  },

  toJSON(message: DomainCertRequest): unknown {
    const obj: any = {};
    message.domain !== undefined && (obj.domain = message.domain);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DomainCertRequest>, I>>(
    object: I
  ): DomainCertRequest {
    const message = createBaseDomainCertRequest();
    message.domain = object.domain ?? "";
    return message;
  },
};

function createBaseDomainCertReply(): DomainCertReply {
  return { cert: "", key: "" };
}

export const DomainCertReply = {
  encode(
    message: DomainCertReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.cert !== "") {
      writer.uint32(10).string(message.cert);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DomainCertReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDomainCertReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cert = reader.string();
          break;
        case 2:
          message.key = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DomainCertReply {
    return {
      cert: isSet(object.cert) ? String(object.cert) : "",
      key: isSet(object.key) ? String(object.key) : "",
    };
  },

  toJSON(message: DomainCertReply): unknown {
    const obj: any = {};
    message.cert !== undefined && (obj.cert = message.cert);
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DomainCertReply>, I>>(
    object: I
  ): DomainCertReply {
    const message = createBaseDomainCertReply();
    message.cert = object.cert ?? "";
    message.key = object.key ?? "";
    return message;
  },
};

function createBaseDevice(): Device {
  return { peerId: "", isOnline: false, deviceApiUrl: "" };
}

export const Device = {
  encode(
    message: Device,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.peerId !== "") {
      writer.uint32(10).string(message.peerId);
    }
    if (message.isOnline === true) {
      writer.uint32(16).bool(message.isOnline);
    }
    if (message.deviceApiUrl !== "") {
      writer.uint32(26).string(message.deviceApiUrl);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Device {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDevice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.peerId = reader.string();
          break;
        case 2:
          message.isOnline = reader.bool();
          break;
        case 3:
          message.deviceApiUrl = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Device {
    return {
      peerId: isSet(object.peerId) ? String(object.peerId) : "",
      isOnline: isSet(object.isOnline) ? Boolean(object.isOnline) : false,
      deviceApiUrl: isSet(object.deviceApiUrl)
        ? String(object.deviceApiUrl)
        : "",
    };
  },

  toJSON(message: Device): unknown {
    const obj: any = {};
    message.peerId !== undefined && (obj.peerId = message.peerId);
    message.isOnline !== undefined && (obj.isOnline = message.isOnline);
    message.deviceApiUrl !== undefined &&
      (obj.deviceApiUrl = message.deviceApiUrl);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Device>, I>>(object: I): Device {
    const message = createBaseDevice();
    message.peerId = object.peerId ?? "";
    message.isOnline = object.isOnline ?? false;
    message.deviceApiUrl = object.deviceApiUrl ?? "";
    return message;
  },
};

function createBaseListDeviceRequest(): ListDeviceRequest {
  return { uid: "" };
}

export const ListDeviceRequest = {
  encode(
    message: ListDeviceRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListDeviceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListDeviceRequest();
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

  fromJSON(object: any): ListDeviceRequest {
    return {
      uid: isSet(object.uid) ? String(object.uid) : "",
    };
  },

  toJSON(message: ListDeviceRequest): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListDeviceRequest>, I>>(
    object: I
  ): ListDeviceRequest {
    const message = createBaseListDeviceRequest();
    message.uid = object.uid ?? "";
    return message;
  },
};

function createBaseListDeviceReply(): ListDeviceReply {
  return { devices: [] };
}

export const ListDeviceReply = {
  encode(
    message: ListDeviceReply,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.devices) {
      Device.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListDeviceReply {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListDeviceReply();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.devices.push(Device.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListDeviceReply {
    return {
      devices: Array.isArray(object?.devices)
        ? object.devices.map((e: any) => Device.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ListDeviceReply): unknown {
    const obj: any = {};
    if (message.devices) {
      obj.devices = message.devices.map((e) =>
        e ? Device.toJSON(e) : undefined
      );
    } else {
      obj.devices = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ListDeviceReply>, I>>(
    object: I
  ): ListDeviceReply {
    const message = createBaseListDeviceReply();
    message.devices = object.devices?.map((e) => Device.fromPartial(e)) || [];
    return message;
  },
};

export interface HPortalSys {
  /** 用auth-token反向查询登陆的设备以及UID */
  QueryLogin(
    request: DeepPartial<AuthToken>,
    metadata?: grpc.Metadata
  ): Promise<LoginInfo>;
  /** 根据UID返回所有的设备列表 */
  ListDevices(
    request: DeepPartial<ListDeviceRequest>,
    metadata?: grpc.Metadata
  ): Promise<ListDeviceReply>;
  QueryDeviceByID(
    request: DeepPartial<PeerID>,
    metadata?: grpc.Metadata
  ): Promise<Device>;
  QueryBoxInfo(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<BoxInfo>;
  /**
   * 获取盒子所属域名下或下一级域名的https证书。
   * 注意不是所有ACME服务器都支持泛域名。
   */
  GetDomainCert(
    request: DeepPartial<DomainCertRequest>,
    metadata?: grpc.Metadata
  ): Promise<DomainCertReply>;
  /**
   * 在部署具体app前，调用此接口获取app证书
   * APP证书格式为:
   *   Issuer: O = $BOX.ORIGIN, CN = $BOX.DOMAIN, serialNumber = $BOX.ID
   *   Subject: O = $BOX.ORIGIN, CN = $APP.DOMAIN, serialNumber = '$UID@$APP.ID || $APP.ID'
   * Issuer为box.crt，通过QueryBoxInfo查询到BoxInfo.BoxCrt。并且box.crt的公钥与box.id是一一对应关系。
   *
   * 盒子内部组件可以直接通过QueryBoxInfo来验证信任链是否合法，盒子外部系统需要通过其他机制比如libp2p.identify去验证box.crt的合法性。
   */
  GetAppCert(
    request: DeepPartial<AppCertRequest>,
    metadata?: grpc.Metadata
  ): Promise<AppCertReply>;
  /** 申请额外的外部可访问的IP,并配置对应访问域名 */
  AllocVirtualExternalIP(
    request: DeepPartial<AllocVEIPRequest>,
    metadata?: grpc.Metadata
  ): Promise<AllocVEIPReply>;
  /** 释放虚拟IP */
  FreeVirtualExternalIP(
    request: DeepPartial<FreeVEIPRequest>,
    metadata?: grpc.Metadata
  ): Promise<FreeVEIPReply>;
  PairDevices(
    request: DeepPartial<PairDevicesRequest>,
    metadata?: grpc.Metadata
  ): Observable<Empty>;
  ListUsers(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<ListUsersReply>;
}

export class HPortalSysClientImpl implements HPortalSys {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.QueryLogin = this.QueryLogin.bind(this);
    this.ListDevices = this.ListDevices.bind(this);
    this.QueryDeviceByID = this.QueryDeviceByID.bind(this);
    this.QueryBoxInfo = this.QueryBoxInfo.bind(this);
    this.GetDomainCert = this.GetDomainCert.bind(this);
    this.GetAppCert = this.GetAppCert.bind(this);
    this.AllocVirtualExternalIP = this.AllocVirtualExternalIP.bind(this);
    this.FreeVirtualExternalIP = this.FreeVirtualExternalIP.bind(this);
    this.PairDevices = this.PairDevices.bind(this);
    this.ListUsers = this.ListUsers.bind(this);
  }

  QueryLogin(
    request: DeepPartial<AuthToken>,
    metadata?: grpc.Metadata
  ): Promise<LoginInfo> {
    return this.rpc.unary(
      HPortalSysQueryLoginDesc,
      AuthToken.fromPartial(request),
      metadata
    );
  }

  ListDevices(
    request: DeepPartial<ListDeviceRequest>,
    metadata?: grpc.Metadata
  ): Promise<ListDeviceReply> {
    return this.rpc.unary(
      HPortalSysListDevicesDesc,
      ListDeviceRequest.fromPartial(request),
      metadata
    );
  }

  QueryDeviceByID(
    request: DeepPartial<PeerID>,
    metadata?: grpc.Metadata
  ): Promise<Device> {
    return this.rpc.unary(
      HPortalSysQueryDeviceByIDDesc,
      PeerID.fromPartial(request),
      metadata
    );
  }

  QueryBoxInfo(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<BoxInfo> {
    return this.rpc.unary(
      HPortalSysQueryBoxInfoDesc,
      Empty.fromPartial(request),
      metadata
    );
  }

  GetDomainCert(
    request: DeepPartial<DomainCertRequest>,
    metadata?: grpc.Metadata
  ): Promise<DomainCertReply> {
    return this.rpc.unary(
      HPortalSysGetDomainCertDesc,
      DomainCertRequest.fromPartial(request),
      metadata
    );
  }

  GetAppCert(
    request: DeepPartial<AppCertRequest>,
    metadata?: grpc.Metadata
  ): Promise<AppCertReply> {
    return this.rpc.unary(
      HPortalSysGetAppCertDesc,
      AppCertRequest.fromPartial(request),
      metadata
    );
  }

  AllocVirtualExternalIP(
    request: DeepPartial<AllocVEIPRequest>,
    metadata?: grpc.Metadata
  ): Promise<AllocVEIPReply> {
    return this.rpc.unary(
      HPortalSysAllocVirtualExternalIPDesc,
      AllocVEIPRequest.fromPartial(request),
      metadata
    );
  }

  FreeVirtualExternalIP(
    request: DeepPartial<FreeVEIPRequest>,
    metadata?: grpc.Metadata
  ): Promise<FreeVEIPReply> {
    return this.rpc.unary(
      HPortalSysFreeVirtualExternalIPDesc,
      FreeVEIPRequest.fromPartial(request),
      metadata
    );
  }

  PairDevices(
    request: DeepPartial<PairDevicesRequest>,
    metadata?: grpc.Metadata
  ): Observable<Empty> {
    return this.rpc.invoke(
      HPortalSysPairDevicesDesc,
      PairDevicesRequest.fromPartial(request),
      metadata
    );
  }

  ListUsers(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata
  ): Promise<ListUsersReply> {
    return this.rpc.unary(
      HPortalSysListUsersDesc,
      Empty.fromPartial(request),
      metadata
    );
  }
}

export const HPortalSysDesc = {
  serviceName: "cloud.lazycat.sys.HPortalSys",
};

export const HPortalSysQueryLoginDesc: UnaryMethodDefinitionish = {
  methodName: "QueryLogin",
  service: HPortalSysDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return AuthToken.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...LoginInfo.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const HPortalSysListDevicesDesc: UnaryMethodDefinitionish = {
  methodName: "ListDevices",
  service: HPortalSysDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ListDeviceRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...ListDeviceReply.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const HPortalSysQueryDeviceByIDDesc: UnaryMethodDefinitionish = {
  methodName: "QueryDeviceByID",
  service: HPortalSysDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return PeerID.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...Device.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const HPortalSysQueryBoxInfoDesc: UnaryMethodDefinitionish = {
  methodName: "QueryBoxInfo",
  service: HPortalSysDesc,
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
        ...BoxInfo.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const HPortalSysGetDomainCertDesc: UnaryMethodDefinitionish = {
  methodName: "GetDomainCert",
  service: HPortalSysDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return DomainCertRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...DomainCertReply.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const HPortalSysGetAppCertDesc: UnaryMethodDefinitionish = {
  methodName: "GetAppCert",
  service: HPortalSysDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return AppCertRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...AppCertReply.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const HPortalSysAllocVirtualExternalIPDesc: UnaryMethodDefinitionish = {
  methodName: "AllocVirtualExternalIP",
  service: HPortalSysDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return AllocVEIPRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...AllocVEIPReply.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const HPortalSysFreeVirtualExternalIPDesc: UnaryMethodDefinitionish = {
  methodName: "FreeVirtualExternalIP",
  service: HPortalSysDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return FreeVEIPRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...FreeVEIPReply.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const HPortalSysPairDevicesDesc: UnaryMethodDefinitionish = {
  methodName: "PairDevices",
  service: HPortalSysDesc,
  requestStream: false,
  responseStream: true,
  requestType: {
    serializeBinary() {
      return PairDevicesRequest.encode(this).finish();
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

export const HPortalSysListUsersDesc: UnaryMethodDefinitionish = {
  methodName: "ListUsers",
  service: HPortalSysDesc,
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
        ...ListUsersReply.decode(data),
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

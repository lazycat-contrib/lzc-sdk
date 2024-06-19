/* eslint-disable */
import { grpc } from "@improbable-eng/grpc-web";
import { BrowserHeaders } from "browser-headers";
import _m0 from "protobufjs/minimal";
import { Observable } from "rxjs";
import { Empty } from "../google/protobuf/empty";

/** /usr/include/linux/input-event-codes.h */
export enum InputEvent {
  KEY_HOME = 0,
  KEY_BACK = 1,
  KEY_F11 = 2,
  KEY_MENU = 3,
  KEY_UP = 4,
  KEY_DOWN = 5,
  KEY_LEFT = 6,
  KEY_RIGHT = 7,
  KEY_PAGEUP = 8,
  KEY_PAGEDOWN = 9,
  KEY_ENTER = 10,
  KEY_BACKSPACE = 11,
  KEY_F10 = 12,
  UNRECOGNIZED = -1,
}

export function inputEventFromJSON(object: any): InputEvent {
  switch (object) {
    case 0:
    case "KEY_HOME":
      return InputEvent.KEY_HOME;
    case 1:
    case "KEY_BACK":
      return InputEvent.KEY_BACK;
    case 2:
    case "KEY_F11":
      return InputEvent.KEY_F11;
    case 3:
    case "KEY_MENU":
      return InputEvent.KEY_MENU;
    case 4:
    case "KEY_UP":
      return InputEvent.KEY_UP;
    case 5:
    case "KEY_DOWN":
      return InputEvent.KEY_DOWN;
    case 6:
    case "KEY_LEFT":
      return InputEvent.KEY_LEFT;
    case 7:
    case "KEY_RIGHT":
      return InputEvent.KEY_RIGHT;
    case 8:
    case "KEY_PAGEUP":
      return InputEvent.KEY_PAGEUP;
    case 9:
    case "KEY_PAGEDOWN":
      return InputEvent.KEY_PAGEDOWN;
    case 10:
    case "KEY_ENTER":
      return InputEvent.KEY_ENTER;
    case 11:
    case "KEY_BACKSPACE":
      return InputEvent.KEY_BACKSPACE;
    case 12:
    case "KEY_F10":
      return InputEvent.KEY_F10;
    case -1:
    case "UNRECOGNIZED":
    default:
      return InputEvent.UNRECOGNIZED;
  }
}

export function inputEventToJSON(object: InputEvent): string {
  switch (object) {
    case InputEvent.KEY_HOME:
      return "KEY_HOME";
    case InputEvent.KEY_BACK:
      return "KEY_BACK";
    case InputEvent.KEY_F11:
      return "KEY_F11";
    case InputEvent.KEY_MENU:
      return "KEY_MENU";
    case InputEvent.KEY_UP:
      return "KEY_UP";
    case InputEvent.KEY_DOWN:
      return "KEY_DOWN";
    case InputEvent.KEY_LEFT:
      return "KEY_LEFT";
    case InputEvent.KEY_RIGHT:
      return "KEY_RIGHT";
    case InputEvent.KEY_PAGEUP:
      return "KEY_PAGEUP";
    case InputEvent.KEY_PAGEDOWN:
      return "KEY_PAGEDOWN";
    case InputEvent.KEY_ENTER:
      return "KEY_ENTER";
    case InputEvent.KEY_BACKSPACE:
      return "KEY_BACKSPACE";
    case InputEvent.KEY_F10:
      return "KEY_F10";
    case InputEvent.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum EventState {
  PRESS = 0,
  RELEASE = 1,
  UNRECOGNIZED = -1,
}

export function eventStateFromJSON(object: any): EventState {
  switch (object) {
    case 0:
    case "PRESS":
      return EventState.PRESS;
    case 1:
    case "RELEASE":
      return EventState.RELEASE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return EventState.UNRECOGNIZED;
  }
}

export function eventStateToJSON(object: EventState): string {
  switch (object) {
    case EventState.PRESS:
      return "PRESS";
    case EventState.RELEASE:
      return "RELEASE";
    case EventState.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

/** uinput TouchPad interface */
export enum TouchEvent {
  MoveTo = 0,
  LeftClick = 1,
  RightClick = 2,
  LeftPress = 3,
  LeftRelease = 4,
  RightPress = 5,
  TouchDown = 6,
  TouchUp = 7,
  UNRECOGNIZED = -1,
}

export function touchEventFromJSON(object: any): TouchEvent {
  switch (object) {
    case 0:
    case "MoveTo":
      return TouchEvent.MoveTo;
    case 1:
    case "LeftClick":
      return TouchEvent.LeftClick;
    case 2:
    case "RightClick":
      return TouchEvent.RightClick;
    case 3:
    case "LeftPress":
      return TouchEvent.LeftPress;
    case 4:
    case "LeftRelease":
      return TouchEvent.LeftRelease;
    case 5:
    case "RightPress":
      return TouchEvent.RightPress;
    case 6:
    case "TouchDown":
      return TouchEvent.TouchDown;
    case 7:
    case "TouchUp":
      return TouchEvent.TouchUp;
    case -1:
    case "UNRECOGNIZED":
    default:
      return TouchEvent.UNRECOGNIZED;
  }
}

export function touchEventToJSON(object: TouchEvent): string {
  switch (object) {
    case TouchEvent.MoveTo:
      return "MoveTo";
    case TouchEvent.LeftClick:
      return "LeftClick";
    case TouchEvent.RightClick:
      return "RightClick";
    case TouchEvent.LeftPress:
      return "LeftPress";
    case TouchEvent.LeftRelease:
      return "LeftRelease";
    case TouchEvent.RightPress:
      return "RightPress";
    case TouchEvent.TouchDown:
      return "TouchDown";
    case TouchEvent.TouchUp:
      return "TouchUp";
    case TouchEvent.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface SendKeyboardEventRequest {
  code: InputEvent;
  state: EventState;
}

export interface MovePosition {
  x: number;
  y: number;
}

export interface SendTouchpadEventRequest {
  event: TouchEvent;
  position?: MovePosition | undefined;
}

export interface SendTouchpadMoveRequest {
  position: MovePosition | undefined;
}

export interface SendMouseMoveByDirectionRequest {
  pixel: number;
}

export interface SendMouseMoveRequest {
  x: number;
  y: number;
}

export interface SendMouseWheelRequest {
  horizontal: boolean;
  delta: number;
}

export interface SetRemoteScreenRectRequest {
  height: number;
  width: number;
}

export interface WriteClipboardRequest {
  text: string;
}

export interface ReadClipboardResponse {
  text: string;
}

export interface ListSinkInputsResponse {
  sinkInputs: ListSinkInputsResponse_SinkInput[];
}

export interface ListSinkInputsResponse_SinkInput {
  index: number;
  name: string;
  sink: number;
  muted: boolean;
  volume: number;
}

export interface ListSinksRepsonse {
  sinks: ListSinksRepsonse_Sink[];
}

export interface ListSinksRepsonse_Sink {
  index: number;
  name: string;
  description: string;
  muted: boolean;
  volume: number;
  cardIndex: number;
}

export interface CardProfile {
  name: string;
  description: string;
  available: boolean;
  priority: number;
  sinks: number;
  sources: number;
}

export interface ListCardsResponse {
  cards: ListCardsResponse_Card[];
}

export interface ListCardsResponse_Card {
  index: number;
  name: string;
  deviceName: string;
  activeProfile: CardProfile | undefined;
  profiles: CardProfile[];
}

export interface SetDefaultSinkRequest {
  sinkName: string;
}

export interface SetCardProfileRequest {
  cardIndex: number;
  profileName: string;
}

export interface SinkRequest {
  sinkName?: string | undefined;
}

export interface GetMuteResponse {
  muted: boolean;
}

export interface SetSinkMuteRequest {
  sinkName?: string | undefined;
  mute: boolean;
}

export interface GetSinkVolumeResponse {
  volume: number;
}

export interface ChangeVolumeRequest {
  value: number;
}

export interface SetSinkInputVolumeRequest {
  sinkInput: string;
  volume: number;
}

export interface BrowserActionRequest {
  action: BrowserActionRequest_Action;
}

export enum BrowserActionRequest_Action {
  CLOSE_TAB = 0,
  NEW_TAB = 1,
  REFRESH_TAB = 2,
  BACKWARD_HISTORY = 3,
  FORWARD_HISTORY = 4,
  UNRECOGNIZED = -1,
}

export function browserActionRequest_ActionFromJSON(object: any): BrowserActionRequest_Action {
  switch (object) {
    case 0:
    case "CLOSE_TAB":
      return BrowserActionRequest_Action.CLOSE_TAB;
    case 1:
    case "NEW_TAB":
      return BrowserActionRequest_Action.NEW_TAB;
    case 2:
    case "REFRESH_TAB":
      return BrowserActionRequest_Action.REFRESH_TAB;
    case 3:
    case "BACKWARD_HISTORY":
      return BrowserActionRequest_Action.BACKWARD_HISTORY;
    case 4:
    case "FORWARD_HISTORY":
      return BrowserActionRequest_Action.FORWARD_HISTORY;
    case -1:
    case "UNRECOGNIZED":
    default:
      return BrowserActionRequest_Action.UNRECOGNIZED;
  }
}

export function browserActionRequest_ActionToJSON(object: BrowserActionRequest_Action): string {
  switch (object) {
    case BrowserActionRequest_Action.CLOSE_TAB:
      return "CLOSE_TAB";
    case BrowserActionRequest_Action.NEW_TAB:
      return "NEW_TAB";
    case BrowserActionRequest_Action.REFRESH_TAB:
      return "REFRESH_TAB";
    case BrowserActionRequest_Action.BACKWARD_HISTORY:
      return "BACKWARD_HISTORY";
    case BrowserActionRequest_Action.FORWARD_HISTORY:
      return "FORWARD_HISTORY";
    case BrowserActionRequest_Action.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

function createBaseSendKeyboardEventRequest(): SendKeyboardEventRequest {
  return { code: 0, state: 0 };
}

export const SendKeyboardEventRequest = {
  encode(message: SendKeyboardEventRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== 0) {
      writer.uint32(8).int32(message.code);
    }
    if (message.state !== 0) {
      writer.uint32(16).int32(message.state);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendKeyboardEventRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendKeyboardEventRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.code = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.state = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SendKeyboardEventRequest {
    return {
      code: isSet(object.code) ? inputEventFromJSON(object.code) : 0,
      state: isSet(object.state) ? eventStateFromJSON(object.state) : 0,
    };
  },

  toJSON(message: SendKeyboardEventRequest): unknown {
    const obj: any = {};
    if (message.code !== 0) {
      obj.code = inputEventToJSON(message.code);
    }
    if (message.state !== 0) {
      obj.state = eventStateToJSON(message.state);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SendKeyboardEventRequest>, I>>(base?: I): SendKeyboardEventRequest {
    return SendKeyboardEventRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SendKeyboardEventRequest>, I>>(object: I): SendKeyboardEventRequest {
    const message = createBaseSendKeyboardEventRequest();
    message.code = object.code ?? 0;
    message.state = object.state ?? 0;
    return message;
  },
};

function createBaseMovePosition(): MovePosition {
  return { x: 0, y: 0 };
}

export const MovePosition = {
  encode(message: MovePosition, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.x !== 0) {
      writer.uint32(8).int32(message.x);
    }
    if (message.y !== 0) {
      writer.uint32(16).int32(message.y);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MovePosition {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMovePosition();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.x = reader.int32();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.y = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MovePosition {
    return { x: isSet(object.x) ? Number(object.x) : 0, y: isSet(object.y) ? Number(object.y) : 0 };
  },

  toJSON(message: MovePosition): unknown {
    const obj: any = {};
    if (message.x !== 0) {
      obj.x = Math.round(message.x);
    }
    if (message.y !== 0) {
      obj.y = Math.round(message.y);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MovePosition>, I>>(base?: I): MovePosition {
    return MovePosition.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<MovePosition>, I>>(object: I): MovePosition {
    const message = createBaseMovePosition();
    message.x = object.x ?? 0;
    message.y = object.y ?? 0;
    return message;
  },
};

function createBaseSendTouchpadEventRequest(): SendTouchpadEventRequest {
  return { event: 0, position: undefined };
}

export const SendTouchpadEventRequest = {
  encode(message: SendTouchpadEventRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.event !== 0) {
      writer.uint32(8).int32(message.event);
    }
    if (message.position !== undefined) {
      MovePosition.encode(message.position, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendTouchpadEventRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendTouchpadEventRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.event = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.position = MovePosition.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SendTouchpadEventRequest {
    return {
      event: isSet(object.event) ? touchEventFromJSON(object.event) : 0,
      position: isSet(object.position) ? MovePosition.fromJSON(object.position) : undefined,
    };
  },

  toJSON(message: SendTouchpadEventRequest): unknown {
    const obj: any = {};
    if (message.event !== 0) {
      obj.event = touchEventToJSON(message.event);
    }
    if (message.position !== undefined) {
      obj.position = MovePosition.toJSON(message.position);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SendTouchpadEventRequest>, I>>(base?: I): SendTouchpadEventRequest {
    return SendTouchpadEventRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SendTouchpadEventRequest>, I>>(object: I): SendTouchpadEventRequest {
    const message = createBaseSendTouchpadEventRequest();
    message.event = object.event ?? 0;
    message.position = (object.position !== undefined && object.position !== null)
      ? MovePosition.fromPartial(object.position)
      : undefined;
    return message;
  },
};

function createBaseSendTouchpadMoveRequest(): SendTouchpadMoveRequest {
  return { position: undefined };
}

export const SendTouchpadMoveRequest = {
  encode(message: SendTouchpadMoveRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.position !== undefined) {
      MovePosition.encode(message.position, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendTouchpadMoveRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendTouchpadMoveRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.position = MovePosition.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SendTouchpadMoveRequest {
    return { position: isSet(object.position) ? MovePosition.fromJSON(object.position) : undefined };
  },

  toJSON(message: SendTouchpadMoveRequest): unknown {
    const obj: any = {};
    if (message.position !== undefined) {
      obj.position = MovePosition.toJSON(message.position);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SendTouchpadMoveRequest>, I>>(base?: I): SendTouchpadMoveRequest {
    return SendTouchpadMoveRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SendTouchpadMoveRequest>, I>>(object: I): SendTouchpadMoveRequest {
    const message = createBaseSendTouchpadMoveRequest();
    message.position = (object.position !== undefined && object.position !== null)
      ? MovePosition.fromPartial(object.position)
      : undefined;
    return message;
  },
};

function createBaseSendMouseMoveByDirectionRequest(): SendMouseMoveByDirectionRequest {
  return { pixel: 0 };
}

export const SendMouseMoveByDirectionRequest = {
  encode(message: SendMouseMoveByDirectionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pixel !== 0) {
      writer.uint32(8).int32(message.pixel);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendMouseMoveByDirectionRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendMouseMoveByDirectionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.pixel = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SendMouseMoveByDirectionRequest {
    return { pixel: isSet(object.pixel) ? Number(object.pixel) : 0 };
  },

  toJSON(message: SendMouseMoveByDirectionRequest): unknown {
    const obj: any = {};
    if (message.pixel !== 0) {
      obj.pixel = Math.round(message.pixel);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SendMouseMoveByDirectionRequest>, I>>(base?: I): SendMouseMoveByDirectionRequest {
    return SendMouseMoveByDirectionRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SendMouseMoveByDirectionRequest>, I>>(
    object: I,
  ): SendMouseMoveByDirectionRequest {
    const message = createBaseSendMouseMoveByDirectionRequest();
    message.pixel = object.pixel ?? 0;
    return message;
  },
};

function createBaseSendMouseMoveRequest(): SendMouseMoveRequest {
  return { x: 0, y: 0 };
}

export const SendMouseMoveRequest = {
  encode(message: SendMouseMoveRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.x !== 0) {
      writer.uint32(8).int32(message.x);
    }
    if (message.y !== 0) {
      writer.uint32(16).int32(message.y);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendMouseMoveRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendMouseMoveRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.x = reader.int32();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.y = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SendMouseMoveRequest {
    return { x: isSet(object.x) ? Number(object.x) : 0, y: isSet(object.y) ? Number(object.y) : 0 };
  },

  toJSON(message: SendMouseMoveRequest): unknown {
    const obj: any = {};
    if (message.x !== 0) {
      obj.x = Math.round(message.x);
    }
    if (message.y !== 0) {
      obj.y = Math.round(message.y);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SendMouseMoveRequest>, I>>(base?: I): SendMouseMoveRequest {
    return SendMouseMoveRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SendMouseMoveRequest>, I>>(object: I): SendMouseMoveRequest {
    const message = createBaseSendMouseMoveRequest();
    message.x = object.x ?? 0;
    message.y = object.y ?? 0;
    return message;
  },
};

function createBaseSendMouseWheelRequest(): SendMouseWheelRequest {
  return { horizontal: false, delta: 0 };
}

export const SendMouseWheelRequest = {
  encode(message: SendMouseWheelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.horizontal === true) {
      writer.uint32(8).bool(message.horizontal);
    }
    if (message.delta !== 0) {
      writer.uint32(21).float(message.delta);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SendMouseWheelRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSendMouseWheelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.horizontal = reader.bool();
          continue;
        case 2:
          if (tag !== 21) {
            break;
          }

          message.delta = reader.float();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SendMouseWheelRequest {
    return {
      horizontal: isSet(object.horizontal) ? Boolean(object.horizontal) : false,
      delta: isSet(object.delta) ? Number(object.delta) : 0,
    };
  },

  toJSON(message: SendMouseWheelRequest): unknown {
    const obj: any = {};
    if (message.horizontal === true) {
      obj.horizontal = message.horizontal;
    }
    if (message.delta !== 0) {
      obj.delta = message.delta;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SendMouseWheelRequest>, I>>(base?: I): SendMouseWheelRequest {
    return SendMouseWheelRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SendMouseWheelRequest>, I>>(object: I): SendMouseWheelRequest {
    const message = createBaseSendMouseWheelRequest();
    message.horizontal = object.horizontal ?? false;
    message.delta = object.delta ?? 0;
    return message;
  },
};

function createBaseSetRemoteScreenRectRequest(): SetRemoteScreenRectRequest {
  return { height: 0, width: 0 };
}

export const SetRemoteScreenRectRequest = {
  encode(message: SetRemoteScreenRectRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.height !== 0) {
      writer.uint32(8).int32(message.height);
    }
    if (message.width !== 0) {
      writer.uint32(16).int32(message.width);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SetRemoteScreenRectRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetRemoteScreenRectRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.height = reader.int32();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.width = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SetRemoteScreenRectRequest {
    return {
      height: isSet(object.height) ? Number(object.height) : 0,
      width: isSet(object.width) ? Number(object.width) : 0,
    };
  },

  toJSON(message: SetRemoteScreenRectRequest): unknown {
    const obj: any = {};
    if (message.height !== 0) {
      obj.height = Math.round(message.height);
    }
    if (message.width !== 0) {
      obj.width = Math.round(message.width);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SetRemoteScreenRectRequest>, I>>(base?: I): SetRemoteScreenRectRequest {
    return SetRemoteScreenRectRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SetRemoteScreenRectRequest>, I>>(object: I): SetRemoteScreenRectRequest {
    const message = createBaseSetRemoteScreenRectRequest();
    message.height = object.height ?? 0;
    message.width = object.width ?? 0;
    return message;
  },
};

function createBaseWriteClipboardRequest(): WriteClipboardRequest {
  return { text: "" };
}

export const WriteClipboardRequest = {
  encode(message: WriteClipboardRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.text !== "") {
      writer.uint32(10).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WriteClipboardRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWriteClipboardRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.text = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): WriteClipboardRequest {
    return { text: isSet(object.text) ? String(object.text) : "" };
  },

  toJSON(message: WriteClipboardRequest): unknown {
    const obj: any = {};
    if (message.text !== "") {
      obj.text = message.text;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<WriteClipboardRequest>, I>>(base?: I): WriteClipboardRequest {
    return WriteClipboardRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<WriteClipboardRequest>, I>>(object: I): WriteClipboardRequest {
    const message = createBaseWriteClipboardRequest();
    message.text = object.text ?? "";
    return message;
  },
};

function createBaseReadClipboardResponse(): ReadClipboardResponse {
  return { text: "" };
}

export const ReadClipboardResponse = {
  encode(message: ReadClipboardResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.text !== "") {
      writer.uint32(10).string(message.text);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ReadClipboardResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseReadClipboardResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.text = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ReadClipboardResponse {
    return { text: isSet(object.text) ? String(object.text) : "" };
  },

  toJSON(message: ReadClipboardResponse): unknown {
    const obj: any = {};
    if (message.text !== "") {
      obj.text = message.text;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ReadClipboardResponse>, I>>(base?: I): ReadClipboardResponse {
    return ReadClipboardResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ReadClipboardResponse>, I>>(object: I): ReadClipboardResponse {
    const message = createBaseReadClipboardResponse();
    message.text = object.text ?? "";
    return message;
  },
};

function createBaseListSinkInputsResponse(): ListSinkInputsResponse {
  return { sinkInputs: [] };
}

export const ListSinkInputsResponse = {
  encode(message: ListSinkInputsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.sinkInputs) {
      ListSinkInputsResponse_SinkInput.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListSinkInputsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListSinkInputsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.sinkInputs.push(ListSinkInputsResponse_SinkInput.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListSinkInputsResponse {
    return {
      sinkInputs: Array.isArray(object?.sinkInputs)
        ? object.sinkInputs.map((e: any) => ListSinkInputsResponse_SinkInput.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ListSinkInputsResponse): unknown {
    const obj: any = {};
    if (message.sinkInputs?.length) {
      obj.sinkInputs = message.sinkInputs.map((e) => ListSinkInputsResponse_SinkInput.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListSinkInputsResponse>, I>>(base?: I): ListSinkInputsResponse {
    return ListSinkInputsResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListSinkInputsResponse>, I>>(object: I): ListSinkInputsResponse {
    const message = createBaseListSinkInputsResponse();
    message.sinkInputs = object.sinkInputs?.map((e) => ListSinkInputsResponse_SinkInput.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListSinkInputsResponse_SinkInput(): ListSinkInputsResponse_SinkInput {
  return { index: 0, name: "", sink: 0, muted: false, volume: 0 };
}

export const ListSinkInputsResponse_SinkInput = {
  encode(message: ListSinkInputsResponse_SinkInput, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== 0) {
      writer.uint32(8).int32(message.index);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.sink !== 0) {
      writer.uint32(24).int32(message.sink);
    }
    if (message.muted === true) {
      writer.uint32(32).bool(message.muted);
    }
    if (message.volume !== 0) {
      writer.uint32(45).float(message.volume);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListSinkInputsResponse_SinkInput {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListSinkInputsResponse_SinkInput();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.index = reader.int32();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.sink = reader.int32();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.muted = reader.bool();
          continue;
        case 5:
          if (tag !== 45) {
            break;
          }

          message.volume = reader.float();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListSinkInputsResponse_SinkInput {
    return {
      index: isSet(object.index) ? Number(object.index) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      sink: isSet(object.sink) ? Number(object.sink) : 0,
      muted: isSet(object.muted) ? Boolean(object.muted) : false,
      volume: isSet(object.volume) ? Number(object.volume) : 0,
    };
  },

  toJSON(message: ListSinkInputsResponse_SinkInput): unknown {
    const obj: any = {};
    if (message.index !== 0) {
      obj.index = Math.round(message.index);
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.sink !== 0) {
      obj.sink = Math.round(message.sink);
    }
    if (message.muted === true) {
      obj.muted = message.muted;
    }
    if (message.volume !== 0) {
      obj.volume = message.volume;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListSinkInputsResponse_SinkInput>, I>>(
    base?: I,
  ): ListSinkInputsResponse_SinkInput {
    return ListSinkInputsResponse_SinkInput.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListSinkInputsResponse_SinkInput>, I>>(
    object: I,
  ): ListSinkInputsResponse_SinkInput {
    const message = createBaseListSinkInputsResponse_SinkInput();
    message.index = object.index ?? 0;
    message.name = object.name ?? "";
    message.sink = object.sink ?? 0;
    message.muted = object.muted ?? false;
    message.volume = object.volume ?? 0;
    return message;
  },
};

function createBaseListSinksRepsonse(): ListSinksRepsonse {
  return { sinks: [] };
}

export const ListSinksRepsonse = {
  encode(message: ListSinksRepsonse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.sinks) {
      ListSinksRepsonse_Sink.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListSinksRepsonse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListSinksRepsonse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.sinks.push(ListSinksRepsonse_Sink.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListSinksRepsonse {
    return {
      sinks: Array.isArray(object?.sinks) ? object.sinks.map((e: any) => ListSinksRepsonse_Sink.fromJSON(e)) : [],
    };
  },

  toJSON(message: ListSinksRepsonse): unknown {
    const obj: any = {};
    if (message.sinks?.length) {
      obj.sinks = message.sinks.map((e) => ListSinksRepsonse_Sink.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListSinksRepsonse>, I>>(base?: I): ListSinksRepsonse {
    return ListSinksRepsonse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListSinksRepsonse>, I>>(object: I): ListSinksRepsonse {
    const message = createBaseListSinksRepsonse();
    message.sinks = object.sinks?.map((e) => ListSinksRepsonse_Sink.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListSinksRepsonse_Sink(): ListSinksRepsonse_Sink {
  return { index: 0, name: "", description: "", muted: false, volume: 0, cardIndex: 0 };
}

export const ListSinksRepsonse_Sink = {
  encode(message: ListSinksRepsonse_Sink, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== 0) {
      writer.uint32(8).int32(message.index);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.muted === true) {
      writer.uint32(32).bool(message.muted);
    }
    if (message.volume !== 0) {
      writer.uint32(45).float(message.volume);
    }
    if (message.cardIndex !== 0) {
      writer.uint32(48).int32(message.cardIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListSinksRepsonse_Sink {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListSinksRepsonse_Sink();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.index = reader.int32();
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

          message.description = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.muted = reader.bool();
          continue;
        case 5:
          if (tag !== 45) {
            break;
          }

          message.volume = reader.float();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.cardIndex = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListSinksRepsonse_Sink {
    return {
      index: isSet(object.index) ? Number(object.index) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      muted: isSet(object.muted) ? Boolean(object.muted) : false,
      volume: isSet(object.volume) ? Number(object.volume) : 0,
      cardIndex: isSet(object.cardIndex) ? Number(object.cardIndex) : 0,
    };
  },

  toJSON(message: ListSinksRepsonse_Sink): unknown {
    const obj: any = {};
    if (message.index !== 0) {
      obj.index = Math.round(message.index);
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.muted === true) {
      obj.muted = message.muted;
    }
    if (message.volume !== 0) {
      obj.volume = message.volume;
    }
    if (message.cardIndex !== 0) {
      obj.cardIndex = Math.round(message.cardIndex);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListSinksRepsonse_Sink>, I>>(base?: I): ListSinksRepsonse_Sink {
    return ListSinksRepsonse_Sink.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListSinksRepsonse_Sink>, I>>(object: I): ListSinksRepsonse_Sink {
    const message = createBaseListSinksRepsonse_Sink();
    message.index = object.index ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.muted = object.muted ?? false;
    message.volume = object.volume ?? 0;
    message.cardIndex = object.cardIndex ?? 0;
    return message;
  },
};

function createBaseCardProfile(): CardProfile {
  return { name: "", description: "", available: false, priority: 0, sinks: 0, sources: 0 };
}

export const CardProfile = {
  encode(message: CardProfile, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    if (message.available === true) {
      writer.uint32(24).bool(message.available);
    }
    if (message.priority !== 0) {
      writer.uint32(32).int32(message.priority);
    }
    if (message.sinks !== 0) {
      writer.uint32(40).int32(message.sinks);
    }
    if (message.sources !== 0) {
      writer.uint32(48).int32(message.sources);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CardProfile {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCardProfile();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.description = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.available = reader.bool();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.priority = reader.int32();
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.sinks = reader.int32();
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.sources = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CardProfile {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      available: isSet(object.available) ? Boolean(object.available) : false,
      priority: isSet(object.priority) ? Number(object.priority) : 0,
      sinks: isSet(object.sinks) ? Number(object.sinks) : 0,
      sources: isSet(object.sources) ? Number(object.sources) : 0,
    };
  },

  toJSON(message: CardProfile): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.available === true) {
      obj.available = message.available;
    }
    if (message.priority !== 0) {
      obj.priority = Math.round(message.priority);
    }
    if (message.sinks !== 0) {
      obj.sinks = Math.round(message.sinks);
    }
    if (message.sources !== 0) {
      obj.sources = Math.round(message.sources);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CardProfile>, I>>(base?: I): CardProfile {
    return CardProfile.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CardProfile>, I>>(object: I): CardProfile {
    const message = createBaseCardProfile();
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.available = object.available ?? false;
    message.priority = object.priority ?? 0;
    message.sinks = object.sinks ?? 0;
    message.sources = object.sources ?? 0;
    return message;
  },
};

function createBaseListCardsResponse(): ListCardsResponse {
  return { cards: [] };
}

export const ListCardsResponse = {
  encode(message: ListCardsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.cards) {
      ListCardsResponse_Card.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListCardsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListCardsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.cards.push(ListCardsResponse_Card.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListCardsResponse {
    return {
      cards: Array.isArray(object?.cards) ? object.cards.map((e: any) => ListCardsResponse_Card.fromJSON(e)) : [],
    };
  },

  toJSON(message: ListCardsResponse): unknown {
    const obj: any = {};
    if (message.cards?.length) {
      obj.cards = message.cards.map((e) => ListCardsResponse_Card.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListCardsResponse>, I>>(base?: I): ListCardsResponse {
    return ListCardsResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListCardsResponse>, I>>(object: I): ListCardsResponse {
    const message = createBaseListCardsResponse();
    message.cards = object.cards?.map((e) => ListCardsResponse_Card.fromPartial(e)) || [];
    return message;
  },
};

function createBaseListCardsResponse_Card(): ListCardsResponse_Card {
  return { index: 0, name: "", deviceName: "", activeProfile: undefined, profiles: [] };
}

export const ListCardsResponse_Card = {
  encode(message: ListCardsResponse_Card, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== 0) {
      writer.uint32(8).int32(message.index);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.deviceName !== "") {
      writer.uint32(26).string(message.deviceName);
    }
    if (message.activeProfile !== undefined) {
      CardProfile.encode(message.activeProfile, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.profiles) {
      CardProfile.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListCardsResponse_Card {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListCardsResponse_Card();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.index = reader.int32();
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

          message.deviceName = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.activeProfile = CardProfile.decode(reader, reader.uint32());
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.profiles.push(CardProfile.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListCardsResponse_Card {
    return {
      index: isSet(object.index) ? Number(object.index) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
      activeProfile: isSet(object.activeProfile) ? CardProfile.fromJSON(object.activeProfile) : undefined,
      profiles: Array.isArray(object?.profiles) ? object.profiles.map((e: any) => CardProfile.fromJSON(e)) : [],
    };
  },

  toJSON(message: ListCardsResponse_Card): unknown {
    const obj: any = {};
    if (message.index !== 0) {
      obj.index = Math.round(message.index);
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.deviceName !== "") {
      obj.deviceName = message.deviceName;
    }
    if (message.activeProfile !== undefined) {
      obj.activeProfile = CardProfile.toJSON(message.activeProfile);
    }
    if (message.profiles?.length) {
      obj.profiles = message.profiles.map((e) => CardProfile.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListCardsResponse_Card>, I>>(base?: I): ListCardsResponse_Card {
    return ListCardsResponse_Card.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ListCardsResponse_Card>, I>>(object: I): ListCardsResponse_Card {
    const message = createBaseListCardsResponse_Card();
    message.index = object.index ?? 0;
    message.name = object.name ?? "";
    message.deviceName = object.deviceName ?? "";
    message.activeProfile = (object.activeProfile !== undefined && object.activeProfile !== null)
      ? CardProfile.fromPartial(object.activeProfile)
      : undefined;
    message.profiles = object.profiles?.map((e) => CardProfile.fromPartial(e)) || [];
    return message;
  },
};

function createBaseSetDefaultSinkRequest(): SetDefaultSinkRequest {
  return { sinkName: "" };
}

export const SetDefaultSinkRequest = {
  encode(message: SetDefaultSinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sinkName !== "") {
      writer.uint32(10).string(message.sinkName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SetDefaultSinkRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetDefaultSinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.sinkName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SetDefaultSinkRequest {
    return { sinkName: isSet(object.sinkName) ? String(object.sinkName) : "" };
  },

  toJSON(message: SetDefaultSinkRequest): unknown {
    const obj: any = {};
    if (message.sinkName !== "") {
      obj.sinkName = message.sinkName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SetDefaultSinkRequest>, I>>(base?: I): SetDefaultSinkRequest {
    return SetDefaultSinkRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SetDefaultSinkRequest>, I>>(object: I): SetDefaultSinkRequest {
    const message = createBaseSetDefaultSinkRequest();
    message.sinkName = object.sinkName ?? "";
    return message;
  },
};

function createBaseSetCardProfileRequest(): SetCardProfileRequest {
  return { cardIndex: 0, profileName: "" };
}

export const SetCardProfileRequest = {
  encode(message: SetCardProfileRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cardIndex !== 0) {
      writer.uint32(8).int32(message.cardIndex);
    }
    if (message.profileName !== "") {
      writer.uint32(18).string(message.profileName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SetCardProfileRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetCardProfileRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.cardIndex = reader.int32();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.profileName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SetCardProfileRequest {
    return {
      cardIndex: isSet(object.cardIndex) ? Number(object.cardIndex) : 0,
      profileName: isSet(object.profileName) ? String(object.profileName) : "",
    };
  },

  toJSON(message: SetCardProfileRequest): unknown {
    const obj: any = {};
    if (message.cardIndex !== 0) {
      obj.cardIndex = Math.round(message.cardIndex);
    }
    if (message.profileName !== "") {
      obj.profileName = message.profileName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SetCardProfileRequest>, I>>(base?: I): SetCardProfileRequest {
    return SetCardProfileRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SetCardProfileRequest>, I>>(object: I): SetCardProfileRequest {
    const message = createBaseSetCardProfileRequest();
    message.cardIndex = object.cardIndex ?? 0;
    message.profileName = object.profileName ?? "";
    return message;
  },
};

function createBaseSinkRequest(): SinkRequest {
  return { sinkName: undefined };
}

export const SinkRequest = {
  encode(message: SinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sinkName !== undefined) {
      writer.uint32(10).string(message.sinkName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SinkRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.sinkName = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SinkRequest {
    return { sinkName: isSet(object.sinkName) ? String(object.sinkName) : undefined };
  },

  toJSON(message: SinkRequest): unknown {
    const obj: any = {};
    if (message.sinkName !== undefined) {
      obj.sinkName = message.sinkName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SinkRequest>, I>>(base?: I): SinkRequest {
    return SinkRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SinkRequest>, I>>(object: I): SinkRequest {
    const message = createBaseSinkRequest();
    message.sinkName = object.sinkName ?? undefined;
    return message;
  },
};

function createBaseGetMuteResponse(): GetMuteResponse {
  return { muted: false };
}

export const GetMuteResponse = {
  encode(message: GetMuteResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.muted === true) {
      writer.uint32(8).bool(message.muted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetMuteResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetMuteResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.muted = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetMuteResponse {
    return { muted: isSet(object.muted) ? Boolean(object.muted) : false };
  },

  toJSON(message: GetMuteResponse): unknown {
    const obj: any = {};
    if (message.muted === true) {
      obj.muted = message.muted;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetMuteResponse>, I>>(base?: I): GetMuteResponse {
    return GetMuteResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetMuteResponse>, I>>(object: I): GetMuteResponse {
    const message = createBaseGetMuteResponse();
    message.muted = object.muted ?? false;
    return message;
  },
};

function createBaseSetSinkMuteRequest(): SetSinkMuteRequest {
  return { sinkName: undefined, mute: false };
}

export const SetSinkMuteRequest = {
  encode(message: SetSinkMuteRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sinkName !== undefined) {
      writer.uint32(10).string(message.sinkName);
    }
    if (message.mute === true) {
      writer.uint32(16).bool(message.mute);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SetSinkMuteRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetSinkMuteRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.sinkName = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.mute = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SetSinkMuteRequest {
    return {
      sinkName: isSet(object.sinkName) ? String(object.sinkName) : undefined,
      mute: isSet(object.mute) ? Boolean(object.mute) : false,
    };
  },

  toJSON(message: SetSinkMuteRequest): unknown {
    const obj: any = {};
    if (message.sinkName !== undefined) {
      obj.sinkName = message.sinkName;
    }
    if (message.mute === true) {
      obj.mute = message.mute;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SetSinkMuteRequest>, I>>(base?: I): SetSinkMuteRequest {
    return SetSinkMuteRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SetSinkMuteRequest>, I>>(object: I): SetSinkMuteRequest {
    const message = createBaseSetSinkMuteRequest();
    message.sinkName = object.sinkName ?? undefined;
    message.mute = object.mute ?? false;
    return message;
  },
};

function createBaseGetSinkVolumeResponse(): GetSinkVolumeResponse {
  return { volume: 0 };
}

export const GetSinkVolumeResponse = {
  encode(message: GetSinkVolumeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.volume !== 0) {
      writer.uint32(13).float(message.volume);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetSinkVolumeResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetSinkVolumeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 13) {
            break;
          }

          message.volume = reader.float();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetSinkVolumeResponse {
    return { volume: isSet(object.volume) ? Number(object.volume) : 0 };
  },

  toJSON(message: GetSinkVolumeResponse): unknown {
    const obj: any = {};
    if (message.volume !== 0) {
      obj.volume = message.volume;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetSinkVolumeResponse>, I>>(base?: I): GetSinkVolumeResponse {
    return GetSinkVolumeResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<GetSinkVolumeResponse>, I>>(object: I): GetSinkVolumeResponse {
    const message = createBaseGetSinkVolumeResponse();
    message.volume = object.volume ?? 0;
    return message;
  },
};

function createBaseChangeVolumeRequest(): ChangeVolumeRequest {
  return { value: 0 };
}

export const ChangeVolumeRequest = {
  encode(message: ChangeVolumeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.value !== 0) {
      writer.uint32(13).float(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChangeVolumeRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChangeVolumeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 13) {
            break;
          }

          message.value = reader.float();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ChangeVolumeRequest {
    return { value: isSet(object.value) ? Number(object.value) : 0 };
  },

  toJSON(message: ChangeVolumeRequest): unknown {
    const obj: any = {};
    if (message.value !== 0) {
      obj.value = message.value;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ChangeVolumeRequest>, I>>(base?: I): ChangeVolumeRequest {
    return ChangeVolumeRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<ChangeVolumeRequest>, I>>(object: I): ChangeVolumeRequest {
    const message = createBaseChangeVolumeRequest();
    message.value = object.value ?? 0;
    return message;
  },
};

function createBaseSetSinkInputVolumeRequest(): SetSinkInputVolumeRequest {
  return { sinkInput: "", volume: 0 };
}

export const SetSinkInputVolumeRequest = {
  encode(message: SetSinkInputVolumeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sinkInput !== "") {
      writer.uint32(10).string(message.sinkInput);
    }
    if (message.volume !== 0) {
      writer.uint32(21).float(message.volume);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SetSinkInputVolumeRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetSinkInputVolumeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.sinkInput = reader.string();
          continue;
        case 2:
          if (tag !== 21) {
            break;
          }

          message.volume = reader.float();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SetSinkInputVolumeRequest {
    return {
      sinkInput: isSet(object.sinkInput) ? String(object.sinkInput) : "",
      volume: isSet(object.volume) ? Number(object.volume) : 0,
    };
  },

  toJSON(message: SetSinkInputVolumeRequest): unknown {
    const obj: any = {};
    if (message.sinkInput !== "") {
      obj.sinkInput = message.sinkInput;
    }
    if (message.volume !== 0) {
      obj.volume = message.volume;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SetSinkInputVolumeRequest>, I>>(base?: I): SetSinkInputVolumeRequest {
    return SetSinkInputVolumeRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<SetSinkInputVolumeRequest>, I>>(object: I): SetSinkInputVolumeRequest {
    const message = createBaseSetSinkInputVolumeRequest();
    message.sinkInput = object.sinkInput ?? "";
    message.volume = object.volume ?? 0;
    return message;
  },
};

function createBaseBrowserActionRequest(): BrowserActionRequest {
  return { action: 0 };
}

export const BrowserActionRequest = {
  encode(message: BrowserActionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.action !== 0) {
      writer.uint32(8).int32(message.action);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BrowserActionRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBrowserActionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.action = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): BrowserActionRequest {
    return { action: isSet(object.action) ? browserActionRequest_ActionFromJSON(object.action) : 0 };
  },

  toJSON(message: BrowserActionRequest): unknown {
    const obj: any = {};
    if (message.action !== 0) {
      obj.action = browserActionRequest_ActionToJSON(message.action);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<BrowserActionRequest>, I>>(base?: I): BrowserActionRequest {
    return BrowserActionRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<BrowserActionRequest>, I>>(object: I): BrowserActionRequest {
    const message = createBaseBrowserActionRequest();
    message.action = object.action ?? 0;
    return message;
  },
};

export interface RemoteControl {
  /** 发送键盘输入事件 */
  SendKeyboardEvent(
    request: DeepPartial<SendKeyboardEventRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /**
   * 发送触摸板输入事件
   * 此接口发送的事件都是源事件，比如发送LeftPress，只会发起按下，不会发起松开
   * 如果是想要单击等操作使用其他接口
   */
  SendTouchpadEvent(
    request: DeepPartial<SendTouchpadEventRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 发送左键 */
  SendTouchpadClick(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<Empty>;
  /** 发送右键 */
  SendTouchpadRightClick(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 发送双击 */
  SendTouchpadDoubleClick(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 发送触摸板滑动 */
  SendTouchpadMove(
    request: DeepPartial<SendTouchpadMoveRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 发起鼠标向左移动 */
  SendMouseMoveLeft(
    request: DeepPartial<SendMouseMoveByDirectionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 发起鼠标向右移动 */
  SendMouseMoveRight(
    request: DeepPartial<SendMouseMoveByDirectionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 发起鼠标向上移动 */
  SendMouseMoveUp(
    request: DeepPartial<SendMouseMoveByDirectionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 发起鼠标向下移动 */
  SendMouseMove(
    request: DeepPartial<SendMouseMoveRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 建立鼠标移动客户端流 */
  MouseMoveStream(
    request: Observable<DeepPartial<SendMouseMoveRequest>>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 发起鼠标左键单击事件 */
  SendMouseLeftClick(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<Empty>;
  /** 发起鼠标右键单击事件 */
  SendMouseRightClick(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<Empty>;
  /** 发起鼠标中键单击事件 */
  SendMouseMiddleClick(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 发起鼠标滚动 */
  SendMouseWheel(
    request: DeepPartial<SendMouseWheelRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 建立流式鼠标滑动 */
  MouseWheelStream(
    request: Observable<DeepPartial<SendMouseWheelRequest>>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 发起鼠标双击左键事件 */
  SendMouseDoubleClick(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 设置远程屏幕的宽高(用于计算鼠标，触控板移动的同比例偏移) */
  SetRemoteScreenRect(
    request: DeepPartial<SetRemoteScreenRectRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 写入剪贴板 */
  WriteClipboard(
    request: DeepPartial<WriteClipboardRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 读取剪贴板 */
  ReadClipboard(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ReadClipboardResponse>;
  /** 粘贴 */
  DoPaste(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<Empty>;
  /** 浏览器操作 */
  BrowserAction(
    request: DeepPartial<BrowserActionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /**
   * 音频管理
   * sink 输入音频设备 (可用于播放音频的音频设备)
   * sink-input 输入音频流 (正在播放音频的程序)
   * card 声卡，对应到某个物理设备
   * card profile 声音可用的配置(立体声、环绕声)，影响sink和source
   * 列出正在播放的输出音频流
   */
  ListSinkInputs(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListSinkInputsResponse>;
  /** 列出当前输入音频设备 */
  ListSinks(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListSinksRepsonse>;
  /** 列出当前可用的声卡 */
  ListCards(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListCardsResponse>;
  /** 设置默认的输出音频设备 */
  SetDefaultSink(
    request: DeepPartial<SetDefaultSinkRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 设置默认的输入音频设备 */
  SetCardProfile(
    request: DeepPartial<SetCardProfileRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 获取是否被静音 */
  GetMute(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<GetMuteResponse>;
  /** 切换静音状态 */
  ToggleMute(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<GetMuteResponse>;
  /** 获取音量 */
  GetVolume(
    request: DeepPartial<SinkRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<GetSinkVolumeResponse>;
  /** 增加音量 */
  IncreaseVolume(
    request: DeepPartial<ChangeVolumeRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 减少音量 */
  DecreaseVolume(
    request: DeepPartial<ChangeVolumeRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
  /** 设置输入音频源音量 */
  SetSinkInputVolume(
    request: DeepPartial<SetSinkInputVolumeRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty>;
}

export class RemoteControlClientImpl implements RemoteControl {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.SendKeyboardEvent = this.SendKeyboardEvent.bind(this);
    this.SendTouchpadEvent = this.SendTouchpadEvent.bind(this);
    this.SendTouchpadClick = this.SendTouchpadClick.bind(this);
    this.SendTouchpadRightClick = this.SendTouchpadRightClick.bind(this);
    this.SendTouchpadDoubleClick = this.SendTouchpadDoubleClick.bind(this);
    this.SendTouchpadMove = this.SendTouchpadMove.bind(this);
    this.SendMouseMoveLeft = this.SendMouseMoveLeft.bind(this);
    this.SendMouseMoveRight = this.SendMouseMoveRight.bind(this);
    this.SendMouseMoveUp = this.SendMouseMoveUp.bind(this);
    this.SendMouseMove = this.SendMouseMove.bind(this);
    this.MouseMoveStream = this.MouseMoveStream.bind(this);
    this.SendMouseLeftClick = this.SendMouseLeftClick.bind(this);
    this.SendMouseRightClick = this.SendMouseRightClick.bind(this);
    this.SendMouseMiddleClick = this.SendMouseMiddleClick.bind(this);
    this.SendMouseWheel = this.SendMouseWheel.bind(this);
    this.MouseWheelStream = this.MouseWheelStream.bind(this);
    this.SendMouseDoubleClick = this.SendMouseDoubleClick.bind(this);
    this.SetRemoteScreenRect = this.SetRemoteScreenRect.bind(this);
    this.WriteClipboard = this.WriteClipboard.bind(this);
    this.ReadClipboard = this.ReadClipboard.bind(this);
    this.DoPaste = this.DoPaste.bind(this);
    this.BrowserAction = this.BrowserAction.bind(this);
    this.ListSinkInputs = this.ListSinkInputs.bind(this);
    this.ListSinks = this.ListSinks.bind(this);
    this.ListCards = this.ListCards.bind(this);
    this.SetDefaultSink = this.SetDefaultSink.bind(this);
    this.SetCardProfile = this.SetCardProfile.bind(this);
    this.GetMute = this.GetMute.bind(this);
    this.ToggleMute = this.ToggleMute.bind(this);
    this.GetVolume = this.GetVolume.bind(this);
    this.IncreaseVolume = this.IncreaseVolume.bind(this);
    this.DecreaseVolume = this.DecreaseVolume.bind(this);
    this.SetSinkInputVolume = this.SetSinkInputVolume.bind(this);
  }

  SendKeyboardEvent(
    request: DeepPartial<SendKeyboardEventRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSendKeyboardEventDesc,
      SendKeyboardEventRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  SendTouchpadEvent(
    request: DeepPartial<SendTouchpadEventRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSendTouchpadEventDesc,
      SendTouchpadEventRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  SendTouchpadClick(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<Empty> {
    return this.rpc.unary(RemoteControlSendTouchpadClickDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  SendTouchpadRightClick(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(RemoteControlSendTouchpadRightClickDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  SendTouchpadDoubleClick(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(RemoteControlSendTouchpadDoubleClickDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  SendTouchpadMove(
    request: DeepPartial<SendTouchpadMoveRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSendTouchpadMoveDesc,
      SendTouchpadMoveRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  SendMouseMoveLeft(
    request: DeepPartial<SendMouseMoveByDirectionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSendMouseMoveLeftDesc,
      SendMouseMoveByDirectionRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  SendMouseMoveRight(
    request: DeepPartial<SendMouseMoveByDirectionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSendMouseMoveRightDesc,
      SendMouseMoveByDirectionRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  SendMouseMoveUp(
    request: DeepPartial<SendMouseMoveByDirectionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSendMouseMoveUpDesc,
      SendMouseMoveByDirectionRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  SendMouseMove(
    request: DeepPartial<SendMouseMoveRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSendMouseMoveDesc,
      SendMouseMoveRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  MouseMoveStream(
    request: Observable<DeepPartial<SendMouseMoveRequest>>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    throw new Error("ts-proto does not yet support client streaming!");
  }

  SendMouseLeftClick(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<Empty> {
    return this.rpc.unary(RemoteControlSendMouseLeftClickDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  SendMouseRightClick(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(RemoteControlSendMouseRightClickDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  SendMouseMiddleClick(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(RemoteControlSendMouseMiddleClickDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  SendMouseWheel(
    request: DeepPartial<SendMouseWheelRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSendMouseWheelDesc,
      SendMouseWheelRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  MouseWheelStream(
    request: Observable<DeepPartial<SendMouseWheelRequest>>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    throw new Error("ts-proto does not yet support client streaming!");
  }

  SendMouseDoubleClick(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(RemoteControlSendMouseDoubleClickDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  SetRemoteScreenRect(
    request: DeepPartial<SetRemoteScreenRectRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSetRemoteScreenRectDesc,
      SetRemoteScreenRectRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  WriteClipboard(
    request: DeepPartial<WriteClipboardRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlWriteClipboardDesc,
      WriteClipboardRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  ReadClipboard(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ReadClipboardResponse> {
    return this.rpc.unary(RemoteControlReadClipboardDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  DoPaste(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<Empty> {
    return this.rpc.unary(RemoteControlDoPasteDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  BrowserAction(
    request: DeepPartial<BrowserActionRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlBrowserActionDesc,
      BrowserActionRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  ListSinkInputs(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListSinkInputsResponse> {
    return this.rpc.unary(RemoteControlListSinkInputsDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  ListSinks(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListSinksRepsonse> {
    return this.rpc.unary(RemoteControlListSinksDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  ListCards(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<ListCardsResponse> {
    return this.rpc.unary(RemoteControlListCardsDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  SetDefaultSink(
    request: DeepPartial<SetDefaultSinkRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSetDefaultSinkDesc,
      SetDefaultSinkRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  SetCardProfile(
    request: DeepPartial<SetCardProfileRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSetCardProfileDesc,
      SetCardProfileRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  GetMute(request: DeepPartial<Empty>, metadata?: grpc.Metadata, abortSignal?: AbortSignal): Promise<GetMuteResponse> {
    return this.rpc.unary(RemoteControlGetMuteDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  ToggleMute(
    request: DeepPartial<Empty>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<GetMuteResponse> {
    return this.rpc.unary(RemoteControlToggleMuteDesc, Empty.fromPartial(request), metadata, abortSignal);
  }

  GetVolume(
    request: DeepPartial<SinkRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<GetSinkVolumeResponse> {
    return this.rpc.unary(RemoteControlGetVolumeDesc, SinkRequest.fromPartial(request), metadata, abortSignal);
  }

  IncreaseVolume(
    request: DeepPartial<ChangeVolumeRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlIncreaseVolumeDesc,
      ChangeVolumeRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  DecreaseVolume(
    request: DeepPartial<ChangeVolumeRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlDecreaseVolumeDesc,
      ChangeVolumeRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }

  SetSinkInputVolume(
    request: DeepPartial<SetSinkInputVolumeRequest>,
    metadata?: grpc.Metadata,
    abortSignal?: AbortSignal,
  ): Promise<Empty> {
    return this.rpc.unary(
      RemoteControlSetSinkInputVolumeDesc,
      SetSinkInputVolumeRequest.fromPartial(request),
      metadata,
      abortSignal,
    );
  }
}

export const RemoteControlDesc = { serviceName: "cloud.lazycat.apis.localdevice.RemoteControl" };

export const RemoteControlSendKeyboardEventDesc: UnaryMethodDefinitionish = {
  methodName: "SendKeyboardEvent",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SendKeyboardEventRequest.encode(this).finish();
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

export const RemoteControlSendTouchpadEventDesc: UnaryMethodDefinitionish = {
  methodName: "SendTouchpadEvent",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SendTouchpadEventRequest.encode(this).finish();
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

export const RemoteControlSendTouchpadClickDesc: UnaryMethodDefinitionish = {
  methodName: "SendTouchpadClick",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
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

export const RemoteControlSendTouchpadRightClickDesc: UnaryMethodDefinitionish = {
  methodName: "SendTouchpadRightClick",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
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

export const RemoteControlSendTouchpadDoubleClickDesc: UnaryMethodDefinitionish = {
  methodName: "SendTouchpadDoubleClick",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
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

export const RemoteControlSendTouchpadMoveDesc: UnaryMethodDefinitionish = {
  methodName: "SendTouchpadMove",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SendTouchpadMoveRequest.encode(this).finish();
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

export const RemoteControlSendMouseMoveLeftDesc: UnaryMethodDefinitionish = {
  methodName: "SendMouseMoveLeft",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SendMouseMoveByDirectionRequest.encode(this).finish();
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

export const RemoteControlSendMouseMoveRightDesc: UnaryMethodDefinitionish = {
  methodName: "SendMouseMoveRight",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SendMouseMoveByDirectionRequest.encode(this).finish();
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

export const RemoteControlSendMouseMoveUpDesc: UnaryMethodDefinitionish = {
  methodName: "SendMouseMoveUp",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SendMouseMoveByDirectionRequest.encode(this).finish();
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

export const RemoteControlSendMouseMoveDesc: UnaryMethodDefinitionish = {
  methodName: "SendMouseMove",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SendMouseMoveRequest.encode(this).finish();
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

export const RemoteControlSendMouseLeftClickDesc: UnaryMethodDefinitionish = {
  methodName: "SendMouseLeftClick",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
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

export const RemoteControlSendMouseRightClickDesc: UnaryMethodDefinitionish = {
  methodName: "SendMouseRightClick",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
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

export const RemoteControlSendMouseMiddleClickDesc: UnaryMethodDefinitionish = {
  methodName: "SendMouseMiddleClick",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
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

export const RemoteControlSendMouseWheelDesc: UnaryMethodDefinitionish = {
  methodName: "SendMouseWheel",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SendMouseWheelRequest.encode(this).finish();
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

export const RemoteControlSendMouseDoubleClickDesc: UnaryMethodDefinitionish = {
  methodName: "SendMouseDoubleClick",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
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

export const RemoteControlSetRemoteScreenRectDesc: UnaryMethodDefinitionish = {
  methodName: "SetRemoteScreenRect",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SetRemoteScreenRectRequest.encode(this).finish();
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

export const RemoteControlWriteClipboardDesc: UnaryMethodDefinitionish = {
  methodName: "WriteClipboard",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return WriteClipboardRequest.encode(this).finish();
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

export const RemoteControlReadClipboardDesc: UnaryMethodDefinitionish = {
  methodName: "ReadClipboard",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ReadClipboardResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RemoteControlDoPasteDesc: UnaryMethodDefinitionish = {
  methodName: "DoPaste",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
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

export const RemoteControlBrowserActionDesc: UnaryMethodDefinitionish = {
  methodName: "BrowserAction",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return BrowserActionRequest.encode(this).finish();
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

export const RemoteControlListSinkInputsDesc: UnaryMethodDefinitionish = {
  methodName: "ListSinkInputs",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ListSinkInputsResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RemoteControlListSinksDesc: UnaryMethodDefinitionish = {
  methodName: "ListSinks",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ListSinksRepsonse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RemoteControlListCardsDesc: UnaryMethodDefinitionish = {
  methodName: "ListCards",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = ListCardsResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RemoteControlSetDefaultSinkDesc: UnaryMethodDefinitionish = {
  methodName: "SetDefaultSink",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SetDefaultSinkRequest.encode(this).finish();
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

export const RemoteControlSetCardProfileDesc: UnaryMethodDefinitionish = {
  methodName: "SetCardProfile",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SetCardProfileRequest.encode(this).finish();
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

export const RemoteControlGetMuteDesc: UnaryMethodDefinitionish = {
  methodName: "GetMute",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = GetMuteResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RemoteControlToggleMuteDesc: UnaryMethodDefinitionish = {
  methodName: "ToggleMute",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return Empty.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = GetMuteResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RemoteControlGetVolumeDesc: UnaryMethodDefinitionish = {
  methodName: "GetVolume",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SinkRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      const value = GetSinkVolumeResponse.decode(data);
      return {
        ...value,
        toObject() {
          return value;
        },
      };
    },
  } as any,
};

export const RemoteControlIncreaseVolumeDesc: UnaryMethodDefinitionish = {
  methodName: "IncreaseVolume",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ChangeVolumeRequest.encode(this).finish();
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

export const RemoteControlDecreaseVolumeDesc: UnaryMethodDefinitionish = {
  methodName: "DecreaseVolume",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ChangeVolumeRequest.encode(this).finish();
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

export const RemoteControlSetSinkInputVolumeDesc: UnaryMethodDefinitionish = {
  methodName: "SetSinkInputVolume",
  service: RemoteControlDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SetSinkInputVolumeRequest.encode(this).finish();
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

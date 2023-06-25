// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: sys/btled.proto

package sys

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ButtonStream_ButtonEvent int32

const (
	// 按钮被按下
	ButtonStream_BUTTON_DOWN ButtonStream_ButtonEvent = 0
	// 按钮被松开
	ButtonStream_BUTTON_UP ButtonStream_ButtonEvent = 1
	// 按钮被点击（按下后被快速松开，两次点击间隔不超过 500 毫秒）
	// 当按钮被点击时，会先产生 BUTTON_DOWN 事件，然后产生 BUTTON_UP 事件，最后产生 BUTTON_CLICK 事件
	// 如果只需要监听按钮被点击事件，可以忽略 BUTTON_DOWN 和 BUTTON_UP 事件
	// 如果需要判断按钮是单击还是双击，需要自行等待 500 毫秒，如果没有收到 BUTTON_DOUBLE_CLICK 事件，则表示是单击
	ButtonStream_BUTTON_CLICK ButtonStream_ButtonEvent = 2
	// 按钮被双击（两次点击间隔不超过 500 毫秒）
	// 当按钮被双击时，会产生两次 BUTTON_CLICK 事件后才产生 BUTTON_DOUBLE_CLICK 事件
	// 如果只需要监听按钮被双击事件，可以忽略 BUTTON_CLICK、BUTTON_DOWN、BUTTON_UP 事件
	ButtonStream_BUTTON_DOUBLE_CLICK ButtonStream_ButtonEvent = 3
	// 按钮被长按（按下时间超过 2 秒还未松开）
	// 按钮被长按的事件过程：会先产生 BUTTON_DOWN 事件，然后产生 BUTTON_LONG_PRESS 事件，等到按钮被松开时，会产生 BUTTON_UP 事件
	ButtonStream_BUTTON_LONG_PRESS ButtonStream_ButtonEvent = 4
	// 按钮被长按（按下时间超过 10 秒还未松开）
	// 按钮被长按超过10秒的事件过程，会先产生 BUTTON_DOWN 事件，然后2秒后产生 BUTTON_LONG_PRESS 事件，然后8秒后产生 BUTTON_LONG_PRESS_10S 事件，等到按钮被松开时，会产生 BUTTON_UP 事件
	ButtonStream_BUTTON_LONG_PRESS_10S ButtonStream_ButtonEvent = 5
)

// Enum value maps for ButtonStream_ButtonEvent.
var (
	ButtonStream_ButtonEvent_name = map[int32]string{
		0: "BUTTON_DOWN",
		1: "BUTTON_UP",
		2: "BUTTON_CLICK",
		3: "BUTTON_DOUBLE_CLICK",
		4: "BUTTON_LONG_PRESS",
		5: "BUTTON_LONG_PRESS_10S",
	}
	ButtonStream_ButtonEvent_value = map[string]int32{
		"BUTTON_DOWN":           0,
		"BUTTON_UP":             1,
		"BUTTON_CLICK":          2,
		"BUTTON_DOUBLE_CLICK":   3,
		"BUTTON_LONG_PRESS":     4,
		"BUTTON_LONG_PRESS_10S": 5,
	}
)

func (x ButtonStream_ButtonEvent) Enum() *ButtonStream_ButtonEvent {
	p := new(ButtonStream_ButtonEvent)
	*p = x
	return p
}

func (x ButtonStream_ButtonEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ButtonStream_ButtonEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_sys_btled_proto_enumTypes[0].Descriptor()
}

func (ButtonStream_ButtonEvent) Type() protoreflect.EnumType {
	return &file_sys_btled_proto_enumTypes[0]
}

func (x ButtonStream_ButtonEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ButtonStream_ButtonEvent.Descriptor instead.
func (ButtonStream_ButtonEvent) EnumDescriptor() ([]byte, []int) {
	return file_sys_btled_proto_rawDescGZIP(), []int{2, 0}
}

type ButtonLedSessionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事务描述，仅用于调试
	Description string `protobuf:"bytes,1,opt,name=Description,proto3" json:"Description,omitempty"`
	// 事务是否可以被抢占
	Preemptable bool `protobuf:"varint,2,opt,name=Preemptable,proto3" json:"Preemptable,omitempty"`
	// 是否需要使用功能按钮（若不使用功能按钮，请不要设置以便节约资源)
	UseButton bool `protobuf:"varint,3,opt,name=UseButton,proto3" json:"UseButton,omitempty"`
}

func (x *ButtonLedSessionInfo) Reset() {
	*x = ButtonLedSessionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_btled_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ButtonLedSessionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ButtonLedSessionInfo) ProtoMessage() {}

func (x *ButtonLedSessionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sys_btled_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ButtonLedSessionInfo.ProtoReflect.Descriptor instead.
func (*ButtonLedSessionInfo) Descriptor() ([]byte, []int) {
	return file_sys_btled_proto_rawDescGZIP(), []int{0}
}

func (x *ButtonLedSessionInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ButtonLedSessionInfo) GetPreemptable() bool {
	if x != nil {
		return x.Preemptable
	}
	return false
}

func (x *ButtonLedSessionInfo) GetUseButton() bool {
	if x != nil {
		return x.UseButton
	}
	return false
}

type LedStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LED 闪烁的亮暗比例（100%/65536）
	//  0 表示一直暗，63356 表示一直亮
	DutyCycle int32 `protobuf:"varint,1,opt,name=DutyCycle,proto3" json:"DutyCycle,omitempty"`
	// LED 闪烁周期（完成一次闪烁的时间，单位：毫秒。设置状态后，按钮会根据所设置的比例先亮后灭）
	Period int32 `protobuf:"varint,2,opt,name=Period,proto3" json:"Period,omitempty"`
	// 事务信息（必须在所有 LED 状态请求之前指定）
	Session *ButtonLedSessionInfo `protobuf:"bytes,3,opt,name=Session,proto3,oneof" json:"Session,omitempty"`
}

func (x *LedStream) Reset() {
	*x = LedStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_btled_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LedStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LedStream) ProtoMessage() {}

func (x *LedStream) ProtoReflect() protoreflect.Message {
	mi := &file_sys_btled_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LedStream.ProtoReflect.Descriptor instead.
func (*LedStream) Descriptor() ([]byte, []int) {
	return file_sys_btled_proto_rawDescGZIP(), []int{1}
}

func (x *LedStream) GetDutyCycle() int32 {
	if x != nil {
		return x.DutyCycle
	}
	return 0
}

func (x *LedStream) GetPeriod() int32 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *LedStream) GetSession() *ButtonLedSessionInfo {
	if x != nil {
		return x.Session
	}
	return nil
}

type ButtonStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 按钮事件
	Event ButtonStream_ButtonEvent `protobuf:"varint,1,opt,name=Event,proto3,enum=cloud.lazycat.apis.sys.ButtonStream_ButtonEvent" json:"Event,omitempty"`
	// 该事务是否已经结束
	//   客户端应该随时检查此消息，如果此值为 true，则说明事务已经被抢占，该事务不再会接收到按钮事件，也不能再发送LED事件
	//   客户端此时唯一的选择是关闭连接，以便释放服务器资源
	Eof bool `protobuf:"varint,2,opt,name=Eof,proto3" json:"Eof,omitempty"`
}

func (x *ButtonStream) Reset() {
	*x = ButtonStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_btled_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ButtonStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ButtonStream) ProtoMessage() {}

func (x *ButtonStream) ProtoReflect() protoreflect.Message {
	mi := &file_sys_btled_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ButtonStream.ProtoReflect.Descriptor instead.
func (*ButtonStream) Descriptor() ([]byte, []int) {
	return file_sys_btled_proto_rawDescGZIP(), []int{2}
}

func (x *ButtonStream) GetEvent() ButtonStream_ButtonEvent {
	if x != nil {
		return x.Event
	}
	return ButtonStream_BUTTON_DOWN
}

func (x *ButtonStream) GetEof() bool {
	if x != nil {
		return x.Eof
	}
	return false
}

var File_sys_btled_proto protoreflect.FileDescriptor

var file_sys_btled_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x79, 0x73, 0x2f, 0x62, 0x74, 0x6c, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x22, 0x78, 0x0a, 0x14, 0x42, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x4c, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x65, 0x6d, 0x70, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x50, 0x72, 0x65, 0x65, 0x6d, 0x70,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x42, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x55, 0x73, 0x65, 0x42, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x09, 0x4c, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x75, 0x74, 0x79, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x44, 0x75, 0x74, 0x79, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x4b, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79,
	0x73, 0x2e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x4c, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0xf5, 0x01, 0x0a, 0x0c, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x46, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x6f, 0x66,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x45, 0x6f, 0x66, 0x22, 0x8a, 0x01, 0x0a, 0x0b,
	0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x42,
	0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x42,
	0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x17, 0x0a,
	0x13, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x5f, 0x43,
	0x4c, 0x49, 0x43, 0x4b, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e,
	0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x53, 0x10, 0x04, 0x12, 0x19, 0x0a,
	0x15, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x5f, 0x50, 0x52, 0x45,
	0x53, 0x53, 0x5f, 0x31, 0x30, 0x53, 0x10, 0x05, 0x32, 0x71, 0x0a, 0x17, 0x42, 0x75, 0x74, 0x74,
	0x6f, 0x6e, 0x4c, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x21,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x4c, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x1a, 0x24, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42, 0x75, 0x74, 0x74, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x28, 0x01, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73,
	0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x79, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_btled_proto_rawDescOnce sync.Once
	file_sys_btled_proto_rawDescData = file_sys_btled_proto_rawDesc
)

func file_sys_btled_proto_rawDescGZIP() []byte {
	file_sys_btled_proto_rawDescOnce.Do(func() {
		file_sys_btled_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_btled_proto_rawDescData)
	})
	return file_sys_btled_proto_rawDescData
}

var file_sys_btled_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sys_btled_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sys_btled_proto_goTypes = []interface{}{
	(ButtonStream_ButtonEvent)(0), // 0: cloud.lazycat.apis.sys.ButtonStream.ButtonEvent
	(*ButtonLedSessionInfo)(nil),  // 1: cloud.lazycat.apis.sys.ButtonLedSessionInfo
	(*LedStream)(nil),             // 2: cloud.lazycat.apis.sys.LedStream
	(*ButtonStream)(nil),          // 3: cloud.lazycat.apis.sys.ButtonStream
}
var file_sys_btled_proto_depIdxs = []int32{
	1, // 0: cloud.lazycat.apis.sys.LedStream.Session:type_name -> cloud.lazycat.apis.sys.ButtonLedSessionInfo
	0, // 1: cloud.lazycat.apis.sys.ButtonStream.Event:type_name -> cloud.lazycat.apis.sys.ButtonStream.ButtonEvent
	2, // 2: cloud.lazycat.apis.sys.ButtonLedSessionService.Connect:input_type -> cloud.lazycat.apis.sys.LedStream
	3, // 3: cloud.lazycat.apis.sys.ButtonLedSessionService.Connect:output_type -> cloud.lazycat.apis.sys.ButtonStream
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sys_btled_proto_init() }
func file_sys_btled_proto_init() {
	if File_sys_btled_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_btled_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ButtonLedSessionInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sys_btled_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LedStream); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sys_btled_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ButtonStream); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_sys_btled_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sys_btled_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_btled_proto_goTypes,
		DependencyIndexes: file_sys_btled_proto_depIdxs,
		EnumInfos:         file_sys_btled_proto_enumTypes,
		MessageInfos:      file_sys_btled_proto_msgTypes,
	}.Build()
	File_sys_btled_proto = out.File
	file_sys_btled_proto_rawDesc = nil
	file_sys_btled_proto_goTypes = nil
	file_sys_btled_proto_depIdxs = nil
}

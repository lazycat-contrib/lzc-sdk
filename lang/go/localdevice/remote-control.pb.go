// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: localdevice/remote-control.proto

package localdevice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InputEventCodes int32

const (
	InputEventCodes_KEY_HOME         InputEventCodes = 0
	InputEventCodes_KEY_BACK         InputEventCodes = 1
	InputEventCodes_KEY_FULL_SCREEN  InputEventCodes = 2
	InputEventCodes_KEY_OPTION       InputEventCodes = 3
	InputEventCodes_BTN_LEFT         InputEventCodes = 4
	InputEventCodes_BTN_RIGHT        InputEventCodes = 5
	InputEventCodes_BTN_MIDDLE       InputEventCodes = 6
	InputEventCodes_REL_X            InputEventCodes = 7
	InputEventCodes_REL_Y            InputEventCodes = 8
	InputEventCodes_REL_WHEEL        InputEventCodes = 9
	InputEventCodes_REL_WHEEL_HI_RES InputEventCodes = 10
	InputEventCodes_KEY_UP           InputEventCodes = 11
	InputEventCodes_KEY_DOWN         InputEventCodes = 12
	InputEventCodes_KEY_LEFT         InputEventCodes = 13
	InputEventCodes_KEY_RIGHT        InputEventCodes = 14
	InputEventCodes_KEY_PAGEUP       InputEventCodes = 15
	InputEventCodes_KEY_PAGEDOWN     InputEventCodes = 16
)

// Enum value maps for InputEventCodes.
var (
	InputEventCodes_name = map[int32]string{
		0:  "KEY_HOME",
		1:  "KEY_BACK",
		2:  "KEY_FULL_SCREEN",
		3:  "KEY_OPTION",
		4:  "BTN_LEFT",
		5:  "BTN_RIGHT",
		6:  "BTN_MIDDLE",
		7:  "REL_X",
		8:  "REL_Y",
		9:  "REL_WHEEL",
		10: "REL_WHEEL_HI_RES",
		11: "KEY_UP",
		12: "KEY_DOWN",
		13: "KEY_LEFT",
		14: "KEY_RIGHT",
		15: "KEY_PAGEUP",
		16: "KEY_PAGEDOWN",
	}
	InputEventCodes_value = map[string]int32{
		"KEY_HOME":         0,
		"KEY_BACK":         1,
		"KEY_FULL_SCREEN":  2,
		"KEY_OPTION":       3,
		"BTN_LEFT":         4,
		"BTN_RIGHT":        5,
		"BTN_MIDDLE":       6,
		"REL_X":            7,
		"REL_Y":            8,
		"REL_WHEEL":        9,
		"REL_WHEEL_HI_RES": 10,
		"KEY_UP":           11,
		"KEY_DOWN":         12,
		"KEY_LEFT":         13,
		"KEY_RIGHT":        14,
		"KEY_PAGEUP":       15,
		"KEY_PAGEDOWN":     16,
	}
)

func (x InputEventCodes) Enum() *InputEventCodes {
	p := new(InputEventCodes)
	*p = x
	return p
}

func (x InputEventCodes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InputEventCodes) Descriptor() protoreflect.EnumDescriptor {
	return file_localdevice_remote_control_proto_enumTypes[0].Descriptor()
}

func (InputEventCodes) Type() protoreflect.EnumType {
	return &file_localdevice_remote_control_proto_enumTypes[0]
}

func (x InputEventCodes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InputEventCodes.Descriptor instead.
func (InputEventCodes) EnumDescriptor() ([]byte, []int) {
	return file_localdevice_remote_control_proto_rawDescGZIP(), []int{0}
}

type EventState int32

const (
	EventState_PRESS   EventState = 0
	EventState_RELEASE EventState = 1
)

// Enum value maps for EventState.
var (
	EventState_name = map[int32]string{
		0: "PRESS",
		1: "RELEASE",
	}
	EventState_value = map[string]int32{
		"PRESS":   0,
		"RELEASE": 1,
	}
)

func (x EventState) Enum() *EventState {
	p := new(EventState)
	*p = x
	return p
}

func (x EventState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventState) Descriptor() protoreflect.EnumDescriptor {
	return file_localdevice_remote_control_proto_enumTypes[1].Descriptor()
}

func (EventState) Type() protoreflect.EnumType {
	return &file_localdevice_remote_control_proto_enumTypes[1]
}

func (x EventState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventState.Descriptor instead.
func (EventState) EnumDescriptor() ([]byte, []int) {
	return file_localdevice_remote_control_proto_rawDescGZIP(), []int{1}
}

type SendEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  InputEventCodes `protobuf:"varint,1,opt,name=code,proto3,enum=cloud.lazycat.apis.localdevice.InputEventCodes" json:"code,omitempty"`
	State EventState      `protobuf:"varint,2,opt,name=state,proto3,enum=cloud.lazycat.apis.localdevice.EventState" json:"state,omitempty"`
}

func (x *SendEventRequest) Reset() {
	*x = SendEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_remote_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEventRequest) ProtoMessage() {}

func (x *SendEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_remote_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEventRequest.ProtoReflect.Descriptor instead.
func (*SendEventRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_remote_control_proto_rawDescGZIP(), []int{0}
}

func (x *SendEventRequest) GetCode() InputEventCodes {
	if x != nil {
		return x.Code
	}
	return InputEventCodes_KEY_HOME
}

func (x *SendEventRequest) GetState() EventState {
	if x != nil {
		return x.State
	}
	return EventState_PRESS
}

var File_localdevice_remote_control_proto protoreflect.FileDescriptor

var file_localdevice_remote_control_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x99, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x93, 0x02, 0x0a, 0x0f,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x0c, 0x0a, 0x08, 0x4b, 0x45, 0x59, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x4b, 0x45, 0x59, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4b,
	0x45, 0x59, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x53, 0x43, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x45, 0x59, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x42, 0x54, 0x4e, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x04, 0x12, 0x0d,
	0x0a, 0x09, 0x42, 0x54, 0x4e, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x05, 0x12, 0x0e, 0x0a,
	0x0a, 0x42, 0x54, 0x4e, 0x5f, 0x4d, 0x49, 0x44, 0x44, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x09, 0x0a,
	0x05, 0x52, 0x45, 0x4c, 0x5f, 0x58, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x4c, 0x5f,
	0x59, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x4c, 0x5f, 0x57, 0x48, 0x45, 0x45, 0x4c,
	0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x45, 0x4c, 0x5f, 0x57, 0x48, 0x45, 0x45, 0x4c, 0x5f,
	0x48, 0x49, 0x5f, 0x52, 0x45, 0x53, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x45, 0x59, 0x5f,
	0x55, 0x50, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x45, 0x59, 0x5f, 0x44, 0x4f, 0x57, 0x4e,
	0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x45, 0x59, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x0d,
	0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x45, 0x59, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x0e, 0x12,
	0x0e, 0x0a, 0x0a, 0x4b, 0x45, 0x59, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x55, 0x50, 0x10, 0x0f, 0x12,
	0x10, 0x0a, 0x0c, 0x4b, 0x45, 0x59, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x44, 0x4f, 0x57, 0x4e, 0x10,
	0x10, 0x2a, 0x24, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x50, 0x52, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45,
	0x4c, 0x45, 0x41, 0x53, 0x45, 0x10, 0x01, 0x32, 0x6d, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x5c, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_localdevice_remote_control_proto_rawDescOnce sync.Once
	file_localdevice_remote_control_proto_rawDescData = file_localdevice_remote_control_proto_rawDesc
)

func file_localdevice_remote_control_proto_rawDescGZIP() []byte {
	file_localdevice_remote_control_proto_rawDescOnce.Do(func() {
		file_localdevice_remote_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_localdevice_remote_control_proto_rawDescData)
	})
	return file_localdevice_remote_control_proto_rawDescData
}

var file_localdevice_remote_control_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_localdevice_remote_control_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_localdevice_remote_control_proto_goTypes = []interface{}{
	(InputEventCodes)(0),     // 0: cloud.lazycat.apis.localdevice.InputEventCodes
	(EventState)(0),          // 1: cloud.lazycat.apis.localdevice.EventState
	(*SendEventRequest)(nil), // 2: cloud.lazycat.apis.localdevice.SendEventRequest
	(*emptypb.Empty)(nil),    // 3: google.protobuf.Empty
}
var file_localdevice_remote_control_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.apis.localdevice.SendEventRequest.code:type_name -> cloud.lazycat.apis.localdevice.InputEventCodes
	1, // 1: cloud.lazycat.apis.localdevice.SendEventRequest.state:type_name -> cloud.lazycat.apis.localdevice.EventState
	2, // 2: cloud.lazycat.apis.localdevice.RemoteControl.SendInputEvent:input_type -> cloud.lazycat.apis.localdevice.SendEventRequest
	3, // 3: cloud.lazycat.apis.localdevice.RemoteControl.SendInputEvent:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_localdevice_remote_control_proto_init() }
func file_localdevice_remote_control_proto_init() {
	if File_localdevice_remote_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_localdevice_remote_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEventRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_localdevice_remote_control_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_localdevice_remote_control_proto_goTypes,
		DependencyIndexes: file_localdevice_remote_control_proto_depIdxs,
		EnumInfos:         file_localdevice_remote_control_proto_enumTypes,
		MessageInfos:      file_localdevice_remote_control_proto_msgTypes,
	}.Build()
	File_localdevice_remote_control_proto = out.File
	file_localdevice_remote_control_proto_rawDesc = nil
	file_localdevice_remote_control_proto_goTypes = nil
	file_localdevice_remote_control_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: localdevice/local-launch.proto

package localdevice

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

type OpenUnsafeAppRequest_Position int32

const (
	OpenUnsafeAppRequest_Left   OpenUnsafeAppRequest_Position = 0
	OpenUnsafeAppRequest_Right  OpenUnsafeAppRequest_Position = 1
	OpenUnsafeAppRequest_Top    OpenUnsafeAppRequest_Position = 2
	OpenUnsafeAppRequest_Bottom OpenUnsafeAppRequest_Position = 3
)

// Enum value maps for OpenUnsafeAppRequest_Position.
var (
	OpenUnsafeAppRequest_Position_name = map[int32]string{
		0: "Left",
		1: "Right",
		2: "Top",
		3: "Bottom",
	}
	OpenUnsafeAppRequest_Position_value = map[string]int32{
		"Left":   0,
		"Right":  1,
		"Top":    2,
		"Bottom": 3,
	}
)

func (x OpenUnsafeAppRequest_Position) Enum() *OpenUnsafeAppRequest_Position {
	p := new(OpenUnsafeAppRequest_Position)
	*p = x
	return p
}

func (x OpenUnsafeAppRequest_Position) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenUnsafeAppRequest_Position) Descriptor() protoreflect.EnumDescriptor {
	return file_localdevice_local_launch_proto_enumTypes[0].Descriptor()
}

func (OpenUnsafeAppRequest_Position) Type() protoreflect.EnumType {
	return &file_localdevice_local_launch_proto_enumTypes[0]
}

func (x OpenUnsafeAppRequest_Position) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenUnsafeAppRequest_Position.Descriptor instead.
func (OpenUnsafeAppRequest_Position) EnumDescriptor() ([]byte, []int) {
	return file_localdevice_local_launch_proto_rawDescGZIP(), []int{6, 0}
}

// 所有的功能都通过 window.lzc_control_api对象注入. (后面以$API代指)
type OpenUnsafeAppRequest_Feature int32

const (
	OpenUnsafeAppRequest_InjectJS   OpenUnsafeAppRequest_Feature = 0 //$API.InjectContentJS(jscontent)
	OpenUnsafeAppRequest_ReadCookie OpenUnsafeAppRequest_Feature = 1 //$API.ReadCookie(domain) -> string
)

// Enum value maps for OpenUnsafeAppRequest_Feature.
var (
	OpenUnsafeAppRequest_Feature_name = map[int32]string{
		0: "InjectJS",
		1: "ReadCookie",
	}
	OpenUnsafeAppRequest_Feature_value = map[string]int32{
		"InjectJS":   0,
		"ReadCookie": 1,
	}
)

func (x OpenUnsafeAppRequest_Feature) Enum() *OpenUnsafeAppRequest_Feature {
	p := new(OpenUnsafeAppRequest_Feature)
	*p = x
	return p
}

func (x OpenUnsafeAppRequest_Feature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenUnsafeAppRequest_Feature) Descriptor() protoreflect.EnumDescriptor {
	return file_localdevice_local_launch_proto_enumTypes[1].Descriptor()
}

func (OpenUnsafeAppRequest_Feature) Type() protoreflect.EnumType {
	return &file_localdevice_local_launch_proto_enumTypes[1]
}

func (x OpenUnsafeAppRequest_Feature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenUnsafeAppRequest_Feature.Descriptor instead.
func (OpenUnsafeAppRequest_Feature) EnumDescriptor() ([]byte, []int) {
	return file_localdevice_local_launch_proto_rawDescGZIP(), []int{6, 1}
}

type PinAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	AppName string `protobuf:"bytes,2,opt,name=appName,proto3" json:"appName,omitempty"`
	// icon地址允许是http类型或者base64类型,如果为空将尝试获取网站的favicon.ico
	IconUrl string `protobuf:"bytes,3,opt,name=iconUrl,proto3" json:"iconUrl,omitempty"`
}

func (x *PinAppRequest) Reset() {
	*x = PinAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_local_launch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PinAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PinAppRequest) ProtoMessage() {}

func (x *PinAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_local_launch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PinAppRequest.ProtoReflect.Descriptor instead.
func (*PinAppRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_local_launch_proto_rawDescGZIP(), []int{0}
}

func (x *PinAppRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PinAppRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *PinAppRequest) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

type PinAppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PinAppReply) Reset() {
	*x = PinAppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_local_launch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PinAppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PinAppReply) ProtoMessage() {}

func (x *PinAppReply) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_local_launch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PinAppReply.ProtoReflect.Descriptor instead.
func (*PinAppReply) Descriptor() ([]byte, []int) {
	return file_localdevice_local_launch_proto_rawDescGZIP(), []int{1}
}

type UnPinAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UnPinAppRequest) Reset() {
	*x = UnPinAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_local_launch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnPinAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnPinAppRequest) ProtoMessage() {}

func (x *UnPinAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_local_launch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnPinAppRequest.ProtoReflect.Descriptor instead.
func (*UnPinAppRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_local_launch_proto_rawDescGZIP(), []int{2}
}

func (x *UnPinAppRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UnPinAppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnPinAppReply) Reset() {
	*x = UnPinAppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_local_launch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnPinAppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnPinAppReply) ProtoMessage() {}

func (x *UnPinAppReply) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_local_launch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnPinAppReply.ProtoReflect.Descriptor instead.
func (*UnPinAppReply) Descriptor() ([]byte, []int) {
	return file_localdevice_local_launch_proto_rawDescGZIP(), []int{3}
}

type OpenAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *OpenAppRequest) Reset() {
	*x = OpenAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_local_launch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAppRequest) ProtoMessage() {}

func (x *OpenAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_local_launch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAppRequest.ProtoReflect.Descriptor instead.
func (*OpenAppRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_local_launch_proto_rawDescGZIP(), []int{4}
}

func (x *OpenAppRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type OpenAppReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OpenAppReply) Reset() {
	*x = OpenAppReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_local_launch_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAppReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAppReply) ProtoMessage() {}

func (x *OpenAppReply) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_local_launch_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAppReply.ProtoReflect.Descriptor instead.
func (*OpenAppReply) Descriptor() ([]byte, []int) {
	return file_localdevice_local_launch_proto_rawDescGZIP(), []int{5}
}

type OpenUnsafeAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// control-view的地址
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// control-view的位置
	Position OpenUnsafeAppRequest_Position `protobuf:"varint,2,opt,name=position,proto3,enum=cloud.lazycat.apis.localdevice.OpenUnsafeAppRequest_Position" json:"position,omitempty"`
	// control-view的高度
	Height string `protobuf:"bytes,3,opt,name=height,proto3" json:"height,omitempty"` // 10%; 50px;
	//
	//此外control-view会接收以下特殊事件, 所有事件均统一以post-message形式发送， msg格式为`{ type: string, msg: string }`
	//control-view需要自行调用addEventListener('message'), 并过滤type==lzc_control_api类型的message.
	//
	//1. OnURLChange(new_url)
	//2. OnNewLinkClick(url)
	//3. OnDownloadLinkRequest(url)
	//4. OnContentMessage(msg string) //在content-view中调用post-message发送的任何消息
	Features []OpenUnsafeAppRequest_Feature `protobuf:"varint,4,rep,packed,name=features,proto3,enum=cloud.lazycat.apis.localdevice.OpenUnsafeAppRequest_Feature" json:"features,omitempty"`
}

func (x *OpenUnsafeAppRequest) Reset() {
	*x = OpenUnsafeAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_local_launch_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenUnsafeAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenUnsafeAppRequest) ProtoMessage() {}

func (x *OpenUnsafeAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_local_launch_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenUnsafeAppRequest.ProtoReflect.Descriptor instead.
func (*OpenUnsafeAppRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_local_launch_proto_rawDescGZIP(), []int{6}
}

func (x *OpenUnsafeAppRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *OpenUnsafeAppRequest) GetPosition() OpenUnsafeAppRequest_Position {
	if x != nil {
		return x.Position
	}
	return OpenUnsafeAppRequest_Left
}

func (x *OpenUnsafeAppRequest) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *OpenUnsafeAppRequest) GetFeatures() []OpenUnsafeAppRequest_Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

var File_localdevice_local_launch_proto protoreflect.FileDescriptor

var file_localdevice_local_launch_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x2d, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x55, 0x0a, 0x0d, 0x50, 0x69, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x0a, 0x0f, 0x55, 0x6e, 0x50, 0x69, 0x6e, 0x41,
	0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x0f, 0x0a, 0x0d, 0x55,
	0x6e, 0x50, 0x69, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x0a, 0x0e,
	0x4f, 0x70, 0x65, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x0e, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0xd4, 0x02, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x6e, 0x55, 0x6e, 0x73, 0x61, 0x66, 0x65, 0x41,
	0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x59, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f,
	0x70, 0x65, 0x6e, 0x55, 0x6e, 0x73, 0x61, 0x66, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x58,
	0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x3c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x55, 0x6e, 0x73, 0x61, 0x66, 0x65, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x65, 0x66, 0x74, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x69, 0x67, 0x68, 0x74, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x6f, 0x70,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x10, 0x03, 0x22, 0x27,
	0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x6a,
	0x65, 0x63, 0x74, 0x4a, 0x53, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x43,
	0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x10, 0x01, 0x32, 0xcc, 0x03, 0x0a, 0x12, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66,
	0x0a, 0x06, 0x50, 0x69, 0x6e, 0x41, 0x70, 0x70, 0x12, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x08, 0x55, 0x6e, 0x50, 0x69, 0x6e, 0x41,
	0x70, 0x70, 0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x6e, 0x50, 0x69, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x6e, 0x50, 0x69, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x07, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x70, 0x70, 0x12,
	0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x75, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x6e, 0x55, 0x6e, 0x73, 0x61, 0x66, 0x65, 0x41, 0x70, 0x70,
	0x12, 0x34, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x55, 0x6e, 0x73, 0x61, 0x66, 0x65, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_localdevice_local_launch_proto_rawDescOnce sync.Once
	file_localdevice_local_launch_proto_rawDescData = file_localdevice_local_launch_proto_rawDesc
)

func file_localdevice_local_launch_proto_rawDescGZIP() []byte {
	file_localdevice_local_launch_proto_rawDescOnce.Do(func() {
		file_localdevice_local_launch_proto_rawDescData = protoimpl.X.CompressGZIP(file_localdevice_local_launch_proto_rawDescData)
	})
	return file_localdevice_local_launch_proto_rawDescData
}

var file_localdevice_local_launch_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_localdevice_local_launch_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_localdevice_local_launch_proto_goTypes = []interface{}{
	(OpenUnsafeAppRequest_Position)(0), // 0: cloud.lazycat.apis.localdevice.OpenUnsafeAppRequest.Position
	(OpenUnsafeAppRequest_Feature)(0),  // 1: cloud.lazycat.apis.localdevice.OpenUnsafeAppRequest.Feature
	(*PinAppRequest)(nil),              // 2: cloud.lazycat.apis.localdevice.PinAppRequest
	(*PinAppReply)(nil),                // 3: cloud.lazycat.apis.localdevice.PinAppReply
	(*UnPinAppRequest)(nil),            // 4: cloud.lazycat.apis.localdevice.UnPinAppRequest
	(*UnPinAppReply)(nil),              // 5: cloud.lazycat.apis.localdevice.UnPinAppReply
	(*OpenAppRequest)(nil),             // 6: cloud.lazycat.apis.localdevice.OpenAppRequest
	(*OpenAppReply)(nil),               // 7: cloud.lazycat.apis.localdevice.OpenAppReply
	(*OpenUnsafeAppRequest)(nil),       // 8: cloud.lazycat.apis.localdevice.OpenUnsafeAppRequest
}
var file_localdevice_local_launch_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.apis.localdevice.OpenUnsafeAppRequest.position:type_name -> cloud.lazycat.apis.localdevice.OpenUnsafeAppRequest.Position
	1, // 1: cloud.lazycat.apis.localdevice.OpenUnsafeAppRequest.features:type_name -> cloud.lazycat.apis.localdevice.OpenUnsafeAppRequest.Feature
	2, // 2: cloud.lazycat.apis.localdevice.LocalLaunchService.PinApp:input_type -> cloud.lazycat.apis.localdevice.PinAppRequest
	4, // 3: cloud.lazycat.apis.localdevice.LocalLaunchService.UnPinApp:input_type -> cloud.lazycat.apis.localdevice.UnPinAppRequest
	6, // 4: cloud.lazycat.apis.localdevice.LocalLaunchService.OpenApp:input_type -> cloud.lazycat.apis.localdevice.OpenAppRequest
	8, // 5: cloud.lazycat.apis.localdevice.LocalLaunchService.OpenUnsafeApp:input_type -> cloud.lazycat.apis.localdevice.OpenUnsafeAppRequest
	3, // 6: cloud.lazycat.apis.localdevice.LocalLaunchService.PinApp:output_type -> cloud.lazycat.apis.localdevice.PinAppReply
	5, // 7: cloud.lazycat.apis.localdevice.LocalLaunchService.UnPinApp:output_type -> cloud.lazycat.apis.localdevice.UnPinAppReply
	7, // 8: cloud.lazycat.apis.localdevice.LocalLaunchService.OpenApp:output_type -> cloud.lazycat.apis.localdevice.OpenAppReply
	7, // 9: cloud.lazycat.apis.localdevice.LocalLaunchService.OpenUnsafeApp:output_type -> cloud.lazycat.apis.localdevice.OpenAppReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_localdevice_local_launch_proto_init() }
func file_localdevice_local_launch_proto_init() {
	if File_localdevice_local_launch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_localdevice_local_launch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PinAppRequest); i {
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
		file_localdevice_local_launch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PinAppReply); i {
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
		file_localdevice_local_launch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnPinAppRequest); i {
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
		file_localdevice_local_launch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnPinAppReply); i {
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
		file_localdevice_local_launch_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAppRequest); i {
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
		file_localdevice_local_launch_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAppReply); i {
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
		file_localdevice_local_launch_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenUnsafeAppRequest); i {
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
			RawDescriptor: file_localdevice_local_launch_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_localdevice_local_launch_proto_goTypes,
		DependencyIndexes: file_localdevice_local_launch_proto_depIdxs,
		EnumInfos:         file_localdevice_local_launch_proto_enumTypes,
		MessageInfos:      file_localdevice_local_launch_proto_msgTypes,
	}.Build()
	File_localdevice_local_launch_proto = out.File
	file_localdevice_local_launch_proto_rawDesc = nil
	file_localdevice_local_launch_proto_goTypes = nil
	file_localdevice_local_launch_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: localdevice/dialog.proto

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

type OpenFileSeletorRequest_SelectType int32

const (
	OpenFileSeletorRequest_Dir  OpenFileSeletorRequest_SelectType = 0
	OpenFileSeletorRequest_File OpenFileSeletorRequest_SelectType = 1
)

// Enum value maps for OpenFileSeletorRequest_SelectType.
var (
	OpenFileSeletorRequest_SelectType_name = map[int32]string{
		0: "Dir",
		1: "File",
	}
	OpenFileSeletorRequest_SelectType_value = map[string]int32{
		"Dir":  0,
		"File": 1,
	}
)

func (x OpenFileSeletorRequest_SelectType) Enum() *OpenFileSeletorRequest_SelectType {
	p := new(OpenFileSeletorRequest_SelectType)
	*p = x
	return p
}

func (x OpenFileSeletorRequest_SelectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenFileSeletorRequest_SelectType) Descriptor() protoreflect.EnumDescriptor {
	return file_localdevice_dialog_proto_enumTypes[0].Descriptor()
}

func (OpenFileSeletorRequest_SelectType) Type() protoreflect.EnumType {
	return &file_localdevice_dialog_proto_enumTypes[0]
}

func (x OpenFileSeletorRequest_SelectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenFileSeletorRequest_SelectType.Descriptor instead.
func (OpenFileSeletorRequest_SelectType) EnumDescriptor() ([]byte, []int) {
	return file_localdevice_dialog_proto_rawDescGZIP(), []int{5, 0}
}

type QuestionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text          string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	DefaultCancel bool   `protobuf:"varint,3,opt,name=defaultCancel,proto3" json:"defaultCancel,omitempty"`
}

func (x *QuestionRequest) Reset() {
	*x = QuestionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_dialog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionRequest) ProtoMessage() {}

func (x *QuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_dialog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionRequest.ProtoReflect.Descriptor instead.
func (*QuestionRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_dialog_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *QuestionRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *QuestionRequest) GetDefaultCancel() bool {
	if x != nil {
		return x.DefaultCancel
	}
	return false
}

type QuestionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yes bool `protobuf:"varint,1,opt,name=yes,proto3" json:"yes,omitempty"`
}

func (x *QuestionResult) Reset() {
	*x = QuestionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_dialog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionResult) ProtoMessage() {}

func (x *QuestionResult) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_dialog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionResult.ProtoReflect.Descriptor instead.
func (*QuestionResult) Descriptor() ([]byte, []int) {
	return file_localdevice_dialog_proto_rawDescGZIP(), []int{1}
}

func (x *QuestionResult) GetYes() bool {
	if x != nil {
		return x.Yes
	}
	return false
}

type MessageBoxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text  string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *MessageBoxRequest) Reset() {
	*x = MessageBoxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_dialog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageBoxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageBoxRequest) ProtoMessage() {}

func (x *MessageBoxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_dialog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageBoxRequest.ProtoReflect.Descriptor instead.
func (*MessageBoxRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_dialog_proto_rawDescGZIP(), []int{2}
}

func (x *MessageBoxRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MessageBoxRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type PasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text  string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PasswordRequest) Reset() {
	*x = PasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_dialog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordRequest) ProtoMessage() {}

func (x *PasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_dialog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordRequest.ProtoReflect.Descriptor instead.
func (*PasswordRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_dialog_proto_rawDescGZIP(), []int{3}
}

func (x *PasswordRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PasswordRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type PasswordResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Ok       bool   `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *PasswordResult) Reset() {
	*x = PasswordResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_dialog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordResult) ProtoMessage() {}

func (x *PasswordResult) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_dialog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordResult.ProtoReflect.Descriptor instead.
func (*PasswordResult) Descriptor() ([]byte, []int) {
	return file_localdevice_dialog_proto_rawDescGZIP(), []int{4}
}

func (x *PasswordResult) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PasswordResult) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type OpenFileSeletorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 选择类型
	Type OpenFileSeletorRequest_SelectType `protobuf:"varint,1,opt,name=type,proto3,enum=cloud.lazycat.apis.localdevice.OpenFileSeletorRequest_SelectType" json:"type,omitempty"`
	// 是否多选（多选在SelectType 为Dir时无效）
	IsSingle bool `protobuf:"varint,2,opt,name=isSingle,proto3" json:"isSingle,omitempty"`
	// 文件后缀过滤器
	// 使用2种类型
	// 第一种 jpb,png,mp3，指定文件的后缀名,使用,分割
	// 支持的大类型有必须 image/* , 或者image/*,audio/*,video/*,document/*, 四种大类可以写在一起，也可以分开写,使用,分割。 如何该字段为空，则不过滤
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *OpenFileSeletorRequest) Reset() {
	*x = OpenFileSeletorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_dialog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenFileSeletorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenFileSeletorRequest) ProtoMessage() {}

func (x *OpenFileSeletorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_dialog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenFileSeletorRequest.ProtoReflect.Descriptor instead.
func (*OpenFileSeletorRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_dialog_proto_rawDescGZIP(), []int{5}
}

func (x *OpenFileSeletorRequest) GetType() OpenFileSeletorRequest_SelectType {
	if x != nil {
		return x.Type
	}
	return OpenFileSeletorRequest_Dir
}

func (x *OpenFileSeletorRequest) GetIsSingle() bool {
	if x != nil {
		return x.IsSingle
	}
	return false
}

func (x *OpenFileSeletorRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type OpenFileSeletorResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件或者目录的路径
	Path []string `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *OpenFileSeletorResult) Reset() {
	*x = OpenFileSeletorResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_dialog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenFileSeletorResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenFileSeletorResult) ProtoMessage() {}

func (x *OpenFileSeletorResult) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_dialog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenFileSeletorResult.ProtoReflect.Descriptor instead.
func (*OpenFileSeletorResult) Descriptor() ([]byte, []int) {
	return file_localdevice_dialog_proto_rawDescGZIP(), []int{6}
}

func (x *OpenFileSeletorResult) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

var File_localdevice_dialog_proto protoreflect.FileDescriptor

var file_localdevice_dialog_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x22, 0x22, 0x0a, 0x0e, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x79, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x79, 0x65, 0x73, 0x22, 0x3d,
	0x0a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3b, 0x0a,
	0x0f, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3c, 0x0a, 0x0e, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0xc4, 0x01, 0x0a, 0x16, 0x4f, 0x70, 0x65,
	0x6e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x41, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x1f,
	0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03,
	0x44, 0x69, 0x72, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x01, 0x22,
	0x2b, 0x0a, 0x15, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x32, 0xcd, 0x03, 0x0a,
	0x0d, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x6d,
	0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x78, 0x12, 0x31, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61,
	0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x6e,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x12, 0x36, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65,
	0x6e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x6c,
	0x65, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65,
	0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_localdevice_dialog_proto_rawDescOnce sync.Once
	file_localdevice_dialog_proto_rawDescData = file_localdevice_dialog_proto_rawDesc
)

func file_localdevice_dialog_proto_rawDescGZIP() []byte {
	file_localdevice_dialog_proto_rawDescOnce.Do(func() {
		file_localdevice_dialog_proto_rawDescData = protoimpl.X.CompressGZIP(file_localdevice_dialog_proto_rawDescData)
	})
	return file_localdevice_dialog_proto_rawDescData
}

var file_localdevice_dialog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_localdevice_dialog_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_localdevice_dialog_proto_goTypes = []interface{}{
	(OpenFileSeletorRequest_SelectType)(0), // 0: cloud.lazycat.apis.localdevice.OpenFileSeletorRequest.SelectType
	(*QuestionRequest)(nil),                // 1: cloud.lazycat.apis.localdevice.QuestionRequest
	(*QuestionResult)(nil),                 // 2: cloud.lazycat.apis.localdevice.QuestionResult
	(*MessageBoxRequest)(nil),              // 3: cloud.lazycat.apis.localdevice.MessageBoxRequest
	(*PasswordRequest)(nil),                // 4: cloud.lazycat.apis.localdevice.PasswordRequest
	(*PasswordResult)(nil),                 // 5: cloud.lazycat.apis.localdevice.PasswordResult
	(*OpenFileSeletorRequest)(nil),         // 6: cloud.lazycat.apis.localdevice.OpenFileSeletorRequest
	(*OpenFileSeletorResult)(nil),          // 7: cloud.lazycat.apis.localdevice.OpenFileSeletorResult
	(*emptypb.Empty)(nil),                  // 8: google.protobuf.Empty
}
var file_localdevice_dialog_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.apis.localdevice.OpenFileSeletorRequest.type:type_name -> cloud.lazycat.apis.localdevice.OpenFileSeletorRequest.SelectType
	1, // 1: cloud.lazycat.apis.localdevice.DialogManager.Question:input_type -> cloud.lazycat.apis.localdevice.QuestionRequest
	3, // 2: cloud.lazycat.apis.localdevice.DialogManager.MessageBox:input_type -> cloud.lazycat.apis.localdevice.MessageBoxRequest
	4, // 3: cloud.lazycat.apis.localdevice.DialogManager.Password:input_type -> cloud.lazycat.apis.localdevice.PasswordRequest
	6, // 4: cloud.lazycat.apis.localdevice.DialogManager.OpenFileSeletor:input_type -> cloud.lazycat.apis.localdevice.OpenFileSeletorRequest
	2, // 5: cloud.lazycat.apis.localdevice.DialogManager.Question:output_type -> cloud.lazycat.apis.localdevice.QuestionResult
	8, // 6: cloud.lazycat.apis.localdevice.DialogManager.MessageBox:output_type -> google.protobuf.Empty
	5, // 7: cloud.lazycat.apis.localdevice.DialogManager.Password:output_type -> cloud.lazycat.apis.localdevice.PasswordResult
	7, // 8: cloud.lazycat.apis.localdevice.DialogManager.OpenFileSeletor:output_type -> cloud.lazycat.apis.localdevice.OpenFileSeletorResult
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_localdevice_dialog_proto_init() }
func file_localdevice_dialog_proto_init() {
	if File_localdevice_dialog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_localdevice_dialog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionRequest); i {
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
		file_localdevice_dialog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionResult); i {
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
		file_localdevice_dialog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageBoxRequest); i {
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
		file_localdevice_dialog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordRequest); i {
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
		file_localdevice_dialog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordResult); i {
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
		file_localdevice_dialog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenFileSeletorRequest); i {
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
		file_localdevice_dialog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenFileSeletorResult); i {
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
			RawDescriptor: file_localdevice_dialog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_localdevice_dialog_proto_goTypes,
		DependencyIndexes: file_localdevice_dialog_proto_depIdxs,
		EnumInfos:         file_localdevice_dialog_proto_enumTypes,
		MessageInfos:      file_localdevice_dialog_proto_msgTypes,
	}.Build()
	File_localdevice_dialog_proto = out.File
	file_localdevice_dialog_proto_rawDesc = nil
	file_localdevice_dialog_proto_goTypes = nil
	file_localdevice_dialog_proto_depIdxs = nil
}

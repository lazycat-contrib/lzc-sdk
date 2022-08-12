// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
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
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xc8, 0x02, 0x0a, 0x0d, 0x44, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x6d, 0x0a, 0x08, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x78, 0x12, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_localdevice_dialog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_localdevice_dialog_proto_goTypes = []interface{}{
	(*QuestionRequest)(nil),   // 0: cloud.lazycat.apis.localdevice.QuestionRequest
	(*QuestionResult)(nil),    // 1: cloud.lazycat.apis.localdevice.QuestionResult
	(*MessageBoxRequest)(nil), // 2: cloud.lazycat.apis.localdevice.MessageBoxRequest
	(*PasswordRequest)(nil),   // 3: cloud.lazycat.apis.localdevice.PasswordRequest
	(*PasswordResult)(nil),    // 4: cloud.lazycat.apis.localdevice.PasswordResult
	(*emptypb.Empty)(nil),     // 5: google.protobuf.Empty
}
var file_localdevice_dialog_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.apis.localdevice.DialogManager.Question:input_type -> cloud.lazycat.apis.localdevice.QuestionRequest
	2, // 1: cloud.lazycat.apis.localdevice.DialogManager.MessageBox:input_type -> cloud.lazycat.apis.localdevice.MessageBoxRequest
	3, // 2: cloud.lazycat.apis.localdevice.DialogManager.Password:input_type -> cloud.lazycat.apis.localdevice.PasswordRequest
	1, // 3: cloud.lazycat.apis.localdevice.DialogManager.Question:output_type -> cloud.lazycat.apis.localdevice.QuestionResult
	5, // 4: cloud.lazycat.apis.localdevice.DialogManager.MessageBox:output_type -> google.protobuf.Empty
	4, // 5: cloud.lazycat.apis.localdevice.DialogManager.Password:output_type -> cloud.lazycat.apis.localdevice.PasswordResult
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_localdevice_dialog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_localdevice_dialog_proto_goTypes,
		DependencyIndexes: file_localdevice_dialog_proto_depIdxs,
		MessageInfos:      file_localdevice_dialog_proto_msgTypes,
	}.Build()
	File_localdevice_dialog_proto = out.File
	file_localdevice_dialog_proto_rawDesc = nil
	file_localdevice_dialog_proto_goTypes = nil
	file_localdevice_dialog_proto_depIdxs = nil
}

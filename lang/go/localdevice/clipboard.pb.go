// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: localdevice/clipboard.proto

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

type ReadClipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mime string `protobuf:"bytes,1,opt,name=mime,proto3" json:"mime,omitempty"` //目前只支持text/plain和image/png
}

func (x *ReadClipRequest) Reset() {
	*x = ReadClipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_clipboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadClipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadClipRequest) ProtoMessage() {}

func (x *ReadClipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_clipboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadClipRequest.ProtoReflect.Descriptor instead.
func (*ReadClipRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_clipboard_proto_rawDescGZIP(), []int{0}
}

func (x *ReadClipRequest) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

type ReadClipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ReadClipResponse) Reset() {
	*x = ReadClipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_clipboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadClipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadClipResponse) ProtoMessage() {}

func (x *ReadClipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_clipboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadClipResponse.ProtoReflect.Descriptor instead.
func (*ReadClipResponse) Descriptor() ([]byte, []int) {
	return file_localdevice_clipboard_proto_rawDescGZIP(), []int{1}
}

func (x *ReadClipResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type WriteClipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mime    string `protobuf:"bytes,1,opt,name=mime,proto3" json:"mime,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *WriteClipRequest) Reset() {
	*x = WriteClipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_clipboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteClipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteClipRequest) ProtoMessage() {}

func (x *WriteClipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_clipboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteClipRequest.ProtoReflect.Descriptor instead.
func (*WriteClipRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_clipboard_proto_rawDescGZIP(), []int{2}
}

func (x *WriteClipRequest) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *WriteClipRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type WriteClipResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *WriteClipResponse) Reset() {
	*x = WriteClipResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_clipboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteClipResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteClipResponse) ProtoMessage() {}

func (x *WriteClipResponse) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_clipboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteClipResponse.ProtoReflect.Descriptor instead.
func (*WriteClipResponse) Descriptor() ([]byte, []int) {
	return file_localdevice_clipboard_proto_rawDescGZIP(), []int{3}
}

func (x *WriteClipResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_localdevice_clipboard_proto protoreflect.FileDescriptor

var file_localdevice_clipboard_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6c,
	0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x25, 0x0a,
	0x0f, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x69, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x40, 0x0a, 0x10, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0xdf, 0x02, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x6b, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64,
	0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x30,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2f,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_localdevice_clipboard_proto_rawDescOnce sync.Once
	file_localdevice_clipboard_proto_rawDescData = file_localdevice_clipboard_proto_rawDesc
)

func file_localdevice_clipboard_proto_rawDescGZIP() []byte {
	file_localdevice_clipboard_proto_rawDescOnce.Do(func() {
		file_localdevice_clipboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_localdevice_clipboard_proto_rawDescData)
	})
	return file_localdevice_clipboard_proto_rawDescData
}

var file_localdevice_clipboard_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_localdevice_clipboard_proto_goTypes = []interface{}{
	(*ReadClipRequest)(nil),   // 0: cloud.lazycat.apis.localdevice.ReadClipRequest
	(*ReadClipResponse)(nil),  // 1: cloud.lazycat.apis.localdevice.ReadClipResponse
	(*WriteClipRequest)(nil),  // 2: cloud.lazycat.apis.localdevice.WriteClipRequest
	(*WriteClipResponse)(nil), // 3: cloud.lazycat.apis.localdevice.WriteClipResponse
}
var file_localdevice_clipboard_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.apis.localdevice.ClipboardManager.Read:input_type -> cloud.lazycat.apis.localdevice.ReadClipRequest
	2, // 1: cloud.lazycat.apis.localdevice.ClipboardManager.Write:input_type -> cloud.lazycat.apis.localdevice.WriteClipRequest
	0, // 2: cloud.lazycat.apis.localdevice.ClipboardManager.Watch:input_type -> cloud.lazycat.apis.localdevice.ReadClipRequest
	1, // 3: cloud.lazycat.apis.localdevice.ClipboardManager.Read:output_type -> cloud.lazycat.apis.localdevice.ReadClipResponse
	3, // 4: cloud.lazycat.apis.localdevice.ClipboardManager.Write:output_type -> cloud.lazycat.apis.localdevice.WriteClipResponse
	1, // 5: cloud.lazycat.apis.localdevice.ClipboardManager.Watch:output_type -> cloud.lazycat.apis.localdevice.ReadClipResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_localdevice_clipboard_proto_init() }
func file_localdevice_clipboard_proto_init() {
	if File_localdevice_clipboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_localdevice_clipboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadClipRequest); i {
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
		file_localdevice_clipboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadClipResponse); i {
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
		file_localdevice_clipboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteClipRequest); i {
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
		file_localdevice_clipboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteClipResponse); i {
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
			RawDescriptor: file_localdevice_clipboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_localdevice_clipboard_proto_goTypes,
		DependencyIndexes: file_localdevice_clipboard_proto_depIdxs,
		MessageInfos:      file_localdevice_clipboard_proto_msgTypes,
	}.Build()
	File_localdevice_clipboard_proto = out.File
	file_localdevice_clipboard_proto_rawDesc = nil
	file_localdevice_clipboard_proto_goTypes = nil
	file_localdevice_clipboard_proto_depIdxs = nil
}

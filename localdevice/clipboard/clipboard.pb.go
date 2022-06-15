// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: localdevice/clipboard/clipboard.proto

package clipboard

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
		mi := &file_localdevice_clipboard_clipboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadClipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadClipRequest) ProtoMessage() {}

func (x *ReadClipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_clipboard_clipboard_proto_msgTypes[0]
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
	return file_localdevice_clipboard_clipboard_proto_rawDescGZIP(), []int{0}
}

func (x *ReadClipRequest) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

type ReadClipReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ReadClipReply) Reset() {
	*x = ReadClipReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_clipboard_clipboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadClipReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadClipReply) ProtoMessage() {}

func (x *ReadClipReply) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_clipboard_clipboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadClipReply.ProtoReflect.Descriptor instead.
func (*ReadClipReply) Descriptor() ([]byte, []int) {
	return file_localdevice_clipboard_clipboard_proto_rawDescGZIP(), []int{1}
}

func (x *ReadClipReply) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type WriteClipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mimie   string `protobuf:"bytes,1,opt,name=mimie,proto3" json:"mimie,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *WriteClipRequest) Reset() {
	*x = WriteClipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_clipboard_clipboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteClipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteClipRequest) ProtoMessage() {}

func (x *WriteClipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_clipboard_clipboard_proto_msgTypes[2]
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
	return file_localdevice_clipboard_clipboard_proto_rawDescGZIP(), []int{2}
}

func (x *WriteClipRequest) GetMimie() string {
	if x != nil {
		return x.Mimie
	}
	return ""
}

func (x *WriteClipRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type WriteClipReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteClipReply) Reset() {
	*x = WriteClipReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_clipboard_clipboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteClipReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteClipReply) ProtoMessage() {}

func (x *WriteClipReply) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_clipboard_clipboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteClipReply.ProtoReflect.Descriptor instead.
func (*WriteClipReply) Descriptor() ([]byte, []int) {
	return file_localdevice_clipboard_clipboard_proto_rawDescGZIP(), []int{3}
}

var File_localdevice_clipboard_clipboard_proto protoreflect.FileDescriptor

var file_localdevice_clipboard_clipboard_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6c,
	0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x22, 0x25, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x43,
	0x6c, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x22, 0x29,
	0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x10, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x69, 0x6d, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69,
	0x6d, 0x69, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x10, 0x0a,
	0x0e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32,
	0xd6, 0x02, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x2f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x6b,
	0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6c,
	0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6c, 0x69, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x05, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61,
	0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x70, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x63, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_localdevice_clipboard_clipboard_proto_rawDescOnce sync.Once
	file_localdevice_clipboard_clipboard_proto_rawDescData = file_localdevice_clipboard_clipboard_proto_rawDesc
)

func file_localdevice_clipboard_clipboard_proto_rawDescGZIP() []byte {
	file_localdevice_clipboard_clipboard_proto_rawDescOnce.Do(func() {
		file_localdevice_clipboard_clipboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_localdevice_clipboard_clipboard_proto_rawDescData)
	})
	return file_localdevice_clipboard_clipboard_proto_rawDescData
}

var file_localdevice_clipboard_clipboard_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_localdevice_clipboard_clipboard_proto_goTypes = []interface{}{
	(*ReadClipRequest)(nil),  // 0: cloud.lazycat.localdevice.apis.ReadClipRequest
	(*ReadClipReply)(nil),    // 1: cloud.lazycat.localdevice.apis.ReadClipReply
	(*WriteClipRequest)(nil), // 2: cloud.lazycat.localdevice.apis.WriteClipRequest
	(*WriteClipReply)(nil),   // 3: cloud.lazycat.localdevice.apis.WriteClipReply
}
var file_localdevice_clipboard_clipboard_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.localdevice.apis.ClipboardManager.Read:input_type -> cloud.lazycat.localdevice.apis.ReadClipRequest
	2, // 1: cloud.lazycat.localdevice.apis.ClipboardManager.Write:input_type -> cloud.lazycat.localdevice.apis.WriteClipRequest
	0, // 2: cloud.lazycat.localdevice.apis.ClipboardManager.Watch:input_type -> cloud.lazycat.localdevice.apis.ReadClipRequest
	1, // 3: cloud.lazycat.localdevice.apis.ClipboardManager.Read:output_type -> cloud.lazycat.localdevice.apis.ReadClipReply
	3, // 4: cloud.lazycat.localdevice.apis.ClipboardManager.Write:output_type -> cloud.lazycat.localdevice.apis.WriteClipReply
	1, // 5: cloud.lazycat.localdevice.apis.ClipboardManager.Watch:output_type -> cloud.lazycat.localdevice.apis.ReadClipReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_localdevice_clipboard_clipboard_proto_init() }
func file_localdevice_clipboard_clipboard_proto_init() {
	if File_localdevice_clipboard_clipboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_localdevice_clipboard_clipboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_localdevice_clipboard_clipboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadClipReply); i {
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
		file_localdevice_clipboard_clipboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_localdevice_clipboard_clipboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteClipReply); i {
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
			RawDescriptor: file_localdevice_clipboard_clipboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_localdevice_clipboard_clipboard_proto_goTypes,
		DependencyIndexes: file_localdevice_clipboard_clipboard_proto_depIdxs,
		MessageInfos:      file_localdevice_clipboard_clipboard_proto_msgTypes,
	}.Build()
	File_localdevice_clipboard_clipboard_proto = out.File
	file_localdevice_clipboard_clipboard_proto_rawDesc = nil
	file_localdevice_clipboard_clipboard_proto_goTypes = nil
	file_localdevice_clipboard_clipboard_proto_depIdxs = nil
}

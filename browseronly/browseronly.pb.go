// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.1
// source: browseronly/browseronly.proto

package browseronly

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string                 `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	DeviceId string                 `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	When     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=when,proto3" json:"when,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_browseronly_browseronly_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_browseronly_browseronly_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_browseronly_browseronly_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserInfo) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *UserInfo) GetWhen() *timestamppb.Timestamp {
	if x != nil {
		return x.When
	}
	return nil
}

type AppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoxId     string `protobuf:"bytes,1,opt,name=box_id,json=boxId,proto3" json:"box_id,omitempty"`
	AppId     string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppDomain string `protobuf:"bytes,3,opt,name=app_domain,json=appDomain,proto3" json:"app_domain,omitempty"`
}

func (x *AppInfo) Reset() {
	*x = AppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_browseronly_browseronly_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfo) ProtoMessage() {}

func (x *AppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_browseronly_browseronly_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfo.ProtoReflect.Descriptor instead.
func (*AppInfo) Descriptor() ([]byte, []int) {
	return file_browseronly_browseronly_proto_rawDescGZIP(), []int{1}
}

func (x *AppInfo) GetBoxId() string {
	if x != nil {
		return x.BoxId
	}
	return ""
}

func (x *AppInfo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *AppInfo) GetAppDomain() string {
	if x != nil {
		return x.AppDomain
	}
	return ""
}

var File_browseronly_browseronly_proto protoreflect.FileDescriptor

var file_browseronly_browseronly_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x79, 0x2f, 0x62, 0x72,
	0x6f, 0x77, 0x73, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x69, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04,
	0x77, 0x68, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x22, 0x56, 0x0a, 0x07,
	0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x6f, 0x78, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x32, 0xe8, 0x01, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72,
	0x4f, 0x6e, 0x6c, 0x79, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x47, 0x0a, 0x0d, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0e, 0x50, 0x61, 0x69,
	0x72, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x0e, 0x5a, 0x0c, 0x2f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x6f, 0x6e, 0x6c, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_browseronly_browseronly_proto_rawDescOnce sync.Once
	file_browseronly_browseronly_proto_rawDescData = file_browseronly_browseronly_proto_rawDesc
)

func file_browseronly_browseronly_proto_rawDescGZIP() []byte {
	file_browseronly_browseronly_proto_rawDescOnce.Do(func() {
		file_browseronly_browseronly_proto_rawDescData = protoimpl.X.CompressGZIP(file_browseronly_browseronly_proto_rawDescData)
	})
	return file_browseronly_browseronly_proto_rawDescData
}

var file_browseronly_browseronly_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_browseronly_browseronly_proto_goTypes = []interface{}{
	(*UserInfo)(nil),              // 0: cloud.lazycat.apis.UserInfo
	(*AppInfo)(nil),               // 1: cloud.lazycat.apis.AppInfo
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_browseronly_browseronly_proto_depIdxs = []int32{
	2, // 0: cloud.lazycat.apis.UserInfo.when:type_name -> google.protobuf.Timestamp
	3, // 1: cloud.lazycat.apis.BrowserOnlyProxy.QueryUserInfo:input_type -> google.protobuf.Empty
	3, // 2: cloud.lazycat.apis.BrowserOnlyProxy.QueryAppInfo:input_type -> google.protobuf.Empty
	3, // 3: cloud.lazycat.apis.BrowserOnlyProxy.PairAllDevices:input_type -> google.protobuf.Empty
	0, // 4: cloud.lazycat.apis.BrowserOnlyProxy.QueryUserInfo:output_type -> cloud.lazycat.apis.UserInfo
	1, // 5: cloud.lazycat.apis.BrowserOnlyProxy.QueryAppInfo:output_type -> cloud.lazycat.apis.AppInfo
	3, // 6: cloud.lazycat.apis.BrowserOnlyProxy.PairAllDevices:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_browseronly_browseronly_proto_init() }
func file_browseronly_browseronly_proto_init() {
	if File_browseronly_browseronly_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_browseronly_browseronly_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_browseronly_browseronly_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfo); i {
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
			RawDescriptor: file_browseronly_browseronly_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_browseronly_browseronly_proto_goTypes,
		DependencyIndexes: file_browseronly_browseronly_proto_depIdxs,
		MessageInfos:      file_browseronly_browseronly_proto_msgTypes,
	}.Build()
	File_browseronly_browseronly_proto = out.File
	file_browseronly_browseronly_proto_rawDesc = nil
	file_browseronly_browseronly_proto_goTypes = nil
	file_browseronly_browseronly_proto_depIdxs = nil
}

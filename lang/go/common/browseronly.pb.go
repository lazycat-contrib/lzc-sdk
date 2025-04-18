// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: common/browseronly.proto

package common

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

type APIServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrontendVersion string `protobuf:"bytes,1,opt,name=frontend_version,json=frontendVersion,proto3" json:"frontend_version,omitempty"`
	BackendVersion  string `protobuf:"bytes,2,opt,name=backend_version,json=backendVersion,proto3" json:"backend_version,omitempty"`
}

func (x *APIServerInfo) Reset() {
	*x = APIServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_browseronly_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIServerInfo) ProtoMessage() {}

func (x *APIServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_browseronly_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIServerInfo.ProtoReflect.Descriptor instead.
func (*APIServerInfo) Descriptor() ([]byte, []int) {
	return file_common_browseronly_proto_rawDescGZIP(), []int{0}
}

func (x *APIServerInfo) GetFrontendVersion() string {
	if x != nil {
		return x.FrontendVersion
	}
	return ""
}

func (x *APIServerInfo) GetBackendVersion() string {
	if x != nil {
		return x.BackendVersion
	}
	return ""
}

type SessionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前登陆用户
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 当前登陆设备
	DeviceId string                 `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	When     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=when,proto3" json:"when,omitempty"`
}

func (x *SessionInfo) Reset() {
	*x = SessionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_browseronly_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionInfo) ProtoMessage() {}

func (x *SessionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_browseronly_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionInfo.ProtoReflect.Descriptor instead.
func (*SessionInfo) Descriptor() ([]byte, []int) {
	return file_common_browseronly_proto_rawDescGZIP(), []int{1}
}

func (x *SessionInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *SessionInfo) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *SessionInfo) GetWhen() *timestamppb.Timestamp {
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
	// APIGateway对所有而http service服务强制要求设置http header类型的
	// AuthInfo. 但有些场景无法使用XHR去设置自定义的http header，导致访问
	// 这些服务很困难，因此由/usr/bin/lzcapp统一提供一个代理服务来设置这个
	// header. 此字段即是对应代理出来的URL前缀。默认为"/_lzc"，可以通过启动
	// 参数调整。
	HttpApiProxyPath string `protobuf:"bytes,4,opt,name=http_api_proxy_path,json=httpApiProxyPath,proto3" json:"http_api_proxy_path,omitempty"`
}

func (x *AppInfo) Reset() {
	*x = AppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_browseronly_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfo) ProtoMessage() {}

func (x *AppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_browseronly_proto_msgTypes[2]
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
	return file_common_browseronly_proto_rawDescGZIP(), []int{2}
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

func (x *AppInfo) GetHttpApiProxyPath() string {
	if x != nil {
		return x.HttpApiProxyPath
	}
	return ""
}

var File_common_browseronly_proto protoreflect.FileDescriptor

var file_common_browseronly_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72,
	0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0d, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x22, 0x85, 0x01, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x2d, 0x0a, 0x13, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x68, 0x74,
	0x74, 0x70, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x61, 0x74, 0x68, 0x32, 0xd6,
	0x02, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x12, 0x54, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0e, 0x50, 0x61, 0x69, 0x72, 0x41,
	0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x58, 0x0a,
	0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x50, 0x49, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x65, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a,
	0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_browseronly_proto_rawDescOnce sync.Once
	file_common_browseronly_proto_rawDescData = file_common_browseronly_proto_rawDesc
)

func file_common_browseronly_proto_rawDescGZIP() []byte {
	file_common_browseronly_proto_rawDescOnce.Do(func() {
		file_common_browseronly_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_browseronly_proto_rawDescData)
	})
	return file_common_browseronly_proto_rawDescData
}

var file_common_browseronly_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_browseronly_proto_goTypes = []interface{}{
	(*APIServerInfo)(nil),         // 0: cloud.lazycat.apis.common.APIServerInfo
	(*SessionInfo)(nil),           // 1: cloud.lazycat.apis.common.SessionInfo
	(*AppInfo)(nil),               // 2: cloud.lazycat.apis.common.AppInfo
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_common_browseronly_proto_depIdxs = []int32{
	3, // 0: cloud.lazycat.apis.common.SessionInfo.when:type_name -> google.protobuf.Timestamp
	4, // 1: cloud.lazycat.apis.common.BrowserOnlyProxy.QuerySessionInfo:input_type -> google.protobuf.Empty
	4, // 2: cloud.lazycat.apis.common.BrowserOnlyProxy.QueryAppInfo:input_type -> google.protobuf.Empty
	4, // 3: cloud.lazycat.apis.common.BrowserOnlyProxy.PairAllDevices:input_type -> google.protobuf.Empty
	4, // 4: cloud.lazycat.apis.common.BrowserOnlyProxy.QueryAPIServerInfo:input_type -> google.protobuf.Empty
	1, // 5: cloud.lazycat.apis.common.BrowserOnlyProxy.QuerySessionInfo:output_type -> cloud.lazycat.apis.common.SessionInfo
	2, // 6: cloud.lazycat.apis.common.BrowserOnlyProxy.QueryAppInfo:output_type -> cloud.lazycat.apis.common.AppInfo
	4, // 7: cloud.lazycat.apis.common.BrowserOnlyProxy.PairAllDevices:output_type -> google.protobuf.Empty
	0, // 8: cloud.lazycat.apis.common.BrowserOnlyProxy.QueryAPIServerInfo:output_type -> cloud.lazycat.apis.common.APIServerInfo
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_browseronly_proto_init() }
func file_common_browseronly_proto_init() {
	if File_common_browseronly_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_browseronly_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIServerInfo); i {
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
		file_common_browseronly_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionInfo); i {
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
		file_common_browseronly_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_common_browseronly_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_browseronly_proto_goTypes,
		DependencyIndexes: file_common_browseronly_proto_depIdxs,
		MessageInfos:      file_common_browseronly_proto_msgTypes,
	}.Build()
	File_common_browseronly_proto = out.File
	file_common_browseronly_proto_rawDesc = nil
	file_common_browseronly_proto_goTypes = nil
	file_common_browseronly_proto_depIdxs = nil
}

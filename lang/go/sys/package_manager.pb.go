// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: sys/package_manager.proto

package sys

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

type PkgURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 内部地址
	//     http://pkgm.api-server.lzcapp/tmp/xxxx-0.2.1.lpk
	// 或外网地址
	//     https://repo.lazycat.cloud/a/c/accc-0.2.1.lpk
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// pkg对应的sha512值，若不为空pm会对此进行cache
	Sha512 string `protobuf:"bytes,2,opt,name=sha512,proto3" json:"sha512,omitempty"`
}

func (x *PkgURL) Reset() {
	*x = PkgURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PkgURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PkgURL) ProtoMessage() {}

func (x *PkgURL) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PkgURL.ProtoReflect.Descriptor instead.
func (*PkgURL) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{0}
}

func (x *PkgURL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PkgURL) GetSha512() string {
	if x != nil {
		return x.Sha512
	}
	return ""
}

type PkgDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PkgDescription) Reset() {
	*x = PkgDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PkgDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PkgDescription) ProtoMessage() {}

func (x *PkgDescription) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PkgDescription.ProtoReflect.Descriptor instead.
func (*PkgDescription) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{1}
}

func (x *PkgDescription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AppDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Icon      string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	Appdomain string `protobuf:"bytes,5,opt,name=appdomain,proto3" json:"appdomain,omitempty"`
	Running   bool   `protobuf:"varint,6,opt,name=running,proto3" json:"running,omitempty"`
}

func (x *AppDescription) Reset() {
	*x = AppDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppDescription) ProtoMessage() {}

func (x *AppDescription) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppDescription.ProtoReflect.Descriptor instead.
func (*AppDescription) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{2}
}

func (x *AppDescription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppDescription) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AppDescription) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AppDescription) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *AppDescription) GetAppdomain() string {
	if x != nil {
		return x.Appdomain
	}
	return ""
}

func (x *AppDescription) GetRunning() bool {
	if x != nil {
		return x.Running
	}
	return false
}

var File_sys_package_manager_proto protoreflect.FileDescriptor

var file_sys_package_manager_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x79, 0x73, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x32, 0x0a, 0x06, 0x50, 0x6b, 0x67, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x68, 0x61, 0x35, 0x31, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68,
	0x61, 0x35, 0x31, 0x32, 0x22, 0x20, 0x0a, 0x0e, 0x50, 0x6b, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x70, 0x70, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x70, 0x70, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x32, 0xfb, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x12, 0x1e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x50, 0x6b, 0x67,
	0x55, 0x52, 0x4c, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a,
	0x09, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x26, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x2e, 0x50, 0x6b, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73,
	0x2e, 0x41, 0x70, 0x70, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x79, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_package_manager_proto_rawDescOnce sync.Once
	file_sys_package_manager_proto_rawDescData = file_sys_package_manager_proto_rawDesc
)

func file_sys_package_manager_proto_rawDescGZIP() []byte {
	file_sys_package_manager_proto_rawDescOnce.Do(func() {
		file_sys_package_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_package_manager_proto_rawDescData)
	})
	return file_sys_package_manager_proto_rawDescData
}

var file_sys_package_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sys_package_manager_proto_goTypes = []interface{}{
	(*PkgURL)(nil),         // 0: cloud.lazycat.apis.sys.PkgURL
	(*PkgDescription)(nil), // 1: cloud.lazycat.apis.sys.PkgDescription
	(*AppDescription)(nil), // 2: cloud.lazycat.apis.sys.AppDescription
	(*emptypb.Empty)(nil),  // 3: google.protobuf.Empty
}
var file_sys_package_manager_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.apis.sys.PackageManager.Install:input_type -> cloud.lazycat.apis.sys.PkgURL
	1, // 1: cloud.lazycat.apis.sys.PackageManager.Uninstall:input_type -> cloud.lazycat.apis.sys.PkgDescription
	3, // 2: cloud.lazycat.apis.sys.PackageManager.ListApplication:input_type -> google.protobuf.Empty
	3, // 3: cloud.lazycat.apis.sys.PackageManager.Install:output_type -> google.protobuf.Empty
	3, // 4: cloud.lazycat.apis.sys.PackageManager.Uninstall:output_type -> google.protobuf.Empty
	2, // 5: cloud.lazycat.apis.sys.PackageManager.ListApplication:output_type -> cloud.lazycat.apis.sys.AppDescription
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sys_package_manager_proto_init() }
func file_sys_package_manager_proto_init() {
	if File_sys_package_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_package_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PkgURL); i {
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
		file_sys_package_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PkgDescription); i {
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
		file_sys_package_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppDescription); i {
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
			RawDescriptor: file_sys_package_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_package_manager_proto_goTypes,
		DependencyIndexes: file_sys_package_manager_proto_depIdxs,
		MessageInfos:      file_sys_package_manager_proto_msgTypes,
	}.Build()
	File_sys_package_manager_proto = out.File
	file_sys_package_manager_proto_rawDesc = nil
	file_sys_package_manager_proto_goTypes = nil
	file_sys_package_manager_proto_depIdxs = nil
}

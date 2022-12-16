// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: localdevice/permission.proto

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

type Permission int32

const (
	// 打开相册需要的权限（读取相册和删除照片）
	Permission_PHOTO_LIBRAY_PERMISSION Permission = 0
	// 打开文件管理器需要的权限(读取文件，删除文件)
	Permission_DOCUMENET_PERMISSION Permission = 1
	// 打开第三方应用的权限(需要使用第三方app来打开文件）
	Permission_OPEN_THIRD_PARTY_APP Permission = 2
	// 将快捷方式放在桌面的权限
	Permission_PIN_APP_PERMISSIN Permission = 3
	// 剪贴板同步（读取、写入剪贴板）
	Permission_SYNC_CLIPBOARD Permission = 4
)

// Enum value maps for Permission.
var (
	Permission_name = map[int32]string{
		0: "PHOTO_LIBRAY_PERMISSION",
		1: "DOCUMENET_PERMISSION",
		2: "OPEN_THIRD_PARTY_APP",
		3: "PIN_APP_PERMISSIN",
		4: "SYNC_CLIPBOARD",
	}
	Permission_value = map[string]int32{
		"PHOTO_LIBRAY_PERMISSION": 0,
		"DOCUMENET_PERMISSION":    1,
		"OPEN_THIRD_PARTY_APP":    2,
		"PIN_APP_PERMISSIN":       3,
		"SYNC_CLIPBOARD":          4,
	}
)

func (x Permission) Enum() *Permission {
	p := new(Permission)
	*p = x
	return p
}

func (x Permission) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permission) Descriptor() protoreflect.EnumDescriptor {
	return file_localdevice_permission_proto_enumTypes[0].Descriptor()
}

func (Permission) Type() protoreflect.EnumType {
	return &file_localdevice_permission_proto_enumTypes[0]
}

func (x Permission) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permission.Descriptor instead.
func (Permission) EnumDescriptor() ([]byte, []int) {
	return file_localdevice_permission_proto_rawDescGZIP(), []int{0}
}

type CheckPermissionReply_PermissionResult int32

const (
	// 已经授予
	CheckPermissionReply_GRANTED CheckPermissionReply_PermissionResult = 0
	// 未授予
	CheckPermissionReply_NO_GRANTED CheckPermissionReply_PermissionResult = 1
	// 未知
	CheckPermissionReply_UNKNOWN CheckPermissionReply_PermissionResult = 2
)

// Enum value maps for CheckPermissionReply_PermissionResult.
var (
	CheckPermissionReply_PermissionResult_name = map[int32]string{
		0: "GRANTED",
		1: "NO_GRANTED",
		2: "UNKNOWN",
	}
	CheckPermissionReply_PermissionResult_value = map[string]int32{
		"GRANTED":    0,
		"NO_GRANTED": 1,
		"UNKNOWN":    2,
	}
)

func (x CheckPermissionReply_PermissionResult) Enum() *CheckPermissionReply_PermissionResult {
	p := new(CheckPermissionReply_PermissionResult)
	*p = x
	return p
}

func (x CheckPermissionReply_PermissionResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckPermissionReply_PermissionResult) Descriptor() protoreflect.EnumDescriptor {
	return file_localdevice_permission_proto_enumTypes[1].Descriptor()
}

func (CheckPermissionReply_PermissionResult) Type() protoreflect.EnumType {
	return &file_localdevice_permission_proto_enumTypes[1]
}

func (x CheckPermissionReply_PermissionResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckPermissionReply_PermissionResult.Descriptor instead.
func (CheckPermissionReply_PermissionResult) EnumDescriptor() ([]byte, []int) {
	return file_localdevice_permission_proto_rawDescGZIP(), []int{1, 0}
}

type CheckPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permission Permission `protobuf:"varint,1,opt,name=permission,proto3,enum=cloud.lazycat.apis.localdevice.Permission" json:"permission,omitempty"`
}

func (x *CheckPermissionRequest) Reset() {
	*x = CheckPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionRequest) ProtoMessage() {}

func (x *CheckPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckPermissionRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_permission_proto_rawDescGZIP(), []int{0}
}

func (x *CheckPermissionRequest) GetPermission() Permission {
	if x != nil {
		return x.Permission
	}
	return Permission_PHOTO_LIBRAY_PERMISSION
}

type CheckPermissionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reulst CheckPermissionReply_PermissionResult `protobuf:"varint,1,opt,name=reulst,proto3,enum=cloud.lazycat.apis.localdevice.CheckPermissionReply_PermissionResult" json:"reulst,omitempty"`
}

func (x *CheckPermissionReply) Reset() {
	*x = CheckPermissionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionReply) ProtoMessage() {}

func (x *CheckPermissionReply) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_permission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionReply.ProtoReflect.Descriptor instead.
func (*CheckPermissionReply) Descriptor() ([]byte, []int) {
	return file_localdevice_permission_proto_rawDescGZIP(), []int{1}
}

func (x *CheckPermissionReply) GetReulst() CheckPermissionReply_PermissionResult {
	if x != nil {
		return x.Reulst
	}
	return CheckPermissionReply_GRANTED
}

var File_localdevice_permission_proto protoreflect.FileDescriptor

var file_localdevice_permission_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x64,
	0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0xb3, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5d, 0x0a,
	0x06, 0x72, 0x65, 0x75, 0x6c, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x75, 0x6c, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x10,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x0b, 0x0a, 0x07, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x4f, 0x5f, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x2a, 0x88, 0x01, 0x0a, 0x0a, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x48, 0x4f,
	0x54, 0x4f, 0x5f, 0x4c, 0x49, 0x42, 0x52, 0x41, 0x59, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x4f, 0x43, 0x55, 0x4d, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01,
	0x12, 0x18, 0x0a, 0x14, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50,
	0x41, 0x52, 0x54, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x49,
	0x4e, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x10,
	0x03, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x43, 0x4c, 0x49, 0x50, 0x42, 0x4f,
	0x41, 0x52, 0x44, 0x10, 0x04, 0x32, 0x97, 0x01, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x81, 0x01, 0x0a, 0x0f,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x36, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e,
	0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_localdevice_permission_proto_rawDescOnce sync.Once
	file_localdevice_permission_proto_rawDescData = file_localdevice_permission_proto_rawDesc
)

func file_localdevice_permission_proto_rawDescGZIP() []byte {
	file_localdevice_permission_proto_rawDescOnce.Do(func() {
		file_localdevice_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_localdevice_permission_proto_rawDescData)
	})
	return file_localdevice_permission_proto_rawDescData
}

var file_localdevice_permission_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_localdevice_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_localdevice_permission_proto_goTypes = []interface{}{
	(Permission)(0), // 0: cloud.lazycat.apis.localdevice.Permission
	(CheckPermissionReply_PermissionResult)(0), // 1: cloud.lazycat.apis.localdevice.CheckPermissionReply.PermissionResult
	(*CheckPermissionRequest)(nil),             // 2: cloud.lazycat.apis.localdevice.CheckPermissionRequest
	(*CheckPermissionReply)(nil),               // 3: cloud.lazycat.apis.localdevice.CheckPermissionReply
}
var file_localdevice_permission_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.apis.localdevice.CheckPermissionRequest.permission:type_name -> cloud.lazycat.apis.localdevice.Permission
	1, // 1: cloud.lazycat.apis.localdevice.CheckPermissionReply.reulst:type_name -> cloud.lazycat.apis.localdevice.CheckPermissionReply.PermissionResult
	2, // 2: cloud.lazycat.apis.localdevice.PermissionManager.CheckPermission:input_type -> cloud.lazycat.apis.localdevice.CheckPermissionRequest
	3, // 3: cloud.lazycat.apis.localdevice.PermissionManager.CheckPermission:output_type -> cloud.lazycat.apis.localdevice.CheckPermissionReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_localdevice_permission_proto_init() }
func file_localdevice_permission_proto_init() {
	if File_localdevice_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_localdevice_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionRequest); i {
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
		file_localdevice_permission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionReply); i {
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
			RawDescriptor: file_localdevice_permission_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_localdevice_permission_proto_goTypes,
		DependencyIndexes: file_localdevice_permission_proto_depIdxs,
		EnumInfos:         file_localdevice_permission_proto_enumTypes,
		MessageInfos:      file_localdevice_permission_proto_msgTypes,
	}.Build()
	File_localdevice_permission_proto = out.File
	file_localdevice_permission_proto_rawDesc = nil
	file_localdevice_permission_proto_goTypes = nil
	file_localdevice_permission_proto_depIdxs = nil
}

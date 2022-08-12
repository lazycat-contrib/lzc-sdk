// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: common/security_context.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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
	// 是否允许挂载用户文档数据
	Permission_USER_DOCUMENT Permission = 0
	// 是否允许安装软件
	Permission_INSTALL_PACKAGE Permission = 1
	// 是否允许挂载可移动存储设备
	Permission_ACCESS_REMOVEABLE_STOREAGE Permission = 2
	// 是否允许生成访问EndDevice的访问凭证, EndDevice会进行具体API的权限验证
	Permission_ACCESS_ENDDEVICE       Permission = 1000
	Permission_ENDDEVICE_CLIPBOARD    Permission = 1001
	Permission_ENDDEVICE_NETWORK_INFO Permission = 1002
	Permission_ENDDEVICE_PHOTOLIBRARY Permission = 1003
)

// Enum value maps for Permission.
var (
	Permission_name = map[int32]string{
		0:    "USER_DOCUMENT",
		1:    "INSTALL_PACKAGE",
		2:    "ACCESS_REMOVEABLE_STOREAGE",
		1000: "ACCESS_ENDDEVICE",
		1001: "ENDDEVICE_CLIPBOARD",
		1002: "ENDDEVICE_NETWORK_INFO",
		1003: "ENDDEVICE_PHOTOLIBRARY",
	}
	Permission_value = map[string]int32{
		"USER_DOCUMENT":              0,
		"INSTALL_PACKAGE":            1,
		"ACCESS_REMOVEABLE_STOREAGE": 2,
		"ACCESS_ENDDEVICE":           1000,
		"ENDDEVICE_CLIPBOARD":        1001,
		"ENDDEVICE_NETWORK_INFO":     1002,
		"ENDDEVICE_PHOTOLIBRARY":     1003,
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
	return file_common_security_context_proto_enumTypes[0].Descriptor()
}

func (Permission) Type() protoreflect.EnumType {
	return &file_common_security_context_proto_enumTypes[0]
}

func (x Permission) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permission.Descriptor instead.
func (Permission) EnumDescriptor() ([]byte, []int) {
	return file_common_security_context_proto_rawDescGZIP(), []int{0}
}

type UserClass int32

const (
	//可以没有用户(比如一个纯后端app)或任何有效用户
	UserClass_ANY_USER UserClass = 0
	//任何有效的用户
	UserClass_NORMAL_USER UserClass = 1
	//必须是管理员用户
	UserClass_ADMIN_USER UserClass = 2
	//请求者必须是资源所有者对应的用户，比如"设置密码"每个用户都有权设置，但仅能设置自身的
	//具体的TARGET_USER通过input message中的target_uid_field option进行标记
	UserClass_TARGET_USER UserClass = 3
)

// Enum value maps for UserClass.
var (
	UserClass_name = map[int32]string{
		0: "ANY_USER",
		1: "NORMAL_USER",
		2: "ADMIN_USER",
		3: "TARGET_USER",
	}
	UserClass_value = map[string]int32{
		"ANY_USER":    0,
		"NORMAL_USER": 1,
		"ADMIN_USER":  2,
		"TARGET_USER": 3,
	}
)

func (x UserClass) Enum() *UserClass {
	p := new(UserClass)
	*p = x
	return p
}

func (x UserClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserClass) Descriptor() protoreflect.EnumDescriptor {
	return file_common_security_context_proto_enumTypes[1].Descriptor()
}

func (UserClass) Type() protoreflect.EnumType {
	return &file_common_security_context_proto_enumTypes[1]
}

func (x UserClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserClass.Descriptor instead.
func (UserClass) EnumDescriptor() ([]byte, []int) {
	return file_common_security_context_proto_rawDescGZIP(), []int{1}
}

type SecurityContextMetaKey int32

const (
	SecurityContextMetaKey_META_KEY_X_LZCAPI_APPID   SecurityContextMetaKey = 0
	SecurityContextMetaKey_META_KEY_X_LZCAPI_REALUID SecurityContextMetaKey = 1
	SecurityContextMetaKey_META_KEY_X_LZCAPI_UID     SecurityContextMetaKey = 2
)

// Enum value maps for SecurityContextMetaKey.
var (
	SecurityContextMetaKey_name = map[int32]string{
		0: "META_KEY_X_LZCAPI_APPID",
		1: "META_KEY_X_LZCAPI_REALUID",
		2: "META_KEY_X_LZCAPI_UID",
	}
	SecurityContextMetaKey_value = map[string]int32{
		"META_KEY_X_LZCAPI_APPID":   0,
		"META_KEY_X_LZCAPI_REALUID": 1,
		"META_KEY_X_LZCAPI_UID":     2,
	}
)

func (x SecurityContextMetaKey) Enum() *SecurityContextMetaKey {
	p := new(SecurityContextMetaKey)
	*p = x
	return p
}

func (x SecurityContextMetaKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityContextMetaKey) Descriptor() protoreflect.EnumDescriptor {
	return file_common_security_context_proto_enumTypes[2].Descriptor()
}

func (SecurityContextMetaKey) Type() protoreflect.EnumType {
	return &file_common_security_context_proto_enumTypes[2]
}

func (x SecurityContextMetaKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityContextMetaKey.Descriptor instead.
func (SecurityContextMetaKey) EnumDescriptor() ([]byte, []int) {
	return file_common_security_context_proto_rawDescGZIP(), []int{2}
}

// APIGateway在进行完权限检测放行后，转发给API Services时会通过metadata附带以下信息
type SecurityContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 请求来源的真实appid，一定是合法有效的
	Appid string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"` //key=X_LZCAPI_APPID
	// 请求来源的真实uid,一定是合法有效的
	RealUid string `protobuf:"bytes,2,opt,name=real_uid,json=realUid,proto3" json:"real_uid,omitempty"` //key=X_LZCAPI_REALUID
	// 请求来源期望的uid，注意此值可以被app server篡改
	// 若uid != real_uid且对应SecurityContextRequire.disable_behalf_of_uid == true
	// 则此请求会被APIGATEWAY阻止转发
	// 若app server没有篡改此值，则此值与real_uid相同
	Uid string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"` //key=X_LZCAPI_UID
}

func (x *SecurityContext) Reset() {
	*x = SecurityContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_security_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityContext) ProtoMessage() {}

func (x *SecurityContext) ProtoReflect() protoreflect.Message {
	mi := &file_common_security_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityContext.ProtoReflect.Descriptor instead.
func (*SecurityContext) Descriptor() ([]byte, []int) {
	return file_common_security_context_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityContext) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *SecurityContext) GetRealUid() string {
	if x != nil {
		return x.RealUid
	}
	return ""
}

func (x *SecurityContext) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

// 以下权限要求通过option的方式声明在proto文件内
// 并由APIGateway进行验证
type SecurityContextRequire struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequirePermissions []Permission `protobuf:"varint,1,rep,packed,name=require_permissions,json=requirePermissions,proto3,enum=cloud.lazycat.apis.common.Permission" json:"require_permissions,omitempty"`
	RequireUserClass   UserClass    `protobuf:"varint,2,opt,name=require_user_class,json=requireUserClass,proto3,enum=cloud.lazycat.apis.common.UserClass" json:"require_user_class,omitempty"`
	// 请求必须直接通过app server发起，而非serverless端发起
	RequireFromBackend bool `protobuf:"varint,3,opt,name=require_from_backend,json=requireFromBackend,proto3" json:"require_from_backend,omitempty"`
	// 禁止app server伪造uid
	DisableBehalfOfUid bool `protobuf:"varint,4,opt,name=disable_behalf_of_uid,json=disableBehalfOfUid,proto3" json:"disable_behalf_of_uid,omitempty"`
}

func (x *SecurityContextRequire) Reset() {
	*x = SecurityContextRequire{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_security_context_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityContextRequire) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityContextRequire) ProtoMessage() {}

func (x *SecurityContextRequire) ProtoReflect() protoreflect.Message {
	mi := &file_common_security_context_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityContextRequire.ProtoReflect.Descriptor instead.
func (*SecurityContextRequire) Descriptor() ([]byte, []int) {
	return file_common_security_context_proto_rawDescGZIP(), []int{1}
}

func (x *SecurityContextRequire) GetRequirePermissions() []Permission {
	if x != nil {
		return x.RequirePermissions
	}
	return nil
}

func (x *SecurityContextRequire) GetRequireUserClass() UserClass {
	if x != nil {
		return x.RequireUserClass
	}
	return UserClass_ANY_USER
}

func (x *SecurityContextRequire) GetRequireFromBackend() bool {
	if x != nil {
		return x.RequireFromBackend
	}
	return false
}

func (x *SecurityContextRequire) GetDisableBehalfOfUid() bool {
	if x != nil {
		return x.DisableBehalfOfUid
	}
	return false
}

type PermissionDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 申请权限的原因说明
	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	// 权限列表
	Permissions []Permission `protobuf:"varint,2,rep,packed,name=permissions,proto3,enum=cloud.lazycat.apis.common.Permission" json:"permissions,omitempty"`
}

func (x *PermissionDesc) Reset() {
	*x = PermissionDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_security_context_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionDesc) ProtoMessage() {}

func (x *PermissionDesc) ProtoReflect() protoreflect.Message {
	mi := &file_common_security_context_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionDesc.ProtoReflect.Descriptor instead.
func (*PermissionDesc) Descriptor() ([]byte, []int) {
	return file_common_security_context_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionDesc) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *PermissionDesc) GetPermissions() []Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// 后续调用lzc-apis时，需要带上此token
type PermissionToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PermissionToken) Reset() {
	*x = PermissionToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_security_context_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionToken) ProtoMessage() {}

func (x *PermissionToken) ProtoReflect() protoreflect.Message {
	mi := &file_common_security_context_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionToken.ProtoReflect.Descriptor instead.
func (*PermissionToken) Descriptor() ([]byte, []int) {
	return file_common_security_context_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PermissionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PermissionStatus) Reset() {
	*x = PermissionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_security_context_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionStatus) ProtoMessage() {}

func (x *PermissionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_common_security_context_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionStatus.ProtoReflect.Descriptor instead.
func (*PermissionStatus) Descriptor() ([]byte, []int) {
	return file_common_security_context_proto_rawDescGZIP(), []int{4}
}

type HasPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Perm Permission `protobuf:"varint,1,opt,name=perm,proto3,enum=cloud.lazycat.apis.common.Permission" json:"perm,omitempty"`
}

func (x *HasPermissionRequest) Reset() {
	*x = HasPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_security_context_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasPermissionRequest) ProtoMessage() {}

func (x *HasPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_security_context_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasPermissionRequest.ProtoReflect.Descriptor instead.
func (*HasPermissionRequest) Descriptor() ([]byte, []int) {
	return file_common_security_context_proto_rawDescGZIP(), []int{5}
}

func (x *HasPermissionRequest) GetPerm() Permission {
	if x != nil {
		return x.Perm
	}
	return Permission_USER_DOCUMENT
}

var file_common_security_context_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*SecurityContextRequire)(nil),
		Field:         70000,
		Name:          "cloud.lazycat.apis.common.scontext",
		Tag:           "bytes,70000,opt,name=scontext",
		Filename:      "common/security_context.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         70001,
		Name:          "cloud.lazycat.apis.common.target_uid_field",
		Tag:           "varint,70001,opt,name=target_uid_field",
		Filename:      "common/security_context.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional cloud.lazycat.apis.common.SecurityContextRequire scontext = 70000;
	E_Scontext = &file_common_security_context_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional int32 target_uid_field = 70001;
	E_TargetUidField = &file_common_security_context_proto_extTypes[1]
)

var File_common_security_context_proto protoreflect.FileDescriptor

var file_common_security_context_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0f,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x55, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x22, 0xa9, 0x02, 0x0a, 0x16, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x12, 0x56, 0x0a,
	0x13, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x52, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x6c, 0x66, 0x5f, 0x6f, 0x66,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x42, 0x65, 0x68, 0x61, 0x6c, 0x66, 0x4f, 0x66, 0x55, 0x69, 0x64, 0x22, 0x71,
	0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x25, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x27, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x51,
	0x0a, 0x14, 0x48, 0x61, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x70, 0x65, 0x72, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x70, 0x65, 0x72,
	0x6d, 0x2a, 0xbf, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x5f, 0x50,
	0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x53, 0x54,
	0x4f, 0x52, 0x45, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x41, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x5f, 0x45, 0x4e, 0x44, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0xe8, 0x07, 0x12,
	0x18, 0x0a, 0x13, 0x45, 0x4e, 0x44, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x4c, 0x49,
	0x50, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0xe9, 0x07, 0x12, 0x1b, 0x0a, 0x16, 0x45, 0x4e, 0x44,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x49,
	0x4e, 0x46, 0x4f, 0x10, 0xea, 0x07, 0x12, 0x1b, 0x0a, 0x16, 0x45, 0x4e, 0x44, 0x44, 0x45, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x50, 0x48, 0x4f, 0x54, 0x4f, 0x4c, 0x49, 0x42, 0x52, 0x41, 0x52, 0x59,
	0x10, 0xeb, 0x07, 0x2a, 0x4b, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x12, 0x0c, 0x0a, 0x08, 0x41, 0x4e, 0x59, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x03,
	0x2a, 0x6f, 0x0a, 0x16, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45,
	0x54, 0x41, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x58, 0x5f, 0x4c, 0x5a, 0x43, 0x41, 0x50, 0x49, 0x5f,
	0x41, 0x50, 0x50, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x54, 0x41, 0x5f,
	0x4b, 0x45, 0x59, 0x5f, 0x58, 0x5f, 0x4c, 0x5a, 0x43, 0x41, 0x50, 0x49, 0x5f, 0x52, 0x45, 0x41,
	0x4c, 0x55, 0x49, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x4b,
	0x45, 0x59, 0x5f, 0x58, 0x5f, 0x4c, 0x5a, 0x43, 0x41, 0x50, 0x49, 0x5f, 0x55, 0x49, 0x44, 0x10,
	0x02, 0x32, 0xde, 0x01, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x62, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x1a, 0x2a, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x03, 0x48,
	0x61, 0x73, 0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48,
	0x61, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x3a, 0x72, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf0,
	0xa2, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x52, 0x08, 0x73, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x3a, 0x4e, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x75, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf1, 0xa2, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_security_context_proto_rawDescOnce sync.Once
	file_common_security_context_proto_rawDescData = file_common_security_context_proto_rawDesc
)

func file_common_security_context_proto_rawDescGZIP() []byte {
	file_common_security_context_proto_rawDescOnce.Do(func() {
		file_common_security_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_security_context_proto_rawDescData)
	})
	return file_common_security_context_proto_rawDescData
}

var file_common_security_context_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_security_context_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_security_context_proto_goTypes = []interface{}{
	(Permission)(0),                     // 0: cloud.lazycat.apis.common.Permission
	(UserClass)(0),                      // 1: cloud.lazycat.apis.common.UserClass
	(SecurityContextMetaKey)(0),         // 2: cloud.lazycat.apis.common.SecurityContextMetaKey
	(*SecurityContext)(nil),             // 3: cloud.lazycat.apis.common.SecurityContext
	(*SecurityContextRequire)(nil),      // 4: cloud.lazycat.apis.common.SecurityContextRequire
	(*PermissionDesc)(nil),              // 5: cloud.lazycat.apis.common.PermissionDesc
	(*PermissionToken)(nil),             // 6: cloud.lazycat.apis.common.PermissionToken
	(*PermissionStatus)(nil),            // 7: cloud.lazycat.apis.common.PermissionStatus
	(*HasPermissionRequest)(nil),        // 8: cloud.lazycat.apis.common.HasPermissionRequest
	(*descriptorpb.MethodOptions)(nil),  // 9: google.protobuf.MethodOptions
	(*descriptorpb.MessageOptions)(nil), // 10: google.protobuf.MessageOptions
}
var file_common_security_context_proto_depIdxs = []int32{
	0,  // 0: cloud.lazycat.apis.common.SecurityContextRequire.require_permissions:type_name -> cloud.lazycat.apis.common.Permission
	1,  // 1: cloud.lazycat.apis.common.SecurityContextRequire.require_user_class:type_name -> cloud.lazycat.apis.common.UserClass
	0,  // 2: cloud.lazycat.apis.common.PermissionDesc.permissions:type_name -> cloud.lazycat.apis.common.Permission
	0,  // 3: cloud.lazycat.apis.common.HasPermissionRequest.perm:type_name -> cloud.lazycat.apis.common.Permission
	9,  // 4: cloud.lazycat.apis.common.scontext:extendee -> google.protobuf.MethodOptions
	10, // 5: cloud.lazycat.apis.common.target_uid_field:extendee -> google.protobuf.MessageOptions
	4,  // 6: cloud.lazycat.apis.common.scontext:type_name -> cloud.lazycat.apis.common.SecurityContextRequire
	5,  // 7: cloud.lazycat.apis.common.PermissionManager.Request:input_type -> cloud.lazycat.apis.common.PermissionDesc
	8,  // 8: cloud.lazycat.apis.common.PermissionManager.Has:input_type -> cloud.lazycat.apis.common.HasPermissionRequest
	6,  // 9: cloud.lazycat.apis.common.PermissionManager.Request:output_type -> cloud.lazycat.apis.common.PermissionToken
	7,  // 10: cloud.lazycat.apis.common.PermissionManager.Has:output_type -> cloud.lazycat.apis.common.PermissionStatus
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	6,  // [6:7] is the sub-list for extension type_name
	4,  // [4:6] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_common_security_context_proto_init() }
func file_common_security_context_proto_init() {
	if File_common_security_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_security_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityContext); i {
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
		file_common_security_context_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityContextRequire); i {
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
		file_common_security_context_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionDesc); i {
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
		file_common_security_context_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionToken); i {
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
		file_common_security_context_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionStatus); i {
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
		file_common_security_context_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasPermissionRequest); i {
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
			RawDescriptor: file_common_security_context_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 2,
			NumServices:   1,
		},
		GoTypes:           file_common_security_context_proto_goTypes,
		DependencyIndexes: file_common_security_context_proto_depIdxs,
		EnumInfos:         file_common_security_context_proto_enumTypes,
		MessageInfos:      file_common_security_context_proto_msgTypes,
		ExtensionInfos:    file_common_security_context_proto_extTypes,
	}.Build()
	File_common_security_context_proto = out.File
	file_common_security_context_proto_rawDesc = nil
	file_common_security_context_proto_goTypes = nil
	file_common_security_context_proto_depIdxs = nil
}

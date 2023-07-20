// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: localdevice/permission.proto

package localdevice

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

type Permission int32

const (
	// 剪贴板
	Permission_CLIPBOARD Permission = 0
	// 设备信息
	Permission_DEVICE_INFO Permission = 1
	// 弹出对话框
	Permission_OPEN_DIALOG Permission = 2
	// 使用第三方app打开文件
	Permission_OPEN_THIRD_PARTY_APP Permission = 3
	// 发送应用到桌面快捷方式
	Permission_PIN_APP Permission = 4
	// 网络信息
	Permission_NETWORK_INFO Permission = 5
	// 相册
	Permission_PHOTO_LIBRARY Permission = 6
	// 文件
	Permission_DOCUMENT Permission = 7
	// 用户配置
	Permission_USER_CONFIG Permission = 8
)

// Enum value maps for Permission.
var (
	Permission_name = map[int32]string{
		0: "CLIPBOARD",
		1: "DEVICE_INFO",
		2: "OPEN_DIALOG",
		3: "OPEN_THIRD_PARTY_APP",
		4: "PIN_APP",
		5: "NETWORK_INFO",
		6: "PHOTO_LIBRARY",
		7: "DOCUMENT",
		8: "USER_CONFIG",
	}
	Permission_value = map[string]int32{
		"CLIPBOARD":            0,
		"DEVICE_INFO":          1,
		"OPEN_DIALOG":          2,
		"OPEN_THIRD_PARTY_APP": 3,
		"PIN_APP":              4,
		"NETWORK_INFO":         5,
		"PHOTO_LIBRARY":        6,
		"DOCUMENT":             7,
		"USER_CONFIG":          8,
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

type PermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permission Permission `protobuf:"varint,1,opt,name=permission,proto3,enum=cloud.lazycat.apis.localdevice.Permission" json:"permission,omitempty"`
}

func (x *PermissionRequest) Reset() {
	*x = PermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionRequest) ProtoMessage() {}

func (x *PermissionRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PermissionRequest.ProtoReflect.Descriptor instead.
func (*PermissionRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_permission_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionRequest) GetPermission() Permission {
	if x != nil {
		return x.Permission
	}
	return Permission_CLIPBOARD
}

type PermissionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PermissionReply) Reset() {
	*x = PermissionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionReply) ProtoMessage() {}

func (x *PermissionReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PermissionReply.ProtoReflect.Descriptor instead.
func (*PermissionReply) Descriptor() ([]byte, []int) {
	return file_localdevice_permission_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionReply) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type ListPermissionsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// map<Permission, bool>
	Result map[int32]bool `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ListPermissionsReply) Reset() {
	*x = ListPermissionsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_permission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPermissionsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionsReply) ProtoMessage() {}

func (x *ListPermissionsReply) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_permission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionsReply.ProtoReflect.Descriptor instead.
func (*ListPermissionsReply) Descriptor() ([]byte, []int) {
	return file_localdevice_permission_proto_rawDescGZIP(), []int{2}
}

func (x *ListPermissionsReply) GetResult() map[int32]bool {
	if x != nil {
		return x.Result
	}
	return nil
}

type RequestAuthTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoxCert []byte `protobuf:"bytes,1,opt,name=box_cert,json=boxCert,proto3" json:"box_cert,omitempty"`
	AppCert []byte `protobuf:"bytes,2,opt,name=app_cert,json=appCert,proto3" json:"app_cert,omitempty"`
	// signature of $LAZYCAT_APPID.$LAZYCAT_BOX_DOMAIN
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *RequestAuthTokenRequest) Reset() {
	*x = RequestAuthTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_permission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAuthTokenRequest) ProtoMessage() {}

func (x *RequestAuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_permission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAuthTokenRequest.ProtoReflect.Descriptor instead.
func (*RequestAuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_permission_proto_rawDescGZIP(), []int{3}
}

func (x *RequestAuthTokenRequest) GetBoxCert() []byte {
	if x != nil {
		return x.BoxCert
	}
	return nil
}

func (x *RequestAuthTokenRequest) GetAppCert() []byte {
	if x != nil {
		return x.AppCert
	}
	return nil
}

func (x *RequestAuthTokenRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type RequestAuthTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Deadline *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *RequestAuthTokenResponse) Reset() {
	*x = RequestAuthTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_permission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAuthTokenResponse) ProtoMessage() {}

func (x *RequestAuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_permission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAuthTokenResponse.ProtoReflect.Descriptor instead.
func (*RequestAuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_localdevice_permission_proto_rawDescGZIP(), []int{4}
}

func (x *RequestAuthTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RequestAuthTokenResponse) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

var File_localdevice_permission_proto protoreflect.FileDescriptor

var file_localdevice_permission_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x11,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4a, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61,
	0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x29, 0x0a,
	0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xab, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x58, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x40, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6d, 0x0a, 0x17, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x78, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6f, 0x78, 0x43, 0x65, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x70, 0x70, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x43, 0x65, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x68, 0x0a, 0x18, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x2a,
	0xa8, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4c, 0x49, 0x50, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x4f, 0x47, 0x10, 0x02, 0x12,
	0x18, 0x0a, 0x14, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41,
	0x52, 0x54, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x49, 0x4e,
	0x5f, 0x41, 0x50, 0x50, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x48, 0x4f, 0x54,
	0x4f, 0x5f, 0x4c, 0x49, 0x42, 0x52, 0x41, 0x52, 0x59, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x08, 0x32, 0xf2, 0x03, 0x0a, 0x11, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x75, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x61, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x34, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x87, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
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

var file_localdevice_permission_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_localdevice_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_localdevice_permission_proto_goTypes = []interface{}{
	(Permission)(0),                  // 0: cloud.lazycat.apis.localdevice.Permission
	(*PermissionRequest)(nil),        // 1: cloud.lazycat.apis.localdevice.PermissionRequest
	(*PermissionReply)(nil),          // 2: cloud.lazycat.apis.localdevice.PermissionReply
	(*ListPermissionsReply)(nil),     // 3: cloud.lazycat.apis.localdevice.ListPermissionsReply
	(*RequestAuthTokenRequest)(nil),  // 4: cloud.lazycat.apis.localdevice.RequestAuthTokenRequest
	(*RequestAuthTokenResponse)(nil), // 5: cloud.lazycat.apis.localdevice.RequestAuthTokenResponse
	nil,                              // 6: cloud.lazycat.apis.localdevice.ListPermissionsReply.ResultEntry
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_localdevice_permission_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.apis.localdevice.PermissionRequest.permission:type_name -> cloud.lazycat.apis.localdevice.Permission
	6, // 1: cloud.lazycat.apis.localdevice.ListPermissionsReply.result:type_name -> cloud.lazycat.apis.localdevice.ListPermissionsReply.ResultEntry
	7, // 2: cloud.lazycat.apis.localdevice.RequestAuthTokenResponse.deadline:type_name -> google.protobuf.Timestamp
	1, // 3: cloud.lazycat.apis.localdevice.PermissionManager.GetPermission:input_type -> cloud.lazycat.apis.localdevice.PermissionRequest
	1, // 4: cloud.lazycat.apis.localdevice.PermissionManager.RequestPermission:input_type -> cloud.lazycat.apis.localdevice.PermissionRequest
	8, // 5: cloud.lazycat.apis.localdevice.PermissionManager.ListPermissions:input_type -> google.protobuf.Empty
	4, // 6: cloud.lazycat.apis.localdevice.PermissionManager.RequestAuthToken:input_type -> cloud.lazycat.apis.localdevice.RequestAuthTokenRequest
	2, // 7: cloud.lazycat.apis.localdevice.PermissionManager.GetPermission:output_type -> cloud.lazycat.apis.localdevice.PermissionReply
	2, // 8: cloud.lazycat.apis.localdevice.PermissionManager.RequestPermission:output_type -> cloud.lazycat.apis.localdevice.PermissionReply
	3, // 9: cloud.lazycat.apis.localdevice.PermissionManager.ListPermissions:output_type -> cloud.lazycat.apis.localdevice.ListPermissionsReply
	5, // 10: cloud.lazycat.apis.localdevice.PermissionManager.RequestAuthToken:output_type -> cloud.lazycat.apis.localdevice.RequestAuthTokenResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_localdevice_permission_proto_init() }
func file_localdevice_permission_proto_init() {
	if File_localdevice_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_localdevice_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionRequest); i {
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
			switch v := v.(*PermissionReply); i {
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
		file_localdevice_permission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPermissionsReply); i {
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
		file_localdevice_permission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAuthTokenRequest); i {
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
		file_localdevice_permission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAuthTokenResponse); i {
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
			NumEnums:      1,
			NumMessages:   6,
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

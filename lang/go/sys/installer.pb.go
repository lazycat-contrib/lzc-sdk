// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: sys/installer.proto

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

type BoxSetupReply_Status int32

const (
	// 等待按钮点击
	BoxSetupReply_PendingButtonClick BoxSetupReply_Status = 0
	// 初始化中，进行申请证书/域名等操作
	BoxSetupReply_Initializing BoxSetupReply_Status = 1
	// 失败
	BoxSetupReply_Failed BoxSetupReply_Status = 2
	// 成功
	BoxSetupReply_Success BoxSetupReply_Status = 3
)

// Enum value maps for BoxSetupReply_Status.
var (
	BoxSetupReply_Status_name = map[int32]string{
		0: "PendingButtonClick",
		1: "Initializing",
		2: "Failed",
		3: "Success",
	}
	BoxSetupReply_Status_value = map[string]int32{
		"PendingButtonClick": 0,
		"Initializing":       1,
		"Failed":             2,
		"Success":            3,
	}
)

func (x BoxSetupReply_Status) Enum() *BoxSetupReply_Status {
	p := new(BoxSetupReply_Status)
	*p = x
	return p
}

func (x BoxSetupReply_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BoxSetupReply_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_sys_installer_proto_enumTypes[0].Descriptor()
}

func (BoxSetupReply_Status) Type() protoreflect.EnumType {
	return &file_sys_installer_proto_enumTypes[0]
}

func (x BoxSetupReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BoxSetupReply_Status.Descriptor instead.
func (BoxSetupReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_sys_installer_proto_rawDescGZIP(), []int{1, 0}
}

type BoxSetupReply_FailedStatus int32

const (
	// 等待点击按钮超时
	BoxSetupReply_WaitButtonTimeout BoxSetupReply_FailedStatus = 0
	// 初始化失败
	BoxSetupReply_InitializeFailed BoxSetupReply_FailedStatus = 1
	// Wi-Fi密码错误
	BoxSetupReply_WifiAuthFailed BoxSetupReply_FailedStatus = 2
)

// Enum value maps for BoxSetupReply_FailedStatus.
var (
	BoxSetupReply_FailedStatus_name = map[int32]string{
		0: "WaitButtonTimeout",
		1: "InitializeFailed",
		2: "WifiAuthFailed",
	}
	BoxSetupReply_FailedStatus_value = map[string]int32{
		"WaitButtonTimeout": 0,
		"InitializeFailed":  1,
		"WifiAuthFailed":    2,
	}
)

func (x BoxSetupReply_FailedStatus) Enum() *BoxSetupReply_FailedStatus {
	p := new(BoxSetupReply_FailedStatus)
	*p = x
	return p
}

func (x BoxSetupReply_FailedStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BoxSetupReply_FailedStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_sys_installer_proto_enumTypes[1].Descriptor()
}

func (BoxSetupReply_FailedStatus) Type() protoreflect.EnumType {
	return &file_sys_installer_proto_enumTypes[1]
}

func (x BoxSetupReply_FailedStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BoxSetupReply_FailedStatus.Descriptor instead.
func (BoxSetupReply_FailedStatus) EnumDescriptor() ([]byte, []int) {
	return file_sys_installer_proto_rawDescGZIP(), []int{1, 1}
}

type BoxSetupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginServer string                    `protobuf:"bytes,1,opt,name=origin_server,json=originServer,proto3" json:"origin_server,omitempty"`
	BoxName      string                    `protobuf:"bytes,2,opt,name=box_name,json=boxName,proto3" json:"box_name,omitempty"`
	UserInfo     *BoxSetupRequest_UserInfo `protobuf:"bytes,3,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *BoxSetupRequest) Reset() {
	*x = BoxSetupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_installer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoxSetupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoxSetupRequest) ProtoMessage() {}

func (x *BoxSetupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_installer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoxSetupRequest.ProtoReflect.Descriptor instead.
func (*BoxSetupRequest) Descriptor() ([]byte, []int) {
	return file_sys_installer_proto_rawDescGZIP(), []int{0}
}

func (x *BoxSetupRequest) GetOriginServer() string {
	if x != nil {
		return x.OriginServer
	}
	return ""
}

func (x *BoxSetupRequest) GetBoxName() string {
	if x != nil {
		return x.BoxName
	}
	return ""
}

func (x *BoxSetupRequest) GetUserInfo() *BoxSetupRequest_UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type BoxSetupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       BoxSetupReply_Status        `protobuf:"varint,1,opt,name=status,proto3,enum=cloud.lazycat.apis.sys.BoxSetupReply_Status" json:"status,omitempty"`
	FailedStatus *BoxSetupReply_FailedStatus `protobuf:"varint,2,opt,name=failed_status,json=failedStatus,proto3,enum=cloud.lazycat.apis.sys.BoxSetupReply_FailedStatus,oneof" json:"failed_status,omitempty"`
	// 如果是InitializeFailed，那么可能有如下失败原因
	// 微服名被别人注册了
	// 微服名不符合规范
	// 账号密码不符合规范
	// 其他的失败就直接用其字面意思就行
	FailedReason *string `protobuf:"bytes,3,opt,name=failed_reason,json=failedReason,proto3,oneof" json:"failed_reason,omitempty"`
}

func (x *BoxSetupReply) Reset() {
	*x = BoxSetupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_installer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoxSetupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoxSetupReply) ProtoMessage() {}

func (x *BoxSetupReply) ProtoReflect() protoreflect.Message {
	mi := &file_sys_installer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoxSetupReply.ProtoReflect.Descriptor instead.
func (*BoxSetupReply) Descriptor() ([]byte, []int) {
	return file_sys_installer_proto_rawDescGZIP(), []int{1}
}

func (x *BoxSetupReply) GetStatus() BoxSetupReply_Status {
	if x != nil {
		return x.Status
	}
	return BoxSetupReply_PendingButtonClick
}

func (x *BoxSetupReply) GetFailedStatus() BoxSetupReply_FailedStatus {
	if x != nil && x.FailedStatus != nil {
		return *x.FailedStatus
	}
	return BoxSetupReply_WaitButtonTimeout
}

func (x *BoxSetupReply) GetFailedReason() string {
	if x != nil && x.FailedReason != nil {
		return *x.FailedReason
	}
	return ""
}

type HasInternetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *HasInternetResponse) Reset() {
	*x = HasInternetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_installer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasInternetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasInternetResponse) ProtoMessage() {}

func (x *HasInternetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_installer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasInternetResponse.ProtoReflect.Descriptor instead.
func (*HasInternetResponse) Descriptor() ([]byte, []int) {
	return file_sys_installer_proto_rawDescGZIP(), []int{2}
}

func (x *HasInternetResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type BoxSetupRequest_UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string  `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password    string  `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Phonenumber *string `protobuf:"bytes,3,opt,name=phonenumber,proto3,oneof" json:"phonenumber,omitempty"`
}

func (x *BoxSetupRequest_UserInfo) Reset() {
	*x = BoxSetupRequest_UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_installer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoxSetupRequest_UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoxSetupRequest_UserInfo) ProtoMessage() {}

func (x *BoxSetupRequest_UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sys_installer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoxSetupRequest_UserInfo.ProtoReflect.Descriptor instead.
func (*BoxSetupRequest_UserInfo) Descriptor() ([]byte, []int) {
	return file_sys_installer_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BoxSetupRequest_UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *BoxSetupRequest_UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *BoxSetupRequest_UserInfo) GetPhonenumber() string {
	if x != nil && x.Phonenumber != nil {
		return *x.Phonenumber
	}
	return ""
}

var File_sys_installer_proto protoreflect.FileDescriptor

var file_sys_installer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x79, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x79, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x0f, 0x42, 0x6f, 0x78, 0x53, 0x65, 0x74,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x6f, 0x78, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x6f, 0x78, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42, 0x6f, 0x78, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x79, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x9f, 0x03, 0x0a, 0x0d, 0x42, 0x6f, 0x78, 0x53, 0x65, 0x74, 0x75, 0x70,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61,
	0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42,
	0x6f, 0x78, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5c, 0x0a, 0x0d, 0x66,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42, 0x6f, 0x78, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x66, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x22, 0x4b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x43, 0x6c,
	0x69, 0x63, 0x6b, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x03,
	0x22, 0x4f, 0x0a, 0x0c, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x15, 0x0a, 0x11, 0x57, 0x61, 0x69, 0x74, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x12, 0x12, 0x0a,
	0x0e, 0x57, 0x69, 0x66, 0x69, 0x41, 0x75, 0x74, 0x68, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10,
	0x02, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x13, 0x48, 0x61, 0x73, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0xb3, 0x03, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x08, 0x42, 0x6f, 0x78,
	0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61,
	0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42,
	0x6f, 0x78, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42, 0x6f, 0x78, 0x53, 0x65, 0x74, 0x75, 0x70,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x12, 0x52, 0x0a, 0x0b, 0x48, 0x61, 0x73, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2b,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x48, 0x61, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x57,
	0x69, 0x66, 0x69, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x60, 0x0a, 0x0b,
	0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x27, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x73, 0x79, 0x73, 0x2e, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x57, 0x69,
	0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a,
	0x0a, 0x08, 0x57, 0x69, 0x66, 0x69, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69,
	0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x79, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_installer_proto_rawDescOnce sync.Once
	file_sys_installer_proto_rawDescData = file_sys_installer_proto_rawDesc
)

func file_sys_installer_proto_rawDescGZIP() []byte {
	file_sys_installer_proto_rawDescOnce.Do(func() {
		file_sys_installer_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_installer_proto_rawDescData)
	})
	return file_sys_installer_proto_rawDescData
}

var file_sys_installer_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sys_installer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sys_installer_proto_goTypes = []interface{}{
	(BoxSetupReply_Status)(0),        // 0: cloud.lazycat.apis.sys.BoxSetupReply.Status
	(BoxSetupReply_FailedStatus)(0),  // 1: cloud.lazycat.apis.sys.BoxSetupReply.FailedStatus
	(*BoxSetupRequest)(nil),          // 2: cloud.lazycat.apis.sys.BoxSetupRequest
	(*BoxSetupReply)(nil),            // 3: cloud.lazycat.apis.sys.BoxSetupReply
	(*HasInternetResponse)(nil),      // 4: cloud.lazycat.apis.sys.HasInternetResponse
	(*BoxSetupRequest_UserInfo)(nil), // 5: cloud.lazycat.apis.sys.BoxSetupRequest.UserInfo
	(*emptypb.Empty)(nil),            // 6: google.protobuf.Empty
	(*WifiConnectInfo)(nil),          // 7: cloud.lazycat.apis.sys.WifiConnectInfo
	(*AccessPointInfoList)(nil),      // 8: cloud.lazycat.apis.sys.AccessPointInfoList
	(*WifiConnectReply)(nil),         // 9: cloud.lazycat.apis.sys.WifiConnectReply
}
var file_sys_installer_proto_depIdxs = []int32{
	5, // 0: cloud.lazycat.apis.sys.BoxSetupRequest.user_info:type_name -> cloud.lazycat.apis.sys.BoxSetupRequest.UserInfo
	0, // 1: cloud.lazycat.apis.sys.BoxSetupReply.status:type_name -> cloud.lazycat.apis.sys.BoxSetupReply.Status
	1, // 2: cloud.lazycat.apis.sys.BoxSetupReply.failed_status:type_name -> cloud.lazycat.apis.sys.BoxSetupReply.FailedStatus
	2, // 3: cloud.lazycat.apis.sys.InstallerService.BoxSetup:input_type -> cloud.lazycat.apis.sys.BoxSetupRequest
	6, // 4: cloud.lazycat.apis.sys.InstallerService.HasInternet:input_type -> google.protobuf.Empty
	6, // 5: cloud.lazycat.apis.sys.InstallerService.WifiList:input_type -> google.protobuf.Empty
	7, // 6: cloud.lazycat.apis.sys.InstallerService.WifiConnect:input_type -> cloud.lazycat.apis.sys.WifiConnectInfo
	6, // 7: cloud.lazycat.apis.sys.InstallerService.WifiScan:input_type -> google.protobuf.Empty
	3, // 8: cloud.lazycat.apis.sys.InstallerService.BoxSetup:output_type -> cloud.lazycat.apis.sys.BoxSetupReply
	4, // 9: cloud.lazycat.apis.sys.InstallerService.HasInternet:output_type -> cloud.lazycat.apis.sys.HasInternetResponse
	8, // 10: cloud.lazycat.apis.sys.InstallerService.WifiList:output_type -> cloud.lazycat.apis.sys.AccessPointInfoList
	9, // 11: cloud.lazycat.apis.sys.InstallerService.WifiConnect:output_type -> cloud.lazycat.apis.sys.WifiConnectReply
	6, // 12: cloud.lazycat.apis.sys.InstallerService.WifiScan:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sys_installer_proto_init() }
func file_sys_installer_proto_init() {
	if File_sys_installer_proto != nil {
		return
	}
	file_sys_network_manager_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sys_installer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoxSetupRequest); i {
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
		file_sys_installer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoxSetupReply); i {
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
		file_sys_installer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasInternetResponse); i {
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
		file_sys_installer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoxSetupRequest_UserInfo); i {
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
	file_sys_installer_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_sys_installer_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sys_installer_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_installer_proto_goTypes,
		DependencyIndexes: file_sys_installer_proto_depIdxs,
		EnumInfos:         file_sys_installer_proto_enumTypes,
		MessageInfos:      file_sys_installer_proto_msgTypes,
	}.Build()
	File_sys_installer_proto = out.File
	file_sys_installer_proto_rawDesc = nil
	file_sys_installer_proto_goTypes = nil
	file_sys_installer_proto_depIdxs = nil
}

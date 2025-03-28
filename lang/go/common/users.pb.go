// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: common/users.proto

package common

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

type Role int32

const (
	Role_ROLE_NORMAL Role = 0
	Role_ROLE_ADMIN  Role = 1
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_NORMAL",
		1: "ROLE_ADMIN",
	}
	Role_value = map[string]int32{
		"ROLE_NORMAL": 0,
		"ROLE_ADMIN":  1,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_common_users_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_common_users_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{0}
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{0}
}

func (x *UserID) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type CheckPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CheckPasswordRequest) Reset() {
	*x = CheckPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPasswordRequest) ProtoMessage() {}

func (x *CheckPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPasswordRequest.ProtoReflect.Descriptor instead.
func (*CheckPasswordRequest) Descriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{1}
}

func (x *CheckPasswordRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CheckPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid                     string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname                string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar                  string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Role                    Role     `protobuf:"varint,4,opt,name=role,proto3,enum=cloud.lazycat.apis.common.Role" json:"role,omitempty"`
	PhoneNumber             []string `protobuf:"bytes,5,rep,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	HasAppInstallPermission bool     `protobuf:"varint,6,opt,name=has_app_install_permission,json=hasAppInstallPermission,proto3" json:"has_app_install_permission,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[2]
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
	return file_common_users_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserInfo) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_NORMAL
}

func (x *UserInfo) GetPhoneNumber() []string {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

func (x *UserInfo) GetHasAppInstallPermission() bool {
	if x != nil {
		return x.HasAppInstallPermission
	}
	return false
}

type ListUIDsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUIDsRequest) Reset() {
	*x = ListUIDsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUIDsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUIDsRequest) ProtoMessage() {}

func (x *ListUIDsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUIDsRequest.ProtoReflect.Descriptor instead.
func (*ListUIDsRequest) Descriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{3}
}

type ListUIDsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uids []string `protobuf:"bytes,1,rep,name=uids,proto3" json:"uids,omitempty"`
}

func (x *ListUIDsReply) Reset() {
	*x = ListUIDsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUIDsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUIDsReply) ProtoMessage() {}

func (x *ListUIDsReply) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUIDsReply.ProtoReflect.Descriptor instead.
func (*ListUIDsReply) Descriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{4}
}

func (x *ListUIDsReply) GetUids() []string {
	if x != nil {
		return x.Uids
	}
	return nil
}

type ChangeRoleReqeust struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Role Role   `protobuf:"varint,2,opt,name=role,proto3,enum=cloud.lazycat.apis.common.Role" json:"role,omitempty"`
}

func (x *ChangeRoleReqeust) Reset() {
	*x = ChangeRoleReqeust{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeRoleReqeust) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeRoleReqeust) ProtoMessage() {}

func (x *ChangeRoleReqeust) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeRoleReqeust.ProtoReflect.Descriptor instead.
func (*ChangeRoleReqeust) Descriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{5}
}

func (x *ChangeRoleReqeust) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ChangeRoleReqeust) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_NORMAL
}

type ResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	OldPassword string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword string `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *ResetPasswordRequest) Reset() {
	*x = ResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordRequest) ProtoMessage() {}

func (x *ResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{6}
}

func (x *ResetPasswordRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ResetPasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ResetPasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 是否清理指定uid用户的数据目录
	ClearUserData bool `protobuf:"varint,2,opt,name=clear_user_data,json=clearUserData,proto3" json:"clear_user_data,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *DeleteUserRequest) GetClearUserData() bool {
	if x != nil {
		return x.ClearUserData
	}
	return false
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role     Role   `protobuf:"varint,3,opt,name=role,proto3,enum=cloud.lazycat.apis.common.Role" json:"role,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{8}
}

func (x *CreateUserRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_NORMAL
}

type UpdateUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname    string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar      string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	PhoneNumber []string `protobuf:"bytes,4,rep,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *UpdateUserInfoRequest) Reset() {
	*x = UpdateUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoRequest) ProtoMessage() {}

func (x *UpdateUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateUserInfoRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UpdateUserInfoRequest) GetPhoneNumber() []string {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

type ForceResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *ForceResetPasswordRequest) Reset() {
	*x = ForceResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_users_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceResetPasswordRequest) ProtoMessage() {}

func (x *ForceResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_users_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ForceResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_common_users_proto_rawDescGZIP(), []int{10}
}

func (x *ForceResetPasswordRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ForceResetPasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

var File_common_users_proto protoreflect.FileDescriptor

var file_common_users_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x3a, 0x04, 0x88, 0x97, 0x22, 0x01, 0x22, 0x44, 0x0a,
	0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0xe5, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x33, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3b,
	0x0a, 0x1a, 0x68, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x17, 0x68, 0x61, 0x73, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x11, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x23,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x69, 0x64, 0x73, 0x22, 0x5a, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x65, 0x75, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22,
	0x6e, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x6c, 0x64,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x4d, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x76,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x33, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x19, 0x46, 0x6f, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2a, 0x27, 0x0a, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d,
	0x41, 0x4c, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x10, 0x01, 0x32, 0x80, 0x07, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x49, 0x44, 0x73,
	0x12, 0x2a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x49, 0x44,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x06, 0x82, 0x97, 0x22, 0x02, 0x10, 0x01, 0x12, 0x5f,
	0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x21, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x1a, 0x23, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x06, 0x82, 0x97, 0x22, 0x02, 0x10, 0x01, 0x12,
	0x62, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0x82, 0x97, 0x22,
	0x02, 0x10, 0x03, 0x12, 0x5a, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x65, 0x75, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0x82, 0x97, 0x22, 0x02, 0x10, 0x02, 0x12,
	0x60, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0x82, 0x97, 0x22, 0x02, 0x10,
	0x03, 0x12, 0x5a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0x82, 0x97, 0x22, 0x02, 0x10, 0x02, 0x12, 0x5a, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x06, 0x82, 0x97, 0x22, 0x02, 0x10, 0x02, 0x12, 0x6a, 0x0a, 0x12, 0x46, 0x6f, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x34, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x6f, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0x82,
	0x97, 0x22, 0x02, 0x10, 0x02, 0x12, 0x60, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x06, 0x82, 0x97, 0x22, 0x02, 0x10, 0x03, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x65, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a,
	0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_users_proto_rawDescOnce sync.Once
	file_common_users_proto_rawDescData = file_common_users_proto_rawDesc
)

func file_common_users_proto_rawDescGZIP() []byte {
	file_common_users_proto_rawDescOnce.Do(func() {
		file_common_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_users_proto_rawDescData)
	})
	return file_common_users_proto_rawDescData
}

var file_common_users_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_users_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_common_users_proto_goTypes = []interface{}{
	(Role)(0),                         // 0: cloud.lazycat.apis.common.Role
	(*UserID)(nil),                    // 1: cloud.lazycat.apis.common.UserID
	(*CheckPasswordRequest)(nil),      // 2: cloud.lazycat.apis.common.CheckPasswordRequest
	(*UserInfo)(nil),                  // 3: cloud.lazycat.apis.common.UserInfo
	(*ListUIDsRequest)(nil),           // 4: cloud.lazycat.apis.common.ListUIDsRequest
	(*ListUIDsReply)(nil),             // 5: cloud.lazycat.apis.common.ListUIDsReply
	(*ChangeRoleReqeust)(nil),         // 6: cloud.lazycat.apis.common.ChangeRoleReqeust
	(*ResetPasswordRequest)(nil),      // 7: cloud.lazycat.apis.common.ResetPasswordRequest
	(*DeleteUserRequest)(nil),         // 8: cloud.lazycat.apis.common.DeleteUserRequest
	(*CreateUserRequest)(nil),         // 9: cloud.lazycat.apis.common.CreateUserRequest
	(*UpdateUserInfoRequest)(nil),     // 10: cloud.lazycat.apis.common.UpdateUserInfoRequest
	(*ForceResetPasswordRequest)(nil), // 11: cloud.lazycat.apis.common.ForceResetPasswordRequest
	(*emptypb.Empty)(nil),             // 12: google.protobuf.Empty
}
var file_common_users_proto_depIdxs = []int32{
	0,  // 0: cloud.lazycat.apis.common.UserInfo.role:type_name -> cloud.lazycat.apis.common.Role
	0,  // 1: cloud.lazycat.apis.common.ChangeRoleReqeust.role:type_name -> cloud.lazycat.apis.common.Role
	0,  // 2: cloud.lazycat.apis.common.CreateUserRequest.role:type_name -> cloud.lazycat.apis.common.Role
	4,  // 3: cloud.lazycat.apis.common.UserManager.ListUIDs:input_type -> cloud.lazycat.apis.common.ListUIDsRequest
	1,  // 4: cloud.lazycat.apis.common.UserManager.QueryUserInfo:input_type -> cloud.lazycat.apis.common.UserID
	10, // 5: cloud.lazycat.apis.common.UserManager.UpdateUserInfo:input_type -> cloud.lazycat.apis.common.UpdateUserInfoRequest
	6,  // 6: cloud.lazycat.apis.common.UserManager.ChangeRole:input_type -> cloud.lazycat.apis.common.ChangeRoleReqeust
	7,  // 7: cloud.lazycat.apis.common.UserManager.ResetPassword:input_type -> cloud.lazycat.apis.common.ResetPasswordRequest
	8,  // 8: cloud.lazycat.apis.common.UserManager.DeleteUser:input_type -> cloud.lazycat.apis.common.DeleteUserRequest
	9,  // 9: cloud.lazycat.apis.common.UserManager.CreateUser:input_type -> cloud.lazycat.apis.common.CreateUserRequest
	11, // 10: cloud.lazycat.apis.common.UserManager.ForceResetPassword:input_type -> cloud.lazycat.apis.common.ForceResetPasswordRequest
	2,  // 11: cloud.lazycat.apis.common.UserManager.CheckPassword:input_type -> cloud.lazycat.apis.common.CheckPasswordRequest
	5,  // 12: cloud.lazycat.apis.common.UserManager.ListUIDs:output_type -> cloud.lazycat.apis.common.ListUIDsReply
	3,  // 13: cloud.lazycat.apis.common.UserManager.QueryUserInfo:output_type -> cloud.lazycat.apis.common.UserInfo
	12, // 14: cloud.lazycat.apis.common.UserManager.UpdateUserInfo:output_type -> google.protobuf.Empty
	12, // 15: cloud.lazycat.apis.common.UserManager.ChangeRole:output_type -> google.protobuf.Empty
	12, // 16: cloud.lazycat.apis.common.UserManager.ResetPassword:output_type -> google.protobuf.Empty
	12, // 17: cloud.lazycat.apis.common.UserManager.DeleteUser:output_type -> google.protobuf.Empty
	12, // 18: cloud.lazycat.apis.common.UserManager.CreateUser:output_type -> google.protobuf.Empty
	12, // 19: cloud.lazycat.apis.common.UserManager.ForceResetPassword:output_type -> google.protobuf.Empty
	12, // 20: cloud.lazycat.apis.common.UserManager.CheckPassword:output_type -> google.protobuf.Empty
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_common_users_proto_init() }
func file_common_users_proto_init() {
	if File_common_users_proto != nil {
		return
	}
	file_common_security_context_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_common_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPasswordRequest); i {
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
		file_common_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUIDsRequest); i {
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
		file_common_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUIDsReply); i {
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
		file_common_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeRoleReqeust); i {
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
		file_common_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordRequest); i {
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
		file_common_users_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_common_users_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_common_users_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoRequest); i {
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
		file_common_users_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceResetPasswordRequest); i {
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
			RawDescriptor: file_common_users_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_users_proto_goTypes,
		DependencyIndexes: file_common_users_proto_depIdxs,
		EnumInfos:         file_common_users_proto_enumTypes,
		MessageInfos:      file_common_users_proto_msgTypes,
	}.Build()
	File_common_users_proto = out.File
	file_common_users_proto_rawDesc = nil
	file_common_users_proto_goTypes = nil
	file_common_users_proto_depIdxs = nil
}

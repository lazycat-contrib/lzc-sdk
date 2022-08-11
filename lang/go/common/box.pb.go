// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: common/box.proto

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

type ShutdownRequest_Action int32

const (
	ShutdownRequest_Poweroff    ShutdownRequest_Action = 0
	ShutdownRequest_Halt        ShutdownRequest_Action = 1
	ShutdownRequest_Reboot      ShutdownRequest_Action = 2
	ShutdownRequest_HybridSleep ShutdownRequest_Action = 3
)

// Enum value maps for ShutdownRequest_Action.
var (
	ShutdownRequest_Action_name = map[int32]string{
		0: "Poweroff",
		1: "Halt",
		2: "Reboot",
		3: "HybridSleep",
	}
	ShutdownRequest_Action_value = map[string]int32{
		"Poweroff":    0,
		"Halt":        1,
		"Reboot":      2,
		"HybridSleep": 3,
	}
)

func (x ShutdownRequest_Action) Enum() *ShutdownRequest_Action {
	p := new(ShutdownRequest_Action)
	*p = x
	return p
}

func (x ShutdownRequest_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShutdownRequest_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_common_box_proto_enumTypes[0].Descriptor()
}

func (ShutdownRequest_Action) Type() protoreflect.EnumType {
	return &file_common_box_proto_enumTypes[0]
}

func (x ShutdownRequest_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShutdownRequest_Action.Descriptor instead.
func (ShutdownRequest_Action) EnumDescriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{2, 0}
}

type BoxInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 硬件唯一ID,出厂后则不再变化
	Udid string `protobuf:"bytes,1,opt,name=udid,proto3" json:"udid,omitempty"`
	// 盒子的唯一名称，此名称是向中心化服务器注册产生，创建后不可修改
	BoxName string `protobuf:"bytes,2,opt,name=box_name,json=boxName,proto3" json:"box_name,omitempty"`
	// 盒子的唯一域名
	BoxDomain string `protobuf:"bytes,3,opt,name=box_domain,json=boxDomain,proto3" json:"box_domain,omitempty"`
	// 盒子内部显示名称，可以随意修改，若不设置则为box_name
	DisplayName string             `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	DiskSpace   *BoxInfo_DiskSpace `protobuf:"bytes,5,opt,name=disk_space,json=diskSpace,proto3" json:"disk_space,omitempty"`
}

func (x *BoxInfo) Reset() {
	*x = BoxInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_box_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoxInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoxInfo) ProtoMessage() {}

func (x *BoxInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_box_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoxInfo.ProtoReflect.Descriptor instead.
func (*BoxInfo) Descriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{0}
}

func (x *BoxInfo) GetUdid() string {
	if x != nil {
		return x.Udid
	}
	return ""
}

func (x *BoxInfo) GetBoxName() string {
	if x != nil {
		return x.BoxName
	}
	return ""
}

func (x *BoxInfo) GetBoxDomain() string {
	if x != nil {
		return x.BoxDomain
	}
	return ""
}

func (x *BoxInfo) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *BoxInfo) GetDiskSpace() *BoxInfo_DiskSpace {
	if x != nil {
		return x.DiskSpace
	}
	return nil
}

type ChangeDisplayNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *ChangeDisplayNameRequest) Reset() {
	*x = ChangeDisplayNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_box_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeDisplayNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeDisplayNameRequest) ProtoMessage() {}

func (x *ChangeDisplayNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_box_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeDisplayNameRequest.ProtoReflect.Descriptor instead.
func (*ChangeDisplayNameRequest) Descriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeDisplayNameRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type ShutdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action  ShutdownRequest_Action `protobuf:"varint,1,opt,name=action,proto3,enum=cloud.lazycat.apis.common.ShutdownRequest_Action" json:"action,omitempty"`
	Reasone string                 `protobuf:"bytes,2,opt,name=reasone,proto3" json:"reasone,omitempty"`
}

func (x *ShutdownRequest) Reset() {
	*x = ShutdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_box_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownRequest) ProtoMessage() {}

func (x *ShutdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_box_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownRequest.ProtoReflect.Descriptor instead.
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{2}
}

func (x *ShutdownRequest) GetAction() ShutdownRequest_Action {
	if x != nil {
		return x.Action
	}
	return ShutdownRequest_Poweroff
}

func (x *ShutdownRequest) GetReasone() string {
	if x != nil {
		return x.Reasone
	}
	return ""
}

type BoxInfo_DiskSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Free  int64 `protobuf:"varint,2,opt,name=free,proto3" json:"free,omitempty"`
}

func (x *BoxInfo_DiskSpace) Reset() {
	*x = BoxInfo_DiskSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_box_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoxInfo_DiskSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoxInfo_DiskSpace) ProtoMessage() {}

func (x *BoxInfo_DiskSpace) ProtoReflect() protoreflect.Message {
	mi := &file_common_box_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoxInfo_DiskSpace.ProtoReflect.Descriptor instead.
func (*BoxInfo_DiskSpace) Descriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BoxInfo_DiskSpace) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *BoxInfo_DiskSpace) GetFree() int64 {
	if x != nil {
		return x.Free
	}
	return 0
}

var File_common_box_proto protoreflect.FileDescriptor

var file_common_box_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x07, 0x42,
	0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x64, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x64, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f,
	0x78, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f,
	0x78, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x78, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x78, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x5f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x1a, 0x35, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x65, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x65, 0x65, 0x22, 0x3d, 0x0a, 0x18, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x0f, 0x53,
	0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64,
	0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x65, 0x22, 0x3d, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a,
	0x08, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x6f, 0x66, 0x66, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48,
	0x61, 0x6c, 0x74, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x79, 0x62, 0x72, 0x69, 0x64, 0x53, 0x6c, 0x65, 0x65, 0x70,
	0x10, 0x03, 0x32, 0x8d, 0x02, 0x0a, 0x0a, 0x42, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x49, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x11,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x33, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x2a, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_box_proto_rawDescOnce sync.Once
	file_common_box_proto_rawDescData = file_common_box_proto_rawDesc
)

func file_common_box_proto_rawDescGZIP() []byte {
	file_common_box_proto_rawDescOnce.Do(func() {
		file_common_box_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_box_proto_rawDescData)
	})
	return file_common_box_proto_rawDescData
}

var file_common_box_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_box_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_box_proto_goTypes = []interface{}{
	(ShutdownRequest_Action)(0),      // 0: cloud.lazycat.apis.common.ShutdownRequest.Action
	(*BoxInfo)(nil),                  // 1: cloud.lazycat.apis.common.BoxInfo
	(*ChangeDisplayNameRequest)(nil), // 2: cloud.lazycat.apis.common.ChangeDisplayNameRequest
	(*ShutdownRequest)(nil),          // 3: cloud.lazycat.apis.common.ShutdownRequest
	(*BoxInfo_DiskSpace)(nil),        // 4: cloud.lazycat.apis.common.BoxInfo.DiskSpace
	(*emptypb.Empty)(nil),            // 5: google.protobuf.Empty
}
var file_common_box_proto_depIdxs = []int32{
	4, // 0: cloud.lazycat.apis.common.BoxInfo.disk_space:type_name -> cloud.lazycat.apis.common.BoxInfo.DiskSpace
	0, // 1: cloud.lazycat.apis.common.ShutdownRequest.action:type_name -> cloud.lazycat.apis.common.ShutdownRequest.Action
	5, // 2: cloud.lazycat.apis.common.BoxService.QueryInfo:input_type -> google.protobuf.Empty
	2, // 3: cloud.lazycat.apis.common.BoxService.ChangeDisplayName:input_type -> cloud.lazycat.apis.common.ChangeDisplayNameRequest
	3, // 4: cloud.lazycat.apis.common.BoxService.Shutdown:input_type -> cloud.lazycat.apis.common.ShutdownRequest
	1, // 5: cloud.lazycat.apis.common.BoxService.QueryInfo:output_type -> cloud.lazycat.apis.common.BoxInfo
	5, // 6: cloud.lazycat.apis.common.BoxService.ChangeDisplayName:output_type -> google.protobuf.Empty
	5, // 7: cloud.lazycat.apis.common.BoxService.Shutdown:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_box_proto_init() }
func file_common_box_proto_init() {
	if File_common_box_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_box_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoxInfo); i {
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
		file_common_box_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeDisplayNameRequest); i {
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
		file_common_box_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownRequest); i {
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
		file_common_box_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoxInfo_DiskSpace); i {
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
			RawDescriptor: file_common_box_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_box_proto_goTypes,
		DependencyIndexes: file_common_box_proto_depIdxs,
		EnumInfos:         file_common_box_proto_enumTypes,
		MessageInfos:      file_common_box_proto_msgTypes,
	}.Build()
	File_common_box_proto = out.File
	file_common_box_proto_rawDesc = nil
	file_common_box_proto_goTypes = nil
	file_common_box_proto_depIdxs = nil
}

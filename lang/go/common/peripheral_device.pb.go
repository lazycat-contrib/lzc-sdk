// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: common/peripheral_device.proto

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

type ListRemovableDiskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//当前APP未挂载的磁盘列表
	Umounted []*RemovableFilesystemPartion `protobuf:"bytes,1,rep,name=umounted,proto3" json:"umounted,omitempty"`
	//当前APP已挂载的磁盘列表。
	Mounted []*RemovableFilesystemPartion `protobuf:"bytes,2,rep,name=mounted,proto3" json:"mounted,omitempty"`
}

func (x *ListRemovableDiskReply) Reset() {
	*x = ListRemovableDiskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRemovableDiskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRemovableDiskReply) ProtoMessage() {}

func (x *ListRemovableDiskReply) ProtoReflect() protoreflect.Message {
	mi := &file_common_peripheral_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRemovableDiskReply.ProtoReflect.Descriptor instead.
func (*ListRemovableDiskReply) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{0}
}

func (x *ListRemovableDiskReply) GetUmounted() []*RemovableFilesystemPartion {
	if x != nil {
		return x.Umounted
	}
	return nil
}

func (x *ListRemovableDiskReply) GetMounted() []*RemovableFilesystemPartion {
	if x != nil {
		return x.Mounted
	}
	return nil
}

type RemovableFilesystemPartion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// partion's uuid.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// size unit in bytes
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// ntfs/fat32/ext4 ...
	Fstype string `protobuf:"bytes,3,opt,name=fstype,proto3" json:"fstype,omitempty"`
	// partion label name or other meaningful name for user
	Name *string `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *RemovableFilesystemPartion) Reset() {
	*x = RemovableFilesystemPartion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovableFilesystemPartion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovableFilesystemPartion) ProtoMessage() {}

func (x *RemovableFilesystemPartion) ProtoReflect() protoreflect.Message {
	mi := &file_common_peripheral_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovableFilesystemPartion.ProtoReflect.Descriptor instead.
func (*RemovableFilesystemPartion) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{1}
}

func (x *RemovableFilesystemPartion) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RemovableFilesystemPartion) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *RemovableFilesystemPartion) GetFstype() string {
	if x != nil {
		return x.Fstype
	}
	return ""
}

func (x *RemovableFilesystemPartion) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type MountFilesystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *MountFilesystemRequest) Reset() {
	*x = MountFilesystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountFilesystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountFilesystemRequest) ProtoMessage() {}

func (x *MountFilesystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_peripheral_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountFilesystemRequest.ProtoReflect.Descriptor instead.
func (*MountFilesystemRequest) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{2}
}

func (x *MountFilesystemRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type UmountFilesystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//	*UmountFilesystemRequest_Uuid
	//	*UmountFilesystemRequest_MountPoint
	Target isUmountFilesystemRequest_Target `protobuf_oneof:"target"`
}

func (x *UmountFilesystemRequest) Reset() {
	*x = UmountFilesystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UmountFilesystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UmountFilesystemRequest) ProtoMessage() {}

func (x *UmountFilesystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_peripheral_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UmountFilesystemRequest.ProtoReflect.Descriptor instead.
func (*UmountFilesystemRequest) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{3}
}

func (m *UmountFilesystemRequest) GetTarget() isUmountFilesystemRequest_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *UmountFilesystemRequest) GetUuid() string {
	if x, ok := x.GetTarget().(*UmountFilesystemRequest_Uuid); ok {
		return x.Uuid
	}
	return ""
}

func (x *UmountFilesystemRequest) GetMountPoint() string {
	if x, ok := x.GetTarget().(*UmountFilesystemRequest_MountPoint); ok {
		return x.MountPoint
	}
	return ""
}

type isUmountFilesystemRequest_Target interface {
	isUmountFilesystemRequest_Target()
}

type UmountFilesystemRequest_Uuid struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3,oneof"`
}

type UmountFilesystemRequest_MountPoint struct {
	MountPoint string `protobuf:"bytes,2,opt,name=mount_point,json=mountPoint,proto3,oneof"`
}

func (*UmountFilesystemRequest_Uuid) isUmountFilesystemRequest_Target() {}

func (*UmountFilesystemRequest_MountPoint) isUmountFilesystemRequest_Target() {}

type WebDavAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *WebDavAuth) Reset() {
	*x = WebDavAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebDavAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebDavAuth) ProtoMessage() {}

func (x *WebDavAuth) ProtoReflect() protoreflect.Message {
	mi := &file_common_peripheral_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebDavAuth.ProtoReflect.Descriptor instead.
func (*WebDavAuth) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{4}
}

func (x *WebDavAuth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *WebDavAuth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type MountWebDavRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string      `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Auth *WebDavAuth `protobuf:"bytes,2,opt,name=auth,proto3,oneof" json:"auth,omitempty"`
}

func (x *MountWebDavRequest) Reset() {
	*x = MountWebDavRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountWebDavRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountWebDavRequest) ProtoMessage() {}

func (x *MountWebDavRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_peripheral_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountWebDavRequest.ProtoReflect.Descriptor instead.
func (*MountWebDavRequest) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{5}
}

func (x *MountWebDavRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MountWebDavRequest) GetAuth() *WebDavAuth {
	if x != nil {
		return x.Auth
	}
	return nil
}

var File_common_peripheral_device_proto protoreflect.FileDescriptor

var file_common_peripheral_device_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x51, 0x0a, 0x08, 0x75, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x75, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x4f, 0x0a, 0x07, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x22, 0x7e, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61,
	0x62, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x73, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x73,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x16, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x17, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x22, 0x44, 0x0a, 0x0a, 0x57, 0x65, 0x62, 0x44, 0x61, 0x76, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6f, 0x0a, 0x12, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x57, 0x65, 0x62, 0x44, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x3e, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x62, 0x44, 0x61,
	0x76, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x9b, 0x03, 0x0a, 0x17, 0x50, 0x65,
	0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0f, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x31, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0x82, 0x97, 0x22, 0x02, 0x08, 0x00, 0x12, 0x60, 0x0a,
	0x10, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x12, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x56, 0x0a, 0x0b, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x65, 0x62, 0x44, 0x61, 0x76, 0x12, 0x2d,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x57, 0x65, 0x62, 0x44, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x65, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a,
	0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_peripheral_device_proto_rawDescOnce sync.Once
	file_common_peripheral_device_proto_rawDescData = file_common_peripheral_device_proto_rawDesc
)

func file_common_peripheral_device_proto_rawDescGZIP() []byte {
	file_common_peripheral_device_proto_rawDescOnce.Do(func() {
		file_common_peripheral_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_peripheral_device_proto_rawDescData)
	})
	return file_common_peripheral_device_proto_rawDescData
}

var file_common_peripheral_device_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_peripheral_device_proto_goTypes = []interface{}{
	(*ListRemovableDiskReply)(nil),     // 0: cloud.lazycat.apis.common.ListRemovableDiskReply
	(*RemovableFilesystemPartion)(nil), // 1: cloud.lazycat.apis.common.RemovableFilesystemPartion
	(*MountFilesystemRequest)(nil),     // 2: cloud.lazycat.apis.common.MountFilesystemRequest
	(*UmountFilesystemRequest)(nil),    // 3: cloud.lazycat.apis.common.UmountFilesystemRequest
	(*WebDavAuth)(nil),                 // 4: cloud.lazycat.apis.common.WebDavAuth
	(*MountWebDavRequest)(nil),         // 5: cloud.lazycat.apis.common.MountWebDavRequest
	(*emptypb.Empty)(nil),              // 6: google.protobuf.Empty
}
var file_common_peripheral_device_proto_depIdxs = []int32{
	1, // 0: cloud.lazycat.apis.common.ListRemovableDiskReply.umounted:type_name -> cloud.lazycat.apis.common.RemovableFilesystemPartion
	1, // 1: cloud.lazycat.apis.common.ListRemovableDiskReply.mounted:type_name -> cloud.lazycat.apis.common.RemovableFilesystemPartion
	4, // 2: cloud.lazycat.apis.common.MountWebDavRequest.auth:type_name -> cloud.lazycat.apis.common.WebDavAuth
	6, // 3: cloud.lazycat.apis.common.PeripheralDeviceService.ListRemovableDisk:input_type -> google.protobuf.Empty
	2, // 4: cloud.lazycat.apis.common.PeripheralDeviceService.MountFilesystem:input_type -> cloud.lazycat.apis.common.MountFilesystemRequest
	3, // 5: cloud.lazycat.apis.common.PeripheralDeviceService.UmountFilesystem:input_type -> cloud.lazycat.apis.common.UmountFilesystemRequest
	5, // 6: cloud.lazycat.apis.common.PeripheralDeviceService.MountWebDav:input_type -> cloud.lazycat.apis.common.MountWebDavRequest
	0, // 7: cloud.lazycat.apis.common.PeripheralDeviceService.ListRemovableDisk:output_type -> cloud.lazycat.apis.common.ListRemovableDiskReply
	6, // 8: cloud.lazycat.apis.common.PeripheralDeviceService.MountFilesystem:output_type -> google.protobuf.Empty
	6, // 9: cloud.lazycat.apis.common.PeripheralDeviceService.UmountFilesystem:output_type -> google.protobuf.Empty
	6, // 10: cloud.lazycat.apis.common.PeripheralDeviceService.MountWebDav:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_peripheral_device_proto_init() }
func file_common_peripheral_device_proto_init() {
	if File_common_peripheral_device_proto != nil {
		return
	}
	file_common_security_context_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_peripheral_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRemovableDiskReply); i {
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
		file_common_peripheral_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovableFilesystemPartion); i {
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
		file_common_peripheral_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountFilesystemRequest); i {
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
		file_common_peripheral_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UmountFilesystemRequest); i {
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
		file_common_peripheral_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebDavAuth); i {
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
		file_common_peripheral_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountWebDavRequest); i {
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
	file_common_peripheral_device_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_common_peripheral_device_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*UmountFilesystemRequest_Uuid)(nil),
		(*UmountFilesystemRequest_MountPoint)(nil),
	}
	file_common_peripheral_device_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_peripheral_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_peripheral_device_proto_goTypes,
		DependencyIndexes: file_common_peripheral_device_proto_depIdxs,
		MessageInfos:      file_common_peripheral_device_proto_msgTypes,
	}.Build()
	File_common_peripheral_device_proto = out.File
	file_common_peripheral_device_proto_rawDesc = nil
	file_common_peripheral_device_proto_goTypes = nil
	file_common_peripheral_device_proto_depIdxs = nil
}

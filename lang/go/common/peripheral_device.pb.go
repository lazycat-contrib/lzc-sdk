// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
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

type FsType int32

const (
	FsType_SAMBA  FsType = 0
	FsType_NFS    FsType = 1
	FsType_WEBDAV FsType = 2
)

// Enum value maps for FsType.
var (
	FsType_name = map[int32]string{
		0: "SAMBA",
		1: "NFS",
		2: "WEBDAV",
	}
	FsType_value = map[string]int32{
		"SAMBA":  0,
		"NFS":    1,
		"WEBDAV": 2,
	}
)

func (x FsType) Enum() *FsType {
	p := new(FsType)
	*p = x
	return p
}

func (x FsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FsType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_peripheral_device_proto_enumTypes[0].Descriptor()
}

func (FsType) Type() protoreflect.EnumType {
	return &file_common_peripheral_device_proto_enumTypes[0]
}

func (x FsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FsType.Descriptor instead.
func (FsType) EnumDescriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{0}
}

type MountRemoteDiskStatus int32

const (
	MountRemoteDiskStatus_SUCCESS    MountRemoteDiskStatus = 0
	MountRemoteDiskStatus_INTERNAL   MountRemoteDiskStatus = 1
	MountRemoteDiskStatus_NULLTARGET MountRemoteDiskStatus = 2
	MountRemoteDiskStatus_NULLMOUNT  MountRemoteDiskStatus = 3
	MountRemoteDiskStatus_INVALCHAR  MountRemoteDiskStatus = 4
	MountRemoteDiskStatus_MOUNTED    MountRemoteDiskStatus = 5
)

// Enum value maps for MountRemoteDiskStatus.
var (
	MountRemoteDiskStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "INTERNAL",
		2: "NULLTARGET",
		3: "NULLMOUNT",
		4: "INVALCHAR",
		5: "MOUNTED",
	}
	MountRemoteDiskStatus_value = map[string]int32{
		"SUCCESS":    0,
		"INTERNAL":   1,
		"NULLTARGET": 2,
		"NULLMOUNT":  3,
		"INVALCHAR":  4,
		"MOUNTED":    5,
	}
)

func (x MountRemoteDiskStatus) Enum() *MountRemoteDiskStatus {
	p := new(MountRemoteDiskStatus)
	*p = x
	return p
}

func (x MountRemoteDiskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MountRemoteDiskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_peripheral_device_proto_enumTypes[1].Descriptor()
}

func (MountRemoteDiskStatus) Type() protoreflect.EnumType {
	return &file_common_peripheral_device_proto_enumTypes[1]
}

func (x MountRemoteDiskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MountRemoteDiskStatus.Descriptor instead.
func (MountRemoteDiskStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{1}
}

type MountArchiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcPath    string `protobuf:"bytes,1,opt,name=src_path,json=srcPath,proto3" json:"src_path,omitempty"` // 要挂载的压缩文件的相对路径
	Mountpoint string `protobuf:"bytes,2,opt,name=mountpoint,proto3" json:"mountpoint,omitempty"`          // 压缩文件要挂载到的目标路径
}

func (x *MountArchiveRequest) Reset() {
	*x = MountArchiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountArchiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountArchiveRequest) ProtoMessage() {}

func (x *MountArchiveRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MountArchiveRequest.ProtoReflect.Descriptor instead.
func (*MountArchiveRequest) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{0}
}

func (x *MountArchiveRequest) GetSrcPath() string {
	if x != nil {
		return x.SrcPath
	}
	return ""
}

func (x *MountArchiveRequest) GetMountpoint() string {
	if x != nil {
		return x.Mountpoint
	}
	return ""
}

type ListFilesystemsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前APP未挂载的磁盘列表
	Umounted []*Filesystem `protobuf:"bytes,1,rep,name=umounted,proto3" json:"umounted,omitempty"`
	// 当前APP已挂载的磁盘列表。
	Mounted []*Filesystem `protobuf:"bytes,2,rep,name=mounted,proto3" json:"mounted,omitempty"`
}

func (x *ListFilesystemsReply) Reset() {
	*x = ListFilesystemsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesystemsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesystemsReply) ProtoMessage() {}

func (x *ListFilesystemsReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListFilesystemsReply.ProtoReflect.Descriptor instead.
func (*ListFilesystemsReply) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{1}
}

func (x *ListFilesystemsReply) GetUmounted() []*Filesystem {
	if x != nil {
		return x.Umounted
	}
	return nil
}

func (x *ListFilesystemsReply) GetMounted() []*Filesystem {
	if x != nil {
		return x.Mounted
	}
	return nil
}

type Filesystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// ntfs/fat32/ext4 ...
	Fstype string `protobuf:"bytes,2,opt,name=fstype,proto3" json:"fstype,omitempty"`
	// partion label name or other meaningful name for user(eg. sda)
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// size unit in bytes
	Size uint64 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	// uuid of filesystem, maybe nil(eg. for fuse)
	Uuid *string `protobuf:"bytes,5,opt,name=uuid,proto3,oneof" json:"uuid,omitempty"`
	// mountpoint of filesystem, nil if it's not mounted
	Mountpoint *string `protobuf:"bytes,6,opt,name=mountpoint,proto3,oneof" json:"mountpoint,omitempty"`
}

func (x *Filesystem) Reset() {
	*x = Filesystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filesystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filesystem) ProtoMessage() {}

func (x *Filesystem) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Filesystem.ProtoReflect.Descriptor instead.
func (*Filesystem) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{2}
}

func (x *Filesystem) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Filesystem) GetFstype() string {
	if x != nil {
		return x.Fstype
	}
	return ""
}

func (x *Filesystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Filesystem) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Filesystem) GetUuid() string {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return ""
}

func (x *Filesystem) GetMountpoint() string {
	if x != nil && x.Mountpoint != nil {
		return *x.Mountpoint
	}
	return ""
}

type MountDiskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *MountDiskRequest) Reset() {
	*x = MountDiskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountDiskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountDiskRequest) ProtoMessage() {}

func (x *MountDiskRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MountDiskRequest.ProtoReflect.Descriptor instead.
func (*MountDiskRequest) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{3}
}

func (x *MountDiskRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *MountDiskRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type MountRemoteDiskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoteTarget string `protobuf:"bytes,1,opt,name=remoteTarget,proto3" json:"remoteTarget,omitempty"`
	Fstype       FsType `protobuf:"varint,2,opt,name=fstype,proto3,enum=cloud.lazycat.apis.common.FsType" json:"fstype,omitempty"`
	Mountpoint   string `protobuf:"bytes,3,opt,name=mountpoint,proto3" json:"mountpoint,omitempty"`
	Port         string `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	Username     string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	Password     string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *MountRemoteDiskRequest) Reset() {
	*x = MountRemoteDiskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountRemoteDiskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountRemoteDiskRequest) ProtoMessage() {}

func (x *MountRemoteDiskRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MountRemoteDiskRequest.ProtoReflect.Descriptor instead.
func (*MountRemoteDiskRequest) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{4}
}

func (x *MountRemoteDiskRequest) GetRemoteTarget() string {
	if x != nil {
		return x.RemoteTarget
	}
	return ""
}

func (x *MountRemoteDiskRequest) GetFstype() FsType {
	if x != nil {
		return x.Fstype
	}
	return FsType_SAMBA
}

func (x *MountRemoteDiskRequest) GetMountpoint() string {
	if x != nil {
		return x.Mountpoint
	}
	return ""
}

func (x *MountRemoteDiskRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *MountRemoteDiskRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MountRemoteDiskRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type MountRemoteDiskResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status MountRemoteDiskStatus `protobuf:"varint,1,opt,name=status,proto3,enum=cloud.lazycat.apis.common.MountRemoteDiskStatus" json:"status,omitempty"`
}

func (x *MountRemoteDiskResp) Reset() {
	*x = MountRemoteDiskResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountRemoteDiskResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountRemoteDiskResp) ProtoMessage() {}

func (x *MountRemoteDiskResp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MountRemoteDiskResp.ProtoReflect.Descriptor instead.
func (*MountRemoteDiskResp) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{5}
}

func (x *MountRemoteDiskResp) GetStatus() MountRemoteDiskStatus {
	if x != nil {
		return x.Status
	}
	return MountRemoteDiskStatus_SUCCESS
}

type UmountFilesystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Target:
	//
	//	*UmountFilesystemRequest_Uuid
	//	*UmountFilesystemRequest_Mountpoint
	Target isUmountFilesystemRequest_Target `protobuf_oneof:"target"`
}

func (x *UmountFilesystemRequest) Reset() {
	*x = UmountFilesystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UmountFilesystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UmountFilesystemRequest) ProtoMessage() {}

func (x *UmountFilesystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_peripheral_device_proto_msgTypes[6]
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
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{6}
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

func (x *UmountFilesystemRequest) GetMountpoint() string {
	if x, ok := x.GetTarget().(*UmountFilesystemRequest_Mountpoint); ok {
		return x.Mountpoint
	}
	return ""
}

type isUmountFilesystemRequest_Target interface {
	isUmountFilesystemRequest_Target()
}

type UmountFilesystemRequest_Uuid struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3,oneof"`
}

type UmountFilesystemRequest_Mountpoint struct {
	Mountpoint string `protobuf:"bytes,2,opt,name=mountpoint,proto3,oneof"`
}

func (*UmountFilesystemRequest_Uuid) isUmountFilesystemRequest_Target() {}

func (*UmountFilesystemRequest_Mountpoint) isUmountFilesystemRequest_Target() {}

type PowerOffDiskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *PowerOffDiskRequest) Reset() {
	*x = PowerOffDiskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_peripheral_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerOffDiskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerOffDiskRequest) ProtoMessage() {}

func (x *PowerOffDiskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_peripheral_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerOffDiskRequest.ProtoReflect.Descriptor instead.
func (*PowerOffDiskRequest) Descriptor() ([]byte, []int) {
	return file_common_peripheral_device_proto_rawDescGZIP(), []int{7}
}

func (x *PowerOffDiskRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
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
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x13, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x41, 0x0a, 0x08, 0x75, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x75, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x22, 0xba, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x73, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x73, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x10, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0xe3, 0x01, 0x0a, 0x16, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x66, 0x73, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x66, 0x73, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5f, 0x0a, 0x13, 0x4d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x48, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5b, 0x0a, 0x17, 0x55, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0a, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x08, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x2d, 0x0a, 0x13, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x4f, 0x66, 0x66, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2a, 0x28, 0x0a, 0x06, 0x46, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x41, 0x4d, 0x42, 0x41, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e,
	0x46, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45, 0x42, 0x44, 0x41, 0x56, 0x10, 0x02,
	0x2a, 0x6d, 0x0a, 0x15, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x44,
	0x69, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e,
	0x41, 0x4c, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x55, 0x4c, 0x4c, 0x54, 0x41, 0x52, 0x47,
	0x45, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x55, 0x4c, 0x4c, 0x4d, 0x4f, 0x55, 0x4e,
	0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x43, 0x48, 0x41, 0x52,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x05, 0x32,
	0xa6, 0x05, 0x0a, 0x17, 0x50, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x5c, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2f, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x58,
	0x0a, 0x09, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x2b, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x06, 0x82, 0x97, 0x22, 0x02, 0x08, 0x00, 0x12, 0x76, 0x0a, 0x0f, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x31, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x60, 0x0a, 0x10, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x55, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0c, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x12, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x58,
	0x0a, 0x0c, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x2e,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x4f, 0x66, 0x66, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c,
	0x7a, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_common_peripheral_device_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_peripheral_device_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_peripheral_device_proto_goTypes = []interface{}{
	(FsType)(0),                     // 0: cloud.lazycat.apis.common.FsType
	(MountRemoteDiskStatus)(0),      // 1: cloud.lazycat.apis.common.MountRemoteDiskStatus
	(*MountArchiveRequest)(nil),     // 2: cloud.lazycat.apis.common.MountArchiveRequest
	(*ListFilesystemsReply)(nil),    // 3: cloud.lazycat.apis.common.ListFilesystemsReply
	(*Filesystem)(nil),              // 4: cloud.lazycat.apis.common.Filesystem
	(*MountDiskRequest)(nil),        // 5: cloud.lazycat.apis.common.MountDiskRequest
	(*MountRemoteDiskRequest)(nil),  // 6: cloud.lazycat.apis.common.MountRemoteDiskRequest
	(*MountRemoteDiskResp)(nil),     // 7: cloud.lazycat.apis.common.MountRemoteDiskResp
	(*UmountFilesystemRequest)(nil), // 8: cloud.lazycat.apis.common.UmountFilesystemRequest
	(*PowerOffDiskRequest)(nil),     // 9: cloud.lazycat.apis.common.PowerOffDiskRequest
	(*emptypb.Empty)(nil),           // 10: google.protobuf.Empty
}
var file_common_peripheral_device_proto_depIdxs = []int32{
	4,  // 0: cloud.lazycat.apis.common.ListFilesystemsReply.umounted:type_name -> cloud.lazycat.apis.common.Filesystem
	4,  // 1: cloud.lazycat.apis.common.ListFilesystemsReply.mounted:type_name -> cloud.lazycat.apis.common.Filesystem
	0,  // 2: cloud.lazycat.apis.common.MountRemoteDiskRequest.fstype:type_name -> cloud.lazycat.apis.common.FsType
	1,  // 3: cloud.lazycat.apis.common.MountRemoteDiskResp.status:type_name -> cloud.lazycat.apis.common.MountRemoteDiskStatus
	10, // 4: cloud.lazycat.apis.common.PeripheralDeviceService.DeviceChanged:input_type -> google.protobuf.Empty
	10, // 5: cloud.lazycat.apis.common.PeripheralDeviceService.ListFilesystems:input_type -> google.protobuf.Empty
	5,  // 6: cloud.lazycat.apis.common.PeripheralDeviceService.MountDisk:input_type -> cloud.lazycat.apis.common.MountDiskRequest
	6,  // 7: cloud.lazycat.apis.common.PeripheralDeviceService.MountRemoteDisk:input_type -> cloud.lazycat.apis.common.MountRemoteDiskRequest
	8,  // 8: cloud.lazycat.apis.common.PeripheralDeviceService.UmountFilesystem:input_type -> cloud.lazycat.apis.common.UmountFilesystemRequest
	2,  // 9: cloud.lazycat.apis.common.PeripheralDeviceService.MountArchive:input_type -> cloud.lazycat.apis.common.MountArchiveRequest
	9,  // 10: cloud.lazycat.apis.common.PeripheralDeviceService.PowerOffDisk:input_type -> cloud.lazycat.apis.common.PowerOffDiskRequest
	10, // 11: cloud.lazycat.apis.common.PeripheralDeviceService.DeviceChanged:output_type -> google.protobuf.Empty
	3,  // 12: cloud.lazycat.apis.common.PeripheralDeviceService.ListFilesystems:output_type -> cloud.lazycat.apis.common.ListFilesystemsReply
	10, // 13: cloud.lazycat.apis.common.PeripheralDeviceService.MountDisk:output_type -> google.protobuf.Empty
	7,  // 14: cloud.lazycat.apis.common.PeripheralDeviceService.MountRemoteDisk:output_type -> cloud.lazycat.apis.common.MountRemoteDiskResp
	10, // 15: cloud.lazycat.apis.common.PeripheralDeviceService.UmountFilesystem:output_type -> google.protobuf.Empty
	10, // 16: cloud.lazycat.apis.common.PeripheralDeviceService.MountArchive:output_type -> google.protobuf.Empty
	10, // 17: cloud.lazycat.apis.common.PeripheralDeviceService.PowerOffDisk:output_type -> google.protobuf.Empty
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_common_peripheral_device_proto_init() }
func file_common_peripheral_device_proto_init() {
	if File_common_peripheral_device_proto != nil {
		return
	}
	file_common_security_context_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_peripheral_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountArchiveRequest); i {
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
			switch v := v.(*ListFilesystemsReply); i {
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
			switch v := v.(*Filesystem); i {
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
			switch v := v.(*MountDiskRequest); i {
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
			switch v := v.(*MountRemoteDiskRequest); i {
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
			switch v := v.(*MountRemoteDiskResp); i {
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
		file_common_peripheral_device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_peripheral_device_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerOffDiskRequest); i {
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
	file_common_peripheral_device_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_common_peripheral_device_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*UmountFilesystemRequest_Uuid)(nil),
		(*UmountFilesystemRequest_Mountpoint)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_peripheral_device_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_peripheral_device_proto_goTypes,
		DependencyIndexes: file_common_peripheral_device_proto_depIdxs,
		EnumInfos:         file_common_peripheral_device_proto_enumTypes,
		MessageInfos:      file_common_peripheral_device_proto_msgTypes,
	}.Build()
	File_common_peripheral_device_proto = out.File
	file_common_peripheral_device_proto_rawDesc = nil
	file_common_peripheral_device_proto_goTypes = nil
	file_common_peripheral_device_proto_depIdxs = nil
}

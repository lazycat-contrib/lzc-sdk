// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: sys/snapd.proto

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

type SnapdAutoStrategy int32

const (
	// 修改该配置项时，此值表示不改变当前配置值
	SnapdAutoStrategy_Default SnapdAutoStrategy = 0
	// 不自动快照（但超过生命周期的历史快照仍会被清理）
	SnapdAutoStrategy_Disabled SnapdAutoStrategy = 1
	// 自动快照，但不自动备份
	SnapdAutoStrategy_SnapOnly SnapdAutoStrategy = 2
	// 自动快照并备份
	SnapdAutoStrategy_SnapAndBackup SnapdAutoStrategy = 3
)

// Enum value maps for SnapdAutoStrategy.
var (
	SnapdAutoStrategy_name = map[int32]string{
		0: "Default",
		1: "Disabled",
		2: "SnapOnly",
		3: "SnapAndBackup",
	}
	SnapdAutoStrategy_value = map[string]int32{
		"Default":       0,
		"Disabled":      1,
		"SnapOnly":      2,
		"SnapAndBackup": 3,
	}
)

func (x SnapdAutoStrategy) Enum() *SnapdAutoStrategy {
	p := new(SnapdAutoStrategy)
	*p = x
	return p
}

func (x SnapdAutoStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SnapdAutoStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_sys_snapd_proto_enumTypes[0].Descriptor()
}

func (SnapdAutoStrategy) Type() protoreflect.EnumType {
	return &file_sys_snapd_proto_enumTypes[0]
}

func (x SnapdAutoStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SnapdAutoStrategy.Descriptor instead.
func (SnapdAutoStrategy) EnumDescriptor() ([]byte, []int) {
	return file_sys_snapd_proto_rawDescGZIP(), []int{0}
}

type SnapdEnableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId string `protobuf:"bytes,1,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	// 需要启用备份的路径列表
	PathList []string `protobuf:"bytes,2,rep,name=PathList,proto3" json:"PathList,omitempty"`
}

func (x *SnapdEnableRequest) Reset() {
	*x = SnapdEnableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_snapd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapdEnableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapdEnableRequest) ProtoMessage() {}

func (x *SnapdEnableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_snapd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapdEnableRequest.ProtoReflect.Descriptor instead.
func (*SnapdEnableRequest) Descriptor() ([]byte, []int) {
	return file_sys_snapd_proto_rawDescGZIP(), []int{0}
}

func (x *SnapdEnableRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *SnapdEnableRequest) GetPathList() []string {
	if x != nil {
		return x.PathList
	}
	return nil
}

type SnapdTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId string `protobuf:"bytes,1,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
}

func (x *SnapdTargetRequest) Reset() {
	*x = SnapdTargetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_snapd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapdTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapdTargetRequest) ProtoMessage() {}

func (x *SnapdTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_snapd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapdTargetRequest.ProtoReflect.Descriptor instead.
func (*SnapdTargetRequest) Descriptor() ([]byte, []int) {
	return file_sys_snapd_proto_rawDescGZIP(), []int{1}
}

func (x *SnapdTargetRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

type SnapdConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 自动快照策略
	AutoStrategy SnapdAutoStrategy `protobuf:"varint,1,opt,name=AutoStrategy,proto3,enum=cloud.lazycat.apis.sys.SnapdAutoStrategy" json:"AutoStrategy,omitempty"`
	// 自动快照的时间间隔，单位为分钟
	AutoSnapInterval uint32 `protobuf:"varint,2,opt,name=AutoSnapInterval,proto3" json:"AutoSnapInterval,omitempty"`
	// 自动快照保留的时长，单位为分钟。超出的快照会被自动删除
	AutoSnapLifetime uint32 `protobuf:"varint,3,opt,name=AutoSnapLifetime,proto3" json:"AutoSnapLifetime,omitempty"`
	// 自动备份快照保留的时长，单位为分钟。当单个备份池中的备份超过该数量时，旧备份将自动被删除
	AutoBackupLifetime uint32 `protobuf:"varint,4,opt,name=AutoBackupLifetime,proto3" json:"AutoBackupLifetime,omitempty"`
}

func (x *SnapdConf) Reset() {
	*x = SnapdConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_snapd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapdConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapdConf) ProtoMessage() {}

func (x *SnapdConf) ProtoReflect() protoreflect.Message {
	mi := &file_sys_snapd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapdConf.ProtoReflect.Descriptor instead.
func (*SnapdConf) Descriptor() ([]byte, []int) {
	return file_sys_snapd_proto_rawDescGZIP(), []int{2}
}

func (x *SnapdConf) GetAutoStrategy() SnapdAutoStrategy {
	if x != nil {
		return x.AutoStrategy
	}
	return SnapdAutoStrategy_Default
}

func (x *SnapdConf) GetAutoSnapInterval() uint32 {
	if x != nil {
		return x.AutoSnapInterval
	}
	return 0
}

func (x *SnapdConf) GetAutoSnapLifetime() uint32 {
	if x != nil {
		return x.AutoSnapLifetime
	}
	return 0
}

func (x *SnapdConf) GetAutoBackupLifetime() uint32 {
	if x != nil {
		return x.AutoBackupLifetime
	}
	return 0
}

type SnapdConfSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId string     `protobuf:"bytes,1,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	Config   *SnapdConf `protobuf:"bytes,2,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (x *SnapdConfSetRequest) Reset() {
	*x = SnapdConfSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_snapd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapdConfSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapdConfSetRequest) ProtoMessage() {}

func (x *SnapdConfSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_snapd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapdConfSetRequest.ProtoReflect.Descriptor instead.
func (*SnapdConfSetRequest) Descriptor() ([]byte, []int) {
	return file_sys_snapd_proto_rawDescGZIP(), []int{3}
}

func (x *SnapdConfSetRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *SnapdConfSetRequest) GetConfig() *SnapdConf {
	if x != nil {
		return x.Config
	}
	return nil
}

type SnapdListSnapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId string `protobuf:"bytes,1,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	// 为空则查看盒子数据盘上快照，否则查看指定备份池中快照
	BackupPool string `protobuf:"bytes,2,opt,name=BackupPool,proto3" json:"BackupPool,omitempty"`
}

func (x *SnapdListSnapRequest) Reset() {
	*x = SnapdListSnapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_snapd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapdListSnapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapdListSnapRequest) ProtoMessage() {}

func (x *SnapdListSnapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_snapd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapdListSnapRequest.ProtoReflect.Descriptor instead.
func (*SnapdListSnapRequest) Descriptor() ([]byte, []int) {
	return file_sys_snapd_proto_rawDescGZIP(), []int{4}
}

func (x *SnapdListSnapRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *SnapdListSnapRequest) GetBackupPool() string {
	if x != nil {
		return x.BackupPool
	}
	return ""
}

type SnapdListSnapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnapshotList []*SnapshotDesc `protobuf:"bytes,1,rep,name=SnapshotList,proto3" json:"SnapshotList,omitempty"`
}

func (x *SnapdListSnapResponse) Reset() {
	*x = SnapdListSnapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_snapd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapdListSnapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapdListSnapResponse) ProtoMessage() {}

func (x *SnapdListSnapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_snapd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapdListSnapResponse.ProtoReflect.Descriptor instead.
func (*SnapdListSnapResponse) Descriptor() ([]byte, []int) {
	return file_sys_snapd_proto_rawDescGZIP(), []int{5}
}

func (x *SnapdListSnapResponse) GetSnapshotList() []*SnapshotDesc {
	if x != nil {
		return x.SnapshotList
	}
	return nil
}

type SnapdSnapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId string `protobuf:"bytes,1,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	// SnapdTakeSnap: 不为空，则同时将快照备份到指定备份池
	// SnapdDelSnap: 为空则删除盒子内快照，否则删除备份池中快照
	// SnapdRestoreSnap：为空则回滚盒子内快照，否则从备份池中读取快照并还原
	BackupPool string `protobuf:"bytes,2,opt,name=BackupPool,proto3" json:"BackupPool,omitempty"`
	SnapName   string `protobuf:"bytes,3,opt,name=SnapName,proto3" json:"SnapName,omitempty"`
}

func (x *SnapdSnapRequest) Reset() {
	*x = SnapdSnapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_snapd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapdSnapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapdSnapRequest) ProtoMessage() {}

func (x *SnapdSnapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_snapd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapdSnapRequest.ProtoReflect.Descriptor instead.
func (*SnapdSnapRequest) Descriptor() ([]byte, []int) {
	return file_sys_snapd_proto_rawDescGZIP(), []int{6}
}

func (x *SnapdSnapRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *SnapdSnapRequest) GetBackupPool() string {
	if x != nil {
		return x.BackupPool
	}
	return ""
}

func (x *SnapdSnapRequest) GetSnapName() string {
	if x != nil {
		return x.SnapName
	}
	return ""
}

var File_sys_snapd_proto protoreflect.FileDescriptor

var file_sys_snapd_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x79, 0x73, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x79, 0x73, 0x2f, 0x4f, 0x53, 0x5f, 0x73,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a,
	0x12, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x12, 0x53,
	0x6e, 0x61, 0x70, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x22, 0xe2, 0x01,
	0x0a, 0x09, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x4d, 0x0a, 0x0c, 0x41,
	0x75, 0x74, 0x6f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x64,
	0x41, 0x75, 0x74, 0x6f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x0c, 0x41, 0x75,
	0x74, 0x6f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x41, 0x75,
	0x74, 0x6f, 0x53, 0x6e, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x6e, 0x61, 0x70, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x6e,
	0x61, 0x70, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x10, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x6e, 0x61, 0x70, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x6f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12,
	0x41, 0x75, 0x74, 0x6f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x6c, 0x0a, 0x13, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61,
	0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53,
	0x6e, 0x61, 0x70, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x52, 0x0a, 0x14, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6f,
	0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x50, 0x6f, 0x6f, 0x6c, 0x22, 0x61, 0x0a, 0x15, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a,
	0x0c, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x44, 0x65, 0x73, 0x63, 0x52, 0x0c, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x6a, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70, 0x64,
	0x53, 0x6e, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x50, 0x6f, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x2a, 0x4f, 0x0a, 0x11, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x41, 0x75, 0x74, 0x6f,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x4f, 0x6e, 0x6c, 0x79, 0x10,
	0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x41, 0x6e, 0x64, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x10, 0x03, 0x32, 0xe1, 0x05, 0x0a, 0x0c, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0b, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x6e,
	0x61, 0x70, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0c, 0x53, 0x6e,
	0x61, 0x70, 0x64, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x5f, 0x0a, 0x0c, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x47, 0x65, 0x74,
	0x12, 0x2a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x22,
	0x00, 0x12, 0x55, 0x0a, 0x0c, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x53, 0x65,
	0x74, 0x12, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x0d, 0x53, 0x6e, 0x61, 0x70,
	0x64, 0x53, 0x6e, 0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73,
	0x79, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73,
	0x2e, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0c, 0x53, 0x6e, 0x61, 0x70,
	0x64, 0x53, 0x6e, 0x61, 0x70, 0x41, 0x64, 0x64, 0x12, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79,
	0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0c,
	0x53, 0x6e, 0x61, 0x70, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x44, 0x65, 0x6c, 0x12, 0x28, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x56, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x52, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x6e,
	0x61, 0x70, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c,
	0x7a, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x73,
	0x79, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_snapd_proto_rawDescOnce sync.Once
	file_sys_snapd_proto_rawDescData = file_sys_snapd_proto_rawDesc
)

func file_sys_snapd_proto_rawDescGZIP() []byte {
	file_sys_snapd_proto_rawDescOnce.Do(func() {
		file_sys_snapd_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_snapd_proto_rawDescData)
	})
	return file_sys_snapd_proto_rawDescData
}

var file_sys_snapd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sys_snapd_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sys_snapd_proto_goTypes = []interface{}{
	(SnapdAutoStrategy)(0),        // 0: cloud.lazycat.apis.sys.SnapdAutoStrategy
	(*SnapdEnableRequest)(nil),    // 1: cloud.lazycat.apis.sys.SnapdEnableRequest
	(*SnapdTargetRequest)(nil),    // 2: cloud.lazycat.apis.sys.SnapdTargetRequest
	(*SnapdConf)(nil),             // 3: cloud.lazycat.apis.sys.SnapdConf
	(*SnapdConfSetRequest)(nil),   // 4: cloud.lazycat.apis.sys.SnapdConfSetRequest
	(*SnapdListSnapRequest)(nil),  // 5: cloud.lazycat.apis.sys.SnapdListSnapRequest
	(*SnapdListSnapResponse)(nil), // 6: cloud.lazycat.apis.sys.SnapdListSnapResponse
	(*SnapdSnapRequest)(nil),      // 7: cloud.lazycat.apis.sys.SnapdSnapRequest
	(*SnapshotDesc)(nil),          // 8: cloud.lazycat.apis.sys.SnapshotDesc
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_sys_snapd_proto_depIdxs = []int32{
	0,  // 0: cloud.lazycat.apis.sys.SnapdConf.AutoStrategy:type_name -> cloud.lazycat.apis.sys.SnapdAutoStrategy
	3,  // 1: cloud.lazycat.apis.sys.SnapdConfSetRequest.Config:type_name -> cloud.lazycat.apis.sys.SnapdConf
	8,  // 2: cloud.lazycat.apis.sys.SnapdListSnapResponse.SnapshotList:type_name -> cloud.lazycat.apis.sys.SnapshotDesc
	1,  // 3: cloud.lazycat.apis.sys.SnapdService.SnapdEnable:input_type -> cloud.lazycat.apis.sys.SnapdEnableRequest
	2,  // 4: cloud.lazycat.apis.sys.SnapdService.SnapdDisable:input_type -> cloud.lazycat.apis.sys.SnapdTargetRequest
	2,  // 5: cloud.lazycat.apis.sys.SnapdService.SnapdConfGet:input_type -> cloud.lazycat.apis.sys.SnapdTargetRequest
	4,  // 6: cloud.lazycat.apis.sys.SnapdService.SnapdConfSet:input_type -> cloud.lazycat.apis.sys.SnapdConfSetRequest
	5,  // 7: cloud.lazycat.apis.sys.SnapdService.SnapdSnapList:input_type -> cloud.lazycat.apis.sys.SnapdListSnapRequest
	7,  // 8: cloud.lazycat.apis.sys.SnapdService.SnapdSnapAdd:input_type -> cloud.lazycat.apis.sys.SnapdSnapRequest
	7,  // 9: cloud.lazycat.apis.sys.SnapdService.SnapdSnapDel:input_type -> cloud.lazycat.apis.sys.SnapdSnapRequest
	7,  // 10: cloud.lazycat.apis.sys.SnapdService.SnapdSnapRestore:input_type -> cloud.lazycat.apis.sys.SnapdSnapRequest
	9,  // 11: cloud.lazycat.apis.sys.SnapdService.SnapdEnable:output_type -> google.protobuf.Empty
	9,  // 12: cloud.lazycat.apis.sys.SnapdService.SnapdDisable:output_type -> google.protobuf.Empty
	3,  // 13: cloud.lazycat.apis.sys.SnapdService.SnapdConfGet:output_type -> cloud.lazycat.apis.sys.SnapdConf
	9,  // 14: cloud.lazycat.apis.sys.SnapdService.SnapdConfSet:output_type -> google.protobuf.Empty
	6,  // 15: cloud.lazycat.apis.sys.SnapdService.SnapdSnapList:output_type -> cloud.lazycat.apis.sys.SnapdListSnapResponse
	9,  // 16: cloud.lazycat.apis.sys.SnapdService.SnapdSnapAdd:output_type -> google.protobuf.Empty
	9,  // 17: cloud.lazycat.apis.sys.SnapdService.SnapdSnapDel:output_type -> google.protobuf.Empty
	9,  // 18: cloud.lazycat.apis.sys.SnapdService.SnapdSnapRestore:output_type -> google.protobuf.Empty
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_sys_snapd_proto_init() }
func file_sys_snapd_proto_init() {
	if File_sys_snapd_proto != nil {
		return
	}
	file_sys_OS_snapshot_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sys_snapd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapdEnableRequest); i {
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
		file_sys_snapd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapdTargetRequest); i {
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
		file_sys_snapd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapdConf); i {
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
		file_sys_snapd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapdConfSetRequest); i {
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
		file_sys_snapd_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapdListSnapRequest); i {
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
		file_sys_snapd_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapdListSnapResponse); i {
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
		file_sys_snapd_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapdSnapRequest); i {
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
			RawDescriptor: file_sys_snapd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_snapd_proto_goTypes,
		DependencyIndexes: file_sys_snapd_proto_depIdxs,
		EnumInfos:         file_sys_snapd_proto_enumTypes,
		MessageInfos:      file_sys_snapd_proto_msgTypes,
	}.Build()
	File_sys_snapd_proto = out.File
	file_sys_snapd_proto_rawDesc = nil
	file_sys_snapd_proto_goTypes = nil
	file_sys_snapd_proto_depIdxs = nil
}

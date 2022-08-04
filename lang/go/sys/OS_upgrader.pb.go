// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: sys/OS_upgrader.proto

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

type DownloaderStatus int32

const (
	DownloaderStatus_ready       DownloaderStatus = 0
	DownloaderStatus_downloading DownloaderStatus = 1
	DownloaderStatus_paused      DownloaderStatus = 2
	DownloaderStatus_completed   DownloaderStatus = 3
	DownloaderStatus_error       DownloaderStatus = 4
)

// Enum value maps for DownloaderStatus.
var (
	DownloaderStatus_name = map[int32]string{
		0: "ready",
		1: "downloading",
		2: "paused",
		3: "completed",
		4: "error",
	}
	DownloaderStatus_value = map[string]int32{
		"ready":       0,
		"downloading": 1,
		"paused":      2,
		"completed":   3,
		"error":       4,
	}
)

func (x DownloaderStatus) Enum() *DownloaderStatus {
	p := new(DownloaderStatus)
	*p = x
	return p
}

func (x DownloaderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DownloaderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_sys_OS_upgrader_proto_enumTypes[0].Descriptor()
}

func (DownloaderStatus) Type() protoreflect.EnumType {
	return &file_sys_OS_upgrader_proto_enumTypes[0]
}

func (x DownloaderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DownloaderStatus.Descriptor instead.
func (DownloaderStatus) EnumDescriptor() ([]byte, []int) {
	return file_sys_OS_upgrader_proto_rawDescGZIP(), []int{0}
}

type LocalSystemVersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前正在运行的系统版本与下一次重启的系统版本，两者通常是一样的
	//    更新之后，会改变 rebootVersion 的值从而两者不一样，表示需要重启才能生效
	CurrentVersion string `protobuf:"bytes,1,opt,name=currentVersion,proto3" json:"currentVersion,omitempty"`
	RebootVersion  string `protobuf:"bytes,2,opt,name=rebootVersion,proto3" json:"rebootVersion,omitempty"`
	// 当前磁盘中所有版本列表（由于产品设计上只允许切换到最新版本，该字段仅供调试）
	Versions []string `protobuf:"bytes,3,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *LocalSystemVersionInfo) Reset() {
	*x = LocalSystemVersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_OS_upgrader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalSystemVersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalSystemVersionInfo) ProtoMessage() {}

func (x *LocalSystemVersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sys_OS_upgrader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalSystemVersionInfo.ProtoReflect.Descriptor instead.
func (*LocalSystemVersionInfo) Descriptor() ([]byte, []int) {
	return file_sys_OS_upgrader_proto_rawDescGZIP(), []int{0}
}

func (x *LocalSystemVersionInfo) GetCurrentVersion() string {
	if x != nil {
		return x.CurrentVersion
	}
	return ""
}

func (x *LocalSystemVersionInfo) GetRebootVersion() string {
	if x != nil {
		return x.RebootVersion
	}
	return ""
}

func (x *LocalSystemVersionInfo) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

type RemoteSystemVersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 最新版本
	//     需要对比该字段与当前运行的版本是否一致来确定当前系统是否是最新的
	LatestVersion string `protobuf:"bytes,1,opt,name=latestVersion,proto3" json:"latestVersion,omitempty"`
	// 线上可下载的所有版本列表，按顺序最新版本在前（由于产品设计上只下载最新版本，该字段仅暂时保留）
	Versions []string `protobuf:"bytes,3,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *RemoteSystemVersionInfo) Reset() {
	*x = RemoteSystemVersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_OS_upgrader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteSystemVersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteSystemVersionInfo) ProtoMessage() {}

func (x *RemoteSystemVersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sys_OS_upgrader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteSystemVersionInfo.ProtoReflect.Descriptor instead.
func (*RemoteSystemVersionInfo) Descriptor() ([]byte, []int) {
	return file_sys_OS_upgrader_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteSystemVersionInfo) GetLatestVersion() string {
	if x != nil {
		return x.LatestVersion
	}
	return ""
}

func (x *RemoteSystemVersionInfo) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

type UpgradeProgressInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 正在下载的系统版本，（为空表示没有正在下载的进度，后面所有字段都没有意义）
	Version string           `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Status  DownloaderStatus `protobuf:"varint,2,opt,name=status,proto3,enum=cloud.lazycat.apis.sys.DownloaderStatus" json:"status,omitempty"`
	// 已下载的大小（字节）
	Complete int64 `protobuf:"varint,3,opt,name=complete,proto3" json:"complete,omitempty"`
	// 总需要下载的大小（字节）
	Total int64 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	// 错误信息（在没有错误时该值为空）
	Error string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpgradeProgressInfo) Reset() {
	*x = UpgradeProgressInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_OS_upgrader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeProgressInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeProgressInfo) ProtoMessage() {}

func (x *UpgradeProgressInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sys_OS_upgrader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeProgressInfo.ProtoReflect.Descriptor instead.
func (*UpgradeProgressInfo) Descriptor() ([]byte, []int) {
	return file_sys_OS_upgrader_proto_rawDescGZIP(), []int{2}
}

func (x *UpgradeProgressInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *UpgradeProgressInfo) GetStatus() DownloaderStatus {
	if x != nil {
		return x.Status
	}
	return DownloaderStatus_ready
}

func (x *UpgradeProgressInfo) GetComplete() int64 {
	if x != nil {
		return x.Complete
	}
	return 0
}

func (x *UpgradeProgressInfo) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UpgradeProgressInfo) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SystemVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SystemVersion) Reset() {
	*x = SystemVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_OS_upgrader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemVersion) ProtoMessage() {}

func (x *SystemVersion) ProtoReflect() protoreflect.Message {
	mi := &file_sys_OS_upgrader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemVersion.ProtoReflect.Descriptor instead.
func (*SystemVersion) Descriptor() ([]byte, []int) {
	return file_sys_OS_upgrader_proto_rawDescGZIP(), []int{3}
}

func (x *SystemVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_sys_OS_upgrader_proto protoreflect.FileDescriptor

var file_sys_OS_upgrader_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x79, 0x73, 0x2f, 0x4f, 0x53, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a,
	0x16, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x0d, 0x72, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x5b, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb9,
	0x01, 0x0a, 0x13, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x54, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x03,
	0x12, 0x09, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x04, 0x32, 0xcd, 0x05, 0x0a, 0x10,
	0x4f, 0x53, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x51, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61,
	0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x05, 0x50, 0x61, 0x75, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2b, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x07,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x12, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x06, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73,
	0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x79, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_OS_upgrader_proto_rawDescOnce sync.Once
	file_sys_OS_upgrader_proto_rawDescData = file_sys_OS_upgrader_proto_rawDesc
)

func file_sys_OS_upgrader_proto_rawDescGZIP() []byte {
	file_sys_OS_upgrader_proto_rawDescOnce.Do(func() {
		file_sys_OS_upgrader_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_OS_upgrader_proto_rawDescData)
	})
	return file_sys_OS_upgrader_proto_rawDescData
}

var file_sys_OS_upgrader_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sys_OS_upgrader_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sys_OS_upgrader_proto_goTypes = []interface{}{
	(DownloaderStatus)(0),           // 0: cloud.lazycat.apis.sys.DownloaderStatus
	(*LocalSystemVersionInfo)(nil),  // 1: cloud.lazycat.apis.sys.LocalSystemVersionInfo
	(*RemoteSystemVersionInfo)(nil), // 2: cloud.lazycat.apis.sys.RemoteSystemVersionInfo
	(*UpgradeProgressInfo)(nil),     // 3: cloud.lazycat.apis.sys.UpgradeProgressInfo
	(*SystemVersion)(nil),           // 4: cloud.lazycat.apis.sys.SystemVersion
	(*emptypb.Empty)(nil),           // 5: google.protobuf.Empty
}
var file_sys_OS_upgrader_proto_depIdxs = []int32{
	0,  // 0: cloud.lazycat.apis.sys.UpgradeProgressInfo.status:type_name -> cloud.lazycat.apis.sys.DownloaderStatus
	5,  // 1: cloud.lazycat.apis.sys.OSUpgradeService.Local:input_type -> google.protobuf.Empty
	5,  // 2: cloud.lazycat.apis.sys.OSUpgradeService.Remote:input_type -> google.protobuf.Empty
	4,  // 3: cloud.lazycat.apis.sys.OSUpgradeService.Select:input_type -> cloud.lazycat.apis.sys.SystemVersion
	5,  // 4: cloud.lazycat.apis.sys.OSUpgradeService.Start:input_type -> google.protobuf.Empty
	5,  // 5: cloud.lazycat.apis.sys.OSUpgradeService.Pause:input_type -> google.protobuf.Empty
	5,  // 6: cloud.lazycat.apis.sys.OSUpgradeService.Progress:input_type -> google.protobuf.Empty
	5,  // 7: cloud.lazycat.apis.sys.OSUpgradeService.Install:input_type -> google.protobuf.Empty
	4,  // 8: cloud.lazycat.apis.sys.OSUpgradeService.Switch:input_type -> cloud.lazycat.apis.sys.SystemVersion
	5,  // 9: cloud.lazycat.apis.sys.OSUpgradeService.Prune:input_type -> google.protobuf.Empty
	5,  // 10: cloud.lazycat.apis.sys.OSUpgradeService.Reboot:input_type -> google.protobuf.Empty
	1,  // 11: cloud.lazycat.apis.sys.OSUpgradeService.Local:output_type -> cloud.lazycat.apis.sys.LocalSystemVersionInfo
	2,  // 12: cloud.lazycat.apis.sys.OSUpgradeService.Remote:output_type -> cloud.lazycat.apis.sys.RemoteSystemVersionInfo
	5,  // 13: cloud.lazycat.apis.sys.OSUpgradeService.Select:output_type -> google.protobuf.Empty
	5,  // 14: cloud.lazycat.apis.sys.OSUpgradeService.Start:output_type -> google.protobuf.Empty
	5,  // 15: cloud.lazycat.apis.sys.OSUpgradeService.Pause:output_type -> google.protobuf.Empty
	3,  // 16: cloud.lazycat.apis.sys.OSUpgradeService.Progress:output_type -> cloud.lazycat.apis.sys.UpgradeProgressInfo
	5,  // 17: cloud.lazycat.apis.sys.OSUpgradeService.Install:output_type -> google.protobuf.Empty
	5,  // 18: cloud.lazycat.apis.sys.OSUpgradeService.Switch:output_type -> google.protobuf.Empty
	5,  // 19: cloud.lazycat.apis.sys.OSUpgradeService.Prune:output_type -> google.protobuf.Empty
	5,  // 20: cloud.lazycat.apis.sys.OSUpgradeService.Reboot:output_type -> google.protobuf.Empty
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_sys_OS_upgrader_proto_init() }
func file_sys_OS_upgrader_proto_init() {
	if File_sys_OS_upgrader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_OS_upgrader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalSystemVersionInfo); i {
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
		file_sys_OS_upgrader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteSystemVersionInfo); i {
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
		file_sys_OS_upgrader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeProgressInfo); i {
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
		file_sys_OS_upgrader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemVersion); i {
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
			RawDescriptor: file_sys_OS_upgrader_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_OS_upgrader_proto_goTypes,
		DependencyIndexes: file_sys_OS_upgrader_proto_depIdxs,
		EnumInfos:         file_sys_OS_upgrader_proto_enumTypes,
		MessageInfos:      file_sys_OS_upgrader_proto_msgTypes,
	}.Build()
	File_sys_OS_upgrader_proto = out.File
	file_sys_OS_upgrader_proto_rawDesc = nil
	file_sys_OS_upgrader_proto_goTypes = nil
	file_sys_OS_upgrader_proto_depIdxs = nil
}

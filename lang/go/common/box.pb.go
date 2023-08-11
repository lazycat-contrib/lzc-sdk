// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
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
	// 关机
	ShutdownRequest_Poweroff ShutdownRequest_Action = 0
	// 重启
	ShutdownRequest_Reboot ShutdownRequest_Action = 1
	// 软重启（仅重启 lzc-os 容器，不关闭 lzc-base-os）
	//
	//	通常用于系统更新、故障修复等场景
	ShutdownRequest_SoftReboot ShutdownRequest_Action = 2
)

// Enum value maps for ShutdownRequest_Action.
var (
	ShutdownRequest_Action_name = map[int32]string{
		0: "Poweroff",
		1: "Reboot",
		2: "SoftReboot",
	}
	ShutdownRequest_Action_value = map[string]int32{
		"Poweroff":   0,
		"Reboot":     1,
		"SoftReboot": 2,
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

type BootOption_BootOptionType int32

const (
	// 无操作（清空所有已设置的操作）
	BootOption_BOOT_OPTION_NONE BootOption_BootOptionType = 0
	// 回滚到上一个版本
	BootOption_BOOT_OPTION_ROLLBACK BootOption_BootOptionType = 1
	// 重置系统（清空系统的 var 数据）
	BootOption_BOOT_OPTION_RESET BootOption_BootOptionType = 2
	// 恢复出厂设置（清空用户信息、系统的 var 数据和用户数据）
	BootOption_BOOT_OPTION_FACTORY_RESET BootOption_BootOptionType = 3
)

// Enum value maps for BootOption_BootOptionType.
var (
	BootOption_BootOptionType_name = map[int32]string{
		0: "BOOT_OPTION_NONE",
		1: "BOOT_OPTION_ROLLBACK",
		2: "BOOT_OPTION_RESET",
		3: "BOOT_OPTION_FACTORY_RESET",
	}
	BootOption_BootOptionType_value = map[string]int32{
		"BOOT_OPTION_NONE":          0,
		"BOOT_OPTION_ROLLBACK":      1,
		"BOOT_OPTION_RESET":         2,
		"BOOT_OPTION_FACTORY_RESET": 3,
	}
)

func (x BootOption_BootOptionType) Enum() *BootOption_BootOptionType {
	p := new(BootOption_BootOptionType)
	*p = x
	return p
}

func (x BootOption_BootOptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BootOption_BootOptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_box_proto_enumTypes[1].Descriptor()
}

func (BootOption_BootOptionType) Type() protoreflect.EnumType {
	return &file_common_box_proto_enumTypes[1]
}

func (x BootOption_BootOptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BootOption_BootOptionType.Descriptor instead.
func (BootOption_BootOptionType) EnumDescriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{3, 0}
}

type DiskInfo_DiskType int32

const (
	DiskInfo_Unknown DiskInfo_DiskType = 0
	DiskInfo_System  DiskInfo_DiskType = 1
	DiskInfo_Data    DiskInfo_DiskType = 2
)

// Enum value maps for DiskInfo_DiskType.
var (
	DiskInfo_DiskType_name = map[int32]string{
		0: "Unknown",
		1: "System",
		2: "Data",
	}
	DiskInfo_DiskType_value = map[string]int32{
		"Unknown": 0,
		"System":  1,
		"Data":    2,
	}
)

func (x DiskInfo_DiskType) Enum() *DiskInfo_DiskType {
	p := new(DiskInfo_DiskType)
	*p = x
	return p
}

func (x DiskInfo_DiskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiskInfo_DiskType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_box_proto_enumTypes[2].Descriptor()
}

func (DiskInfo_DiskType) Type() protoreflect.EnumType {
	return &file_common_box_proto_enumTypes[2]
}

func (x DiskInfo_DiskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiskInfo_DiskType.Descriptor instead.
func (DiskInfo_DiskType) EnumDescriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{4, 0}
}

// 健康状态
type DiskInfo_Health int32

const (
	DiskInfo_Normal  DiskInfo_Health = 0
	DiskInfo_Caution DiskInfo_Health = 1
	DiskInfo_Bad     DiskInfo_Health = 2
)

// Enum value maps for DiskInfo_Health.
var (
	DiskInfo_Health_name = map[int32]string{
		0: "Normal",
		1: "Caution",
		2: "Bad",
	}
	DiskInfo_Health_value = map[string]int32{
		"Normal":  0,
		"Caution": 1,
		"Bad":     2,
	}
)

func (x DiskInfo_Health) Enum() *DiskInfo_Health {
	p := new(DiskInfo_Health)
	*p = x
	return p
}

func (x DiskInfo_Health) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiskInfo_Health) Descriptor() protoreflect.EnumDescriptor {
	return file_common_box_proto_enumTypes[3].Descriptor()
}

func (DiskInfo_Health) Type() protoreflect.EnumType {
	return &file_common_box_proto_enumTypes[3]
}

func (x DiskInfo_Health) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiskInfo_Health.Descriptor instead.
func (DiskInfo_Health) EnumDescriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{4, 1}
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

	Action ShutdownRequest_Action `protobuf:"varint,1,opt,name=action,proto3,enum=cloud.lazycat.apis.common.ShutdownRequest_Action" json:"action,omitempty"`
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

type BootOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type BootOption_BootOptionType `protobuf:"varint,1,opt,name=type,proto3,enum=cloud.lazycat.apis.common.BootOption_BootOptionType" json:"type,omitempty"`
}

func (x *BootOption) Reset() {
	*x = BootOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_box_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootOption) ProtoMessage() {}

func (x *BootOption) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BootOption.ProtoReflect.Descriptor instead.
func (*BootOption) Descriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{3}
}

func (x *BootOption) GetType() BootOption_BootOptionType {
	if x != nil {
		return x.Type
	}
	return BootOption_BOOT_OPTION_NONE
}

type DiskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 磁盘类型
	Type DiskInfo_DiskType `protobuf:"varint,1,opt,name=type,proto3,enum=cloud.lazycat.apis.common.DiskInfo_DiskType" json:"type,omitempty"`
	// 型号
	Model string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	// 序列号
	Serial string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	// 磁盘容量（字节）
	Size int64 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	// 剩余空间（字节）
	Free int64 `protobuf:"varint,5,opt,name=free,proto3" json:"free,omitempty"`
	// 温度
	Temperature int32 `protobuf:"varint,6,opt,name=temperature,proto3" json:"temperature,omitempty"`
	// 运行时间（小时）
	PowerOnHours int32           `protobuf:"varint,7,opt,name=power_on_hours,json=powerOnHours,proto3" json:"power_on_hours,omitempty"`
	Health       DiskInfo_Health `protobuf:"varint,8,opt,name=health,proto3,enum=cloud.lazycat.apis.common.DiskInfo_Health" json:"health,omitempty"`
	// 健康状态理由
	HealthReason string `protobuf:"bytes,9,opt,name=health_reason,json=healthReason,proto3" json:"health_reason,omitempty"`
}

func (x *DiskInfo) Reset() {
	*x = DiskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_box_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskInfo) ProtoMessage() {}

func (x *DiskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_box_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskInfo.ProtoReflect.Descriptor instead.
func (*DiskInfo) Descriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{4}
}

func (x *DiskInfo) GetType() DiskInfo_DiskType {
	if x != nil {
		return x.Type
	}
	return DiskInfo_Unknown
}

func (x *DiskInfo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *DiskInfo) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *DiskInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *DiskInfo) GetFree() int64 {
	if x != nil {
		return x.Free
	}
	return 0
}

func (x *DiskInfo) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *DiskInfo) GetPowerOnHours() int32 {
	if x != nil {
		return x.PowerOnHours
	}
	return 0
}

func (x *DiskInfo) GetHealth() DiskInfo_Health {
	if x != nil {
		return x.Health
	}
	return DiskInfo_Normal
}

func (x *DiskInfo) GetHealthReason() string {
	if x != nil {
		return x.HealthReason
	}
	return ""
}

type DisksInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disks []*DiskInfo `protobuf:"bytes,1,rep,name=disks,proto3" json:"disks,omitempty"`
}

func (x *DisksInfo) Reset() {
	*x = DisksInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_box_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisksInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisksInfo) ProtoMessage() {}

func (x *DisksInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_box_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisksInfo.ProtoReflect.Descriptor instead.
func (*DisksInfo) Descriptor() ([]byte, []int) {
	return file_common_box_proto_rawDescGZIP(), []int{5}
}

func (x *DisksInfo) GetDisks() []*DiskInfo {
	if x != nil {
		return x.Disks
	}
	return nil
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
		mi := &file_common_box_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoxInfo_DiskSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoxInfo_DiskSpace) ProtoMessage() {}

func (x *BoxInfo_DiskSpace) ProtoReflect() protoreflect.Message {
	mi := &file_common_box_proto_msgTypes[6]
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
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x0f, 0x53,
	0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64,
	0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x6f, 0x66, 0x66, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x6f, 0x66, 0x74, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x10, 0x02, 0x22, 0xce, 0x01,
	0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x76, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x4f, 0x4f, 0x54,
	0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x18,
	0x0a, 0x14, 0x42, 0x4f, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x4f,
	0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x4f, 0x4f, 0x54,
	0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x02, 0x12,
	0x1d, 0x0a, 0x19, 0x42, 0x4f, 0x4f, 0x54, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46,
	0x41, 0x43, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x03, 0x22, 0xae,
	0x03, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44,
	0x69, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66,
	0x72, 0x65, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6f,
	0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x4f, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x23, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x10, 0x02, 0x22, 0x2a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x0a, 0x0a,
	0x06, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x61, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x61, 0x64, 0x10, 0x02, 0x22,
	0x46, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x05,
	0x64, 0x69, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x32, 0xb1, 0x03, 0x0a, 0x0a, 0x42, 0x6f, 0x78, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x00, 0x12, 0x62, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74, 0x64,
	0x6f, 0x77, 0x6e, 0x12, 0x2a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x44, 0x69, 0x73, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x44, 0x69, 0x73, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73,
	0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_common_box_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_common_box_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_box_proto_goTypes = []interface{}{
	(ShutdownRequest_Action)(0),      // 0: cloud.lazycat.apis.common.ShutdownRequest.Action
	(BootOption_BootOptionType)(0),   // 1: cloud.lazycat.apis.common.BootOption.BootOptionType
	(DiskInfo_DiskType)(0),           // 2: cloud.lazycat.apis.common.DiskInfo.DiskType
	(DiskInfo_Health)(0),             // 3: cloud.lazycat.apis.common.DiskInfo.Health
	(*BoxInfo)(nil),                  // 4: cloud.lazycat.apis.common.BoxInfo
	(*ChangeDisplayNameRequest)(nil), // 5: cloud.lazycat.apis.common.ChangeDisplayNameRequest
	(*ShutdownRequest)(nil),          // 6: cloud.lazycat.apis.common.ShutdownRequest
	(*BootOption)(nil),               // 7: cloud.lazycat.apis.common.BootOption
	(*DiskInfo)(nil),                 // 8: cloud.lazycat.apis.common.DiskInfo
	(*DisksInfo)(nil),                // 9: cloud.lazycat.apis.common.DisksInfo
	(*BoxInfo_DiskSpace)(nil),        // 10: cloud.lazycat.apis.common.BoxInfo.DiskSpace
	(*emptypb.Empty)(nil),            // 11: google.protobuf.Empty
}
var file_common_box_proto_depIdxs = []int32{
	10, // 0: cloud.lazycat.apis.common.BoxInfo.disk_space:type_name -> cloud.lazycat.apis.common.BoxInfo.DiskSpace
	0,  // 1: cloud.lazycat.apis.common.ShutdownRequest.action:type_name -> cloud.lazycat.apis.common.ShutdownRequest.Action
	1,  // 2: cloud.lazycat.apis.common.BootOption.type:type_name -> cloud.lazycat.apis.common.BootOption.BootOptionType
	2,  // 3: cloud.lazycat.apis.common.DiskInfo.type:type_name -> cloud.lazycat.apis.common.DiskInfo.DiskType
	3,  // 4: cloud.lazycat.apis.common.DiskInfo.health:type_name -> cloud.lazycat.apis.common.DiskInfo.Health
	8,  // 5: cloud.lazycat.apis.common.DisksInfo.disks:type_name -> cloud.lazycat.apis.common.DiskInfo
	11, // 6: cloud.lazycat.apis.common.BoxService.QueryInfo:input_type -> google.protobuf.Empty
	5,  // 7: cloud.lazycat.apis.common.BoxService.ChangeDisplayName:input_type -> cloud.lazycat.apis.common.ChangeDisplayNameRequest
	7,  // 8: cloud.lazycat.apis.common.BoxService.SetBootOption:input_type -> cloud.lazycat.apis.common.BootOption
	6,  // 9: cloud.lazycat.apis.common.BoxService.Shutdown:input_type -> cloud.lazycat.apis.common.ShutdownRequest
	11, // 10: cloud.lazycat.apis.common.BoxService.QueryDisksInfo:input_type -> google.protobuf.Empty
	4,  // 11: cloud.lazycat.apis.common.BoxService.QueryInfo:output_type -> cloud.lazycat.apis.common.BoxInfo
	11, // 12: cloud.lazycat.apis.common.BoxService.ChangeDisplayName:output_type -> google.protobuf.Empty
	11, // 13: cloud.lazycat.apis.common.BoxService.SetBootOption:output_type -> google.protobuf.Empty
	11, // 14: cloud.lazycat.apis.common.BoxService.Shutdown:output_type -> google.protobuf.Empty
	9,  // 15: cloud.lazycat.apis.common.BoxService.QueryDisksInfo:output_type -> cloud.lazycat.apis.common.DisksInfo
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
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
			switch v := v.(*BootOption); i {
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
		file_common_box_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskInfo); i {
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
		file_common_box_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisksInfo); i {
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
		file_common_box_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			NumEnums:      4,
			NumMessages:   7,
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

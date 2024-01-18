// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: common/file_handler.proto

package common

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

type IntentAction int32

const (
	// 未知
	IntentAction_UN_KNOWN IntentAction = 0
	// 打开
	IntentAction_OPEN IntentAction = 1
	// 分享
	IntentAction_SHARE IntentAction = 2
)

// Enum value maps for IntentAction.
var (
	IntentAction_name = map[int32]string{
		0: "UN_KNOWN",
		1: "OPEN",
		2: "SHARE",
	}
	IntentAction_value = map[string]int32{
		"UN_KNOWN": 0,
		"OPEN":     1,
		"SHARE":    2,
	}
)

func (x IntentAction) Enum() *IntentAction {
	p := new(IntentAction)
	*p = x
	return p
}

func (x IntentAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IntentAction) Descriptor() protoreflect.EnumDescriptor {
	return file_common_file_handler_proto_enumTypes[0].Descriptor()
}

func (IntentAction) Type() protoreflect.EnumType {
	return &file_common_file_handler_proto_enumTypes[0]
}

func (x IntentAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IntentAction.Descriptor instead.
func (IntentAction) EnumDescriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{0}
}

type WalkDirRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 目录路径
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// 过滤条件
	Pattern *string `protobuf:"bytes,2,opt,name=pattern,proto3,oneof" json:"pattern,omitempty"`
	// 并行线程数，不传或<=0的话，就等于$ nproc
	WorkerCount *int32 `protobuf:"varint,3,opt,name=worker_count,json=workerCount,proto3,oneof" json:"worker_count,omitempty"`
}

func (x *WalkDirRequest) Reset() {
	*x = WalkDirRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalkDirRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalkDirRequest) ProtoMessage() {}

func (x *WalkDirRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalkDirRequest.ProtoReflect.Descriptor instead.
func (*WalkDirRequest) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{0}
}

func (x *WalkDirRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *WalkDirRequest) GetPattern() string {
	if x != nil && x.Pattern != nil {
		return *x.Pattern
	}
	return ""
}

func (x *WalkDirRequest) GetWorkerCount() int32 {
	if x != nil && x.WorkerCount != nil {
		return *x.WorkerCount
	}
	return 0
}

type WalkDirReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 全路径 /home/lnks/hello.txt
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// 文件名 hello.txt
	Base string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	// 文件大小
	Size int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// 是否为目录
	IsDir bool `protobuf:"varint,4,opt,name=isDir,proto3" json:"isDir,omitempty"`
	// 文件修改时间
	Mtime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=mtime,proto3" json:"mtime,omitempty"`
}

func (x *WalkDirReply) Reset() {
	*x = WalkDirReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalkDirReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalkDirReply) ProtoMessage() {}

func (x *WalkDirReply) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalkDirReply.ProtoReflect.Descriptor instead.
func (*WalkDirReply) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{1}
}

func (x *WalkDirReply) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *WalkDirReply) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *WalkDirReply) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *WalkDirReply) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

func (x *WalkDirReply) GetMtime() *timestamppb.Timestamp {
	if x != nil {
		return x.Mtime
	}
	return nil
}

type CreateDirRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CreateDirRequest) Reset() {
	*x = CreateDirRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirRequest) ProtoMessage() {}

func (x *CreateDirRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirRequest.ProtoReflect.Descriptor instead.
func (*CreateDirRequest) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDirRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type OpenFileManagerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件(夹)绝对路径
	// 1. window 下面的必须以盘符开头，如 C:\Users\lnks\Desktop
	// 2. unix   下面的必须以 / 开头，如 /home/lnks/Desktop
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *OpenFileManagerRequest) Reset() {
	*x = OpenFileManagerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenFileManagerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenFileManagerRequest) ProtoMessage() {}

func (x *OpenFileManagerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenFileManagerRequest.ProtoReflect.Descriptor instead.
func (*OpenFileManagerRequest) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{3}
}

func (x *OpenFileManagerRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type OpenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 应用id
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// 文件路径
	Path         string       `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	IntentAction IntentAction `protobuf:"varint,4,opt,name=intent_action,json=intentAction,proto3,enum=cloud.lazycat.apis.common.IntentAction" json:"intent_action,omitempty"`
}

func (x *OpenRequest) Reset() {
	*x = OpenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenRequest) ProtoMessage() {}

func (x *OpenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenRequest.ProtoReflect.Descriptor instead.
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{4}
}

func (x *OpenRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *OpenRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *OpenRequest) GetIntentAction() IntentAction {
	if x != nil {
		return x.IntentAction
	}
	return IntentAction_UN_KNOWN
}

type OpenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 打开应用的地址
	Url *string `protobuf:"bytes,1,opt,name=url,proto3,oneof" json:"url,omitempty"`
}

func (x *OpenReply) Reset() {
	*x = OpenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenReply) ProtoMessage() {}

func (x *OpenReply) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenReply.ProtoReflect.Descriptor instead.
func (*OpenReply) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{5}
}

func (x *OpenReply) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件mime类型
	Mime string `protobuf:"bytes,1,opt,name=mime,proto3" json:"mime,omitempty"`
	// 文件路径(可选 android下面使用path查询，体验更好 ）
	Path *string `protobuf:"bytes,2,opt,name=path,proto3,oneof" json:"path,omitempty"`
	// 期望的 app icon 尺寸(宽高一致）
	IconSize *int32 `protobuf:"varint,3,opt,name=icon_size,json=iconSize,proto3,oneof" json:"icon_size,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{6}
}

func (x *QueryRequest) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *QueryRequest) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *QueryRequest) GetIconSize() int32 {
	if x != nil && x.IconSize != nil {
		return *x.IconSize
	}
	return 0
}

type QueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppList []*AppShortcut `protobuf:"bytes,1,rep,name=app_list,json=appList,proto3" json:"app_list,omitempty"`
}

func (x *QueryReply) Reset() {
	*x = QueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReply) ProtoMessage() {}

func (x *QueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReply.ProtoReflect.Descriptor instead.
func (*QueryReply) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{7}
}

func (x *QueryReply) GetAppList() []*AppShortcut {
	if x != nil {
		return x.AppList
	}
	return nil
}

type AppShortcut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId        string       `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Title        string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Icon         string       `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	IntentAction IntentAction `protobuf:"varint,4,opt,name=intent_action,json=intentAction,proto3,enum=cloud.lazycat.apis.common.IntentAction" json:"intent_action,omitempty"`
}

func (x *AppShortcut) Reset() {
	*x = AppShortcut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppShortcut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppShortcut) ProtoMessage() {}

func (x *AppShortcut) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppShortcut.ProtoReflect.Descriptor instead.
func (*AppShortcut) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{8}
}

func (x *AppShortcut) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *AppShortcut) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AppShortcut) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *AppShortcut) GetIntentAction() IntentAction {
	if x != nil {
		return x.IntentAction
	}
	return IntentAction_UN_KNOWN
}

type StatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *StatRequest) Reset() {
	*x = StatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatRequest) ProtoMessage() {}

func (x *StatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatRequest.ProtoReflect.Descriptor instead.
func (*StatRequest) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{9}
}

func (x *StatRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// StatReply 和 golang fs.FileInfo 中字段保持一致
type StatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size    int64                  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Mode    uint32                 `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	ModTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=modTime,proto3" json:"modTime,omitempty"`
	IsDir   bool                   `protobuf:"varint,5,opt,name=isDir,proto3" json:"isDir,omitempty"`
}

func (x *StatReply) Reset() {
	*x = StatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_file_handler_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatReply) ProtoMessage() {}

func (x *StatReply) ProtoReflect() protoreflect.Message {
	mi := &file_common_file_handler_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatReply.ProtoReflect.Descriptor instead.
func (*StatReply) Descriptor() ([]byte, []int) {
	return file_common_file_handler_proto_rawDescGZIP(), []int{10}
}

func (x *StatReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StatReply) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *StatReply) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *StatReply) GetModTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ModTime
	}
	return nil
}

func (x *StatReply) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

var File_common_file_handler_proto protoreflect.FileDescriptor

var file_common_file_handler_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x57, 0x61, 0x6c, 0x6b, 0x44, 0x69, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x01, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x92, 0x01, 0x0a, 0x0c, 0x57, 0x61, 0x6c, 0x6b, 0x44, 0x69, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x73, 0x44, 0x69, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44,
	0x69, 0x72, 0x12, 0x30, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x6d,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2c, 0x0a, 0x16,
	0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x86, 0x01, 0x0a, 0x0b, 0x4f,
	0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x4c, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x22,
	0x74, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09,
	0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x01, 0x52, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x63, 0x6f, 0x6e,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x4f, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61,
	0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x52, 0x07, 0x61,
	0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x93, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x61,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x6d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x44, 0x69,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44, 0x69, 0x72, 0x2a, 0x31,
	0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c,
	0x0a, 0x08, 0x55, 0x4e, 0x5f, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x48, 0x41, 0x52, 0x45, 0x10,
	0x02, 0x32, 0xb1, 0x04, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x12, 0x59, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x27, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x04,
	0x6f, 0x70, 0x65, 0x6e, 0x12, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x07, 0x77, 0x61, 0x6c, 0x6b, 0x44, 0x69, 0x72, 0x12,
	0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x57, 0x61, 0x6c, 0x6b,
	0x44, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x57, 0x61, 0x6c, 0x6b, 0x44, 0x69, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x54, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x72, 0x12, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x56, 0x0a,
	0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61,
	0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_file_handler_proto_rawDescOnce sync.Once
	file_common_file_handler_proto_rawDescData = file_common_file_handler_proto_rawDesc
)

func file_common_file_handler_proto_rawDescGZIP() []byte {
	file_common_file_handler_proto_rawDescOnce.Do(func() {
		file_common_file_handler_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_file_handler_proto_rawDescData)
	})
	return file_common_file_handler_proto_rawDescData
}

var file_common_file_handler_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_file_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_common_file_handler_proto_goTypes = []interface{}{
	(IntentAction)(0),              // 0: cloud.lazycat.apis.common.IntentAction
	(*WalkDirRequest)(nil),         // 1: cloud.lazycat.apis.common.WalkDirRequest
	(*WalkDirReply)(nil),           // 2: cloud.lazycat.apis.common.WalkDirReply
	(*CreateDirRequest)(nil),       // 3: cloud.lazycat.apis.common.CreateDirRequest
	(*OpenFileManagerRequest)(nil), // 4: cloud.lazycat.apis.common.OpenFileManagerRequest
	(*OpenRequest)(nil),            // 5: cloud.lazycat.apis.common.OpenRequest
	(*OpenReply)(nil),              // 6: cloud.lazycat.apis.common.OpenReply
	(*QueryRequest)(nil),           // 7: cloud.lazycat.apis.common.QueryRequest
	(*QueryReply)(nil),             // 8: cloud.lazycat.apis.common.QueryReply
	(*AppShortcut)(nil),            // 9: cloud.lazycat.apis.common.AppShortcut
	(*StatRequest)(nil),            // 10: cloud.lazycat.apis.common.StatRequest
	(*StatReply)(nil),              // 11: cloud.lazycat.apis.common.StatReply
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 13: google.protobuf.Empty
}
var file_common_file_handler_proto_depIdxs = []int32{
	12, // 0: cloud.lazycat.apis.common.WalkDirReply.mtime:type_name -> google.protobuf.Timestamp
	0,  // 1: cloud.lazycat.apis.common.OpenRequest.intent_action:type_name -> cloud.lazycat.apis.common.IntentAction
	9,  // 2: cloud.lazycat.apis.common.QueryReply.app_list:type_name -> cloud.lazycat.apis.common.AppShortcut
	0,  // 3: cloud.lazycat.apis.common.AppShortcut.intent_action:type_name -> cloud.lazycat.apis.common.IntentAction
	12, // 4: cloud.lazycat.apis.common.StatReply.modTime:type_name -> google.protobuf.Timestamp
	7,  // 5: cloud.lazycat.apis.common.FileHandler.query:input_type -> cloud.lazycat.apis.common.QueryRequest
	5,  // 6: cloud.lazycat.apis.common.FileHandler.open:input_type -> cloud.lazycat.apis.common.OpenRequest
	4,  // 7: cloud.lazycat.apis.common.FileHandler.openFileManager:input_type -> cloud.lazycat.apis.common.OpenFileManagerRequest
	1,  // 8: cloud.lazycat.apis.common.FileHandler.walkDir:input_type -> cloud.lazycat.apis.common.WalkDirRequest
	3,  // 9: cloud.lazycat.apis.common.FileHandler.createDir:input_type -> cloud.lazycat.apis.common.CreateDirRequest
	10, // 10: cloud.lazycat.apis.common.FileHandler.stat:input_type -> cloud.lazycat.apis.common.StatRequest
	8,  // 11: cloud.lazycat.apis.common.FileHandler.query:output_type -> cloud.lazycat.apis.common.QueryReply
	6,  // 12: cloud.lazycat.apis.common.FileHandler.open:output_type -> cloud.lazycat.apis.common.OpenReply
	13, // 13: cloud.lazycat.apis.common.FileHandler.openFileManager:output_type -> google.protobuf.Empty
	2,  // 14: cloud.lazycat.apis.common.FileHandler.walkDir:output_type -> cloud.lazycat.apis.common.WalkDirReply
	13, // 15: cloud.lazycat.apis.common.FileHandler.createDir:output_type -> google.protobuf.Empty
	11, // 16: cloud.lazycat.apis.common.FileHandler.stat:output_type -> cloud.lazycat.apis.common.StatReply
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_common_file_handler_proto_init() }
func file_common_file_handler_proto_init() {
	if File_common_file_handler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_file_handler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalkDirRequest); i {
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
		file_common_file_handler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalkDirReply); i {
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
		file_common_file_handler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDirRequest); i {
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
		file_common_file_handler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenFileManagerRequest); i {
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
		file_common_file_handler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenRequest); i {
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
		file_common_file_handler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenReply); i {
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
		file_common_file_handler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_common_file_handler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReply); i {
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
		file_common_file_handler_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppShortcut); i {
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
		file_common_file_handler_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatRequest); i {
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
		file_common_file_handler_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatReply); i {
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
	file_common_file_handler_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_common_file_handler_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_common_file_handler_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_file_handler_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_file_handler_proto_goTypes,
		DependencyIndexes: file_common_file_handler_proto_depIdxs,
		EnumInfos:         file_common_file_handler_proto_enumTypes,
		MessageInfos:      file_common_file_handler_proto_msgTypes,
	}.Build()
	File_common_file_handler_proto = out.File
	file_common_file_handler_proto_rawDesc = nil
	file_common_file_handler_proto_goTypes = nil
	file_common_file_handler_proto_depIdxs = nil
}

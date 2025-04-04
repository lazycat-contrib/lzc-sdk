// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: sys/btrfs.proto

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

type BtrfsSubvolCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Force bool   `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *BtrfsSubvolCreateRequest) Reset() {
	*x = BtrfsSubvolCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_btrfs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BtrfsSubvolCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BtrfsSubvolCreateRequest) ProtoMessage() {}

func (x *BtrfsSubvolCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_btrfs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BtrfsSubvolCreateRequest.ProtoReflect.Descriptor instead.
func (*BtrfsSubvolCreateRequest) Descriptor() ([]byte, []int) {
	return file_sys_btrfs_proto_rawDescGZIP(), []int{0}
}

func (x *BtrfsSubvolCreateRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *BtrfsSubvolCreateRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type BtrfsSubvolInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *BtrfsSubvolInfoRequest) Reset() {
	*x = BtrfsSubvolInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_btrfs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BtrfsSubvolInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BtrfsSubvolInfoRequest) ProtoMessage() {}

func (x *BtrfsSubvolInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_btrfs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BtrfsSubvolInfoRequest.ProtoReflect.Descriptor instead.
func (*BtrfsSubvolInfoRequest) Descriptor() ([]byte, []int) {
	return file_sys_btrfs_proto_rawDescGZIP(), []int{1}
}

func (x *BtrfsSubvolInfoRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type BtrfsSubvolInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Gen  uint64 `protobuf:"varint,2,opt,name=gen,proto3" json:"gen,omitempty"`
}

func (x *BtrfsSubvolInfoResponse) Reset() {
	*x = BtrfsSubvolInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_btrfs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BtrfsSubvolInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BtrfsSubvolInfoResponse) ProtoMessage() {}

func (x *BtrfsSubvolInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_btrfs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BtrfsSubvolInfoResponse.ProtoReflect.Descriptor instead.
func (*BtrfsSubvolInfoResponse) Descriptor() ([]byte, []int) {
	return file_sys_btrfs_proto_rawDescGZIP(), []int{2}
}

func (x *BtrfsSubvolInfoResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *BtrfsSubvolInfoResponse) GetGen() uint64 {
	if x != nil {
		return x.Gen
	}
	return 0
}

type BtrfsSubvolFindNewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Gen  uint64 `protobuf:"varint,2,opt,name=gen,proto3" json:"gen,omitempty"`
}

func (x *BtrfsSubvolFindNewRequest) Reset() {
	*x = BtrfsSubvolFindNewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_btrfs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BtrfsSubvolFindNewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BtrfsSubvolFindNewRequest) ProtoMessage() {}

func (x *BtrfsSubvolFindNewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_btrfs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BtrfsSubvolFindNewRequest.ProtoReflect.Descriptor instead.
func (*BtrfsSubvolFindNewRequest) Descriptor() ([]byte, []int) {
	return file_sys_btrfs_proto_rawDescGZIP(), []int{3}
}

func (x *BtrfsSubvolFindNewRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *BtrfsSubvolFindNewRequest) GetGen() uint64 {
	if x != nil {
		return x.Gen
	}
	return 0
}

type BtrfsSubvolFindNewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Gen   uint64   `protobuf:"varint,2,opt,name=gen,proto3" json:"gen,omitempty"`
}

func (x *BtrfsSubvolFindNewResponse) Reset() {
	*x = BtrfsSubvolFindNewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_btrfs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BtrfsSubvolFindNewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BtrfsSubvolFindNewResponse) ProtoMessage() {}

func (x *BtrfsSubvolFindNewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_btrfs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BtrfsSubvolFindNewResponse.ProtoReflect.Descriptor instead.
func (*BtrfsSubvolFindNewResponse) Descriptor() ([]byte, []int) {
	return file_sys_btrfs_proto_rawDescGZIP(), []int{4}
}

func (x *BtrfsSubvolFindNewResponse) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *BtrfsSubvolFindNewResponse) GetGen() uint64 {
	if x != nil {
		return x.Gen
	}
	return 0
}

type BtrfsRenameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcPath string `protobuf:"bytes,1,opt,name=src_path,json=srcPath,proto3" json:"src_path,omitempty"`
	// src and dest path must reside in same subvolume
	// otherwise rename/reflink will fail
	DestPath string `protobuf:"bytes,2,opt,name=dest_path,json=destPath,proto3" json:"dest_path,omitempty"`
	// reflink instead of rename
	Reflink bool `protobuf:"varint,3,opt,name=reflink,proto3" json:"reflink,omitempty"`
}

func (x *BtrfsRenameRequest) Reset() {
	*x = BtrfsRenameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_btrfs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BtrfsRenameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BtrfsRenameRequest) ProtoMessage() {}

func (x *BtrfsRenameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_btrfs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BtrfsRenameRequest.ProtoReflect.Descriptor instead.
func (*BtrfsRenameRequest) Descriptor() ([]byte, []int) {
	return file_sys_btrfs_proto_rawDescGZIP(), []int{5}
}

func (x *BtrfsRenameRequest) GetSrcPath() string {
	if x != nil {
		return x.SrcPath
	}
	return ""
}

func (x *BtrfsRenameRequest) GetDestPath() string {
	if x != nil {
		return x.DestPath
	}
	return ""
}

func (x *BtrfsRenameRequest) GetReflink() bool {
	if x != nil {
		return x.Reflink
	}
	return false
}

var File_sys_btrfs_proto protoreflect.FileDescriptor

var file_sys_btrfs_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x79, 0x73, 0x2f, 0x62, 0x74, 0x72, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x18, 0x42, 0x74, 0x72, 0x66, 0x73, 0x53,
	0x75, 0x62, 0x76, 0x6f, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x2c, 0x0a, 0x16,
	0x42, 0x74, 0x72, 0x66, 0x73, 0x53, 0x75, 0x62, 0x76, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x3f, 0x0a, 0x17, 0x42, 0x74,
	0x72, 0x66, 0x73, 0x53, 0x75, 0x62, 0x76, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x67, 0x65, 0x6e, 0x22, 0x41, 0x0a, 0x19, 0x42,
	0x74, 0x72, 0x66, 0x73, 0x53, 0x75, 0x62, 0x76, 0x6f, 0x6c, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03,
	0x67, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x67, 0x65, 0x6e, 0x22, 0x44,
	0x0a, 0x1a, 0x42, 0x74, 0x72, 0x66, 0x73, 0x53, 0x75, 0x62, 0x76, 0x6f, 0x6c, 0x46, 0x69, 0x6e,
	0x64, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x67, 0x65, 0x6e, 0x22, 0x66, 0x0a, 0x12, 0x42, 0x74, 0x72, 0x66, 0x73, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72,
	0x63, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72,
	0x63, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0xa2, 0x03, 0x0a,
	0x09, 0x42, 0x74, 0x72, 0x66, 0x73, 0x55, 0x74, 0x69, 0x6c, 0x12, 0x5a, 0x0a, 0x0c, 0x53, 0x75,
	0x62, 0x76, 0x6f, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x2e, 0x42, 0x74, 0x72, 0x66, 0x73, 0x53, 0x75, 0x62, 0x76, 0x6f, 0x6c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x76, 0x6f, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42, 0x74,
	0x72, 0x66, 0x73, 0x53, 0x75, 0x62, 0x76, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42, 0x74,
	0x72, 0x66, 0x73, 0x53, 0x75, 0x62, 0x76, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x76, 0x6f,
	0x6c, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x12, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79,
	0x73, 0x2e, 0x42, 0x74, 0x72, 0x66, 0x73, 0x53, 0x75, 0x62, 0x76, 0x6f, 0x6c, 0x46, 0x69, 0x6e,
	0x64, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42, 0x74, 0x72, 0x66, 0x73, 0x53, 0x75, 0x62, 0x76, 0x6f, 0x6c,
	0x46, 0x69, 0x6e, 0x64, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x06, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x73, 0x79, 0x73, 0x2e, 0x42, 0x74, 0x72, 0x66, 0x73, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x79, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sys_btrfs_proto_rawDescOnce sync.Once
	file_sys_btrfs_proto_rawDescData = file_sys_btrfs_proto_rawDesc
)

func file_sys_btrfs_proto_rawDescGZIP() []byte {
	file_sys_btrfs_proto_rawDescOnce.Do(func() {
		file_sys_btrfs_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_btrfs_proto_rawDescData)
	})
	return file_sys_btrfs_proto_rawDescData
}

var file_sys_btrfs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sys_btrfs_proto_goTypes = []interface{}{
	(*BtrfsSubvolCreateRequest)(nil),   // 0: cloud.lazycat.apis.sys.BtrfsSubvolCreateRequest
	(*BtrfsSubvolInfoRequest)(nil),     // 1: cloud.lazycat.apis.sys.BtrfsSubvolInfoRequest
	(*BtrfsSubvolInfoResponse)(nil),    // 2: cloud.lazycat.apis.sys.BtrfsSubvolInfoResponse
	(*BtrfsSubvolFindNewRequest)(nil),  // 3: cloud.lazycat.apis.sys.BtrfsSubvolFindNewRequest
	(*BtrfsSubvolFindNewResponse)(nil), // 4: cloud.lazycat.apis.sys.BtrfsSubvolFindNewResponse
	(*BtrfsRenameRequest)(nil),         // 5: cloud.lazycat.apis.sys.BtrfsRenameRequest
	(*emptypb.Empty)(nil),              // 6: google.protobuf.Empty
}
var file_sys_btrfs_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.apis.sys.BtrfsUtil.SubvolCreate:input_type -> cloud.lazycat.apis.sys.BtrfsSubvolCreateRequest
	1, // 1: cloud.lazycat.apis.sys.BtrfsUtil.SubvolInfo:input_type -> cloud.lazycat.apis.sys.BtrfsSubvolInfoRequest
	3, // 2: cloud.lazycat.apis.sys.BtrfsUtil.SubvolFindNew:input_type -> cloud.lazycat.apis.sys.BtrfsSubvolFindNewRequest
	5, // 3: cloud.lazycat.apis.sys.BtrfsUtil.Rename:input_type -> cloud.lazycat.apis.sys.BtrfsRenameRequest
	6, // 4: cloud.lazycat.apis.sys.BtrfsUtil.SubvolCreate:output_type -> google.protobuf.Empty
	2, // 5: cloud.lazycat.apis.sys.BtrfsUtil.SubvolInfo:output_type -> cloud.lazycat.apis.sys.BtrfsSubvolInfoResponse
	4, // 6: cloud.lazycat.apis.sys.BtrfsUtil.SubvolFindNew:output_type -> cloud.lazycat.apis.sys.BtrfsSubvolFindNewResponse
	6, // 7: cloud.lazycat.apis.sys.BtrfsUtil.Rename:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sys_btrfs_proto_init() }
func file_sys_btrfs_proto_init() {
	if File_sys_btrfs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_btrfs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BtrfsSubvolCreateRequest); i {
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
		file_sys_btrfs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BtrfsSubvolInfoRequest); i {
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
		file_sys_btrfs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BtrfsSubvolInfoResponse); i {
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
		file_sys_btrfs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BtrfsSubvolFindNewRequest); i {
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
		file_sys_btrfs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BtrfsSubvolFindNewResponse); i {
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
		file_sys_btrfs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BtrfsRenameRequest); i {
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
			RawDescriptor: file_sys_btrfs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_btrfs_proto_goTypes,
		DependencyIndexes: file_sys_btrfs_proto_depIdxs,
		MessageInfos:      file_sys_btrfs_proto_msgTypes,
	}.Build()
	File_sys_btrfs_proto = out.File
	file_sys_btrfs_proto_rawDesc = nil
	file_sys_btrfs_proto_goTypes = nil
	file_sys_btrfs_proto_depIdxs = nil
}

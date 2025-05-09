// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: sys/ingress.proto

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

type AppAccessPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string           `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"` //若为空则从context中获取实际UID
	Policy *AppAccessPolicy `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *AppAccessPolicyRequest) Reset() {
	*x = AppAccessPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_ingress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppAccessPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppAccessPolicyRequest) ProtoMessage() {}

func (x *AppAccessPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_ingress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppAccessPolicyRequest.ProtoReflect.Descriptor instead.
func (*AppAccessPolicyRequest) Descriptor() ([]byte, []int) {
	return file_sys_ingress_proto_rawDescGZIP(), []int{0}
}

func (x *AppAccessPolicyRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AppAccessPolicyRequest) GetPolicy() *AppAccessPolicy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type AppAccessPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否允许访问所有应用。
	// 此字段如果出现，则allow_access_apps字段会被忽略，即旧值不会有任何调整变化。
	//
	// 当目标用户有安装应用的权限时，无法设置no_limit=false。其他情况no_limit的值无任何限制。
	NoLimit *bool `protobuf:"varint,1,opt,name=no_limit,json=noLimit,proto3,oneof" json:"no_limit,omitempty"`
	// 允许访问的appid列表
	AllowAccessAppids []string `protobuf:"bytes,2,rep,name=allow_access_appids,json=allowAccessAppids,proto3" json:"allow_access_appids,omitempty"`
}

func (x *AppAccessPolicy) Reset() {
	*x = AppAccessPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_ingress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppAccessPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppAccessPolicy) ProtoMessage() {}

func (x *AppAccessPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_sys_ingress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppAccessPolicy.ProtoReflect.Descriptor instead.
func (*AppAccessPolicy) Descriptor() ([]byte, []int) {
	return file_sys_ingress_proto_rawDescGZIP(), []int{1}
}

func (x *AppAccessPolicy) GetNoLimit() bool {
	if x != nil && x.NoLimit != nil {
		return *x.NoLimit
	}
	return false
}

func (x *AppAccessPolicy) GetAllowAccessAppids() []string {
	if x != nil {
		return x.AllowAccessAppids
	}
	return nil
}

type IngressAppLastAccessTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=AppId,proto3" json:"AppId,omitempty"`
}

func (x *IngressAppLastAccessTimeRequest) Reset() {
	*x = IngressAppLastAccessTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_ingress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngressAppLastAccessTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngressAppLastAccessTimeRequest) ProtoMessage() {}

func (x *IngressAppLastAccessTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_ingress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngressAppLastAccessTimeRequest.ProtoReflect.Descriptor instead.
func (*IngressAppLastAccessTimeRequest) Descriptor() ([]byte, []int) {
	return file_sys_ingress_proto_rawDescGZIP(), []int{2}
}

func (x *IngressAppLastAccessTimeRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type IngressAppLastAccessTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnixTime int64 `protobuf:"varint,1,opt,name=UnixTime,proto3" json:"UnixTime,omitempty"`
}

func (x *IngressAppLastAccessTimeResponse) Reset() {
	*x = IngressAppLastAccessTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_ingress_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngressAppLastAccessTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngressAppLastAccessTimeResponse) ProtoMessage() {}

func (x *IngressAppLastAccessTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_ingress_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngressAppLastAccessTimeResponse.ProtoReflect.Descriptor instead.
func (*IngressAppLastAccessTimeResponse) Descriptor() ([]byte, []int) {
	return file_sys_ingress_proto_rawDescGZIP(), []int{3}
}

func (x *IngressAppLastAccessTimeResponse) GetUnixTime() int64 {
	if x != nil {
		return x.UnixTime
	}
	return 0
}

type SubscribeEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeEventRequest) Reset() {
	*x = SubscribeEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_ingress_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeEventRequest) ProtoMessage() {}

func (x *SubscribeEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_ingress_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeEventRequest.ProtoReflect.Descriptor instead.
func (*SubscribeEventRequest) Descriptor() ([]byte, []int) {
	return file_sys_ingress_proto_rawDescGZIP(), []int{4}
}

type SubscribeEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeEventResponse) Reset() {
	*x = SubscribeEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_ingress_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeEventResponse) ProtoMessage() {}

func (x *SubscribeEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_ingress_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeEventResponse.ProtoReflect.Descriptor instead.
func (*SubscribeEventResponse) Descriptor() ([]byte, []int) {
	return file_sys_ingress_proto_rawDescGZIP(), []int{5}
}

var File_sys_ingress_proto protoreflect.FileDescriptor

var file_sys_ingress_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x79, 0x73, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x70,
	0x70, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x6e, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x08, 0x6e, 0x6f, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x6f,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x70, 0x70, 0x69, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x41, 0x70, 0x70, 0x69, 0x64, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x6f, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x37, 0x0a, 0x1f, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x41, 0x70, 0x70, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x22, 0x3e,
	0x0a, 0x20, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x70, 0x70, 0x4c, 0x61, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x17,
	0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xf5, 0x02, 0x0a, 0x16, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x12,
	0x53, 0x65, 0x74, 0x41, 0x70, 0x70, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x6f, 0x0a, 0x14, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x8b, 0x01, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x70, 0x70, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x70,
	0x70, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x87, 0x01, 0x0a, 0x12, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x71, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x79, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_ingress_proto_rawDescOnce sync.Once
	file_sys_ingress_proto_rawDescData = file_sys_ingress_proto_rawDesc
)

func file_sys_ingress_proto_rawDescGZIP() []byte {
	file_sys_ingress_proto_rawDescOnce.Do(func() {
		file_sys_ingress_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_ingress_proto_rawDescData)
	})
	return file_sys_ingress_proto_rawDescData
}

var file_sys_ingress_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sys_ingress_proto_goTypes = []interface{}{
	(*AppAccessPolicyRequest)(nil),           // 0: cloud.lazycat.apis.sys.AppAccessPolicyRequest
	(*AppAccessPolicy)(nil),                  // 1: cloud.lazycat.apis.sys.AppAccessPolicy
	(*IngressAppLastAccessTimeRequest)(nil),  // 2: cloud.lazycat.apis.sys.IngressAppLastAccessTimeRequest
	(*IngressAppLastAccessTimeResponse)(nil), // 3: cloud.lazycat.apis.sys.IngressAppLastAccessTimeResponse
	(*SubscribeEventRequest)(nil),            // 4: cloud.lazycat.apis.sys.SubscribeEventRequest
	(*SubscribeEventResponse)(nil),           // 5: cloud.lazycat.apis.sys.SubscribeEventResponse
	(*emptypb.Empty)(nil),                    // 6: google.protobuf.Empty
}
var file_sys_ingress_proto_depIdxs = []int32{
	1, // 0: cloud.lazycat.apis.sys.AppAccessPolicyRequest.policy:type_name -> cloud.lazycat.apis.sys.AppAccessPolicy
	0, // 1: cloud.lazycat.apis.sys.AccessControlerService.SetAppAccessPolicy:input_type -> cloud.lazycat.apis.sys.AppAccessPolicyRequest
	0, // 2: cloud.lazycat.apis.sys.AccessControlerService.QueryAppAccessPolicy:input_type -> cloud.lazycat.apis.sys.AppAccessPolicyRequest
	2, // 3: cloud.lazycat.apis.sys.AccessControlerService.GetAppLastAccessTime:input_type -> cloud.lazycat.apis.sys.IngressAppLastAccessTimeRequest
	4, // 4: cloud.lazycat.apis.sys.UserSessionService.SubscribeEvent:input_type -> cloud.lazycat.apis.sys.SubscribeEventRequest
	6, // 5: cloud.lazycat.apis.sys.AccessControlerService.SetAppAccessPolicy:output_type -> google.protobuf.Empty
	1, // 6: cloud.lazycat.apis.sys.AccessControlerService.QueryAppAccessPolicy:output_type -> cloud.lazycat.apis.sys.AppAccessPolicy
	3, // 7: cloud.lazycat.apis.sys.AccessControlerService.GetAppLastAccessTime:output_type -> cloud.lazycat.apis.sys.IngressAppLastAccessTimeResponse
	5, // 8: cloud.lazycat.apis.sys.UserSessionService.SubscribeEvent:output_type -> cloud.lazycat.apis.sys.SubscribeEventResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sys_ingress_proto_init() }
func file_sys_ingress_proto_init() {
	if File_sys_ingress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_ingress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppAccessPolicyRequest); i {
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
		file_sys_ingress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppAccessPolicy); i {
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
		file_sys_ingress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngressAppLastAccessTimeRequest); i {
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
		file_sys_ingress_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngressAppLastAccessTimeResponse); i {
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
		file_sys_ingress_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeEventRequest); i {
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
		file_sys_ingress_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeEventResponse); i {
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
	file_sys_ingress_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sys_ingress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_sys_ingress_proto_goTypes,
		DependencyIndexes: file_sys_ingress_proto_depIdxs,
		MessageInfos:      file_sys_ingress_proto_msgTypes,
	}.Build()
	File_sys_ingress_proto = out.File
	file_sys_ingress_proto_rawDesc = nil
	file_sys_ingress_proto_goTypes = nil
	file_sys_ingress_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: sys/network_manager.proto

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

type NetworkDeviceStatus int32

const (
	NetworkDeviceStatus_NetworkDeviceStatusUnavailable   NetworkDeviceStatus = 0 // 不可用
	NetworkDeviceStatus_NetworkDeviceStatusDisconnected  NetworkDeviceStatus = 1 // 未连接
	NetworkDeviceStatus_NetworkDeviceStatusConnecting    NetworkDeviceStatus = 2 // 正在连接
	NetworkDeviceStatus_NetworkDeviceStatusConnected     NetworkDeviceStatus = 3 // 已连接
	NetworkDeviceStatus_NetworkDeviceStatusDisconnecting NetworkDeviceStatus = 4 // 正在断开
	NetworkDeviceStatus_NetworkDeviceStatusDisabled      NetworkDeviceStatus = 5 // 已禁用
)

// Enum value maps for NetworkDeviceStatus.
var (
	NetworkDeviceStatus_name = map[int32]string{
		0: "NetworkDeviceStatusUnavailable",
		1: "NetworkDeviceStatusDisconnected",
		2: "NetworkDeviceStatusConnecting",
		3: "NetworkDeviceStatusConnected",
		4: "NetworkDeviceStatusDisconnecting",
		5: "NetworkDeviceStatusDisabled",
	}
	NetworkDeviceStatus_value = map[string]int32{
		"NetworkDeviceStatusUnavailable":   0,
		"NetworkDeviceStatusDisconnected":  1,
		"NetworkDeviceStatusConnecting":    2,
		"NetworkDeviceStatusConnected":     3,
		"NetworkDeviceStatusDisconnecting": 4,
		"NetworkDeviceStatusDisabled":      5,
	}
)

func (x NetworkDeviceStatus) Enum() *NetworkDeviceStatus {
	p := new(NetworkDeviceStatus)
	*p = x
	return p
}

func (x NetworkDeviceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkDeviceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_sys_network_manager_proto_enumTypes[0].Descriptor()
}

func (NetworkDeviceStatus) Type() protoreflect.EnumType {
	return &file_sys_network_manager_proto_enumTypes[0]
}

func (x NetworkDeviceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkDeviceStatus.Descriptor instead.
func (NetworkDeviceStatus) EnumDescriptor() ([]byte, []int) {
	return file_sys_network_manager_proto_rawDescGZIP(), []int{0}
}

type KeyMgmt int32

const (
	KeyMgmt_KeyMgmtNone    KeyMgmt = 0
	KeyMgmt_KeyMgmtWEP     KeyMgmt = 1 // WEP  （最老的协议了，目前几乎没人用）
	KeyMgmt_KeyMgmtWPA_PSK KeyMgmt = 2 // WPA/WPA2-Personal  （一般大概率都是这个，应当作为默认值）
	KeyMgmt_KeyMgmtSAE     KeyMgmt = 3 // WPA3-Personal  （新一代协议，用的人比较少）
)

// Enum value maps for KeyMgmt.
var (
	KeyMgmt_name = map[int32]string{
		0: "KeyMgmtNone",
		1: "KeyMgmtWEP",
		2: "KeyMgmtWPA_PSK",
		3: "KeyMgmtSAE",
	}
	KeyMgmt_value = map[string]int32{
		"KeyMgmtNone":    0,
		"KeyMgmtWEP":     1,
		"KeyMgmtWPA_PSK": 2,
		"KeyMgmtSAE":     3,
	}
)

func (x KeyMgmt) Enum() *KeyMgmt {
	p := new(KeyMgmt)
	*p = x
	return p
}

func (x KeyMgmt) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyMgmt) Descriptor() protoreflect.EnumDescriptor {
	return file_sys_network_manager_proto_enumTypes[1].Descriptor()
}

func (KeyMgmt) Type() protoreflect.EnumType {
	return &file_sys_network_manager_proto_enumTypes[1]
}

func (x KeyMgmt) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyMgmt.Descriptor instead.
func (KeyMgmt) EnumDescriptor() ([]byte, []int) {
	return file_sys_network_manager_proto_rawDescGZIP(), []int{1}
}

type WifiConnectResult int32

const (
	WifiConnectResult_WifiConnectResultSuccess       WifiConnectResult = 0
	WifiConnectResult_WifiConnectResultNoSuchBssid   WifiConnectResult = 1
	WifiConnectResult_WifiConnectResultWrongPassword WifiConnectResult = 2
	WifiConnectResult_WifiConnectResultUnknownError  WifiConnectResult = 3
)

// Enum value maps for WifiConnectResult.
var (
	WifiConnectResult_name = map[int32]string{
		0: "WifiConnectResultSuccess",
		1: "WifiConnectResultNoSuchBssid",
		2: "WifiConnectResultWrongPassword",
		3: "WifiConnectResultUnknownError",
	}
	WifiConnectResult_value = map[string]int32{
		"WifiConnectResultSuccess":       0,
		"WifiConnectResultNoSuchBssid":   1,
		"WifiConnectResultWrongPassword": 2,
		"WifiConnectResultUnknownError":  3,
	}
)

func (x WifiConnectResult) Enum() *WifiConnectResult {
	p := new(WifiConnectResult)
	*p = x
	return p
}

func (x WifiConnectResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WifiConnectResult) Descriptor() protoreflect.EnumDescriptor {
	return file_sys_network_manager_proto_enumTypes[2].Descriptor()
}

func (WifiConnectResult) Type() protoreflect.EnumType {
	return &file_sys_network_manager_proto_enumTypes[2]
}

func (x WifiConnectResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WifiConnectResult.Descriptor instead.
func (WifiConnectResult) EnumDescriptor() ([]byte, []int) {
	return file_sys_network_manager_proto_rawDescGZIP(), []int{2}
}

type AccessPointInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 热点的网卡 mac 地址（由于 ssid 可能重复，所以将此字段作为整个列表的 index）
	Bssid string `protobuf:"bytes,1,opt,name=bssid,proto3" json:"bssid,omitempty"`
	// 热点的 ssid
	Ssid string `protobuf:"bytes,2,opt,name=ssid,proto3" json:"ssid,omitempty"`
	// 信号强度（范围 0 <= signal <= 100）
	Signal int32 `protobuf:"varint,3,opt,name=signal,proto3" json:"signal,omitempty"`
	// 是否需要密码
	Security bool `protobuf:"varint,4,opt,name=security,proto3" json:"security,omitempty"`
	// 是否已连接
	Connected bool `protobuf:"varint,5,opt,name=connected,proto3" json:"connected,omitempty"`
	// 是否已保存密码
	Saved bool `protobuf:"varint,6,opt,name=saved,proto3" json:"saved,omitempty"`
	// 是否自动连接
	AutoConnected bool `protobuf:"varint,7,opt,name=auto_connected,json=autoConnected,proto3" json:"auto_connected,omitempty"`
}

func (x *AccessPointInfo) Reset() {
	*x = AccessPointInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_network_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessPointInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessPointInfo) ProtoMessage() {}

func (x *AccessPointInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sys_network_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessPointInfo.ProtoReflect.Descriptor instead.
func (*AccessPointInfo) Descriptor() ([]byte, []int) {
	return file_sys_network_manager_proto_rawDescGZIP(), []int{0}
}

func (x *AccessPointInfo) GetBssid() string {
	if x != nil {
		return x.Bssid
	}
	return ""
}

func (x *AccessPointInfo) GetSsid() string {
	if x != nil {
		return x.Ssid
	}
	return ""
}

func (x *AccessPointInfo) GetSignal() int32 {
	if x != nil {
		return x.Signal
	}
	return 0
}

func (x *AccessPointInfo) GetSecurity() bool {
	if x != nil {
		return x.Security
	}
	return false
}

func (x *AccessPointInfo) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

func (x *AccessPointInfo) GetSaved() bool {
	if x != nil {
		return x.Saved
	}
	return false
}

func (x *AccessPointInfo) GetAutoConnected() bool {
	if x != nil {
		return x.AutoConnected
	}
	return false
}

type AccessPointInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*AccessPointInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AccessPointInfoList) Reset() {
	*x = AccessPointInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_network_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessPointInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessPointInfoList) ProtoMessage() {}

func (x *AccessPointInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_sys_network_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessPointInfoList.ProtoReflect.Descriptor instead.
func (*AccessPointInfoList) Descriptor() ([]byte, []int) {
	return file_sys_network_manager_proto_rawDescGZIP(), []int{1}
}

func (x *AccessPointInfoList) GetList() []*AccessPointInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type NetworkDeviceStatusInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否已连接到互联网
	HasInternet bool `protobuf:"varint,1,opt,name=hasInternet,proto3" json:"hasInternet,omitempty"`
	// 有线连接状态（已假设只有一块有线网卡）
	WiredDevice NetworkDeviceStatus `protobuf:"varint,2,opt,name=WiredDevice,proto3,enum=cloud.lazycat.apis.sys.NetworkDeviceStatus" json:"WiredDevice,omitempty"`
	// 无线设备状态（已假设只有一块无线网卡）
	WirelessDevice NetworkDeviceStatus `protobuf:"varint,3,opt,name=WirelessDevice,proto3,enum=cloud.lazycat.apis.sys.NetworkDeviceStatus" json:"WirelessDevice,omitempty"`
	// 若无线设备已连接，则该字段表示已连接的 wifi 的信息
	Info *AccessPointInfo `protobuf:"bytes,4,opt,name=info,proto3,oneof" json:"info,omitempty"`
}

func (x *NetworkDeviceStatusInfo) Reset() {
	*x = NetworkDeviceStatusInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_network_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkDeviceStatusInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkDeviceStatusInfo) ProtoMessage() {}

func (x *NetworkDeviceStatusInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sys_network_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkDeviceStatusInfo.ProtoReflect.Descriptor instead.
func (*NetworkDeviceStatusInfo) Descriptor() ([]byte, []int) {
	return file_sys_network_manager_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkDeviceStatusInfo) GetHasInternet() bool {
	if x != nil {
		return x.HasInternet
	}
	return false
}

func (x *NetworkDeviceStatusInfo) GetWiredDevice() NetworkDeviceStatus {
	if x != nil {
		return x.WiredDevice
	}
	return NetworkDeviceStatus_NetworkDeviceStatusUnavailable
}

func (x *NetworkDeviceStatusInfo) GetWirelessDevice() NetworkDeviceStatus {
	if x != nil {
		return x.WirelessDevice
	}
	return NetworkDeviceStatus_NetworkDeviceStatusUnavailable
}

func (x *NetworkDeviceStatusInfo) GetInfo() *AccessPointInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type WifiConnectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// bssid 和 ssid 指定其一即可
	Bssid    string   `protobuf:"bytes,1,opt,name=bssid,proto3" json:"bssid,omitempty"`
	Ssid     string   `protobuf:"bytes,2,opt,name=ssid,proto3" json:"ssid,omitempty"`
	Password string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	KeyMgmt  *KeyMgmt `protobuf:"varint,4,opt,name=key_mgmt,json=keyMgmt,proto3,enum=cloud.lazycat.apis.sys.KeyMgmt,oneof" json:"key_mgmt,omitempty"`
}

func (x *WifiConnectInfo) Reset() {
	*x = WifiConnectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_network_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiConnectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiConnectInfo) ProtoMessage() {}

func (x *WifiConnectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sys_network_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiConnectInfo.ProtoReflect.Descriptor instead.
func (*WifiConnectInfo) Descriptor() ([]byte, []int) {
	return file_sys_network_manager_proto_rawDescGZIP(), []int{3}
}

func (x *WifiConnectInfo) GetBssid() string {
	if x != nil {
		return x.Bssid
	}
	return ""
}

func (x *WifiConnectInfo) GetSsid() string {
	if x != nil {
		return x.Ssid
	}
	return ""
}

func (x *WifiConnectInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *WifiConnectInfo) GetKeyMgmt() KeyMgmt {
	if x != nil && x.KeyMgmt != nil {
		return *x.KeyMgmt
	}
	return KeyMgmt_KeyMgmtNone
}

type WifiConfigInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ssid     string  `protobuf:"bytes,1,opt,name=ssid,proto3" json:"ssid,omitempty"`
	Password string  `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	KeyMgmt  KeyMgmt `protobuf:"varint,3,opt,name=key_mgmt,json=keyMgmt,proto3,enum=cloud.lazycat.apis.sys.KeyMgmt" json:"key_mgmt,omitempty"`
}

func (x *WifiConfigInfo) Reset() {
	*x = WifiConfigInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_network_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiConfigInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiConfigInfo) ProtoMessage() {}

func (x *WifiConfigInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sys_network_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiConfigInfo.ProtoReflect.Descriptor instead.
func (*WifiConfigInfo) Descriptor() ([]byte, []int) {
	return file_sys_network_manager_proto_rawDescGZIP(), []int{4}
}

func (x *WifiConfigInfo) GetSsid() string {
	if x != nil {
		return x.Ssid
	}
	return ""
}

func (x *WifiConfigInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *WifiConfigInfo) GetKeyMgmt() KeyMgmt {
	if x != nil {
		return x.KeyMgmt
	}
	return KeyMgmt_KeyMgmtNone
}

type WifiConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result WifiConnectResult `protobuf:"varint,1,opt,name=result,proto3,enum=cloud.lazycat.apis.sys.WifiConnectResult" json:"result,omitempty"`
}

func (x *WifiConnectReply) Reset() {
	*x = WifiConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_network_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiConnectReply) ProtoMessage() {}

func (x *WifiConnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_sys_network_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiConnectReply.ProtoReflect.Descriptor instead.
func (*WifiConnectReply) Descriptor() ([]byte, []int) {
	return file_sys_network_manager_proto_rawDescGZIP(), []int{5}
}

func (x *WifiConnectReply) GetResult() WifiConnectResult {
	if x != nil {
		return x.Result
	}
	return WifiConnectResult_WifiConnectResultSuccess
}

var File_sys_network_manager_proto protoreflect.FileDescriptor

var file_sys_network_manager_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x79, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xca, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x73, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x73, 0x61, 0x76, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x61, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x52, 0x0a,
	0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0xaa, 0x02, 0x0a, 0x17, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x68, 0x61, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12,
	0x4d, 0x0a, 0x0b, 0x57, 0x69, 0x72, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x0b, 0x57, 0x69, 0x72, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53,
	0x0a, 0x0e, 0x57, 0x69, 0x72, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0e, 0x57, 0x69, 0x72, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xa5,
	0x01, 0x0a, 0x0f, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x73, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3f, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f,
	0x6d, 0x67, 0x6d, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x4d, 0x67, 0x6d, 0x74, 0x48, 0x00, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x4d, 0x67, 0x6d, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x6d, 0x67, 0x6d, 0x74, 0x22, 0x7c, 0x0a, 0x0e, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x73, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x73, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f,
	0x6d, 0x67, 0x6d, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x4d, 0x67, 0x6d, 0x74, 0x52, 0x07, 0x6b, 0x65, 0x79,
	0x4d, 0x67, 0x6d, 0x74, 0x22, 0x55, 0x0a, 0x10, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79,
	0x73, 0x2e, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0xea, 0x01, 0x0a, 0x13,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x1e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12,
	0x20, 0x0a, 0x1c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10,
	0x03, 0x12, 0x24, 0x0a, 0x20, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0x05, 0x2a, 0x4e, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x4d,
	0x67, 0x6d, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x4d, 0x67, 0x6d, 0x74, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x4d, 0x67, 0x6d, 0x74, 0x57,
	0x45, 0x50, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x4d, 0x67, 0x6d, 0x74, 0x57,
	0x50, 0x41, 0x5f, 0x50, 0x53, 0x4b, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x4d,
	0x67, 0x6d, 0x74, 0x53, 0x41, 0x45, 0x10, 0x03, 0x2a, 0x9a, 0x01, 0x0a, 0x11, 0x57, 0x69, 0x66,
	0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c,
	0x0a, 0x18, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c,
	0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x4e, 0x6f, 0x53, 0x75, 0x63, 0x68, 0x42, 0x73, 0x73, 0x69, 0x64, 0x10, 0x01, 0x12, 0x22,
	0x0a, 0x1e, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x57, 0x72, 0x6f, 0x6e, 0x67, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x10, 0x03, 0x32, 0xbf, 0x03, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x08, 0x57, 0x69, 0x66, 0x69, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x08, 0x57,
	0x69, 0x66, 0x69, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x62,
	0x0a, 0x0b, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x27, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e,
	0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x63, 0x0a, 0x0d, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x41, 0x64, 0x64, 0x12, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x57, 0x69, 0x66,
	0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x28, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x73, 0x79, 0x73, 0x2e, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x65, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a,
	0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x79,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_network_manager_proto_rawDescOnce sync.Once
	file_sys_network_manager_proto_rawDescData = file_sys_network_manager_proto_rawDesc
)

func file_sys_network_manager_proto_rawDescGZIP() []byte {
	file_sys_network_manager_proto_rawDescOnce.Do(func() {
		file_sys_network_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_network_manager_proto_rawDescData)
	})
	return file_sys_network_manager_proto_rawDescData
}

var file_sys_network_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_sys_network_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sys_network_manager_proto_goTypes = []interface{}{
	(NetworkDeviceStatus)(0),        // 0: cloud.lazycat.apis.sys.NetworkDeviceStatus
	(KeyMgmt)(0),                    // 1: cloud.lazycat.apis.sys.KeyMgmt
	(WifiConnectResult)(0),          // 2: cloud.lazycat.apis.sys.WifiConnectResult
	(*AccessPointInfo)(nil),         // 3: cloud.lazycat.apis.sys.AccessPointInfo
	(*AccessPointInfoList)(nil),     // 4: cloud.lazycat.apis.sys.AccessPointInfoList
	(*NetworkDeviceStatusInfo)(nil), // 5: cloud.lazycat.apis.sys.NetworkDeviceStatusInfo
	(*WifiConnectInfo)(nil),         // 6: cloud.lazycat.apis.sys.WifiConnectInfo
	(*WifiConfigInfo)(nil),          // 7: cloud.lazycat.apis.sys.WifiConfigInfo
	(*WifiConnectReply)(nil),        // 8: cloud.lazycat.apis.sys.WifiConnectReply
	(*emptypb.Empty)(nil),           // 9: google.protobuf.Empty
}
var file_sys_network_manager_proto_depIdxs = []int32{
	3,  // 0: cloud.lazycat.apis.sys.AccessPointInfoList.list:type_name -> cloud.lazycat.apis.sys.AccessPointInfo
	0,  // 1: cloud.lazycat.apis.sys.NetworkDeviceStatusInfo.WiredDevice:type_name -> cloud.lazycat.apis.sys.NetworkDeviceStatus
	0,  // 2: cloud.lazycat.apis.sys.NetworkDeviceStatusInfo.WirelessDevice:type_name -> cloud.lazycat.apis.sys.NetworkDeviceStatus
	3,  // 3: cloud.lazycat.apis.sys.NetworkDeviceStatusInfo.info:type_name -> cloud.lazycat.apis.sys.AccessPointInfo
	1,  // 4: cloud.lazycat.apis.sys.WifiConnectInfo.key_mgmt:type_name -> cloud.lazycat.apis.sys.KeyMgmt
	1,  // 5: cloud.lazycat.apis.sys.WifiConfigInfo.key_mgmt:type_name -> cloud.lazycat.apis.sys.KeyMgmt
	2,  // 6: cloud.lazycat.apis.sys.WifiConnectReply.result:type_name -> cloud.lazycat.apis.sys.WifiConnectResult
	9,  // 7: cloud.lazycat.apis.sys.NetworkManager.Status:input_type -> google.protobuf.Empty
	9,  // 8: cloud.lazycat.apis.sys.NetworkManager.WifiScan:input_type -> google.protobuf.Empty
	9,  // 9: cloud.lazycat.apis.sys.NetworkManager.WifiList:input_type -> google.protobuf.Empty
	6,  // 10: cloud.lazycat.apis.sys.NetworkManager.WifiConnect:input_type -> cloud.lazycat.apis.sys.WifiConnectInfo
	7,  // 11: cloud.lazycat.apis.sys.NetworkManager.WifiConfigAdd:input_type -> cloud.lazycat.apis.sys.WifiConfigInfo
	5,  // 12: cloud.lazycat.apis.sys.NetworkManager.Status:output_type -> cloud.lazycat.apis.sys.NetworkDeviceStatusInfo
	9,  // 13: cloud.lazycat.apis.sys.NetworkManager.WifiScan:output_type -> google.protobuf.Empty
	4,  // 14: cloud.lazycat.apis.sys.NetworkManager.WifiList:output_type -> cloud.lazycat.apis.sys.AccessPointInfoList
	8,  // 15: cloud.lazycat.apis.sys.NetworkManager.WifiConnect:output_type -> cloud.lazycat.apis.sys.WifiConnectReply
	8,  // 16: cloud.lazycat.apis.sys.NetworkManager.WifiConfigAdd:output_type -> cloud.lazycat.apis.sys.WifiConnectReply
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_sys_network_manager_proto_init() }
func file_sys_network_manager_proto_init() {
	if File_sys_network_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_network_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessPointInfo); i {
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
		file_sys_network_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessPointInfoList); i {
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
		file_sys_network_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkDeviceStatusInfo); i {
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
		file_sys_network_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiConnectInfo); i {
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
		file_sys_network_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiConfigInfo); i {
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
		file_sys_network_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiConnectReply); i {
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
	file_sys_network_manager_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_sys_network_manager_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sys_network_manager_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_network_manager_proto_goTypes,
		DependencyIndexes: file_sys_network_manager_proto_depIdxs,
		EnumInfos:         file_sys_network_manager_proto_enumTypes,
		MessageInfos:      file_sys_network_manager_proto_msgTypes,
	}.Build()
	File_sys_network_manager_proto = out.File
	file_sys_network_manager_proto_rawDesc = nil
	file_sys_network_manager_proto_goTypes = nil
	file_sys_network_manager_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: sys/package_manager.proto

package sys

import (
	_ "gitee.com/linakesi/lzc-sdk/lang/go/common"
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

type PkgURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 内部地址
	//     http://pkgm.api-server.lzcapp/tmp/xxxx-0.2.1.lpk
	// 或外网地址
	//     https://repo.lazycat.cloud/a/c/accc-0.2.1.lpk
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// pkg对应的sha256值，若不为空，
	// 1. 本地有对应包的缓存，则会直接使用缓存
	// 2. 本地若没有缓存，则会从 url 下载包，并校验包的 sha256 值
	Sha256 string `protobuf:"bytes,2,opt,name=sha256,proto3" json:"sha256,omitempty"`
}

func (x *PkgURL) Reset() {
	*x = PkgURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PkgURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PkgURL) ProtoMessage() {}

func (x *PkgURL) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PkgURL.ProtoReflect.Descriptor instead.
func (*PkgURL) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{0}
}

func (x *PkgURL) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PkgURL) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

type UninstallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	// 卸载后是否清空 data 目录 (/lzcapp/var)
	ClearData bool `protobuf:"varint,2,opt,name=clear_data,json=clearData,proto3" json:"clear_data,omitempty"`
}

func (x *UninstallRequest) Reset() {
	*x = UninstallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UninstallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UninstallRequest) ProtoMessage() {}

func (x *UninstallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UninstallRequest.ProtoReflect.Descriptor instead.
func (*UninstallRequest) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{1}
}

func (x *UninstallRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *UninstallRequest) GetClearData() bool {
	if x != nil {
		return x.ClearData
	}
	return false
}

type AppDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid   string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// 应用图标所在的 url，如 https://apis.$boxdomain/icon/$appid.png
	Icon string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	// 应用所在的域名，如 app.box.heiyu.space
	Domain string `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *AppDescription) Reset() {
	*x = AppDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppDescription) ProtoMessage() {}

func (x *AppDescription) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppDescription.ProtoReflect.Descriptor instead.
func (*AppDescription) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{2}
}

func (x *AppDescription) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *AppDescription) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AppDescription) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AppDescription) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *AppDescription) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type QueryApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 需要查询的 AppId 的列表，如果列表为空，则查询所有的应用
	AppidList []string `protobuf:"bytes,1,rep,name=appid_list,json=appidList,proto3" json:"appid_list,omitempty"`
}

func (x *QueryApplicationRequest) Reset() {
	*x = QueryApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryApplicationRequest) ProtoMessage() {}

func (x *QueryApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryApplicationRequest.ProtoReflect.Descriptor instead.
func (*QueryApplicationRequest) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{3}
}

func (x *QueryApplicationRequest) GetAppidList() []string {
	if x != nil {
		return x.AppidList
	}
	return nil
}

type QueryApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DescList []*AppDescription `protobuf:"bytes,1,rep,name=desc_list,json=descList,proto3" json:"desc_list,omitempty"`
}

func (x *QueryApplicationResponse) Reset() {
	*x = QueryApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryApplicationResponse) ProtoMessage() {}

func (x *QueryApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryApplicationResponse.ProtoReflect.Descriptor instead.
func (*QueryApplicationResponse) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{4}
}

func (x *QueryApplicationResponse) GetDescList() []*AppDescription {
	if x != nil {
		return x.DescList
	}
	return nil
}

type QueryAppStorageUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid        string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	NeedPkg      bool   `protobuf:"varint,2,opt,name=need_pkg,json=needPkg,proto3" json:"need_pkg,omitempty"`
	NeedData     bool   `protobuf:"varint,3,opt,name=need_data,json=needData,proto3" json:"need_data,omitempty"`
	NeedCache    bool   `protobuf:"varint,4,opt,name=need_cache,json=needCache,proto3" json:"need_cache,omitempty"`
	NeedTmp      bool   `protobuf:"varint,5,opt,name=need_tmp,json=needTmp,proto3" json:"need_tmp,omitempty"`
	NeedUserdata bool   `protobuf:"varint,6,opt,name=need_userdata,json=needUserdata,proto3" json:"need_userdata,omitempty"`
}

func (x *QueryAppStorageUsageRequest) Reset() {
	*x = QueryAppStorageUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAppStorageUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAppStorageUsageRequest) ProtoMessage() {}

func (x *QueryAppStorageUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAppStorageUsageRequest.ProtoReflect.Descriptor instead.
func (*QueryAppStorageUsageRequest) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{5}
}

func (x *QueryAppStorageUsageRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *QueryAppStorageUsageRequest) GetNeedPkg() bool {
	if x != nil {
		return x.NeedPkg
	}
	return false
}

func (x *QueryAppStorageUsageRequest) GetNeedData() bool {
	if x != nil {
		return x.NeedData
	}
	return false
}

func (x *QueryAppStorageUsageRequest) GetNeedCache() bool {
	if x != nil {
		return x.NeedCache
	}
	return false
}

func (x *QueryAppStorageUsageRequest) GetNeedTmp() bool {
	if x != nil {
		return x.NeedTmp
	}
	return false
}

func (x *QueryAppStorageUsageRequest) GetNeedUserdata() bool {
	if x != nil {
		return x.NeedUserdata
	}
	return false
}

type AppStorageUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 应用本身占用的大小 (/lzcapp/pkg)
	Pkg uint64 `protobuf:"varint,1,opt,name=pkg,proto3" json:"pkg,omitempty"`
	// 应用数据的大小 (所有用户产生的应用数据大小) (/lzcapp/var)
	Data uint64 `protobuf:"varint,2,opt,name=data,proto3" json:"data,omitempty"`
	// 应用缓存的大小 (/lzcapp/cache)
	Cache uint64 `protobuf:"varint,3,opt,name=cache,proto3" json:"cache,omitempty"`
	// 应用临时文件的大小 (/tmp)
	Tmp uint64 `protobuf:"varint,4,opt,name=tmp,proto3" json:"tmp,omitempty"`
	// 各个用户产生的应用数据大小 (/lzcapp/var/userdata/$uid)
	Userdata map[string]uint64 `protobuf:"bytes,5,rep,name=userdata,proto3" json:"userdata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *AppStorageUsage) Reset() {
	*x = AppStorageUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppStorageUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppStorageUsage) ProtoMessage() {}

func (x *AppStorageUsage) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppStorageUsage.ProtoReflect.Descriptor instead.
func (*AppStorageUsage) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{6}
}

func (x *AppStorageUsage) GetPkg() uint64 {
	if x != nil {
		return x.Pkg
	}
	return 0
}

func (x *AppStorageUsage) GetData() uint64 {
	if x != nil {
		return x.Data
	}
	return 0
}

func (x *AppStorageUsage) GetCache() uint64 {
	if x != nil {
		return x.Cache
	}
	return 0
}

func (x *AppStorageUsage) GetTmp() uint64 {
	if x != nil {
		return x.Tmp
	}
	return 0
}

func (x *AppStorageUsage) GetUserdata() map[string]uint64 {
	if x != nil {
		return x.Userdata
	}
	return nil
}

type SetUserPermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户的 uid
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// 是否允许 uid 安装应用
	AllowInstallApp bool `protobuf:"varint,2,opt,name=allow_install_app,json=allowInstallApp,proto3" json:"allow_install_app,omitempty"`
	// 是否允许 uid 卸载应用
	AllowUninstallApp bool `protobuf:"varint,3,opt,name=allow_uninstall_app,json=allowUninstallApp,proto3" json:"allow_uninstall_app,omitempty"`
}

func (x *SetUserPermissionsRequest) Reset() {
	*x = SetUserPermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sys_package_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserPermissionsRequest) ProtoMessage() {}

func (x *SetUserPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sys_package_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserPermissionsRequest.ProtoReflect.Descriptor instead.
func (*SetUserPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_sys_package_manager_proto_rawDescGZIP(), []int{7}
}

func (x *SetUserPermissionsRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *SetUserPermissionsRequest) GetAllowInstallApp() bool {
	if x != nil {
		return x.AllowInstallApp
	}
	return false
}

func (x *SetUserPermissionsRequest) GetAllowUninstallApp() bool {
	if x != nil {
		return x.AllowUninstallApp
	}
	return false
}

var File_sys_package_manager_proto protoreflect.FileDescriptor

var file_sys_package_manager_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x79, 0x73, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x32, 0x0a, 0x06, 0x50, 0x6b, 0x67, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61,
	0x32, 0x35, 0x36, 0x22, 0x47, 0x0a, 0x10, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x82, 0x01, 0x0a,
	0x0e, 0x41, 0x70, 0x70, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x38, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x70, 0x70, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x70, 0x70, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x18, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x64, 0x65, 0x73, 0x63, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x73, 0x79, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xca, 0x01, 0x0a,
	0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x70, 0x6b, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x65, 0x65, 0x64, 0x50, 0x6b, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x6e, 0x65, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65,
	0x65, 0x64, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x6e, 0x65, 0x65, 0x64, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x65,
	0x64, 0x5f, 0x74, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x65, 0x65,
	0x64, 0x54, 0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e, 0x65, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x22, 0xef, 0x01, 0x0a, 0x0f, 0x41, 0x70,
	0x70, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x6b, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x6b, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6d, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74, 0x6d, 0x70, 0x12, 0x51, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b,
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x89, 0x01, 0x0a, 0x19,
	0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x70, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x75, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x70, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x6e, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x32, 0x8c, 0x04, 0x0a, 0x0e, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x07, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x1e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61,
	0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x50,
	0x6b, 0x67, 0x55, 0x52, 0x4c, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0x82,
	0x97, 0x22, 0x02, 0x08, 0x01, 0x12, 0x55, 0x0a, 0x09, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x06, 0x82, 0x97, 0x22, 0x02, 0x08, 0x01, 0x12, 0x77, 0x0a, 0x10,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x70,
	0x70, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63,
	0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a,
	0x12, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79,
	0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x53, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x06,
	0x82, 0x97, 0x22, 0x02, 0x10, 0x02, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61, 0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x79, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sys_package_manager_proto_rawDescOnce sync.Once
	file_sys_package_manager_proto_rawDescData = file_sys_package_manager_proto_rawDesc
)

func file_sys_package_manager_proto_rawDescGZIP() []byte {
	file_sys_package_manager_proto_rawDescOnce.Do(func() {
		file_sys_package_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_sys_package_manager_proto_rawDescData)
	})
	return file_sys_package_manager_proto_rawDescData
}

var file_sys_package_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_sys_package_manager_proto_goTypes = []interface{}{
	(*PkgURL)(nil),                      // 0: cloud.lazycat.apis.sys.PkgURL
	(*UninstallRequest)(nil),            // 1: cloud.lazycat.apis.sys.UninstallRequest
	(*AppDescription)(nil),              // 2: cloud.lazycat.apis.sys.AppDescription
	(*QueryApplicationRequest)(nil),     // 3: cloud.lazycat.apis.sys.QueryApplicationRequest
	(*QueryApplicationResponse)(nil),    // 4: cloud.lazycat.apis.sys.QueryApplicationResponse
	(*QueryAppStorageUsageRequest)(nil), // 5: cloud.lazycat.apis.sys.QueryAppStorageUsageRequest
	(*AppStorageUsage)(nil),             // 6: cloud.lazycat.apis.sys.AppStorageUsage
	(*SetUserPermissionsRequest)(nil),   // 7: cloud.lazycat.apis.sys.SetUserPermissionsRequest
	nil,                                 // 8: cloud.lazycat.apis.sys.AppStorageUsage.UserdataEntry
	(*emptypb.Empty)(nil),               // 9: google.protobuf.Empty
}
var file_sys_package_manager_proto_depIdxs = []int32{
	2, // 0: cloud.lazycat.apis.sys.QueryApplicationResponse.desc_list:type_name -> cloud.lazycat.apis.sys.AppDescription
	8, // 1: cloud.lazycat.apis.sys.AppStorageUsage.userdata:type_name -> cloud.lazycat.apis.sys.AppStorageUsage.UserdataEntry
	0, // 2: cloud.lazycat.apis.sys.PackageManager.Install:input_type -> cloud.lazycat.apis.sys.PkgURL
	1, // 3: cloud.lazycat.apis.sys.PackageManager.Uninstall:input_type -> cloud.lazycat.apis.sys.UninstallRequest
	3, // 4: cloud.lazycat.apis.sys.PackageManager.QueryApplication:input_type -> cloud.lazycat.apis.sys.QueryApplicationRequest
	5, // 5: cloud.lazycat.apis.sys.PackageManager.QueryAppStorageUsage:input_type -> cloud.lazycat.apis.sys.QueryAppStorageUsageRequest
	7, // 6: cloud.lazycat.apis.sys.PackageManager.SetUserPermissions:input_type -> cloud.lazycat.apis.sys.SetUserPermissionsRequest
	9, // 7: cloud.lazycat.apis.sys.PackageManager.Install:output_type -> google.protobuf.Empty
	9, // 8: cloud.lazycat.apis.sys.PackageManager.Uninstall:output_type -> google.protobuf.Empty
	4, // 9: cloud.lazycat.apis.sys.PackageManager.QueryApplication:output_type -> cloud.lazycat.apis.sys.QueryApplicationResponse
	6, // 10: cloud.lazycat.apis.sys.PackageManager.QueryAppStorageUsage:output_type -> cloud.lazycat.apis.sys.AppStorageUsage
	9, // 11: cloud.lazycat.apis.sys.PackageManager.SetUserPermissions:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sys_package_manager_proto_init() }
func file_sys_package_manager_proto_init() {
	if File_sys_package_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sys_package_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PkgURL); i {
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
		file_sys_package_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UninstallRequest); i {
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
		file_sys_package_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppDescription); i {
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
		file_sys_package_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryApplicationRequest); i {
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
		file_sys_package_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryApplicationResponse); i {
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
		file_sys_package_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAppStorageUsageRequest); i {
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
		file_sys_package_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppStorageUsage); i {
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
		file_sys_package_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserPermissionsRequest); i {
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
			RawDescriptor: file_sys_package_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sys_package_manager_proto_goTypes,
		DependencyIndexes: file_sys_package_manager_proto_depIdxs,
		MessageInfos:      file_sys_package_manager_proto_msgTypes,
	}.Build()
	File_sys_package_manager_proto = out.File
	file_sys_package_manager_proto_rawDesc = nil
	file_sys_package_manager_proto_goTypes = nil
	file_sys_package_manager_proto_depIdxs = nil
}

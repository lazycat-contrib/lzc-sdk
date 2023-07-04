// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: sys/portal-server/portal-server.proto

package portal_server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HPortalSysClient is the client API for HPortalSys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HPortalSysClient interface {
	// 用auth-token反向查询登陆的设备以及UID
	QueryLogin(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*LoginInfo, error)
	QueryBoxInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BoxInfo, error)
	// 获取盒子所属域名下或下一级域名的https证书。
	// 注意不是所有ACME服务器都支持泛域名。
	GetDomainCert(ctx context.Context, in *DomainCertRequest, opts ...grpc.CallOption) (*DomainCertReply, error)
	// 申请额外的外部可访问的IP,并配置对应访问域名
	AllocVirtualExternalIP(ctx context.Context, in *AllocVEIPRequest, opts ...grpc.CallOption) (*AllocVEIPReply, error)
	// 释放虚拟IP
	FreeVirtualExternalIP(ctx context.Context, in *FreeVEIPRequest, opts ...grpc.CallOption) (*FreeVEIPReply, error)
	//  查询所有UID
	ListUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUsersReply, error)
	//  创建用户信息
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//  删除用户信息
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//  修改新的密码
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//  校验用户密码是否正确
	CheckPassword(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//  根据用户uid查询用户信息
	QueryRole(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*QueryRoleReply, error)
	//  修改指定uid的用户角色
	ChangeRole(ctx context.Context, in *ChangeRoleReqeust, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//  添加或删除受信任设备
	ChangeTrustEndDevice(ctx context.Context, in *ChangeTrustEndDeviceRequest, opts ...grpc.CallOption) (*ChangeTrustEndDeviceReply, error)
	// 根据UID返回所有的设备列表
	ListDevices(ctx context.Context, in *ListDeviceRequest, opts ...grpc.CallOption) (*ListDeviceReply, error)
	QueryDeviceByID(ctx context.Context, in *DeviceID, opts ...grpc.CallOption) (*Device, error)
	// 删除登陆的会话状态
	ClearLoginSession(ctx context.Context, in *ClearLoginSessionRequest, opts ...grpc.CallOption) (*ClearLoginSessionReply, error)
	// 获取remotesocks服务器地址
	RemoteSocks(ctx context.Context, in *RemoteSocksRequest, opts ...grpc.CallOption) (*RemoteSocksReply, error)
	// hserver重启后默认设置BoxSystem为booting状态
	// 实际的BoxSystem需要定期(建议两到三秒)设置其实际状态，避免hserver被手动或自动重启后设置的盒子系统状态错误
	UpdateBoxSystemStatus(ctx context.Context, in *UpdateBoxSystemStatusRequest, opts ...grpc.CallOption) (*UpdateBoxSystemStatusResponse, error)
	// 仅在盒子未初始化时，可以被调用。
	SetupHServer(ctx context.Context, in *SetupHServerRequest, opts ...grpc.CallOption) (*SetupHServerReply, error)
	// 重置盒子
	// 1. 向Origin请求释放盒子名下的所有域名
	// 2. 清除本地的box.name
	// 3. 进入为初始化状态
	ResetHServer(ctx context.Context, in *ResetHServerRequest, opts ...grpc.CallOption) (*ResetHServerReply, error)
	// Deprecated: Do not use.
	// ----------------------------- 以下为准备废弃的接口 --------------------------------------
	GetDomainSelfCert(ctx context.Context, in *DomainCertRequest, opts ...grpc.CallOption) (*DomainCertReply, error)
	// Deprecated: Do not use.
	// 以下接口要改名字
	// 强制将特定设备加到受信任列表中
	TrustUserDevice(ctx context.Context, in *TrustUserDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type hPortalSysClient struct {
	cc grpc.ClientConnInterface
}

func NewHPortalSysClient(cc grpc.ClientConnInterface) HPortalSysClient {
	return &hPortalSysClient{cc}
}

func (c *hPortalSysClient) QueryLogin(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*LoginInfo, error) {
	out := new(LoginInfo)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/QueryLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) QueryBoxInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BoxInfo, error) {
	out := new(BoxInfo)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/QueryBoxInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) GetDomainCert(ctx context.Context, in *DomainCertRequest, opts ...grpc.CallOption) (*DomainCertReply, error) {
	out := new(DomainCertReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/GetDomainCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) AllocVirtualExternalIP(ctx context.Context, in *AllocVEIPRequest, opts ...grpc.CallOption) (*AllocVEIPReply, error) {
	out := new(AllocVEIPReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/AllocVirtualExternalIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) FreeVirtualExternalIP(ctx context.Context, in *FreeVEIPRequest, opts ...grpc.CallOption) (*FreeVEIPReply, error) {
	out := new(FreeVEIPReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/FreeVirtualExternalIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) ListUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUsersReply, error) {
	out := new(ListUsersReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) CheckPassword(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/CheckPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) QueryRole(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*QueryRoleReply, error) {
	out := new(QueryRoleReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/QueryRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) ChangeRole(ctx context.Context, in *ChangeRoleReqeust, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/ChangeRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) ChangeTrustEndDevice(ctx context.Context, in *ChangeTrustEndDeviceRequest, opts ...grpc.CallOption) (*ChangeTrustEndDeviceReply, error) {
	out := new(ChangeTrustEndDeviceReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/ChangeTrustEndDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) ListDevices(ctx context.Context, in *ListDeviceRequest, opts ...grpc.CallOption) (*ListDeviceReply, error) {
	out := new(ListDeviceReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/ListDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) QueryDeviceByID(ctx context.Context, in *DeviceID, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/QueryDeviceByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) ClearLoginSession(ctx context.Context, in *ClearLoginSessionRequest, opts ...grpc.CallOption) (*ClearLoginSessionReply, error) {
	out := new(ClearLoginSessionReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/ClearLoginSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) RemoteSocks(ctx context.Context, in *RemoteSocksRequest, opts ...grpc.CallOption) (*RemoteSocksReply, error) {
	out := new(RemoteSocksReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/RemoteSocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) UpdateBoxSystemStatus(ctx context.Context, in *UpdateBoxSystemStatusRequest, opts ...grpc.CallOption) (*UpdateBoxSystemStatusResponse, error) {
	out := new(UpdateBoxSystemStatusResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/UpdateBoxSystemStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) SetupHServer(ctx context.Context, in *SetupHServerRequest, opts ...grpc.CallOption) (*SetupHServerReply, error) {
	out := new(SetupHServerReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/SetupHServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) ResetHServer(ctx context.Context, in *ResetHServerRequest, opts ...grpc.CallOption) (*ResetHServerReply, error) {
	out := new(ResetHServerReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/ResetHServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *hPortalSysClient) GetDomainSelfCert(ctx context.Context, in *DomainCertRequest, opts ...grpc.CallOption) (*DomainCertReply, error) {
	out := new(DomainCertReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/GetDomainSelfCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *hPortalSysClient) TrustUserDevice(ctx context.Context, in *TrustUserDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/TrustUserDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HPortalSysServer is the server API for HPortalSys service.
// All implementations must embed UnimplementedHPortalSysServer
// for forward compatibility
type HPortalSysServer interface {
	// 用auth-token反向查询登陆的设备以及UID
	QueryLogin(context.Context, *AuthToken) (*LoginInfo, error)
	QueryBoxInfo(context.Context, *emptypb.Empty) (*BoxInfo, error)
	// 获取盒子所属域名下或下一级域名的https证书。
	// 注意不是所有ACME服务器都支持泛域名。
	GetDomainCert(context.Context, *DomainCertRequest) (*DomainCertReply, error)
	// 申请额外的外部可访问的IP,并配置对应访问域名
	AllocVirtualExternalIP(context.Context, *AllocVEIPRequest) (*AllocVEIPReply, error)
	// 释放虚拟IP
	FreeVirtualExternalIP(context.Context, *FreeVEIPRequest) (*FreeVEIPReply, error)
	//  查询所有UID
	ListUsers(context.Context, *emptypb.Empty) (*ListUsersReply, error)
	//  创建用户信息
	CreateUser(context.Context, *CreateUserRequest) (*emptypb.Empty, error)
	//  删除用户信息
	DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	//  修改新的密码
	ResetPassword(context.Context, *ResetPasswordRequest) (*emptypb.Empty, error)
	//  校验用户密码是否正确
	CheckPassword(context.Context, *CheckPasswordRequest) (*emptypb.Empty, error)
	//  根据用户uid查询用户信息
	QueryRole(context.Context, *UserID) (*QueryRoleReply, error)
	//  修改指定uid的用户角色
	ChangeRole(context.Context, *ChangeRoleReqeust) (*emptypb.Empty, error)
	//  添加或删除受信任设备
	ChangeTrustEndDevice(context.Context, *ChangeTrustEndDeviceRequest) (*ChangeTrustEndDeviceReply, error)
	// 根据UID返回所有的设备列表
	ListDevices(context.Context, *ListDeviceRequest) (*ListDeviceReply, error)
	QueryDeviceByID(context.Context, *DeviceID) (*Device, error)
	// 删除登陆的会话状态
	ClearLoginSession(context.Context, *ClearLoginSessionRequest) (*ClearLoginSessionReply, error)
	// 获取remotesocks服务器地址
	RemoteSocks(context.Context, *RemoteSocksRequest) (*RemoteSocksReply, error)
	// hserver重启后默认设置BoxSystem为booting状态
	// 实际的BoxSystem需要定期(建议两到三秒)设置其实际状态，避免hserver被手动或自动重启后设置的盒子系统状态错误
	UpdateBoxSystemStatus(context.Context, *UpdateBoxSystemStatusRequest) (*UpdateBoxSystemStatusResponse, error)
	// 仅在盒子未初始化时，可以被调用。
	SetupHServer(context.Context, *SetupHServerRequest) (*SetupHServerReply, error)
	// 重置盒子
	// 1. 向Origin请求释放盒子名下的所有域名
	// 2. 清除本地的box.name
	// 3. 进入为初始化状态
	ResetHServer(context.Context, *ResetHServerRequest) (*ResetHServerReply, error)
	// Deprecated: Do not use.
	// ----------------------------- 以下为准备废弃的接口 --------------------------------------
	GetDomainSelfCert(context.Context, *DomainCertRequest) (*DomainCertReply, error)
	// Deprecated: Do not use.
	// 以下接口要改名字
	// 强制将特定设备加到受信任列表中
	TrustUserDevice(context.Context, *TrustUserDeviceRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedHPortalSysServer()
}

// UnimplementedHPortalSysServer must be embedded to have forward compatible implementations.
type UnimplementedHPortalSysServer struct {
}

func (UnimplementedHPortalSysServer) QueryLogin(context.Context, *AuthToken) (*LoginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLogin not implemented")
}
func (UnimplementedHPortalSysServer) QueryBoxInfo(context.Context, *emptypb.Empty) (*BoxInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBoxInfo not implemented")
}
func (UnimplementedHPortalSysServer) GetDomainCert(context.Context, *DomainCertRequest) (*DomainCertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainCert not implemented")
}
func (UnimplementedHPortalSysServer) AllocVirtualExternalIP(context.Context, *AllocVEIPRequest) (*AllocVEIPReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocVirtualExternalIP not implemented")
}
func (UnimplementedHPortalSysServer) FreeVirtualExternalIP(context.Context, *FreeVEIPRequest) (*FreeVEIPReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreeVirtualExternalIP not implemented")
}
func (UnimplementedHPortalSysServer) ListUsers(context.Context, *emptypb.Empty) (*ListUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedHPortalSysServer) CreateUser(context.Context, *CreateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedHPortalSysServer) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedHPortalSysServer) ResetPassword(context.Context, *ResetPasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedHPortalSysServer) CheckPassword(context.Context, *CheckPasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassword not implemented")
}
func (UnimplementedHPortalSysServer) QueryRole(context.Context, *UserID) (*QueryRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRole not implemented")
}
func (UnimplementedHPortalSysServer) ChangeRole(context.Context, *ChangeRoleReqeust) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeRole not implemented")
}
func (UnimplementedHPortalSysServer) ChangeTrustEndDevice(context.Context, *ChangeTrustEndDeviceRequest) (*ChangeTrustEndDeviceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeTrustEndDevice not implemented")
}
func (UnimplementedHPortalSysServer) ListDevices(context.Context, *ListDeviceRequest) (*ListDeviceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevices not implemented")
}
func (UnimplementedHPortalSysServer) QueryDeviceByID(context.Context, *DeviceID) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceByID not implemented")
}
func (UnimplementedHPortalSysServer) ClearLoginSession(context.Context, *ClearLoginSessionRequest) (*ClearLoginSessionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearLoginSession not implemented")
}
func (UnimplementedHPortalSysServer) RemoteSocks(context.Context, *RemoteSocksRequest) (*RemoteSocksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoteSocks not implemented")
}
func (UnimplementedHPortalSysServer) UpdateBoxSystemStatus(context.Context, *UpdateBoxSystemStatusRequest) (*UpdateBoxSystemStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBoxSystemStatus not implemented")
}
func (UnimplementedHPortalSysServer) SetupHServer(context.Context, *SetupHServerRequest) (*SetupHServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupHServer not implemented")
}
func (UnimplementedHPortalSysServer) ResetHServer(context.Context, *ResetHServerRequest) (*ResetHServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetHServer not implemented")
}
func (UnimplementedHPortalSysServer) GetDomainSelfCert(context.Context, *DomainCertRequest) (*DomainCertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainSelfCert not implemented")
}
func (UnimplementedHPortalSysServer) TrustUserDevice(context.Context, *TrustUserDeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrustUserDevice not implemented")
}
func (UnimplementedHPortalSysServer) mustEmbedUnimplementedHPortalSysServer() {}

// UnsafeHPortalSysServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HPortalSysServer will
// result in compilation errors.
type UnsafeHPortalSysServer interface {
	mustEmbedUnimplementedHPortalSysServer()
}

func RegisterHPortalSysServer(s grpc.ServiceRegistrar, srv HPortalSysServer) {
	s.RegisterService(&HPortalSys_ServiceDesc, srv)
}

func _HPortalSys_QueryLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).QueryLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/QueryLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).QueryLogin(ctx, req.(*AuthToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_QueryBoxInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).QueryBoxInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/QueryBoxInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).QueryBoxInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_GetDomainCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DomainCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).GetDomainCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/GetDomainCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).GetDomainCert(ctx, req.(*DomainCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_AllocVirtualExternalIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocVEIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).AllocVirtualExternalIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/AllocVirtualExternalIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).AllocVirtualExternalIP(ctx, req.(*AllocVEIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_FreeVirtualExternalIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeVEIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).FreeVirtualExternalIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/FreeVirtualExternalIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).FreeVirtualExternalIP(ctx, req.(*FreeVEIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).ListUsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_CheckPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).CheckPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/CheckPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).CheckPassword(ctx, req.(*CheckPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_QueryRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).QueryRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/QueryRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).QueryRole(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_ChangeRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeRoleReqeust)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).ChangeRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/ChangeRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).ChangeRole(ctx, req.(*ChangeRoleReqeust))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_ChangeTrustEndDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeTrustEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).ChangeTrustEndDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/ChangeTrustEndDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).ChangeTrustEndDevice(ctx, req.(*ChangeTrustEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/ListDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).ListDevices(ctx, req.(*ListDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_QueryDeviceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).QueryDeviceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/QueryDeviceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).QueryDeviceByID(ctx, req.(*DeviceID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_ClearLoginSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearLoginSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).ClearLoginSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/ClearLoginSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).ClearLoginSession(ctx, req.(*ClearLoginSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_RemoteSocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteSocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).RemoteSocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/RemoteSocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).RemoteSocks(ctx, req.(*RemoteSocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_UpdateBoxSystemStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBoxSystemStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).UpdateBoxSystemStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/UpdateBoxSystemStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).UpdateBoxSystemStatus(ctx, req.(*UpdateBoxSystemStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_SetupHServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupHServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).SetupHServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/SetupHServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).SetupHServer(ctx, req.(*SetupHServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_ResetHServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetHServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).ResetHServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/ResetHServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).ResetHServer(ctx, req.(*ResetHServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_GetDomainSelfCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DomainCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).GetDomainSelfCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/GetDomainSelfCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).GetDomainSelfCert(ctx, req.(*DomainCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_TrustUserDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustUserDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).TrustUserDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/TrustUserDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).TrustUserDevice(ctx, req.(*TrustUserDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HPortalSys_ServiceDesc is the grpc.ServiceDesc for HPortalSys service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HPortalSys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.sys.HPortalSys",
	HandlerType: (*HPortalSysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryLogin",
			Handler:    _HPortalSys_QueryLogin_Handler,
		},
		{
			MethodName: "QueryBoxInfo",
			Handler:    _HPortalSys_QueryBoxInfo_Handler,
		},
		{
			MethodName: "GetDomainCert",
			Handler:    _HPortalSys_GetDomainCert_Handler,
		},
		{
			MethodName: "AllocVirtualExternalIP",
			Handler:    _HPortalSys_AllocVirtualExternalIP_Handler,
		},
		{
			MethodName: "FreeVirtualExternalIP",
			Handler:    _HPortalSys_FreeVirtualExternalIP_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _HPortalSys_ListUsers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _HPortalSys_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _HPortalSys_DeleteUser_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _HPortalSys_ResetPassword_Handler,
		},
		{
			MethodName: "CheckPassword",
			Handler:    _HPortalSys_CheckPassword_Handler,
		},
		{
			MethodName: "QueryRole",
			Handler:    _HPortalSys_QueryRole_Handler,
		},
		{
			MethodName: "ChangeRole",
			Handler:    _HPortalSys_ChangeRole_Handler,
		},
		{
			MethodName: "ChangeTrustEndDevice",
			Handler:    _HPortalSys_ChangeTrustEndDevice_Handler,
		},
		{
			MethodName: "ListDevices",
			Handler:    _HPortalSys_ListDevices_Handler,
		},
		{
			MethodName: "QueryDeviceByID",
			Handler:    _HPortalSys_QueryDeviceByID_Handler,
		},
		{
			MethodName: "ClearLoginSession",
			Handler:    _HPortalSys_ClearLoginSession_Handler,
		},
		{
			MethodName: "RemoteSocks",
			Handler:    _HPortalSys_RemoteSocks_Handler,
		},
		{
			MethodName: "UpdateBoxSystemStatus",
			Handler:    _HPortalSys_UpdateBoxSystemStatus_Handler,
		},
		{
			MethodName: "SetupHServer",
			Handler:    _HPortalSys_SetupHServer_Handler,
		},
		{
			MethodName: "ResetHServer",
			Handler:    _HPortalSys_ResetHServer_Handler,
		},
		{
			MethodName: "GetDomainSelfCert",
			Handler:    _HPortalSys_GetDomainSelfCert_Handler,
		},
		{
			MethodName: "TrustUserDevice",
			Handler:    _HPortalSys_TrustUserDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys/portal-server/portal-server.proto",
}

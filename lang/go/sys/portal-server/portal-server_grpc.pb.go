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
	// 根据UID返回所有的设备列表
	ListDevices(ctx context.Context, in *ListDeviceRequest, opts ...grpc.CallOption) (*ListDeviceReply, error)
	QueryDeviceByID(ctx context.Context, in *DeviceID, opts ...grpc.CallOption) (*Device, error)
	QueryBoxInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BoxInfo, error)
	// 获取盒子所属域名下或下一级域名的https证书。
	// 注意不是所有ACME服务器都支持泛域名。
	GetDomainCert(ctx context.Context, in *DomainCertRequest, opts ...grpc.CallOption) (*DomainCertReply, error)
	GetDomainSelfCert(ctx context.Context, in *DomainCertRequest, opts ...grpc.CallOption) (*DomainCertReply, error)
	// 在部署具体app前，调用此接口获取app证书
	// APP证书格式为:
	//   Issuer: O = $BOX.ORIGIN, CN = $BOX.DOMAIN, serialNumber = $BOX.ID
	//   Subject: O = $BOX.ORIGIN, CN = $APP.DOMAIN, serialNumber = '$UID@$APP.ID || $APP.ID'
	// Issuer为box.crt，通过QueryBoxInfo查询到BoxInfo.BoxCrt。并且box.crt的公钥与box.id是一一对应关系。
	//
	// 盒子内部组件可以直接通过QueryBoxInfo来验证信任链是否合法，盒子外部系统需要通过其他机制比如libp2p.identify去验证box.crt的合法性。
	//
	GetAppCert(ctx context.Context, in *AppCertRequest, opts ...grpc.CallOption) (*AppCertReply, error)
	// 申请额外的外部可访问的IP,并配置对应访问域名
	AllocVirtualExternalIP(ctx context.Context, in *AllocVEIPRequest, opts ...grpc.CallOption) (*AllocVEIPReply, error)
	// 释放虚拟IP
	FreeVirtualExternalIP(ctx context.Context, in *FreeVEIPRequest, opts ...grpc.CallOption) (*FreeVEIPReply, error)
	PairDevices(ctx context.Context, in *PairDevicesRequest, opts ...grpc.CallOption) (HPortalSys_PairDevicesClient, error)
	//  查询所有UID
	ListUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUsersReply, error)
	//  根据用户uid查询用户信息
	QueryRole(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*QueryRoleReply, error)
	//  修改指定uid的用户角色
	ChangeRole(ctx context.Context, in *ChangeRoleReqeust, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//  通过验证旧密码修改新的密码
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//  删除用户信息
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//  创建用户信息
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//  强制重置用户密码
	ForceResetPassword(ctx context.Context, in *ForceResetPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 生成用户注册token,以便上层实现各类用户注册机制
	GenUserInvitation(ctx context.Context, in *GenUserInvitationRequest, opts ...grpc.CallOption) (*UserInvitation, error)
	// 注销当前用户指定设备，同时将连接断开
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//校验用户密码是否正确
	CheckPassword(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoteSocks(ctx context.Context, in *RemoteSocksRequest, opts ...grpc.CallOption) (*RemoteSocksReply, error)
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

func (c *hPortalSysClient) GetDomainSelfCert(ctx context.Context, in *DomainCertRequest, opts ...grpc.CallOption) (*DomainCertReply, error) {
	out := new(DomainCertReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/GetDomainSelfCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) GetAppCert(ctx context.Context, in *AppCertRequest, opts ...grpc.CallOption) (*AppCertReply, error) {
	out := new(AppCertReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/GetAppCert", in, out, opts...)
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

func (c *hPortalSysClient) PairDevices(ctx context.Context, in *PairDevicesRequest, opts ...grpc.CallOption) (HPortalSys_PairDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &HPortalSys_ServiceDesc.Streams[0], "/cloud.lazycat.apis.sys.HPortalSys/PairDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &hPortalSysPairDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HPortalSys_PairDevicesClient interface {
	Recv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type hPortalSysPairDevicesClient struct {
	grpc.ClientStream
}

func (x *hPortalSysPairDevicesClient) Recv() (*emptypb.Empty, error) {
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hPortalSysClient) ListUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListUsersReply, error) {
	out := new(ListUsersReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/ListUsers", in, out, opts...)
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

func (c *hPortalSysClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/ResetPassword", in, out, opts...)
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

func (c *hPortalSysClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) ForceResetPassword(ctx context.Context, in *ForceResetPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/ForceResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) GenUserInvitation(ctx context.Context, in *GenUserInvitationRequest, opts ...grpc.CallOption) (*UserInvitation, error) {
	out := new(UserInvitation)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/GenUserInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hPortalSysClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/Logout", in, out, opts...)
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

func (c *hPortalSysClient) RemoteSocks(ctx context.Context, in *RemoteSocksRequest, opts ...grpc.CallOption) (*RemoteSocksReply, error) {
	out := new(RemoteSocksReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.HPortalSys/RemoteSocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

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
	// 根据UID返回所有的设备列表
	ListDevices(context.Context, *ListDeviceRequest) (*ListDeviceReply, error)
	QueryDeviceByID(context.Context, *DeviceID) (*Device, error)
	QueryBoxInfo(context.Context, *emptypb.Empty) (*BoxInfo, error)
	// 获取盒子所属域名下或下一级域名的https证书。
	// 注意不是所有ACME服务器都支持泛域名。
	GetDomainCert(context.Context, *DomainCertRequest) (*DomainCertReply, error)
	GetDomainSelfCert(context.Context, *DomainCertRequest) (*DomainCertReply, error)
	// 在部署具体app前，调用此接口获取app证书
	// APP证书格式为:
	//   Issuer: O = $BOX.ORIGIN, CN = $BOX.DOMAIN, serialNumber = $BOX.ID
	//   Subject: O = $BOX.ORIGIN, CN = $APP.DOMAIN, serialNumber = '$UID@$APP.ID || $APP.ID'
	// Issuer为box.crt，通过QueryBoxInfo查询到BoxInfo.BoxCrt。并且box.crt的公钥与box.id是一一对应关系。
	//
	// 盒子内部组件可以直接通过QueryBoxInfo来验证信任链是否合法，盒子外部系统需要通过其他机制比如libp2p.identify去验证box.crt的合法性。
	//
	GetAppCert(context.Context, *AppCertRequest) (*AppCertReply, error)
	// 申请额外的外部可访问的IP,并配置对应访问域名
	AllocVirtualExternalIP(context.Context, *AllocVEIPRequest) (*AllocVEIPReply, error)
	// 释放虚拟IP
	FreeVirtualExternalIP(context.Context, *FreeVEIPRequest) (*FreeVEIPReply, error)
	PairDevices(*PairDevicesRequest, HPortalSys_PairDevicesServer) error
	//  查询所有UID
	ListUsers(context.Context, *emptypb.Empty) (*ListUsersReply, error)
	//  根据用户uid查询用户信息
	QueryRole(context.Context, *UserID) (*QueryRoleReply, error)
	//  修改指定uid的用户角色
	ChangeRole(context.Context, *ChangeRoleReqeust) (*emptypb.Empty, error)
	//  通过验证旧密码修改新的密码
	ResetPassword(context.Context, *ResetPasswordRequest) (*emptypb.Empty, error)
	//  删除用户信息
	DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error)
	//  创建用户信息
	CreateUser(context.Context, *CreateUserRequest) (*emptypb.Empty, error)
	//  强制重置用户密码
	ForceResetPassword(context.Context, *ForceResetPasswordRequest) (*emptypb.Empty, error)
	// 生成用户注册token,以便上层实现各类用户注册机制
	GenUserInvitation(context.Context, *GenUserInvitationRequest) (*UserInvitation, error)
	// 注销当前用户指定设备，同时将连接断开
	Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error)
	//校验用户密码是否正确
	CheckPassword(context.Context, *CheckPasswordRequest) (*emptypb.Empty, error)
	RemoteSocks(context.Context, *RemoteSocksRequest) (*RemoteSocksReply, error)
	TrustUserDevice(context.Context, *TrustUserDeviceRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedHPortalSysServer()
}

// UnimplementedHPortalSysServer must be embedded to have forward compatible implementations.
type UnimplementedHPortalSysServer struct {
}

func (UnimplementedHPortalSysServer) QueryLogin(context.Context, *AuthToken) (*LoginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLogin not implemented")
}
func (UnimplementedHPortalSysServer) ListDevices(context.Context, *ListDeviceRequest) (*ListDeviceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevices not implemented")
}
func (UnimplementedHPortalSysServer) QueryDeviceByID(context.Context, *DeviceID) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceByID not implemented")
}
func (UnimplementedHPortalSysServer) QueryBoxInfo(context.Context, *emptypb.Empty) (*BoxInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBoxInfo not implemented")
}
func (UnimplementedHPortalSysServer) GetDomainCert(context.Context, *DomainCertRequest) (*DomainCertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainCert not implemented")
}
func (UnimplementedHPortalSysServer) GetDomainSelfCert(context.Context, *DomainCertRequest) (*DomainCertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomainSelfCert not implemented")
}
func (UnimplementedHPortalSysServer) GetAppCert(context.Context, *AppCertRequest) (*AppCertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppCert not implemented")
}
func (UnimplementedHPortalSysServer) AllocVirtualExternalIP(context.Context, *AllocVEIPRequest) (*AllocVEIPReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocVirtualExternalIP not implemented")
}
func (UnimplementedHPortalSysServer) FreeVirtualExternalIP(context.Context, *FreeVEIPRequest) (*FreeVEIPReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreeVirtualExternalIP not implemented")
}
func (UnimplementedHPortalSysServer) PairDevices(*PairDevicesRequest, HPortalSys_PairDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method PairDevices not implemented")
}
func (UnimplementedHPortalSysServer) ListUsers(context.Context, *emptypb.Empty) (*ListUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedHPortalSysServer) QueryRole(context.Context, *UserID) (*QueryRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRole not implemented")
}
func (UnimplementedHPortalSysServer) ChangeRole(context.Context, *ChangeRoleReqeust) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeRole not implemented")
}
func (UnimplementedHPortalSysServer) ResetPassword(context.Context, *ResetPasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedHPortalSysServer) DeleteUser(context.Context, *DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedHPortalSysServer) CreateUser(context.Context, *CreateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedHPortalSysServer) ForceResetPassword(context.Context, *ForceResetPasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceResetPassword not implemented")
}
func (UnimplementedHPortalSysServer) GenUserInvitation(context.Context, *GenUserInvitationRequest) (*UserInvitation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenUserInvitation not implemented")
}
func (UnimplementedHPortalSysServer) Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedHPortalSysServer) CheckPassword(context.Context, *CheckPasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassword not implemented")
}
func (UnimplementedHPortalSysServer) RemoteSocks(context.Context, *RemoteSocksRequest) (*RemoteSocksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoteSocks not implemented")
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

func _HPortalSys_GetAppCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).GetAppCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/GetAppCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).GetAppCert(ctx, req.(*AppCertRequest))
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

func _HPortalSys_PairDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PairDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HPortalSysServer).PairDevices(m, &hPortalSysPairDevicesServer{stream})
}

type HPortalSys_PairDevicesServer interface {
	Send(*emptypb.Empty) error
	grpc.ServerStream
}

type hPortalSysPairDevicesServer struct {
	grpc.ServerStream
}

func (x *hPortalSysPairDevicesServer) Send(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
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

func _HPortalSys_ForceResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForceResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).ForceResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/ForceResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).ForceResetPassword(ctx, req.(*ForceResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_GenUserInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenUserInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).GenUserInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/GenUserInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).GenUserInvitation(ctx, req.(*GenUserInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HPortalSys_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HPortalSysServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.HPortalSys/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HPortalSysServer).Logout(ctx, req.(*LogoutRequest))
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
			MethodName: "ListDevices",
			Handler:    _HPortalSys_ListDevices_Handler,
		},
		{
			MethodName: "QueryDeviceByID",
			Handler:    _HPortalSys_QueryDeviceByID_Handler,
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
			MethodName: "GetDomainSelfCert",
			Handler:    _HPortalSys_GetDomainSelfCert_Handler,
		},
		{
			MethodName: "GetAppCert",
			Handler:    _HPortalSys_GetAppCert_Handler,
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
			MethodName: "QueryRole",
			Handler:    _HPortalSys_QueryRole_Handler,
		},
		{
			MethodName: "ChangeRole",
			Handler:    _HPortalSys_ChangeRole_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _HPortalSys_ResetPassword_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _HPortalSys_DeleteUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _HPortalSys_CreateUser_Handler,
		},
		{
			MethodName: "ForceResetPassword",
			Handler:    _HPortalSys_ForceResetPassword_Handler,
		},
		{
			MethodName: "GenUserInvitation",
			Handler:    _HPortalSys_GenUserInvitation_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _HPortalSys_Logout_Handler,
		},
		{
			MethodName: "CheckPassword",
			Handler:    _HPortalSys_CheckPassword_Handler,
		},
		{
			MethodName: "RemoteSocks",
			Handler:    _HPortalSys_RemoteSocks_Handler,
		},
		{
			MethodName: "TrustUserDevice",
			Handler:    _HPortalSys_TrustUserDevice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PairDevices",
			Handler:       _HPortalSys_PairDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sys/portal-server/portal-server.proto",
}

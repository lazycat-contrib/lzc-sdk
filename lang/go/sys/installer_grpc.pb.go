// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: sys/installer.proto

package sys

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

// InstallerServiceClient is the client API for InstallerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstallerServiceClient interface {
	// 初始化微服
	BoxSetup(ctx context.Context, in *BoxSetupRequest, opts ...grpc.CallOption) (InstallerService_BoxSetupClient, error)
	// 微服是否有互联网
	HasInternet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HasInternetResponse, error)
	// 列出内部缓存中的 Wi-Fi 列表， 如果没有则会扫描后返回
	WifiList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AccessPointInfoList, error)
	// 连接一个Wi-Fi
	WifiConnect(ctx context.Context, in *WifiConnectInfo, opts ...grpc.CallOption) (*WifiConnectReply, error)
	// 手动链接一个Wi-Fi（用于连接隐藏网络）
	WifiConfigAdd(ctx context.Context, in *WifiConnectInfo, opts ...grpc.CallOption) (*WifiConnectReply, error)
}

type installerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstallerServiceClient(cc grpc.ClientConnInterface) InstallerServiceClient {
	return &installerServiceClient{cc}
}

func (c *installerServiceClient) BoxSetup(ctx context.Context, in *BoxSetupRequest, opts ...grpc.CallOption) (InstallerService_BoxSetupClient, error) {
	stream, err := c.cc.NewStream(ctx, &InstallerService_ServiceDesc.Streams[0], "/cloud.lazycat.apis.sys.InstallerService/BoxSetup", opts...)
	if err != nil {
		return nil, err
	}
	x := &installerServiceBoxSetupClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InstallerService_BoxSetupClient interface {
	Recv() (*BoxSetupReply, error)
	grpc.ClientStream
}

type installerServiceBoxSetupClient struct {
	grpc.ClientStream
}

func (x *installerServiceBoxSetupClient) Recv() (*BoxSetupReply, error) {
	m := new(BoxSetupReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *installerServiceClient) HasInternet(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HasInternetResponse, error) {
	out := new(HasInternetResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.InstallerService/HasInternet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installerServiceClient) WifiList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AccessPointInfoList, error) {
	out := new(AccessPointInfoList)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.InstallerService/WifiList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installerServiceClient) WifiConnect(ctx context.Context, in *WifiConnectInfo, opts ...grpc.CallOption) (*WifiConnectReply, error) {
	out := new(WifiConnectReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.InstallerService/WifiConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installerServiceClient) WifiConfigAdd(ctx context.Context, in *WifiConnectInfo, opts ...grpc.CallOption) (*WifiConnectReply, error) {
	out := new(WifiConnectReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.InstallerService/WifiConfigAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstallerServiceServer is the server API for InstallerService service.
// All implementations must embed UnimplementedInstallerServiceServer
// for forward compatibility
type InstallerServiceServer interface {
	// 初始化微服
	BoxSetup(*BoxSetupRequest, InstallerService_BoxSetupServer) error
	// 微服是否有互联网
	HasInternet(context.Context, *emptypb.Empty) (*HasInternetResponse, error)
	// 列出内部缓存中的 Wi-Fi 列表， 如果没有则会扫描后返回
	WifiList(context.Context, *emptypb.Empty) (*AccessPointInfoList, error)
	// 连接一个Wi-Fi
	WifiConnect(context.Context, *WifiConnectInfo) (*WifiConnectReply, error)
	// 手动链接一个Wi-Fi（用于连接隐藏网络）
	WifiConfigAdd(context.Context, *WifiConnectInfo) (*WifiConnectReply, error)
	mustEmbedUnimplementedInstallerServiceServer()
}

// UnimplementedInstallerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInstallerServiceServer struct {
}

func (UnimplementedInstallerServiceServer) BoxSetup(*BoxSetupRequest, InstallerService_BoxSetupServer) error {
	return status.Errorf(codes.Unimplemented, "method BoxSetup not implemented")
}
func (UnimplementedInstallerServiceServer) HasInternet(context.Context, *emptypb.Empty) (*HasInternetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasInternet not implemented")
}
func (UnimplementedInstallerServiceServer) WifiList(context.Context, *emptypb.Empty) (*AccessPointInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WifiList not implemented")
}
func (UnimplementedInstallerServiceServer) WifiConnect(context.Context, *WifiConnectInfo) (*WifiConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WifiConnect not implemented")
}
func (UnimplementedInstallerServiceServer) WifiConfigAdd(context.Context, *WifiConnectInfo) (*WifiConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WifiConfigAdd not implemented")
}
func (UnimplementedInstallerServiceServer) mustEmbedUnimplementedInstallerServiceServer() {}

// UnsafeInstallerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstallerServiceServer will
// result in compilation errors.
type UnsafeInstallerServiceServer interface {
	mustEmbedUnimplementedInstallerServiceServer()
}

func RegisterInstallerServiceServer(s grpc.ServiceRegistrar, srv InstallerServiceServer) {
	s.RegisterService(&InstallerService_ServiceDesc, srv)
}

func _InstallerService_BoxSetup_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BoxSetupRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InstallerServiceServer).BoxSetup(m, &installerServiceBoxSetupServer{stream})
}

type InstallerService_BoxSetupServer interface {
	Send(*BoxSetupReply) error
	grpc.ServerStream
}

type installerServiceBoxSetupServer struct {
	grpc.ServerStream
}

func (x *installerServiceBoxSetupServer) Send(m *BoxSetupReply) error {
	return x.ServerStream.SendMsg(m)
}

func _InstallerService_HasInternet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallerServiceServer).HasInternet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.InstallerService/HasInternet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallerServiceServer).HasInternet(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstallerService_WifiList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallerServiceServer).WifiList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.InstallerService/WifiList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallerServiceServer).WifiList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstallerService_WifiConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiConnectInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallerServiceServer).WifiConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.InstallerService/WifiConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallerServiceServer).WifiConnect(ctx, req.(*WifiConnectInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstallerService_WifiConfigAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiConnectInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallerServiceServer).WifiConfigAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.InstallerService/WifiConfigAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallerServiceServer).WifiConfigAdd(ctx, req.(*WifiConnectInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// InstallerService_ServiceDesc is the grpc.ServiceDesc for InstallerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InstallerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.sys.InstallerService",
	HandlerType: (*InstallerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HasInternet",
			Handler:    _InstallerService_HasInternet_Handler,
		},
		{
			MethodName: "WifiList",
			Handler:    _InstallerService_WifiList_Handler,
		},
		{
			MethodName: "WifiConnect",
			Handler:    _InstallerService_WifiConnect_Handler,
		},
		{
			MethodName: "WifiConfigAdd",
			Handler:    _InstallerService_WifiConfigAdd_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BoxSetup",
			Handler:       _InstallerService_BoxSetup_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sys/installer.proto",
}

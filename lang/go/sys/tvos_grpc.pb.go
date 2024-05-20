// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: sys/tvos.proto

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

const (
	TvOS_Run_FullMethodName             = "/cloud.lazycat.apis.sys.TvOS/Run"
	TvOS_Status_FullMethodName          = "/cloud.lazycat.apis.sys.TvOS/Status"
	TvOS_Stop_FullMethodName            = "/cloud.lazycat.apis.sys.TvOS/Stop"
	TvOS_IsHDMIConnected_FullMethodName = "/cloud.lazycat.apis.sys.TvOS/IsHDMIConnected"
)

// TvOSClient is the client API for TvOS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TvOSClient interface {
	// 启动
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 状态
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	// 停止
	Stop(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 是否连接HDMI
	IsHDMIConnected(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IsHDMIConnectedResponse, error)
}

type tvOSClient struct {
	cc grpc.ClientConnInterface
}

func NewTvOSClient(cc grpc.ClientConnInterface) TvOSClient {
	return &tvOSClient{cc}
}

func (c *tvOSClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TvOS_Run_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tvOSClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, TvOS_Status_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tvOSClient) Stop(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TvOS_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tvOSClient) IsHDMIConnected(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IsHDMIConnectedResponse, error) {
	out := new(IsHDMIConnectedResponse)
	err := c.cc.Invoke(ctx, TvOS_IsHDMIConnected_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TvOSServer is the server API for TvOS service.
// All implementations must embed UnimplementedTvOSServer
// for forward compatibility
type TvOSServer interface {
	// 启动
	Run(context.Context, *RunRequest) (*emptypb.Empty, error)
	// 状态
	Status(context.Context, *emptypb.Empty) (*StatusResponse, error)
	// 停止
	Stop(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// 是否连接HDMI
	IsHDMIConnected(context.Context, *emptypb.Empty) (*IsHDMIConnectedResponse, error)
	mustEmbedUnimplementedTvOSServer()
}

// UnimplementedTvOSServer must be embedded to have forward compatible implementations.
type UnimplementedTvOSServer struct {
}

func (UnimplementedTvOSServer) Run(context.Context, *RunRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedTvOSServer) Status(context.Context, *emptypb.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedTvOSServer) Stop(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedTvOSServer) IsHDMIConnected(context.Context, *emptypb.Empty) (*IsHDMIConnectedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsHDMIConnected not implemented")
}
func (UnimplementedTvOSServer) mustEmbedUnimplementedTvOSServer() {}

// UnsafeTvOSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TvOSServer will
// result in compilation errors.
type UnsafeTvOSServer interface {
	mustEmbedUnimplementedTvOSServer()
}

func RegisterTvOSServer(s grpc.ServiceRegistrar, srv TvOSServer) {
	s.RegisterService(&TvOS_ServiceDesc, srv)
}

func _TvOS_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TvOSServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TvOS_Run_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TvOSServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TvOS_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TvOSServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TvOS_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TvOSServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TvOS_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TvOSServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TvOS_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TvOSServer).Stop(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TvOS_IsHDMIConnected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TvOSServer).IsHDMIConnected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TvOS_IsHDMIConnected_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TvOSServer).IsHDMIConnected(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TvOS_ServiceDesc is the grpc.ServiceDesc for TvOS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TvOS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.sys.TvOS",
	HandlerType: (*TvOSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _TvOS_Run_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _TvOS_Status_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _TvOS_Stop_Handler,
		},
		{
			MethodName: "IsHDMIConnected",
			Handler:    _TvOS_IsHDMIConnected_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys/tvos.proto",
}

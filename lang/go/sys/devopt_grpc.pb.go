// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: sys/devopt.proto

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
	DevOptService_GetDeveloperOptions_FullMethodName = "/cloud.lazycat.apis.sys.DevOptService/GetDeveloperOptions"
	DevOptService_SetDeveloperOptions_FullMethodName = "/cloud.lazycat.apis.sys.DevOptService/SetDeveloperOptions"
	DevOptService_SshdEnable_FullMethodName          = "/cloud.lazycat.apis.sys.DevOptService/SshdEnable"
	DevOptService_SshdEnabled_FullMethodName         = "/cloud.lazycat.apis.sys.DevOptService/SshdEnabled"
)

// DevOptServiceClient is the client API for DevOptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevOptServiceClient interface {
	GetDeveloperOptions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DeveloperOptions, error)
	SetDeveloperOptions(ctx context.Context, in *DeveloperOptions, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SshdEnable(ctx context.Context, in *SshdEnableRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SshdEnabled(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EnableSshdResponse, error)
}

type devOptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDevOptServiceClient(cc grpc.ClientConnInterface) DevOptServiceClient {
	return &devOptServiceClient{cc}
}

func (c *devOptServiceClient) GetDeveloperOptions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DeveloperOptions, error) {
	out := new(DeveloperOptions)
	err := c.cc.Invoke(ctx, DevOptService_GetDeveloperOptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devOptServiceClient) SetDeveloperOptions(ctx context.Context, in *DeveloperOptions, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DevOptService_SetDeveloperOptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devOptServiceClient) SshdEnable(ctx context.Context, in *SshdEnableRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DevOptService_SshdEnable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devOptServiceClient) SshdEnabled(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EnableSshdResponse, error) {
	out := new(EnableSshdResponse)
	err := c.cc.Invoke(ctx, DevOptService_SshdEnabled_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevOptServiceServer is the server API for DevOptService service.
// All implementations must embed UnimplementedDevOptServiceServer
// for forward compatibility
type DevOptServiceServer interface {
	GetDeveloperOptions(context.Context, *emptypb.Empty) (*DeveloperOptions, error)
	SetDeveloperOptions(context.Context, *DeveloperOptions) (*emptypb.Empty, error)
	SshdEnable(context.Context, *SshdEnableRequest) (*emptypb.Empty, error)
	SshdEnabled(context.Context, *emptypb.Empty) (*EnableSshdResponse, error)
	mustEmbedUnimplementedDevOptServiceServer()
}

// UnimplementedDevOptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDevOptServiceServer struct {
}

func (UnimplementedDevOptServiceServer) GetDeveloperOptions(context.Context, *emptypb.Empty) (*DeveloperOptions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeveloperOptions not implemented")
}
func (UnimplementedDevOptServiceServer) SetDeveloperOptions(context.Context, *DeveloperOptions) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeveloperOptions not implemented")
}
func (UnimplementedDevOptServiceServer) SshdEnable(context.Context, *SshdEnableRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SshdEnable not implemented")
}
func (UnimplementedDevOptServiceServer) SshdEnabled(context.Context, *emptypb.Empty) (*EnableSshdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SshdEnabled not implemented")
}
func (UnimplementedDevOptServiceServer) mustEmbedUnimplementedDevOptServiceServer() {}

// UnsafeDevOptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevOptServiceServer will
// result in compilation errors.
type UnsafeDevOptServiceServer interface {
	mustEmbedUnimplementedDevOptServiceServer()
}

func RegisterDevOptServiceServer(s grpc.ServiceRegistrar, srv DevOptServiceServer) {
	s.RegisterService(&DevOptService_ServiceDesc, srv)
}

func _DevOptService_GetDeveloperOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevOptServiceServer).GetDeveloperOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevOptService_GetDeveloperOptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevOptServiceServer).GetDeveloperOptions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevOptService_SetDeveloperOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeveloperOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevOptServiceServer).SetDeveloperOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevOptService_SetDeveloperOptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevOptServiceServer).SetDeveloperOptions(ctx, req.(*DeveloperOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevOptService_SshdEnable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SshdEnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevOptServiceServer).SshdEnable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevOptService_SshdEnable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevOptServiceServer).SshdEnable(ctx, req.(*SshdEnableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevOptService_SshdEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevOptServiceServer).SshdEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevOptService_SshdEnabled_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevOptServiceServer).SshdEnabled(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DevOptService_ServiceDesc is the grpc.ServiceDesc for DevOptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DevOptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.sys.DevOptService",
	HandlerType: (*DevOptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeveloperOptions",
			Handler:    _DevOptService_GetDeveloperOptions_Handler,
		},
		{
			MethodName: "SetDeveloperOptions",
			Handler:    _DevOptService_SetDeveloperOptions_Handler,
		},
		{
			MethodName: "SshdEnable",
			Handler:    _DevOptService_SshdEnable_Handler,
		},
		{
			MethodName: "SshdEnabled",
			Handler:    _DevOptService_SshdEnabled_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys/devopt.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: common/box.proto

package common

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
	BoxService_QueryInfo_FullMethodName         = "/cloud.lazycat.apis.common.BoxService/QueryInfo"
	BoxService_ChangeDisplayName_FullMethodName = "/cloud.lazycat.apis.common.BoxService/ChangeDisplayName"
	BoxService_SetBootOption_FullMethodName     = "/cloud.lazycat.apis.common.BoxService/SetBootOption"
	BoxService_Shutdown_FullMethodName          = "/cloud.lazycat.apis.common.BoxService/Shutdown"
	BoxService_QueryDisksInfo_FullMethodName    = "/cloud.lazycat.apis.common.BoxService/QueryDisksInfo"
)

// BoxServiceClient is the client API for BoxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoxServiceClient interface {
	QueryInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BoxInfo, error)
	ChangeDisplayName(ctx context.Context, in *ChangeDisplayNameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetBootOption(ctx context.Context, in *BootOption, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	QueryDisksInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DisksInfo, error)
}

type boxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoxServiceClient(cc grpc.ClientConnInterface) BoxServiceClient {
	return &boxServiceClient{cc}
}

func (c *boxServiceClient) QueryInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BoxInfo, error) {
	out := new(BoxInfo)
	err := c.cc.Invoke(ctx, BoxService_QueryInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) ChangeDisplayName(ctx context.Context, in *ChangeDisplayNameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BoxService_ChangeDisplayName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) SetBootOption(ctx context.Context, in *BootOption, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BoxService_SetBootOption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BoxService_Shutdown_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) QueryDisksInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DisksInfo, error) {
	out := new(DisksInfo)
	err := c.cc.Invoke(ctx, BoxService_QueryDisksInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoxServiceServer is the server API for BoxService service.
// All implementations must embed UnimplementedBoxServiceServer
// for forward compatibility
type BoxServiceServer interface {
	QueryInfo(context.Context, *emptypb.Empty) (*BoxInfo, error)
	ChangeDisplayName(context.Context, *ChangeDisplayNameRequest) (*emptypb.Empty, error)
	SetBootOption(context.Context, *BootOption) (*emptypb.Empty, error)
	Shutdown(context.Context, *ShutdownRequest) (*emptypb.Empty, error)
	QueryDisksInfo(context.Context, *emptypb.Empty) (*DisksInfo, error)
	mustEmbedUnimplementedBoxServiceServer()
}

// UnimplementedBoxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBoxServiceServer struct {
}

func (UnimplementedBoxServiceServer) QueryInfo(context.Context, *emptypb.Empty) (*BoxInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryInfo not implemented")
}
func (UnimplementedBoxServiceServer) ChangeDisplayName(context.Context, *ChangeDisplayNameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeDisplayName not implemented")
}
func (UnimplementedBoxServiceServer) SetBootOption(context.Context, *BootOption) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBootOption not implemented")
}
func (UnimplementedBoxServiceServer) Shutdown(context.Context, *ShutdownRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedBoxServiceServer) QueryDisksInfo(context.Context, *emptypb.Empty) (*DisksInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDisksInfo not implemented")
}
func (UnimplementedBoxServiceServer) mustEmbedUnimplementedBoxServiceServer() {}

// UnsafeBoxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoxServiceServer will
// result in compilation errors.
type UnsafeBoxServiceServer interface {
	mustEmbedUnimplementedBoxServiceServer()
}

func RegisterBoxServiceServer(s grpc.ServiceRegistrar, srv BoxServiceServer) {
	s.RegisterService(&BoxService_ServiceDesc, srv)
}

func _BoxService_QueryInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).QueryInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_QueryInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).QueryInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_ChangeDisplayName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeDisplayNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).ChangeDisplayName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_ChangeDisplayName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).ChangeDisplayName(ctx, req.(*ChangeDisplayNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_SetBootOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).SetBootOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_SetBootOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).SetBootOption(ctx, req.(*BootOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_Shutdown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_QueryDisksInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).QueryDisksInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_QueryDisksInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).QueryDisksInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BoxService_ServiceDesc is the grpc.ServiceDesc for BoxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.common.BoxService",
	HandlerType: (*BoxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryInfo",
			Handler:    _BoxService_QueryInfo_Handler,
		},
		{
			MethodName: "ChangeDisplayName",
			Handler:    _BoxService_ChangeDisplayName_Handler,
		},
		{
			MethodName: "SetBootOption",
			Handler:    _BoxService_SetBootOption_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _BoxService_Shutdown_Handler,
		},
		{
			MethodName: "QueryDisksInfo",
			Handler:    _BoxService_QueryDisksInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/box.proto",
}

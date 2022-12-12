// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: localdevice/clipboard.proto

package localdevice

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

// ClipboardManagerClient is the client API for ClipboardManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClipboardManagerClient interface {
	GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClipConfig, error)
	Read(ctx context.Context, in *ReadClipRequest, opts ...grpc.CallOption) (*ReadClipReply, error)
	Write(ctx context.Context, in *WriteClipRequest, opts ...grpc.CallOption) (*WriteClipReply, error)
	Watch(ctx context.Context, in *ReadClipRequest, opts ...grpc.CallOption) (ClipboardManager_WatchClient, error)
}

type clipboardManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewClipboardManagerClient(cc grpc.ClientConnInterface) ClipboardManagerClient {
	return &clipboardManagerClient{cc}
}

func (c *clipboardManagerClient) GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ClipConfig, error) {
	out := new(ClipConfig)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.localdevice.ClipboardManager/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clipboardManagerClient) Read(ctx context.Context, in *ReadClipRequest, opts ...grpc.CallOption) (*ReadClipReply, error) {
	out := new(ReadClipReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.localdevice.ClipboardManager/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clipboardManagerClient) Write(ctx context.Context, in *WriteClipRequest, opts ...grpc.CallOption) (*WriteClipReply, error) {
	out := new(WriteClipReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.localdevice.ClipboardManager/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clipboardManagerClient) Watch(ctx context.Context, in *ReadClipRequest, opts ...grpc.CallOption) (ClipboardManager_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClipboardManager_ServiceDesc.Streams[0], "/cloud.lazycat.apis.localdevice.ClipboardManager/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &clipboardManagerWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClipboardManager_WatchClient interface {
	Recv() (*ReadClipReply, error)
	grpc.ClientStream
}

type clipboardManagerWatchClient struct {
	grpc.ClientStream
}

func (x *clipboardManagerWatchClient) Recv() (*ReadClipReply, error) {
	m := new(ReadClipReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClipboardManagerServer is the server API for ClipboardManager service.
// All implementations must embed UnimplementedClipboardManagerServer
// for forward compatibility
type ClipboardManagerServer interface {
	GetConfig(context.Context, *emptypb.Empty) (*ClipConfig, error)
	Read(context.Context, *ReadClipRequest) (*ReadClipReply, error)
	Write(context.Context, *WriteClipRequest) (*WriteClipReply, error)
	Watch(*ReadClipRequest, ClipboardManager_WatchServer) error
	mustEmbedUnimplementedClipboardManagerServer()
}

// UnimplementedClipboardManagerServer must be embedded to have forward compatible implementations.
type UnimplementedClipboardManagerServer struct {
}

func (UnimplementedClipboardManagerServer) GetConfig(context.Context, *emptypb.Empty) (*ClipConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedClipboardManagerServer) Read(context.Context, *ReadClipRequest) (*ReadClipReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedClipboardManagerServer) Write(context.Context, *WriteClipRequest) (*WriteClipReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedClipboardManagerServer) Watch(*ReadClipRequest, ClipboardManager_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedClipboardManagerServer) mustEmbedUnimplementedClipboardManagerServer() {}

// UnsafeClipboardManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClipboardManagerServer will
// result in compilation errors.
type UnsafeClipboardManagerServer interface {
	mustEmbedUnimplementedClipboardManagerServer()
}

func RegisterClipboardManagerServer(s grpc.ServiceRegistrar, srv ClipboardManagerServer) {
	s.RegisterService(&ClipboardManager_ServiceDesc, srv)
}

func _ClipboardManager_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClipboardManagerServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.localdevice.ClipboardManager/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClipboardManagerServer).GetConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClipboardManager_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadClipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClipboardManagerServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.localdevice.ClipboardManager/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClipboardManagerServer).Read(ctx, req.(*ReadClipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClipboardManager_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteClipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClipboardManagerServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.localdevice.ClipboardManager/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClipboardManagerServer).Write(ctx, req.(*WriteClipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClipboardManager_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadClipRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClipboardManagerServer).Watch(m, &clipboardManagerWatchServer{stream})
}

type ClipboardManager_WatchServer interface {
	Send(*ReadClipReply) error
	grpc.ServerStream
}

type clipboardManagerWatchServer struct {
	grpc.ServerStream
}

func (x *clipboardManagerWatchServer) Send(m *ReadClipReply) error {
	return x.ServerStream.SendMsg(m)
}

// ClipboardManager_ServiceDesc is the grpc.ServiceDesc for ClipboardManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClipboardManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.localdevice.ClipboardManager",
	HandlerType: (*ClipboardManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _ClipboardManager_GetConfig_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ClipboardManager_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _ClipboardManager_Write_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _ClipboardManager_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "localdevice/clipboard.proto",
}

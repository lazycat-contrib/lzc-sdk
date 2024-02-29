// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: localdevice/remote-input-method.proto

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

const (
	Rim_ListenInputFocus_FullMethodName = "/cloud.lazycat.apis.localdevice.Rim/ListenInputFocus"
	Rim_ListenInputBlur_FullMethodName  = "/cloud.lazycat.apis.localdevice.Rim/ListenInputBlur"
	Rim_GetInputText_FullMethodName     = "/cloud.lazycat.apis.localdevice.Rim/GetInputText"
	Rim_SetInputText_FullMethodName     = "/cloud.lazycat.apis.localdevice.Rim/SetInputText"
	Rim_IsInputFocus_FullMethodName     = "/cloud.lazycat.apis.localdevice.Rim/IsInputFocus"
)

// RimClient is the client API for Rim service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RimClient interface {
	// 监听输入框聚焦
	ListenInputFocus(ctx context.Context, in *UserDeviceId, opts ...grpc.CallOption) (Rim_ListenInputFocusClient, error)
	// 监听输入框失焦
	ListenInputBlur(ctx context.Context, in *UserDeviceId, opts ...grpc.CallOption) (Rim_ListenInputBlurClient, error)
	// 获取当前聚焦输入框文本
	GetInputText(ctx context.Context, in *UserDeviceId, opts ...grpc.CallOption) (*InputContentReply, error)
	// 设置当前聚焦输入框文本
	SetInputText(ctx context.Context, in *SetInputTextRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 获取当前是否有聚焦输入框
	IsInputFocus(ctx context.Context, in *UserDeviceId, opts ...grpc.CallOption) (*IsInputFocusResponse, error)
}

type rimClient struct {
	cc grpc.ClientConnInterface
}

func NewRimClient(cc grpc.ClientConnInterface) RimClient {
	return &rimClient{cc}
}

func (c *rimClient) ListenInputFocus(ctx context.Context, in *UserDeviceId, opts ...grpc.CallOption) (Rim_ListenInputFocusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rim_ServiceDesc.Streams[0], Rim_ListenInputFocus_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rimListenInputFocusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rim_ListenInputFocusClient interface {
	Recv() (*InputContentReply, error)
	grpc.ClientStream
}

type rimListenInputFocusClient struct {
	grpc.ClientStream
}

func (x *rimListenInputFocusClient) Recv() (*InputContentReply, error) {
	m := new(InputContentReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rimClient) ListenInputBlur(ctx context.Context, in *UserDeviceId, opts ...grpc.CallOption) (Rim_ListenInputBlurClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rim_ServiceDesc.Streams[1], Rim_ListenInputBlur_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rimListenInputBlurClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rim_ListenInputBlurClient interface {
	Recv() (*InputContentReply, error)
	grpc.ClientStream
}

type rimListenInputBlurClient struct {
	grpc.ClientStream
}

func (x *rimListenInputBlurClient) Recv() (*InputContentReply, error) {
	m := new(InputContentReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rimClient) GetInputText(ctx context.Context, in *UserDeviceId, opts ...grpc.CallOption) (*InputContentReply, error) {
	out := new(InputContentReply)
	err := c.cc.Invoke(ctx, Rim_GetInputText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rimClient) SetInputText(ctx context.Context, in *SetInputTextRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Rim_SetInputText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rimClient) IsInputFocus(ctx context.Context, in *UserDeviceId, opts ...grpc.CallOption) (*IsInputFocusResponse, error) {
	out := new(IsInputFocusResponse)
	err := c.cc.Invoke(ctx, Rim_IsInputFocus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RimServer is the server API for Rim service.
// All implementations must embed UnimplementedRimServer
// for forward compatibility
type RimServer interface {
	// 监听输入框聚焦
	ListenInputFocus(*UserDeviceId, Rim_ListenInputFocusServer) error
	// 监听输入框失焦
	ListenInputBlur(*UserDeviceId, Rim_ListenInputBlurServer) error
	// 获取当前聚焦输入框文本
	GetInputText(context.Context, *UserDeviceId) (*InputContentReply, error)
	// 设置当前聚焦输入框文本
	SetInputText(context.Context, *SetInputTextRequest) (*emptypb.Empty, error)
	// 获取当前是否有聚焦输入框
	IsInputFocus(context.Context, *UserDeviceId) (*IsInputFocusResponse, error)
	mustEmbedUnimplementedRimServer()
}

// UnimplementedRimServer must be embedded to have forward compatible implementations.
type UnimplementedRimServer struct {
}

func (UnimplementedRimServer) ListenInputFocus(*UserDeviceId, Rim_ListenInputFocusServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenInputFocus not implemented")
}
func (UnimplementedRimServer) ListenInputBlur(*UserDeviceId, Rim_ListenInputBlurServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenInputBlur not implemented")
}
func (UnimplementedRimServer) GetInputText(context.Context, *UserDeviceId) (*InputContentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInputText not implemented")
}
func (UnimplementedRimServer) SetInputText(context.Context, *SetInputTextRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInputText not implemented")
}
func (UnimplementedRimServer) IsInputFocus(context.Context, *UserDeviceId) (*IsInputFocusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsInputFocus not implemented")
}
func (UnimplementedRimServer) mustEmbedUnimplementedRimServer() {}

// UnsafeRimServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RimServer will
// result in compilation errors.
type UnsafeRimServer interface {
	mustEmbedUnimplementedRimServer()
}

func RegisterRimServer(s grpc.ServiceRegistrar, srv RimServer) {
	s.RegisterService(&Rim_ServiceDesc, srv)
}

func _Rim_ListenInputFocus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserDeviceId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RimServer).ListenInputFocus(m, &rimListenInputFocusServer{stream})
}

type Rim_ListenInputFocusServer interface {
	Send(*InputContentReply) error
	grpc.ServerStream
}

type rimListenInputFocusServer struct {
	grpc.ServerStream
}

func (x *rimListenInputFocusServer) Send(m *InputContentReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Rim_ListenInputBlur_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserDeviceId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RimServer).ListenInputBlur(m, &rimListenInputBlurServer{stream})
}

type Rim_ListenInputBlurServer interface {
	Send(*InputContentReply) error
	grpc.ServerStream
}

type rimListenInputBlurServer struct {
	grpc.ServerStream
}

func (x *rimListenInputBlurServer) Send(m *InputContentReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Rim_GetInputText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeviceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RimServer).GetInputText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rim_GetInputText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RimServer).GetInputText(ctx, req.(*UserDeviceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rim_SetInputText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInputTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RimServer).SetInputText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rim_SetInputText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RimServer).SetInputText(ctx, req.(*SetInputTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rim_IsInputFocus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeviceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RimServer).IsInputFocus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rim_IsInputFocus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RimServer).IsInputFocus(ctx, req.(*UserDeviceId))
	}
	return interceptor(ctx, in, info, handler)
}

// Rim_ServiceDesc is the grpc.ServiceDesc for Rim service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rim_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.localdevice.Rim",
	HandlerType: (*RimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInputText",
			Handler:    _Rim_GetInputText_Handler,
		},
		{
			MethodName: "SetInputText",
			Handler:    _Rim_SetInputText_Handler,
		},
		{
			MethodName: "IsInputFocus",
			Handler:    _Rim_IsInputFocus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenInputFocus",
			Handler:       _Rim_ListenInputFocus_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenInputBlur",
			Handler:       _Rim_ListenInputBlur_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "localdevice/remote-input-method.proto",
}

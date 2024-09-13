// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: common/end_device.proto

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
	EndDeviceService_ListEndDevices_FullMethodName     = "/cloud.lazycat.apis.common.EndDeviceService/ListEndDevices"
	EndDeviceService_PairAllEndDevices__FullMethodName = "/cloud.lazycat.apis.common.EndDeviceService/PairAllEndDevices_"
	EndDeviceService_RemoveEndDevice_FullMethodName    = "/cloud.lazycat.apis.common.EndDeviceService/RemoveEndDevice"
)

// EndDeviceServiceClient is the client API for EndDeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndDeviceServiceClient interface {
	// 枚举当前登陆用户所有的设备信息
	ListEndDevices(ctx context.Context, in *ListEndDeviceRequest, opts ...grpc.CallOption) (*ListEndDeviceReply, error)
	// 将发起请求的设备与登陆用户的其他所有设备建立其虚拟网络隧道
	// 以便发起请求的faile浏览器可以绕过盒子直接访问其他节点上的设备API，比如剪贴板、文件拷贝等
	// 后端代码本身就与任意设备建立好了虚拟网络隧道，因此不需要使用此API。
	// 前端代码可以调用browseronly.proto:PairAllDeivce自动设置相关参数。
	PairAllEndDevices_(ctx context.Context, in *PairEndDeviceRequest, opts ...grpc.CallOption) (EndDeviceService_PairAllEndDevices_Client, error)
	// 移除指定uid下面的指定设备(如果设备id为空，将会移除所有的设备）
	RemoveEndDevice(ctx context.Context, in *RemoveEndDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type endDeviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEndDeviceServiceClient(cc grpc.ClientConnInterface) EndDeviceServiceClient {
	return &endDeviceServiceClient{cc}
}

func (c *endDeviceServiceClient) ListEndDevices(ctx context.Context, in *ListEndDeviceRequest, opts ...grpc.CallOption) (*ListEndDeviceReply, error) {
	out := new(ListEndDeviceReply)
	err := c.cc.Invoke(ctx, EndDeviceService_ListEndDevices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceServiceClient) PairAllEndDevices_(ctx context.Context, in *PairEndDeviceRequest, opts ...grpc.CallOption) (EndDeviceService_PairAllEndDevices_Client, error) {
	stream, err := c.cc.NewStream(ctx, &EndDeviceService_ServiceDesc.Streams[0], EndDeviceService_PairAllEndDevices__FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &endDeviceServicePairAllEndDevices_Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EndDeviceService_PairAllEndDevices_Client interface {
	Recv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type endDeviceServicePairAllEndDevices_Client struct {
	grpc.ClientStream
}

func (x *endDeviceServicePairAllEndDevices_Client) Recv() (*emptypb.Empty, error) {
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endDeviceServiceClient) RemoveEndDevice(ctx context.Context, in *RemoveEndDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EndDeviceService_RemoveEndDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndDeviceServiceServer is the server API for EndDeviceService service.
// All implementations must embed UnimplementedEndDeviceServiceServer
// for forward compatibility
type EndDeviceServiceServer interface {
	// 枚举当前登陆用户所有的设备信息
	ListEndDevices(context.Context, *ListEndDeviceRequest) (*ListEndDeviceReply, error)
	// 将发起请求的设备与登陆用户的其他所有设备建立其虚拟网络隧道
	// 以便发起请求的faile浏览器可以绕过盒子直接访问其他节点上的设备API，比如剪贴板、文件拷贝等
	// 后端代码本身就与任意设备建立好了虚拟网络隧道，因此不需要使用此API。
	// 前端代码可以调用browseronly.proto:PairAllDeivce自动设置相关参数。
	PairAllEndDevices_(*PairEndDeviceRequest, EndDeviceService_PairAllEndDevices_Server) error
	// 移除指定uid下面的指定设备(如果设备id为空，将会移除所有的设备）
	RemoveEndDevice(context.Context, *RemoveEndDeviceRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedEndDeviceServiceServer()
}

// UnimplementedEndDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEndDeviceServiceServer struct {
}

func (UnimplementedEndDeviceServiceServer) ListEndDevices(context.Context, *ListEndDeviceRequest) (*ListEndDeviceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEndDevices not implemented")
}
func (UnimplementedEndDeviceServiceServer) PairAllEndDevices_(*PairEndDeviceRequest, EndDeviceService_PairAllEndDevices_Server) error {
	return status.Errorf(codes.Unimplemented, "method PairAllEndDevices_ not implemented")
}
func (UnimplementedEndDeviceServiceServer) RemoveEndDevice(context.Context, *RemoveEndDeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEndDevice not implemented")
}
func (UnimplementedEndDeviceServiceServer) mustEmbedUnimplementedEndDeviceServiceServer() {}

// UnsafeEndDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndDeviceServiceServer will
// result in compilation errors.
type UnsafeEndDeviceServiceServer interface {
	mustEmbedUnimplementedEndDeviceServiceServer()
}

func RegisterEndDeviceServiceServer(s grpc.ServiceRegistrar, srv EndDeviceServiceServer) {
	s.RegisterService(&EndDeviceService_ServiceDesc, srv)
}

func _EndDeviceService_ListEndDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceServiceServer).ListEndDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceService_ListEndDevices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceServiceServer).ListEndDevices(ctx, req.(*ListEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceService_PairAllEndDevices__Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PairEndDeviceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EndDeviceServiceServer).PairAllEndDevices_(m, &endDeviceServicePairAllEndDevices_Server{stream})
}

type EndDeviceService_PairAllEndDevices_Server interface {
	Send(*emptypb.Empty) error
	grpc.ServerStream
}

type endDeviceServicePairAllEndDevices_Server struct {
	grpc.ServerStream
}

func (x *endDeviceServicePairAllEndDevices_Server) Send(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func _EndDeviceService_RemoveEndDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveEndDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceServiceServer).RemoveEndDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndDeviceService_RemoveEndDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceServiceServer).RemoveEndDevice(ctx, req.(*RemoveEndDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EndDeviceService_ServiceDesc is the grpc.ServiceDesc for EndDeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndDeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.common.EndDeviceService",
	HandlerType: (*EndDeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEndDevices",
			Handler:    _EndDeviceService_ListEndDevices_Handler,
		},
		{
			MethodName: "RemoveEndDevice",
			Handler:    _EndDeviceService_RemoveEndDevice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PairAllEndDevices_",
			Handler:       _EndDeviceService_PairAllEndDevices__Handler,
			ServerStreams: true,
		},
	},
	Metadata: "common/end_device.proto",
}

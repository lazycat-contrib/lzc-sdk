// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: common/peripheral_device.proto

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

// PeripheralDeviceServiceClient is the client API for PeripheralDeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeripheralDeviceServiceClient interface {
	// 列举当前盒子所连接的移动磁盘。以及对应磁盘在当前APP中的挂载情况。
	// 挂载情况应该反应当下的实际情况。
	ListRemovableDisk(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRemovableDiskReply, error)
	// 挂载/卸载特定移动磁盘的某个分区到 $APPID/lzcapp/run/mnt/media/$partion_uuid目录下
	MountFilesystem(ctx context.Context, in *MountFilesystemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UmountFilesystem(ctx context.Context, in *UmountFilesystemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type peripheralDeviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPeripheralDeviceServiceClient(cc grpc.ClientConnInterface) PeripheralDeviceServiceClient {
	return &peripheralDeviceServiceClient{cc}
}

func (c *peripheralDeviceServiceClient) ListRemovableDisk(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRemovableDiskReply, error) {
	out := new(ListRemovableDiskReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.PeripheralDeviceService/ListRemovableDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peripheralDeviceServiceClient) MountFilesystem(ctx context.Context, in *MountFilesystemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.PeripheralDeviceService/MountFilesystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peripheralDeviceServiceClient) UmountFilesystem(ctx context.Context, in *UmountFilesystemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.PeripheralDeviceService/UmountFilesystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeripheralDeviceServiceServer is the server API for PeripheralDeviceService service.
// All implementations must embed UnimplementedPeripheralDeviceServiceServer
// for forward compatibility
type PeripheralDeviceServiceServer interface {
	// 列举当前盒子所连接的移动磁盘。以及对应磁盘在当前APP中的挂载情况。
	// 挂载情况应该反应当下的实际情况。
	ListRemovableDisk(context.Context, *emptypb.Empty) (*ListRemovableDiskReply, error)
	// 挂载/卸载特定移动磁盘的某个分区到 $APPID/lzcapp/run/mnt/media/$partion_uuid目录下
	MountFilesystem(context.Context, *MountFilesystemRequest) (*emptypb.Empty, error)
	UmountFilesystem(context.Context, *UmountFilesystemRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPeripheralDeviceServiceServer()
}

// UnimplementedPeripheralDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPeripheralDeviceServiceServer struct {
}

func (UnimplementedPeripheralDeviceServiceServer) ListRemovableDisk(context.Context, *emptypb.Empty) (*ListRemovableDiskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRemovableDisk not implemented")
}
func (UnimplementedPeripheralDeviceServiceServer) MountFilesystem(context.Context, *MountFilesystemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MountFilesystem not implemented")
}
func (UnimplementedPeripheralDeviceServiceServer) UmountFilesystem(context.Context, *UmountFilesystemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UmountFilesystem not implemented")
}
func (UnimplementedPeripheralDeviceServiceServer) mustEmbedUnimplementedPeripheralDeviceServiceServer() {
}

// UnsafePeripheralDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeripheralDeviceServiceServer will
// result in compilation errors.
type UnsafePeripheralDeviceServiceServer interface {
	mustEmbedUnimplementedPeripheralDeviceServiceServer()
}

func RegisterPeripheralDeviceServiceServer(s grpc.ServiceRegistrar, srv PeripheralDeviceServiceServer) {
	s.RegisterService(&PeripheralDeviceService_ServiceDesc, srv)
}

func _PeripheralDeviceService_ListRemovableDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeripheralDeviceServiceServer).ListRemovableDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.PeripheralDeviceService/ListRemovableDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeripheralDeviceServiceServer).ListRemovableDisk(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeripheralDeviceService_MountFilesystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MountFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeripheralDeviceServiceServer).MountFilesystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.PeripheralDeviceService/MountFilesystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeripheralDeviceServiceServer).MountFilesystem(ctx, req.(*MountFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeripheralDeviceService_UmountFilesystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UmountFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeripheralDeviceServiceServer).UmountFilesystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.PeripheralDeviceService/UmountFilesystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeripheralDeviceServiceServer).UmountFilesystem(ctx, req.(*UmountFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PeripheralDeviceService_ServiceDesc is the grpc.ServiceDesc for PeripheralDeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeripheralDeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.common.PeripheralDeviceService",
	HandlerType: (*PeripheralDeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRemovableDisk",
			Handler:    _PeripheralDeviceService_ListRemovableDisk_Handler,
		},
		{
			MethodName: "MountFilesystem",
			Handler:    _PeripheralDeviceService_MountFilesystem_Handler,
		},
		{
			MethodName: "UmountFilesystem",
			Handler:    _PeripheralDeviceService_UmountFilesystem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/peripheral_device.proto",
}

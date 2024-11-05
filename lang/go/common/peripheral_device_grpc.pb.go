// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
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

const (
	PeripheralDeviceService_DeviceChanged_FullMethodName    = "/cloud.lazycat.apis.common.PeripheralDeviceService/DeviceChanged"
	PeripheralDeviceService_ListFilesystems_FullMethodName  = "/cloud.lazycat.apis.common.PeripheralDeviceService/ListFilesystems"
	PeripheralDeviceService_MountDisk_FullMethodName        = "/cloud.lazycat.apis.common.PeripheralDeviceService/MountDisk"
	PeripheralDeviceService_MountRemoteDisk_FullMethodName  = "/cloud.lazycat.apis.common.PeripheralDeviceService/MountRemoteDisk"
	PeripheralDeviceService_UmountFilesystem_FullMethodName = "/cloud.lazycat.apis.common.PeripheralDeviceService/UmountFilesystem"
	PeripheralDeviceService_MountArchive_FullMethodName     = "/cloud.lazycat.apis.common.PeripheralDeviceService/MountArchive"
	PeripheralDeviceService_PowerOffDisk_FullMethodName     = "/cloud.lazycat.apis.common.PeripheralDeviceService/PowerOffDisk"
)

// PeripheralDeviceServiceClient is the client API for PeripheralDeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeripheralDeviceServiceClient interface {
	// 如果有设备变动，就返回一下 Empty，使用方需要 ListFilesystems 来获取具体信息
	DeviceChanged(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PeripheralDeviceService_DeviceChangedClient, error)
	// 列出当前盒子已挂载 和 可以挂载但未挂载的文件系统
	ListFilesystems(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListFilesystemsReply, error)
	// 挂载/卸载特定移动磁盘的某个分区到
	// $APPID/lzcapp/run/mnt/media/$partition_uuid 目录上
	MountDisk(ctx context.Context, in *MountDiskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 挂载 Smb/NFS/WebDAV 到 $APPID/lzcapp/run/mnt/media/$uid/.remotefs/$mountpoint 目录下
	MountRemoteDisk(ctx context.Context, in *MountRemoteDiskRequest, opts ...grpc.CallOption) (*MountRemoteDiskResp, error)
	// 通过 uuid 或 mountpoint 卸载文件系统
	UmountFilesystem(ctx context.Context, in *UmountFilesystemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MountArchive(ctx context.Context, in *MountArchiveRequest, opts ...grpc.CallOption) (PeripheralDeviceService_MountArchiveClient, error)
	// 弹出外部存储设备
	PowerOffDisk(ctx context.Context, in *PowerOffDiskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type peripheralDeviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPeripheralDeviceServiceClient(cc grpc.ClientConnInterface) PeripheralDeviceServiceClient {
	return &peripheralDeviceServiceClient{cc}
}

func (c *peripheralDeviceServiceClient) DeviceChanged(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PeripheralDeviceService_DeviceChangedClient, error) {
	stream, err := c.cc.NewStream(ctx, &PeripheralDeviceService_ServiceDesc.Streams[0], PeripheralDeviceService_DeviceChanged_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &peripheralDeviceServiceDeviceChangedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PeripheralDeviceService_DeviceChangedClient interface {
	Recv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type peripheralDeviceServiceDeviceChangedClient struct {
	grpc.ClientStream
}

func (x *peripheralDeviceServiceDeviceChangedClient) Recv() (*emptypb.Empty, error) {
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *peripheralDeviceServiceClient) ListFilesystems(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListFilesystemsReply, error) {
	out := new(ListFilesystemsReply)
	err := c.cc.Invoke(ctx, PeripheralDeviceService_ListFilesystems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peripheralDeviceServiceClient) MountDisk(ctx context.Context, in *MountDiskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PeripheralDeviceService_MountDisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peripheralDeviceServiceClient) MountRemoteDisk(ctx context.Context, in *MountRemoteDiskRequest, opts ...grpc.CallOption) (*MountRemoteDiskResp, error) {
	out := new(MountRemoteDiskResp)
	err := c.cc.Invoke(ctx, PeripheralDeviceService_MountRemoteDisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peripheralDeviceServiceClient) UmountFilesystem(ctx context.Context, in *UmountFilesystemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PeripheralDeviceService_UmountFilesystem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peripheralDeviceServiceClient) MountArchive(ctx context.Context, in *MountArchiveRequest, opts ...grpc.CallOption) (PeripheralDeviceService_MountArchiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &PeripheralDeviceService_ServiceDesc.Streams[1], PeripheralDeviceService_MountArchive_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &peripheralDeviceServiceMountArchiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PeripheralDeviceService_MountArchiveClient interface {
	Recv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type peripheralDeviceServiceMountArchiveClient struct {
	grpc.ClientStream
}

func (x *peripheralDeviceServiceMountArchiveClient) Recv() (*emptypb.Empty, error) {
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *peripheralDeviceServiceClient) PowerOffDisk(ctx context.Context, in *PowerOffDiskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PeripheralDeviceService_PowerOffDisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeripheralDeviceServiceServer is the server API for PeripheralDeviceService service.
// All implementations must embed UnimplementedPeripheralDeviceServiceServer
// for forward compatibility
type PeripheralDeviceServiceServer interface {
	// 如果有设备变动，就返回一下 Empty，使用方需要 ListFilesystems 来获取具体信息
	DeviceChanged(*emptypb.Empty, PeripheralDeviceService_DeviceChangedServer) error
	// 列出当前盒子已挂载 和 可以挂载但未挂载的文件系统
	ListFilesystems(context.Context, *emptypb.Empty) (*ListFilesystemsReply, error)
	// 挂载/卸载特定移动磁盘的某个分区到
	// $APPID/lzcapp/run/mnt/media/$partition_uuid 目录上
	MountDisk(context.Context, *MountDiskRequest) (*emptypb.Empty, error)
	// 挂载 Smb/NFS/WebDAV 到 $APPID/lzcapp/run/mnt/media/$uid/.remotefs/$mountpoint 目录下
	MountRemoteDisk(context.Context, *MountRemoteDiskRequest) (*MountRemoteDiskResp, error)
	// 通过 uuid 或 mountpoint 卸载文件系统
	UmountFilesystem(context.Context, *UmountFilesystemRequest) (*emptypb.Empty, error)
	MountArchive(*MountArchiveRequest, PeripheralDeviceService_MountArchiveServer) error
	// 弹出外部存储设备
	PowerOffDisk(context.Context, *PowerOffDiskRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPeripheralDeviceServiceServer()
}

// UnimplementedPeripheralDeviceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPeripheralDeviceServiceServer struct {
}

func (UnimplementedPeripheralDeviceServiceServer) DeviceChanged(*emptypb.Empty, PeripheralDeviceService_DeviceChangedServer) error {
	return status.Errorf(codes.Unimplemented, "method DeviceChanged not implemented")
}
func (UnimplementedPeripheralDeviceServiceServer) ListFilesystems(context.Context, *emptypb.Empty) (*ListFilesystemsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFilesystems not implemented")
}
func (UnimplementedPeripheralDeviceServiceServer) MountDisk(context.Context, *MountDiskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MountDisk not implemented")
}
func (UnimplementedPeripheralDeviceServiceServer) MountRemoteDisk(context.Context, *MountRemoteDiskRequest) (*MountRemoteDiskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MountRemoteDisk not implemented")
}
func (UnimplementedPeripheralDeviceServiceServer) UmountFilesystem(context.Context, *UmountFilesystemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UmountFilesystem not implemented")
}
func (UnimplementedPeripheralDeviceServiceServer) MountArchive(*MountArchiveRequest, PeripheralDeviceService_MountArchiveServer) error {
	return status.Errorf(codes.Unimplemented, "method MountArchive not implemented")
}
func (UnimplementedPeripheralDeviceServiceServer) PowerOffDisk(context.Context, *PowerOffDiskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PowerOffDisk not implemented")
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

func _PeripheralDeviceService_DeviceChanged_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PeripheralDeviceServiceServer).DeviceChanged(m, &peripheralDeviceServiceDeviceChangedServer{stream})
}

type PeripheralDeviceService_DeviceChangedServer interface {
	Send(*emptypb.Empty) error
	grpc.ServerStream
}

type peripheralDeviceServiceDeviceChangedServer struct {
	grpc.ServerStream
}

func (x *peripheralDeviceServiceDeviceChangedServer) Send(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func _PeripheralDeviceService_ListFilesystems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeripheralDeviceServiceServer).ListFilesystems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeripheralDeviceService_ListFilesystems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeripheralDeviceServiceServer).ListFilesystems(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeripheralDeviceService_MountDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MountDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeripheralDeviceServiceServer).MountDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeripheralDeviceService_MountDisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeripheralDeviceServiceServer).MountDisk(ctx, req.(*MountDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeripheralDeviceService_MountRemoteDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MountRemoteDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeripheralDeviceServiceServer).MountRemoteDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeripheralDeviceService_MountRemoteDisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeripheralDeviceServiceServer).MountRemoteDisk(ctx, req.(*MountRemoteDiskRequest))
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
		FullMethod: PeripheralDeviceService_UmountFilesystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeripheralDeviceServiceServer).UmountFilesystem(ctx, req.(*UmountFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeripheralDeviceService_MountArchive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MountArchiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PeripheralDeviceServiceServer).MountArchive(m, &peripheralDeviceServiceMountArchiveServer{stream})
}

type PeripheralDeviceService_MountArchiveServer interface {
	Send(*emptypb.Empty) error
	grpc.ServerStream
}

type peripheralDeviceServiceMountArchiveServer struct {
	grpc.ServerStream
}

func (x *peripheralDeviceServiceMountArchiveServer) Send(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func _PeripheralDeviceService_PowerOffDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PowerOffDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeripheralDeviceServiceServer).PowerOffDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeripheralDeviceService_PowerOffDisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeripheralDeviceServiceServer).PowerOffDisk(ctx, req.(*PowerOffDiskRequest))
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
			MethodName: "ListFilesystems",
			Handler:    _PeripheralDeviceService_ListFilesystems_Handler,
		},
		{
			MethodName: "MountDisk",
			Handler:    _PeripheralDeviceService_MountDisk_Handler,
		},
		{
			MethodName: "MountRemoteDisk",
			Handler:    _PeripheralDeviceService_MountRemoteDisk_Handler,
		},
		{
			MethodName: "UmountFilesystem",
			Handler:    _PeripheralDeviceService_UmountFilesystem_Handler,
		},
		{
			MethodName: "PowerOffDisk",
			Handler:    _PeripheralDeviceService_PowerOffDisk_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeviceChanged",
			Handler:       _PeripheralDeviceService_DeviceChanged_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MountArchive",
			Handler:       _PeripheralDeviceService_MountArchive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "common/peripheral_device.proto",
}

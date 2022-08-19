// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: sys/package_manager.proto

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

// PackageManagerClient is the client API for PackageManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackageManagerClient interface {
	// 根据 URL 和 校验码（可选），安装应用
	Install(ctx context.Context, in *InstallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 根据 AppId 卸载应用
	Uninstall(ctx context.Context, in *UninstallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 查询用户安装的特定应用的详情
	QueryApplication(ctx context.Context, in *QueryApplicationRequest, opts ...grpc.CallOption) (*QueryApplicationResponse, error)
	// 获取应用占用的存储空间详情
	QueryAppStorageUsage(ctx context.Context, in *QueryAppStorageUsageRequest, opts ...grpc.CallOption) (*AppStorageUsage, error)
	// 设置某个用户是否可以安装应用
	SetUserPermissions(ctx context.Context, in *SetUserPermissionsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 暂停下载特定应用
	PauseAppDownload(ctx context.Context, in *Appid, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type packageManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageManagerClient(cc grpc.ClientConnInterface) PackageManagerClient {
	return &packageManagerClient{cc}
}

func (c *packageManagerClient) Install(ctx context.Context, in *InstallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.PackageManager/Install", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageManagerClient) Uninstall(ctx context.Context, in *UninstallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.PackageManager/Uninstall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageManagerClient) QueryApplication(ctx context.Context, in *QueryApplicationRequest, opts ...grpc.CallOption) (*QueryApplicationResponse, error) {
	out := new(QueryApplicationResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.PackageManager/QueryApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageManagerClient) QueryAppStorageUsage(ctx context.Context, in *QueryAppStorageUsageRequest, opts ...grpc.CallOption) (*AppStorageUsage, error) {
	out := new(AppStorageUsage)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.PackageManager/QueryAppStorageUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageManagerClient) SetUserPermissions(ctx context.Context, in *SetUserPermissionsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.PackageManager/SetUserPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageManagerClient) PauseAppDownload(ctx context.Context, in *Appid, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.PackageManager/PauseAppDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackageManagerServer is the server API for PackageManager service.
// All implementations must embed UnimplementedPackageManagerServer
// for forward compatibility
type PackageManagerServer interface {
	// 根据 URL 和 校验码（可选），安装应用
	Install(context.Context, *InstallRequest) (*emptypb.Empty, error)
	// 根据 AppId 卸载应用
	Uninstall(context.Context, *UninstallRequest) (*emptypb.Empty, error)
	// 查询用户安装的特定应用的详情
	QueryApplication(context.Context, *QueryApplicationRequest) (*QueryApplicationResponse, error)
	// 获取应用占用的存储空间详情
	QueryAppStorageUsage(context.Context, *QueryAppStorageUsageRequest) (*AppStorageUsage, error)
	// 设置某个用户是否可以安装应用
	SetUserPermissions(context.Context, *SetUserPermissionsRequest) (*emptypb.Empty, error)
	// 暂停下载特定应用
	PauseAppDownload(context.Context, *Appid) (*emptypb.Empty, error)
	mustEmbedUnimplementedPackageManagerServer()
}

// UnimplementedPackageManagerServer must be embedded to have forward compatible implementations.
type UnimplementedPackageManagerServer struct {
}

func (UnimplementedPackageManagerServer) Install(context.Context, *InstallRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedPackageManagerServer) Uninstall(context.Context, *UninstallRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Uninstall not implemented")
}
func (UnimplementedPackageManagerServer) QueryApplication(context.Context, *QueryApplicationRequest) (*QueryApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryApplication not implemented")
}
func (UnimplementedPackageManagerServer) QueryAppStorageUsage(context.Context, *QueryAppStorageUsageRequest) (*AppStorageUsage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAppStorageUsage not implemented")
}
func (UnimplementedPackageManagerServer) SetUserPermissions(context.Context, *SetUserPermissionsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserPermissions not implemented")
}
func (UnimplementedPackageManagerServer) PauseAppDownload(context.Context, *Appid) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseAppDownload not implemented")
}
func (UnimplementedPackageManagerServer) mustEmbedUnimplementedPackageManagerServer() {}

// UnsafePackageManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackageManagerServer will
// result in compilation errors.
type UnsafePackageManagerServer interface {
	mustEmbedUnimplementedPackageManagerServer()
}

func RegisterPackageManagerServer(s grpc.ServiceRegistrar, srv PackageManagerServer) {
	s.RegisterService(&PackageManager_ServiceDesc, srv)
}

func _PackageManager_Install_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageManagerServer).Install(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.PackageManager/Install",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageManagerServer).Install(ctx, req.(*InstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageManager_Uninstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UninstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageManagerServer).Uninstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.PackageManager/Uninstall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageManagerServer).Uninstall(ctx, req.(*UninstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageManager_QueryApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageManagerServer).QueryApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.PackageManager/QueryApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageManagerServer).QueryApplication(ctx, req.(*QueryApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageManager_QueryAppStorageUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAppStorageUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageManagerServer).QueryAppStorageUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.PackageManager/QueryAppStorageUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageManagerServer).QueryAppStorageUsage(ctx, req.(*QueryAppStorageUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageManager_SetUserPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageManagerServer).SetUserPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.PackageManager/SetUserPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageManagerServer).SetUserPermissions(ctx, req.(*SetUserPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageManager_PauseAppDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Appid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageManagerServer).PauseAppDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.PackageManager/PauseAppDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageManagerServer).PauseAppDownload(ctx, req.(*Appid))
	}
	return interceptor(ctx, in, info, handler)
}

// PackageManager_ServiceDesc is the grpc.ServiceDesc for PackageManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PackageManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.sys.PackageManager",
	HandlerType: (*PackageManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Install",
			Handler:    _PackageManager_Install_Handler,
		},
		{
			MethodName: "Uninstall",
			Handler:    _PackageManager_Uninstall_Handler,
		},
		{
			MethodName: "QueryApplication",
			Handler:    _PackageManager_QueryApplication_Handler,
		},
		{
			MethodName: "QueryAppStorageUsage",
			Handler:    _PackageManager_QueryAppStorageUsage_Handler,
		},
		{
			MethodName: "SetUserPermissions",
			Handler:    _PackageManager_SetUserPermissions_Handler,
		},
		{
			MethodName: "PauseAppDownload",
			Handler:    _PackageManager_PauseAppDownload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys/package_manager.proto",
}

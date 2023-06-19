// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: common/apk_shell.proto

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

// ApkShellClient is the client API for ApkShell service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApkShellClient interface {
	// 基于模板，根据传入的参数构建 APK
	// 同一时间最多只能有一个构建任务在执行，并发的请求会阻塞
	// 此方法不负责缓存生成的 APK，每次请求都会重新构建
	BuildApk(ctx context.Context, in *BuildApkRequest, opts ...grpc.CallOption) (*BuildApkResponse, error)
	GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type apkShellClient struct {
	cc grpc.ClientConnInterface
}

func NewApkShellClient(cc grpc.ClientConnInterface) ApkShellClient {
	return &apkShellClient{cc}
}

func (c *apkShellClient) BuildApk(ctx context.Context, in *BuildApkRequest, opts ...grpc.CallOption) (*BuildApkResponse, error) {
	out := new(BuildApkResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.ApkShell/BuildApk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apkShellClient) GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.ApkShell/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApkShellServer is the server API for ApkShell service.
// All implementations must embed UnimplementedApkShellServer
// for forward compatibility
type ApkShellServer interface {
	// 基于模板，根据传入的参数构建 APK
	// 同一时间最多只能有一个构建任务在执行，并发的请求会阻塞
	// 此方法不负责缓存生成的 APK，每次请求都会重新构建
	BuildApk(context.Context, *BuildApkRequest) (*BuildApkResponse, error)
	GetVersion(context.Context, *emptypb.Empty) (*GetVersionResponse, error)
	mustEmbedUnimplementedApkShellServer()
}

// UnimplementedApkShellServer must be embedded to have forward compatible implementations.
type UnimplementedApkShellServer struct {
}

func (UnimplementedApkShellServer) BuildApk(context.Context, *BuildApkRequest) (*BuildApkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildApk not implemented")
}
func (UnimplementedApkShellServer) GetVersion(context.Context, *emptypb.Empty) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedApkShellServer) mustEmbedUnimplementedApkShellServer() {}

// UnsafeApkShellServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApkShellServer will
// result in compilation errors.
type UnsafeApkShellServer interface {
	mustEmbedUnimplementedApkShellServer()
}

func RegisterApkShellServer(s grpc.ServiceRegistrar, srv ApkShellServer) {
	s.RegisterService(&ApkShell_ServiceDesc, srv)
}

func _ApkShell_BuildApk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildApkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApkShellServer).BuildApk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.ApkShell/BuildApk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApkShellServer).BuildApk(ctx, req.(*BuildApkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApkShell_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApkShellServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.ApkShell/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApkShellServer).GetVersion(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ApkShell_ServiceDesc is the grpc.ServiceDesc for ApkShell service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApkShell_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.common.ApkShell",
	HandlerType: (*ApkShellServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildApk",
			Handler:    _ApkShell_BuildApk_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ApkShell_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/apk_shell.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: localdevice/permission.proto

package localdevice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PermissionManagerClient is the client API for PermissionManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionManagerClient interface {
	// 检测权限
	GetPermission(ctx context.Context, in *PermissionRequest, opts ...grpc.CallOption) (*PermissionReply, error)
	// 申请权限（会弹出对话框让用户决定是否同意）
	RequestPermission(ctx context.Context, in *PermissionRequest, opts ...grpc.CallOption) (*PermissionReply, error)
}

type permissionManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionManagerClient(cc grpc.ClientConnInterface) PermissionManagerClient {
	return &permissionManagerClient{cc}
}

func (c *permissionManagerClient) GetPermission(ctx context.Context, in *PermissionRequest, opts ...grpc.CallOption) (*PermissionReply, error) {
	out := new(PermissionReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.localdevice.PermissionManager/GetPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionManagerClient) RequestPermission(ctx context.Context, in *PermissionRequest, opts ...grpc.CallOption) (*PermissionReply, error) {
	out := new(PermissionReply)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.localdevice.PermissionManager/RequestPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionManagerServer is the server API for PermissionManager service.
// All implementations must embed UnimplementedPermissionManagerServer
// for forward compatibility
type PermissionManagerServer interface {
	// 检测权限
	GetPermission(context.Context, *PermissionRequest) (*PermissionReply, error)
	// 申请权限（会弹出对话框让用户决定是否同意）
	RequestPermission(context.Context, *PermissionRequest) (*PermissionReply, error)
	mustEmbedUnimplementedPermissionManagerServer()
}

// UnimplementedPermissionManagerServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionManagerServer struct {
}

func (UnimplementedPermissionManagerServer) GetPermission(context.Context, *PermissionRequest) (*PermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermission not implemented")
}
func (UnimplementedPermissionManagerServer) RequestPermission(context.Context, *PermissionRequest) (*PermissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPermission not implemented")
}
func (UnimplementedPermissionManagerServer) mustEmbedUnimplementedPermissionManagerServer() {}

// UnsafePermissionManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionManagerServer will
// result in compilation errors.
type UnsafePermissionManagerServer interface {
	mustEmbedUnimplementedPermissionManagerServer()
}

func RegisterPermissionManagerServer(s grpc.ServiceRegistrar, srv PermissionManagerServer) {
	s.RegisterService(&PermissionManager_ServiceDesc, srv)
}

func _PermissionManager_GetPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionManagerServer).GetPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.localdevice.PermissionManager/GetPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionManagerServer).GetPermission(ctx, req.(*PermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionManager_RequestPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionManagerServer).RequestPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.localdevice.PermissionManager/RequestPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionManagerServer).RequestPermission(ctx, req.(*PermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionManager_ServiceDesc is the grpc.ServiceDesc for PermissionManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.localdevice.PermissionManager",
	HandlerType: (*PermissionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPermission",
			Handler:    _PermissionManager_GetPermission_Handler,
		},
		{
			MethodName: "RequestPermission",
			Handler:    _PermissionManager_RequestPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "localdevice/permission.proto",
}

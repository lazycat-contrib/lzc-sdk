// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: localdevice/config.proto

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

const (
	UserConfig_GetUserConfig_FullMethodName = "/cloud.lazycat.apis.localdevice.UserConfig/GetUserConfig"
	UserConfig_SetUserConfig_FullMethodName = "/cloud.lazycat.apis.localdevice.UserConfig/SetUserConfig"
)

// UserConfigClient is the client API for UserConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserConfigClient interface {
	GetUserConfig(ctx context.Context, in *GetUserConfigRequest, opts ...grpc.CallOption) (*GetUserConfigResponse, error)
	SetUserConfig(ctx context.Context, in *SetUserConfigRequest, opts ...grpc.CallOption) (*SetUserConfigResponse, error)
}

type userConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewUserConfigClient(cc grpc.ClientConnInterface) UserConfigClient {
	return &userConfigClient{cc}
}

func (c *userConfigClient) GetUserConfig(ctx context.Context, in *GetUserConfigRequest, opts ...grpc.CallOption) (*GetUserConfigResponse, error) {
	out := new(GetUserConfigResponse)
	err := c.cc.Invoke(ctx, UserConfig_GetUserConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userConfigClient) SetUserConfig(ctx context.Context, in *SetUserConfigRequest, opts ...grpc.CallOption) (*SetUserConfigResponse, error) {
	out := new(SetUserConfigResponse)
	err := c.cc.Invoke(ctx, UserConfig_SetUserConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserConfigServer is the server API for UserConfig service.
// All implementations must embed UnimplementedUserConfigServer
// for forward compatibility
type UserConfigServer interface {
	GetUserConfig(context.Context, *GetUserConfigRequest) (*GetUserConfigResponse, error)
	SetUserConfig(context.Context, *SetUserConfigRequest) (*SetUserConfigResponse, error)
	mustEmbedUnimplementedUserConfigServer()
}

// UnimplementedUserConfigServer must be embedded to have forward compatible implementations.
type UnimplementedUserConfigServer struct {
}

func (UnimplementedUserConfigServer) GetUserConfig(context.Context, *GetUserConfigRequest) (*GetUserConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserConfig not implemented")
}
func (UnimplementedUserConfigServer) SetUserConfig(context.Context, *SetUserConfigRequest) (*SetUserConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserConfig not implemented")
}
func (UnimplementedUserConfigServer) mustEmbedUnimplementedUserConfigServer() {}

// UnsafeUserConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserConfigServer will
// result in compilation errors.
type UnsafeUserConfigServer interface {
	mustEmbedUnimplementedUserConfigServer()
}

func RegisterUserConfigServer(s grpc.ServiceRegistrar, srv UserConfigServer) {
	s.RegisterService(&UserConfig_ServiceDesc, srv)
}

func _UserConfig_GetUserConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConfigServer).GetUserConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserConfig_GetUserConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConfigServer).GetUserConfig(ctx, req.(*GetUserConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserConfig_SetUserConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserConfigServer).SetUserConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserConfig_SetUserConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserConfigServer).SetUserConfig(ctx, req.(*SetUserConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserConfig_ServiceDesc is the grpc.ServiceDesc for UserConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.localdevice.UserConfig",
	HandlerType: (*UserConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserConfig",
			Handler:    _UserConfig_GetUserConfig_Handler,
		},
		{
			MethodName: "SetUserConfig",
			Handler:    _UserConfig_SetUserConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "localdevice/config.proto",
}

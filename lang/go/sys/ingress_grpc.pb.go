// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: sys/ingress.proto

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

// IngressServiceClient is the client API for IngressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngressServiceClient interface {
	// 允许一个用户对指定 app 的访问
	AllowAdd(ctx context.Context, in *IngressAllowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 阻止一个用户对指定 app 的访问权限
	AllowDel(ctx context.Context, in *IngressAllowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 阻止一个用户对所有 app 的访问权限
	AllowDelAll(ctx context.Context, in *IngressAllowDelAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 查询一个用户是否被允许访问指定 app
	AllowGet(ctx context.Context, in *IngressAllowGetRequest, opts ...grpc.CallOption) (*IngressAllowGetResponse, error)
	// 列出一个用户的所有被允许访问的 app 列表
	AllowList(ctx context.Context, in *IngressAllowListRequest, opts ...grpc.CallOption) (*IngressAllowListResponse, error)
	// 清空白名单
	AllowClear(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 启用/禁用白名单
	AllowManage(ctx context.Context, in *IngressAllowManageRequest, opts ...grpc.CallOption) (*IngressAllowManageResponse, error)
	// 获取指定 app 最后一次被访问的时刻
	GetAppLastAccessTime(ctx context.Context, in *IngressAppLastAccessTimeRequest, opts ...grpc.CallOption) (*IngressAppLastAccessTimeResponse, error)
}

type ingressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIngressServiceClient(cc grpc.ClientConnInterface) IngressServiceClient {
	return &ingressServiceClient{cc}
}

func (c *ingressServiceClient) AllowAdd(ctx context.Context, in *IngressAllowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/AllowAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowDel(ctx context.Context, in *IngressAllowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/AllowDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowDelAll(ctx context.Context, in *IngressAllowDelAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/AllowDelAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowGet(ctx context.Context, in *IngressAllowGetRequest, opts ...grpc.CallOption) (*IngressAllowGetResponse, error) {
	out := new(IngressAllowGetResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/AllowGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowList(ctx context.Context, in *IngressAllowListRequest, opts ...grpc.CallOption) (*IngressAllowListResponse, error) {
	out := new(IngressAllowListResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/AllowList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowClear(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/AllowClear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowManage(ctx context.Context, in *IngressAllowManageRequest, opts ...grpc.CallOption) (*IngressAllowManageResponse, error) {
	out := new(IngressAllowManageResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/AllowManage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) GetAppLastAccessTime(ctx context.Context, in *IngressAppLastAccessTimeRequest, opts ...grpc.CallOption) (*IngressAppLastAccessTimeResponse, error) {
	out := new(IngressAppLastAccessTimeResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/GetAppLastAccessTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngressServiceServer is the server API for IngressService service.
// All implementations must embed UnimplementedIngressServiceServer
// for forward compatibility
type IngressServiceServer interface {
	// 允许一个用户对指定 app 的访问
	AllowAdd(context.Context, *IngressAllowRequest) (*emptypb.Empty, error)
	// 阻止一个用户对指定 app 的访问权限
	AllowDel(context.Context, *IngressAllowRequest) (*emptypb.Empty, error)
	// 阻止一个用户对所有 app 的访问权限
	AllowDelAll(context.Context, *IngressAllowDelAllRequest) (*emptypb.Empty, error)
	// 查询一个用户是否被允许访问指定 app
	AllowGet(context.Context, *IngressAllowGetRequest) (*IngressAllowGetResponse, error)
	// 列出一个用户的所有被允许访问的 app 列表
	AllowList(context.Context, *IngressAllowListRequest) (*IngressAllowListResponse, error)
	// 清空白名单
	AllowClear(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// 启用/禁用白名单
	AllowManage(context.Context, *IngressAllowManageRequest) (*IngressAllowManageResponse, error)
	// 获取指定 app 最后一次被访问的时刻
	GetAppLastAccessTime(context.Context, *IngressAppLastAccessTimeRequest) (*IngressAppLastAccessTimeResponse, error)
	mustEmbedUnimplementedIngressServiceServer()
}

// UnimplementedIngressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIngressServiceServer struct {
}

func (UnimplementedIngressServiceServer) AllowAdd(context.Context, *IngressAllowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowAdd not implemented")
}
func (UnimplementedIngressServiceServer) AllowDel(context.Context, *IngressAllowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowDel not implemented")
}
func (UnimplementedIngressServiceServer) AllowDelAll(context.Context, *IngressAllowDelAllRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowDelAll not implemented")
}
func (UnimplementedIngressServiceServer) AllowGet(context.Context, *IngressAllowGetRequest) (*IngressAllowGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowGet not implemented")
}
func (UnimplementedIngressServiceServer) AllowList(context.Context, *IngressAllowListRequest) (*IngressAllowListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowList not implemented")
}
func (UnimplementedIngressServiceServer) AllowClear(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowClear not implemented")
}
func (UnimplementedIngressServiceServer) AllowManage(context.Context, *IngressAllowManageRequest) (*IngressAllowManageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowManage not implemented")
}
func (UnimplementedIngressServiceServer) GetAppLastAccessTime(context.Context, *IngressAppLastAccessTimeRequest) (*IngressAppLastAccessTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppLastAccessTime not implemented")
}
func (UnimplementedIngressServiceServer) mustEmbedUnimplementedIngressServiceServer() {}

// UnsafeIngressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngressServiceServer will
// result in compilation errors.
type UnsafeIngressServiceServer interface {
	mustEmbedUnimplementedIngressServiceServer()
}

func RegisterIngressServiceServer(s grpc.ServiceRegistrar, srv IngressServiceServer) {
	s.RegisterService(&IngressService_ServiceDesc, srv)
}

func _IngressService_AllowAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressAllowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).AllowAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/AllowAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).AllowAdd(ctx, req.(*IngressAllowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_AllowDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressAllowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).AllowDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/AllowDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).AllowDel(ctx, req.(*IngressAllowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_AllowDelAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressAllowDelAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).AllowDelAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/AllowDelAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).AllowDelAll(ctx, req.(*IngressAllowDelAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_AllowGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressAllowGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).AllowGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/AllowGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).AllowGet(ctx, req.(*IngressAllowGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_AllowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressAllowListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).AllowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/AllowList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).AllowList(ctx, req.(*IngressAllowListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_AllowClear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).AllowClear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/AllowClear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).AllowClear(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_AllowManage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressAllowManageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).AllowManage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/AllowManage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).AllowManage(ctx, req.(*IngressAllowManageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_GetAppLastAccessTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressAppLastAccessTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).GetAppLastAccessTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/GetAppLastAccessTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).GetAppLastAccessTime(ctx, req.(*IngressAppLastAccessTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IngressService_ServiceDesc is the grpc.ServiceDesc for IngressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IngressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.sys.IngressService",
	HandlerType: (*IngressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllowAdd",
			Handler:    _IngressService_AllowAdd_Handler,
		},
		{
			MethodName: "AllowDel",
			Handler:    _IngressService_AllowDel_Handler,
		},
		{
			MethodName: "AllowDelAll",
			Handler:    _IngressService_AllowDelAll_Handler,
		},
		{
			MethodName: "AllowGet",
			Handler:    _IngressService_AllowGet_Handler,
		},
		{
			MethodName: "AllowList",
			Handler:    _IngressService_AllowList_Handler,
		},
		{
			MethodName: "AllowClear",
			Handler:    _IngressService_AllowClear_Handler,
		},
		{
			MethodName: "AllowManage",
			Handler:    _IngressService_AllowManage_Handler,
		},
		{
			MethodName: "GetAppLastAccessTime",
			Handler:    _IngressService_GetAppLastAccessTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys/ingress.proto",
}

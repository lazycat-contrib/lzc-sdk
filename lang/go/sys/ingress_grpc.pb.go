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
	// 阻止一个用户对指定 app 的访问
	BlockAdd(ctx context.Context, in *IngressBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 恢复一个用户对指定 app 的访问权限
	BlockDel(ctx context.Context, in *IngressBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 恢复一个用户对所有 app 的访问权限
	BlockDelAll(ctx context.Context, in *IngressBlockDelAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 查询一个用户是否被阻止访问指定 app
	BlockGet(ctx context.Context, in *IngressBlockGetRequest, opts ...grpc.CallOption) (*IngressBlockGetResponse, error)
	// 列出一个用户的所有被阻止访问的 app 列表
	BlockList(ctx context.Context, in *IngressBlockListRequest, opts ...grpc.CallOption) (*IngressBlockListResponse, error)
	// 清空黑名单
	BlockClear(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 获取指定 app 最后一次被访问的时刻
	GetAppLastAccessTime(ctx context.Context, in *IngressAppLastAccessTimeRequest, opts ...grpc.CallOption) (*IngressAppLastAccessTimeResponse, error)
}

type ingressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIngressServiceClient(cc grpc.ClientConnInterface) IngressServiceClient {
	return &ingressServiceClient{cc}
}

func (c *ingressServiceClient) BlockAdd(ctx context.Context, in *IngressBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/BlockAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) BlockDel(ctx context.Context, in *IngressBlockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/BlockDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) BlockDelAll(ctx context.Context, in *IngressBlockDelAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/BlockDelAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) BlockGet(ctx context.Context, in *IngressBlockGetRequest, opts ...grpc.CallOption) (*IngressBlockGetResponse, error) {
	out := new(IngressBlockGetResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/BlockGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) BlockList(ctx context.Context, in *IngressBlockListRequest, opts ...grpc.CallOption) (*IngressBlockListResponse, error) {
	out := new(IngressBlockListResponse)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/BlockList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) BlockClear(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.sys.IngressService/BlockClear", in, out, opts...)
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
	// 阻止一个用户对指定 app 的访问
	BlockAdd(context.Context, *IngressBlockRequest) (*emptypb.Empty, error)
	// 恢复一个用户对指定 app 的访问权限
	BlockDel(context.Context, *IngressBlockRequest) (*emptypb.Empty, error)
	// 恢复一个用户对所有 app 的访问权限
	BlockDelAll(context.Context, *IngressBlockDelAllRequest) (*emptypb.Empty, error)
	// 查询一个用户是否被阻止访问指定 app
	BlockGet(context.Context, *IngressBlockGetRequest) (*IngressBlockGetResponse, error)
	// 列出一个用户的所有被阻止访问的 app 列表
	BlockList(context.Context, *IngressBlockListRequest) (*IngressBlockListResponse, error)
	// 清空黑名单
	BlockClear(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// 获取指定 app 最后一次被访问的时刻
	GetAppLastAccessTime(context.Context, *IngressAppLastAccessTimeRequest) (*IngressAppLastAccessTimeResponse, error)
	mustEmbedUnimplementedIngressServiceServer()
}

// UnimplementedIngressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIngressServiceServer struct {
}

func (UnimplementedIngressServiceServer) BlockAdd(context.Context, *IngressBlockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockAdd not implemented")
}
func (UnimplementedIngressServiceServer) BlockDel(context.Context, *IngressBlockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockDel not implemented")
}
func (UnimplementedIngressServiceServer) BlockDelAll(context.Context, *IngressBlockDelAllRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockDelAll not implemented")
}
func (UnimplementedIngressServiceServer) BlockGet(context.Context, *IngressBlockGetRequest) (*IngressBlockGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockGet not implemented")
}
func (UnimplementedIngressServiceServer) BlockList(context.Context, *IngressBlockListRequest) (*IngressBlockListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockList not implemented")
}
func (UnimplementedIngressServiceServer) BlockClear(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockClear not implemented")
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

func _IngressService_BlockAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).BlockAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/BlockAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).BlockAdd(ctx, req.(*IngressBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_BlockDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).BlockDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/BlockDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).BlockDel(ctx, req.(*IngressBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_BlockDelAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressBlockDelAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).BlockDelAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/BlockDelAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).BlockDelAll(ctx, req.(*IngressBlockDelAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_BlockGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressBlockGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).BlockGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/BlockGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).BlockGet(ctx, req.(*IngressBlockGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_BlockList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressBlockListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).BlockList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/BlockList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).BlockList(ctx, req.(*IngressBlockListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IngressService_BlockClear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServiceServer).BlockClear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.sys.IngressService/BlockClear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServiceServer).BlockClear(ctx, req.(*emptypb.Empty))
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
			MethodName: "BlockAdd",
			Handler:    _IngressService_BlockAdd_Handler,
		},
		{
			MethodName: "BlockDel",
			Handler:    _IngressService_BlockDel_Handler,
		},
		{
			MethodName: "BlockDelAll",
			Handler:    _IngressService_BlockDelAll_Handler,
		},
		{
			MethodName: "BlockGet",
			Handler:    _IngressService_BlockGet_Handler,
		},
		{
			MethodName: "BlockList",
			Handler:    _IngressService_BlockList_Handler,
		},
		{
			MethodName: "BlockClear",
			Handler:    _IngressService_BlockClear_Handler,
		},
		{
			MethodName: "GetAppLastAccessTime",
			Handler:    _IngressService_GetAppLastAccessTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys/ingress.proto",
}

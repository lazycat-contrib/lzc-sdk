// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
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

const (
	AccessControlerService_SetAppAccessPolicy_FullMethodName   = "/cloud.lazycat.apis.sys.AccessControlerService/SetAppAccessPolicy"
	AccessControlerService_QueryAppAccessPolicy_FullMethodName = "/cloud.lazycat.apis.sys.AccessControlerService/QueryAppAccessPolicy"
	AccessControlerService_GetAppLastAccessTime_FullMethodName = "/cloud.lazycat.apis.sys.AccessControlerService/GetAppLastAccessTime"
)

// AccessControlerServiceClient is the client API for AccessControlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessControlerServiceClient interface {
	SetAppAccessPolicy(ctx context.Context, in *AppAccessPolicyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	QueryAppAccessPolicy(ctx context.Context, in *AppAccessPolicyRequest, opts ...grpc.CallOption) (*AppAccessPolicy, error)
	GetAppLastAccessTime(ctx context.Context, in *IngressAppLastAccessTimeRequest, opts ...grpc.CallOption) (*IngressAppLastAccessTimeResponse, error)
}

type accessControlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessControlerServiceClient(cc grpc.ClientConnInterface) AccessControlerServiceClient {
	return &accessControlerServiceClient{cc}
}

func (c *accessControlerServiceClient) SetAppAccessPolicy(ctx context.Context, in *AppAccessPolicyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AccessControlerService_SetAppAccessPolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlerServiceClient) QueryAppAccessPolicy(ctx context.Context, in *AppAccessPolicyRequest, opts ...grpc.CallOption) (*AppAccessPolicy, error) {
	out := new(AppAccessPolicy)
	err := c.cc.Invoke(ctx, AccessControlerService_QueryAppAccessPolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlerServiceClient) GetAppLastAccessTime(ctx context.Context, in *IngressAppLastAccessTimeRequest, opts ...grpc.CallOption) (*IngressAppLastAccessTimeResponse, error) {
	out := new(IngressAppLastAccessTimeResponse)
	err := c.cc.Invoke(ctx, AccessControlerService_GetAppLastAccessTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessControlerServiceServer is the server API for AccessControlerService service.
// All implementations must embed UnimplementedAccessControlerServiceServer
// for forward compatibility
type AccessControlerServiceServer interface {
	SetAppAccessPolicy(context.Context, *AppAccessPolicyRequest) (*emptypb.Empty, error)
	QueryAppAccessPolicy(context.Context, *AppAccessPolicyRequest) (*AppAccessPolicy, error)
	GetAppLastAccessTime(context.Context, *IngressAppLastAccessTimeRequest) (*IngressAppLastAccessTimeResponse, error)
	mustEmbedUnimplementedAccessControlerServiceServer()
}

// UnimplementedAccessControlerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccessControlerServiceServer struct {
}

func (UnimplementedAccessControlerServiceServer) SetAppAccessPolicy(context.Context, *AppAccessPolicyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAppAccessPolicy not implemented")
}
func (UnimplementedAccessControlerServiceServer) QueryAppAccessPolicy(context.Context, *AppAccessPolicyRequest) (*AppAccessPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAppAccessPolicy not implemented")
}
func (UnimplementedAccessControlerServiceServer) GetAppLastAccessTime(context.Context, *IngressAppLastAccessTimeRequest) (*IngressAppLastAccessTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppLastAccessTime not implemented")
}
func (UnimplementedAccessControlerServiceServer) mustEmbedUnimplementedAccessControlerServiceServer() {
}

// UnsafeAccessControlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessControlerServiceServer will
// result in compilation errors.
type UnsafeAccessControlerServiceServer interface {
	mustEmbedUnimplementedAccessControlerServiceServer()
}

func RegisterAccessControlerServiceServer(s grpc.ServiceRegistrar, srv AccessControlerServiceServer) {
	s.RegisterService(&AccessControlerService_ServiceDesc, srv)
}

func _AccessControlerService_SetAppAccessPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppAccessPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlerServiceServer).SetAppAccessPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlerService_SetAppAccessPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlerServiceServer).SetAppAccessPolicy(ctx, req.(*AppAccessPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlerService_QueryAppAccessPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppAccessPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlerServiceServer).QueryAppAccessPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlerService_QueryAppAccessPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlerServiceServer).QueryAppAccessPolicy(ctx, req.(*AppAccessPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlerService_GetAppLastAccessTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IngressAppLastAccessTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlerServiceServer).GetAppLastAccessTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlerService_GetAppLastAccessTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlerServiceServer).GetAppLastAccessTime(ctx, req.(*IngressAppLastAccessTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessControlerService_ServiceDesc is the grpc.ServiceDesc for AccessControlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessControlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.sys.AccessControlerService",
	HandlerType: (*AccessControlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAppAccessPolicy",
			Handler:    _AccessControlerService_SetAppAccessPolicy_Handler,
		},
		{
			MethodName: "QueryAppAccessPolicy",
			Handler:    _AccessControlerService_QueryAppAccessPolicy_Handler,
		},
		{
			MethodName: "GetAppLastAccessTime",
			Handler:    _AccessControlerService_GetAppLastAccessTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys/ingress.proto",
}

const (
	IngressService_AllowAdd_FullMethodName             = "/cloud.lazycat.apis.sys.IngressService/AllowAdd"
	IngressService_AllowDel_FullMethodName             = "/cloud.lazycat.apis.sys.IngressService/AllowDel"
	IngressService_AllowDelAll_FullMethodName          = "/cloud.lazycat.apis.sys.IngressService/AllowDelAll"
	IngressService_AllowGet_FullMethodName             = "/cloud.lazycat.apis.sys.IngressService/AllowGet"
	IngressService_AllowList_FullMethodName            = "/cloud.lazycat.apis.sys.IngressService/AllowList"
	IngressService_AllowClear_FullMethodName           = "/cloud.lazycat.apis.sys.IngressService/AllowClear"
	IngressService_AllowManage_FullMethodName          = "/cloud.lazycat.apis.sys.IngressService/AllowManage"
	IngressService_GetAppLastAccessTime_FullMethodName = "/cloud.lazycat.apis.sys.IngressService/GetAppLastAccessTime"
)

// IngressServiceClient is the client API for IngressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Deprecated: Do not use.
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

// Deprecated: Do not use.
func NewIngressServiceClient(cc grpc.ClientConnInterface) IngressServiceClient {
	return &ingressServiceClient{cc}
}

func (c *ingressServiceClient) AllowAdd(ctx context.Context, in *IngressAllowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IngressService_AllowAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowDel(ctx context.Context, in *IngressAllowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IngressService_AllowDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowDelAll(ctx context.Context, in *IngressAllowDelAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IngressService_AllowDelAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowGet(ctx context.Context, in *IngressAllowGetRequest, opts ...grpc.CallOption) (*IngressAllowGetResponse, error) {
	out := new(IngressAllowGetResponse)
	err := c.cc.Invoke(ctx, IngressService_AllowGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowList(ctx context.Context, in *IngressAllowListRequest, opts ...grpc.CallOption) (*IngressAllowListResponse, error) {
	out := new(IngressAllowListResponse)
	err := c.cc.Invoke(ctx, IngressService_AllowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowClear(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IngressService_AllowClear_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) AllowManage(ctx context.Context, in *IngressAllowManageRequest, opts ...grpc.CallOption) (*IngressAllowManageResponse, error) {
	out := new(IngressAllowManageResponse)
	err := c.cc.Invoke(ctx, IngressService_AllowManage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ingressServiceClient) GetAppLastAccessTime(ctx context.Context, in *IngressAppLastAccessTimeRequest, opts ...grpc.CallOption) (*IngressAppLastAccessTimeResponse, error) {
	out := new(IngressAppLastAccessTimeResponse)
	err := c.cc.Invoke(ctx, IngressService_GetAppLastAccessTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngressServiceServer is the server API for IngressService service.
// All implementations must embed UnimplementedIngressServiceServer
// for forward compatibility
//
// Deprecated: Do not use.
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

// Deprecated: Do not use.
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
		FullMethod: IngressService_AllowAdd_FullMethodName,
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
		FullMethod: IngressService_AllowDel_FullMethodName,
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
		FullMethod: IngressService_AllowDelAll_FullMethodName,
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
		FullMethod: IngressService_AllowGet_FullMethodName,
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
		FullMethod: IngressService_AllowList_FullMethodName,
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
		FullMethod: IngressService_AllowClear_FullMethodName,
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
		FullMethod: IngressService_AllowManage_FullMethodName,
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
		FullMethod: IngressService_GetAppLastAccessTime_FullMethodName,
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

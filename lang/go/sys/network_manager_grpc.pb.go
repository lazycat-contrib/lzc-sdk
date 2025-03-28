// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: sys/network_manager.proto

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
	NetworkManager_Status_FullMethodName          = "/cloud.lazycat.apis.sys.NetworkManager/Status"
	NetworkManager_WifiScan_FullMethodName        = "/cloud.lazycat.apis.sys.NetworkManager/WifiScan"
	NetworkManager_WifiList_FullMethodName        = "/cloud.lazycat.apis.sys.NetworkManager/WifiList"
	NetworkManager_WifiConnect_FullMethodName     = "/cloud.lazycat.apis.sys.NetworkManager/WifiConnect"
	NetworkManager_WifiConnectTemp_FullMethodName = "/cloud.lazycat.apis.sys.NetworkManager/WifiConnectTemp"
	NetworkManager_WifiForget_FullMethodName      = "/cloud.lazycat.apis.sys.NetworkManager/WifiForget"
	NetworkManager_WifiConfigAdd_FullMethodName   = "/cloud.lazycat.apis.sys.NetworkManager/WifiConfigAdd"
	NetworkManager_GetConnectivity_FullMethodName = "/cloud.lazycat.apis.sys.NetworkManager/GetConnectivity"
	NetworkManager_Connectivity_FullMethodName    = "/cloud.lazycat.apis.sys.NetworkManager/Connectivity"
	NetworkManager_NmcliCall_FullMethodName       = "/cloud.lazycat.apis.sys.NetworkManager/NmcliCall"
)

// NetworkManagerClient is the client API for NetworkManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkManagerClient interface {
	// 获取网络设备的状态（是否已连接，连接了哪个）
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NetworkDeviceStatusInfo, error)
	// Scan 扫描附近wifi热点信息，扫描结果在内部缓存里（阻塞，可能耗费数秒）
	WifiScan(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List 列出内部缓存中的 wifi 列表
	WifiList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AccessPointInfoList, error)
	// 连接一个 wifi 热点
	//
	//	连接失败会删除已保存的配置，并自动连回上一次连接的 wifi（如果有的话），防止失联
	WifiConnect(ctx context.Context, in *WifiConnectInfo, opts ...grpc.CallOption) (*WifiConnectReply, error)
	// 暂时连接一个 wifi 热点
	// 时间到了之后会Revert回指定的 wifi 热点
	// 如果在上一个调用的duration时间范围内再次调用，则会取消上次调用时间到后对fallback_bssid的连接
	WifiConnectTemp(ctx context.Context, in *WifiConnectTempInfo, opts ...grpc.CallOption) (*WifiConnectReply, error)
	// 忘记一个 wifi 热点
	WifiForget(ctx context.Context, in *WifiForgetInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 手动添加和连接一个 wifi 热点配置（用于连接隐藏网络）
	WifiConfigAdd(ctx context.Context, in *WifiConfigInfo, opts ...grpc.CallOption) (*WifiConnectReply, error)
	// (不建议使用) nmcli networking connectivity check
	GetConnectivity(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConnectivityReply, error)
	// 自己实现的，返回格式和 nmcli networking connectivity check 一样
	Connectivity(ctx context.Context, in *ConnectivityRequest, opts ...grpc.CallOption) (*ConnectivityReply, error)
	// 直接调用 nmcli
	NmcliCall(ctx context.Context, in *NmcliCallRequest, opts ...grpc.CallOption) (*NmcliCallReply, error)
}

type networkManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkManagerClient(cc grpc.ClientConnInterface) NetworkManagerClient {
	return &networkManagerClient{cc}
}

func (c *networkManagerClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NetworkDeviceStatusInfo, error) {
	out := new(NetworkDeviceStatusInfo)
	err := c.cc.Invoke(ctx, NetworkManager_Status_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkManagerClient) WifiScan(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NetworkManager_WifiScan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkManagerClient) WifiList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AccessPointInfoList, error) {
	out := new(AccessPointInfoList)
	err := c.cc.Invoke(ctx, NetworkManager_WifiList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkManagerClient) WifiConnect(ctx context.Context, in *WifiConnectInfo, opts ...grpc.CallOption) (*WifiConnectReply, error) {
	out := new(WifiConnectReply)
	err := c.cc.Invoke(ctx, NetworkManager_WifiConnect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkManagerClient) WifiConnectTemp(ctx context.Context, in *WifiConnectTempInfo, opts ...grpc.CallOption) (*WifiConnectReply, error) {
	out := new(WifiConnectReply)
	err := c.cc.Invoke(ctx, NetworkManager_WifiConnectTemp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkManagerClient) WifiForget(ctx context.Context, in *WifiForgetInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, NetworkManager_WifiForget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkManagerClient) WifiConfigAdd(ctx context.Context, in *WifiConfigInfo, opts ...grpc.CallOption) (*WifiConnectReply, error) {
	out := new(WifiConnectReply)
	err := c.cc.Invoke(ctx, NetworkManager_WifiConfigAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkManagerClient) GetConnectivity(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetConnectivityReply, error) {
	out := new(GetConnectivityReply)
	err := c.cc.Invoke(ctx, NetworkManager_GetConnectivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkManagerClient) Connectivity(ctx context.Context, in *ConnectivityRequest, opts ...grpc.CallOption) (*ConnectivityReply, error) {
	out := new(ConnectivityReply)
	err := c.cc.Invoke(ctx, NetworkManager_Connectivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkManagerClient) NmcliCall(ctx context.Context, in *NmcliCallRequest, opts ...grpc.CallOption) (*NmcliCallReply, error) {
	out := new(NmcliCallReply)
	err := c.cc.Invoke(ctx, NetworkManager_NmcliCall_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkManagerServer is the server API for NetworkManager service.
// All implementations must embed UnimplementedNetworkManagerServer
// for forward compatibility
type NetworkManagerServer interface {
	// 获取网络设备的状态（是否已连接，连接了哪个）
	Status(context.Context, *emptypb.Empty) (*NetworkDeviceStatusInfo, error)
	// Scan 扫描附近wifi热点信息，扫描结果在内部缓存里（阻塞，可能耗费数秒）
	WifiScan(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// List 列出内部缓存中的 wifi 列表
	WifiList(context.Context, *emptypb.Empty) (*AccessPointInfoList, error)
	// 连接一个 wifi 热点
	//
	//	连接失败会删除已保存的配置，并自动连回上一次连接的 wifi（如果有的话），防止失联
	WifiConnect(context.Context, *WifiConnectInfo) (*WifiConnectReply, error)
	// 暂时连接一个 wifi 热点
	// 时间到了之后会Revert回指定的 wifi 热点
	// 如果在上一个调用的duration时间范围内再次调用，则会取消上次调用时间到后对fallback_bssid的连接
	WifiConnectTemp(context.Context, *WifiConnectTempInfo) (*WifiConnectReply, error)
	// 忘记一个 wifi 热点
	WifiForget(context.Context, *WifiForgetInfo) (*emptypb.Empty, error)
	// 手动添加和连接一个 wifi 热点配置（用于连接隐藏网络）
	WifiConfigAdd(context.Context, *WifiConfigInfo) (*WifiConnectReply, error)
	// (不建议使用) nmcli networking connectivity check
	GetConnectivity(context.Context, *emptypb.Empty) (*GetConnectivityReply, error)
	// 自己实现的，返回格式和 nmcli networking connectivity check 一样
	Connectivity(context.Context, *ConnectivityRequest) (*ConnectivityReply, error)
	// 直接调用 nmcli
	NmcliCall(context.Context, *NmcliCallRequest) (*NmcliCallReply, error)
	mustEmbedUnimplementedNetworkManagerServer()
}

// UnimplementedNetworkManagerServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkManagerServer struct {
}

func (UnimplementedNetworkManagerServer) Status(context.Context, *emptypb.Empty) (*NetworkDeviceStatusInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedNetworkManagerServer) WifiScan(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WifiScan not implemented")
}
func (UnimplementedNetworkManagerServer) WifiList(context.Context, *emptypb.Empty) (*AccessPointInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WifiList not implemented")
}
func (UnimplementedNetworkManagerServer) WifiConnect(context.Context, *WifiConnectInfo) (*WifiConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WifiConnect not implemented")
}
func (UnimplementedNetworkManagerServer) WifiConnectTemp(context.Context, *WifiConnectTempInfo) (*WifiConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WifiConnectTemp not implemented")
}
func (UnimplementedNetworkManagerServer) WifiForget(context.Context, *WifiForgetInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WifiForget not implemented")
}
func (UnimplementedNetworkManagerServer) WifiConfigAdd(context.Context, *WifiConfigInfo) (*WifiConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WifiConfigAdd not implemented")
}
func (UnimplementedNetworkManagerServer) GetConnectivity(context.Context, *emptypb.Empty) (*GetConnectivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectivity not implemented")
}
func (UnimplementedNetworkManagerServer) Connectivity(context.Context, *ConnectivityRequest) (*ConnectivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connectivity not implemented")
}
func (UnimplementedNetworkManagerServer) NmcliCall(context.Context, *NmcliCallRequest) (*NmcliCallReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NmcliCall not implemented")
}
func (UnimplementedNetworkManagerServer) mustEmbedUnimplementedNetworkManagerServer() {}

// UnsafeNetworkManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkManagerServer will
// result in compilation errors.
type UnsafeNetworkManagerServer interface {
	mustEmbedUnimplementedNetworkManagerServer()
}

func RegisterNetworkManagerServer(s grpc.ServiceRegistrar, srv NetworkManagerServer) {
	s.RegisterService(&NetworkManager_ServiceDesc, srv)
}

func _NetworkManager_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkManager_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkManager_WifiScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).WifiScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkManager_WifiScan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).WifiScan(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkManager_WifiList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).WifiList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkManager_WifiList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).WifiList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkManager_WifiConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiConnectInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).WifiConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkManager_WifiConnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).WifiConnect(ctx, req.(*WifiConnectInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkManager_WifiConnectTemp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiConnectTempInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).WifiConnectTemp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkManager_WifiConnectTemp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).WifiConnectTemp(ctx, req.(*WifiConnectTempInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkManager_WifiForget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiForgetInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).WifiForget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkManager_WifiForget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).WifiForget(ctx, req.(*WifiForgetInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkManager_WifiConfigAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiConfigInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).WifiConfigAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkManager_WifiConfigAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).WifiConfigAdd(ctx, req.(*WifiConfigInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkManager_GetConnectivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).GetConnectivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkManager_GetConnectivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).GetConnectivity(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkManager_Connectivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).Connectivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkManager_Connectivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).Connectivity(ctx, req.(*ConnectivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkManager_NmcliCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NmcliCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkManagerServer).NmcliCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkManager_NmcliCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkManagerServer).NmcliCall(ctx, req.(*NmcliCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkManager_ServiceDesc is the grpc.ServiceDesc for NetworkManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.sys.NetworkManager",
	HandlerType: (*NetworkManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _NetworkManager_Status_Handler,
		},
		{
			MethodName: "WifiScan",
			Handler:    _NetworkManager_WifiScan_Handler,
		},
		{
			MethodName: "WifiList",
			Handler:    _NetworkManager_WifiList_Handler,
		},
		{
			MethodName: "WifiConnect",
			Handler:    _NetworkManager_WifiConnect_Handler,
		},
		{
			MethodName: "WifiConnectTemp",
			Handler:    _NetworkManager_WifiConnectTemp_Handler,
		},
		{
			MethodName: "WifiForget",
			Handler:    _NetworkManager_WifiForget_Handler,
		},
		{
			MethodName: "WifiConfigAdd",
			Handler:    _NetworkManager_WifiConfigAdd_Handler,
		},
		{
			MethodName: "GetConnectivity",
			Handler:    _NetworkManager_GetConnectivity_Handler,
		},
		{
			MethodName: "Connectivity",
			Handler:    _NetworkManager_Connectivity_Handler,
		},
		{
			MethodName: "NmcliCall",
			Handler:    _NetworkManager_NmcliCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys/network_manager.proto",
}

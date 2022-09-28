// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: common/filetrans.proto

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

// FileSyncServiceClient is the client API for FileSyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileSyncServiceClient interface {
	// 创建队列
	CreateQueue(ctx context.Context, in *TaskQueueID, opts ...grpc.CallOption) (*FileTaskQueueResp, error)
	// 列出所有 QueueID
	ListQueue(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TaskQueueListResp, error)
	// 通过队列的 ID 和 Status 获取任务列表
	QueryQueue(ctx context.Context, in *TaskQueueQueryReq, opts ...grpc.CallOption) (FileSyncService_QueryQueueClient, error)
	// 通过队列的 ID 和 Status 清除任务
	ClearQueue(ctx context.Context, in *TaskQueueQueryReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 根据队列 ID 设置队列的速率并发等设置
	ConfigQueue(ctx context.Context, in *TaskQueueConfigReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 根据队列 ID 暂停队列
	PauseQueue(ctx context.Context, in *TaskQueueID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 根据队列 ID 开始队列
	StartQuque(ctx context.Context, in *TaskQueueID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 根据队列 ID 获取队列动态信息（比如 pending 状态的增加了 task1）
	QueryQueueMessage(ctx context.Context, in *TaskQueueQueryReq, opts ...grpc.CallOption) (FileSyncService_QueryQueueMessageClient, error)
	// 暂时不支持创建Task时创建任务，需要提前创建好任务。queue_id不存在则报错
	CreateTask(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*Task, error)
	// 根据 ID 获取单个任务状态
	QueryTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (FileSyncService_QueryTaskClient, error)
	// 根据 ID 开始单个任务
	ResumeTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 根据 ID 暂停单个任务
	PauseTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 根据 ID 删除单个任务
	DeleteTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type fileSyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileSyncServiceClient(cc grpc.ClientConnInterface) FileSyncServiceClient {
	return &fileSyncServiceClient{cc}
}

func (c *fileSyncServiceClient) CreateQueue(ctx context.Context, in *TaskQueueID, opts ...grpc.CallOption) (*FileTaskQueueResp, error) {
	out := new(FileTaskQueueResp)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.FileSyncService/CreateQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) ListQueue(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TaskQueueListResp, error) {
	out := new(TaskQueueListResp)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.FileSyncService/ListQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) QueryQueue(ctx context.Context, in *TaskQueueQueryReq, opts ...grpc.CallOption) (FileSyncService_QueryQueueClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSyncService_ServiceDesc.Streams[0], "/cloud.lazycat.apis.common.FileSyncService/QueryQueue", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSyncServiceQueryQueueClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSyncService_QueryQueueClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type fileSyncServiceQueryQueueClient struct {
	grpc.ClientStream
}

func (x *fileSyncServiceQueryQueueClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSyncServiceClient) ClearQueue(ctx context.Context, in *TaskQueueQueryReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.FileSyncService/ClearQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) ConfigQueue(ctx context.Context, in *TaskQueueConfigReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.FileSyncService/ConfigQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) PauseQueue(ctx context.Context, in *TaskQueueID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.FileSyncService/PauseQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) StartQuque(ctx context.Context, in *TaskQueueID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.FileSyncService/StartQuque", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) QueryQueueMessage(ctx context.Context, in *TaskQueueQueryReq, opts ...grpc.CallOption) (FileSyncService_QueryQueueMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSyncService_ServiceDesc.Streams[1], "/cloud.lazycat.apis.common.FileSyncService/QueryQueueMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSyncServiceQueryQueueMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSyncService_QueryQueueMessageClient interface {
	Recv() (*QueueMessageResp, error)
	grpc.ClientStream
}

type fileSyncServiceQueryQueueMessageClient struct {
	grpc.ClientStream
}

func (x *fileSyncServiceQueryQueueMessageClient) Recv() (*QueueMessageResp, error) {
	m := new(QueueMessageResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSyncServiceClient) CreateTask(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.FileSyncService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) QueryTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (FileSyncService_QueryTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSyncService_ServiceDesc.Streams[2], "/cloud.lazycat.apis.common.FileSyncService/QueryTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSyncServiceQueryTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSyncService_QueryTaskClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type fileSyncServiceQueryTaskClient struct {
	grpc.ClientStream
}

func (x *fileSyncServiceQueryTaskClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSyncServiceClient) ResumeTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.FileSyncService/ResumeTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) PauseTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.FileSyncService/PauseTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) DeleteTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/cloud.lazycat.apis.common.FileSyncService/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileSyncServiceServer is the server API for FileSyncService service.
// All implementations must embed UnimplementedFileSyncServiceServer
// for forward compatibility
type FileSyncServiceServer interface {
	// 创建队列
	CreateQueue(context.Context, *TaskQueueID) (*FileTaskQueueResp, error)
	// 列出所有 QueueID
	ListQueue(context.Context, *emptypb.Empty) (*TaskQueueListResp, error)
	// 通过队列的 ID 和 Status 获取任务列表
	QueryQueue(*TaskQueueQueryReq, FileSyncService_QueryQueueServer) error
	// 通过队列的 ID 和 Status 清除任务
	ClearQueue(context.Context, *TaskQueueQueryReq) (*emptypb.Empty, error)
	// 根据队列 ID 设置队列的速率并发等设置
	ConfigQueue(context.Context, *TaskQueueConfigReq) (*emptypb.Empty, error)
	// 根据队列 ID 暂停队列
	PauseQueue(context.Context, *TaskQueueID) (*emptypb.Empty, error)
	// 根据队列 ID 开始队列
	StartQuque(context.Context, *TaskQueueID) (*emptypb.Empty, error)
	// 根据队列 ID 获取队列动态信息（比如 pending 状态的增加了 task1）
	QueryQueueMessage(*TaskQueueQueryReq, FileSyncService_QueryQueueMessageServer) error
	// 暂时不支持创建Task时创建任务，需要提前创建好任务。queue_id不存在则报错
	CreateTask(context.Context, *TaskCreateRequest) (*Task, error)
	// 根据 ID 获取单个任务状态
	QueryTask(*TaskID, FileSyncService_QueryTaskServer) error
	// 根据 ID 开始单个任务
	ResumeTask(context.Context, *TaskID) (*emptypb.Empty, error)
	// 根据 ID 暂停单个任务
	PauseTask(context.Context, *TaskID) (*emptypb.Empty, error)
	// 根据 ID 删除单个任务
	DeleteTask(context.Context, *TaskID) (*emptypb.Empty, error)
	mustEmbedUnimplementedFileSyncServiceServer()
}

// UnimplementedFileSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileSyncServiceServer struct {
}

func (UnimplementedFileSyncServiceServer) CreateQueue(context.Context, *TaskQueueID) (*FileTaskQueueResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQueue not implemented")
}
func (UnimplementedFileSyncServiceServer) ListQueue(context.Context, *emptypb.Empty) (*TaskQueueListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQueue not implemented")
}
func (UnimplementedFileSyncServiceServer) QueryQueue(*TaskQueueQueryReq, FileSyncService_QueryQueueServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryQueue not implemented")
}
func (UnimplementedFileSyncServiceServer) ClearQueue(context.Context, *TaskQueueQueryReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearQueue not implemented")
}
func (UnimplementedFileSyncServiceServer) ConfigQueue(context.Context, *TaskQueueConfigReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigQueue not implemented")
}
func (UnimplementedFileSyncServiceServer) PauseQueue(context.Context, *TaskQueueID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseQueue not implemented")
}
func (UnimplementedFileSyncServiceServer) StartQuque(context.Context, *TaskQueueID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartQuque not implemented")
}
func (UnimplementedFileSyncServiceServer) QueryQueueMessage(*TaskQueueQueryReq, FileSyncService_QueryQueueMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryQueueMessage not implemented")
}
func (UnimplementedFileSyncServiceServer) CreateTask(context.Context, *TaskCreateRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedFileSyncServiceServer) QueryTask(*TaskID, FileSyncService_QueryTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryTask not implemented")
}
func (UnimplementedFileSyncServiceServer) ResumeTask(context.Context, *TaskID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeTask not implemented")
}
func (UnimplementedFileSyncServiceServer) PauseTask(context.Context, *TaskID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseTask not implemented")
}
func (UnimplementedFileSyncServiceServer) DeleteTask(context.Context, *TaskID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedFileSyncServiceServer) mustEmbedUnimplementedFileSyncServiceServer() {}

// UnsafeFileSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileSyncServiceServer will
// result in compilation errors.
type UnsafeFileSyncServiceServer interface {
	mustEmbedUnimplementedFileSyncServiceServer()
}

func RegisterFileSyncServiceServer(s grpc.ServiceRegistrar, srv FileSyncServiceServer) {
	s.RegisterService(&FileSyncService_ServiceDesc, srv)
}

func _FileSyncService_CreateQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskQueueID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).CreateQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.FileSyncService/CreateQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).CreateQueue(ctx, req.(*TaskQueueID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_ListQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).ListQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.FileSyncService/ListQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).ListQueue(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_QueryQueue_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskQueueQueryReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSyncServiceServer).QueryQueue(m, &fileSyncServiceQueryQueueServer{stream})
}

type FileSyncService_QueryQueueServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type fileSyncServiceQueryQueueServer struct {
	grpc.ServerStream
}

func (x *fileSyncServiceQueryQueueServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSyncService_ClearQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskQueueQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).ClearQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.FileSyncService/ClearQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).ClearQueue(ctx, req.(*TaskQueueQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_ConfigQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskQueueConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).ConfigQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.FileSyncService/ConfigQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).ConfigQueue(ctx, req.(*TaskQueueConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_PauseQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskQueueID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).PauseQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.FileSyncService/PauseQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).PauseQueue(ctx, req.(*TaskQueueID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_StartQuque_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskQueueID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).StartQuque(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.FileSyncService/StartQuque",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).StartQuque(ctx, req.(*TaskQueueID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_QueryQueueMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskQueueQueryReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSyncServiceServer).QueryQueueMessage(m, &fileSyncServiceQueryQueueMessageServer{stream})
}

type FileSyncService_QueryQueueMessageServer interface {
	Send(*QueueMessageResp) error
	grpc.ServerStream
}

type fileSyncServiceQueryQueueMessageServer struct {
	grpc.ServerStream
}

func (x *fileSyncServiceQueryQueueMessageServer) Send(m *QueueMessageResp) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSyncService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.FileSyncService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).CreateTask(ctx, req.(*TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_QueryTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSyncServiceServer).QueryTask(m, &fileSyncServiceQueryTaskServer{stream})
}

type FileSyncService_QueryTaskServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type fileSyncServiceQueryTaskServer struct {
	grpc.ServerStream
}

func (x *fileSyncServiceQueryTaskServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSyncService_ResumeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).ResumeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.FileSyncService/ResumeTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).ResumeTask(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_PauseTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).PauseTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.FileSyncService/PauseTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).PauseTask(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.lazycat.apis.common.FileSyncService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).DeleteTask(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

// FileSyncService_ServiceDesc is the grpc.ServiceDesc for FileSyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileSyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.common.FileSyncService",
	HandlerType: (*FileSyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQueue",
			Handler:    _FileSyncService_CreateQueue_Handler,
		},
		{
			MethodName: "ListQueue",
			Handler:    _FileSyncService_ListQueue_Handler,
		},
		{
			MethodName: "ClearQueue",
			Handler:    _FileSyncService_ClearQueue_Handler,
		},
		{
			MethodName: "ConfigQueue",
			Handler:    _FileSyncService_ConfigQueue_Handler,
		},
		{
			MethodName: "PauseQueue",
			Handler:    _FileSyncService_PauseQueue_Handler,
		},
		{
			MethodName: "StartQuque",
			Handler:    _FileSyncService_StartQuque_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _FileSyncService_CreateTask_Handler,
		},
		{
			MethodName: "ResumeTask",
			Handler:    _FileSyncService_ResumeTask_Handler,
		},
		{
			MethodName: "PauseTask",
			Handler:    _FileSyncService_PauseTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _FileSyncService_DeleteTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryQueue",
			Handler:       _FileSyncService_QueryQueue_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "QueryQueueMessage",
			Handler:       _FileSyncService_QueryQueueMessage_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "QueryTask",
			Handler:       _FileSyncService_QueryTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "common/filetrans.proto",
}

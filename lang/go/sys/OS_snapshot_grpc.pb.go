// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: sys/OS_snapshot.proto

package sys

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OSSnapshotServiceClient is the client API for OSSnapshotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OSSnapshotServiceClient interface {
}

type oSSnapshotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOSSnapshotServiceClient(cc grpc.ClientConnInterface) OSSnapshotServiceClient {
	return &oSSnapshotServiceClient{cc}
}

// OSSnapshotServiceServer is the server API for OSSnapshotService service.
// All implementations must embed UnimplementedOSSnapshotServiceServer
// for forward compatibility
type OSSnapshotServiceServer interface {
	mustEmbedUnimplementedOSSnapshotServiceServer()
}

// UnimplementedOSSnapshotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOSSnapshotServiceServer struct {
}

func (UnimplementedOSSnapshotServiceServer) mustEmbedUnimplementedOSSnapshotServiceServer() {}

// UnsafeOSSnapshotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OSSnapshotServiceServer will
// result in compilation errors.
type UnsafeOSSnapshotServiceServer interface {
	mustEmbedUnimplementedOSSnapshotServiceServer()
}

func RegisterOSSnapshotServiceServer(s grpc.ServiceRegistrar, srv OSSnapshotServiceServer) {
	s.RegisterService(&OSSnapshotService_ServiceDesc, srv)
}

// OSSnapshotService_ServiceDesc is the grpc.ServiceDesc for OSSnapshotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OSSnapshotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.sys.OSSnapshotService",
	HandlerType: (*OSSnapshotServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "sys/OS_snapshot.proto",
}

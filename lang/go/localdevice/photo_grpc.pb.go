// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: localdevice/photo.proto

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
	PhotoLibrary_MakeAlbum_FullMethodName      = "/cloud.lazycat.apis.localdevice.PhotoLibrary/MakeAlbum"
	PhotoLibrary_ListAlbums_FullMethodName     = "/cloud.lazycat.apis.localdevice.PhotoLibrary/ListAlbums"
	PhotoLibrary_PutPhoto_FullMethodName       = "/cloud.lazycat.apis.localdevice.PhotoLibrary/PutPhoto"
	PhotoLibrary_DeletePhoto_FullMethodName    = "/cloud.lazycat.apis.localdevice.PhotoLibrary/DeletePhoto"
	PhotoLibrary_ListPhotoMetas_FullMethodName = "/cloud.lazycat.apis.localdevice.PhotoLibrary/ListPhotoMetas"
	PhotoLibrary_ListPhotos_FullMethodName     = "/cloud.lazycat.apis.localdevice.PhotoLibrary/ListPhotos"
	PhotoLibrary_QueryPhoto_FullMethodName     = "/cloud.lazycat.apis.localdevice.PhotoLibrary/QueryPhoto"
)

// PhotoLibraryClient is the client API for PhotoLibrary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhotoLibraryClient interface {
	MakeAlbum(ctx context.Context, in *MakeAlbumRequest, opts ...grpc.CallOption) (*Album, error)
	// 列举所有的系统相册
	ListAlbums(ctx context.Context, in *ListAlbumsRequest, opts ...grpc.CallOption) (*ListAlbumsReply, error)
	// 存储一张图片到某个相册中
	PutPhoto(ctx context.Context, in *PutPhotoRequest, opts ...grpc.CallOption) (PhotoLibrary_PutPhotoClient, error)
	DeletePhoto(ctx context.Context, in *DeletePhotoRequest, opts ...grpc.CallOption) (*DeletePhotoReply, error)
	// 枚举具体相册中的图片元信息
	ListPhotoMetas(ctx context.Context, in *ListPhotoMetasRequest, opts ...grpc.CallOption) (PhotoLibrary_ListPhotoMetasClient, error)
	// 列举所有的系统图片
	ListPhotos(ctx context.Context, in *ListPhotoMetasRequest, opts ...grpc.CallOption) (*ListPhotosReply, error)
	QueryPhoto(ctx context.Context, in *QueryPhotoRequest, opts ...grpc.CallOption) (*PhotoMeta, error)
}

type photoLibraryClient struct {
	cc grpc.ClientConnInterface
}

func NewPhotoLibraryClient(cc grpc.ClientConnInterface) PhotoLibraryClient {
	return &photoLibraryClient{cc}
}

func (c *photoLibraryClient) MakeAlbum(ctx context.Context, in *MakeAlbumRequest, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := c.cc.Invoke(ctx, PhotoLibrary_MakeAlbum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoLibraryClient) ListAlbums(ctx context.Context, in *ListAlbumsRequest, opts ...grpc.CallOption) (*ListAlbumsReply, error) {
	out := new(ListAlbumsReply)
	err := c.cc.Invoke(ctx, PhotoLibrary_ListAlbums_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoLibraryClient) PutPhoto(ctx context.Context, in *PutPhotoRequest, opts ...grpc.CallOption) (PhotoLibrary_PutPhotoClient, error) {
	stream, err := c.cc.NewStream(ctx, &PhotoLibrary_ServiceDesc.Streams[0], PhotoLibrary_PutPhoto_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &photoLibraryPutPhotoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PhotoLibrary_PutPhotoClient interface {
	Recv() (*PutPhotoReply, error)
	grpc.ClientStream
}

type photoLibraryPutPhotoClient struct {
	grpc.ClientStream
}

func (x *photoLibraryPutPhotoClient) Recv() (*PutPhotoReply, error) {
	m := new(PutPhotoReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *photoLibraryClient) DeletePhoto(ctx context.Context, in *DeletePhotoRequest, opts ...grpc.CallOption) (*DeletePhotoReply, error) {
	out := new(DeletePhotoReply)
	err := c.cc.Invoke(ctx, PhotoLibrary_DeletePhoto_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoLibraryClient) ListPhotoMetas(ctx context.Context, in *ListPhotoMetasRequest, opts ...grpc.CallOption) (PhotoLibrary_ListPhotoMetasClient, error) {
	stream, err := c.cc.NewStream(ctx, &PhotoLibrary_ServiceDesc.Streams[1], PhotoLibrary_ListPhotoMetas_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &photoLibraryListPhotoMetasClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PhotoLibrary_ListPhotoMetasClient interface {
	Recv() (*PhotoMeta, error)
	grpc.ClientStream
}

type photoLibraryListPhotoMetasClient struct {
	grpc.ClientStream
}

func (x *photoLibraryListPhotoMetasClient) Recv() (*PhotoMeta, error) {
	m := new(PhotoMeta)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *photoLibraryClient) ListPhotos(ctx context.Context, in *ListPhotoMetasRequest, opts ...grpc.CallOption) (*ListPhotosReply, error) {
	out := new(ListPhotosReply)
	err := c.cc.Invoke(ctx, PhotoLibrary_ListPhotos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoLibraryClient) QueryPhoto(ctx context.Context, in *QueryPhotoRequest, opts ...grpc.CallOption) (*PhotoMeta, error) {
	out := new(PhotoMeta)
	err := c.cc.Invoke(ctx, PhotoLibrary_QueryPhoto_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhotoLibraryServer is the server API for PhotoLibrary service.
// All implementations must embed UnimplementedPhotoLibraryServer
// for forward compatibility
type PhotoLibraryServer interface {
	MakeAlbum(context.Context, *MakeAlbumRequest) (*Album, error)
	// 列举所有的系统相册
	ListAlbums(context.Context, *ListAlbumsRequest) (*ListAlbumsReply, error)
	// 存储一张图片到某个相册中
	PutPhoto(*PutPhotoRequest, PhotoLibrary_PutPhotoServer) error
	DeletePhoto(context.Context, *DeletePhotoRequest) (*DeletePhotoReply, error)
	// 枚举具体相册中的图片元信息
	ListPhotoMetas(*ListPhotoMetasRequest, PhotoLibrary_ListPhotoMetasServer) error
	// 列举所有的系统图片
	ListPhotos(context.Context, *ListPhotoMetasRequest) (*ListPhotosReply, error)
	QueryPhoto(context.Context, *QueryPhotoRequest) (*PhotoMeta, error)
	mustEmbedUnimplementedPhotoLibraryServer()
}

// UnimplementedPhotoLibraryServer must be embedded to have forward compatible implementations.
type UnimplementedPhotoLibraryServer struct {
}

func (UnimplementedPhotoLibraryServer) MakeAlbum(context.Context, *MakeAlbumRequest) (*Album, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeAlbum not implemented")
}
func (UnimplementedPhotoLibraryServer) ListAlbums(context.Context, *ListAlbumsRequest) (*ListAlbumsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlbums not implemented")
}
func (UnimplementedPhotoLibraryServer) PutPhoto(*PutPhotoRequest, PhotoLibrary_PutPhotoServer) error {
	return status.Errorf(codes.Unimplemented, "method PutPhoto not implemented")
}
func (UnimplementedPhotoLibraryServer) DeletePhoto(context.Context, *DeletePhotoRequest) (*DeletePhotoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePhoto not implemented")
}
func (UnimplementedPhotoLibraryServer) ListPhotoMetas(*ListPhotoMetasRequest, PhotoLibrary_ListPhotoMetasServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPhotoMetas not implemented")
}
func (UnimplementedPhotoLibraryServer) ListPhotos(context.Context, *ListPhotoMetasRequest) (*ListPhotosReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPhotos not implemented")
}
func (UnimplementedPhotoLibraryServer) QueryPhoto(context.Context, *QueryPhotoRequest) (*PhotoMeta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPhoto not implemented")
}
func (UnimplementedPhotoLibraryServer) mustEmbedUnimplementedPhotoLibraryServer() {}

// UnsafePhotoLibraryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhotoLibraryServer will
// result in compilation errors.
type UnsafePhotoLibraryServer interface {
	mustEmbedUnimplementedPhotoLibraryServer()
}

func RegisterPhotoLibraryServer(s grpc.ServiceRegistrar, srv PhotoLibraryServer) {
	s.RegisterService(&PhotoLibrary_ServiceDesc, srv)
}

func _PhotoLibrary_MakeAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoLibraryServer).MakeAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoLibrary_MakeAlbum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoLibraryServer).MakeAlbum(ctx, req.(*MakeAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoLibrary_ListAlbums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlbumsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoLibraryServer).ListAlbums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoLibrary_ListAlbums_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoLibraryServer).ListAlbums(ctx, req.(*ListAlbumsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoLibrary_PutPhoto_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PutPhotoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PhotoLibraryServer).PutPhoto(m, &photoLibraryPutPhotoServer{stream})
}

type PhotoLibrary_PutPhotoServer interface {
	Send(*PutPhotoReply) error
	grpc.ServerStream
}

type photoLibraryPutPhotoServer struct {
	grpc.ServerStream
}

func (x *photoLibraryPutPhotoServer) Send(m *PutPhotoReply) error {
	return x.ServerStream.SendMsg(m)
}

func _PhotoLibrary_DeletePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoLibraryServer).DeletePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoLibrary_DeletePhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoLibraryServer).DeletePhoto(ctx, req.(*DeletePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoLibrary_ListPhotoMetas_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPhotoMetasRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PhotoLibraryServer).ListPhotoMetas(m, &photoLibraryListPhotoMetasServer{stream})
}

type PhotoLibrary_ListPhotoMetasServer interface {
	Send(*PhotoMeta) error
	grpc.ServerStream
}

type photoLibraryListPhotoMetasServer struct {
	grpc.ServerStream
}

func (x *photoLibraryListPhotoMetasServer) Send(m *PhotoMeta) error {
	return x.ServerStream.SendMsg(m)
}

func _PhotoLibrary_ListPhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPhotoMetasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoLibraryServer).ListPhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoLibrary_ListPhotos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoLibraryServer).ListPhotos(ctx, req.(*ListPhotoMetasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoLibrary_QueryPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoLibraryServer).QueryPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoLibrary_QueryPhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoLibraryServer).QueryPhoto(ctx, req.(*QueryPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PhotoLibrary_ServiceDesc is the grpc.ServiceDesc for PhotoLibrary service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhotoLibrary_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.lazycat.apis.localdevice.PhotoLibrary",
	HandlerType: (*PhotoLibraryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeAlbum",
			Handler:    _PhotoLibrary_MakeAlbum_Handler,
		},
		{
			MethodName: "ListAlbums",
			Handler:    _PhotoLibrary_ListAlbums_Handler,
		},
		{
			MethodName: "DeletePhoto",
			Handler:    _PhotoLibrary_DeletePhoto_Handler,
		},
		{
			MethodName: "ListPhotos",
			Handler:    _PhotoLibrary_ListPhotos_Handler,
		},
		{
			MethodName: "QueryPhoto",
			Handler:    _PhotoLibrary_QueryPhoto_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutPhoto",
			Handler:       _PhotoLibrary_PutPhoto_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListPhotoMetas",
			Handler:       _PhotoLibrary_ListPhotoMetas_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "localdevice/photo.proto",
}

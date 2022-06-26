// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.1
// source: localdevice/photo/photo.proto

package photo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListPhotoMetasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 相册ID, 若为空，则表示返回所有相册中的图片
	AlbumIds        []string `protobuf:"bytes,1,rep,name=album_ids,json=albumIds,proto3" json:"album_ids,omitempty"`
	ThumbnailWidth  int32    `protobuf:"varint,2,opt,name=thumbnail_width,json=thumbnailWidth,proto3" json:"thumbnail_width,omitempty"`
	ThumbnailHeight int32    `protobuf:"varint,3,opt,name=thumbnail_height,json=thumbnailHeight,proto3" json:"thumbnail_height,omitempty"`
	NeedFileName    bool     `protobuf:"varint,4,opt,name=need_file_name,json=needFileName,proto3" json:"need_file_name,omitempty"`
	NeedAlbumIds    bool     `protobuf:"varint,5,opt,name=need_album_ids,json=needAlbumIds,proto3" json:"need_album_ids,omitempty"`
}

func (x *ListPhotoMetasRequest) Reset() {
	*x = ListPhotoMetasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_photo_photo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPhotoMetasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPhotoMetasRequest) ProtoMessage() {}

func (x *ListPhotoMetasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_photo_photo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPhotoMetasRequest.ProtoReflect.Descriptor instead.
func (*ListPhotoMetasRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_photo_photo_proto_rawDescGZIP(), []int{0}
}

func (x *ListPhotoMetasRequest) GetAlbumIds() []string {
	if x != nil {
		return x.AlbumIds
	}
	return nil
}

func (x *ListPhotoMetasRequest) GetThumbnailWidth() int32 {
	if x != nil {
		return x.ThumbnailWidth
	}
	return 0
}

func (x *ListPhotoMetasRequest) GetThumbnailHeight() int32 {
	if x != nil {
		return x.ThumbnailHeight
	}
	return 0
}

func (x *ListPhotoMetasRequest) GetNeedFileName() bool {
	if x != nil {
		return x.NeedFileName
	}
	return false
}

func (x *ListPhotoMetasRequest) GetNeedAlbumIds() bool {
	if x != nil {
		return x.NeedAlbumIds
	}
	return false
}

type PhotoMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PhotoUrl     string                 `protobuf:"bytes,2,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"` //浏览器直接可以使用的url, 可能是device domain下的一个文件或是一个data url
	ThumbnailUrl string                 `protobuf:"bytes,3,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	Width        string                 `protobuf:"bytes,4,opt,name=width,proto3" json:"width,omitempty"`
	Height       string                 `protobuf:"bytes,5,opt,name=height,proto3" json:"height,omitempty"`
	AlbumIds     []string               `protobuf:"bytes,6,rep,name=album_ids,json=albumIds,proto3" json:"album_ids,omitempty"`
	FileName     *string                `protobuf:"bytes,7,opt,name=file_name,json=fileName,proto3,oneof" json:"file_name,omitempty"`
	CreationDate *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=creation_date,json=creationDate,proto3,oneof" json:"creation_date,omitempty"`
	Latitude     *float32               `protobuf:"fixed32,9,opt,name=latitude,proto3,oneof" json:"latitude,omitempty"`
	Longitude    *float32               `protobuf:"fixed32,10,opt,name=longitude,proto3,oneof" json:"longitude,omitempty"`
}

func (x *PhotoMeta) Reset() {
	*x = PhotoMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_photo_photo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoMeta) ProtoMessage() {}

func (x *PhotoMeta) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_photo_photo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoMeta.ProtoReflect.Descriptor instead.
func (*PhotoMeta) Descriptor() ([]byte, []int) {
	return file_localdevice_photo_photo_proto_rawDescGZIP(), []int{1}
}

func (x *PhotoMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PhotoMeta) GetPhotoUrl() string {
	if x != nil {
		return x.PhotoUrl
	}
	return ""
}

func (x *PhotoMeta) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *PhotoMeta) GetWidth() string {
	if x != nil {
		return x.Width
	}
	return ""
}

func (x *PhotoMeta) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *PhotoMeta) GetAlbumIds() []string {
	if x != nil {
		return x.AlbumIds
	}
	return nil
}

func (x *PhotoMeta) GetFileName() string {
	if x != nil && x.FileName != nil {
		return *x.FileName
	}
	return ""
}

func (x *PhotoMeta) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

func (x *PhotoMeta) GetLatitude() float32 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *PhotoMeta) GetLongitude() float32 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Album) Reset() {
	*x = Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_photo_photo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_photo_photo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Album.ProtoReflect.Descriptor instead.
func (*Album) Descriptor() ([]byte, []int) {
	return file_localdevice_photo_photo_proto_rawDescGZIP(), []int{2}
}

func (x *Album) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Album) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ListAlbumsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAlbumsRequest) Reset() {
	*x = ListAlbumsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_photo_photo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlbumsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlbumsRequest) ProtoMessage() {}

func (x *ListAlbumsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_photo_photo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlbumsRequest.ProtoReflect.Descriptor instead.
func (*ListAlbumsRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_photo_photo_proto_rawDescGZIP(), []int{3}
}

type ListAlbumsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Albums []*Album `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty"`
}

func (x *ListAlbumsReply) Reset() {
	*x = ListAlbumsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_photo_photo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlbumsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlbumsReply) ProtoMessage() {}

func (x *ListAlbumsReply) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_photo_photo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlbumsReply.ProtoReflect.Descriptor instead.
func (*ListAlbumsReply) Descriptor() ([]byte, []int) {
	return file_localdevice_photo_photo_proto_rawDescGZIP(), []int{4}
}

func (x *ListAlbumsReply) GetAlbums() []*Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

type PutPhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlbumId string `protobuf:"bytes,1,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
	// 图片路径, 支持dataurl
	Url      string  `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	FileName *string `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3,oneof" json:"file_name,omitempty"`
}

func (x *PutPhotoRequest) Reset() {
	*x = PutPhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_photo_photo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutPhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutPhotoRequest) ProtoMessage() {}

func (x *PutPhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_photo_photo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutPhotoRequest.ProtoReflect.Descriptor instead.
func (*PutPhotoRequest) Descriptor() ([]byte, []int) {
	return file_localdevice_photo_photo_proto_rawDescGZIP(), []int{5}
}

func (x *PutPhotoRequest) GetAlbumId() string {
	if x != nil {
		return x.AlbumId
	}
	return ""
}

func (x *PutPhotoRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PutPhotoRequest) GetFileName() string {
	if x != nil && x.FileName != nil {
		return *x.FileName
	}
	return ""
}

var File_localdevice_photo_photo_proto protoreflect.FileDescriptor

var file_localdevice_photo_photo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01,
	0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x49, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x29, 0x0a,
	0x10, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x65, 0x65, 0x64,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0e, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x49, 0x64, 0x73, 0x22, 0x8f, 0x03, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x73, 0x12,
	0x20, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x44, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x48, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x48, 0x03, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x2d, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d, 0x0a,
	0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x22, 0x6e, 0x0a, 0x0f,
	0x50, 0x75, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xd1, 0x02, 0x0a,
	0x0c, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x72, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x55, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x50,
	0x75, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x12, 0x35, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61,
	0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_localdevice_photo_photo_proto_rawDescOnce sync.Once
	file_localdevice_photo_photo_proto_rawDescData = file_localdevice_photo_photo_proto_rawDesc
)

func file_localdevice_photo_photo_proto_rawDescGZIP() []byte {
	file_localdevice_photo_photo_proto_rawDescOnce.Do(func() {
		file_localdevice_photo_photo_proto_rawDescData = protoimpl.X.CompressGZIP(file_localdevice_photo_photo_proto_rawDescData)
	})
	return file_localdevice_photo_photo_proto_rawDescData
}

var file_localdevice_photo_photo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_localdevice_photo_photo_proto_goTypes = []interface{}{
	(*ListPhotoMetasRequest)(nil), // 0: cloud.lazycat.localdevice.apis.ListPhotoMetasRequest
	(*PhotoMeta)(nil),             // 1: cloud.lazycat.localdevice.apis.PhotoMeta
	(*Album)(nil),                 // 2: cloud.lazycat.localdevice.apis.Album
	(*ListAlbumsRequest)(nil),     // 3: cloud.lazycat.localdevice.apis.ListAlbumsRequest
	(*ListAlbumsReply)(nil),       // 4: cloud.lazycat.localdevice.apis.ListAlbumsReply
	(*PutPhotoRequest)(nil),       // 5: cloud.lazycat.localdevice.apis.PutPhotoRequest
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_localdevice_photo_photo_proto_depIdxs = []int32{
	6, // 0: cloud.lazycat.localdevice.apis.PhotoMeta.creation_date:type_name -> google.protobuf.Timestamp
	2, // 1: cloud.lazycat.localdevice.apis.ListAlbumsReply.albums:type_name -> cloud.lazycat.localdevice.apis.Album
	3, // 2: cloud.lazycat.localdevice.apis.PhotoLibrary.ListAlbums:input_type -> cloud.lazycat.localdevice.apis.ListAlbumsRequest
	5, // 3: cloud.lazycat.localdevice.apis.PhotoLibrary.PutPhoto:input_type -> cloud.lazycat.localdevice.apis.PutPhotoRequest
	0, // 4: cloud.lazycat.localdevice.apis.PhotoLibrary.ListPhotoMetas:input_type -> cloud.lazycat.localdevice.apis.ListPhotoMetasRequest
	4, // 5: cloud.lazycat.localdevice.apis.PhotoLibrary.ListAlbums:output_type -> cloud.lazycat.localdevice.apis.ListAlbumsReply
	7, // 6: cloud.lazycat.localdevice.apis.PhotoLibrary.PutPhoto:output_type -> google.protobuf.Empty
	1, // 7: cloud.lazycat.localdevice.apis.PhotoLibrary.ListPhotoMetas:output_type -> cloud.lazycat.localdevice.apis.PhotoMeta
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_localdevice_photo_photo_proto_init() }
func file_localdevice_photo_photo_proto_init() {
	if File_localdevice_photo_photo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_localdevice_photo_photo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPhotoMetasRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_localdevice_photo_photo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_localdevice_photo_photo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Album); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_localdevice_photo_photo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlbumsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_localdevice_photo_photo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlbumsReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_localdevice_photo_photo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutPhotoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_localdevice_photo_photo_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_localdevice_photo_photo_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_localdevice_photo_photo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_localdevice_photo_photo_proto_goTypes,
		DependencyIndexes: file_localdevice_photo_photo_proto_depIdxs,
		MessageInfos:      file_localdevice_photo_photo_proto_msgTypes,
	}.Build()
	File_localdevice_photo_photo_proto = out.File
	file_localdevice_photo_photo_proto_rawDesc = nil
	file_localdevice_photo_photo_proto_goTypes = nil
	file_localdevice_photo_photo_proto_depIdxs = nil
}

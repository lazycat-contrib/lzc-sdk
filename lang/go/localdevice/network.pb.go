// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: localdevice/network.proto

package localdevice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConnectionType int32

const (
	ConnectionType_Unknown   ConnectionType = 0
	ConnectionType_CELL_NONE ConnectionType = 1
	ConnectionType_ETHERNET  ConnectionType = 2
	ConnectionType_WIFI      ConnectionType = 3
	ConnectionType_CELL      ConnectionType = 4
	ConnectionType_CELL_2G   ConnectionType = 5
	ConnectionType_CELL_3G   ConnectionType = 6
	ConnectionType_CELL_4G   ConnectionType = 7
	ConnectionType_CELL_5G   ConnectionType = 8
)

// Enum value maps for ConnectionType.
var (
	ConnectionType_name = map[int32]string{
		0: "Unknown",
		1: "CELL_NONE",
		2: "ETHERNET",
		3: "WIFI",
		4: "CELL",
		5: "CELL_2G",
		6: "CELL_3G",
		7: "CELL_4G",
		8: "CELL_5G",
	}
	ConnectionType_value = map[string]int32{
		"Unknown":   0,
		"CELL_NONE": 1,
		"ETHERNET":  2,
		"WIFI":      3,
		"CELL":      4,
		"CELL_2G":   5,
		"CELL_3G":   6,
		"CELL_4G":   7,
		"CELL_5G":   8,
	}
)

func (x ConnectionType) Enum() *ConnectionType {
	p := new(ConnectionType)
	*p = x
	return p
}

func (x ConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_localdevice_network_proto_enumTypes[0].Descriptor()
}

func (ConnectionType) Type() protoreflect.EnumType {
	return &file_localdevice_network_proto_enumTypes[0]
}

func (x ConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionType.Descriptor instead.
func (ConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_localdevice_network_proto_rawDescGZIP(), []int{0}
}

type NetworkInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctype    ConnectionType `protobuf:"varint,1,opt,name=ctype,proto3,enum=cloud.lazycat.apis.localdevice.ConnectionType" json:"ctype,omitempty"`
	IsOnline bool           `protobuf:"varint,2,opt,name=IsOnline,proto3" json:"IsOnline,omitempty"`
}

func (x *NetworkInformation) Reset() {
	*x = NetworkInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_localdevice_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInformation) ProtoMessage() {}

func (x *NetworkInformation) ProtoReflect() protoreflect.Message {
	mi := &file_localdevice_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInformation.ProtoReflect.Descriptor instead.
func (*NetworkInformation) Descriptor() ([]byte, []int) {
	return file_localdevice_network_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkInformation) GetCtype() ConnectionType {
	if x != nil {
		return x.Ctype
	}
	return ConnectionType_Unknown
}

func (x *NetworkInformation) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

var File_localdevice_network_proto protoreflect.FileDescriptor

var file_localdevice_network_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x12, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44,
	0x0a, 0x05, 0x63, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x63,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x2a, 0x82, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x45, 0x54, 0x48, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x57, 0x49, 0x46, 0x49, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x45, 0x4c, 0x4c, 0x10,
	0x04, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x32, 0x47, 0x10, 0x05, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x33, 0x47, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x45, 0x4c, 0x4c, 0x5f, 0x34, 0x47, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x45, 0x4c, 0x4c,
	0x5f, 0x35, 0x47, 0x10, 0x08, 0x32, 0x67, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x32, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x63, 0x61, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x38,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x61,
	0x6b, 0x65, 0x73, 0x69, 0x2f, 0x6c, 0x7a, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_localdevice_network_proto_rawDescOnce sync.Once
	file_localdevice_network_proto_rawDescData = file_localdevice_network_proto_rawDesc
)

func file_localdevice_network_proto_rawDescGZIP() []byte {
	file_localdevice_network_proto_rawDescOnce.Do(func() {
		file_localdevice_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_localdevice_network_proto_rawDescData)
	})
	return file_localdevice_network_proto_rawDescData
}

var file_localdevice_network_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_localdevice_network_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_localdevice_network_proto_goTypes = []interface{}{
	(ConnectionType)(0),        // 0: cloud.lazycat.apis.localdevice.ConnectionType
	(*NetworkInformation)(nil), // 1: cloud.lazycat.apis.localdevice.NetworkInformation
	(*emptypb.Empty)(nil),      // 2: google.protobuf.Empty
}
var file_localdevice_network_proto_depIdxs = []int32{
	0, // 0: cloud.lazycat.apis.localdevice.NetworkInformation.ctype:type_name -> cloud.lazycat.apis.localdevice.ConnectionType
	2, // 1: cloud.lazycat.apis.localdevice.NetworkManager.Query:input_type -> google.protobuf.Empty
	1, // 2: cloud.lazycat.apis.localdevice.NetworkManager.Query:output_type -> cloud.lazycat.apis.localdevice.NetworkInformation
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_localdevice_network_proto_init() }
func file_localdevice_network_proto_init() {
	if File_localdevice_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_localdevice_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInformation); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_localdevice_network_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_localdevice_network_proto_goTypes,
		DependencyIndexes: file_localdevice_network_proto_depIdxs,
		EnumInfos:         file_localdevice_network_proto_enumTypes,
		MessageInfos:      file_localdevice_network_proto_msgTypes,
	}.Build()
	File_localdevice_network_proto = out.File
	file_localdevice_network_proto_rawDesc = nil
	file_localdevice_network_proto_goTypes = nil
	file_localdevice_network_proto_depIdxs = nil
}

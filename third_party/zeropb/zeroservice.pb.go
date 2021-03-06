// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: zeroservice.proto

package zeropb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_zeroservice_proto protoreflect.FileDescriptor

var file_zeroservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x7a, 0x65, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x7a, 0x65, 0x72, 0x6f, 0x32, 0x0d, 0x0a, 0x0b, 0x5a, 0x65, 0x72,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x7a, 0x65,
	0x72, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_zeroservice_proto_goTypes = []interface{}{}
var file_zeroservice_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zeroservice_proto_init() }
func file_zeroservice_proto_init() {
	if File_zeroservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zeroservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zeroservice_proto_goTypes,
		DependencyIndexes: file_zeroservice_proto_depIdxs,
	}.Build()
	File_zeroservice_proto = out.File
	file_zeroservice_proto_rawDesc = nil
	file_zeroservice_proto_goTypes = nil
	file_zeroservice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ZeroServiceClient is the client API for ZeroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ZeroServiceClient interface {
}

type zeroServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewZeroServiceClient(cc grpc.ClientConnInterface) ZeroServiceClient {
	return &zeroServiceClient{cc}
}

// ZeroServiceServer is the server API for ZeroService service.
type ZeroServiceServer interface {
}

// UnimplementedZeroServiceServer can be embedded to have forward compatible implementations.
type UnimplementedZeroServiceServer struct {
}

func RegisterZeroServiceServer(s *grpc.Server, srv ZeroServiceServer) {
	s.RegisterService(&_ZeroService_serviceDesc, srv)
}

var _ZeroService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zero.ZeroService",
	HandlerType: (*ZeroServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "zeroservice.proto",
}

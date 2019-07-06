// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/grpc.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

type Request struct {
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	return file_grpc_grpc_proto_msgTypes[0].MessageOf(x)
}

func (m *Request) XXX_Methods() *protoiface.Methods {
	return file_grpc_grpc_proto_msgTypes[0].Methods()
}

// Deprecated: Use Request.ProtoReflect.Type instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	return file_grpc_grpc_proto_msgTypes[1].MessageOf(x)
}

func (m *Response) XXX_Methods() *protoiface.Methods {
	return file_grpc_grpc_proto_msgTypes[1].Methods()
}

// Deprecated: Use Response.ProtoReflect.Type instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{1}
}

var File_grpc_grpc_proto protoreflect.FileDescriptor

var file_grpc_grpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc9, 0x02,
	0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x0a, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x64, 0x6f, 0x77,
	0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0d, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x4c, 0x0a, 0x09, 0x62,
	0x69, 0x64, 0x69, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_grpc_proto_rawDescOnce sync.Once
	file_grpc_grpc_proto_rawDescData = file_grpc_grpc_proto_rawDesc
)

func file_grpc_grpc_proto_rawDescGZIP() []byte {
	file_grpc_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_grpc_proto_rawDescData)
	})
	return file_grpc_grpc_proto_rawDescData
}

var file_grpc_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_grpc_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: goproto.protoc.grpc.Request
	(*Response)(nil), // 1: goproto.protoc.grpc.Response
}
var file_grpc_grpc_proto_depIdxs = []int32{
	0, // goproto.protoc.grpc.test_service.unary_call:input_type -> goproto.protoc.grpc.Request
	0, // goproto.protoc.grpc.test_service.downstream_call:input_type -> goproto.protoc.grpc.Request
	0, // goproto.protoc.grpc.test_service.upstream_call:input_type -> goproto.protoc.grpc.Request
	0, // goproto.protoc.grpc.test_service.bidi_call:input_type -> goproto.protoc.grpc.Request
	1, // goproto.protoc.grpc.test_service.unary_call:output_type -> goproto.protoc.grpc.Response
	1, // goproto.protoc.grpc.test_service.downstream_call:output_type -> goproto.protoc.grpc.Response
	1, // goproto.protoc.grpc.test_service.upstream_call:output_type -> goproto.protoc.grpc.Response
	1, // goproto.protoc.grpc.test_service.bidi_call:output_type -> goproto.protoc.grpc.Response
	4, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_grpc_grpc_proto_init() }
func file_grpc_grpc_proto_init() {
	if File_grpc_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.sizeCache
			case 1:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.sizeCache
			case 1:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			RawDescriptor: file_grpc_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_grpc_proto_msgTypes,
	}.Build()
	File_grpc_grpc_proto = out.File
	file_grpc_grpc_proto_rawDesc = nil
	file_grpc_grpc_proto_goTypes = nil
	file_grpc_grpc_proto_depIdxs = nil
}

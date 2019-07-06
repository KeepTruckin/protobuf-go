// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/empty.proto

package emptypb

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

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_empty_proto_msgTypes[0].MessageOf(x)
}

func (m *Empty) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_empty_proto_msgTypes[0].Methods()
}

// Deprecated: Use Empty.ProtoReflect.Type instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_google_protobuf_empty_proto_rawDescGZIP(), []int{0}
}

func (*Empty) XXX_WellKnownType() string { return "Empty" }

var File_google_protobuf_empty_proto protoreflect.FileDescriptor

var file_google_protobuf_empty_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x7d, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x0a,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2,
	0x02, 0x03, 0x47, 0x50, 0x42, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_protobuf_empty_proto_rawDescOnce sync.Once
	file_google_protobuf_empty_proto_rawDescData = file_google_protobuf_empty_proto_rawDesc
)

func file_google_protobuf_empty_proto_rawDescGZIP() []byte {
	file_google_protobuf_empty_proto_rawDescOnce.Do(func() {
		file_google_protobuf_empty_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_empty_proto_rawDescData)
	})
	return file_google_protobuf_empty_proto_rawDescData
}

var file_google_protobuf_empty_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_protobuf_empty_proto_goTypes = []interface{}{
	(*Empty)(nil), // 0: google.protobuf.Empty
}
var file_google_protobuf_empty_proto_depIdxs = []int32{
	0, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_google_protobuf_empty_proto_init() }
func file_google_protobuf_empty_proto_init() {
	if File_google_protobuf_empty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_protobuf_empty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_google_protobuf_empty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_empty_proto_goTypes,
		DependencyIndexes: file_google_protobuf_empty_proto_depIdxs,
		MessageInfos:      file_google_protobuf_empty_proto_msgTypes,
	}.Build()
	File_google_protobuf_empty_proto = out.File
	file_google_protobuf_empty_proto_rawDesc = nil
	file_google_protobuf_empty_proto_goTypes = nil
	file_google_protobuf_empty_proto_depIdxs = nil
}

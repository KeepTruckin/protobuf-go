// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/source_context.proto

package sourcecontextpb

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

// `SourceContext` represents information about the source of a
// protobuf element, like the file in which it is defined.
type SourceContext struct {
	// The path-qualified name of the .proto file that contained the associated
	// protobuf element.  For example: `"google/protobuf/source_context.proto"`.
	FileName      string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SourceContext) Reset() {
	*x = SourceContext{}
}

func (x *SourceContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceContext) ProtoMessage() {}

func (x *SourceContext) ProtoReflect() protoreflect.Message {
	return file_google_protobuf_source_context_proto_msgTypes[0].MessageOf(x)
}

func (m *SourceContext) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_source_context_proto_msgTypes[0].Methods()
}

// Deprecated: Use SourceContext.ProtoReflect.Type instead.
func (*SourceContext) Descriptor() ([]byte, []int) {
	return file_google_protobuf_source_context_proto_rawDescGZIP(), []int{0}
}

func (x *SourceContext) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var File_google_protobuf_source_context_proto protoreflect.FileDescriptor

var file_google_protobuf_source_context_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x2c, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x8a, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x12, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x47, 0x50,
	0x42, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_protobuf_source_context_proto_rawDescOnce sync.Once
	file_google_protobuf_source_context_proto_rawDescData = file_google_protobuf_source_context_proto_rawDesc
)

func file_google_protobuf_source_context_proto_rawDescGZIP() []byte {
	file_google_protobuf_source_context_proto_rawDescOnce.Do(func() {
		file_google_protobuf_source_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_source_context_proto_rawDescData)
	})
	return file_google_protobuf_source_context_proto_rawDescData
}

var file_google_protobuf_source_context_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_protobuf_source_context_proto_goTypes = []interface{}{
	(*SourceContext)(nil), // 0: google.protobuf.SourceContext
}
var file_google_protobuf_source_context_proto_depIdxs = []int32{
	0, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_google_protobuf_source_context_proto_init() }
func file_google_protobuf_source_context_proto_init() {
	if File_google_protobuf_source_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_protobuf_source_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceContext); i {
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			RawDescriptor: file_google_protobuf_source_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_source_context_proto_goTypes,
		DependencyIndexes: file_google_protobuf_source_context_proto_depIdxs,
		MessageInfos:      file_google_protobuf_source_context_proto_msgTypes,
	}.Build()
	File_google_protobuf_source_context_proto = out.File
	file_google_protobuf_source_context_proto_rawDesc = nil
	file_google_protobuf_source_context_proto_goTypes = nil
	file_google_protobuf_source_context_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/b.proto

package import_public

import (
	sub "google.golang.org/protobuf/cmd/protoc-gen-go/testdata/import_public/sub"
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

type Local struct {
	M             *sub.M `protobuf:"bytes,1,opt,name=m" json:"m,omitempty"`
	E             *sub.E `protobuf:"varint,2,opt,name=e,enum=goproto.protoc.import_public.sub.E" json:"e,omitempty"`
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Local) Reset() {
	*x = Local{}
}

func (x *Local) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Local) ProtoMessage() {}

func (x *Local) ProtoReflect() protoreflect.Message {
	return file_import_public_b_proto_msgTypes[0].MessageOf(x)
}

func (m *Local) XXX_Methods() *protoiface.Methods {
	return file_import_public_b_proto_msgTypes[0].Methods()
}

// Deprecated: Use Local.ProtoReflect.Type instead.
func (*Local) Descriptor() ([]byte, []int) {
	return file_import_public_b_proto_rawDescGZIP(), []int{0}
}

func (x *Local) GetM() *sub.M {
	if x != nil {
		return x.M
	}
	return nil
}

func (x *Local) GetE() sub.E {
	if x != nil && x.E != nil {
		return *x.E
	}
	return sub.E_ZERO
}

var File_import_public_b_proto protoreflect.FileDescriptor

var file_import_public_b_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x1a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x75, 0x62, 0x2f, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6d, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x01, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x4d, 0x52, 0x01, 0x6d, 0x12, 0x31, 0x0a, 0x01,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x45, 0x52, 0x01, 0x65, 0x42,
	0x45, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
}

var (
	file_import_public_b_proto_rawDescOnce sync.Once
	file_import_public_b_proto_rawDescData = file_import_public_b_proto_rawDesc
)

func file_import_public_b_proto_rawDescGZIP() []byte {
	file_import_public_b_proto_rawDescOnce.Do(func() {
		file_import_public_b_proto_rawDescData = protoimpl.X.CompressGZIP(file_import_public_b_proto_rawDescData)
	})
	return file_import_public_b_proto_rawDescData
}

var file_import_public_b_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_import_public_b_proto_goTypes = []interface{}{
	(*Local)(nil), // 0: goproto.protoc.import_public.Local
	(*sub.M)(nil), // 1: goproto.protoc.import_public.sub.M
	(sub.E)(0),    // 2: goproto.protoc.import_public.sub.E
}
var file_import_public_b_proto_depIdxs = []int32{
	1, // goproto.protoc.import_public.Local.m:type_name -> goproto.protoc.import_public.sub.M
	2, // goproto.protoc.import_public.Local.e:type_name -> goproto.protoc.import_public.sub.E
	2, // starting offset of method output_type sub-list
	2, // starting offset of method input_type sub-list
	2, // starting offset of extension type_name sub-list
	2, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_import_public_b_proto_init() }
func file_import_public_b_proto_init() {
	if File_import_public_b_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_import_public_b_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Local); i {
			case 2:
				return &v.sizeCache
			case 3:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			RawDescriptor: file_import_public_b_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_import_public_b_proto_goTypes,
		DependencyIndexes: file_import_public_b_proto_depIdxs,
		MessageInfos:      file_import_public_b_proto_msgTypes,
	}.Build()
	File_import_public_b_proto = out.File
	file_import_public_b_proto_rawDesc = nil
	file_import_public_b_proto_goTypes = nil
	file_import_public_b_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: benchmarks.proto

package benchmarks

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

type BenchmarkDataset struct {
	// Name of the benchmark dataset.  This should be unique across all datasets.
	// Should only contain word characters: [a-zA-Z0-9_]
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Fully-qualified name of the protobuf message for this dataset.
	// It will be one of the messages defined benchmark_messages_proto2.proto
	// or benchmark_messages_proto3.proto.
	//
	// Implementations that do not support reflection can implement this with
	// an explicit "if/else" chain that lists every known message defined
	// in those files.
	MessageName string `protobuf:"bytes,2,opt,name=message_name,json=messageName,proto3" json:"message_name,omitempty"`
	// The payload(s) for this dataset.  They should be parsed or serialized
	// in sequence, in a loop, ie.
	//
	//  while (!benchmarkDone) {  // Benchmark runner decides when to exit.
	//    for (i = 0; i < benchmark.payload.length; i++) {
	//      parse(benchmark.payload[i])
	//    }
	//  }
	//
	// This is intended to let datasets include a variety of data to provide
	// potentially more realistic results than just parsing the same message
	// over and over.  A single message parsed repeatedly could yield unusually
	// good branch prediction performance.
	Payload       [][]byte `protobuf:"bytes,3,rep,name=payload,proto3" json:"payload,omitempty"`
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BenchmarkDataset) Reset() {
	*x = BenchmarkDataset{}
}

func (x *BenchmarkDataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BenchmarkDataset) ProtoMessage() {}

func (x *BenchmarkDataset) ProtoReflect() protoreflect.Message {
	return file_benchmarks_proto_msgTypes[0].MessageOf(x)
}

func (m *BenchmarkDataset) XXX_Methods() *protoiface.Methods {
	return file_benchmarks_proto_msgTypes[0].Methods()
}

// Deprecated: Use BenchmarkDataset.ProtoReflect.Type instead.
func (*BenchmarkDataset) Descriptor() ([]byte, []int) {
	return file_benchmarks_proto_rawDescGZIP(), []int{0}
}

func (x *BenchmarkDataset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BenchmarkDataset) GetMessageName() string {
	if x != nil {
		return x.MessageName
	}
	return ""
}

func (x *BenchmarkDataset) GetPayload() [][]byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_benchmarks_proto protoreflect.FileDescriptor

var file_benchmarks_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x22, 0x63,
	0x0a, 0x10, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x5b, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x73, 0x5a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_benchmarks_proto_rawDescOnce sync.Once
	file_benchmarks_proto_rawDescData = file_benchmarks_proto_rawDesc
)

func file_benchmarks_proto_rawDescGZIP() []byte {
	file_benchmarks_proto_rawDescOnce.Do(func() {
		file_benchmarks_proto_rawDescData = protoimpl.X.CompressGZIP(file_benchmarks_proto_rawDescData)
	})
	return file_benchmarks_proto_rawDescData
}

var file_benchmarks_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_benchmarks_proto_goTypes = []interface{}{
	(*BenchmarkDataset)(nil), // 0: benchmarks.BenchmarkDataset
}
var file_benchmarks_proto_depIdxs = []int32{
	0, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_benchmarks_proto_init() }
func file_benchmarks_proto_init() {
	if File_benchmarks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_benchmarks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BenchmarkDataset); i {
			case 3:
				return &v.sizeCache
			case 4:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			RawDescriptor: file_benchmarks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_benchmarks_proto_goTypes,
		DependencyIndexes: file_benchmarks_proto_depIdxs,
		MessageInfos:      file_benchmarks_proto_msgTypes,
	}.Build()
	File_benchmarks_proto = out.File
	file_benchmarks_proto_rawDesc = nil
	file_benchmarks_proto_goTypes = nil
	file_benchmarks_proto_depIdxs = nil
}

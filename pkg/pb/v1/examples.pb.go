// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: service/examples/pkg/proto/examples.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_examples_pkg_proto_examples_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_service_examples_pkg_proto_examples_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_service_examples_pkg_proto_examples_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_examples_pkg_proto_examples_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_examples_pkg_proto_examples_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_service_examples_pkg_proto_examples_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_service_examples_pkg_proto_examples_proto protoreflect.FileDescriptor

var file_service_examples_pkg_proto_examples_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x2e,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x68,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xf8,
	0x01, 0x0a, 0x08, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x04, 0x54,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1d, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x25, 0x2e, 0x67,
	0x6f, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x22, 0x00,
	0x12, 0x57, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x25, 0x2e,
	0x67, 0x6f, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_examples_pkg_proto_examples_proto_rawDescOnce sync.Once
	file_service_examples_pkg_proto_examples_proto_rawDescData = file_service_examples_pkg_proto_examples_proto_rawDesc
)

func file_service_examples_pkg_proto_examples_proto_rawDescGZIP() []byte {
	file_service_examples_pkg_proto_examples_proto_rawDescOnce.Do(func() {
		file_service_examples_pkg_proto_examples_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_examples_pkg_proto_examples_proto_rawDescData)
	})
	return file_service_examples_pkg_proto_examples_proto_rawDescData
}

var file_service_examples_pkg_proto_examples_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_examples_pkg_proto_examples_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: go.park.examples.v1.Empty
	(*Response)(nil),         // 1: go.park.examples.v1.Response
	(*FetchEchoRequest)(nil), // 2: go.park.examples.v1.FetchEchoRequest
	(*Echo)(nil),             // 3: go.park.examples.v1.Echo
	(*EchoResponse)(nil),     // 4: go.park.examples.v1.EchoResponse
}
var file_service_examples_pkg_proto_examples_proto_depIdxs = []int32{
	0, // 0: go.park.examples.v1.Examples.Test:input_type -> go.park.examples.v1.Empty
	2, // 1: go.park.examples.v1.Examples.ReadEcho:input_type -> go.park.examples.v1.FetchEchoRequest
	2, // 2: go.park.examples.v1.Examples.FetchEcho:input_type -> go.park.examples.v1.FetchEchoRequest
	1, // 3: go.park.examples.v1.Examples.Test:output_type -> go.park.examples.v1.Response
	3, // 4: go.park.examples.v1.Examples.ReadEcho:output_type -> go.park.examples.v1.Echo
	4, // 5: go.park.examples.v1.Examples.FetchEcho:output_type -> go.park.examples.v1.EchoResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_examples_pkg_proto_examples_proto_init() }
func file_service_examples_pkg_proto_examples_proto_init() {
	if File_service_examples_pkg_proto_examples_proto != nil {
		return
	}
	file_service_examples_pkg_proto_echo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_examples_pkg_proto_examples_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_service_examples_pkg_proto_examples_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_service_examples_pkg_proto_examples_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_examples_pkg_proto_examples_proto_goTypes,
		DependencyIndexes: file_service_examples_pkg_proto_examples_proto_depIdxs,
		MessageInfos:      file_service_examples_pkg_proto_examples_proto_msgTypes,
	}.Build()
	File_service_examples_pkg_proto_examples_proto = out.File
	file_service_examples_pkg_proto_examples_proto_rawDesc = nil
	file_service_examples_pkg_proto_examples_proto_goTypes = nil
	file_service_examples_pkg_proto_examples_proto_depIdxs = nil
}

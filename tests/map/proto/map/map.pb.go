// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: map.proto

package _map

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

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{0}
}

func (x *Foo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type MapMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strings  map[string]string `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Strings2 map[string]int32  `protobuf:"bytes,2,rep,name=strings2,proto3" json:"strings2,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *MapMessage) Reset() {
	*x = MapMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_map_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMessage) ProtoMessage() {}

func (x *MapMessage) ProtoReflect() protoreflect.Message {
	mi := &file_map_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapMessage.ProtoReflect.Descriptor instead.
func (*MapMessage) Descriptor() ([]byte, []int) {
	return file_map_proto_rawDescGZIP(), []int{1}
}

func (x *MapMessage) GetStrings() map[string]string {
	if x != nil {
		return x.Strings
	}
	return nil
}

func (x *MapMessage) GetStrings2() map[string]int32 {
	if x != nil {
		return x.Strings2
	}
	return nil
}

var File_map_proto protoreflect.FileDescriptor

var file_map_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x22, 0x19, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0xfa, 0x01, 0x0a,
	0x0a, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x32,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61,
	0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73,
	0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x32,
	0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_map_proto_rawDescOnce sync.Once
	file_map_proto_rawDescData = file_map_proto_rawDesc
)

func file_map_proto_rawDescGZIP() []byte {
	file_map_proto_rawDescOnce.Do(func() {
		file_map_proto_rawDescData = protoimpl.X.CompressGZIP(file_map_proto_rawDescData)
	})
	return file_map_proto_rawDescData
}

var file_map_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_map_proto_goTypes = []interface{}{
	(*Foo)(nil),        // 0: test.Foo
	(*MapMessage)(nil), // 1: test.MapMessage
	nil,                // 2: test.MapMessage.StringsEntry
	nil,                // 3: test.MapMessage.Strings2Entry
}
var file_map_proto_depIdxs = []int32{
	2, // 0: test.MapMessage.strings:type_name -> test.MapMessage.StringsEntry
	3, // 1: test.MapMessage.strings2:type_name -> test.MapMessage.Strings2Entry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_map_proto_init() }
func file_map_proto_init() {
	if File_map_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_map_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
		file_map_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapMessage); i {
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
			RawDescriptor: file_map_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_map_proto_goTypes,
		DependencyIndexes: file_map_proto_depIdxs,
		MessageInfos:      file_map_proto_msgTypes,
	}.Build()
	File_map_proto = out.File
	file_map_proto_rawDesc = nil
	file_map_proto_goTypes = nil
	file_map_proto_depIdxs = nil
}

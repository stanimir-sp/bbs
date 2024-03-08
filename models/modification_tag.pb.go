// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: modification_tag.proto

package models

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

type ModificationTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Epoch string `protobuf:"bytes,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Index uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *ModificationTag) Reset() {
	*x = ModificationTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modification_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModificationTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModificationTag) ProtoMessage() {}

func (x *ModificationTag) ProtoReflect() protoreflect.Message {
	mi := &file_modification_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModificationTag.ProtoReflect.Descriptor instead.
func (*ModificationTag) Descriptor() ([]byte, []int) {
	return file_modification_tag_proto_rawDescGZIP(), []int{0}
}

func (x *ModificationTag) GetEpoch() string {
	if x != nil {
		return x.Epoch
	}
	return ""
}

func (x *ModificationTag) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_modification_tag_proto protoreflect.FileDescriptor

var file_modification_tag_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x22, 0x3d, 0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42,
	0x22, 0x5a, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x72, 0x79, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x62, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modification_tag_proto_rawDescOnce sync.Once
	file_modification_tag_proto_rawDescData = file_modification_tag_proto_rawDesc
)

func file_modification_tag_proto_rawDescGZIP() []byte {
	file_modification_tag_proto_rawDescOnce.Do(func() {
		file_modification_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_modification_tag_proto_rawDescData)
	})
	return file_modification_tag_proto_rawDescData
}

var file_modification_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_modification_tag_proto_goTypes = []interface{}{
	(*ModificationTag)(nil), // 0: models.ModificationTag
}
var file_modification_tag_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modification_tag_proto_init() }
func file_modification_tag_proto_init() {
	if File_modification_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modification_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModificationTag); i {
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
			RawDescriptor: file_modification_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_modification_tag_proto_goTypes,
		DependencyIndexes: file_modification_tag_proto_depIdxs,
		MessageInfos:      file_modification_tag_proto_msgTypes,
	}.Build()
	File_modification_tag_proto = out.File
	file_modification_tag_proto_rawDesc = nil
	file_modification_tag_proto_goTypes = nil
	file_modification_tag_proto_depIdxs = nil
}

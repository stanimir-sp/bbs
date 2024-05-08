// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: sidecar.proto

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

type ProtoSidecar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action   *ProtoAction `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	DiskMb   int32        `protobuf:"varint,2,opt,name=disk_mb,proto3" json:"disk_mb,omitempty"`
	MemoryMb int32        `protobuf:"varint,3,opt,name=memory_mb,proto3" json:"memory_mb,omitempty"`
}

func (x *ProtoSidecar) Reset() {
	*x = ProtoSidecar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sidecar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoSidecar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoSidecar) ProtoMessage() {}

func (x *ProtoSidecar) ProtoReflect() protoreflect.Message {
	mi := &file_sidecar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoSidecar.ProtoReflect.Descriptor instead.
func (*ProtoSidecar) Descriptor() ([]byte, []int) {
	return file_sidecar_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoSidecar) GetAction() *ProtoAction {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *ProtoSidecar) GetDiskMb() int32 {
	if x != nil {
		return x.DiskMb
	}
	return 0
}

func (x *ProtoSidecar) GetMemoryMb() int32 {
	if x != nil {
		return x.MemoryMb
	}
	return 0
}

var File_sidecar_proto protoreflect.FileDescriptor

var file_sidecar_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53,
	0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x12, 0x2b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x62, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x62, 0x42, 0x22, 0x5a, 0x20, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x62, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sidecar_proto_rawDescOnce sync.Once
	file_sidecar_proto_rawDescData = file_sidecar_proto_rawDesc
)

func file_sidecar_proto_rawDescGZIP() []byte {
	file_sidecar_proto_rawDescOnce.Do(func() {
		file_sidecar_proto_rawDescData = protoimpl.X.CompressGZIP(file_sidecar_proto_rawDescData)
	})
	return file_sidecar_proto_rawDescData
}

var file_sidecar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sidecar_proto_goTypes = []interface{}{
	(*ProtoSidecar)(nil), // 0: models.ProtoSidecar
	(*ProtoAction)(nil),  // 1: models.ProtoAction
}
var file_sidecar_proto_depIdxs = []int32{
	1, // 0: models.ProtoSidecar.action:type_name -> models.ProtoAction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sidecar_proto_init() }
func file_sidecar_proto_init() {
	if File_sidecar_proto != nil {
		return
	}
	file_actions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sidecar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoSidecar); i {
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
			RawDescriptor: file_sidecar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sidecar_proto_goTypes,
		DependencyIndexes: file_sidecar_proto_depIdxs,
		MessageInfos:      file_sidecar_proto_msgTypes,
	}.Build()
	File_sidecar_proto = out.File
	file_sidecar_proto_rawDesc = nil
	file_sidecar_proto_goTypes = nil
	file_sidecar_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.2
// source: security_group.proto

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

type ProtoPortRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Start         uint32                 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End           uint32                 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProtoPortRange) Reset() {
	*x = ProtoPortRange{}
	mi := &file_security_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoPortRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoPortRange) ProtoMessage() {}

func (x *ProtoPortRange) ProtoReflect() protoreflect.Message {
	mi := &file_security_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoPortRange.ProtoReflect.Descriptor instead.
func (*ProtoPortRange) Descriptor() ([]byte, []int) {
	return file_security_group_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoPortRange) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ProtoPortRange) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

type ProtoICMPInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          int32                  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Code          int32                  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProtoICMPInfo) Reset() {
	*x = ProtoICMPInfo{}
	mi := &file_security_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoICMPInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoICMPInfo) ProtoMessage() {}

func (x *ProtoICMPInfo) ProtoReflect() protoreflect.Message {
	mi := &file_security_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoICMPInfo.ProtoReflect.Descriptor instead.
func (*ProtoICMPInfo) Descriptor() ([]byte, []int) {
	return file_security_group_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoICMPInfo) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ProtoICMPInfo) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type ProtoSecurityGroupRule struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Protocol      string                 `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Destinations  []string               `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	Ports         []uint32               `protobuf:"varint,3,rep,name=ports,proto3" json:"ports,omitempty"`
	PortRange     *ProtoPortRange        `protobuf:"bytes,4,opt,name=port_range,proto3" json:"port_range,omitempty"`
	IcmpInfo      *ProtoICMPInfo         `protobuf:"bytes,5,opt,name=icmp_info,proto3" json:"icmp_info,omitempty"`
	Log           bool                   `protobuf:"varint,6,opt,name=log,proto3" json:"log,omitempty"`
	Annotations   []string               `protobuf:"bytes,7,rep,name=annotations,proto3" json:"annotations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProtoSecurityGroupRule) Reset() {
	*x = ProtoSecurityGroupRule{}
	mi := &file_security_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoSecurityGroupRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoSecurityGroupRule) ProtoMessage() {}

func (x *ProtoSecurityGroupRule) ProtoReflect() protoreflect.Message {
	mi := &file_security_group_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoSecurityGroupRule.ProtoReflect.Descriptor instead.
func (*ProtoSecurityGroupRule) Descriptor() ([]byte, []int) {
	return file_security_group_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoSecurityGroupRule) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ProtoSecurityGroupRule) GetDestinations() []string {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *ProtoSecurityGroupRule) GetPorts() []uint32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *ProtoSecurityGroupRule) GetPortRange() *ProtoPortRange {
	if x != nil {
		return x.PortRange
	}
	return nil
}

func (x *ProtoSecurityGroupRule) GetIcmpInfo() *ProtoICMPInfo {
	if x != nil {
		return x.IcmpInfo
	}
	return nil
}

func (x *ProtoSecurityGroupRule) GetLog() bool {
	if x != nil {
		return x.Log
	}
	return false
}

func (x *ProtoSecurityGroupRule) GetAnnotations() []string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_security_group_proto protoreflect.FileDescriptor

var file_security_group_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x09,
	0x62, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x41, 0x0a,
	0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x43, 0x4d, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xc0, 0x3e,
	0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x98, 0x02, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x05, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x00, 0x52, 0x05,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x33, 0x0a,
	0x09, 0x69, 0x63, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x49,
	0x43, 0x4d, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x69, 0x63, 0x6d, 0x70, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x15, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x03, 0xc0, 0x3e, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x22, 0x5a, 0x20, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x62, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_security_group_proto_rawDescOnce sync.Once
	file_security_group_proto_rawDescData = file_security_group_proto_rawDesc
)

func file_security_group_proto_rawDescGZIP() []byte {
	file_security_group_proto_rawDescOnce.Do(func() {
		file_security_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_group_proto_rawDescData)
	})
	return file_security_group_proto_rawDescData
}

var file_security_group_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_security_group_proto_goTypes = []any{
	(*ProtoPortRange)(nil),         // 0: models.ProtoPortRange
	(*ProtoICMPInfo)(nil),          // 1: models.ProtoICMPInfo
	(*ProtoSecurityGroupRule)(nil), // 2: models.ProtoSecurityGroupRule
}
var file_security_group_proto_depIdxs = []int32{
	0, // 0: models.ProtoSecurityGroupRule.port_range:type_name -> models.ProtoPortRange
	1, // 1: models.ProtoSecurityGroupRule.icmp_info:type_name -> models.ProtoICMPInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_security_group_proto_init() }
func file_security_group_proto_init() {
	if File_security_group_proto != nil {
		return
	}
	file_bbs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_security_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_group_proto_goTypes,
		DependencyIndexes: file_security_group_proto_depIdxs,
		MessageInfos:      file_security_group_proto_msgTypes,
	}.Build()
	File_security_group_proto = out.File
	file_security_group_proto_rawDesc = nil
	file_security_group_proto_goTypes = nil
	file_security_group_proto_depIdxs = nil
}

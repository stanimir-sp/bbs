// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
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

type PortRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start uint32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   uint32 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *PortRange) Reset() {
	*x = PortRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortRange) ProtoMessage() {}

func (x *PortRange) ProtoReflect() protoreflect.Message {
	mi := &file_security_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortRange.ProtoReflect.Descriptor instead.
func (*PortRange) Descriptor() ([]byte, []int) {
	return file_security_group_proto_rawDescGZIP(), []int{0}
}

func (x *PortRange) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *PortRange) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

type ICMPInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ICMPInfo) Reset() {
	*x = ICMPInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ICMPInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ICMPInfo) ProtoMessage() {}

func (x *ICMPInfo) ProtoReflect() protoreflect.Message {
	mi := &file_security_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ICMPInfo.ProtoReflect.Descriptor instead.
func (*ICMPInfo) Descriptor() ([]byte, []int) {
	return file_security_group_proto_rawDescGZIP(), []int{1}
}

func (x *ICMPInfo) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ICMPInfo) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type SecurityGroupRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol     string     `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Destinations []string   `protobuf:"bytes,2,rep,name=destinations,proto3" json:"destinations,omitempty"`
	Ports        []uint32   `protobuf:"varint,3,rep,name=ports,proto3" json:"ports,omitempty"`
	PortRange    *PortRange `protobuf:"bytes,4,opt,name=port_range,json=portRange,proto3" json:"port_range,omitempty"`
	IcmpInfo     *ICMPInfo  `protobuf:"bytes,5,opt,name=icmp_info,json=icmpInfo,proto3" json:"icmp_info,omitempty"`
	Log          bool       `protobuf:"varint,6,opt,name=log,proto3" json:"log,omitempty"`
	Annotations  []string   `protobuf:"bytes,7,rep,name=annotations,proto3" json:"annotations,omitempty"`
}

func (x *SecurityGroupRule) Reset() {
	*x = SecurityGroupRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityGroupRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityGroupRule) ProtoMessage() {}

func (x *SecurityGroupRule) ProtoReflect() protoreflect.Message {
	mi := &file_security_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityGroupRule.ProtoReflect.Descriptor instead.
func (*SecurityGroupRule) Descriptor() ([]byte, []int) {
	return file_security_group_proto_rawDescGZIP(), []int{2}
}

func (x *SecurityGroupRule) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *SecurityGroupRule) GetDestinations() []string {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *SecurityGroupRule) GetPorts() []uint32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *SecurityGroupRule) GetPortRange() *PortRange {
	if x != nil {
		return x.PortRange
	}
	return nil
}

func (x *SecurityGroupRule) GetIcmpInfo() *ICMPInfo {
	if x != nil {
		return x.IcmpInfo
	}
	return nil
}

func (x *SecurityGroupRule) GetLog() bool {
	if x != nil {
		return x.Log
	}
	return false
}

func (x *SecurityGroupRule) GetAnnotations() []string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_security_group_proto protoreflect.FileDescriptor

var file_security_group_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x33,
	0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x22, 0x32, 0x0a, 0x08, 0x49, 0x43, 0x4d, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x82, 0x02, 0x0a, 0x11, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a,
	0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x00,
	0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x69, 0x63, 0x6d,
	0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x43, 0x4d, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x69, 0x63, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x24, 0x5a, 0x22,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x2f, 0x62, 0x62, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_security_group_proto_goTypes = []interface{}{
	(*PortRange)(nil),         // 0: models.PortRange
	(*ICMPInfo)(nil),          // 1: models.ICMPInfo
	(*SecurityGroupRule)(nil), // 2: models.SecurityGroupRule
}
var file_security_group_proto_depIdxs = []int32{
	0, // 0: models.SecurityGroupRule.port_range:type_name -> models.PortRange
	1, // 1: models.SecurityGroupRule.icmp_info:type_name -> models.ICMPInfo
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
	if !protoimpl.UnsafeEnabled {
		file_security_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortRange); i {
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
		file_security_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ICMPInfo); i {
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
		file_security_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityGroupRule); i {
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

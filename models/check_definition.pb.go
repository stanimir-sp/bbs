// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: check_definition.proto

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

type ProtoCheckDefinition struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Checks          []*ProtoCheck          `protobuf:"bytes,1,rep,name=checks,proto3" json:"checks,omitempty"`
	LogSource       string                 `protobuf:"bytes,2,opt,name=log_source,proto3" json:"log_source,omitempty"`
	ReadinessChecks []*ProtoCheck          `protobuf:"bytes,3,rep,name=readiness_checks,proto3" json:"readiness_checks,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ProtoCheckDefinition) Reset() {
	*x = ProtoCheckDefinition{}
	mi := &file_check_definition_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoCheckDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoCheckDefinition) ProtoMessage() {}

func (x *ProtoCheckDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_check_definition_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoCheckDefinition.ProtoReflect.Descriptor instead.
func (*ProtoCheckDefinition) Descriptor() ([]byte, []int) {
	return file_check_definition_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoCheckDefinition) GetChecks() []*ProtoCheck {
	if x != nil {
		return x.Checks
	}
	return nil
}

func (x *ProtoCheckDefinition) GetLogSource() string {
	if x != nil {
		return x.LogSource
	}
	return ""
}

func (x *ProtoCheckDefinition) GetReadinessChecks() []*ProtoCheck {
	if x != nil {
		return x.ReadinessChecks
	}
	return nil
}

type ProtoCheck struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// oneof is hard to use right now, instead we can do this check in validation
	// oneof check {
	TcpCheck      *ProtoTCPCheck  `protobuf:"bytes,1,opt,name=tcp_check,proto3" json:"tcp_check,omitempty"`
	HttpCheck     *ProtoHTTPCheck `protobuf:"bytes,2,opt,name=http_check,proto3" json:"http_check,omitempty"` // }
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProtoCheck) Reset() {
	*x = ProtoCheck{}
	mi := &file_check_definition_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoCheck) ProtoMessage() {}

func (x *ProtoCheck) ProtoReflect() protoreflect.Message {
	mi := &file_check_definition_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoCheck.ProtoReflect.Descriptor instead.
func (*ProtoCheck) Descriptor() ([]byte, []int) {
	return file_check_definition_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoCheck) GetTcpCheck() *ProtoTCPCheck {
	if x != nil {
		return x.TcpCheck
	}
	return nil
}

func (x *ProtoCheck) GetHttpCheck() *ProtoHTTPCheck {
	if x != nil {
		return x.HttpCheck
	}
	return nil
}

type ProtoTCPCheck struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Port             uint32                 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	ConnectTimeoutMs uint64                 `protobuf:"varint,2,opt,name=connect_timeout_ms,proto3" json:"connect_timeout_ms,omitempty"`
	IntervalMs       uint64                 `protobuf:"varint,3,opt,name=interval_ms,proto3" json:"interval_ms,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ProtoTCPCheck) Reset() {
	*x = ProtoTCPCheck{}
	mi := &file_check_definition_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoTCPCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoTCPCheck) ProtoMessage() {}

func (x *ProtoTCPCheck) ProtoReflect() protoreflect.Message {
	mi := &file_check_definition_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoTCPCheck.ProtoReflect.Descriptor instead.
func (*ProtoTCPCheck) Descriptor() ([]byte, []int) {
	return file_check_definition_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoTCPCheck) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ProtoTCPCheck) GetConnectTimeoutMs() uint64 {
	if x != nil {
		return x.ConnectTimeoutMs
	}
	return 0
}

func (x *ProtoTCPCheck) GetIntervalMs() uint64 {
	if x != nil {
		return x.IntervalMs
	}
	return 0
}

type ProtoHTTPCheck struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Port             uint32                 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	RequestTimeoutMs uint64                 `protobuf:"varint,2,opt,name=request_timeout_ms,proto3" json:"request_timeout_ms,omitempty"`
	Path             string                 `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	IntervalMs       uint64                 `protobuf:"varint,4,opt,name=interval_ms,proto3" json:"interval_ms,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ProtoHTTPCheck) Reset() {
	*x = ProtoHTTPCheck{}
	mi := &file_check_definition_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoHTTPCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoHTTPCheck) ProtoMessage() {}

func (x *ProtoHTTPCheck) ProtoReflect() protoreflect.Message {
	mi := &file_check_definition_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoHTTPCheck.ProtoReflect.Descriptor instead.
func (*ProtoHTTPCheck) Descriptor() ([]byte, []int) {
	return file_check_definition_proto_rawDescGZIP(), []int{3}
}

func (x *ProtoHTTPCheck) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ProtoHTTPCheck) GetRequestTimeoutMs() uint64 {
	if x != nil {
		return x.RequestTimeoutMs
	}
	return 0
}

func (x *ProtoHTTPCheck) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ProtoHTTPCheck) GetIntervalMs() uint64 {
	if x != nil {
		return x.IntervalMs
	}
	return 0
}

var File_check_definition_proto protoreflect.FileDescriptor

var file_check_definition_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x1a, 0x09, 0x62, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x14,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73,
	0x12, 0x23, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x10, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x10, 0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x22, 0x79, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x63, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x43, 0x50, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x09, 0x74,
	0x63, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x36, 0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x54, 0x54, 0x50, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x22, 0x7a, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x43, 0x50, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x03, 0xc0, 0x3e, 0x01, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x73, 0x22, 0x94, 0x01, 0x0a,
	0x0e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x54, 0x54, 0x50, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x17, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x03, 0xc0,
	0x3e, 0x01, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x5f, 0x6d, 0x73, 0x42, 0x22, 0x5a, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x62, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_check_definition_proto_rawDescOnce sync.Once
	file_check_definition_proto_rawDescData = file_check_definition_proto_rawDesc
)

func file_check_definition_proto_rawDescGZIP() []byte {
	file_check_definition_proto_rawDescOnce.Do(func() {
		file_check_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_definition_proto_rawDescData)
	})
	return file_check_definition_proto_rawDescData
}

var file_check_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_check_definition_proto_goTypes = []any{
	(*ProtoCheckDefinition)(nil), // 0: models.ProtoCheckDefinition
	(*ProtoCheck)(nil),           // 1: models.ProtoCheck
	(*ProtoTCPCheck)(nil),        // 2: models.ProtoTCPCheck
	(*ProtoHTTPCheck)(nil),       // 3: models.ProtoHTTPCheck
}
var file_check_definition_proto_depIdxs = []int32{
	1, // 0: models.ProtoCheckDefinition.checks:type_name -> models.ProtoCheck
	1, // 1: models.ProtoCheckDefinition.readiness_checks:type_name -> models.ProtoCheck
	2, // 2: models.ProtoCheck.tcp_check:type_name -> models.ProtoTCPCheck
	3, // 3: models.ProtoCheck.http_check:type_name -> models.ProtoHTTPCheck
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_check_definition_proto_init() }
func file_check_definition_proto_init() {
	if File_check_definition_proto != nil {
		return
	}
	file_bbs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_check_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_check_definition_proto_goTypes,
		DependencyIndexes: file_check_definition_proto_depIdxs,
		MessageInfos:      file_check_definition_proto_msgTypes,
	}.Build()
	File_check_definition_proto = out.File
	file_check_definition_proto_rawDesc = nil
	file_check_definition_proto_goTypes = nil
	file_check_definition_proto_depIdxs = nil
}

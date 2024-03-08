// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: domain.proto

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

type DomainsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Domains []string `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
}

func (x *DomainsResponse) Reset() {
	*x = DomainsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainsResponse) ProtoMessage() {}

func (x *DomainsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainsResponse.ProtoReflect.Descriptor instead.
func (*DomainsResponse) Descriptor() ([]byte, []int) {
	return file_domain_proto_rawDescGZIP(), []int{0}
}

func (x *DomainsResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *DomainsResponse) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

type UpsertDomainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpsertDomainResponse) Reset() {
	*x = UpsertDomainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertDomainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertDomainResponse) ProtoMessage() {}

func (x *UpsertDomainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertDomainResponse.ProtoReflect.Descriptor instead.
func (*UpsertDomainResponse) Descriptor() ([]byte, []int) {
	return file_domain_proto_rawDescGZIP(), []int{1}
}

func (x *UpsertDomainResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type UpsertDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Ttl    uint32 `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *UpsertDomainRequest) Reset() {
	*x = UpsertDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertDomainRequest) ProtoMessage() {}

func (x *UpsertDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_domain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertDomainRequest.ProtoReflect.Descriptor instead.
func (*UpsertDomainRequest) Descriptor() ([]byte, []int) {
	return file_domain_proto_rawDescGZIP(), []int{2}
}

func (x *UpsertDomainRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *UpsertDomainRequest) GetTtl() uint32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

var File_domain_proto protoreflect.FileDescriptor

var file_domain_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x3b, 0x0a, 0x14, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x13, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x74, 0x74, 0x6c, 0x42, 0x22, 0x5a, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x62, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_proto_rawDescOnce sync.Once
	file_domain_proto_rawDescData = file_domain_proto_rawDesc
)

func file_domain_proto_rawDescGZIP() []byte {
	file_domain_proto_rawDescOnce.Do(func() {
		file_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_proto_rawDescData)
	})
	return file_domain_proto_rawDescData
}

var file_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_domain_proto_goTypes = []interface{}{
	(*DomainsResponse)(nil),      // 0: models.DomainsResponse
	(*UpsertDomainResponse)(nil), // 1: models.UpsertDomainResponse
	(*UpsertDomainRequest)(nil),  // 2: models.UpsertDomainRequest
	(*Error)(nil),                // 3: models.Error
}
var file_domain_proto_depIdxs = []int32{
	3, // 0: models.DomainsResponse.error:type_name -> models.Error
	3, // 1: models.UpsertDomainResponse.error:type_name -> models.Error
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_domain_proto_init() }
func file_domain_proto_init() {
	if File_domain_proto != nil {
		return
	}
	file_error_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainsResponse); i {
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
		file_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertDomainResponse); i {
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
		file_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertDomainRequest); i {
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
			RawDescriptor: file_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_proto_goTypes,
		DependencyIndexes: file_domain_proto_depIdxs,
		MessageInfos:      file_domain_proto_msgTypes,
	}.Build()
	File_domain_proto = out.File
	file_domain_proto_rawDesc = nil
	file_domain_proto_goTypes = nil
	file_domain_proto_depIdxs = nil
}

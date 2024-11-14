// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: task.proto

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

type ProtoTask_State int32

const (
	ProtoTask_Invalid   ProtoTask_State = 0
	ProtoTask_Pending   ProtoTask_State = 1
	ProtoTask_Running   ProtoTask_State = 2
	ProtoTask_Completed ProtoTask_State = 3
	ProtoTask_Resolving ProtoTask_State = 4
)

// Enum value maps for ProtoTask_State.
var (
	ProtoTask_State_name = map[int32]string{
		0: "Invalid",
		1: "Pending",
		2: "Running",
		3: "Completed",
		4: "Resolving",
	}
	ProtoTask_State_value = map[string]int32{
		"Invalid":   0,
		"Pending":   1,
		"Running":   2,
		"Completed": 3,
		"Resolving": 4,
	}
)

func (x ProtoTask_State) Enum() *ProtoTask_State {
	p := new(ProtoTask_State)
	*p = x
	return p
}

func (x ProtoTask_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoTask_State) Descriptor() protoreflect.EnumDescriptor {
	return file_task_proto_enumTypes[0].Descriptor()
}

func (ProtoTask_State) Type() protoreflect.EnumType {
	return &file_task_proto_enumTypes[0]
}

func (x ProtoTask_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtoTask_State.Descriptor instead.
func (ProtoTask_State) EnumDescriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1, 0}
}

type ProtoTaskDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RootFs                string                      `protobuf:"bytes,1,opt,name=root_fs,json=rootfs,proto3" json:"root_fs,omitempty"`
	EnvironmentVariables  []*ProtoEnvironmentVariable `protobuf:"bytes,2,rep,name=environment_variables,json=env,proto3" json:"environment_variables,omitempty"`
	Action                *ProtoAction                `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	DiskMb                int32                       `protobuf:"varint,4,opt,name=disk_mb,proto3" json:"disk_mb,omitempty"`
	MemoryMb              int32                       `protobuf:"varint,5,opt,name=memory_mb,proto3" json:"memory_mb,omitempty"`
	CpuWeight             uint32                      `protobuf:"varint,6,opt,name=cpu_weight,proto3" json:"cpu_weight,omitempty"`
	Privileged            bool                        `protobuf:"varint,7,opt,name=privileged,proto3" json:"privileged,omitempty"`
	LogSource             string                      `protobuf:"bytes,8,opt,name=log_source,proto3" json:"log_source,omitempty"`
	LogGuid               string                      `protobuf:"bytes,9,opt,name=log_guid,proto3" json:"log_guid,omitempty"`
	MetricsGuid           string                      `protobuf:"bytes,10,opt,name=metrics_guid,proto3" json:"metrics_guid,omitempty"`
	ResultFile            string                      `protobuf:"bytes,11,opt,name=result_file,proto3" json:"result_file,omitempty"`
	CompletionCallbackUrl string                      `protobuf:"bytes,12,opt,name=completion_callback_url,proto3" json:"completion_callback_url,omitempty"`
	Annotation            string                      `protobuf:"bytes,13,opt,name=annotation,proto3" json:"annotation,omitempty"`
	EgressRules           []*ProtoSecurityGroupRule   `protobuf:"bytes,14,rep,name=egress_rules,proto3" json:"egress_rules,omitempty"`
	CachedDependencies    []*ProtoCachedDependency    `protobuf:"bytes,15,rep,name=cached_dependencies,proto3" json:"cached_dependencies,omitempty"`
	// Deprecated: Marked as deprecated in task.proto.
	LegacyDownloadUser            string                          `protobuf:"bytes,16,opt,name=legacy_download_user,proto3" json:"legacy_download_user,omitempty"`
	TrustedSystemCertificatesPath string                          `protobuf:"bytes,17,opt,name=trusted_system_certificates_path,proto3" json:"trusted_system_certificates_path,omitempty"`
	VolumeMounts                  []*ProtoVolumeMount             `protobuf:"bytes,18,rep,name=volume_mounts,proto3" json:"volume_mounts,omitempty"`
	Network                       *ProtoNetwork                   `protobuf:"bytes,19,opt,name=network,proto3" json:"network,omitempty"`
	PlacementTags                 []string                        `protobuf:"bytes,20,rep,name=placement_tags,proto3" json:"placement_tags,omitempty"`
	MaxPids                       int32                           `protobuf:"varint,21,opt,name=max_pids,proto3" json:"max_pids,omitempty"`
	CertificateProperties         *ProtoCertificateProperties     `protobuf:"bytes,22,opt,name=certificate_properties,proto3" json:"certificate_properties,omitempty"`
	ImageUsername                 string                          `protobuf:"bytes,23,opt,name=image_username,proto3" json:"image_username,omitempty"`
	ImagePassword                 string                          `protobuf:"bytes,24,opt,name=image_password,proto3" json:"image_password,omitempty"`
	ImageLayers                   []*ProtoImageLayer              `protobuf:"bytes,25,rep,name=image_layers,proto3" json:"image_layers,omitempty"`
	LogRateLimit                  *ProtoLogRateLimit              `protobuf:"bytes,26,opt,name=log_rate_limit,proto3" json:"log_rate_limit,omitempty"`
	MetricTags                    map[string]*ProtoMetricTagValue `protobuf:"bytes,27,rep,name=metric_tags,proto3" json:"metric_tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProtoTaskDefinition) Reset() {
	*x = ProtoTaskDefinition{}
	mi := &file_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoTaskDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoTaskDefinition) ProtoMessage() {}

func (x *ProtoTaskDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoTaskDefinition.ProtoReflect.Descriptor instead.
func (*ProtoTaskDefinition) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoTaskDefinition) GetRootFs() string {
	if x != nil {
		return x.RootFs
	}
	return ""
}

func (x *ProtoTaskDefinition) GetEnvironmentVariables() []*ProtoEnvironmentVariable {
	if x != nil {
		return x.EnvironmentVariables
	}
	return nil
}

func (x *ProtoTaskDefinition) GetAction() *ProtoAction {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *ProtoTaskDefinition) GetDiskMb() int32 {
	if x != nil {
		return x.DiskMb
	}
	return 0
}

func (x *ProtoTaskDefinition) GetMemoryMb() int32 {
	if x != nil {
		return x.MemoryMb
	}
	return 0
}

func (x *ProtoTaskDefinition) GetCpuWeight() uint32 {
	if x != nil {
		return x.CpuWeight
	}
	return 0
}

func (x *ProtoTaskDefinition) GetPrivileged() bool {
	if x != nil {
		return x.Privileged
	}
	return false
}

func (x *ProtoTaskDefinition) GetLogSource() string {
	if x != nil {
		return x.LogSource
	}
	return ""
}

func (x *ProtoTaskDefinition) GetLogGuid() string {
	if x != nil {
		return x.LogGuid
	}
	return ""
}

func (x *ProtoTaskDefinition) GetMetricsGuid() string {
	if x != nil {
		return x.MetricsGuid
	}
	return ""
}

func (x *ProtoTaskDefinition) GetResultFile() string {
	if x != nil {
		return x.ResultFile
	}
	return ""
}

func (x *ProtoTaskDefinition) GetCompletionCallbackUrl() string {
	if x != nil {
		return x.CompletionCallbackUrl
	}
	return ""
}

func (x *ProtoTaskDefinition) GetAnnotation() string {
	if x != nil {
		return x.Annotation
	}
	return ""
}

func (x *ProtoTaskDefinition) GetEgressRules() []*ProtoSecurityGroupRule {
	if x != nil {
		return x.EgressRules
	}
	return nil
}

func (x *ProtoTaskDefinition) GetCachedDependencies() []*ProtoCachedDependency {
	if x != nil {
		return x.CachedDependencies
	}
	return nil
}

// Deprecated: Marked as deprecated in task.proto.
func (x *ProtoTaskDefinition) GetLegacyDownloadUser() string {
	if x != nil {
		return x.LegacyDownloadUser
	}
	return ""
}

func (x *ProtoTaskDefinition) GetTrustedSystemCertificatesPath() string {
	if x != nil {
		return x.TrustedSystemCertificatesPath
	}
	return ""
}

func (x *ProtoTaskDefinition) GetVolumeMounts() []*ProtoVolumeMount {
	if x != nil {
		return x.VolumeMounts
	}
	return nil
}

func (x *ProtoTaskDefinition) GetNetwork() *ProtoNetwork {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *ProtoTaskDefinition) GetPlacementTags() []string {
	if x != nil {
		return x.PlacementTags
	}
	return nil
}

func (x *ProtoTaskDefinition) GetMaxPids() int32 {
	if x != nil {
		return x.MaxPids
	}
	return 0
}

func (x *ProtoTaskDefinition) GetCertificateProperties() *ProtoCertificateProperties {
	if x != nil {
		return x.CertificateProperties
	}
	return nil
}

func (x *ProtoTaskDefinition) GetImageUsername() string {
	if x != nil {
		return x.ImageUsername
	}
	return ""
}

func (x *ProtoTaskDefinition) GetImagePassword() string {
	if x != nil {
		return x.ImagePassword
	}
	return ""
}

func (x *ProtoTaskDefinition) GetImageLayers() []*ProtoImageLayer {
	if x != nil {
		return x.ImageLayers
	}
	return nil
}

func (x *ProtoTaskDefinition) GetLogRateLimit() *ProtoLogRateLimit {
	if x != nil {
		return x.LogRateLimit
	}
	return nil
}

func (x *ProtoTaskDefinition) GetMetricTags() map[string]*ProtoMetricTagValue {
	if x != nil {
		return x.MetricTags
	}
	return nil
}

type ProtoTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskDefinition   *ProtoTaskDefinition `protobuf:"bytes,1,opt,name=task_definition,proto3" json:"task_definition,omitempty"`
	TaskGuid         string               `protobuf:"bytes,2,opt,name=task_guid,proto3" json:"task_guid,omitempty"`
	Domain           string               `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	CreatedAt        int64                `protobuf:"varint,4,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt        int64                `protobuf:"varint,5,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	FirstCompletedAt int64                `protobuf:"varint,6,opt,name=first_completed_at,proto3" json:"first_completed_at,omitempty"`
	State            ProtoTask_State      `protobuf:"varint,7,opt,name=state,proto3,enum=models.ProtoTask_State" json:"state,omitempty"`
	CellId           string               `protobuf:"bytes,8,opt,name=cell_id,proto3" json:"cell_id,omitempty"`
	Result           string               `protobuf:"bytes,9,opt,name=result,proto3" json:"result,omitempty"`
	Failed           bool                 `protobuf:"varint,10,opt,name=failed,proto3" json:"failed,omitempty"`
	FailureReason    string               `protobuf:"bytes,11,opt,name=failure_reason,proto3" json:"failure_reason,omitempty"`
	RejectionCount   int32                `protobuf:"varint,12,opt,name=rejection_count,proto3" json:"rejection_count,omitempty"`
	RejectionReason  string               `protobuf:"bytes,13,opt,name=rejection_reason,proto3" json:"rejection_reason,omitempty"`
}

func (x *ProtoTask) Reset() {
	*x = ProtoTask{}
	mi := &file_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoTask) ProtoMessage() {}

func (x *ProtoTask) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoTask.ProtoReflect.Descriptor instead.
func (*ProtoTask) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoTask) GetTaskDefinition() *ProtoTaskDefinition {
	if x != nil {
		return x.TaskDefinition
	}
	return nil
}

func (x *ProtoTask) GetTaskGuid() string {
	if x != nil {
		return x.TaskGuid
	}
	return ""
}

func (x *ProtoTask) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ProtoTask) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ProtoTask) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ProtoTask) GetFirstCompletedAt() int64 {
	if x != nil {
		return x.FirstCompletedAt
	}
	return 0
}

func (x *ProtoTask) GetState() ProtoTask_State {
	if x != nil {
		return x.State
	}
	return ProtoTask_Invalid
}

func (x *ProtoTask) GetCellId() string {
	if x != nil {
		return x.CellId
	}
	return ""
}

func (x *ProtoTask) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *ProtoTask) GetFailed() bool {
	if x != nil {
		return x.Failed
	}
	return false
}

func (x *ProtoTask) GetFailureReason() string {
	if x != nil {
		return x.FailureReason
	}
	return ""
}

func (x *ProtoTask) GetRejectionCount() int32 {
	if x != nil {
		return x.RejectionCount
	}
	return 0
}

func (x *ProtoTask) GetRejectionReason() string {
	if x != nil {
		return x.RejectionReason
	}
	return ""
}

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x1a, 0x09, 0x62, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74, 0x61, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x0b, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x66, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x74, 0x66, 0x73, 0x12, 0x44, 0x0a,
	0x15, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x03,
	0x65, 0x6e, 0x76, 0x12, 0x2b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x07, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x6d, 0x62, 0x12,
	0x21, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x62, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f,
	0x6d, 0x62, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x70, 0x75, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x0a, 0x63, 0x70, 0x75,
	0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xc0, 0x3e, 0x01,
	0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0a,
	0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x67, 0x75,
	0x69, 0x64, 0x12, 0x27, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x67, 0x75,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x0c, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0c,
	0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x75,
	0x6c, 0x65, 0x52, 0x0c, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x12, 0x4f, 0x0a, 0x13, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x13, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x36, 0x0a, 0x14, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x14, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x20, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x20, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x0d, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0d, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x69, 0x64, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x03, 0xc0, 0x3e, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x69, 0x64, 0x73, 0x12, 0x5a,
	0x0a, 0x16, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x16, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xc0, 0x3e, 0x01, 0x52, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x3b, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x41, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x6f, 0x67, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x52, 0x0e, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x4d, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x61, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74,
	0x61, 0x67, 0x73, 0x1a, 0x5a, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x61, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x61, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xfe, 0x04, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x4a, 0x0a,
	0x0f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x09, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x3e,
	0x01, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x3e,
	0x01, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xc0,
	0x3e, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x23,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x12, 0x33, 0x0a, 0x12, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x03, 0xc0, 0x3e, 0x01, 0x52, 0x12, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x07,
	0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0,
	0x3e, 0x01, 0x52, 0x07, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x3e, 0x01,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x06, 0x66,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0,
	0x3e, 0x01, 0x52, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x0f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xc0, 0x3e, 0x01,
	0x52, 0x0f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2f, 0x0a, 0x10, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x3e, 0x01,
	0x52, 0x10, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10,
	0x03, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x69, 0x6e, 0x67, 0x10, 0x04,
	0x42, 0x22, 0x5a, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x72, 0x79, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x62, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData = file_task_proto_rawDesc
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_proto_rawDescData)
	})
	return file_task_proto_rawDescData
}

var file_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_task_proto_goTypes = []any{
	(ProtoTask_State)(0),               // 0: models.ProtoTask.State
	(*ProtoTaskDefinition)(nil),        // 1: models.ProtoTaskDefinition
	(*ProtoTask)(nil),                  // 2: models.ProtoTask
	nil,                                // 3: models.ProtoTaskDefinition.MetricTagsEntry
	(*ProtoEnvironmentVariable)(nil),   // 4: models.ProtoEnvironmentVariable
	(*ProtoAction)(nil),                // 5: models.ProtoAction
	(*ProtoSecurityGroupRule)(nil),     // 6: models.ProtoSecurityGroupRule
	(*ProtoCachedDependency)(nil),      // 7: models.ProtoCachedDependency
	(*ProtoVolumeMount)(nil),           // 8: models.ProtoVolumeMount
	(*ProtoNetwork)(nil),               // 9: models.ProtoNetwork
	(*ProtoCertificateProperties)(nil), // 10: models.ProtoCertificateProperties
	(*ProtoImageLayer)(nil),            // 11: models.ProtoImageLayer
	(*ProtoLogRateLimit)(nil),          // 12: models.ProtoLogRateLimit
	(*ProtoMetricTagValue)(nil),        // 13: models.ProtoMetricTagValue
}
var file_task_proto_depIdxs = []int32{
	4,  // 0: models.ProtoTaskDefinition.environment_variables:type_name -> models.ProtoEnvironmentVariable
	5,  // 1: models.ProtoTaskDefinition.action:type_name -> models.ProtoAction
	6,  // 2: models.ProtoTaskDefinition.egress_rules:type_name -> models.ProtoSecurityGroupRule
	7,  // 3: models.ProtoTaskDefinition.cached_dependencies:type_name -> models.ProtoCachedDependency
	8,  // 4: models.ProtoTaskDefinition.volume_mounts:type_name -> models.ProtoVolumeMount
	9,  // 5: models.ProtoTaskDefinition.network:type_name -> models.ProtoNetwork
	10, // 6: models.ProtoTaskDefinition.certificate_properties:type_name -> models.ProtoCertificateProperties
	11, // 7: models.ProtoTaskDefinition.image_layers:type_name -> models.ProtoImageLayer
	12, // 8: models.ProtoTaskDefinition.log_rate_limit:type_name -> models.ProtoLogRateLimit
	3,  // 9: models.ProtoTaskDefinition.metric_tags:type_name -> models.ProtoTaskDefinition.MetricTagsEntry
	1,  // 10: models.ProtoTask.task_definition:type_name -> models.ProtoTaskDefinition
	0,  // 11: models.ProtoTask.state:type_name -> models.ProtoTask.State
	13, // 12: models.ProtoTaskDefinition.MetricTagsEntry.value:type_name -> models.ProtoMetricTagValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	file_bbs_proto_init()
	file_actions_proto_init()
	file_environment_variables_proto_init()
	file_security_group_proto_init()
	file_cached_dependency_proto_init()
	file_volume_mount_proto_init()
	file_network_proto_init()
	file_certificate_properties_proto_init()
	file_image_layer_proto_init()
	file_log_rate_limit_proto_init()
	file_metric_tags_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		EnumInfos:         file_task_proto_enumTypes,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}

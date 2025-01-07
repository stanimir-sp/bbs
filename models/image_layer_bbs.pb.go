// Code generated by protoc-gen-go-bbs. DO NOT EDIT.
// versions:
// - protoc-gen-go-bbs v0.0.1
// - protoc            v5.29.2
// source: image_layer.proto

package models

import (
	strconv "strconv"
)

type ImageLayer_DigestAlgorithm int32

const (
	ImageLayer_DigestAlgorithmInvalid ImageLayer_DigestAlgorithm = 0
	ImageLayer_DigestAlgorithmSha256  ImageLayer_DigestAlgorithm = 1
	// Deprecated: marked deprecated in proto file
	ImageLayer_DigestAlgorithmSha512 ImageLayer_DigestAlgorithm = 2
)

// Enum value maps for ImageLayer_DigestAlgorithm
var (
	ImageLayer_DigestAlgorithm_name = map[int32]string{
		0: "DigestAlgorithmInvalid",
		1: "SHA256",
		2: "SHA512",
	}
	ImageLayer_DigestAlgorithm_value = map[string]int32{
		"DigestAlgorithmInvalid": 0,
		"SHA256":                 1,
		"SHA512":                 2,
	}
)

func (m ImageLayer_DigestAlgorithm) String() string {
	s, ok := ImageLayer_DigestAlgorithm_name[int32(m)]
	if ok {
		return s
	}
	return strconv.Itoa(int(m))
}

type ImageLayer_MediaType int32

const (
	ImageLayer_MediaTypeInvalid ImageLayer_MediaType = 0
	ImageLayer_MediaTypeTgz     ImageLayer_MediaType = 1
	ImageLayer_MediaTypeTar     ImageLayer_MediaType = 2
	ImageLayer_MediaTypeZip     ImageLayer_MediaType = 3
)

// Enum value maps for ImageLayer_MediaType
var (
	ImageLayer_MediaType_name = map[int32]string{
		0: "MediaTypeInvalid",
		1: "TGZ",
		2: "TAR",
		3: "ZIP",
	}
	ImageLayer_MediaType_value = map[string]int32{
		"MediaTypeInvalid": 0,
		"TGZ":              1,
		"TAR":              2,
		"ZIP":              3,
	}
)

func (m ImageLayer_MediaType) String() string {
	s, ok := ImageLayer_MediaType_name[int32(m)]
	if ok {
		return s
	}
	return strconv.Itoa(int(m))
}

type ImageLayer_Type int32

const (
	ImageLayer_LayerTypeInvalid   ImageLayer_Type = 0
	ImageLayer_LayerTypeShared    ImageLayer_Type = 1
	ImageLayer_LayerTypeExclusive ImageLayer_Type = 2
)

// Enum value maps for ImageLayer_Type
var (
	ImageLayer_Type_name = map[int32]string{
		0: "LayerTypeInvalid",
		1: "SHARED",
		2: "EXCLUSIVE",
	}
	ImageLayer_Type_value = map[string]int32{
		"LayerTypeInvalid": 0,
		"SHARED":           1,
		"EXCLUSIVE":        2,
	}
)

func (m ImageLayer_Type) String() string {
	s, ok := ImageLayer_Type_name[int32(m)]
	if ok {
		return s
	}
	return strconv.Itoa(int(m))
}

// Prevent copylock errors when using ProtoImageLayer directly
type ImageLayer struct {
	Name            string                     `json:"name,omitempty"`
	Url             string                     `json:"url"`
	DestinationPath string                     `json:"destination_path"`
	LayerType       ImageLayer_Type            `json:"layer_type"`
	MediaType       ImageLayer_MediaType       `json:"media_type"`
	DigestAlgorithm ImageLayer_DigestAlgorithm `json:"digest_algorithm,omitempty"`
	DigestValue     string                     `json:"digest_value,omitempty"`
}

func (this *ImageLayer) Equal(that interface{}) bool {

	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ImageLayer)
	if !ok {
		that2, ok := that.(ImageLayer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}

	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}

	if this.Name != that1.Name {
		return false
	}
	if this.Url != that1.Url {
		return false
	}
	if this.DestinationPath != that1.DestinationPath {
		return false
	}
	if this.LayerType != that1.LayerType {
		return false
	}
	if this.MediaType != that1.MediaType {
		return false
	}
	if this.DigestAlgorithm != that1.DigestAlgorithm {
		return false
	}
	if this.DigestValue != that1.DigestValue {
		return false
	}
	return true
}
func (m *ImageLayer) GetName() string {
	if m != nil {
		return m.Name
	}
	var defaultValue string
	defaultValue = ""
	return defaultValue
}
func (m *ImageLayer) SetName(value string) {
	if m != nil {
		m.Name = value
	}
}
func (m *ImageLayer) GetUrl() string {
	if m != nil {
		return m.Url
	}
	var defaultValue string
	defaultValue = ""
	return defaultValue
}
func (m *ImageLayer) SetUrl(value string) {
	if m != nil {
		m.Url = value
	}
}
func (m *ImageLayer) GetDestinationPath() string {
	if m != nil {
		return m.DestinationPath
	}
	var defaultValue string
	defaultValue = ""
	return defaultValue
}
func (m *ImageLayer) SetDestinationPath(value string) {
	if m != nil {
		m.DestinationPath = value
	}
}
func (m *ImageLayer) GetLayerType() ImageLayer_Type {
	if m != nil {
		return m.LayerType
	}
	var defaultValue ImageLayer_Type
	defaultValue = 0
	return defaultValue
}
func (m *ImageLayer) SetLayerType(value ImageLayer_Type) {
	if m != nil {
		m.LayerType = value
	}
}
func (m *ImageLayer) GetMediaType() ImageLayer_MediaType {
	if m != nil {
		return m.MediaType
	}
	var defaultValue ImageLayer_MediaType
	defaultValue = 0
	return defaultValue
}
func (m *ImageLayer) SetMediaType(value ImageLayer_MediaType) {
	if m != nil {
		m.MediaType = value
	}
}
func (m *ImageLayer) GetDigestAlgorithm() ImageLayer_DigestAlgorithm {
	if m != nil {
		return m.DigestAlgorithm
	}
	var defaultValue ImageLayer_DigestAlgorithm
	defaultValue = 0
	return defaultValue
}
func (m *ImageLayer) SetDigestAlgorithm(value ImageLayer_DigestAlgorithm) {
	if m != nil {
		m.DigestAlgorithm = value
	}
}
func (m *ImageLayer) GetDigestValue() string {
	if m != nil {
		return m.DigestValue
	}
	var defaultValue string
	defaultValue = ""
	return defaultValue
}
func (m *ImageLayer) SetDigestValue(value string) {
	if m != nil {
		m.DigestValue = value
	}
}
func (x *ImageLayer) ToProto() *ProtoImageLayer {
	if x == nil {
		return nil
	}

	proto := &ProtoImageLayer{
		Name:            x.Name,
		Url:             x.Url,
		DestinationPath: x.DestinationPath,
		LayerType:       ProtoImageLayer_Type(x.LayerType),
		MediaType:       ProtoImageLayer_MediaType(x.MediaType),
		DigestAlgorithm: ProtoImageLayer_DigestAlgorithm(x.DigestAlgorithm),
		DigestValue:     x.DigestValue,
	}
	return proto
}

func (x *ProtoImageLayer) FromProto() *ImageLayer {
	if x == nil {
		return nil
	}

	copysafe := &ImageLayer{
		Name:            x.Name,
		Url:             x.Url,
		DestinationPath: x.DestinationPath,
		LayerType:       ImageLayer_Type(x.LayerType),
		MediaType:       ImageLayer_MediaType(x.MediaType),
		DigestAlgorithm: ImageLayer_DigestAlgorithm(x.DigestAlgorithm),
		DigestValue:     x.DigestValue,
	}
	return copysafe
}

func ImageLayerToProtoSlice(values []*ImageLayer) []*ProtoImageLayer {
	if values == nil {
		return nil
	}
	result := make([]*ProtoImageLayer, len(values))
	for i, val := range values {
		result[i] = val.ToProto()
	}
	return result
}

func ImageLayerFromProtoSlice(values []*ProtoImageLayer) []*ImageLayer {
	if values == nil {
		return nil
	}
	result := make([]*ImageLayer, len(values))
	for i, val := range values {
		result[i] = val.FromProto()
	}
	return result
}

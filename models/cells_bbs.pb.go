// Code generated by protoc-gen-go-bbs. DO NOT EDIT.
// versions:
// - protoc-gen-go-bbs v0.0.1
// - protoc            v5.29.3
// source: cells.proto

package models

// Prevent copylock errors when using ProtoCellCapacity directly
type CellCapacity struct {
	MemoryMb   int32 `json:"memory_mb"`
	DiskMb     int32 `json:"disk_mb"`
	Containers int32 `json:"containers"`
}

func (this *CellCapacity) Equal(that interface{}) bool {

	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CellCapacity)
	if !ok {
		that2, ok := that.(CellCapacity)
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

	if this.MemoryMb != that1.MemoryMb {
		return false
	}
	if this.DiskMb != that1.DiskMb {
		return false
	}
	if this.Containers != that1.Containers {
		return false
	}
	return true
}
func (m *CellCapacity) GetMemoryMb() int32 {
	if m != nil {
		return m.MemoryMb
	}
	var defaultValue int32
	defaultValue = 0
	return defaultValue
}
func (m *CellCapacity) SetMemoryMb(value int32) {
	if m != nil {
		m.MemoryMb = value
	}
}
func (m *CellCapacity) GetDiskMb() int32 {
	if m != nil {
		return m.DiskMb
	}
	var defaultValue int32
	defaultValue = 0
	return defaultValue
}
func (m *CellCapacity) SetDiskMb(value int32) {
	if m != nil {
		m.DiskMb = value
	}
}
func (m *CellCapacity) GetContainers() int32 {
	if m != nil {
		return m.Containers
	}
	var defaultValue int32
	defaultValue = 0
	return defaultValue
}
func (m *CellCapacity) SetContainers(value int32) {
	if m != nil {
		m.Containers = value
	}
}
func (x *CellCapacity) ToProto() *ProtoCellCapacity {
	if x == nil {
		return nil
	}

	proto := &ProtoCellCapacity{
		MemoryMb:   x.MemoryMb,
		DiskMb:     x.DiskMb,
		Containers: x.Containers,
	}
	return proto
}

func (x *ProtoCellCapacity) FromProto() *CellCapacity {
	if x == nil {
		return nil
	}

	copysafe := &CellCapacity{
		MemoryMb:   x.MemoryMb,
		DiskMb:     x.DiskMb,
		Containers: x.Containers,
	}
	return copysafe
}

func CellCapacityToProtoSlice(values []*CellCapacity) []*ProtoCellCapacity {
	if values == nil {
		return nil
	}
	result := make([]*ProtoCellCapacity, len(values))
	for i, val := range values {
		result[i] = val.ToProto()
	}
	return result
}

func CellCapacityFromProtoSlice(values []*ProtoCellCapacity) []*CellCapacity {
	if values == nil {
		return nil
	}
	result := make([]*CellCapacity, len(values))
	for i, val := range values {
		result[i] = val.FromProto()
	}
	return result
}

// Prevent copylock errors when using ProtoCellPresence directly
type CellPresence struct {
	CellId                string        `json:"cell_id"`
	RepAddress            string        `json:"rep_address"`
	Zone                  string        `json:"zone"`
	Capacity              *CellCapacity `json:"capacity,omitempty"`
	RootfsProviders       []*Provider   `json:"rootfs_provider_list,omitempty"`
	PlacementTags         []string      `json:"placement_tags,omitempty"`
	OptionalPlacementTags []string      `json:"optional_placement_tags,omitempty"`
	RepUrl                string        `json:"rep_url"`
}

func (this *CellPresence) Equal(that interface{}) bool {

	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CellPresence)
	if !ok {
		that2, ok := that.(CellPresence)
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

	if this.CellId != that1.CellId {
		return false
	}
	if this.RepAddress != that1.RepAddress {
		return false
	}
	if this.Zone != that1.Zone {
		return false
	}
	if this.Capacity == nil {
		if that1.Capacity != nil {
			return false
		}
	} else if !this.Capacity.Equal(*that1.Capacity) {
		return false
	}
	if this.RootfsProviders == nil {
		if that1.RootfsProviders != nil {
			return false
		}
	} else if len(this.RootfsProviders) != len(that1.RootfsProviders) {
		return false
	}
	for i := range this.RootfsProviders {
		if !this.RootfsProviders[i].Equal(that1.RootfsProviders[i]) {
			return false
		}
	}
	if this.PlacementTags == nil {
		if that1.PlacementTags != nil {
			return false
		}
	} else if len(this.PlacementTags) != len(that1.PlacementTags) {
		return false
	}
	for i := range this.PlacementTags {
		if this.PlacementTags[i] != that1.PlacementTags[i] {
			return false
		}
	}
	if this.OptionalPlacementTags == nil {
		if that1.OptionalPlacementTags != nil {
			return false
		}
	} else if len(this.OptionalPlacementTags) != len(that1.OptionalPlacementTags) {
		return false
	}
	for i := range this.OptionalPlacementTags {
		if this.OptionalPlacementTags[i] != that1.OptionalPlacementTags[i] {
			return false
		}
	}
	if this.RepUrl != that1.RepUrl {
		return false
	}
	return true
}
func (m *CellPresence) GetCellId() string {
	if m != nil {
		return m.CellId
	}
	var defaultValue string
	defaultValue = ""
	return defaultValue
}
func (m *CellPresence) SetCellId(value string) {
	if m != nil {
		m.CellId = value
	}
}
func (m *CellPresence) GetRepAddress() string {
	if m != nil {
		return m.RepAddress
	}
	var defaultValue string
	defaultValue = ""
	return defaultValue
}
func (m *CellPresence) SetRepAddress(value string) {
	if m != nil {
		m.RepAddress = value
	}
}
func (m *CellPresence) GetZone() string {
	if m != nil {
		return m.Zone
	}
	var defaultValue string
	defaultValue = ""
	return defaultValue
}
func (m *CellPresence) SetZone(value string) {
	if m != nil {
		m.Zone = value
	}
}
func (m *CellPresence) GetCapacity() *CellCapacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}
func (m *CellPresence) SetCapacity(value *CellCapacity) {
	if m != nil {
		m.Capacity = value
	}
}
func (m *CellPresence) GetRootfsProviders() []*Provider {
	if m != nil {
		return m.RootfsProviders
	}
	return nil
}
func (m *CellPresence) SetRootfsProviders(value []*Provider) {
	if m != nil {
		m.RootfsProviders = value
	}
}
func (m *CellPresence) GetPlacementTags() []string {
	if m != nil {
		return m.PlacementTags
	}
	return nil
}
func (m *CellPresence) SetPlacementTags(value []string) {
	if m != nil {
		m.PlacementTags = value
	}
}
func (m *CellPresence) GetOptionalPlacementTags() []string {
	if m != nil {
		return m.OptionalPlacementTags
	}
	return nil
}
func (m *CellPresence) SetOptionalPlacementTags(value []string) {
	if m != nil {
		m.OptionalPlacementTags = value
	}
}
func (m *CellPresence) GetRepUrl() string {
	if m != nil {
		return m.RepUrl
	}
	var defaultValue string
	defaultValue = ""
	return defaultValue
}
func (m *CellPresence) SetRepUrl(value string) {
	if m != nil {
		m.RepUrl = value
	}
}
func (x *CellPresence) ToProto() *ProtoCellPresence {
	if x == nil {
		return nil
	}

	proto := &ProtoCellPresence{
		CellId:                x.CellId,
		RepAddress:            x.RepAddress,
		Zone:                  x.Zone,
		Capacity:              x.Capacity.ToProto(),
		RootfsProviders:       ProviderToProtoSlice(x.RootfsProviders),
		PlacementTags:         x.PlacementTags,
		OptionalPlacementTags: x.OptionalPlacementTags,
		RepUrl:                x.RepUrl,
	}
	return proto
}

func (x *ProtoCellPresence) FromProto() *CellPresence {
	if x == nil {
		return nil
	}

	copysafe := &CellPresence{
		CellId:                x.CellId,
		RepAddress:            x.RepAddress,
		Zone:                  x.Zone,
		Capacity:              x.Capacity.FromProto(),
		RootfsProviders:       ProviderFromProtoSlice(x.RootfsProviders),
		PlacementTags:         x.PlacementTags,
		OptionalPlacementTags: x.OptionalPlacementTags,
		RepUrl:                x.RepUrl,
	}
	return copysafe
}

func CellPresenceToProtoSlice(values []*CellPresence) []*ProtoCellPresence {
	if values == nil {
		return nil
	}
	result := make([]*ProtoCellPresence, len(values))
	for i, val := range values {
		result[i] = val.ToProto()
	}
	return result
}

func CellPresenceFromProtoSlice(values []*ProtoCellPresence) []*CellPresence {
	if values == nil {
		return nil
	}
	result := make([]*CellPresence, len(values))
	for i, val := range values {
		result[i] = val.FromProto()
	}
	return result
}

// Prevent copylock errors when using ProtoProvider directly
type Provider struct {
	Name       string   `json:"name"`
	Properties []string `json:"properties,omitempty"`
}

func (this *Provider) Equal(that interface{}) bool {

	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Provider)
	if !ok {
		that2, ok := that.(Provider)
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
	if this.Properties == nil {
		if that1.Properties != nil {
			return false
		}
	} else if len(this.Properties) != len(that1.Properties) {
		return false
	}
	for i := range this.Properties {
		if this.Properties[i] != that1.Properties[i] {
			return false
		}
	}
	return true
}
func (m *Provider) GetName() string {
	if m != nil {
		return m.Name
	}
	var defaultValue string
	defaultValue = ""
	return defaultValue
}
func (m *Provider) SetName(value string) {
	if m != nil {
		m.Name = value
	}
}
func (m *Provider) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}
func (m *Provider) SetProperties(value []string) {
	if m != nil {
		m.Properties = value
	}
}
func (x *Provider) ToProto() *ProtoProvider {
	if x == nil {
		return nil
	}

	proto := &ProtoProvider{
		Name:       x.Name,
		Properties: x.Properties,
	}
	return proto
}

func (x *ProtoProvider) FromProto() *Provider {
	if x == nil {
		return nil
	}

	copysafe := &Provider{
		Name:       x.Name,
		Properties: x.Properties,
	}
	return copysafe
}

func ProviderToProtoSlice(values []*Provider) []*ProtoProvider {
	if values == nil {
		return nil
	}
	result := make([]*ProtoProvider, len(values))
	for i, val := range values {
		result[i] = val.ToProto()
	}
	return result
}

func ProviderFromProtoSlice(values []*ProtoProvider) []*Provider {
	if values == nil {
		return nil
	}
	result := make([]*Provider, len(values))
	for i, val := range values {
		result[i] = val.FromProto()
	}
	return result
}

// Prevent copylock errors when using ProtoCellsResponse directly
type CellsResponse struct {
	Error *Error          `json:"error,omitempty"`
	Cells []*CellPresence `json:"cells,omitempty"`
}

func (this *CellsResponse) Equal(that interface{}) bool {

	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CellsResponse)
	if !ok {
		that2, ok := that.(CellsResponse)
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

	if this.Error == nil {
		if that1.Error != nil {
			return false
		}
	} else if !this.Error.Equal(*that1.Error) {
		return false
	}
	if this.Cells == nil {
		if that1.Cells != nil {
			return false
		}
	} else if len(this.Cells) != len(that1.Cells) {
		return false
	}
	for i := range this.Cells {
		if !this.Cells[i].Equal(that1.Cells[i]) {
			return false
		}
	}
	return true
}
func (m *CellsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}
func (m *CellsResponse) SetError(value *Error) {
	if m != nil {
		m.Error = value
	}
}
func (m *CellsResponse) GetCells() []*CellPresence {
	if m != nil {
		return m.Cells
	}
	return nil
}
func (m *CellsResponse) SetCells(value []*CellPresence) {
	if m != nil {
		m.Cells = value
	}
}
func (x *CellsResponse) ToProto() *ProtoCellsResponse {
	if x == nil {
		return nil
	}

	proto := &ProtoCellsResponse{
		Error: x.Error.ToProto(),
		Cells: CellPresenceToProtoSlice(x.Cells),
	}
	return proto
}

func (x *ProtoCellsResponse) FromProto() *CellsResponse {
	if x == nil {
		return nil
	}

	copysafe := &CellsResponse{
		Error: x.Error.FromProto(),
		Cells: CellPresenceFromProtoSlice(x.Cells),
	}
	return copysafe
}

func CellsResponseToProtoSlice(values []*CellsResponse) []*ProtoCellsResponse {
	if values == nil {
		return nil
	}
	result := make([]*ProtoCellsResponse, len(values))
	for i, val := range values {
		result[i] = val.ToProto()
	}
	return result
}

func CellsResponseFromProtoSlice(values []*ProtoCellsResponse) []*CellsResponse {
	if values == nil {
		return nil
	}
	result := make([]*CellsResponse, len(values))
	for i, val := range values {
		result[i] = val.FromProto()
	}
	return result
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/failover_service.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//A FailoverService creates a new hostname to which services can send requests.
//Requests will be routed based on a list of backing traffic targets ordered by
//decreasing priority. When outlier detection detects that a traffic target in the list is
//in an unhealthy state, requests sent to the FailoverService will be routed
//to the next healthy traffic target in the list. For each traffic target referenced in the
//FailoverService's BackingServices list, outlier detection must be configured using a TrafficPolicy.
//
//Currently this feature only supports Services backed by Istio.
type FailoverServiceSpec struct {
	//
	//The DNS name of the FailoverService. Must be unique within the service mesh instance
	//since it is used as the hostname with which clients communicate.
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// The port on which the FailoverService listens.
	Port *FailoverServiceSpec_Port `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	// The meshes that this FailoverService will be visible to.
	Meshes []*v1.ObjectRef `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty"`
	//
	//The list of services backing the FailoverService, ordered by decreasing priority.
	//All services must be backed by either the same service mesh instance or
	//backed by service meshes that are grouped under a common VirtualMesh.
	BackingServices      []*FailoverServiceSpec_BackingService `protobuf:"bytes,4,rep,name=backing_services,json=backingServices,proto3" json:"backing_services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *FailoverServiceSpec) Reset()         { *m = FailoverServiceSpec{} }
func (m *FailoverServiceSpec) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceSpec) ProtoMessage()    {}
func (*FailoverServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{0}
}
func (m *FailoverServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceSpec.Unmarshal(m, b)
}
func (m *FailoverServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceSpec.Marshal(b, m, deterministic)
}
func (m *FailoverServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceSpec.Merge(m, src)
}
func (m *FailoverServiceSpec) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceSpec.Size(m)
}
func (m *FailoverServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceSpec proto.InternalMessageInfo

func (m *FailoverServiceSpec) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *FailoverServiceSpec) GetPort() *FailoverServiceSpec_Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *FailoverServiceSpec) GetMeshes() []*v1.ObjectRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *FailoverServiceSpec) GetBackingServices() []*FailoverServiceSpec_BackingService {
	if m != nil {
		return m.BackingServices
	}
	return nil
}

// The port on which the FailoverService listens.
type FailoverServiceSpec_Port struct {
	// Port number.
	Number uint32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// Protocol of the requests sent to the FailoverService, must be one of HTTP, HTTPS, GRPC, HTTP2, MONGO, TCP, TLS.
	Protocol             string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailoverServiceSpec_Port) Reset()         { *m = FailoverServiceSpec_Port{} }
func (m *FailoverServiceSpec_Port) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceSpec_Port) ProtoMessage()    {}
func (*FailoverServiceSpec_Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{0, 0}
}
func (m *FailoverServiceSpec_Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceSpec_Port.Unmarshal(m, b)
}
func (m *FailoverServiceSpec_Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceSpec_Port.Marshal(b, m, deterministic)
}
func (m *FailoverServiceSpec_Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceSpec_Port.Merge(m, src)
}
func (m *FailoverServiceSpec_Port) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceSpec_Port.Size(m)
}
func (m *FailoverServiceSpec_Port) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceSpec_Port.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceSpec_Port proto.InternalMessageInfo

func (m *FailoverServiceSpec_Port) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *FailoverServiceSpec_Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

// The traffic targets that comprise the FailoverService.
type FailoverServiceSpec_BackingService struct {
	// Different traffic target types can be selected as backing services.
	//
	// Types that are valid to be assigned to BackingServiceType:
	//	*FailoverServiceSpec_BackingService_KubeService
	BackingServiceType   isFailoverServiceSpec_BackingService_BackingServiceType `protobuf_oneof:"backing_service_type"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *FailoverServiceSpec_BackingService) Reset()         { *m = FailoverServiceSpec_BackingService{} }
func (m *FailoverServiceSpec_BackingService) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceSpec_BackingService) ProtoMessage()    {}
func (*FailoverServiceSpec_BackingService) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{0, 1}
}
func (m *FailoverServiceSpec_BackingService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceSpec_BackingService.Unmarshal(m, b)
}
func (m *FailoverServiceSpec_BackingService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceSpec_BackingService.Marshal(b, m, deterministic)
}
func (m *FailoverServiceSpec_BackingService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceSpec_BackingService.Merge(m, src)
}
func (m *FailoverServiceSpec_BackingService) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceSpec_BackingService.Size(m)
}
func (m *FailoverServiceSpec_BackingService) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceSpec_BackingService.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceSpec_BackingService proto.InternalMessageInfo

type isFailoverServiceSpec_BackingService_BackingServiceType interface {
	isFailoverServiceSpec_BackingService_BackingServiceType()
	Equal(interface{}) bool
}

type FailoverServiceSpec_BackingService_KubeService struct {
	KubeService *v1.ClusterObjectRef `protobuf:"bytes,1,opt,name=kube_service,json=kubeService,proto3,oneof" json:"kube_service,omitempty"`
}

func (*FailoverServiceSpec_BackingService_KubeService) isFailoverServiceSpec_BackingService_BackingServiceType() {
}

func (m *FailoverServiceSpec_BackingService) GetBackingServiceType() isFailoverServiceSpec_BackingService_BackingServiceType {
	if m != nil {
		return m.BackingServiceType
	}
	return nil
}

func (m *FailoverServiceSpec_BackingService) GetKubeService() *v1.ClusterObjectRef {
	if x, ok := m.GetBackingServiceType().(*FailoverServiceSpec_BackingService_KubeService); ok {
		return x.KubeService
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FailoverServiceSpec_BackingService) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FailoverServiceSpec_BackingService_KubeService)(nil),
	}
}

type FailoverServiceStatus struct {
	//
	//The most recent generation observed in the the FailoverService metadata.
	//If the observedGeneration does not match generation, the controller has not received the most
	//recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	//
	//The state of the overall resource, will only show accepted if it has been successfully
	//applied to all target meshes.
	State ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=networking.smh.solo.io.ApprovalState" json:"state,omitempty"`
	// The status of the FailoverService for each Mesh to which it has been applied.
	Meshes map[string]*ApprovalStatus `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Any errors found while processing this generation of the resource.
	Errors               []string `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailoverServiceStatus) Reset()         { *m = FailoverServiceStatus{} }
func (m *FailoverServiceStatus) String() string { return proto.CompactTextString(m) }
func (*FailoverServiceStatus) ProtoMessage()    {}
func (*FailoverServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2a3822bd950167, []int{1}
}
func (m *FailoverServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailoverServiceStatus.Unmarshal(m, b)
}
func (m *FailoverServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailoverServiceStatus.Marshal(b, m, deterministic)
}
func (m *FailoverServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverServiceStatus.Merge(m, src)
}
func (m *FailoverServiceStatus) XXX_Size() int {
	return xxx_messageInfo_FailoverServiceStatus.Size(m)
}
func (m *FailoverServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverServiceStatus proto.InternalMessageInfo

func (m *FailoverServiceStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *FailoverServiceStatus) GetState() ApprovalState {
	if m != nil {
		return m.State
	}
	return ApprovalState_PENDING
}

func (m *FailoverServiceStatus) GetMeshes() map[string]*ApprovalStatus {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *FailoverServiceStatus) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*FailoverServiceSpec)(nil), "networking.smh.solo.io.FailoverServiceSpec")
	proto.RegisterType((*FailoverServiceSpec_Port)(nil), "networking.smh.solo.io.FailoverServiceSpec.Port")
	proto.RegisterType((*FailoverServiceSpec_BackingService)(nil), "networking.smh.solo.io.FailoverServiceSpec.BackingService")
	proto.RegisterType((*FailoverServiceStatus)(nil), "networking.smh.solo.io.FailoverServiceStatus")
	proto.RegisterMapType((map[string]*ApprovalStatus)(nil), "networking.smh.solo.io.FailoverServiceStatus.MeshesEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/networking/v1alpha2/failover_service.proto", fileDescriptor_4c2a3822bd950167)
}

var fileDescriptor_4c2a3822bd950167 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0xf2, 0xd3, 0xe8, 0xcb, 0x06, 0x4a, 0xb5, 0x2d, 0x51, 0x64, 0x10, 0x8a, 0x8a, 0x40,
	0xb9, 0xc9, 0x9a, 0x1a, 0x2e, 0x20, 0x20, 0x21, 0xc2, 0x5f, 0x25, 0x84, 0xa0, 0xdb, 0x3b, 0x6e,
	0xa2, 0xb5, 0x3b, 0xb5, 0x4d, 0x6c, 0x8f, 0xb5, 0xbb, 0x36, 0xca, 0x1b, 0xf1, 0x20, 0xbc, 0x01,
	0x6f, 0xc0, 0x93, 0xa0, 0x5d, 0x3b, 0xa6, 0x8d, 0x82, 0x14, 0xae, 0xb2, 0xe3, 0x9d, 0x73, 0x66,
	0xce, 0x99, 0xd9, 0x10, 0x1e, 0xc6, 0x3a, 0x2a, 0x7c, 0x16, 0x60, 0xea, 0x2a, 0x4c, 0x70, 0x1a,
	0xa3, 0xab, 0x40, 0x96, 0x71, 0x00, 0xd3, 0x14, 0x54, 0x34, 0x8d, 0x0a, 0xdf, 0x15, 0x79, 0xec,
	0x66, 0xa0, 0xbf, 0xa1, 0x5c, 0xc6, 0x59, 0xe8, 0x96, 0x27, 0x22, 0xc9, 0x23, 0xe1, 0xb9, 0x97,
	0x22, 0x4e, 0xb0, 0x04, 0xb9, 0xa8, 0x11, 0x2c, 0x97, 0xa8, 0x91, 0x0e, 0xff, 0xe4, 0x32, 0x95,
	0x46, 0xcc, 0xf0, 0xb2, 0x18, 0x9d, 0x3b, 0x6a, 0x59, 0x7a, 0x96, 0x2c, 0x40, 0x09, 0x6e, 0x79,
	0x62, 0x7f, 0x2b, 0x90, 0xf3, 0x72, 0xe7, 0xaa, 0xa5, 0x48, 0xe2, 0x0b, 0xa1, 0x63, 0xcc, 0x16,
	0x4a, 0x0b, 0xbd, 0x26, 0x38, 0x0a, 0x31, 0x44, 0x7b, 0x74, 0xcd, 0xa9, 0xfa, 0x7a, 0xfc, 0xa3,
	0x43, 0x0e, 0xdf, 0xd5, 0x6d, 0x9e, 0x57, 0x15, 0xce, 0x73, 0x08, 0xa8, 0x43, 0xfe, 0x8f, 0x50,
	0xe9, 0x4c, 0xa4, 0x30, 0x6a, 0x8d, 0x5b, 0x93, 0x3e, 0x6f, 0x62, 0xfa, 0x86, 0x74, 0x73, 0x94,
	0x7a, 0xd4, 0x1e, 0xb7, 0x26, 0x03, 0xef, 0x11, 0xdb, 0x2e, 0x87, 0x6d, 0xa1, 0x65, 0x9f, 0x51,
	0x6a, 0x6e, 0xd1, 0xf4, 0x09, 0xe9, 0x19, 0x29, 0xa0, 0x46, 0x9d, 0x71, 0x67, 0x32, 0xf0, 0xee,
	0x32, 0xab, 0xd6, 0x78, 0xd0, 0x50, 0x7c, 0xf2, 0xbf, 0x42, 0xa0, 0x39, 0x5c, 0xf2, 0x3a, 0x97,
	0x02, 0x39, 0xf0, 0x45, 0x60, 0x6a, 0xad, 0x4d, 0x55, 0xa3, 0xae, 0xc5, 0xcf, 0xfe, 0xa5, 0x8f,
	0x79, 0xc5, 0x51, 0x7f, 0xe2, 0xb7, 0xfc, 0x6b, 0xb1, 0x72, 0x66, 0xa4, 0x6b, 0x5a, 0xa5, 0x43,
	0xd2, 0xcb, 0x8a, 0xd4, 0x07, 0x69, 0x4d, 0xb8, 0xc9, 0xeb, 0xc8, 0xd8, 0x63, 0xfd, 0x0b, 0x30,
	0xb1, 0x36, 0xf4, 0x79, 0x13, 0x3b, 0x92, 0xec, 0x5f, 0xa7, 0xa7, 0xa7, 0xe4, 0xc6, 0xb2, 0xf0,
	0x61, 0xdd, 0xb1, 0xe5, 0x1a, 0x78, 0xf7, 0xb7, 0x08, 0x7e, 0x9d, 0x14, 0x4a, 0x83, 0x6c, 0x74,
	0x9f, 0xfe, 0xc7, 0x07, 0x06, 0x5a, 0x33, 0xcd, 0x87, 0xe4, 0x68, 0x43, 0xfe, 0x42, 0xaf, 0x72,
	0x38, 0xfe, 0xd9, 0x26, 0xb7, 0x37, 0x75, 0x6a, 0xa1, 0x0b, 0x45, 0x5d, 0x72, 0x88, 0xbe, 0xc9,
	0x85, 0x8b, 0x45, 0x08, 0x19, 0x48, 0xbb, 0x19, 0xb6, 0x85, 0x0e, 0xa7, 0xeb, 0xab, 0xf7, 0xcd,
	0x0d, 0x7d, 0x4e, 0xf6, 0xec, 0xda, 0x58, 0x5d, 0xfb, 0xde, 0x83, 0xbf, 0xd9, 0xfa, 0x2a, 0xcf,
	0x25, 0x96, 0x22, 0x31, 0x75, 0x80, 0x57, 0x18, 0x7a, 0xb6, 0x31, 0xd4, 0x67, 0xbb, 0x0e, 0xc5,
	0x36, 0xcb, 0x3e, 0x5a, 0xec, 0xdb, 0x4c, 0xcb, 0x55, 0x33, 0xf1, 0x21, 0xe9, 0x81, 0x94, 0x28,
	0xab, 0x39, 0xf7, 0x79, 0x1d, 0x39, 0x82, 0x0c, 0xae, 0xa4, 0xd3, 0x03, 0xd2, 0x59, 0xc2, 0xaa,
	0xde, 0x55, 0x73, 0xa4, 0x2f, 0xc8, 0x5e, 0x29, 0x92, 0x02, 0xea, 0x3d, 0x7d, 0xb8, 0x8b, 0x90,
	0x42, 0xf1, 0x0a, 0x34, 0x6b, 0x3f, 0x6d, 0xcd, 0xcf, 0xbe, 0xff, 0xba, 0xd7, 0xfa, 0xf2, 0x61,
	0x97, 0xbf, 0x80, 0x7c, 0x19, 0x6e, 0x3c, 0xc8, 0xab, 0x35, 0x9a, 0xc7, 0xe9, 0xf7, 0xec, 0x9a,
	0x3c, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x28, 0x15, 0x0b, 0x2d, 0x58, 0x04, 0x00, 0x00,
}

func (this *FailoverServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceSpec)
	if !ok {
		that2, ok := that.(FailoverServiceSpec)
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
	if this.Hostname != that1.Hostname {
		return false
	}
	if !this.Port.Equal(that1.Port) {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if len(this.BackingServices) != len(that1.BackingServices) {
		return false
	}
	for i := range this.BackingServices {
		if !this.BackingServices[i].Equal(that1.BackingServices[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverServiceSpec_Port) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceSpec_Port)
	if !ok {
		that2, ok := that.(FailoverServiceSpec_Port)
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
	if this.Number != that1.Number {
		return false
	}
	if this.Protocol != that1.Protocol {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverServiceSpec_BackingService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceSpec_BackingService)
	if !ok {
		that2, ok := that.(FailoverServiceSpec_BackingService)
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
	if that1.BackingServiceType == nil {
		if this.BackingServiceType != nil {
			return false
		}
	} else if this.BackingServiceType == nil {
		return false
	} else if !this.BackingServiceType.Equal(that1.BackingServiceType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FailoverServiceSpec_BackingService_KubeService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceSpec_BackingService_KubeService)
	if !ok {
		that2, ok := that.(FailoverServiceSpec_BackingService_KubeService)
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
	if !this.KubeService.Equal(that1.KubeService) {
		return false
	}
	return true
}
func (this *FailoverServiceStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverServiceStatus)
	if !ok {
		that2, ok := that.(FailoverServiceStatus)
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
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if len(this.Errors) != len(that1.Errors) {
		return false
	}
	for i := range this.Errors {
		if this.Errors[i] != that1.Errors[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

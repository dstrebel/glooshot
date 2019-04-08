// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/glooshot/api/v1/glooshot.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	faultinjection "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/faultinjection"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// TODO: this should probably go in the status. consult with ilackarms
type ExperimentSpec_State int32

const (
	ExperimentSpec_Init    ExperimentSpec_State = 0
	ExperimentSpec_Running ExperimentSpec_State = 1
	ExperimentSpec_Done    ExperimentSpec_State = 2
)

var ExperimentSpec_State_name = map[int32]string{
	0: "Init",
	1: "Running",
	2: "Done",
}

var ExperimentSpec_State_value = map[string]int32{
	"Init":    0,
	"Running": 1,
	"Done":    2,
}

func (x ExperimentSpec_State) String() string {
	return proto.EnumName(ExperimentSpec_State_name, int32(x))
}

func (ExperimentSpec_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{1, 0}
}

type Experiment struct {
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
	Spec     *ExperimentSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Status indicates the validation status of the resource. Status is read-only by clients, and set by gloo during validation
	Status               core.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Experiment) Reset()         { *m = Experiment{} }
func (m *Experiment) String() string { return proto.CompactTextString(m) }
func (*Experiment) ProtoMessage()    {}
func (*Experiment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{0}
}
func (m *Experiment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Experiment.Unmarshal(m, b)
}
func (m *Experiment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Experiment.Marshal(b, m, deterministic)
}
func (m *Experiment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Experiment.Merge(m, src)
}
func (m *Experiment) XXX_Size() int {
	return xxx_messageInfo_Experiment.Size(m)
}
func (m *Experiment) XXX_DiscardUnknown() {
	xxx_messageInfo_Experiment.DiscardUnknown(m)
}

var xxx_messageInfo_Experiment proto.InternalMessageInfo

func (m *Experiment) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Experiment) GetSpec() *ExperimentSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Experiment) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

type ExperimentSpec struct {
	Faults        []*ExperimentSpec_InjectedFault `protobuf:"bytes,1,rep,name=faults,proto3" json:"faults,omitempty"`
	StopCondition *StopCondition                  `protobuf:"bytes,2,opt,name=stop_condition,json=stopCondition,proto3" json:"stop_condition,omitempty"`
	// State is the enum indicating the state of the resource
	State                ExperimentSpec_State `protobuf:"varint,4,opt,name=state,proto3,enum=glooshot.solo.io.ExperimentSpec_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ExperimentSpec) Reset()         { *m = ExperimentSpec{} }
func (m *ExperimentSpec) String() string { return proto.CompactTextString(m) }
func (*ExperimentSpec) ProtoMessage()    {}
func (*ExperimentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{1}
}
func (m *ExperimentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExperimentSpec.Unmarshal(m, b)
}
func (m *ExperimentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExperimentSpec.Marshal(b, m, deterministic)
}
func (m *ExperimentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExperimentSpec.Merge(m, src)
}
func (m *ExperimentSpec) XXX_Size() int {
	return xxx_messageInfo_ExperimentSpec.Size(m)
}
func (m *ExperimentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ExperimentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ExperimentSpec proto.InternalMessageInfo

func (m *ExperimentSpec) GetFaults() []*ExperimentSpec_InjectedFault {
	if m != nil {
		return m.Faults
	}
	return nil
}

func (m *ExperimentSpec) GetStopCondition() *StopCondition {
	if m != nil {
		return m.StopCondition
	}
	return nil
}

func (m *ExperimentSpec) GetState() ExperimentSpec_State {
	if m != nil {
		return m.State
	}
	return ExperimentSpec_Init
}

type ExperimentSpec_InjectedFault struct {
	// TODO(yuval-k) should this be an upstream ref?
	Service              *v1.Destination             `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Fault                *faultinjection.RouteFaults `protobuf:"bytes,2,opt,name=fault,proto3" json:"fault,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ExperimentSpec_InjectedFault) Reset()         { *m = ExperimentSpec_InjectedFault{} }
func (m *ExperimentSpec_InjectedFault) String() string { return proto.CompactTextString(m) }
func (*ExperimentSpec_InjectedFault) ProtoMessage()    {}
func (*ExperimentSpec_InjectedFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{1, 0}
}
func (m *ExperimentSpec_InjectedFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExperimentSpec_InjectedFault.Unmarshal(m, b)
}
func (m *ExperimentSpec_InjectedFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExperimentSpec_InjectedFault.Marshal(b, m, deterministic)
}
func (m *ExperimentSpec_InjectedFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExperimentSpec_InjectedFault.Merge(m, src)
}
func (m *ExperimentSpec_InjectedFault) XXX_Size() int {
	return xxx_messageInfo_ExperimentSpec_InjectedFault.Size(m)
}
func (m *ExperimentSpec_InjectedFault) XXX_DiscardUnknown() {
	xxx_messageInfo_ExperimentSpec_InjectedFault.DiscardUnknown(m)
}

var xxx_messageInfo_ExperimentSpec_InjectedFault proto.InternalMessageInfo

func (m *ExperimentSpec_InjectedFault) GetService() *v1.Destination {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ExperimentSpec_InjectedFault) GetFault() *faultinjection.RouteFaults {
	if m != nil {
		return m.Fault
	}
	return nil
}

type StopCondition struct {
	Duration             *time.Duration     `protobuf:"bytes,1,opt,name=duration,proto3,stdduration" json:"duration,omitempty"`
	Metric               []*MetricThreshold `protobuf:"bytes,2,rep,name=metric,proto3" json:"metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StopCondition) Reset()         { *m = StopCondition{} }
func (m *StopCondition) String() string { return proto.CompactTextString(m) }
func (*StopCondition) ProtoMessage()    {}
func (*StopCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{2}
}
func (m *StopCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopCondition.Unmarshal(m, b)
}
func (m *StopCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopCondition.Marshal(b, m, deterministic)
}
func (m *StopCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopCondition.Merge(m, src)
}
func (m *StopCondition) XXX_Size() int {
	return xxx_messageInfo_StopCondition.Size(m)
}
func (m *StopCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_StopCondition.DiscardUnknown(m)
}

var xxx_messageInfo_StopCondition proto.InternalMessageInfo

func (m *StopCondition) GetDuration() *time.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *StopCondition) GetMetric() []*MetricThreshold {
	if m != nil {
		return m.Metric
	}
	return nil
}

type MetricThreshold struct {
	MetricName           string   `protobuf:"bytes,1,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricThreshold) Reset()         { *m = MetricThreshold{} }
func (m *MetricThreshold) String() string { return proto.CompactTextString(m) }
func (*MetricThreshold) ProtoMessage()    {}
func (*MetricThreshold) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{3}
}
func (m *MetricThreshold) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricThreshold.Unmarshal(m, b)
}
func (m *MetricThreshold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricThreshold.Marshal(b, m, deterministic)
}
func (m *MetricThreshold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricThreshold.Merge(m, src)
}
func (m *MetricThreshold) XXX_Size() int {
	return xxx_messageInfo_MetricThreshold.Size(m)
}
func (m *MetricThreshold) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricThreshold.DiscardUnknown(m)
}

var xxx_messageInfo_MetricThreshold proto.InternalMessageInfo

func (m *MetricThreshold) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *MetricThreshold) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ChaosSettings struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaosSettings) Reset()         { *m = ChaosSettings{} }
func (m *ChaosSettings) String() string { return proto.CompactTextString(m) }
func (*ChaosSettings) ProtoMessage()    {}
func (*ChaosSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9da8418b9c75752, []int{4}
}
func (m *ChaosSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaosSettings.Unmarshal(m, b)
}
func (m *ChaosSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaosSettings.Marshal(b, m, deterministic)
}
func (m *ChaosSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaosSettings.Merge(m, src)
}
func (m *ChaosSettings) XXX_Size() int {
	return xxx_messageInfo_ChaosSettings.Size(m)
}
func (m *ChaosSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaosSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ChaosSettings proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("glooshot.solo.io.ExperimentSpec_State", ExperimentSpec_State_name, ExperimentSpec_State_value)
	proto.RegisterType((*Experiment)(nil), "glooshot.solo.io.Experiment")
	proto.RegisterType((*ExperimentSpec)(nil), "glooshot.solo.io.ExperimentSpec")
	proto.RegisterType((*ExperimentSpec_InjectedFault)(nil), "glooshot.solo.io.ExperimentSpec.InjectedFault")
	proto.RegisterType((*StopCondition)(nil), "glooshot.solo.io.StopCondition")
	proto.RegisterType((*MetricThreshold)(nil), "glooshot.solo.io.MetricThreshold")
	proto.RegisterType((*ChaosSettings)(nil), "glooshot.solo.io.ChaosSettings")
}

func init() {
	proto.RegisterFile("github.com/solo-io/glooshot/api/v1/glooshot.proto", fileDescriptor_b9da8418b9c75752)
}

var fileDescriptor_b9da8418b9c75752 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x52, 0x13, 0x4f,
	0x10, 0xc7, 0x59, 0xb2, 0x09, 0xfc, 0x26, 0x15, 0xc8, 0x6f, 0x8a, 0xb2, 0x42, 0x0e, 0x24, 0xee,
	0x01, 0x73, 0xd0, 0xdd, 0x02, 0x3c, 0x20, 0xea, 0x25, 0x20, 0xca, 0x01, 0xad, 0x9a, 0x78, 0xf2,
	0x42, 0x2d, 0x9b, 0xce, 0x66, 0x64, 0x33, 0xbd, 0xb5, 0x33, 0x4b, 0xc5, 0x2b, 0x1e, 0xbc, 0x79,
	0xf6, 0x11, 0x7c, 0x14, 0x9f, 0x02, 0x4b, 0xdf, 0x00, 0x9f, 0xc0, 0xda, 0xd9, 0xd9, 0xc5, 0x80,
	0x8a, 0x9c, 0x92, 0x9e, 0xee, 0x4f, 0xff, 0xf9, 0xce, 0xf4, 0x92, 0x8d, 0x90, 0xab, 0x71, 0x7a,
	0xec, 0x06, 0x38, 0xf1, 0x24, 0x46, 0xf8, 0x80, 0xa3, 0x17, 0x46, 0x88, 0x72, 0x8c, 0xca, 0xf3,
	0x63, 0xee, 0x9d, 0x6e, 0x94, 0xb6, 0x1b, 0x27, 0xa8, 0x90, 0x36, 0x4b, 0x3b, 0x03, 0x5c, 0x8e,
	0xed, 0x95, 0x10, 0x43, 0xd4, 0x4e, 0x2f, 0xfb, 0x97, 0xc7, 0xb5, 0xd7, 0x42, 0xc4, 0x30, 0x02,
	0x4f, 0x5b, 0xc7, 0xe9, 0xc8, 0x1b, 0xa6, 0x89, 0xaf, 0x38, 0x0a, 0xe3, 0xf7, 0x7e, 0x53, 0x5a,
	0xff, 0x9e, 0xf0, 0xb2, 0xb4, 0x54, 0xbe, 0x4a, 0xa5, 0x01, 0x36, 0xfe, 0x01, 0x98, 0x80, 0xf2,
	0x87, 0xbe, 0xf2, 0x6f, 0x81, 0x14, 0xb6, 0x41, 0xb6, 0xff, 0xa0, 0x48, 0x36, 0xc7, 0x5b, 0x08,
	0x94, 0xcc, 0x2d, 0x03, 0xc7, 0x09, 0x4e, 0xdf, 0x19, 0xf2, 0xd5, 0xed, 0xc8, 0x28, 0x0d, 0xb9,
	0x90, 0xde, 0xc8, 0x4f, 0x23, 0xc5, 0x45, 0x16, 0xc0, 0x51, 0xe4, 0x66, 0x9e, 0xd0, 0xf9, 0x66,
	0x11, 0xf2, 0x6c, 0x1a, 0x43, 0xc2, 0x27, 0x20, 0x14, 0xdd, 0x26, 0x8b, 0xc5, 0x78, 0x2d, 0xab,
	0x6b, 0xf5, 0xea, 0x9b, 0x77, 0xdc, 0x00, 0x13, 0x28, 0xee, 0xc1, 0x3d, 0x34, 0xde, 0xbe, 0xfd,
	0xe5, 0xbc, 0x33, 0xc7, 0xca, 0x68, 0xfa, 0x90, 0xd8, 0x32, 0x86, 0xa0, 0x35, 0xaf, 0xa9, 0xae,
	0x7b, 0xf5, 0x06, 0xdd, 0xcb, 0x2a, 0x83, 0x18, 0x02, 0xa6, 0xa3, 0xe9, 0x73, 0x52, 0xcb, 0xf5,
	0x6f, 0x55, 0x34, 0xb7, 0x32, 0x5b, 0x6d, 0xa0, 0x7d, 0xfd, 0xd5, 0xac, 0xd6, 0x8f, 0xf3, 0xce,
	0xff, 0x0a, 0xa4, 0x1a, 0xf2, 0xd1, 0x68, 0xc7, 0xe1, 0xa1, 0xc0, 0x04, 0x1c, 0x66, 0xf0, 0x9d,
	0xf6, 0xd9, 0x85, 0x5d, 0x25, 0x15, 0x98, 0xc6, 0x67, 0x17, 0x76, 0x83, 0xd6, 0xa1, 0xac, 0x26,
	0x9d, 0x8f, 0x15, 0xb2, 0x34, 0x5b, 0x9d, 0xee, 0x93, 0x9a, 0x56, 0x41, 0xb6, 0xac, 0x6e, 0xa5,
	0x57, 0xdf, 0x74, 0x6f, 0xea, 0xd7, 0x3d, 0xd0, 0xf2, 0xc1, 0x70, 0x3f, 0xc3, 0x98, 0xa1, 0xe9,
	0x3e, 0x59, 0x92, 0x0a, 0xe3, 0xa3, 0x00, 0xc5, 0x90, 0x67, 0xe2, 0x9a, 0xf9, 0x3b, 0xd7, 0xf3,
	0x0d, 0x14, 0xc6, 0xbb, 0x45, 0x18, 0x6b, 0xc8, 0x5f, 0x4d, 0xfa, 0x84, 0x54, 0xb3, 0x41, 0xa0,
	0x65, 0x77, 0xad, 0xde, 0xd2, 0xe6, 0xfa, 0x8d, 0xed, 0x64, 0xca, 0x00, 0xcb, 0xa1, 0xf6, 0x7b,
	0x8b, 0x34, 0x66, 0xfa, 0xa3, 0x5b, 0x64, 0x41, 0x42, 0x72, 0xca, 0x03, 0x30, 0xd7, 0xb8, 0xaa,
	0x33, 0x96, 0xd9, 0xf6, 0x40, 0x2a, 0x2e, 0xf4, 0xaa, 0xb0, 0x22, 0x92, 0x3e, 0x25, 0x55, 0x3d,
	0x96, 0x99, 0xe1, 0x9e, 0x6b, 0x1e, 0x4a, 0xfe, 0x8a, 0x66, 0x13, 0x30, 0x4c, 0x15, 0xe8, 0x5a,
	0x92, 0xe5, 0x94, 0xd3, 0x23, 0x55, 0xdd, 0x15, 0x5d, 0x24, 0xf6, 0x81, 0xe0, 0xaa, 0x39, 0x47,
	0xeb, 0x64, 0x81, 0xa5, 0x42, 0x70, 0x11, 0x36, 0xad, 0xec, 0x78, 0x0f, 0x05, 0x34, 0xe7, 0x9d,
	0x0f, 0x16, 0x69, 0xcc, 0xc8, 0x41, 0x1f, 0x93, 0xc5, 0x62, 0x75, 0x2f, 0x1b, 0xd6, 0xbb, 0xed,
	0x16, 0xbb, 0xed, 0xee, 0x99, 0x80, 0xbe, 0xfd, 0xe9, 0x6b, 0xc7, 0x62, 0x25, 0x40, 0x1f, 0x91,
	0xda, 0x04, 0x54, 0xc2, 0xb3, 0xc7, 0x97, 0x5d, 0xe6, 0xdd, 0xeb, 0xea, 0x1d, 0x6a, 0xff, 0xeb,
	0x71, 0x02, 0x72, 0x8c, 0xd1, 0x90, 0x19, 0xc0, 0x79, 0x41, 0x96, 0xaf, 0xb8, 0x68, 0x87, 0xd4,
	0x73, 0xe7, 0x91, 0xf0, 0x27, 0xb9, 0x7c, 0xff, 0x31, 0x92, 0x1f, 0xbd, 0xf4, 0x27, 0x40, 0x57,
	0x48, 0xf5, 0xd4, 0x8f, 0x52, 0xd0, 0x32, 0x59, 0x2c, 0x37, 0x9c, 0x65, 0xd2, 0xd8, 0x1d, 0xfb,
	0x28, 0x07, 0xa0, 0x14, 0x17, 0xa1, 0xec, 0xdf, 0xff, 0xfc, 0x7d, 0xcd, 0x7a, 0xb3, 0xfe, 0xb7,
	0x8f, 0x5f, 0x7c, 0x12, 0x9a, 0x55, 0x3d, 0xae, 0xe9, 0x31, 0xb7, 0x7e, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x2f, 0x12, 0x2f, 0x90, 0x2d, 0x05, 0x00, 0x00,
}

func (this *Experiment) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Experiment)
	if !ok {
		that2, ok := that.(Experiment)
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
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ExperimentSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ExperimentSpec)
	if !ok {
		that2, ok := that.(ExperimentSpec)
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
	if len(this.Faults) != len(that1.Faults) {
		return false
	}
	for i := range this.Faults {
		if !this.Faults[i].Equal(that1.Faults[i]) {
			return false
		}
	}
	if !this.StopCondition.Equal(that1.StopCondition) {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ExperimentSpec_InjectedFault) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ExperimentSpec_InjectedFault)
	if !ok {
		that2, ok := that.(ExperimentSpec_InjectedFault)
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
	if !this.Service.Equal(that1.Service) {
		return false
	}
	if !this.Fault.Equal(that1.Fault) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *StopCondition) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StopCondition)
	if !ok {
		that2, ok := that.(StopCondition)
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
	if this.Duration != nil && that1.Duration != nil {
		if *this.Duration != *that1.Duration {
			return false
		}
	} else if this.Duration != nil {
		return false
	} else if that1.Duration != nil {
		return false
	}
	if len(this.Metric) != len(that1.Metric) {
		return false
	}
	for i := range this.Metric {
		if !this.Metric[i].Equal(that1.Metric[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MetricThreshold) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetricThreshold)
	if !ok {
		that2, ok := that.(MetricThreshold)
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
	if this.MetricName != that1.MetricName {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ChaosSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChaosSettings)
	if !ok {
		that2, ok := that.(ChaosSettings)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

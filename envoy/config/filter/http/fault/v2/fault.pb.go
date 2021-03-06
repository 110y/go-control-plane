// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/fault/v2/fault.proto

package envoy_config_filter_http_fault_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/fault/v2"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FaultAbort struct {
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	//	*FaultAbort_HeaderAbort_
	ErrorType            isFaultAbort_ErrorType   `protobuf_oneof:"error_type"`
	Percentage           *_type.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{0}
}

func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort.Unmarshal(m, b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
}
func (m *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(m, src)
}
func (m *FaultAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort.Size(m)
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

type FaultAbort_HeaderAbort_ struct {
	HeaderAbort *FaultAbort_HeaderAbort `protobuf:"bytes,4,opt,name=header_abort,json=headerAbort,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (*FaultAbort_HeaderAbort_) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetHeaderAbort() *FaultAbort_HeaderAbort {
	if x, ok := m.GetErrorType().(*FaultAbort_HeaderAbort_); ok {
		return x.HeaderAbort
	}
	return nil
}

func (m *FaultAbort) GetPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultAbort_HttpStatus)(nil),
		(*FaultAbort_HeaderAbort_)(nil),
	}
}

type FaultAbort_HeaderAbort struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultAbort_HeaderAbort) Reset()         { *m = FaultAbort_HeaderAbort{} }
func (m *FaultAbort_HeaderAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort_HeaderAbort) ProtoMessage()    {}
func (*FaultAbort_HeaderAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{0, 0}
}

func (m *FaultAbort_HeaderAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort_HeaderAbort.Unmarshal(m, b)
}
func (m *FaultAbort_HeaderAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort_HeaderAbort.Marshal(b, m, deterministic)
}
func (m *FaultAbort_HeaderAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort_HeaderAbort.Merge(m, src)
}
func (m *FaultAbort_HeaderAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort_HeaderAbort.Size(m)
}
func (m *FaultAbort_HeaderAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort_HeaderAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort_HeaderAbort proto.InternalMessageInfo

type HTTPFault struct {
	Delay                           *v2.FaultDelay         `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	Abort                           *FaultAbort            `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	UpstreamCluster                 string                 `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	Headers                         []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	DownstreamNodes                 []string               `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	MaxActiveFaults                 *wrappers.UInt32Value  `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	ResponseRateLimit               *v2.FaultRateLimit     `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	DelayPercentRuntime             string                 `protobuf:"bytes,8,opt,name=delay_percent_runtime,json=delayPercentRuntime,proto3" json:"delay_percent_runtime,omitempty"`
	AbortPercentRuntime             string                 `protobuf:"bytes,9,opt,name=abort_percent_runtime,json=abortPercentRuntime,proto3" json:"abort_percent_runtime,omitempty"`
	DelayDurationRuntime            string                 `protobuf:"bytes,10,opt,name=delay_duration_runtime,json=delayDurationRuntime,proto3" json:"delay_duration_runtime,omitempty"`
	AbortHttpStatusRuntime          string                 `protobuf:"bytes,11,opt,name=abort_http_status_runtime,json=abortHttpStatusRuntime,proto3" json:"abort_http_status_runtime,omitempty"`
	MaxActiveFaultsRuntime          string                 `protobuf:"bytes,12,opt,name=max_active_faults_runtime,json=maxActiveFaultsRuntime,proto3" json:"max_active_faults_runtime,omitempty"`
	ResponseRateLimitPercentRuntime string                 `protobuf:"bytes,13,opt,name=response_rate_limit_percent_runtime,json=responseRateLimitPercentRuntime,proto3" json:"response_rate_limit_percent_runtime,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}               `json:"-"`
	XXX_unrecognized                []byte                 `json:"-"`
	XXX_sizecache                   int32                  `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{1}
}

func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFault.Unmarshal(m, b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
}
func (m *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(m, src)
}
func (m *HTTPFault) XXX_Size() int {
	return xxx_messageInfo_HTTPFault.Size(m)
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func (m *HTTPFault) GetMaxActiveFaults() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxActiveFaults
	}
	return nil
}

func (m *HTTPFault) GetResponseRateLimit() *v2.FaultRateLimit {
	if m != nil {
		return m.ResponseRateLimit
	}
	return nil
}

func (m *HTTPFault) GetDelayPercentRuntime() string {
	if m != nil {
		return m.DelayPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortPercentRuntime() string {
	if m != nil {
		return m.AbortPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetDelayDurationRuntime() string {
	if m != nil {
		return m.DelayDurationRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortHttpStatusRuntime() string {
	if m != nil {
		return m.AbortHttpStatusRuntime
	}
	return ""
}

func (m *HTTPFault) GetMaxActiveFaultsRuntime() string {
	if m != nil {
		return m.MaxActiveFaultsRuntime
	}
	return ""
}

func (m *HTTPFault) GetResponseRateLimitPercentRuntime() string {
	if m != nil {
		return m.ResponseRateLimitPercentRuntime
	}
	return ""
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.config.filter.http.fault.v2.FaultAbort")
	proto.RegisterType((*FaultAbort_HeaderAbort)(nil), "envoy.config.filter.http.fault.v2.FaultAbort.HeaderAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.config.filter.http.fault.v2.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/fault/v2/fault.proto", fileDescriptor_26070db6b6576d5c)
}

var fileDescriptor_26070db6b6576d5c = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x4e, 0x13, 0x41,
	0x14, 0xc6, 0xd9, 0xfe, 0x03, 0xa6, 0x10, 0xca, 0xa2, 0xb8, 0x36, 0x82, 0x05, 0x13, 0x53, 0x8c,
	0xec, 0x26, 0xc5, 0x1b, 0x62, 0x34, 0xa1, 0x10, 0x52, 0x0d, 0x9a, 0x66, 0x45, 0xaf, 0x8c, 0x9b,
	0xa1, 0x3d, 0x6d, 0x37, 0xd9, 0xce, 0x6c, 0x66, 0x67, 0x4b, 0x7b, 0x67, 0xe2, 0x03, 0x78, 0xeb,
	0x4b, 0xf8, 0x02, 0x3e, 0x01, 0xb7, 0xde, 0xf9, 0x1c, 0x5e, 0x19, 0x2f, 0x8c, 0x99, 0x33, 0xbb,
	0xa5, 0xd0, 0x26, 0xe8, 0x4d, 0x33, 0x9d, 0x73, 0x7e, 0xe7, 0x9b, 0xf3, 0xcd, 0x99, 0x25, 0xbb,
	0xc0, 0x06, 0x7c, 0xe4, 0xb4, 0x38, 0xeb, 0xf8, 0x5d, 0xa7, 0xe3, 0x07, 0x12, 0x84, 0xd3, 0x93,
	0x32, 0x74, 0x3a, 0x34, 0x0e, 0xa4, 0x33, 0xa8, 0xe9, 0x85, 0x1d, 0x0a, 0x2e, 0xb9, 0xb9, 0x85,
	0xe9, 0xb6, 0x4e, 0xb7, 0x75, 0xba, 0xad, 0xd2, 0x6d, 0x9d, 0x35, 0xa8, 0x95, 0x77, 0x74, 0x45,
	0x1a, 0xfa, 0x0a, 0x16, 0x3c, 0x96, 0xa0, 0x7f, 0xbd, 0x16, 0xef, 0x87, 0x9c, 0x01, 0x93, 0x91,
	0xae, 0x56, 0xae, 0xce, 0x12, 0x9f, 0xa5, 0x5b, 0xb6, 0x74, 0xa6, 0x1c, 0x85, 0xe0, 0x84, 0x20,
	0x5a, 0xc0, 0xd2, 0xc8, 0x66, 0x97, 0xf3, 0x6e, 0x00, 0x0e, 0xfe, 0x3b, 0x8b, 0x3b, 0xce, 0xb9,
	0xa0, 0x61, 0x08, 0x22, 0xd5, 0xd8, 0x8c, 0xdb, 0x21, 0x75, 0x28, 0x63, 0x5c, 0x52, 0xe9, 0x73,
	0x16, 0x39, 0x7d, 0xbf, 0x2b, 0xa8, 0x84, 0x24, 0xbe, 0x31, 0x15, 0x8f, 0x24, 0x95, 0x71, 0x8a,
	0xdf, 0x19, 0xd0, 0xc0, 0x6f, 0x53, 0x09, 0x4e, 0xba, 0xd0, 0x81, 0xed, 0x4f, 0x19, 0x42, 0x8e,
	0xd5, 0x09, 0x0f, 0xce, 0xb8, 0x90, 0xa6, 0x4d, 0x8a, 0xca, 0x06, 0x4f, 0xc3, 0x56, 0xa6, 0x62,
	0x54, 0x97, 0xeb, 0xc5, 0xdf, 0xf5, 0x85, 0x47, 0x85, 0xd2, 0x8f, 0x5c, 0xf5, 0xc2, 0x68, 0xcc,
	0xb9, 0x44, 0x65, 0xbc, 0xc1, 0x04, 0xf3, 0x03, 0x59, 0xea, 0x01, 0x6d, 0x83, 0xf0, 0xa8, 0xe2,
	0xad, 0x5c, 0xc5, 0xa8, 0x16, 0x6b, 0xfb, 0xf6, 0x8d, 0xfe, 0xda, 0x97, 0xa2, 0x76, 0x03, 0x2b,
	0xe0, 0xba, 0x31, 0xe7, 0x16, 0x7b, 0x97, 0x7f, 0xcd, 0x67, 0x84, 0x24, 0x3e, 0xd1, 0x2e, 0x58,
	0x59, 0xac, 0xbe, 0x91, 0x54, 0x57, 0x2e, 0xda, 0xc7, 0x82, 0xb6, 0x54, 0xbf, 0x34, 0x68, 0xea,
	0x3c, 0x77, 0x02, 0x28, 0x2f, 0x93, 0xe2, 0x44, 0xf1, 0xfa, 0x2a, 0x21, 0x20, 0x04, 0x17, 0x9e,
	0x42, 0xcd, 0xec, 0xaf, 0xba, 0xf1, 0x32, 0xb7, 0x60, 0x94, 0x32, 0xdb, 0x5f, 0x0b, 0x64, 0xb1,
	0x71, 0x7a, 0xda, 0xc4, 0x43, 0x99, 0xcf, 0x49, 0xbe, 0x0d, 0x01, 0x1d, 0x59, 0x06, 0xea, 0x55,
	0x67, 0x76, 0x73, 0xb5, 0x91, 0x23, 0x95, 0xef, 0x6a, 0xcc, 0x3c, 0x24, 0x79, 0xed, 0x46, 0x06,
	0xf9, 0xdd, 0xff, 0x72, 0xc3, 0xd5, 0xac, 0xb9, 0x43, 0x4a, 0x71, 0x18, 0x49, 0x01, 0xb4, 0xef,
	0xb5, 0x82, 0x38, 0x92, 0x20, 0xb0, 0xff, 0x45, 0x77, 0x25, 0xdd, 0x3f, 0xd4, 0xdb, 0xe6, 0x53,
	0x32, 0xaf, 0x3d, 0x8b, 0xac, 0x5c, 0x25, 0x5b, 0x2d, 0xd6, 0xb6, 0x12, 0x45, 0x1a, 0xfa, 0xaa,
	0x38, 0x8e, 0x6d, 0xe2, 0xf2, 0x2b, 0x2a, 0x5b, 0x3d, 0x10, 0x6e, 0x4a, 0x28, 0x9d, 0x36, 0x3f,
	0x67, 0x89, 0x12, 0xe3, 0x6d, 0x88, 0xac, 0x7c, 0x25, 0xab, 0x74, 0x2e, 0xf7, 0x5f, 0xab, 0x6d,
	0xb3, 0x41, 0x56, 0xfb, 0x74, 0xe8, 0x29, 0xc3, 0x07, 0xe0, 0xe1, 0xd9, 0x23, 0xab, 0x80, 0x3d,
	0xde, 0xb3, 0xf5, 0xfc, 0xda, 0xe9, 0xfc, 0xda, 0x6f, 0x5f, 0x30, 0xb9, 0x57, 0x7b, 0x47, 0x83,
	0x18, 0xdc, 0x95, 0x3e, 0x1d, 0x1e, 0x20, 0x85, 0x7d, 0x46, 0xe6, 0x7b, 0xb2, 0x26, 0x20, 0x0a,
	0x39, 0x8b, 0xc0, 0x53, 0x43, 0xec, 0x05, 0x7e, 0xdf, 0x97, 0xd6, 0x3c, 0xd6, 0x7a, 0xfc, 0x0f,
	0x7e, 0xbb, 0x54, 0xc2, 0x89, 0x62, 0xdc, 0xd5, 0xb4, 0xd0, 0x78, 0xcb, 0xac, 0x91, 0xdb, 0x78,
	0x11, 0x5e, 0x32, 0x09, 0x9e, 0x88, 0x99, 0xf4, 0xfb, 0x60, 0x2d, 0xa0, 0x7f, 0x6b, 0x18, 0x4c,
	0xc7, 0x45, 0x87, 0x14, 0x83, 0xbe, 0x4f, 0x31, 0x8b, 0x9a, 0xc1, 0xe0, 0x35, 0xe6, 0x09, 0x59,
	0xd7, 0x3a, 0xed, 0x58, 0xe0, 0xa3, 0x1b, 0x43, 0x04, 0xa1, 0x5b, 0x18, 0x3d, 0x4a, 0x82, 0x29,
	0xb5, 0x4f, 0xee, 0x6a, 0xa5, 0x89, 0x87, 0x36, 0x06, 0x8b, 0x08, 0xae, 0x63, 0x42, 0x63, 0xfc,
	0xcc, 0x26, 0xd0, 0xa9, 0x0b, 0x18, 0xa3, 0x4b, 0x1a, 0xbd, 0x66, 0x75, 0x8a, 0x9e, 0x90, 0x07,
	0x33, 0x1c, 0x9f, 0xea, 0x76, 0x19, 0x8b, 0xdc, 0x9f, 0xf2, 0xf4, 0x6a, 0xe7, 0xf5, 0xf8, 0xe7,
	0x97, 0x3f, 0x9f, 0xf3, 0x55, 0xf3, 0xa1, 0xbe, 0x29, 0x18, 0x4a, 0x60, 0x91, 0xfa, 0xea, 0x24,
	0xb7, 0x15, 0x5d, 0x19, 0xef, 0xbd, 0x6f, 0x1f, 0x2f, 0xbe, 0x17, 0x32, 0x25, 0x83, 0x38, 0x3e,
	0xd7, 0x97, 0x1b, 0x0a, 0x3e, 0x1c, 0xdd, 0xfc, 0x2e, 0xea, 0xfa, 0xdb, 0xd4, 0x54, 0x43, 0xd5,
	0x34, 0xce, 0x0a, 0x38, 0x5d, 0x7b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x95, 0x3e, 0x4f,
	0xee, 0x05, 0x00, 0x00,
}

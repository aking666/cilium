// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/health_check/v2/health_check.proto

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	envoy/config/filter/http/health_check/v2/health_check.proto

It has these top-level messages:
	HealthCheck
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"
import envoy_api_v2_route "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/route"
import envoy_type1 "github.com/cilium/cilium/pkg/envoy/envoy/type"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HealthCheck struct {
	// Specifies whether the filter operates in pass through mode or not.
	PassThroughMode *google_protobuf1.BoolValue `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode" json:"pass_through_mode,omitempty"`
	// Specifies the incoming HTTP endpoint that should be considered the
	// health check endpoint. For example */healthcheck*.
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint" json:"endpoint,omitempty"`
	// If operating in pass through mode, the amount of time in milliseconds
	// that the filter should cache the upstream response.
	CacheTime *google_protobuf.Duration `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime" json:"cache_time,omitempty"`
	// If operating in non-pass-through mode, specifies a set of upstream cluster
	// names and the minimum percentage of servers in each of those clusters that
	// must be healthy in order for the filter to return a 200.
	ClusterMinHealthyPercentages map[string]*envoy_type1.Percent `protobuf:"bytes,4,rep,name=cluster_min_healthy_percentages,json=clusterMinHealthyPercentages" json:"cluster_min_healthy_percentages,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// [#not-implemented-hide:]
	// Specifies a set of health check request headers to match on. The health check filter will
	// check a request’s headers against all the specified headers. To specify the health check
	// endpoint, set the ``:path`` header to match on. Note that if the
	// :ref:`endpoint <envoy_api_field_config.filter.http.health_check.v2.HealthCheck.endpoint>`
	// field is set, it will overwrite any ``:path`` header to match.
	Headers []*envoy_api_v2_route.HeaderMatcher `protobuf:"bytes,5,rep,name=headers" json:"headers,omitempty"`
}

func (m *HealthCheck) Reset()                    { *m = HealthCheck{} }
func (m *HealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()               {}
func (*HealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HealthCheck) GetPassThroughMode() *google_protobuf1.BoolValue {
	if m != nil {
		return m.PassThroughMode
	}
	return nil
}

func (m *HealthCheck) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *HealthCheck) GetCacheTime() *google_protobuf.Duration {
	if m != nil {
		return m.CacheTime
	}
	return nil
}

func (m *HealthCheck) GetClusterMinHealthyPercentages() map[string]*envoy_type1.Percent {
	if m != nil {
		return m.ClusterMinHealthyPercentages
	}
	return nil
}

func (m *HealthCheck) GetHeaders() []*envoy_api_v2_route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.filter.http.health_check.v2.HealthCheck")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/health_check/v2/health_check.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x8b, 0xd4, 0x30,
	0x18, 0x25, 0x9d, 0x8e, 0x3a, 0xe9, 0xc1, 0xb5, 0x0a, 0xd6, 0x41, 0xdc, 0x59, 0x2f, 0x8e, 0x97,
	0x04, 0xea, 0x45, 0x5c, 0xf0, 0xd0, 0x55, 0xd8, 0xcb, 0x80, 0x94, 0x45, 0xc1, 0x4b, 0xc9, 0xb6,
	0xdf, 0xb4, 0x61, 0x3b, 0x49, 0x48, 0xd3, 0x4a, 0xff, 0x82, 0x07, 0xcf, 0x9e, 0xfc, 0x21, 0x9e,
	0xfc, 0x27, 0x9e, 0xfd, 0x17, 0x92, 0xa4, 0xa3, 0x2e, 0x03, 0xea, 0x65, 0xf8, 0x26, 0xef, 0x7b,
	0xdf, 0x7b, 0xbc, 0x57, 0x7c, 0x0a, 0x62, 0x90, 0x23, 0x2d, 0xa5, 0xd8, 0xf2, 0x9a, 0x6e, 0x79,
	0x6b, 0x40, 0xd3, 0xc6, 0x18, 0x45, 0x1b, 0x60, 0xad, 0x69, 0x8a, 0xb2, 0x81, 0xf2, 0x8a, 0x0e,
	0xe9, 0xb5, 0xff, 0x44, 0x69, 0x69, 0x64, 0xbc, 0x76, 0x64, 0xe2, 0xc9, 0xc4, 0x93, 0x89, 0x25,
	0x93, 0x6b, 0xcb, 0x43, 0xba, 0x7c, 0x54, 0x4b, 0x59, 0xb7, 0x40, 0x1d, 0xef, 0xb2, 0xdf, 0xd2,
	0xaa, 0xd7, 0xcc, 0x70, 0x29, 0xfc, 0xa5, 0x43, 0xfc, 0x83, 0x66, 0x4a, 0x81, 0xee, 0xf6, 0xb8,
	0xb7, 0xc9, 0x14, 0xb7, 0x56, 0xb4, 0xec, 0x0d, 0xf8, 0xdf, 0x09, 0x4f, 0x3c, 0x6e, 0x46, 0x05,
	0x54, 0x81, 0x2e, 0x41, 0x98, 0x09, 0xb9, 0x3f, 0xb0, 0x96, 0x57, 0xcc, 0x00, 0xdd, 0x0f, 0x13,
	0x70, 0xaf, 0x96, 0xb5, 0x74, 0x23, 0xb5, 0x93, 0x7f, 0x7d, 0xfc, 0x29, 0xc4, 0xd1, 0xb9, 0x33,
	0x7f, 0x66, 0xbd, 0xc7, 0x39, 0xbe, 0xa3, 0x58, 0xd7, 0x15, 0xa6, 0xd1, 0xb2, 0xaf, 0x9b, 0x62,
	0x27, 0x2b, 0x48, 0xd0, 0x0a, 0xad, 0xa3, 0x74, 0x49, 0xbc, 0x69, 0xb2, 0x37, 0x4d, 0x32, 0x29,
	0xdb, 0xb7, 0xac, 0xed, 0x21, 0xc3, 0x5f, 0x7f, 0x7c, 0x9b, 0xcd, 0x3f, 0xa2, 0xe0, 0x08, 0xe5,
	0xb7, 0xed, 0x81, 0x0b, 0xcf, 0xdf, 0xc8, 0x0a, 0xe2, 0x27, 0xf8, 0x16, 0x88, 0x4a, 0x49, 0x2e,
	0x4c, 0x12, 0xac, 0xd0, 0x7a, 0x91, 0x45, 0x76, 0x3d, 0xd4, 0xc1, 0x0a, 0x25, 0x28, 0xff, 0x05,
	0xc6, 0x2f, 0x31, 0x2e, 0x59, 0xd9, 0x40, 0x61, 0xf8, 0x0e, 0x92, 0x99, 0x53, 0x7d, 0x70, 0xa0,
	0xfa, 0x6a, 0x8a, 0x32, 0x0b, 0x3f, 0x7f, 0x3f, 0x46, 0xf9, 0xc2, 0x51, 0x2e, 0xf8, 0x0e, 0xe2,
	0x2f, 0x08, 0x1f, 0x97, 0x6d, 0xdf, 0x19, 0xd0, 0xc5, 0x8e, 0x8b, 0xc2, 0xb7, 0x32, 0x16, 0x53,
	0x42, 0xac, 0x86, 0x2e, 0x09, 0x57, 0xb3, 0x75, 0x94, 0xbe, 0x23, 0xff, 0x5b, 0x25, 0xf9, 0x23,
	0x1d, 0x72, 0xe6, 0x8f, 0x6f, 0xb8, 0xf0, 0xaf, 0xe3, 0x9b, 0xdf, 0x97, 0x5f, 0x0b, 0xa3, 0xc7,
	0xfc, 0x61, 0xf9, 0x97, 0x95, 0xf8, 0x14, 0xdf, 0x6c, 0x80, 0x55, 0xa0, 0xbb, 0x64, 0xee, 0x7c,
	0x9c, 0x4c, 0x3e, 0x98, 0xe2, 0x56, 0xcb, 0x57, 0x7c, 0xee, 0x56, 0x36, 0xcc, 0x94, 0x0d, 0xe8,
	0x7c, 0xcf, 0x58, 0x56, 0xf8, 0xe4, 0x9f, 0xfa, 0xf1, 0x11, 0x9e, 0x5d, 0xc1, 0xe8, 0x1a, 0x5b,
	0xe4, 0x76, 0x8c, 0x9f, 0xe2, 0xf9, 0x60, 0x3b, 0x72, 0xd1, 0x47, 0xe9, 0xdd, 0x49, 0xd1, 0x7e,
	0x3a, 0x64, 0xa2, 0xe7, 0x7e, 0xe3, 0x45, 0xf0, 0x1c, 0x65, 0xe1, 0xfb, 0x60, 0x48, 0x2f, 0x6f,
	0xb8, 0xb4, 0x9f, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xda, 0x62, 0xf0, 0x2f, 0x03, 0x00,
	0x00,
}

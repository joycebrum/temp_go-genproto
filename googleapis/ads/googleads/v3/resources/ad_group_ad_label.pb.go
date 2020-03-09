// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/ad_group_ad_label.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A relationship between an ad group ad and a label.
type AdGroupAdLabel struct {
	// The resource name of the ad group ad label.
	// Ad group ad label resource names have the form:
	// `customers/{customer_id}/adGroupAdLabels/{ad_group_id}~{ad_id}~{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ad group ad to which the label is attached.
	AdGroupAd *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
	// The label assigned to the ad group ad.
	Label                *wrappers.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdGroupAdLabel) Reset()         { *m = AdGroupAdLabel{} }
func (m *AdGroupAdLabel) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdLabel) ProtoMessage()    {}
func (*AdGroupAdLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_668144ba7a9ba42f, []int{0}
}

func (m *AdGroupAdLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdLabel.Unmarshal(m, b)
}
func (m *AdGroupAdLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdLabel.Marshal(b, m, deterministic)
}
func (m *AdGroupAdLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdLabel.Merge(m, src)
}
func (m *AdGroupAdLabel) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdLabel.Size(m)
}
func (m *AdGroupAdLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdLabel.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdLabel proto.InternalMessageInfo

func (m *AdGroupAdLabel) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupAdLabel) GetAdGroupAd() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupAd
	}
	return nil
}

func (m *AdGroupAdLabel) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*AdGroupAdLabel)(nil), "google.ads.googleads.v3.resources.AdGroupAdLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/ad_group_ad_label.proto", fileDescriptor_668144ba7a9ba42f)
}

var fileDescriptor_668144ba7a9ba42f = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0xb1, 0xc3, 0x06, 0x71, 0xb6, 0x1d, 0xbc, 0x4b, 0x16, 0xc2, 0x48, 0x36, 0xc2, 0x72,
	0x92, 0x20, 0xbe, 0x6c, 0xda, 0x2e, 0xca, 0x25, 0x30, 0x4a, 0x09, 0x29, 0xf8, 0x50, 0x0c, 0x46,
	0xb1, 0x14, 0x61, 0xb0, 0x2d, 0x23, 0xd9, 0xe9, 0x21, 0x04, 0xfa, 0x0d, 0xfa, 0x1d, 0x7a, 0xec,
	0x47, 0xe9, 0x47, 0xc9, 0xa7, 0x28, 0xfe, 0x23, 0x25, 0x69, 0xa1, 0xed, 0xed, 0xb1, 0xdf, 0xdf,
	0xf3, 0xbe, 0xcf, 0x2b, 0xc9, 0xf9, 0xc3, 0x85, 0xe0, 0x09, 0x83, 0x84, 0x2a, 0xd8, 0xc8, 0x4a,
	0x6d, 0x3d, 0x28, 0x99, 0x12, 0xa5, 0x8c, 0x98, 0x82, 0x84, 0x86, 0x5c, 0x8a, 0x32, 0x0f, 0x09,
	0x0d, 0x13, 0xb2, 0x66, 0x09, 0xc8, 0xa5, 0x28, 0x84, 0x3b, 0x6e, 0x78, 0x40, 0xa8, 0x02, 0xc6,
	0x0a, 0xb6, 0x1e, 0x30, 0xd6, 0xc1, 0x37, 0xdd, 0x3d, 0x8f, 0x4d, 0xc3, 0xc6, 0x3d, 0xf8, 0xde,
	0x96, 0xea, 0xaf, 0x75, 0xb9, 0x81, 0x37, 0x92, 0xe4, 0x39, 0x93, 0xaa, 0xad, 0x0f, 0x4f, 0xac,
	0x24, 0xcb, 0x44, 0x41, 0x8a, 0x58, 0x64, 0x6d, 0xf5, 0xc7, 0x9d, 0xed, 0x7c, 0xc1, 0x74, 0x51,
	0xc5, 0xc2, 0xf4, 0xa2, 0x0a, 0xe5, 0xfe, 0x74, 0x3e, 0xeb, 0x11, 0x61, 0x46, 0x52, 0xd6, 0xb7,
	0x46, 0xd6, 0xb4, 0xbb, 0xfa, 0xa4, 0x7f, 0x5e, 0x92, 0x94, 0xb9, 0xff, 0x9c, 0xde, 0xc9, 0x3a,
	0x7d, 0x7b, 0x64, 0x4d, 0x7b, 0xb3, 0x61, 0x1b, 0x1f, 0xe8, 0x2c, 0xe0, 0xaa, 0x90, 0x71, 0xc6,
	0x7d, 0x92, 0x94, 0x6c, 0xd5, 0x25, 0x7a, 0x8e, 0x3b, 0x73, 0x3e, 0xd4, 0x07, 0xd0, 0xef, 0xbc,
	0xc3, 0xd7, 0xa0, 0x68, 0x73, 0xc0, 0x91, 0xf3, 0xeb, 0x78, 0x3e, 0xad, 0xca, 0x63, 0x05, 0x22,
	0x91, 0xc2, 0x67, 0x4b, 0xfc, 0x8e, 0x4a, 0x55, 0x88, 0x94, 0x49, 0x05, 0x77, 0x5a, 0xee, 0x21,
	0x39, 0x83, 0x14, 0xdc, 0xbd, 0xb8, 0x92, 0xfd, 0xfc, 0xd6, 0x76, 0x26, 0x91, 0x48, 0xc1, 0x9b,
	0x97, 0x32, 0xff, 0x7a, 0x3e, 0x73, 0x59, 0x85, 0x5f, 0x5a, 0xd7, 0xff, 0x5b, 0x27, 0x17, 0x09,
	0xc9, 0x38, 0x10, 0x92, 0x43, 0xce, 0xb2, 0x7a, 0x35, 0x78, 0x0c, 0xfd, 0xca, 0x43, 0xf9, 0x6b,
	0xd4, 0xbd, 0xdd, 0x59, 0x60, 0xfc, 0x60, 0x8f, 0x17, 0x4d, 0x4b, 0x4c, 0x15, 0x68, 0x64, 0xa5,
	0x7c, 0x0f, 0xac, 0x34, 0xf9, 0xa8, 0x99, 0x00, 0x53, 0x15, 0x18, 0x26, 0xf0, 0xbd, 0xc0, 0x30,
	0x07, 0x7b, 0xd2, 0x14, 0x10, 0xc2, 0x54, 0x21, 0x64, 0x28, 0x84, 0x7c, 0x0f, 0x21, 0xc3, 0xad,
	0x3f, 0xd6, 0x61, 0xbd, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xbe, 0x8e, 0xe2, 0xd4, 0x02,
	0x00, 0x00,
}
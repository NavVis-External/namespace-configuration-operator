// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/hotel_date_selection_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible hotel date selection types.
type HotelDateSelectionTypeEnum_HotelDateSelectionType int32

const (
	// Not specified.
	HotelDateSelectionTypeEnum_UNSPECIFIED HotelDateSelectionTypeEnum_HotelDateSelectionType = 0
	// Used for return value only. Represents value unknown in this version.
	HotelDateSelectionTypeEnum_UNKNOWN HotelDateSelectionTypeEnum_HotelDateSelectionType = 1
	// Dates selected by default.
	HotelDateSelectionTypeEnum_DEFAULT_SELECTION HotelDateSelectionTypeEnum_HotelDateSelectionType = 50
	// Dates selected by the user.
	HotelDateSelectionTypeEnum_USER_SELECTED HotelDateSelectionTypeEnum_HotelDateSelectionType = 51
)

var HotelDateSelectionTypeEnum_HotelDateSelectionType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	50: "DEFAULT_SELECTION",
	51: "USER_SELECTED",
}
var HotelDateSelectionTypeEnum_HotelDateSelectionType_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"DEFAULT_SELECTION": 50,
	"USER_SELECTED":     51,
}

func (x HotelDateSelectionTypeEnum_HotelDateSelectionType) String() string {
	return proto.EnumName(HotelDateSelectionTypeEnum_HotelDateSelectionType_name, int32(x))
}
func (HotelDateSelectionTypeEnum_HotelDateSelectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_hotel_date_selection_type_736ea29770a336eb, []int{0, 0}
}

// Container for enum describing possible hotel date selection types
type HotelDateSelectionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelDateSelectionTypeEnum) Reset()         { *m = HotelDateSelectionTypeEnum{} }
func (m *HotelDateSelectionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*HotelDateSelectionTypeEnum) ProtoMessage()    {}
func (*HotelDateSelectionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_hotel_date_selection_type_736ea29770a336eb, []int{0}
}
func (m *HotelDateSelectionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelDateSelectionTypeEnum.Unmarshal(m, b)
}
func (m *HotelDateSelectionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelDateSelectionTypeEnum.Marshal(b, m, deterministic)
}
func (dst *HotelDateSelectionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelDateSelectionTypeEnum.Merge(dst, src)
}
func (m *HotelDateSelectionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_HotelDateSelectionTypeEnum.Size(m)
}
func (m *HotelDateSelectionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelDateSelectionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_HotelDateSelectionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HotelDateSelectionTypeEnum)(nil), "google.ads.googleads.v0.enums.HotelDateSelectionTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.HotelDateSelectionTypeEnum_HotelDateSelectionType", HotelDateSelectionTypeEnum_HotelDateSelectionType_name, HotelDateSelectionTypeEnum_HotelDateSelectionType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/hotel_date_selection_type.proto", fileDescriptor_hotel_date_selection_type_736ea29770a336eb)
}

var fileDescriptor_hotel_date_selection_type_736ea29770a336eb = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0x8c, 0xfc, 0x92, 0xd4, 0x9c, 0xf8, 0x94, 0xc4, 0x92, 0xd4, 0xf8, 0xe2,
	0xd4, 0x9c, 0xd4, 0xe4, 0x92, 0xcc, 0xfc, 0xbc, 0xf8, 0x92, 0xca, 0x82, 0x54, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x1e, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x76, 0xbd, 0x32,
	0x03, 0x3d, 0xb0, 0x76, 0xa5, 0x3a, 0x2e, 0x29, 0x0f, 0x90, 0x09, 0x2e, 0x89, 0x25, 0xa9, 0xc1,
	0x30, 0xfd, 0x21, 0x95, 0x05, 0xa9, 0xae, 0x79, 0xa5, 0xb9, 0x4a, 0x09, 0x5c, 0x62, 0xd8, 0x65,
	0x85, 0xf8, 0xb9, 0xb8, 0x43, 0xfd, 0x82, 0x03, 0x5c, 0x9d, 0x3d, 0xdd, 0x3c, 0x5d, 0x5d, 0x04,
	0x18, 0x84, 0xb8, 0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x85, 0x44,
	0xb9, 0x04, 0x5d, 0x5c, 0xdd, 0x1c, 0x43, 0x7d, 0x42, 0xe2, 0x83, 0x5d, 0x7d, 0x5c, 0x9d, 0x43,
	0x3c, 0xfd, 0xfd, 0x04, 0x8c, 0x84, 0x04, 0xb9, 0x78, 0x43, 0x83, 0x5d, 0x83, 0xa0, 0x62, 0xae,
	0x2e, 0x02, 0xc6, 0x4e, 0x1f, 0x18, 0xb9, 0x14, 0x93, 0xf3, 0x73, 0xf5, 0xf0, 0xba, 0xd2, 0x49,
	0x1a, 0xbb, 0x2b, 0x02, 0x40, 0x3e, 0x0c, 0x60, 0x8c, 0x72, 0x82, 0xea, 0x4e, 0xcf, 0xcf, 0x49,
	0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f, 0xcd, 0x03, 0xfb, 0x1f, 0x16, 0x64, 0x05,
	0x99, 0xc5, 0x38, 0x42, 0xd0, 0x1a, 0x4c, 0x2e, 0x62, 0x62, 0x76, 0x77, 0x74, 0x5c, 0xc5, 0x24,
	0xeb, 0x0e, 0x31, 0xca, 0x31, 0xa5, 0x58, 0x0f, 0xc2, 0x04, 0xb1, 0xc2, 0x0c, 0xf4, 0x40, 0xe1,
	0x51, 0x7c, 0x0a, 0x26, 0x1f, 0xe3, 0x98, 0x52, 0x1c, 0x03, 0x97, 0x8f, 0x09, 0x33, 0x88, 0x01,
	0xcb, 0xbf, 0x62, 0x52, 0x84, 0x08, 0x5a, 0x59, 0x39, 0xa6, 0x14, 0x5b, 0x59, 0xc1, 0x55, 0x58,
	0x59, 0x85, 0x19, 0x58, 0x59, 0x81, 0xd5, 0x24, 0xb1, 0x81, 0x1d, 0x66, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xbf, 0xf6, 0x09, 0x3d, 0xd9, 0x01, 0x00, 0x00,
}

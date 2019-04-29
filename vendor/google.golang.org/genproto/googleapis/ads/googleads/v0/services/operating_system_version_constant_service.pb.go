// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/operating_system_version_constant_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for
// [OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant][google.ads.googleads.v0.services.OperatingSystemVersionConstantService.GetOperatingSystemVersionConstant].
type GetOperatingSystemVersionConstantRequest struct {
	// Resource name of the OS version to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperatingSystemVersionConstantRequest) Reset() {
	*m = GetOperatingSystemVersionConstantRequest{}
}
func (m *GetOperatingSystemVersionConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperatingSystemVersionConstantRequest) ProtoMessage()    {}
func (*GetOperatingSystemVersionConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_operating_system_version_constant_service_3ab31064b18c22b0, []int{0}
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Unmarshal(m, b)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Marshal(b, m, deterministic)
}
func (dst *GetOperatingSystemVersionConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Merge(dst, src)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperatingSystemVersionConstantRequest.Size(m)
}
func (m *GetOperatingSystemVersionConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperatingSystemVersionConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperatingSystemVersionConstantRequest proto.InternalMessageInfo

func (m *GetOperatingSystemVersionConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOperatingSystemVersionConstantRequest)(nil), "google.ads.googleads.v0.services.GetOperatingSystemVersionConstantRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperatingSystemVersionConstantServiceClient is the client API for OperatingSystemVersionConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperatingSystemVersionConstantServiceClient interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error)
}

type operatingSystemVersionConstantServiceClient struct {
	cc *grpc.ClientConn
}

func NewOperatingSystemVersionConstantServiceClient(cc *grpc.ClientConn) OperatingSystemVersionConstantServiceClient {
	return &operatingSystemVersionConstantServiceClient{cc}
}

func (c *operatingSystemVersionConstantServiceClient) GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error) {
	out := new(resources.OperatingSystemVersionConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatingSystemVersionConstantServiceServer is the server API for OperatingSystemVersionConstantService service.
type OperatingSystemVersionConstantServiceServer interface {
	// Returns the requested OS version constant in full detail.
	GetOperatingSystemVersionConstant(context.Context, *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error)
}

func RegisterOperatingSystemVersionConstantServiceServer(s *grpc.Server, srv OperatingSystemVersionConstantServiceServer) {
	s.RegisterService(&_OperatingSystemVersionConstantService_serviceDesc, srv)
}

func _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatingSystemVersionConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, req.(*GetOperatingSystemVersionConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperatingSystemVersionConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.OperatingSystemVersionConstantService",
	HandlerType: (*OperatingSystemVersionConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperatingSystemVersionConstant",
			Handler:    _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/operating_system_version_constant_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/operating_system_version_constant_service.proto", fileDescriptor_operating_system_version_constant_service_3ab31064b18c22b0)
}

var fileDescriptor_operating_system_version_constant_service_3ab31064b18c22b0 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0x3e, 0xf8, 0xc0, 0xa0, 0x9b, 0xac, 0xa4, 0xb8, 0xa8, 0xb5, 0x42, 0xe9, 0x62,
	0x12, 0x14, 0x11, 0x46, 0xba, 0x48, 0x5d, 0x54, 0x5d, 0xd8, 0xd2, 0x42, 0x16, 0x12, 0x08, 0x63,
	0x33, 0x84, 0x40, 0x33, 0x37, 0xe6, 0x4e, 0x03, 0x22, 0x6e, 0x7c, 0x03, 0x71, 0xe9, 0xce, 0xa5,
	0x8f, 0xe2, 0xd6, 0x57, 0x70, 0xe5, 0xce, 0x37, 0x90, 0x74, 0x32, 0x01, 0x17, 0xfd, 0xb3, 0x3b,
	0x4c, 0x4e, 0x7e, 0x67, 0xee, 0xb9, 0x63, 0x8d, 0x62, 0x80, 0x78, 0xc6, 0x1d, 0x16, 0xa1, 0xa3,
	0x64, 0xa9, 0x0a, 0xd7, 0x41, 0x9e, 0x17, 0xc9, 0x94, 0xa3, 0x03, 0x19, 0xcf, 0x99, 0x4c, 0x44,
	0x1c, 0xe2, 0x3d, 0x4a, 0x9e, 0x86, 0x05, 0xcf, 0x31, 0x01, 0x11, 0x4e, 0x41, 0xa0, 0x64, 0x42,
	0x86, 0x95, 0x95, 0x64, 0x39, 0x48, 0xb0, 0x9b, 0x0a, 0x43, 0x58, 0x84, 0xa4, 0x26, 0x92, 0xc2,
	0x25, 0x9a, 0xd8, 0xb8, 0x5c, 0x96, 0x99, 0x73, 0x84, 0x79, 0xbe, 0x51, 0xa8, 0x0a, 0x6b, 0xec,
	0x69, 0x54, 0x96, 0x38, 0x4c, 0x08, 0x90, 0x4c, 0x26, 0x20, 0x50, 0x7d, 0x6d, 0x0d, 0xad, 0xce,
	0x80, 0xcb, 0xa1, 0x66, 0x4d, 0x16, 0x28, 0x5f, 0x91, 0xce, 0x2b, 0xd0, 0x98, 0xdf, 0xcd, 0x39,
	0x4a, 0xfb, 0xc0, 0xda, 0xd1, 0xf1, 0xa1, 0x60, 0x29, 0xdf, 0x35, 0x9a, 0x46, 0x67, 0x6b, 0xbc,
	0xad, 0x0f, 0xaf, 0x59, 0xca, 0x8f, 0x5e, 0x4d, 0xeb, 0x70, 0x35, 0x6e, 0xa2, 0x86, 0xb4, 0x7f,
	0x0c, 0x6b, 0x7f, 0x6d, 0xb6, 0x7d, 0x45, 0xd6, 0x95, 0x45, 0x36, 0x1d, 0xa0, 0xe1, 0x2d, 0x65,
	0xd5, 0xb5, 0x92, 0xd5, 0xa4, 0x56, 0xef, 0xe9, 0xf3, 0xeb, 0xc5, 0x3c, 0xb5, 0x4f, 0xca, 0x65,
	0x3c, 0xfc, 0xa9, 0xa3, 0x07, 0x2b, 0x7f, 0x45, 0xa7, 0xfb, 0xd8, 0x7f, 0x36, 0xad, 0xf6, 0x14,
	0xd2, 0xb5, 0x33, 0xf5, 0xbb, 0x1b, 0x75, 0x38, 0x2a, 0x77, 0x38, 0x32, 0x6e, 0x2e, 0x2a, 0x5e,
	0x0c, 0x33, 0x26, 0x62, 0x02, 0x79, 0xec, 0xc4, 0x5c, 0x2c, 0x36, 0xac, 0x9f, 0x4f, 0x96, 0xe0,
	0xf2, 0x17, 0x7c, 0xa6, 0xc5, 0x9b, 0xf9, 0x6f, 0xe0, 0x79, 0xef, 0x66, 0x73, 0xa0, 0x80, 0x5e,
	0x84, 0x44, 0xc9, 0x52, 0xf9, 0x2e, 0xa9, 0x82, 0xf1, 0x43, 0x5b, 0x02, 0x2f, 0xc2, 0xa0, 0xb6,
	0x04, 0xbe, 0x1b, 0x68, 0xcb, 0xb7, 0xd9, 0x56, 0xe7, 0x94, 0x7a, 0x11, 0x52, 0x5a, 0x9b, 0x28,
	0xf5, 0x5d, 0x4a, 0xb5, 0xed, 0xf6, 0xff, 0xe2, 0x9e, 0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x15, 0x9b, 0xe8, 0x12, 0x68, 0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/managed_placement_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [ManagedPlacementViewService.GetManagedPlacementView][google.ads.googleads.v2.services.ManagedPlacementViewService.GetManagedPlacementView].
type GetManagedPlacementViewRequest struct {
	// The resource name of the Managed Placement View to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetManagedPlacementViewRequest) Reset()         { *m = GetManagedPlacementViewRequest{} }
func (m *GetManagedPlacementViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetManagedPlacementViewRequest) ProtoMessage()    {}
func (*GetManagedPlacementViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_48879207a9e4f5e2, []int{0}
}

func (m *GetManagedPlacementViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Unmarshal(m, b)
}
func (m *GetManagedPlacementViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Marshal(b, m, deterministic)
}
func (m *GetManagedPlacementViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetManagedPlacementViewRequest.Merge(m, src)
}
func (m *GetManagedPlacementViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetManagedPlacementViewRequest.Size(m)
}
func (m *GetManagedPlacementViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetManagedPlacementViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetManagedPlacementViewRequest proto.InternalMessageInfo

func (m *GetManagedPlacementViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetManagedPlacementViewRequest)(nil), "google.ads.googleads.v2.services.GetManagedPlacementViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/managed_placement_view_service.proto", fileDescriptor_48879207a9e4f5e2)
}

var fileDescriptor_48879207a9e4f5e2 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0x8b, 0xdb, 0x40,
	0x10, 0x45, 0x0a, 0x04, 0x22, 0x92, 0x46, 0x8d, 0x8d, 0x1c, 0x82, 0x70, 0x5c, 0x04, 0x17, 0xbb,
	0xa0, 0x14, 0x26, 0x1b, 0xf2, 0x21, 0x83, 0x71, 0x9a, 0x04, 0xe3, 0x80, 0x8a, 0x20, 0x10, 0x1b,
	0x69, 0x10, 0x02, 0x69, 0x57, 0xd1, 0xae, 0xe5, 0x22, 0xa4, 0x49, 0x93, 0x3e, 0xf9, 0x07, 0x29,
	0xef, 0xa7, 0xb8, 0x3d, 0xee, 0x1f, 0x5c, 0x75, 0xbf, 0xe2, 0x90, 0x57, 0x2b, 0x9f, 0xc1, 0xb2,
	0xbb, 0xc7, 0xce, 0x9b, 0xf7, 0x66, 0xde, 0xac, 0xb5, 0x48, 0x39, 0x4f, 0x73, 0xc0, 0x34, 0x11,
	0x58, 0xc1, 0x06, 0xd5, 0x1e, 0x16, 0x50, 0xd5, 0x59, 0x0c, 0x02, 0x17, 0x94, 0xd1, 0x14, 0x92,
	0xa8, 0xcc, 0x69, 0x0c, 0x05, 0x30, 0x19, 0xd5, 0x19, 0x6c, 0xa3, 0xb6, 0x8e, 0xca, 0x8a, 0x4b,
	0x6e, 0xbb, 0xaa, 0x17, 0xd1, 0x44, 0xa0, 0x4e, 0x06, 0xd5, 0x1e, 0xd2, 0x32, 0xce, 0xfb, 0x3e,
	0xa3, 0x0a, 0x04, 0xdf, 0x54, 0xfd, 0x4e, 0xca, 0xc1, 0x79, 0xae, 0xfb, 0xcb, 0x0c, 0x53, 0xc6,
	0xb8, 0xa4, 0x32, 0xe3, 0x4c, 0xb4, 0xd5, 0xc1, 0x83, 0x6a, 0x9c, 0x67, 0xc0, 0xa4, 0x2a, 0x8c,
	0x17, 0xd6, 0x8b, 0x25, 0xc8, 0xcf, 0x4a, 0x79, 0xa5, 0x85, 0x83, 0x0c, 0xb6, 0x6b, 0xf8, 0xb1,
	0x01, 0x21, 0xed, 0x97, 0xd6, 0x33, 0x3d, 0x42, 0xc4, 0x68, 0x01, 0x43, 0xc3, 0x35, 0x5e, 0x3d,
	0x59, 0x3f, 0xd5, 0x8f, 0x5f, 0x68, 0x01, 0xde, 0x5f, 0xd3, 0x1a, 0x9d, 0x12, 0xf9, 0xaa, 0xd6,
	0xb3, 0x6f, 0x0c, 0x6b, 0xd0, 0xe3, 0x63, 0x7f, 0x44, 0x97, 0xc2, 0x41, 0xe7, 0x47, 0x74, 0x66,
	0xbd, 0x0a, 0x5d, 0x78, 0xe8, 0x54, 0xff, 0xf8, 0xc3, 0xef, 0xeb, 0xdb, 0x7f, 0xe6, 0x1b, 0x7b,
	0xd6, 0x04, 0xfd, 0xf3, 0x68, 0xcd, 0x77, 0xf1, 0x46, 0x48, 0x5e, 0x40, 0x25, 0xf0, 0x54, 0x27,
	0x7f, 0xd4, 0x2c, 0xf0, 0xf4, 0x97, 0x33, 0xda, 0xf9, 0xc3, 0x83, 0x61, 0x8b, 0xca, 0x4c, 0xa0,
	0x98, 0x17, 0xf3, 0x3f, 0xa6, 0x35, 0x89, 0x79, 0x71, 0x71, 0xbd, 0xb9, 0x7b, 0x26, 0xba, 0x55,
	0x73, 0xa6, 0x95, 0xf1, 0xed, 0x53, 0xab, 0x92, 0xf2, 0x9c, 0xb2, 0x14, 0xf1, 0x2a, 0xc5, 0x29,
	0xb0, 0xfd, 0x11, 0xf1, 0xc1, 0xb7, 0xff, 0x9f, 0xbe, 0xd5, 0xe0, 0xbf, 0xf9, 0x68, 0xe9, 0xfb,
	0x57, 0xa6, 0xbb, 0x54, 0x82, 0x7e, 0x22, 0x90, 0x82, 0x0d, 0x0a, 0x3c, 0xd4, 0x1a, 0x8b, 0x9d,
	0xa6, 0x84, 0x7e, 0x22, 0xc2, 0x8e, 0x12, 0x06, 0x5e, 0xa8, 0x29, 0x77, 0xe6, 0x44, 0xbd, 0x13,
	0xe2, 0x27, 0x82, 0x90, 0x8e, 0x44, 0x48, 0xe0, 0x11, 0xa2, 0x69, 0xdf, 0x1f, 0xef, 0xe7, 0x7c,
	0x7d, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xba, 0x51, 0x3c, 0x4e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ManagedPlacementViewServiceClient is the client API for ManagedPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagedPlacementViewServiceClient interface {
	// Returns the requested Managed Placement view in full detail.
	GetManagedPlacementView(ctx context.Context, in *GetManagedPlacementViewRequest, opts ...grpc.CallOption) (*resources.ManagedPlacementView, error)
}

type managedPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagedPlacementViewServiceClient(cc grpc.ClientConnInterface) ManagedPlacementViewServiceClient {
	return &managedPlacementViewServiceClient{cc}
}

func (c *managedPlacementViewServiceClient) GetManagedPlacementView(ctx context.Context, in *GetManagedPlacementViewRequest, opts ...grpc.CallOption) (*resources.ManagedPlacementView, error) {
	out := new(resources.ManagedPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.ManagedPlacementViewService/GetManagedPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagedPlacementViewServiceServer is the server API for ManagedPlacementViewService service.
type ManagedPlacementViewServiceServer interface {
	// Returns the requested Managed Placement view in full detail.
	GetManagedPlacementView(context.Context, *GetManagedPlacementViewRequest) (*resources.ManagedPlacementView, error)
}

// UnimplementedManagedPlacementViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedManagedPlacementViewServiceServer struct {
}

func (*UnimplementedManagedPlacementViewServiceServer) GetManagedPlacementView(ctx context.Context, req *GetManagedPlacementViewRequest) (*resources.ManagedPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManagedPlacementView not implemented")
}

func RegisterManagedPlacementViewServiceServer(s *grpc.Server, srv ManagedPlacementViewServiceServer) {
	s.RegisterService(&_ManagedPlacementViewService_serviceDesc, srv)
}

func _ManagedPlacementViewService_GetManagedPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManagedPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagedPlacementViewServiceServer).GetManagedPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.ManagedPlacementViewService/GetManagedPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagedPlacementViewServiceServer).GetManagedPlacementView(ctx, req.(*GetManagedPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ManagedPlacementViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.ManagedPlacementViewService",
	HandlerType: (*ManagedPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManagedPlacementView",
			Handler:    _ManagedPlacementViewService_GetManagedPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/managed_placement_view_service.proto",
}

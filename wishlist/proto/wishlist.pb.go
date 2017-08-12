// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/wishlist.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/wishlist.proto

It has these top-level messages:
	GetRequest
	GetResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRequest) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

type GetResponse struct {
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto1.RegisterType((*GetRequest)(nil), "proto.GetRequest")
	proto1.RegisterType((*GetResponse)(nil), "proto.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WishlistService service

type WishlistServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type wishlistServiceClient struct {
	cc *grpc.ClientConn
}

func NewWishlistServiceClient(cc *grpc.ClientConn) WishlistServiceClient {
	return &wishlistServiceClient{cc}
}

func (c *wishlistServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/proto.WishlistService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WishlistService service

type WishlistServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterWishlistServiceServer(s *grpc.Server, srv WishlistServiceServer) {
	s.RegisterService(&_WishlistService_serviceDesc, srv)
}

func _WishlistService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WishlistService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WishlistService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WishlistService",
	HandlerType: (*WishlistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _WishlistService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/wishlist.proto",
}

func init() { proto1.RegisterFile("proto/wishlist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xcf, 0x2c, 0xce, 0xc8, 0xc9, 0x2c, 0x2e, 0xd1, 0x03, 0x73, 0x85, 0x58, 0xc1,
	0x94, 0x92, 0x2e, 0x17, 0x97, 0x7b, 0x6a, 0x49, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90,
	0x3c, 0x17, 0x77, 0x72, 0x69, 0x71, 0x49, 0x7e, 0x6e, 0x6a, 0x51, 0x7c, 0x66, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x73, 0x10, 0x17, 0x4c, 0xc8, 0x33, 0x45, 0x89, 0x97, 0x8b, 0x1b, 0xac, 0xbc,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0xd5, 0xc8, 0x9e, 0x8b, 0x3f, 0x1c, 0x6a, 0x6c, 0x70, 0x6a, 0x51,
	0x59, 0x66, 0x72, 0xaa, 0x90, 0x0e, 0x17, 0xb3, 0x7b, 0x6a, 0x89, 0x90, 0x20, 0xc4, 0x1a, 0x3d,
	0x84, 0xe1, 0x52, 0x42, 0xc8, 0x42, 0x10, 0x03, 0x9c, 0x0c, 0xa2, 0xf4, 0xd2, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xcb, 0x12, 0x8b, 0x33, 0x73, 0x32, 0x2b, 0x75, 0x4b,
	0xf4, 0x73, 0x33, 0x93, 0x8b, 0xf2, 0x4b, 0x8a, 0x12, 0x33, 0xf3, 0xe0, 0xee, 0xd6, 0x07, 0xeb,
	0x4e, 0x62, 0x03, 0x53, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0xba, 0xd0, 0x21, 0xd6,
	0x00, 0x00, 0x00,
}

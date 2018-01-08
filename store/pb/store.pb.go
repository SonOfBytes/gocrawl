// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store/pb/store.proto

/*
Package storepb is a generated protocol buffer package.

It is generated from these files:
	store/pb/store.proto

It has these top-level messages:
	StoreSubmitRequest
	StoreSubmitReply
	StoreGetRequest
	StoreGetReply
*/
package storepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Store url crawl results with session credentials
type StoreSubmitRequest struct {
	Session string   `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	Url     string   `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Urls    []string `protobuf:"bytes,3,rep,name=urls" json:"urls,omitempty"`
}

func (m *StoreSubmitRequest) Reset()                    { *m = StoreSubmitRequest{} }
func (m *StoreSubmitRequest) String() string            { return proto.CompactTextString(m) }
func (*StoreSubmitRequest) ProtoMessage()               {}
func (*StoreSubmitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StoreSubmitRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *StoreSubmitRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *StoreSubmitRequest) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

type StoreSubmitReply struct {
}

func (m *StoreSubmitReply) Reset()                    { *m = StoreSubmitReply{} }
func (m *StoreSubmitReply) String() string            { return proto.CompactTextString(m) }
func (*StoreSubmitReply) ProtoMessage()               {}
func (*StoreSubmitReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Get stored url crawl results with session credentials
type StoreGetRequest struct {
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *StoreGetRequest) Reset()                    { *m = StoreGetRequest{} }
func (m *StoreGetRequest) String() string            { return proto.CompactTextString(m) }
func (*StoreGetRequest) ProtoMessage()               {}
func (*StoreGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StoreGetRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *StoreGetRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type StoreGetReply struct {
	Urls []string `protobuf:"bytes,1,rep,name=urls" json:"urls,omitempty"`
}

func (m *StoreGetReply) Reset()                    { *m = StoreGetReply{} }
func (m *StoreGetReply) String() string            { return proto.CompactTextString(m) }
func (*StoreGetReply) ProtoMessage()               {}
func (*StoreGetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StoreGetReply) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func init() {
	proto.RegisterType((*StoreSubmitRequest)(nil), "StoreSubmitRequest")
	proto.RegisterType((*StoreSubmitReply)(nil), "StoreSubmitReply")
	proto.RegisterType((*StoreGetRequest)(nil), "StoreGetRequest")
	proto.RegisterType((*StoreGetReply)(nil), "StoreGetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Store service

type StoreClient interface {
	Submit(ctx context.Context, in *StoreSubmitRequest, opts ...grpc.CallOption) (*StoreSubmitReply, error)
	Get(ctx context.Context, in *StoreGetRequest, opts ...grpc.CallOption) (*StoreGetReply, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Submit(ctx context.Context, in *StoreSubmitRequest, opts ...grpc.CallOption) (*StoreSubmitReply, error) {
	out := new(StoreSubmitReply)
	err := grpc.Invoke(ctx, "/Store/Submit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Get(ctx context.Context, in *StoreGetRequest, opts ...grpc.CallOption) (*StoreGetReply, error) {
	out := new(StoreGetReply)
	err := grpc.Invoke(ctx, "/Store/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Store service

type StoreServer interface {
	Submit(context.Context, *StoreSubmitRequest) (*StoreSubmitReply, error)
	Get(context.Context, *StoreGetRequest) (*StoreGetReply, error)
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreSubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Store/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Submit(ctx, req.(*StoreSubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Store/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Get(ctx, req.(*StoreGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _Store_Submit_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Store_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store/pb/store.proto",
}

func init() { proto.RegisterFile("store/pb/store.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2e, 0xc9, 0x2f,
	0x4a, 0xd5, 0x2f, 0x48, 0xd2, 0x07, 0x33, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x95, 0x42, 0xb8,
	0x84, 0x82, 0x41, 0xdc, 0xe0, 0xd2, 0xa4, 0xdc, 0xcc, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x09, 0x2e, 0xf6, 0xe2, 0xd4, 0xe2, 0xe2, 0xcc, 0xfc, 0x3c, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x80, 0x8b, 0xb9, 0xb4, 0x28, 0x47, 0x82, 0x09, 0x2c, 0x0a,
	0x62, 0x0a, 0x09, 0x71, 0xb1, 0x94, 0x16, 0xe5, 0x14, 0x4b, 0x30, 0x2b, 0x30, 0x6b, 0x70, 0x06,
	0x81, 0xd9, 0x4a, 0x42, 0x5c, 0x02, 0x28, 0xa6, 0x16, 0xe4, 0x54, 0x2a, 0xd9, 0x72, 0xf1, 0x83,
	0xc5, 0xdc, 0x53, 0xc9, 0xb1, 0x46, 0x49, 0x99, 0x8b, 0x17, 0xa1, 0xbd, 0x20, 0xa7, 0x12, 0x6e,
	0x2f, 0x23, 0xc2, 0x5e, 0xa3, 0x34, 0x2e, 0x56, 0xb0, 0x22, 0x21, 0x23, 0x2e, 0x36, 0x88, 0xdd,
	0x42, 0xc2, 0x7a, 0x98, 0xfe, 0x93, 0x12, 0xd4, 0xc3, 0x70, 0x1e, 0x83, 0x90, 0x26, 0x17, 0xb3,
	0x7b, 0x6a, 0x89, 0x90, 0x80, 0x1e, 0x9a, 0x33, 0xa5, 0xf8, 0xf4, 0x50, 0x6c, 0x56, 0x62, 0x70,
	0xe2, 0x8c, 0x62, 0x07, 0x07, 0x62, 0x41, 0x52, 0x12, 0x1b, 0x38, 0x1c, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xf5, 0x5f, 0x42, 0xbd, 0x5f, 0x01, 0x00, 0x00,
}

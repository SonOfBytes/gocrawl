// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/pb/api.proto

/*
Package apipb is a generated protocol buffer package.

It is generated from these files:
	api/pb/api.proto

It has these top-level messages:
	APIAuthenticateRequest
	APIAuthenticateReply
	APIValidateRequest
	APIValidateReply
	APISubmitRequest
	APISubmitReply
	APIGetRequest
	APIGetReply
*/
package apipb

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

type APIAuthenticateRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *APIAuthenticateRequest) Reset()                    { *m = APIAuthenticateRequest{} }
func (m *APIAuthenticateRequest) String() string            { return proto.CompactTextString(m) }
func (*APIAuthenticateRequest) ProtoMessage()               {}
func (*APIAuthenticateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *APIAuthenticateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *APIAuthenticateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type APIAuthenticateReply struct {
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *APIAuthenticateReply) Reset()                    { *m = APIAuthenticateReply{} }
func (m *APIAuthenticateReply) String() string            { return proto.CompactTextString(m) }
func (*APIAuthenticateReply) ProtoMessage()               {}
func (*APIAuthenticateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *APIAuthenticateReply) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type APIValidateRequest struct {
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *APIValidateRequest) Reset()                    { *m = APIValidateRequest{} }
func (m *APIValidateRequest) String() string            { return proto.CompactTextString(m) }
func (*APIValidateRequest) ProtoMessage()               {}
func (*APIValidateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *APIValidateRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type APIValidateReply struct {
	Valid bool `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
}

func (m *APIValidateReply) Reset()                    { *m = APIValidateReply{} }
func (m *APIValidateReply) String() string            { return proto.CompactTextString(m) }
func (*APIValidateReply) ProtoMessage()               {}
func (*APIValidateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *APIValidateReply) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

// Request url crawl with session credentials
type APISubmitRequest struct {
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *APISubmitRequest) Reset()                    { *m = APISubmitRequest{} }
func (m *APISubmitRequest) String() string            { return proto.CompactTextString(m) }
func (*APISubmitRequest) ProtoMessage()               {}
func (*APISubmitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *APISubmitRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *APISubmitRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type APISubmitReply struct {
	Job string `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
}

func (m *APISubmitReply) Reset()                    { *m = APISubmitReply{} }
func (m *APISubmitReply) String() string            { return proto.CompactTextString(m) }
func (*APISubmitReply) ProtoMessage()               {}
func (*APISubmitReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *APISubmitReply) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

// Get url crawl results with session credentials
type APIGetRequest struct {
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *APIGetRequest) Reset()                    { *m = APIGetRequest{} }
func (m *APIGetRequest) String() string            { return proto.CompactTextString(m) }
func (*APIGetRequest) ProtoMessage()               {}
func (*APIGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *APIGetRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *APIGetRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type APIGetReply struct {
	Urls []string `protobuf:"bytes,1,rep,name=urls" json:"urls,omitempty"`
}

func (m *APIGetReply) Reset()                    { *m = APIGetReply{} }
func (m *APIGetReply) String() string            { return proto.CompactTextString(m) }
func (*APIGetReply) ProtoMessage()               {}
func (*APIGetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *APIGetReply) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func init() {
	proto.RegisterType((*APIAuthenticateRequest)(nil), "APIAuthenticateRequest")
	proto.RegisterType((*APIAuthenticateReply)(nil), "APIAuthenticateReply")
	proto.RegisterType((*APIValidateRequest)(nil), "APIValidateRequest")
	proto.RegisterType((*APIValidateReply)(nil), "APIValidateReply")
	proto.RegisterType((*APISubmitRequest)(nil), "APISubmitRequest")
	proto.RegisterType((*APISubmitReply)(nil), "APISubmitReply")
	proto.RegisterType((*APIGetRequest)(nil), "APIGetRequest")
	proto.RegisterType((*APIGetReply)(nil), "APIGetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	// Sends a api request to process an URL
	Submit(ctx context.Context, in *APISubmitRequest, opts ...grpc.CallOption) (*APISubmitReply, error)
	// Gets a response for an URL
	Get(ctx context.Context, in *APIGetRequest, opts ...grpc.CallOption) (*APIGetReply, error)
	// Sends a authentication request
	Authenticate(ctx context.Context, in *APIAuthenticateRequest, opts ...grpc.CallOption) (*APIAuthenticateReply, error)
	// Sends a validaton request
	Validate(ctx context.Context, in *APIValidateRequest, opts ...grpc.CallOption) (*APIValidateReply, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Submit(ctx context.Context, in *APISubmitRequest, opts ...grpc.CallOption) (*APISubmitReply, error) {
	out := new(APISubmitReply)
	err := grpc.Invoke(ctx, "/API/Submit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Get(ctx context.Context, in *APIGetRequest, opts ...grpc.CallOption) (*APIGetReply, error) {
	out := new(APIGetReply)
	err := grpc.Invoke(ctx, "/API/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Authenticate(ctx context.Context, in *APIAuthenticateRequest, opts ...grpc.CallOption) (*APIAuthenticateReply, error) {
	out := new(APIAuthenticateReply)
	err := grpc.Invoke(ctx, "/API/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Validate(ctx context.Context, in *APIValidateRequest, opts ...grpc.CallOption) (*APIValidateReply, error) {
	out := new(APIValidateReply)
	err := grpc.Invoke(ctx, "/API/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	// Sends a api request to process an URL
	Submit(context.Context, *APISubmitRequest) (*APISubmitReply, error)
	// Gets a response for an URL
	Get(context.Context, *APIGetRequest) (*APIGetReply, error)
	// Sends a authentication request
	Authenticate(context.Context, *APIAuthenticateRequest) (*APIAuthenticateReply, error)
	// Sends a validaton request
	Validate(context.Context, *APIValidateRequest) (*APIValidateReply, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APISubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Submit(ctx, req.(*APISubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Get(ctx, req.(*APIGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIAuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Authenticate(ctx, req.(*APIAuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Validate(ctx, req.(*APIValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _API_Submit_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _API_Get_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _API_Authenticate_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _API_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/pb/api.proto",
}

func init() { proto.RegisterFile("api/pb/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x4d, 0x7e, 0xf9, 0xf5, 0xdf, 0x58, 0x6b, 0x3a, 0x56, 0x0d, 0x39, 0xd5, 0x05, 0xa1, 0xa7,
	0xad, 0xa8, 0x37, 0x41, 0x8c, 0x97, 0x92, 0x5b, 0x88, 0xe0, 0xc1, 0xdb, 0xc6, 0x2e, 0xb8, 0x92,
	0x36, 0x6b, 0x76, 0xa3, 0xf8, 0x3d, 0xfd, 0x40, 0xb2, 0x69, 0x23, 0x49, 0x5b, 0x11, 0xbc, 0xcd,
	0xdb, 0x7d, 0xf3, 0xde, 0x0c, 0x6f, 0xc0, 0x65, 0x52, 0x4c, 0x65, 0x32, 0x65, 0x52, 0x50, 0x99,
	0x67, 0x3a, 0x23, 0x11, 0x1c, 0x07, 0x51, 0x18, 0x14, 0xfa, 0x99, 0x2f, 0xb5, 0x78, 0x62, 0x9a,
	0xc7, 0xfc, 0xb5, 0xe0, 0x4a, 0xa3, 0x0f, 0xdd, 0x42, 0xf1, 0x7c, 0xc9, 0x16, 0xdc, 0xb3, 0xc7,
	0xf6, 0xa4, 0x17, 0x7f, 0x63, 0xf3, 0x27, 0x99, 0x52, 0xef, 0x59, 0x3e, 0xf7, 0xfe, 0xad, 0xfe,
	0x2a, 0x4c, 0xce, 0x61, 0xb4, 0xa5, 0x28, 0xd3, 0x0f, 0xf4, 0xa0, 0xa3, 0xb8, 0x52, 0x22, 0x5b,
	0xae, 0xe5, 0x2a, 0x48, 0x28, 0x60, 0x10, 0x85, 0x0f, 0x2c, 0x15, 0xf3, 0x9a, 0xff, 0xcf, 0xfc,
	0x09, 0xb8, 0x0d, 0xbe, 0x51, 0x1f, 0x41, 0xeb, 0xcd, 0x3c, 0x94, 0xdc, 0x6e, 0xbc, 0x02, 0xe4,
	0xa6, 0x64, 0xde, 0x17, 0xc9, 0x42, 0xe8, 0x5f, 0x75, 0xd1, 0x05, 0xa7, 0xc8, 0xd3, 0xf5, 0x42,
	0xa6, 0x24, 0x04, 0x06, 0xb5, 0x7e, 0xe3, 0xe3, 0x82, 0xf3, 0x92, 0x25, 0xeb, 0x4e, 0x53, 0x92,
	0x6b, 0xd8, 0x0f, 0xa2, 0x70, 0xc6, 0xff, 0x64, 0x70, 0x0a, 0x7b, 0x55, 0xb3, 0x51, 0x47, 0xf8,
	0x5f, 0xe4, 0xa9, 0xf2, 0xec, 0xb1, 0x33, 0xe9, 0xc5, 0x65, 0x7d, 0xf1, 0x69, 0x83, 0x13, 0x44,
	0x21, 0x52, 0x68, 0xaf, 0x06, 0xc1, 0x21, 0xdd, 0x5c, 0xca, 0x3f, 0xa0, 0xcd, 0x39, 0x89, 0x85,
	0x67, 0xe0, 0xcc, 0xb8, 0xc6, 0x01, 0x6d, 0x4c, 0xe7, 0xf7, 0x69, 0xcd, 0x90, 0x58, 0x78, 0x0b,
	0xfd, 0x7a, 0x56, 0x78, 0x42, 0x77, 0xdf, 0x83, 0x7f, 0x44, 0x77, 0xc5, 0x4a, 0x2c, 0xbc, 0x82,
	0x6e, 0x95, 0x05, 0x1e, 0xd2, 0xed, 0x24, 0xfd, 0x21, 0xdd, 0x8c, 0x8b, 0x58, 0x77, 0x9d, 0xc7,
	0x16, 0x93, 0x42, 0x26, 0x49, 0xbb, 0x3c, 0xc4, 0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0x0b, 0x46, 0xd4, 0x9c, 0x02, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queue/pb/queue.proto

/*
Package queuepb is a generated protocol buffer package.

It is generated from these files:
	queue/pb/queue.proto

It has these top-level messages:
	QueueSubmitRequest
	QueueSubmitReply
	QueueGetRequest
	QueueGetReply
*/
package queuepb

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

// Queue url crawl with session credentials
type QueueSubmitRequest struct {
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Depth   int32  `protobuf:"varint,3,opt,name=depth" json:"depth,omitempty"`
	Job     string `protobuf:"bytes,4,opt,name=job" json:"job,omitempty"`
}

func (m *QueueSubmitRequest) Reset()                    { *m = QueueSubmitRequest{} }
func (m *QueueSubmitRequest) String() string            { return proto.CompactTextString(m) }
func (*QueueSubmitRequest) ProtoMessage()               {}
func (*QueueSubmitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueueSubmitRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *QueueSubmitRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *QueueSubmitRequest) GetDepth() int32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *QueueSubmitRequest) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

type QueueSubmitReply struct {
}

func (m *QueueSubmitReply) Reset()                    { *m = QueueSubmitReply{} }
func (m *QueueSubmitReply) String() string            { return proto.CompactTextString(m) }
func (*QueueSubmitReply) ProtoMessage()               {}
func (*QueueSubmitReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Queue url crawl with session credentials
type QueueGetRequest struct {
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *QueueGetRequest) Reset()                    { *m = QueueGetRequest{} }
func (m *QueueGetRequest) String() string            { return proto.CompactTextString(m) }
func (*QueueGetRequest) ProtoMessage()               {}
func (*QueueGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueueGetRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type QueueGetReply struct {
	Url   string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Depth int32  `protobuf:"varint,2,opt,name=depth" json:"depth,omitempty"`
	Job   string `protobuf:"bytes,3,opt,name=job" json:"job,omitempty"`
}

func (m *QueueGetReply) Reset()                    { *m = QueueGetReply{} }
func (m *QueueGetReply) String() string            { return proto.CompactTextString(m) }
func (*QueueGetReply) ProtoMessage()               {}
func (*QueueGetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueueGetReply) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *QueueGetReply) GetDepth() int32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *QueueGetReply) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func init() {
	proto.RegisterType((*QueueSubmitRequest)(nil), "QueueSubmitRequest")
	proto.RegisterType((*QueueSubmitReply)(nil), "QueueSubmitReply")
	proto.RegisterType((*QueueGetRequest)(nil), "QueueGetRequest")
	proto.RegisterType((*QueueGetReply)(nil), "QueueGetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Queue service

type QueueClient interface {
	// Sends a api request
	Submit(ctx context.Context, in *QueueSubmitRequest, opts ...grpc.CallOption) (*QueueSubmitReply, error)
	Get(ctx context.Context, in *QueueGetRequest, opts ...grpc.CallOption) (*QueueGetReply, error)
}

type queueClient struct {
	cc *grpc.ClientConn
}

func NewQueueClient(cc *grpc.ClientConn) QueueClient {
	return &queueClient{cc}
}

func (c *queueClient) Submit(ctx context.Context, in *QueueSubmitRequest, opts ...grpc.CallOption) (*QueueSubmitReply, error) {
	out := new(QueueSubmitReply)
	err := grpc.Invoke(ctx, "/Queue/Submit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) Get(ctx context.Context, in *QueueGetRequest, opts ...grpc.CallOption) (*QueueGetReply, error) {
	out := new(QueueGetReply)
	err := grpc.Invoke(ctx, "/Queue/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Queue service

type QueueServer interface {
	// Sends a api request
	Submit(context.Context, *QueueSubmitRequest) (*QueueSubmitReply, error)
	Get(context.Context, *QueueGetRequest) (*QueueGetReply, error)
}

func RegisterQueueServer(s *grpc.Server, srv QueueServer) {
	s.RegisterService(&_Queue_serviceDesc, srv)
}

func _Queue_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueSubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Queue/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Submit(ctx, req.(*QueueSubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Queue/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Get(ctx, req.(*QueueGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Queue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Queue",
	HandlerType: (*QueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _Queue_Submit_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Queue_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queue/pb/queue.proto",
}

func init() { proto.RegisterFile("queue/pb/queue.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2c, 0x4d, 0x2d,
	0x4d, 0xd5, 0x2f, 0x48, 0xd2, 0x07, 0x33, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x95, 0x32, 0xb8,
	0x84, 0x02, 0x41, 0xdc, 0xe0, 0xd2, 0xa4, 0xdc, 0xcc, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x09, 0x2e, 0xf6, 0xe2, 0xd4, 0xe2, 0xe2, 0xcc, 0xfc, 0x3c, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x80, 0x8b, 0xb9, 0xb4, 0x28, 0x47, 0x82, 0x09, 0x2c, 0x0a,
	0x62, 0x0a, 0x89, 0x70, 0xb1, 0xa6, 0xa4, 0x16, 0x94, 0x64, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0xb0,
	0x06, 0x41, 0x38, 0x20, 0x75, 0x59, 0xf9, 0x49, 0x12, 0x2c, 0x10, 0x75, 0x59, 0xf9, 0x49, 0x4a,
	0x42, 0x5c, 0x02, 0x28, 0x36, 0x15, 0xe4, 0x54, 0x2a, 0x69, 0x73, 0xf1, 0x83, 0xc5, 0xdc, 0x53,
	0x09, 0x5b, 0xad, 0xe4, 0xc9, 0xc5, 0x8b, 0x50, 0x5c, 0x90, 0x53, 0x09, 0x73, 0x0b, 0x23, 0x16,
	0xb7, 0x30, 0x61, 0x71, 0x0b, 0x33, 0xdc, 0x2d, 0x46, 0x69, 0x5c, 0xac, 0x60, 0xa3, 0x84, 0x8c,
	0xb8, 0xd8, 0x20, 0xee, 0x11, 0x12, 0xd6, 0xc3, 0x0c, 0x07, 0x29, 0x41, 0x3d, 0x0c, 0x27, 0x33,
	0x08, 0x69, 0x72, 0x31, 0xbb, 0xa7, 0x96, 0x08, 0x09, 0xe8, 0xa1, 0x39, 0x5d, 0x8a, 0x4f, 0x0f,
	0xc5, 0x7d, 0x4a, 0x0c, 0x4e, 0x9c, 0x51, 0xec, 0xe0, 0xc0, 0x2e, 0x48, 0x4a, 0x62, 0x03, 0x87,
	0xb7, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xdb, 0xc4, 0x4f, 0x87, 0x01, 0x00, 0x00,
}
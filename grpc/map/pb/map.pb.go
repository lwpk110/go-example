// Code generated by protoc-gen-go. DO NOT EDIT.
// source: map.proto

package pb

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

type HelloRequest struct {
	Conditions           map[string][]byte `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_map_b444b07772d47403, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (dst *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(dst, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetConditions() map[string][]byte {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_map_b444b07772d47403, []int{1}
}
func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (dst *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(dst, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterMapType((map[string][]byte)(nil), "HelloRequest.ConditionsEntry")
	proto.RegisterType((*HelloReply)(nil), "HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoClient is the client API for Demo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type demoClient struct {
	cc *grpc.ClientConn
}

func NewDemoClient(cc *grpc.ClientConn) DemoClient {
	return &demoClient{cc}
}

func (c *demoClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Demo/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServer is the server API for Demo service.
type DemoServer interface {
	Hello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterDemoServer(s *grpc.Server, srv DemoServer) {
	s.RegisterService(&_Demo_serviceDesc, srv)
}

func _Demo_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Demo/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Demo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Demo",
	HandlerType: (*DemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Demo_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "map.proto",
}

func init() { proto.RegisterFile("map.proto", fileDescriptor_map_b444b07772d47403) }

var fileDescriptor_map_b444b07772d47403 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x4d, 0x2c, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xea, 0x61, 0xe4, 0xe2, 0xf1, 0x48, 0xcd, 0xc9, 0xc9, 0x0f,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xb2, 0xe5, 0xe2, 0x4a, 0xce, 0xcf, 0x4b, 0xc9, 0x2c,
	0xc9, 0xcc, 0xcf, 0x2b, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd5, 0x43, 0x56, 0xa2,
	0xe7, 0x0c, 0x97, 0x77, 0xcd, 0x2b, 0x29, 0xaa, 0x0c, 0x42, 0xd2, 0x20, 0x65, 0xcb, 0xc5, 0x8f,
	0x26, 0x2d, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0x62, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0,
	0x04, 0x41, 0x38, 0x56, 0x4c, 0x16, 0x8c, 0x4a, 0x6a, 0x5c, 0x5c, 0x50, 0xab, 0x0a, 0x72, 0x2a,
	0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xa1, 0xba, 0x61, 0x5c, 0x23,
	0x5d, 0x2e, 0x16, 0x97, 0xd4, 0xdc, 0x7c, 0x21, 0x55, 0x2e, 0x56, 0xb0, 0x7a, 0x21, 0x5e, 0x14,
	0x27, 0x4a, 0x71, 0xeb, 0x21, 0x8c, 0x51, 0x62, 0x70, 0x62, 0x89, 0x62, 0x2a, 0x48, 0x4a, 0x62,
	0x03, 0x7b, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x27, 0xd8, 0xd8, 0x3a, 0xff, 0x00, 0x00,
	0x00,
}
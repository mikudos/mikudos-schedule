// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/users/users.proto

package users

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb78b70ef3ad6f73, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
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
	return fileDescriptor_bb78b70ef3ad6f73, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
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
	proto.RegisterType((*HelloRequest)(nil), "users.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "users.HelloReply")
}

func init() { proto.RegisterFile("proto/users/users.proto", fileDescriptor_bb78b70ef3ad6f73) }

var fileDescriptor_bb78b70ef3ad6f73 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2d, 0x4e, 0x2d, 0x2a, 0x86, 0x90, 0x7a, 0x60, 0x11, 0x21, 0x56, 0x30, 0x47,
	0x49, 0x89, 0x8b, 0xc7, 0x23, 0x35, 0x27, 0x27, 0x3f, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44,
	0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc,
	0x56, 0x52, 0xe3, 0xe2, 0x82, 0xaa, 0x29, 0xc8, 0xa9, 0x14, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x87, 0x29, 0x82, 0x71, 0x8d, 0xca, 0xb9, 0x78, 0x42, 0x41, 0x86, 0x06, 0xa7,
	0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x99, 0x70, 0x71, 0x04, 0x27, 0x56, 0x82, 0xb5, 0x0a, 0x09,
	0xeb, 0x41, 0x2c, 0x47, 0xb6, 0x4c, 0x4a, 0x10, 0x55, 0xb0, 0x20, 0xa7, 0x52, 0x89, 0x41, 0xc8,
	0x90, 0x8b, 0x15, 0xa4, 0x2b, 0x93, 0x78, 0x2d, 0x49, 0x6c, 0x60, 0x2f, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xbc, 0x3e, 0x5f, 0x5d, 0xed, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHi(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type usersServiceClient struct {
	cc *grpc.ClientConn
}

func NewUsersServiceClient(cc *grpc.ClientConn) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/users.UsersService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) SayHi(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/users.UsersService/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
type UsersServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHi(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedUsersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (*UnimplementedUsersServiceServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedUsersServiceServer) SayHi(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}

func RegisterUsersServiceServer(s *grpc.Server, srv UsersServiceServer) {
	s.RegisterService(&_UsersService_serviceDesc, srv)
}

func _UsersService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UsersService/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).SayHi(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _UsersService_SayHello_Handler,
		},
		{
			MethodName: "SayHi",
			Handler:    _UsersService_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/users/users.proto",
}

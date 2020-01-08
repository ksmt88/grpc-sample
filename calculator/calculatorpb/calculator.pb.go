// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

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

type SumRequest struct {
	NumberA              int32    `protobuf:"varint,1,opt,name=numberA,proto3" json:"numberA,omitempty"`
	NumberB              int32    `protobuf:"varint,2,opt,name=numberB,proto3" json:"numberB,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetNumberA() int32 {
	if m != nil {
		return m.NumberA
	}
	return 0
}

func (m *SumRequest) GetNumberB() int32 {
	if m != nil {
		return m.NumberB
	}
	return 0
}

type SumResponse struct {
	Res                  int32    `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetRes() int32 {
	if m != nil {
		return m.Res
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x47, 0x30, 0x0b, 0x92, 0x90, 0x38, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x08, 0x11, 0x25, 0x07, 0x2e, 0xae, 0xe0, 0xd2, 0xdc, 0xa0, 0xd4,
	0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e, 0xf6, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x47,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x18, 0x17, 0x21, 0xe3, 0x24, 0xc1, 0x84, 0x2c, 0xe3,
	0xa4, 0x24, 0xcf, 0xc5, 0x0d, 0x36, 0xa1, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x80, 0x8b,
	0xb9, 0x28, 0xb5, 0x18, 0xaa, 0x1d, 0xc4, 0x34, 0x72, 0xe5, 0xe2, 0x74, 0x86, 0x5a, 0x98, 0x2a,
	0x64, 0xc1, 0xc5, 0x1c, 0x5c, 0x9a, 0x2b, 0x24, 0xa6, 0x87, 0xe4, 0x2a, 0x84, 0x03, 0xa4, 0xc4,
	0x31, 0xc4, 0x21, 0xc6, 0x2a, 0x31, 0x38, 0xf1, 0x45, 0xf1, 0x20, 0x7b, 0x2b, 0x89, 0x0d, 0xec,
	0x19, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0x6d, 0xfc, 0xcd, 0xf8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculateClient is the client API for Calculate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type calculateClient struct {
	cc *grpc.ClientConn
}

func NewCalculateClient(cc *grpc.ClientConn) CalculateClient {
	return &calculateClient{cc}
}

func (c *calculateClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.Calculate/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServer is the server API for Calculate service.
type CalculateServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

// UnimplementedCalculateServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateServer struct {
}

func (*UnimplementedCalculateServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterCalculateServer(s *grpc.Server, srv CalculateServer) {
	s.RegisterService(&_Calculate_serviceDesc, srv)
}

func _Calculate_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculate/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculate",
	HandlerType: (*CalculateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Calculate_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator/calculatorpb/calculator.proto",
}

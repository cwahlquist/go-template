// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package gotemplateService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type ServiceCommand struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Completed            bool     `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceCommand) Reset()         { *m = ServiceCommand{} }
func (m *ServiceCommand) String() string { return proto.CompactTextString(m) }
func (*ServiceCommand) ProtoMessage()    {}
func (*ServiceCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *ServiceCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceCommand.Unmarshal(m, b)
}
func (m *ServiceCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceCommand.Marshal(b, m, deterministic)
}
func (m *ServiceCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceCommand.Merge(m, src)
}
func (m *ServiceCommand) XXX_Size() int {
	return xxx_messageInfo_ServiceCommand.Size(m)
}
func (m *ServiceCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceCommand proto.InternalMessageInfo

func (m *ServiceCommand) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceCommand) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceCommand) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type Id struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceCommand)(nil), "gotemplateService.ServiceCommand")
	proto.RegisterType((*Id)(nil), "gotemplateService.Id")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0xcf, 0x2f, 0x49, 0xcd,
	0x2d, 0xc8, 0x49, 0x2c, 0x49, 0x0d, 0x86, 0x48, 0x28, 0x05, 0x71, 0xf1, 0x41, 0x99, 0xce, 0xf9,
	0xb9, 0xb9, 0x89, 0x79, 0x29, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60,
	0x11, 0x30, 0x5b, 0x48, 0x86, 0x8b, 0x33, 0x39, 0x3f, 0xb7, 0x20, 0x27, 0xb5, 0x24, 0x35, 0x45,
	0x82, 0x59, 0x81, 0x51, 0x83, 0x23, 0x08, 0x21, 0xa0, 0x24, 0xc2, 0xc5, 0xe4, 0x89, 0x61, 0x8e,
	0x51, 0x2a, 0x97, 0xa0, 0x3b, 0xba, 0xf5, 0x42, 0x01, 0x5c, 0x6c, 0x1e, 0xa9, 0x89, 0x39, 0x25,
	0x19, 0x42, 0x8a, 0x7a, 0x18, 0x8e, 0xd3, 0x43, 0x75, 0x99, 0x14, 0x61, 0x25, 0x4a, 0x0c, 0x49,
	0x6c, 0x60, 0xaf, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x35, 0x2b, 0x23, 0x35, 0xfb, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GotemplateServiceClient is the client API for GotemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GotemplateServiceClient interface {
	Health(ctx context.Context, in *ServiceCommand, opts ...grpc.CallOption) (*ServiceCommand, error)
}

type gotemplateServiceClient struct {
	cc *grpc.ClientConn
}

func NewGotemplateServiceClient(cc *grpc.ClientConn) GotemplateServiceClient {
	return &gotemplateServiceClient{cc}
}

func (c *gotemplateServiceClient) Health(ctx context.Context, in *ServiceCommand, opts ...grpc.CallOption) (*ServiceCommand, error) {
	out := new(ServiceCommand)
	err := c.cc.Invoke(ctx, "/gotemplateService.GotemplateService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GotemplateServiceServer is the server API for GotemplateService service.
type GotemplateServiceServer interface {
	Health(context.Context, *ServiceCommand) (*ServiceCommand, error)
}

func RegisterGotemplateServiceServer(s *grpc.Server, srv GotemplateServiceServer) {
	s.RegisterService(&_GotemplateService_serviceDesc, srv)
}

func _GotemplateService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GotemplateServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gotemplateService.GotemplateService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GotemplateServiceServer).Health(ctx, req.(*ServiceCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _GotemplateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gotemplateService.GotemplateService",
	HandlerType: (*GotemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _GotemplateService_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

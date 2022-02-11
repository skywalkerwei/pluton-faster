// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pay.proto

package pay

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

// 支付创建
type CreateRequest struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid"`
	Oid                  int64    `protobuf:"varint,2,opt,name=Oid,proto3" json:"Oid"`
	Amount               int64    `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0564d675d5c516e0, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CreateRequest) GetOid() int64 {
	if m != nil {
		return m.Oid
	}
	return 0
}

func (m *CreateRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type CreateResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0564d675d5c516e0, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 支付详情
type DetailRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailRequest) Reset()         { *m = DetailRequest{} }
func (m *DetailRequest) String() string { return proto.CompactTextString(m) }
func (*DetailRequest) ProtoMessage()    {}
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0564d675d5c516e0, []int{2}
}

func (m *DetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailRequest.Unmarshal(m, b)
}
func (m *DetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailRequest.Marshal(b, m, deterministic)
}
func (m *DetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailRequest.Merge(m, src)
}
func (m *DetailRequest) XXX_Size() int {
	return xxx_messageInfo_DetailRequest.Size(m)
}
func (m *DetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetailRequest proto.InternalMessageInfo

func (m *DetailRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DetailResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Uid                  int64    `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid"`
	Oid                  int64    `protobuf:"varint,3,opt,name=Oid,proto3" json:"Oid"`
	Amount               int64    `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount"`
	Source               int64    `protobuf:"varint,5,opt,name=Source,proto3" json:"Source"`
	Status               int64    `protobuf:"varint,6,opt,name=Status,proto3" json:"Status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailResponse) Reset()         { *m = DetailResponse{} }
func (m *DetailResponse) String() string { return proto.CompactTextString(m) }
func (*DetailResponse) ProtoMessage()    {}
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0564d675d5c516e0, []int{3}
}

func (m *DetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailResponse.Unmarshal(m, b)
}
func (m *DetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailResponse.Marshal(b, m, deterministic)
}
func (m *DetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailResponse.Merge(m, src)
}
func (m *DetailResponse) XXX_Size() int {
	return xxx_messageInfo_DetailResponse.Size(m)
}
func (m *DetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetailResponse proto.InternalMessageInfo

func (m *DetailResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DetailResponse) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *DetailResponse) GetOid() int64 {
	if m != nil {
		return m.Oid
	}
	return 0
}

func (m *DetailResponse) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *DetailResponse) GetSource() int64 {
	if m != nil {
		return m.Source
	}
	return 0
}

func (m *DetailResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

// 支付详情
type CallbackRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Uid                  int64    `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid"`
	Oid                  int64    `protobuf:"varint,3,opt,name=Oid,proto3" json:"Oid"`
	Amount               int64    `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount"`
	Source               int64    `protobuf:"varint,5,opt,name=Source,proto3" json:"Source"`
	Status               int64    `protobuf:"varint,6,opt,name=Status,proto3" json:"Status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallbackRequest) Reset()         { *m = CallbackRequest{} }
func (m *CallbackRequest) String() string { return proto.CompactTextString(m) }
func (*CallbackRequest) ProtoMessage()    {}
func (*CallbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0564d675d5c516e0, []int{4}
}

func (m *CallbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackRequest.Unmarshal(m, b)
}
func (m *CallbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackRequest.Marshal(b, m, deterministic)
}
func (m *CallbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackRequest.Merge(m, src)
}
func (m *CallbackRequest) XXX_Size() int {
	return xxx_messageInfo_CallbackRequest.Size(m)
}
func (m *CallbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackRequest proto.InternalMessageInfo

func (m *CallbackRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CallbackRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CallbackRequest) GetOid() int64 {
	if m != nil {
		return m.Oid
	}
	return 0
}

func (m *CallbackRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CallbackRequest) GetSource() int64 {
	if m != nil {
		return m.Source
	}
	return 0
}

func (m *CallbackRequest) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type CallbackResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallbackResponse) Reset()         { *m = CallbackResponse{} }
func (m *CallbackResponse) String() string { return proto.CompactTextString(m) }
func (*CallbackResponse) ProtoMessage()    {}
func (*CallbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0564d675d5c516e0, []int{5}
}

func (m *CallbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackResponse.Unmarshal(m, b)
}
func (m *CallbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackResponse.Marshal(b, m, deterministic)
}
func (m *CallbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackResponse.Merge(m, src)
}
func (m *CallbackResponse) XXX_Size() int {
	return xxx_messageInfo_CallbackResponse.Size(m)
}
func (m *CallbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateRequest)(nil), "payclient.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "payclient.CreateResponse")
	proto.RegisterType((*DetailRequest)(nil), "payclient.DetailRequest")
	proto.RegisterType((*DetailResponse)(nil), "payclient.DetailResponse")
	proto.RegisterType((*CallbackRequest)(nil), "payclient.CallbackRequest")
	proto.RegisterType((*CallbackResponse)(nil), "payclient.CallbackResponse")
}

func init() { proto.RegisterFile("pay.proto", fileDescriptor_0564d675d5c516e0) }

var fileDescriptor_0564d675d5c516e0 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0xe5, 0x98, 0x5a, 0xf0, 0x4b, 0x0d, 0x95, 0x07, 0x64, 0xc2, 0x40, 0x95, 0x89, 0x29,
	0x03, 0xcc, 0x1d, 0x20, 0x6c, 0x0c, 0x45, 0x45, 0x2c, 0x6c, 0x6e, 0xe2, 0xc1, 0x22, 0xc4, 0x21,
	0xf9, 0x33, 0xe4, 0x1d, 0xe0, 0xbd, 0x78, 0x2c, 0x94, 0x38, 0x2e, 0x75, 0x09, 0x2b, 0x9b, 0xff,
	0xb3, 0x7d, 0xfe, 0x74, 0x67, 0x38, 0xa9, 0x64, 0x97, 0x54, 0xb5, 0x41, 0xc3, 0xfb, 0x65, 0x56,
	0x68, 0x55, 0x62, 0xfc, 0x00, 0xf3, 0xb4, 0x56, 0x12, 0xd5, 0x46, 0xbd, 0xb7, 0xaa, 0x41, 0xbe,
	0x00, 0xfa, 0xac, 0x73, 0x41, 0x96, 0xe4, 0x8a, 0x6e, 0xfa, 0x65, 0xaf, 0xac, 0x75, 0x2e, 0x02,
	0xab, 0xac, 0x75, 0xce, 0xcf, 0x80, 0xdd, 0xbe, 0x99, 0xb6, 0x44, 0x41, 0x07, 0x71, 0x9c, 0xe2,
	0x25, 0x84, 0xce, 0xac, 0xa9, 0x4c, 0xd9, 0x28, 0x1e, 0x42, 0xb0, 0x33, 0x0b, 0x74, 0x1e, 0x5f,
	0xc2, 0xfc, 0x5e, 0xa1, 0xd4, 0x85, 0x7b, 0xee, 0xf0, 0xc0, 0x07, 0x81, 0xd0, 0x9d, 0x98, 0xf6,
	0x70, 0x84, 0xc1, 0x2f, 0x42, 0x3a, 0x45, 0x78, 0xb4, 0x4f, 0xd8, 0xeb, 0x4f, 0xa6, 0xad, 0x33,
	0x25, 0x66, 0x56, 0xb7, 0xd3, 0xa0, 0xa3, 0xc4, 0xb6, 0x11, 0x6c, 0xd4, 0x87, 0x29, 0xfe, 0x24,
	0x70, 0x9a, 0xca, 0xa2, 0xd8, 0xca, 0xec, 0xf5, 0x0f, 0xe4, 0x7f, 0xe5, 0xe1, 0xb0, 0xf8, 0xc1,
	0xb1, 0xf9, 0x5c, 0x7f, 0x11, 0xa0, 0x8f, 0xb2, 0xe3, 0x2b, 0x60, 0x36, 0x7d, 0x2e, 0x92, 0x5d,
	0xc1, 0x89, 0xd7, 0x6e, 0x74, 0x3e, 0xb1, 0x33, 0xc6, 0xbc, 0x02, 0x66, 0x83, 0xf7, 0xae, 0x7b,
	0x6d, 0x79, 0xd7, 0x0f, 0x5a, 0x4a, 0xe1, 0xd8, 0x91, 0xf1, 0x68, 0xff, 0x15, 0x3f, 0xbd, 0xe8,
	0x62, 0x72, 0xcf, 0x9a, 0xdc, 0xcd, 0x5e, 0x68, 0x25, 0xbb, 0x2d, 0x1b, 0xbe, 0xe9, 0xcd, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x91, 0x8e, 0x0e, 0xb3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PayClient is the client API for Pay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PayClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
	Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error)
}

type payClient struct {
	cc *grpc.ClientConn
}

func NewPayClient(cc *grpc.ClientConn) PayClient {
	return &payClient{cc}
}

func (c *payClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/payclient.Pay/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	out := new(DetailResponse)
	err := c.cc.Invoke(ctx, "/payclient.Pay/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error) {
	out := new(CallbackResponse)
	err := c.cc.Invoke(ctx, "/payclient.Pay/Callback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayServer is the server API for Pay service.
type PayServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Detail(context.Context, *DetailRequest) (*DetailResponse, error)
	Callback(context.Context, *CallbackRequest) (*CallbackResponse, error)
}

// UnimplementedPayServer can be embedded to have forward compatible implementations.
type UnimplementedPayServer struct {
}

func (*UnimplementedPayServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedPayServer) Detail(ctx context.Context, req *DetailRequest) (*DetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (*UnimplementedPayServer) Callback(ctx context.Context, req *CallbackRequest) (*CallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
}

func RegisterPayServer(s *grpc.Server, srv PayServer) {
	s.RegisterService(&_Pay_serviceDesc, srv)
}

func _Pay_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payclient.Pay/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payclient.Pay/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).Detail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payclient.Pay/Callback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).Callback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "payclient.Pay",
	HandlerType: (*PayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Pay_Create_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _Pay_Detail_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _Pay_Callback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay.proto",
}

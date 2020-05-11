// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package pb

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type StringValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringValue) Reset()         { *m = StringValue{} }
func (m *StringValue) String() string { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()    {}
func (*StringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *StringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringValue.Unmarshal(m, b)
}
func (m *StringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringValue.Marshal(b, m, deterministic)
}
func (m *StringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringValue.Merge(m, src)
}
func (m *StringValue) XXX_Size() int {
	return xxx_messageInfo_StringValue.Size(m)
}
func (m *StringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringValue proto.InternalMessageInfo

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Token struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ServerInfo struct {
	ServerHost           string   `protobuf:"bytes,1,opt,name=ServerHost,proto3" json:"ServerHost,omitempty"`
	LoginKey             string   `protobuf:"bytes,2,opt,name=LoginKey,proto3" json:"LoginKey,omitempty"`
	ConnectionType       string   `protobuf:"bytes,3,opt,name=ConnectionType,proto3" json:"ConnectionType,omitempty"`
	LastId               string   `protobuf:"bytes,4,opt,name=LastId,proto3" json:"LastId,omitempty"`
	TcpPort              int64    `protobuf:"varint,5,opt,name=TcpPort,proto3" json:"TcpPort,omitempty"`
	KcpPort              int64    `protobuf:"varint,6,opt,name=KcpPort,proto3" json:"KcpPort,omitempty"`
	UdpApiPort           int64    `protobuf:"varint,7,opt,name=UdpApiPort,proto3" json:"UdpApiPort,omitempty"`
	KcpApiPort           int64    `protobuf:"varint,8,opt,name=KcpApiPort,proto3" json:"KcpApiPort,omitempty"`
	TlsPort              int64    `protobuf:"varint,9,opt,name=TlsPort,proto3" json:"TlsPort,omitempty"`
	GrpcPort             int64    `protobuf:"varint,10,opt,name=GrpcPort,proto3" json:"GrpcPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetServerHost() string {
	if m != nil {
		return m.ServerHost
	}
	return ""
}

func (m *ServerInfo) GetLoginKey() string {
	if m != nil {
		return m.LoginKey
	}
	return ""
}

func (m *ServerInfo) GetConnectionType() string {
	if m != nil {
		return m.ConnectionType
	}
	return ""
}

func (m *ServerInfo) GetLastId() string {
	if m != nil {
		return m.LastId
	}
	return ""
}

func (m *ServerInfo) GetTcpPort() int64 {
	if m != nil {
		return m.TcpPort
	}
	return 0
}

func (m *ServerInfo) GetKcpPort() int64 {
	if m != nil {
		return m.KcpPort
	}
	return 0
}

func (m *ServerInfo) GetUdpApiPort() int64 {
	if m != nil {
		return m.UdpApiPort
	}
	return 0
}

func (m *ServerInfo) GetKcpApiPort() int64 {
	if m != nil {
		return m.KcpApiPort
	}
	return 0
}

func (m *ServerInfo) GetTlsPort() int64 {
	if m != nil {
		return m.TlsPort
	}
	return 0
}

func (m *ServerInfo) GetGrpcPort() int64 {
	if m != nil {
		return m.GrpcPort
	}
	return 0
}

type LoginResponse struct {
	Code                 string   `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	LoginStatus          bool     `protobuf:"varint,3,opt,name=LoginStatus,proto3" json:"LoginStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *LoginResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LoginResponse) GetLoginStatus() bool {
	if m != nil {
		return m.LoginStatus
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*StringValue)(nil), "pb.StringValue")
	proto.RegisterType((*Token)(nil), "pb.Token")
	proto.RegisterType((*ServerInfo)(nil), "pb.ServerInfo")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x8b, 0xd3, 0x40,
	0x10, 0xc0, 0x4d, 0xef, 0xd2, 0x3f, 0x13, 0xee, 0xc0, 0x45, 0x34, 0x1c, 0x28, 0x25, 0x82, 0xde,
	0xcb, 0xe5, 0x41, 0x7d, 0xf0, 0x49, 0xb0, 0x87, 0xd8, 0xd2, 0x3b, 0x2c, 0x49, 0xf4, 0x55, 0x36,
	0xe9, 0x18, 0x82, 0xed, 0xee, 0xb2, 0xbb, 0xad, 0xe4, 0xd9, 0x6f, 0xe4, 0x27, 0x94, 0x9d, 0x24,
	0x6d, 0x2a, 0xde, 0xdb, 0xfc, 0xe6, 0x37, 0x93, 0xec, 0xce, 0x0e, 0x5c, 0x18, 0xd4, 0xfb, 0xaa,
	0xc0, 0x58, 0x69, 0x69, 0x25, 0x1b, 0xa8, 0x3c, 0x1a, 0x81, 0xff, 0x69, 0xab, 0x6c, 0x1d, 0xbd,
	0x84, 0x20, 0xb5, 0xba, 0x12, 0xe5, 0x37, 0xbe, 0xd9, 0x21, 0x7b, 0x02, 0x3e, 0x05, 0xa1, 0x37,
	0xf5, 0xae, 0x27, 0x49, 0x03, 0xd1, 0x73, 0xf0, 0x33, 0xf9, 0x13, 0xc5, 0x03, 0xfa, 0xcf, 0x00,
	0x20, 0x45, 0xbd, 0x47, 0xbd, 0x10, 0x3f, 0x24, 0x7b, 0xd1, 0xd1, 0x5c, 0x1a, 0xdb, 0x56, 0xf6,
	0x32, 0xec, 0x0a, 0xc6, 0x77, 0xb2, 0xac, 0xc4, 0x12, 0xeb, 0x70, 0x40, 0xf6, 0xc0, 0xec, 0x15,
	0x5c, 0xde, 0x4a, 0x21, 0xb0, 0xb0, 0x95, 0x14, 0x59, 0xad, 0x30, 0x3c, 0xa3, 0x8a, 0x7f, 0xb2,
	0xec, 0x29, 0x0c, 0xef, 0xb8, 0xb1, 0x8b, 0x75, 0x78, 0x4e, 0xbe, 0x25, 0x16, 0xc2, 0x28, 0x2b,
	0xd4, 0x4a, 0x6a, 0x1b, 0xfa, 0x53, 0xef, 0xfa, 0x2c, 0xe9, 0xd0, 0x99, 0x65, 0x6b, 0x86, 0x8d,
	0x69, 0xd1, 0x9d, 0xf7, 0xeb, 0x5a, 0x7d, 0x54, 0x15, 0xc9, 0x11, 0xc9, 0x5e, 0xc6, 0xf9, 0x65,
	0x71, 0xf0, 0xe3, 0xc6, 0x1f, 0x33, 0xf4, 0xcf, 0x8d, 0x21, 0x39, 0x69, 0xff, 0xd9, 0xa0, 0xbb,
	0xe9, 0x67, 0xad, 0x0a, 0x52, 0x40, 0xea, 0xc0, 0xd1, 0x77, 0xb8, 0xa0, 0x5b, 0x27, 0x68, 0x94,
	0x14, 0x06, 0x19, 0x83, 0xf3, 0x5b, 0xb9, 0xee, 0x46, 0x4b, 0xb1, 0xfb, 0xf4, 0x3d, 0x1a, 0xc3,
	0x4b, 0x6c, 0x27, 0xd5, 0x21, 0x9b, 0x42, 0x40, 0xed, 0xa9, 0xe5, 0x76, 0x67, 0x68, 0x4a, 0xe3,
	0xa4, 0x9f, 0x7a, 0xf3, 0xdb, 0x83, 0x60, 0x6e, 0xad, 0xba, 0xe7, 0x82, 0x97, 0xa8, 0xd9, 0x07,
	0x78, 0xd6, 0x68, 0x7a, 0x89, 0x59, 0xdd, 0x7b, 0xb1, 0xcb, 0x58, 0xe5, 0xf1, 0x91, 0xaf, 0x1e,
	0x3b, 0x3e, 0x39, 0x5d, 0xf4, 0x88, 0xbd, 0x03, 0x76, 0xd2, 0xdf, 0x6c, 0xc4, 0xc4, 0x95, 0x52,
	0xf8, 0xdf, 0xae, 0xd9, 0x7b, 0x78, 0x5d, 0xc8, 0x6d, 0x5c, 0x49, 0xeb, 0x96, 0x30, 0xfe, 0xa2,
	0x50, 0x2c, 0x64, 0x36, 0xdf, 0xe5, 0x71, 0xc9, 0x2d, 0xfe, 0xe2, 0xf5, 0x4d, 0xa9, 0x55, 0x71,
	0xc3, 0x55, 0x35, 0x0b, 0xdc, 0x6c, 0xd2, 0x66, 0x55, 0x57, 0x5e, 0x3e, 0xa4, 0x6d, 0x7d, 0xfb,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x17, 0x71, 0x3a, 0xbe, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HttpManagerClient is the client API for HttpManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HttpManagerClient interface {
	LoginServerByServerInfo(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*LoginResponse, error)
	LoginServerByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*LoginResponse, error)
}

type httpManagerClient struct {
	cc *grpc.ClientConn
}

func NewHttpManagerClient(cc *grpc.ClientConn) HttpManagerClient {
	return &httpManagerClient{cc}
}

func (c *httpManagerClient) LoginServerByServerInfo(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.HttpManager/LoginServerByServerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *httpManagerClient) LoginServerByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.HttpManager/LoginServerByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HttpManagerServer is the server API for HttpManager service.
type HttpManagerServer interface {
	LoginServerByServerInfo(context.Context, *ServerInfo) (*LoginResponse, error)
	LoginServerByToken(context.Context, *Token) (*LoginResponse, error)
}

// UnimplementedHttpManagerServer can be embedded to have forward compatible implementations.
type UnimplementedHttpManagerServer struct {
}

func (*UnimplementedHttpManagerServer) LoginServerByServerInfo(ctx context.Context, req *ServerInfo) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginServerByServerInfo not implemented")
}
func (*UnimplementedHttpManagerServer) LoginServerByToken(ctx context.Context, req *Token) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginServerByToken not implemented")
}

func RegisterHttpManagerServer(s *grpc.Server, srv HttpManagerServer) {
	s.RegisterService(&_HttpManager_serviceDesc, srv)
}

func _HttpManager_LoginServerByServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpManagerServer).LoginServerByServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HttpManager/LoginServerByServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpManagerServer).LoginServerByServerInfo(ctx, req.(*ServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _HttpManager_LoginServerByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpManagerServer).LoginServerByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HttpManager/LoginServerByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpManagerServer).LoginServerByToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _HttpManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HttpManager",
	HandlerType: (*HttpManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginServerByServerInfo",
			Handler:    _HttpManager_LoginServerByServerInfo_Handler,
		},
		{
			MethodName: "LoginServerByToken",
			Handler:    _HttpManager_LoginServerByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

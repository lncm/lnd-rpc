// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v0.8.0/watchtowerrpc/watchtower.proto

package watchtowerrpc

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

type GetInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_072fafef223e33aa, []int{0}
}

func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoRequest.Unmarshal(m, b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(m, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoRequest.Size(m)
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

type GetInfoResponse struct {
	/// The public key of the watchtower.
	Pubkey []byte `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	/// The listening addresses of the watchtower.
	Listeners []string `protobuf:"bytes,2,rep,name=listeners,proto3" json:"listeners,omitempty"`
	/// The URIs of the watchtower.
	Uris                 []string `protobuf:"bytes,3,rep,name=uris,proto3" json:"uris,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoResponse) Reset()         { *m = GetInfoResponse{} }
func (m *GetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetInfoResponse) ProtoMessage()    {}
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_072fafef223e33aa, []int{1}
}

func (m *GetInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoResponse.Unmarshal(m, b)
}
func (m *GetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResponse.Merge(m, src)
}
func (m *GetInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetInfoResponse.Size(m)
}
func (m *GetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResponse proto.InternalMessageInfo

func (m *GetInfoResponse) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *GetInfoResponse) GetListeners() []string {
	if m != nil {
		return m.Listeners
	}
	return nil
}

func (m *GetInfoResponse) GetUris() []string {
	if m != nil {
		return m.Uris
	}
	return nil
}

func init() {
	proto.RegisterType((*GetInfoRequest)(nil), "watchtowerrpc.GetInfoRequest")
	proto.RegisterType((*GetInfoResponse)(nil), "watchtowerrpc.GetInfoResponse")
}

func init() {
	proto.RegisterFile("v0.8.0/watchtowerrpc/watchtower.proto", fileDescriptor_072fafef223e33aa)
}

var fileDescriptor_072fafef223e33aa = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x33, 0xd0, 0xb3,
	0xd0, 0x33, 0xd0, 0x2f, 0x4f, 0x2c, 0x49, 0xce, 0x28, 0xc9, 0x2f, 0x4f, 0x2d, 0x2a, 0x2a, 0x48,
	0x46, 0xe2, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa2, 0xc8, 0x2b, 0x09, 0x70, 0xf1,
	0xb9, 0xa7, 0x96, 0x78, 0xe6, 0xa5, 0xe5, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0x45,
	0x73, 0xf1, 0xc3, 0x45, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x0a, 0x4a,
	0x93, 0xb2, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xa0, 0x3c, 0x21, 0x19, 0x2e,
	0xce, 0x9c, 0xcc, 0xe2, 0x92, 0xd4, 0xbc, 0xd4, 0xa2, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0xce,
	0x20, 0x84, 0x80, 0x90, 0x10, 0x17, 0x4b, 0x69, 0x51, 0x66, 0xb1, 0x04, 0x33, 0x58, 0x02, 0xcc,
	0x36, 0x0a, 0xe3, 0xe2, 0x0a, 0x87, 0xdb, 0x2f, 0xe4, 0xc1, 0xc5, 0x0e, 0xb5, 0x4a, 0x48, 0x56,
	0x0f, 0xc5, 0x5d, 0x7a, 0xa8, 0x8e, 0x92, 0x92, 0xc3, 0x25, 0x0d, 0x71, 0xa1, 0x93, 0x69, 0x94,
	0x71, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x4e, 0x66, 0x7a, 0x46,
	0x49, 0x5e, 0x66, 0x5e, 0x7a, 0x5e, 0x6a, 0x49, 0x79, 0x7e, 0x51, 0xb6, 0x7e, 0x4e, 0x5e, 0x8a,
	0x7e, 0x4e, 0x1e, 0x6a, 0x78, 0x14, 0x15, 0x24, 0x27, 0xb1, 0x81, 0xc3, 0xc4, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x11, 0x6f, 0x08, 0xca, 0x3c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WatchtowerClient is the client API for Watchtower service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WatchtowerClient interface {
	//* lncli: tower info
	//GetInfo returns general information concerning the companion watchtower
	//including it's public key and URIs where the server is currently
	//listening for clients.
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
}

type watchtowerClient struct {
	cc *grpc.ClientConn
}

func NewWatchtowerClient(cc *grpc.ClientConn) WatchtowerClient {
	return &watchtowerClient{cc}
}

func (c *watchtowerClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	out := new(GetInfoResponse)
	err := c.cc.Invoke(ctx, "/watchtowerrpc.Watchtower/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatchtowerServer is the server API for Watchtower service.
type WatchtowerServer interface {
	//* lncli: tower info
	//GetInfo returns general information concerning the companion watchtower
	//including it's public key and URIs where the server is currently
	//listening for clients.
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)
}

// UnimplementedWatchtowerServer can be embedded to have forward compatible implementations.
type UnimplementedWatchtowerServer struct {
}

func (*UnimplementedWatchtowerServer) GetInfo(ctx context.Context, req *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}

func RegisterWatchtowerServer(s *grpc.Server, srv WatchtowerServer) {
	s.RegisterService(&_Watchtower_serviceDesc, srv)
}

func _Watchtower_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatchtowerServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/watchtowerrpc.Watchtower/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatchtowerServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Watchtower_serviceDesc = grpc.ServiceDesc{
	ServiceName: "watchtowerrpc.Watchtower",
	HandlerType: (*WatchtowerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Watchtower_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v0.8.0/watchtowerrpc/watchtower.proto",
}

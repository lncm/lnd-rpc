// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v0.6.1/routerrpc/router.proto

package routerrpc

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

type PaymentRequest struct {
	//*
	//A serialized BOLT-11 payment request that contains all information
	//required to dispatch the payment. If the pay req is invalid, or expired,
	//an error will be returned.
	PayReq string `protobuf:"bytes,1,opt,name=pay_req,json=payReq,proto3" json:"pay_req,omitempty"`
	//*
	//An absolute limit on the highest fee we should pay when looking for a route
	//to the destination. Routes with fees higher than this will be ignored, if
	//there are no routes with a fee below this amount, an error will be
	//returned.
	FeeLimitSat int64 `protobuf:"varint,2,opt,name=fee_limit_sat,json=feeLimitSat,proto3" json:"fee_limit_sat,omitempty"`
	//*
	//An absolute limit on the cumulative CLTV value along the route for this
	//payment. Routes with total CLTV values higher than this will be ignored,
	//if there are no routes with a CLTV value below this amount, an error will
	//be returned.
	CltvLimit int32 `protobuf:"varint,3,opt,name=cltv_limit,json=cltvLimit,proto3" json:"cltv_limit,omitempty"`
	//*
	//An upper limit on the amount of time we should spend when attempting to
	//fulfill the payment. This is expressed in seconds. If we cannot make a
	//successful payment within this time frame, an error will be returned.
	TimeoutSeconds int32 `protobuf:"varint,4,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty"`
	//*
	//The channel id of the channel that must be taken to the first hop. If zero,
	//any channel may be used.
	OutgoingChannelId    int64    `protobuf:"varint,5,opt,name=outgoing_channel_id,json=outgoingChannelId,proto3" json:"outgoing_channel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentRequest) Reset()         { *m = PaymentRequest{} }
func (m *PaymentRequest) String() string { return proto.CompactTextString(m) }
func (*PaymentRequest) ProtoMessage()    {}
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_883560724cbd9db1, []int{0}
}

func (m *PaymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentRequest.Unmarshal(m, b)
}
func (m *PaymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentRequest.Marshal(b, m, deterministic)
}
func (m *PaymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentRequest.Merge(m, src)
}
func (m *PaymentRequest) XXX_Size() int {
	return xxx_messageInfo_PaymentRequest.Size(m)
}
func (m *PaymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentRequest proto.InternalMessageInfo

func (m *PaymentRequest) GetPayReq() string {
	if m != nil {
		return m.PayReq
	}
	return ""
}

func (m *PaymentRequest) GetFeeLimitSat() int64 {
	if m != nil {
		return m.FeeLimitSat
	}
	return 0
}

func (m *PaymentRequest) GetCltvLimit() int32 {
	if m != nil {
		return m.CltvLimit
	}
	return 0
}

func (m *PaymentRequest) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *PaymentRequest) GetOutgoingChannelId() int64 {
	if m != nil {
		return m.OutgoingChannelId
	}
	return 0
}

type PaymentResponse struct {
	//*
	//The payment hash that we paid to. Provided so callers are able to map
	//responses (which may be streaming) back to their original requests.
	PayHash []byte `protobuf:"bytes,1,opt,name=pay_hash,json=payHash,proto3" json:"pay_hash,omitempty"`
	//*
	//The pre-image of the payment successfully completed.
	PreImage []byte `protobuf:"bytes,2,opt,name=pre_image,json=preImage,proto3" json:"pre_image,omitempty"`
	//*
	//If not an empty string, then a string representation of the payment error.
	PaymentErr           string   `protobuf:"bytes,3,opt,name=payment_err,json=paymentErr,proto3" json:"payment_err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentResponse) Reset()         { *m = PaymentResponse{} }
func (m *PaymentResponse) String() string { return proto.CompactTextString(m) }
func (*PaymentResponse) ProtoMessage()    {}
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_883560724cbd9db1, []int{1}
}

func (m *PaymentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentResponse.Unmarshal(m, b)
}
func (m *PaymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentResponse.Marshal(b, m, deterministic)
}
func (m *PaymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentResponse.Merge(m, src)
}
func (m *PaymentResponse) XXX_Size() int {
	return xxx_messageInfo_PaymentResponse.Size(m)
}
func (m *PaymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentResponse proto.InternalMessageInfo

func (m *PaymentResponse) GetPayHash() []byte {
	if m != nil {
		return m.PayHash
	}
	return nil
}

func (m *PaymentResponse) GetPreImage() []byte {
	if m != nil {
		return m.PreImage
	}
	return nil
}

func (m *PaymentResponse) GetPaymentErr() string {
	if m != nil {
		return m.PaymentErr
	}
	return ""
}

type RouteFeeRequest struct {
	//*
	//The destination once wishes to obtain a routing fee quote to.
	Dest []byte `protobuf:"bytes,1,opt,name=dest,proto3" json:"dest,omitempty"`
	//*
	//The amount one wishes to send to the target destination.
	AmtSat               int64    `protobuf:"varint,2,opt,name=amt_sat,json=amtSat,proto3" json:"amt_sat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteFeeRequest) Reset()         { *m = RouteFeeRequest{} }
func (m *RouteFeeRequest) String() string { return proto.CompactTextString(m) }
func (*RouteFeeRequest) ProtoMessage()    {}
func (*RouteFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_883560724cbd9db1, []int{2}
}

func (m *RouteFeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteFeeRequest.Unmarshal(m, b)
}
func (m *RouteFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteFeeRequest.Marshal(b, m, deterministic)
}
func (m *RouteFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteFeeRequest.Merge(m, src)
}
func (m *RouteFeeRequest) XXX_Size() int {
	return xxx_messageInfo_RouteFeeRequest.Size(m)
}
func (m *RouteFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteFeeRequest proto.InternalMessageInfo

func (m *RouteFeeRequest) GetDest() []byte {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *RouteFeeRequest) GetAmtSat() int64 {
	if m != nil {
		return m.AmtSat
	}
	return 0
}

type RouteFeeResponse struct {
	//*
	//A lower bound of the estimated fee to the target destination within the
	//network, expressed in milli-satoshis.
	RoutingFeeMsat int64 `protobuf:"varint,1,opt,name=routing_fee_msat,json=routingFeeMsat,proto3" json:"routing_fee_msat,omitempty"`
	//*
	//An estimate of the worst case time delay that can occur. Note that callers
	//will still need to factor in the final CLTV delta of the last hop into this
	//value.
	TimeLockDelay        int64    `protobuf:"varint,2,opt,name=time_lock_delay,json=timeLockDelay,proto3" json:"time_lock_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteFeeResponse) Reset()         { *m = RouteFeeResponse{} }
func (m *RouteFeeResponse) String() string { return proto.CompactTextString(m) }
func (*RouteFeeResponse) ProtoMessage()    {}
func (*RouteFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_883560724cbd9db1, []int{3}
}

func (m *RouteFeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteFeeResponse.Unmarshal(m, b)
}
func (m *RouteFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteFeeResponse.Marshal(b, m, deterministic)
}
func (m *RouteFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteFeeResponse.Merge(m, src)
}
func (m *RouteFeeResponse) XXX_Size() int {
	return xxx_messageInfo_RouteFeeResponse.Size(m)
}
func (m *RouteFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RouteFeeResponse proto.InternalMessageInfo

func (m *RouteFeeResponse) GetRoutingFeeMsat() int64 {
	if m != nil {
		return m.RoutingFeeMsat
	}
	return 0
}

func (m *RouteFeeResponse) GetTimeLockDelay() int64 {
	if m != nil {
		return m.TimeLockDelay
	}
	return 0
}

func init() {
	proto.RegisterType((*PaymentRequest)(nil), "routerrpc.PaymentRequest")
	proto.RegisterType((*PaymentResponse)(nil), "routerrpc.PaymentResponse")
	proto.RegisterType((*RouteFeeRequest)(nil), "routerrpc.RouteFeeRequest")
	proto.RegisterType((*RouteFeeResponse)(nil), "routerrpc.RouteFeeResponse")
}

func init() { proto.RegisterFile("v0.6.1/routerrpc/router.proto", fileDescriptor_883560724cbd9db1) }

var fileDescriptor_883560724cbd9db1 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x15, 0xb6, 0x85, 0xe5, 0x74, 0x6b, 0x8b, 0xb9, 0x20, 0xeb, 0x34, 0x51, 0xe5, 0x02,
	0x72, 0x15, 0xbe, 0x24, 0x2e, 0xb9, 0x61, 0x9b, 0x98, 0x18, 0x12, 0x72, 0x1f, 0xc0, 0x32, 0xc9,
	0x59, 0x13, 0x16, 0xc7, 0xae, 0xed, 0x4c, 0xca, 0xb3, 0xf0, 0x3c, 0xbc, 0x17, 0xb2, 0xe3, 0x96,
	0x81, 0x7a, 0xe7, 0xfc, 0xcf, 0xe7, 0xff, 0x77, 0x02, 0x17, 0x0f, 0x6f, 0x8b, 0x8f, 0xc5, 0xbb,
	0x37, 0x5a, 0xf6, 0x16, 0xb5, 0x56, 0x65, 0x78, 0x15, 0x4a, 0x4b, 0x2b, 0x49, 0xb2, 0xd3, 0xb3,
	0xdf, 0x11, 0x4c, 0xbf, 0xf3, 0x41, 0x60, 0x67, 0x29, 0x6e, 0x7a, 0x34, 0x96, 0xbc, 0x80, 0xa7,
	0x8a, 0x0f, 0x4c, 0xe3, 0x26, 0x8d, 0x96, 0x51, 0x9e, 0xd0, 0x58, 0xf1, 0x81, 0xe2, 0x86, 0x64,
	0x70, 0x7a, 0x87, 0xc8, 0xda, 0x46, 0x34, 0x96, 0x19, 0x6e, 0xd3, 0x27, 0xcb, 0x28, 0x3f, 0xa0,
	0x93, 0x3b, 0xc4, 0x5b, 0xa7, 0xad, 0xb8, 0x25, 0x17, 0x00, 0x65, 0x6b, 0x1f, 0xc6, 0xa4, 0xf4,
	0x60, 0x19, 0xe5, 0x47, 0x34, 0x71, 0x8a, 0xcf, 0x20, 0xaf, 0x61, 0x66, 0x1b, 0x81, 0xb2, 0xb7,
	0xcc, 0x60, 0x29, 0xbb, 0xca, 0xa4, 0x87, 0x3e, 0x67, 0x1a, 0xe4, 0xd5, 0xa8, 0x92, 0x02, 0x9e,
	0xcb, 0xde, 0xae, 0x65, 0xd3, 0xad, 0x59, 0x59, 0xf3, 0xae, 0xc3, 0x96, 0x35, 0x55, 0x7a, 0xe4,
	0x27, 0x3e, 0xdb, 0x86, 0x3e, 0x8f, 0x91, 0x9b, 0x2a, 0xfb, 0x09, 0xb3, 0x9d, 0x0d, 0xa3, 0x64,
	0x67, 0x90, 0x9c, 0xc1, 0xb1, 0xf3, 0x51, 0x73, 0x53, 0x7b, 0x23, 0x27, 0xd4, 0xf9, 0xfa, 0xc2,
	0x4d, 0x4d, 0xce, 0x21, 0x51, 0x1a, 0x59, 0x23, 0xf8, 0x1a, 0xbd, 0x8b, 0x13, 0x7a, 0xac, 0x34,
	0xde, 0xb8, 0x6f, 0xf2, 0x12, 0x26, 0x6a, 0x6c, 0xc5, 0x50, 0x6b, 0xef, 0x21, 0xa1, 0x10, 0xa4,
	0x2b, 0xad, 0xb3, 0x4f, 0x30, 0xa3, 0x0e, 0xe0, 0x35, 0xe2, 0x96, 0x19, 0x81, 0xc3, 0x0a, 0x8d,
	0x0d, 0x73, 0xfc, 0xdb, 0x71, 0xe4, 0xe2, 0x31, 0xa8, 0x98, 0x0b, 0xc7, 0x28, 0xab, 0x60, 0xfe,
	0xb7, 0x3e, 0x2c, 0x9b, 0xc3, 0xdc, 0x1d, 0xc5, 0xd9, 0x75, 0x8c, 0x85, 0xab, 0x8a, 0x7c, 0xd5,
	0x34, 0xe8, 0xd7, 0x88, 0xdf, 0x0c, 0xb7, 0xe4, 0xd5, 0x88, 0x90, 0xb5, 0xb2, 0xbc, 0x67, 0x15,
	0xb6, 0x7c, 0x08, 0xed, 0x4f, 0x9d, 0x7c, 0x2b, 0xcb, 0xfb, 0x4b, 0x27, 0xbe, 0xff, 0x15, 0x41,
	0xec, 0xc7, 0x68, 0x72, 0x09, 0x93, 0x15, 0x76, 0x55, 0x00, 0x44, 0xce, 0x8a, 0xdd, 0xfd, 0x8b,
	0x7f, 0x6f, 0xbf, 0x58, 0xec, 0x0b, 0x85, 0x15, 0xbf, 0xc2, 0xfc, 0xca, 0xd8, 0x46, 0x70, 0x8b,
	0xdb, 0xf5, 0xc9, 0xe3, 0xfc, 0xff, 0x98, 0x2c, 0xce, 0xf7, 0xc6, 0xc6, 0x66, 0x3f, 0x62, 0xff,
	0x27, 0x7e, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0x81, 0xfb, 0x6a, 0x42, 0xaa, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	//*
	//SendPayment attempts to route a payment described by the passed
	//PaymentRequest to the final destination. If we are unable to route the
	//payment, or cannot find a route that satisfies the constraints in the
	//PaymentRequest, then an error will be returned. Otherwise, the payment
	//pre-image, along with the final route will be returned.
	SendPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	//*
	//EstimateRouteFee allows callers to obtain a lower bound w.r.t how much it
	//may cost to send an HTLC to the target end destination.
	EstimateRouteFee(ctx context.Context, in *RouteFeeRequest, opts ...grpc.CallOption) (*RouteFeeResponse, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) SendPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/routerrpc.Router/SendPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) EstimateRouteFee(ctx context.Context, in *RouteFeeRequest, opts ...grpc.CallOption) (*RouteFeeResponse, error) {
	out := new(RouteFeeResponse)
	err := c.cc.Invoke(ctx, "/routerrpc.Router/EstimateRouteFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	//*
	//SendPayment attempts to route a payment described by the passed
	//PaymentRequest to the final destination. If we are unable to route the
	//payment, or cannot find a route that satisfies the constraints in the
	//PaymentRequest, then an error will be returned. Otherwise, the payment
	//pre-image, along with the final route will be returned.
	SendPayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
	//*
	//EstimateRouteFee allows callers to obtain a lower bound w.r.t how much it
	//may cost to send an HTLC to the target end destination.
	EstimateRouteFee(context.Context, *RouteFeeRequest) (*RouteFeeResponse, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) SendPayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPayment not implemented")
}
func (*UnimplementedRouterServer) EstimateRouteFee(ctx context.Context, req *RouteFeeRequest) (*RouteFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateRouteFee not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_SendPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).SendPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerrpc.Router/SendPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).SendPayment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_EstimateRouteFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).EstimateRouteFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerrpc.Router/EstimateRouteFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).EstimateRouteFee(ctx, req.(*RouteFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routerrpc.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPayment",
			Handler:    _Router_SendPayment_Handler,
		},
		{
			MethodName: "EstimateRouteFee",
			Handler:    _Router_EstimateRouteFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v0.6.1/routerrpc/router.proto",
}

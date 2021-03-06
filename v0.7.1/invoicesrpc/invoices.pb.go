// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v0.7.1/invoicesrpc/invoices.proto

package invoicesrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	lnrpc "github.com/lncm/lnd-rpc/v0.7.1/lnrpc"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CancelInvoiceMsg struct {
	/// Hash corresponding to the (hold) invoice to cancel.
	PaymentHash          []byte   `protobuf:"bytes,1,opt,name=payment_hash,json=paymentHash,proto3" json:"payment_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelInvoiceMsg) Reset()         { *m = CancelInvoiceMsg{} }
func (m *CancelInvoiceMsg) String() string { return proto.CompactTextString(m) }
func (*CancelInvoiceMsg) ProtoMessage()    {}
func (*CancelInvoiceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_812a5d18f4719e7a, []int{0}
}

func (m *CancelInvoiceMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelInvoiceMsg.Unmarshal(m, b)
}
func (m *CancelInvoiceMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelInvoiceMsg.Marshal(b, m, deterministic)
}
func (m *CancelInvoiceMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelInvoiceMsg.Merge(m, src)
}
func (m *CancelInvoiceMsg) XXX_Size() int {
	return xxx_messageInfo_CancelInvoiceMsg.Size(m)
}
func (m *CancelInvoiceMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelInvoiceMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CancelInvoiceMsg proto.InternalMessageInfo

func (m *CancelInvoiceMsg) GetPaymentHash() []byte {
	if m != nil {
		return m.PaymentHash
	}
	return nil
}

type CancelInvoiceResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelInvoiceResp) Reset()         { *m = CancelInvoiceResp{} }
func (m *CancelInvoiceResp) String() string { return proto.CompactTextString(m) }
func (*CancelInvoiceResp) ProtoMessage()    {}
func (*CancelInvoiceResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_812a5d18f4719e7a, []int{1}
}

func (m *CancelInvoiceResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelInvoiceResp.Unmarshal(m, b)
}
func (m *CancelInvoiceResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelInvoiceResp.Marshal(b, m, deterministic)
}
func (m *CancelInvoiceResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelInvoiceResp.Merge(m, src)
}
func (m *CancelInvoiceResp) XXX_Size() int {
	return xxx_messageInfo_CancelInvoiceResp.Size(m)
}
func (m *CancelInvoiceResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelInvoiceResp.DiscardUnknown(m)
}

var xxx_messageInfo_CancelInvoiceResp proto.InternalMessageInfo

type AddHoldInvoiceRequest struct {
	//*
	//An optional memo to attach along with the invoice. Used for record keeping
	//purposes for the invoice's creator, and will also be set in the description
	//field of the encoded payment request if the description_hash field is not
	//being used.
	Memo string `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
	/// The hash of the preimage
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	/// The value of this invoice in satoshis
	Value int64 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	//*
	//Hash (SHA-256) of a description of the payment. Used if the description of
	//payment (memo) is too long to naturally fit within the description field
	//of an encoded payment request.
	DescriptionHash []byte `protobuf:"bytes,4,opt,name=description_hash,proto3" json:"description_hash,omitempty"`
	/// Payment request expiry time in seconds. Default is 3600 (1 hour).
	Expiry int64 `protobuf:"varint,5,opt,name=expiry,proto3" json:"expiry,omitempty"`
	/// Fallback on-chain address.
	FallbackAddr string `protobuf:"bytes,6,opt,name=fallback_addr,proto3" json:"fallback_addr,omitempty"`
	/// Delta to use for the time-lock of the CLTV extended to the final hop.
	CltvExpiry uint64 `protobuf:"varint,7,opt,name=cltv_expiry,proto3" json:"cltv_expiry,omitempty"`
	//*
	//Route hints that can each be individually used to assist in reaching the
	//invoice's destination.
	RouteHints []*lnrpc.RouteHint `protobuf:"bytes,8,rep,name=route_hints,proto3" json:"route_hints,omitempty"`
	/// Whether this invoice should include routing hints for private channels.
	Private              bool     `protobuf:"varint,9,opt,name=private,proto3" json:"private,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddHoldInvoiceRequest) Reset()         { *m = AddHoldInvoiceRequest{} }
func (m *AddHoldInvoiceRequest) String() string { return proto.CompactTextString(m) }
func (*AddHoldInvoiceRequest) ProtoMessage()    {}
func (*AddHoldInvoiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_812a5d18f4719e7a, []int{2}
}

func (m *AddHoldInvoiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddHoldInvoiceRequest.Unmarshal(m, b)
}
func (m *AddHoldInvoiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddHoldInvoiceRequest.Marshal(b, m, deterministic)
}
func (m *AddHoldInvoiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddHoldInvoiceRequest.Merge(m, src)
}
func (m *AddHoldInvoiceRequest) XXX_Size() int {
	return xxx_messageInfo_AddHoldInvoiceRequest.Size(m)
}
func (m *AddHoldInvoiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddHoldInvoiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddHoldInvoiceRequest proto.InternalMessageInfo

func (m *AddHoldInvoiceRequest) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *AddHoldInvoiceRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *AddHoldInvoiceRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *AddHoldInvoiceRequest) GetDescriptionHash() []byte {
	if m != nil {
		return m.DescriptionHash
	}
	return nil
}

func (m *AddHoldInvoiceRequest) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *AddHoldInvoiceRequest) GetFallbackAddr() string {
	if m != nil {
		return m.FallbackAddr
	}
	return ""
}

func (m *AddHoldInvoiceRequest) GetCltvExpiry() uint64 {
	if m != nil {
		return m.CltvExpiry
	}
	return 0
}

func (m *AddHoldInvoiceRequest) GetRouteHints() []*lnrpc.RouteHint {
	if m != nil {
		return m.RouteHints
	}
	return nil
}

func (m *AddHoldInvoiceRequest) GetPrivate() bool {
	if m != nil {
		return m.Private
	}
	return false
}

type AddHoldInvoiceResp struct {
	//*
	//A bare-bones invoice for a payment within the Lightning Network.  With the
	//details of the invoice, the sender has all the data necessary to send a
	//payment to the recipient.
	PaymentRequest       string   `protobuf:"bytes,1,opt,name=payment_request,proto3" json:"payment_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddHoldInvoiceResp) Reset()         { *m = AddHoldInvoiceResp{} }
func (m *AddHoldInvoiceResp) String() string { return proto.CompactTextString(m) }
func (*AddHoldInvoiceResp) ProtoMessage()    {}
func (*AddHoldInvoiceResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_812a5d18f4719e7a, []int{3}
}

func (m *AddHoldInvoiceResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddHoldInvoiceResp.Unmarshal(m, b)
}
func (m *AddHoldInvoiceResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddHoldInvoiceResp.Marshal(b, m, deterministic)
}
func (m *AddHoldInvoiceResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddHoldInvoiceResp.Merge(m, src)
}
func (m *AddHoldInvoiceResp) XXX_Size() int {
	return xxx_messageInfo_AddHoldInvoiceResp.Size(m)
}
func (m *AddHoldInvoiceResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddHoldInvoiceResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddHoldInvoiceResp proto.InternalMessageInfo

func (m *AddHoldInvoiceResp) GetPaymentRequest() string {
	if m != nil {
		return m.PaymentRequest
	}
	return ""
}

type SettleInvoiceMsg struct {
	/// Externally discovered pre-image that should be used to settle the hold invoice.
	Preimage             []byte   `protobuf:"bytes,1,opt,name=preimage,proto3" json:"preimage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettleInvoiceMsg) Reset()         { *m = SettleInvoiceMsg{} }
func (m *SettleInvoiceMsg) String() string { return proto.CompactTextString(m) }
func (*SettleInvoiceMsg) ProtoMessage()    {}
func (*SettleInvoiceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_812a5d18f4719e7a, []int{4}
}

func (m *SettleInvoiceMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleInvoiceMsg.Unmarshal(m, b)
}
func (m *SettleInvoiceMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleInvoiceMsg.Marshal(b, m, deterministic)
}
func (m *SettleInvoiceMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleInvoiceMsg.Merge(m, src)
}
func (m *SettleInvoiceMsg) XXX_Size() int {
	return xxx_messageInfo_SettleInvoiceMsg.Size(m)
}
func (m *SettleInvoiceMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleInvoiceMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SettleInvoiceMsg proto.InternalMessageInfo

func (m *SettleInvoiceMsg) GetPreimage() []byte {
	if m != nil {
		return m.Preimage
	}
	return nil
}

type SettleInvoiceResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettleInvoiceResp) Reset()         { *m = SettleInvoiceResp{} }
func (m *SettleInvoiceResp) String() string { return proto.CompactTextString(m) }
func (*SettleInvoiceResp) ProtoMessage()    {}
func (*SettleInvoiceResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_812a5d18f4719e7a, []int{5}
}

func (m *SettleInvoiceResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettleInvoiceResp.Unmarshal(m, b)
}
func (m *SettleInvoiceResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettleInvoiceResp.Marshal(b, m, deterministic)
}
func (m *SettleInvoiceResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettleInvoiceResp.Merge(m, src)
}
func (m *SettleInvoiceResp) XXX_Size() int {
	return xxx_messageInfo_SettleInvoiceResp.Size(m)
}
func (m *SettleInvoiceResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SettleInvoiceResp.DiscardUnknown(m)
}

var xxx_messageInfo_SettleInvoiceResp proto.InternalMessageInfo

type SubscribeSingleInvoiceRequest struct {
	/// Hash corresponding to the (hold) invoice to subscribe to.
	RHash                []byte   `protobuf:"bytes,2,opt,name=r_hash,proto3" json:"r_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeSingleInvoiceRequest) Reset()         { *m = SubscribeSingleInvoiceRequest{} }
func (m *SubscribeSingleInvoiceRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeSingleInvoiceRequest) ProtoMessage()    {}
func (*SubscribeSingleInvoiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_812a5d18f4719e7a, []int{6}
}

func (m *SubscribeSingleInvoiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeSingleInvoiceRequest.Unmarshal(m, b)
}
func (m *SubscribeSingleInvoiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeSingleInvoiceRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeSingleInvoiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeSingleInvoiceRequest.Merge(m, src)
}
func (m *SubscribeSingleInvoiceRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeSingleInvoiceRequest.Size(m)
}
func (m *SubscribeSingleInvoiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeSingleInvoiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeSingleInvoiceRequest proto.InternalMessageInfo

func (m *SubscribeSingleInvoiceRequest) GetRHash() []byte {
	if m != nil {
		return m.RHash
	}
	return nil
}

func init() {
	proto.RegisterType((*CancelInvoiceMsg)(nil), "invoicesrpc.CancelInvoiceMsg")
	proto.RegisterType((*CancelInvoiceResp)(nil), "invoicesrpc.CancelInvoiceResp")
	proto.RegisterType((*AddHoldInvoiceRequest)(nil), "invoicesrpc.AddHoldInvoiceRequest")
	proto.RegisterType((*AddHoldInvoiceResp)(nil), "invoicesrpc.AddHoldInvoiceResp")
	proto.RegisterType((*SettleInvoiceMsg)(nil), "invoicesrpc.SettleInvoiceMsg")
	proto.RegisterType((*SettleInvoiceResp)(nil), "invoicesrpc.SettleInvoiceResp")
	proto.RegisterType((*SubscribeSingleInvoiceRequest)(nil), "invoicesrpc.SubscribeSingleInvoiceRequest")
}

func init() { proto.RegisterFile("v0.7.1/invoicesrpc/invoices.proto", fileDescriptor_812a5d18f4719e7a) }

var fileDescriptor_812a5d18f4719e7a = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xda, 0x6e, 0xb7, 0x3b, 0xdd, 0x8f, 0x62, 0xa0, 0x8a, 0x22, 0x16, 0xb2, 0x11, 0x87,
	0x68, 0x25, 0xd2, 0x52, 0x84, 0x38, 0x81, 0x04, 0x5c, 0x0a, 0x12, 0x1c, 0x52, 0xc1, 0x81, 0x4b,
	0xe5, 0x26, 0x26, 0xb5, 0x70, 0x6d, 0x63, 0xbb, 0x15, 0xfb, 0xab, 0xf8, 0x3d, 0xfc, 0x1b, 0x14,
	0xc7, 0x5b, 0x25, 0xd9, 0x8f, 0x4b, 0xe4, 0x79, 0x33, 0xf3, 0x32, 0x79, 0x6f, 0x1c, 0xb8, 0xd8,
	0x4d, 0x93, 0x37, 0xc9, 0xcb, 0x09, 0xe5, 0x3b, 0x41, 0x33, 0xa2, 0x95, 0xcc, 0xf6, 0xe7, 0x44,
	0x2a, 0x61, 0x04, 0x1a, 0xd6, 0x72, 0xc1, 0x93, 0x42, 0x88, 0x82, 0x91, 0x09, 0x96, 0x74, 0x82,
	0x39, 0x17, 0x06, 0x1b, 0x2a, 0xb8, 0x2b, 0x0d, 0xce, 0x18, 0x2f, 0x09, 0x94, 0xcc, 0x2a, 0x20,
	0x7a, 0x0d, 0xa3, 0x8f, 0x98, 0x67, 0x84, 0x7d, 0xaa, 0x38, 0xbe, 0xe8, 0x02, 0x5d, 0xc0, 0xb1,
	0xc4, 0x57, 0x1b, 0xc2, 0xcd, 0x72, 0x8d, 0xf5, 0xda, 0xf7, 0x42, 0x2f, 0x3e, 0x4e, 0x87, 0x0e,
	0x9b, 0x63, 0xbd, 0x8e, 0x1e, 0xc2, 0x83, 0x46, 0x5b, 0x4a, 0xb4, 0x8c, 0xfe, 0x76, 0xe0, 0xf1,
	0xfb, 0x3c, 0x9f, 0x0b, 0x96, 0xef, 0xe1, 0xdf, 0x5b, 0xa2, 0x0d, 0x42, 0xd0, 0xdb, 0x90, 0x8d,
	0xb0, 0x4c, 0x47, 0xa9, 0x3d, 0x97, 0x98, 0x65, 0xef, 0x58, 0x76, 0x7b, 0x46, 0x8f, 0xe0, 0x60,
	0x87, 0xd9, 0x96, 0xf8, 0xdd, 0xd0, 0x8b, 0xbb, 0x69, 0x15, 0xa0, 0x4b, 0x18, 0xe5, 0x44, 0x67,
	0x8a, 0xca, 0xf2, 0x53, 0xaa, 0x99, 0x7a, 0xb6, 0xeb, 0x06, 0x8e, 0xc6, 0xd0, 0x27, 0x7f, 0x24,
	0x55, 0x57, 0xfe, 0x81, 0xa5, 0x70, 0x11, 0x7a, 0x0e, 0x27, 0x3f, 0x31, 0x63, 0x2b, 0x9c, 0xfd,
	0x5a, 0xe2, 0x3c, 0x57, 0x7e, 0xdf, 0x8e, 0xd2, 0x04, 0x51, 0x08, 0xc3, 0x8c, 0x99, 0xdd, 0xd2,
	0x51, 0x1c, 0x86, 0x5e, 0xdc, 0x4b, 0xeb, 0x10, 0x9a, 0xc1, 0x50, 0x89, 0xad, 0x21, 0xcb, 0x35,
	0xe5, 0x46, 0xfb, 0x83, 0xb0, 0x1b, 0x0f, 0x67, 0xa3, 0xc4, 0xca, 0x9a, 0xa4, 0x65, 0x66, 0x4e,
	0xb9, 0x49, 0xeb, 0x45, 0xc8, 0x87, 0x43, 0xa9, 0xe8, 0x0e, 0x1b, 0xe2, 0x1f, 0x85, 0x5e, 0x3c,
	0x48, 0xaf, 0xc3, 0xe8, 0x1d, 0xa0, 0xb6, 0x60, 0x5a, 0xa2, 0x18, 0xce, 0xae, 0xf5, 0x57, 0x95,
	0x80, 0x4e, 0xb8, 0x36, 0x1c, 0x25, 0x30, 0x5a, 0x10, 0x63, 0x18, 0xa9, 0xb9, 0x17, 0xc0, 0x40,
	0x2a, 0x42, 0x37, 0xb8, 0x20, 0xce, 0xb9, 0x7d, 0x5c, 0xda, 0xd6, 0xa8, 0xb7, 0xb6, 0xbd, 0x85,
	0xf3, 0xc5, 0x76, 0x55, 0xea, 0xb8, 0x22, 0x0b, 0xca, 0x8b, 0x5a, 0xb6, 0x72, 0x6f, 0x0c, 0x7d,
	0xb5, 0xac, 0x79, 0xe5, 0xa2, 0xcf, 0xbd, 0x81, 0x37, 0xea, 0xcc, 0xfe, 0x75, 0x60, 0xe0, 0x1a,
	0x34, 0xfa, 0x0e, 0xe3, 0xdb, 0xb9, 0xd0, 0x65, 0x52, 0xdb, 0xd2, 0xe4, 0xde, 0x17, 0x06, 0xa7,
	0x4e, 0x4f, 0x07, 0x4f, 0x3d, 0xf4, 0x15, 0x4e, 0x1a, 0xfb, 0x86, 0xce, 0x1b, 0x74, 0xed, 0x15,
	0x0e, 0x9e, 0xde, 0x9d, 0xb6, 0x12, 0x7f, 0x83, 0xd3, 0xa6, 0xf0, 0x28, 0x6a, 0x74, 0xdc, 0xba,
	0xc6, 0xc1, 0xb3, 0x7b, 0x6b, 0xb4, 0x2c, 0xc7, 0x6c, 0xe8, 0xdb, 0x1a, 0xb3, 0xed, 0x55, 0x6b,
	0xcc, 0x1b, 0xd6, 0x7c, 0x98, 0xfd, 0x98, 0x16, 0xd4, 0xac, 0xb7, 0xab, 0x24, 0x13, 0x9b, 0x09,
	0xe3, 0x59, 0xf9, 0xc8, 0x5f, 0x94, 0x57, 0xd8, 0xfd, 0x16, 0xaa, 0xfb, 0x5c, 0x63, 0x59, 0xf5,
	0xed, 0xc5, 0x7e, 0xf5, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x2d, 0xa1, 0x82, 0x39, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InvoicesClient is the client API for Invoices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InvoicesClient interface {
	//*
	//SubscribeSingleInvoice returns a uni-directional stream (server -> client)
	//to notify the client of state transitions of the specified invoice.
	//Initially the current invoice state is always sent out.
	SubscribeSingleInvoice(ctx context.Context, in *SubscribeSingleInvoiceRequest, opts ...grpc.CallOption) (Invoices_SubscribeSingleInvoiceClient, error)
	//*
	//CancelInvoice cancels a currently open invoice. If the invoice is already
	//canceled, this call will succeed. If the invoice is already settled, it will
	//fail.
	CancelInvoice(ctx context.Context, in *CancelInvoiceMsg, opts ...grpc.CallOption) (*CancelInvoiceResp, error)
	//*
	//AddHoldInvoice creates a hold invoice. It ties the invoice to the hash
	//supplied in the request.
	AddHoldInvoice(ctx context.Context, in *AddHoldInvoiceRequest, opts ...grpc.CallOption) (*AddHoldInvoiceResp, error)
	//*
	//SettleInvoice settles an accepted invoice. If the invoice is already
	//settled, this call will succeed.
	SettleInvoice(ctx context.Context, in *SettleInvoiceMsg, opts ...grpc.CallOption) (*SettleInvoiceResp, error)
}

type invoicesClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoicesClient(cc grpc.ClientConnInterface) InvoicesClient {
	return &invoicesClient{cc}
}

func (c *invoicesClient) SubscribeSingleInvoice(ctx context.Context, in *SubscribeSingleInvoiceRequest, opts ...grpc.CallOption) (Invoices_SubscribeSingleInvoiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Invoices_serviceDesc.Streams[0], "/invoicesrpc.Invoices/SubscribeSingleInvoice", opts...)
	if err != nil {
		return nil, err
	}
	x := &invoicesSubscribeSingleInvoiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Invoices_SubscribeSingleInvoiceClient interface {
	Recv() (*lnrpc.Invoice, error)
	grpc.ClientStream
}

type invoicesSubscribeSingleInvoiceClient struct {
	grpc.ClientStream
}

func (x *invoicesSubscribeSingleInvoiceClient) Recv() (*lnrpc.Invoice, error) {
	m := new(lnrpc.Invoice)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *invoicesClient) CancelInvoice(ctx context.Context, in *CancelInvoiceMsg, opts ...grpc.CallOption) (*CancelInvoiceResp, error) {
	out := new(CancelInvoiceResp)
	err := c.cc.Invoke(ctx, "/invoicesrpc.Invoices/CancelInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicesClient) AddHoldInvoice(ctx context.Context, in *AddHoldInvoiceRequest, opts ...grpc.CallOption) (*AddHoldInvoiceResp, error) {
	out := new(AddHoldInvoiceResp)
	err := c.cc.Invoke(ctx, "/invoicesrpc.Invoices/AddHoldInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *invoicesClient) SettleInvoice(ctx context.Context, in *SettleInvoiceMsg, opts ...grpc.CallOption) (*SettleInvoiceResp, error) {
	out := new(SettleInvoiceResp)
	err := c.cc.Invoke(ctx, "/invoicesrpc.Invoices/SettleInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoicesServer is the server API for Invoices service.
type InvoicesServer interface {
	//*
	//SubscribeSingleInvoice returns a uni-directional stream (server -> client)
	//to notify the client of state transitions of the specified invoice.
	//Initially the current invoice state is always sent out.
	SubscribeSingleInvoice(*SubscribeSingleInvoiceRequest, Invoices_SubscribeSingleInvoiceServer) error
	//*
	//CancelInvoice cancels a currently open invoice. If the invoice is already
	//canceled, this call will succeed. If the invoice is already settled, it will
	//fail.
	CancelInvoice(context.Context, *CancelInvoiceMsg) (*CancelInvoiceResp, error)
	//*
	//AddHoldInvoice creates a hold invoice. It ties the invoice to the hash
	//supplied in the request.
	AddHoldInvoice(context.Context, *AddHoldInvoiceRequest) (*AddHoldInvoiceResp, error)
	//*
	//SettleInvoice settles an accepted invoice. If the invoice is already
	//settled, this call will succeed.
	SettleInvoice(context.Context, *SettleInvoiceMsg) (*SettleInvoiceResp, error)
}

// UnimplementedInvoicesServer can be embedded to have forward compatible implementations.
type UnimplementedInvoicesServer struct {
}

func (*UnimplementedInvoicesServer) SubscribeSingleInvoice(req *SubscribeSingleInvoiceRequest, srv Invoices_SubscribeSingleInvoiceServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeSingleInvoice not implemented")
}
func (*UnimplementedInvoicesServer) CancelInvoice(ctx context.Context, req *CancelInvoiceMsg) (*CancelInvoiceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelInvoice not implemented")
}
func (*UnimplementedInvoicesServer) AddHoldInvoice(ctx context.Context, req *AddHoldInvoiceRequest) (*AddHoldInvoiceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHoldInvoice not implemented")
}
func (*UnimplementedInvoicesServer) SettleInvoice(ctx context.Context, req *SettleInvoiceMsg) (*SettleInvoiceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SettleInvoice not implemented")
}

func RegisterInvoicesServer(s *grpc.Server, srv InvoicesServer) {
	s.RegisterService(&_Invoices_serviceDesc, srv)
}

func _Invoices_SubscribeSingleInvoice_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeSingleInvoiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InvoicesServer).SubscribeSingleInvoice(m, &invoicesSubscribeSingleInvoiceServer{stream})
}

type Invoices_SubscribeSingleInvoiceServer interface {
	Send(*lnrpc.Invoice) error
	grpc.ServerStream
}

type invoicesSubscribeSingleInvoiceServer struct {
	grpc.ServerStream
}

func (x *invoicesSubscribeSingleInvoiceServer) Send(m *lnrpc.Invoice) error {
	return x.ServerStream.SendMsg(m)
}

func _Invoices_CancelInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelInvoiceMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicesServer).CancelInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoicesrpc.Invoices/CancelInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicesServer).CancelInvoice(ctx, req.(*CancelInvoiceMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoices_AddHoldInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHoldInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicesServer).AddHoldInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoicesrpc.Invoices/AddHoldInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicesServer).AddHoldInvoice(ctx, req.(*AddHoldInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Invoices_SettleInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettleInvoiceMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicesServer).SettleInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/invoicesrpc.Invoices/SettleInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicesServer).SettleInvoice(ctx, req.(*SettleInvoiceMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Invoices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "invoicesrpc.Invoices",
	HandlerType: (*InvoicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelInvoice",
			Handler:    _Invoices_CancelInvoice_Handler,
		},
		{
			MethodName: "AddHoldInvoice",
			Handler:    _Invoices_AddHoldInvoice_Handler,
		},
		{
			MethodName: "SettleInvoice",
			Handler:    _Invoices_SettleInvoice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeSingleInvoice",
			Handler:       _Invoices_SubscribeSingleInvoice_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v0.7.1/invoicesrpc/invoices.proto",
}

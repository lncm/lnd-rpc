// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v0.8.2/chainrpc/chainnotifier.proto

package chainrpc

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

type ConfRequest struct {
	//
	//The transaction hash for which we should request a confirmation notification
	//for. If set to a hash of all zeros, then the confirmation notification will
	//be requested for the script instead.
	Txid []byte `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	//
	//An output script within a transaction with the hash above which will be used
	//by light clients to match block filters. If the transaction hash is set to a
	//hash of all zeros, then a confirmation notification will be requested for
	//this script instead.
	Script []byte `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	//
	//The number of desired confirmations the transaction/output script should
	//reach before dispatching a confirmation notification.
	NumConfs uint32 `protobuf:"varint,3,opt,name=num_confs,json=numConfs,proto3" json:"num_confs,omitempty"`
	//
	//The earliest height in the chain for which the transaction/output script
	//could have been included in a block. This should in most cases be set to the
	//broadcast height of the transaction/output script.
	HeightHint           uint32   `protobuf:"varint,4,opt,name=height_hint,json=heightHint,proto3" json:"height_hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfRequest) Reset()         { *m = ConfRequest{} }
func (m *ConfRequest) String() string { return proto.CompactTextString(m) }
func (*ConfRequest) ProtoMessage()    {}
func (*ConfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aabde7db09ab05a3, []int{0}
}

func (m *ConfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfRequest.Unmarshal(m, b)
}
func (m *ConfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfRequest.Marshal(b, m, deterministic)
}
func (m *ConfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfRequest.Merge(m, src)
}
func (m *ConfRequest) XXX_Size() int {
	return xxx_messageInfo_ConfRequest.Size(m)
}
func (m *ConfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfRequest proto.InternalMessageInfo

func (m *ConfRequest) GetTxid() []byte {
	if m != nil {
		return m.Txid
	}
	return nil
}

func (m *ConfRequest) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *ConfRequest) GetNumConfs() uint32 {
	if m != nil {
		return m.NumConfs
	}
	return 0
}

func (m *ConfRequest) GetHeightHint() uint32 {
	if m != nil {
		return m.HeightHint
	}
	return 0
}

type ConfDetails struct {
	// The raw bytes of the confirmed transaction.
	RawTx []byte `protobuf:"bytes,1,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	// The hash of the block in which the confirmed transaction was included in.
	BlockHash []byte `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// The height of the block in which the confirmed transaction was included in.
	BlockHeight uint32 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// The index of the confirmed transaction within the transaction.
	TxIndex              uint32   `protobuf:"varint,4,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfDetails) Reset()         { *m = ConfDetails{} }
func (m *ConfDetails) String() string { return proto.CompactTextString(m) }
func (*ConfDetails) ProtoMessage()    {}
func (*ConfDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_aabde7db09ab05a3, []int{1}
}

func (m *ConfDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfDetails.Unmarshal(m, b)
}
func (m *ConfDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfDetails.Marshal(b, m, deterministic)
}
func (m *ConfDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfDetails.Merge(m, src)
}
func (m *ConfDetails) XXX_Size() int {
	return xxx_messageInfo_ConfDetails.Size(m)
}
func (m *ConfDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ConfDetails proto.InternalMessageInfo

func (m *ConfDetails) GetRawTx() []byte {
	if m != nil {
		return m.RawTx
	}
	return nil
}

func (m *ConfDetails) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *ConfDetails) GetBlockHeight() uint32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *ConfDetails) GetTxIndex() uint32 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

type Reorg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reorg) Reset()         { *m = Reorg{} }
func (m *Reorg) String() string { return proto.CompactTextString(m) }
func (*Reorg) ProtoMessage()    {}
func (*Reorg) Descriptor() ([]byte, []int) {
	return fileDescriptor_aabde7db09ab05a3, []int{2}
}

func (m *Reorg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reorg.Unmarshal(m, b)
}
func (m *Reorg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reorg.Marshal(b, m, deterministic)
}
func (m *Reorg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reorg.Merge(m, src)
}
func (m *Reorg) XXX_Size() int {
	return xxx_messageInfo_Reorg.Size(m)
}
func (m *Reorg) XXX_DiscardUnknown() {
	xxx_messageInfo_Reorg.DiscardUnknown(m)
}

var xxx_messageInfo_Reorg proto.InternalMessageInfo

type ConfEvent struct {
	// Types that are valid to be assigned to Event:
	//	*ConfEvent_Conf
	//	*ConfEvent_Reorg
	Event                isConfEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConfEvent) Reset()         { *m = ConfEvent{} }
func (m *ConfEvent) String() string { return proto.CompactTextString(m) }
func (*ConfEvent) ProtoMessage()    {}
func (*ConfEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_aabde7db09ab05a3, []int{3}
}

func (m *ConfEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfEvent.Unmarshal(m, b)
}
func (m *ConfEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfEvent.Marshal(b, m, deterministic)
}
func (m *ConfEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfEvent.Merge(m, src)
}
func (m *ConfEvent) XXX_Size() int {
	return xxx_messageInfo_ConfEvent.Size(m)
}
func (m *ConfEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ConfEvent proto.InternalMessageInfo

type isConfEvent_Event interface {
	isConfEvent_Event()
}

type ConfEvent_Conf struct {
	Conf *ConfDetails `protobuf:"bytes,1,opt,name=conf,proto3,oneof"`
}

type ConfEvent_Reorg struct {
	Reorg *Reorg `protobuf:"bytes,2,opt,name=reorg,proto3,oneof"`
}

func (*ConfEvent_Conf) isConfEvent_Event() {}

func (*ConfEvent_Reorg) isConfEvent_Event() {}

func (m *ConfEvent) GetEvent() isConfEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ConfEvent) GetConf() *ConfDetails {
	if x, ok := m.GetEvent().(*ConfEvent_Conf); ok {
		return x.Conf
	}
	return nil
}

func (m *ConfEvent) GetReorg() *Reorg {
	if x, ok := m.GetEvent().(*ConfEvent_Reorg); ok {
		return x.Reorg
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConfEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConfEvent_Conf)(nil),
		(*ConfEvent_Reorg)(nil),
	}
}

type Outpoint struct {
	// The hash of the transaction.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// The index of the output within the transaction.
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Outpoint) Reset()         { *m = Outpoint{} }
func (m *Outpoint) String() string { return proto.CompactTextString(m) }
func (*Outpoint) ProtoMessage()    {}
func (*Outpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_aabde7db09ab05a3, []int{4}
}

func (m *Outpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Outpoint.Unmarshal(m, b)
}
func (m *Outpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Outpoint.Marshal(b, m, deterministic)
}
func (m *Outpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outpoint.Merge(m, src)
}
func (m *Outpoint) XXX_Size() int {
	return xxx_messageInfo_Outpoint.Size(m)
}
func (m *Outpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Outpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Outpoint proto.InternalMessageInfo

func (m *Outpoint) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Outpoint) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type SpendRequest struct {
	//
	//The outpoint for which we should request a spend notification for. If set to
	//a zero outpoint, then the spend notification will be requested for the
	//script instead.
	Outpoint *Outpoint `protobuf:"bytes,1,opt,name=outpoint,proto3" json:"outpoint,omitempty"`
	//
	//The output script for the outpoint above. This will be used by light clients
	//to match block filters. If the outpoint is set to a zero outpoint, then a
	//spend notification will be requested for this script instead.
	Script []byte `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	//
	//The earliest height in the chain for which the outpoint/output script could
	//have been spent. This should in most cases be set to the broadcast height of
	//the outpoint/output script.
	HeightHint           uint32   `protobuf:"varint,3,opt,name=height_hint,json=heightHint,proto3" json:"height_hint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpendRequest) Reset()         { *m = SpendRequest{} }
func (m *SpendRequest) String() string { return proto.CompactTextString(m) }
func (*SpendRequest) ProtoMessage()    {}
func (*SpendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aabde7db09ab05a3, []int{5}
}

func (m *SpendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpendRequest.Unmarshal(m, b)
}
func (m *SpendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpendRequest.Marshal(b, m, deterministic)
}
func (m *SpendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpendRequest.Merge(m, src)
}
func (m *SpendRequest) XXX_Size() int {
	return xxx_messageInfo_SpendRequest.Size(m)
}
func (m *SpendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpendRequest proto.InternalMessageInfo

func (m *SpendRequest) GetOutpoint() *Outpoint {
	if m != nil {
		return m.Outpoint
	}
	return nil
}

func (m *SpendRequest) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *SpendRequest) GetHeightHint() uint32 {
	if m != nil {
		return m.HeightHint
	}
	return 0
}

type SpendDetails struct {
	// The outpoint was that spent.
	SpendingOutpoint *Outpoint `protobuf:"bytes,1,opt,name=spending_outpoint,json=spendingOutpoint,proto3" json:"spending_outpoint,omitempty"`
	// The raw bytes of the spending transaction.
	RawSpendingTx []byte `protobuf:"bytes,2,opt,name=raw_spending_tx,json=rawSpendingTx,proto3" json:"raw_spending_tx,omitempty"`
	// The hash of the spending transaction.
	SpendingTxHash []byte `protobuf:"bytes,3,opt,name=spending_tx_hash,json=spendingTxHash,proto3" json:"spending_tx_hash,omitempty"`
	// The input of the spending transaction that fulfilled the spend request.
	SpendingInputIndex uint32 `protobuf:"varint,4,opt,name=spending_input_index,json=spendingInputIndex,proto3" json:"spending_input_index,omitempty"`
	// The height at which the spending transaction was included in a block.
	SpendingHeight       uint32   `protobuf:"varint,5,opt,name=spending_height,json=spendingHeight,proto3" json:"spending_height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpendDetails) Reset()         { *m = SpendDetails{} }
func (m *SpendDetails) String() string { return proto.CompactTextString(m) }
func (*SpendDetails) ProtoMessage()    {}
func (*SpendDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_aabde7db09ab05a3, []int{6}
}

func (m *SpendDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpendDetails.Unmarshal(m, b)
}
func (m *SpendDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpendDetails.Marshal(b, m, deterministic)
}
func (m *SpendDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpendDetails.Merge(m, src)
}
func (m *SpendDetails) XXX_Size() int {
	return xxx_messageInfo_SpendDetails.Size(m)
}
func (m *SpendDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_SpendDetails.DiscardUnknown(m)
}

var xxx_messageInfo_SpendDetails proto.InternalMessageInfo

func (m *SpendDetails) GetSpendingOutpoint() *Outpoint {
	if m != nil {
		return m.SpendingOutpoint
	}
	return nil
}

func (m *SpendDetails) GetRawSpendingTx() []byte {
	if m != nil {
		return m.RawSpendingTx
	}
	return nil
}

func (m *SpendDetails) GetSpendingTxHash() []byte {
	if m != nil {
		return m.SpendingTxHash
	}
	return nil
}

func (m *SpendDetails) GetSpendingInputIndex() uint32 {
	if m != nil {
		return m.SpendingInputIndex
	}
	return 0
}

func (m *SpendDetails) GetSpendingHeight() uint32 {
	if m != nil {
		return m.SpendingHeight
	}
	return 0
}

type SpendEvent struct {
	// Types that are valid to be assigned to Event:
	//	*SpendEvent_Spend
	//	*SpendEvent_Reorg
	Event                isSpendEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SpendEvent) Reset()         { *m = SpendEvent{} }
func (m *SpendEvent) String() string { return proto.CompactTextString(m) }
func (*SpendEvent) ProtoMessage()    {}
func (*SpendEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_aabde7db09ab05a3, []int{7}
}

func (m *SpendEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpendEvent.Unmarshal(m, b)
}
func (m *SpendEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpendEvent.Marshal(b, m, deterministic)
}
func (m *SpendEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpendEvent.Merge(m, src)
}
func (m *SpendEvent) XXX_Size() int {
	return xxx_messageInfo_SpendEvent.Size(m)
}
func (m *SpendEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SpendEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SpendEvent proto.InternalMessageInfo

type isSpendEvent_Event interface {
	isSpendEvent_Event()
}

type SpendEvent_Spend struct {
	Spend *SpendDetails `protobuf:"bytes,1,opt,name=spend,proto3,oneof"`
}

type SpendEvent_Reorg struct {
	Reorg *Reorg `protobuf:"bytes,2,opt,name=reorg,proto3,oneof"`
}

func (*SpendEvent_Spend) isSpendEvent_Event() {}

func (*SpendEvent_Reorg) isSpendEvent_Event() {}

func (m *SpendEvent) GetEvent() isSpendEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *SpendEvent) GetSpend() *SpendDetails {
	if x, ok := m.GetEvent().(*SpendEvent_Spend); ok {
		return x.Spend
	}
	return nil
}

func (m *SpendEvent) GetReorg() *Reorg {
	if x, ok := m.GetEvent().(*SpendEvent_Reorg); ok {
		return x.Reorg
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SpendEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SpendEvent_Spend)(nil),
		(*SpendEvent_Reorg)(nil),
	}
}

type BlockEpoch struct {
	// The hash of the block.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// The height of the block.
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockEpoch) Reset()         { *m = BlockEpoch{} }
func (m *BlockEpoch) String() string { return proto.CompactTextString(m) }
func (*BlockEpoch) ProtoMessage()    {}
func (*BlockEpoch) Descriptor() ([]byte, []int) {
	return fileDescriptor_aabde7db09ab05a3, []int{8}
}

func (m *BlockEpoch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockEpoch.Unmarshal(m, b)
}
func (m *BlockEpoch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockEpoch.Marshal(b, m, deterministic)
}
func (m *BlockEpoch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockEpoch.Merge(m, src)
}
func (m *BlockEpoch) XXX_Size() int {
	return xxx_messageInfo_BlockEpoch.Size(m)
}
func (m *BlockEpoch) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockEpoch.DiscardUnknown(m)
}

var xxx_messageInfo_BlockEpoch proto.InternalMessageInfo

func (m *BlockEpoch) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockEpoch) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*ConfRequest)(nil), "chainrpc.ConfRequest")
	proto.RegisterType((*ConfDetails)(nil), "chainrpc.ConfDetails")
	proto.RegisterType((*Reorg)(nil), "chainrpc.Reorg")
	proto.RegisterType((*ConfEvent)(nil), "chainrpc.ConfEvent")
	proto.RegisterType((*Outpoint)(nil), "chainrpc.Outpoint")
	proto.RegisterType((*SpendRequest)(nil), "chainrpc.SpendRequest")
	proto.RegisterType((*SpendDetails)(nil), "chainrpc.SpendDetails")
	proto.RegisterType((*SpendEvent)(nil), "chainrpc.SpendEvent")
	proto.RegisterType((*BlockEpoch)(nil), "chainrpc.BlockEpoch")
}

func init() {
	proto.RegisterFile("v0.8.2/chainrpc/chainnotifier.proto", fileDescriptor_aabde7db09ab05a3)
}

var fileDescriptor_aabde7db09ab05a3 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x6f, 0xda, 0x3e,
	0x14, 0x6d, 0x5a, 0xc2, 0x9f, 0x0b, 0xfc, 0x68, 0xfd, 0xa3, 0x88, 0x76, 0x9a, 0xd6, 0x65, 0xd2,
	0x8a, 0x34, 0x29, 0x43, 0x6c, 0x0f, 0x7d, 0x9b, 0x54, 0xd6, 0x09, 0x5e, 0x3a, 0x29, 0xed, 0x7b,
	0x94, 0x06, 0x43, 0xbc, 0x15, 0x27, 0x8b, 0x4d, 0xc9, 0xeb, 0x3e, 0xed, 0xbe, 0xc4, 0x1e, 0x26,
	0x5f, 0xdb, 0x81, 0xb2, 0x4e, 0x9a, 0xf6, 0xe6, 0x7b, 0xef, 0xc9, 0xf1, 0xb9, 0x3e, 0x07, 0xe0,
	0xd5, 0xc3, 0xd0, 0xbf, 0xf0, 0x47, 0x6f, 0xe3, 0x24, 0x62, 0x3c, 0xcf, 0x62, 0x7d, 0xe0, 0xa9,
	0x64, 0x73, 0x46, 0x73, 0x3f, 0xcb, 0x53, 0x99, 0x92, 0xba, 0x9d, 0x7a, 0x6b, 0x68, 0x8e, 0x53,
	0x3e, 0x0f, 0xe8, 0xb7, 0x15, 0x15, 0x92, 0x10, 0xa8, 0xc8, 0x82, 0xcd, 0xfa, 0xce, 0x99, 0x33,
	0x68, 0x05, 0x78, 0x26, 0x3d, 0xa8, 0x8a, 0x38, 0x67, 0x99, 0xec, 0xef, 0x63, 0xd7, 0x54, 0xe4,
	0x19, 0x34, 0xf8, 0x6a, 0x19, 0xc6, 0x29, 0x9f, 0x8b, 0xfe, 0xc1, 0x99, 0x33, 0x68, 0x07, 0x75,
	0xbe, 0x5a, 0x2a, 0x3a, 0x41, 0x5e, 0x40, 0x33, 0xa1, 0x6c, 0x91, 0xc8, 0x30, 0x61, 0x5c, 0xf6,
	0x2b, 0x38, 0x06, 0xdd, 0x9a, 0x30, 0x2e, 0xbd, 0xef, 0x8e, 0xbe, 0xf9, 0x23, 0x95, 0x11, 0xbb,
	0x17, 0xe4, 0x18, 0xaa, 0x79, 0xb4, 0x0e, 0x65, 0x61, 0xee, 0x76, 0xf3, 0x68, 0x7d, 0x5b, 0x90,
	0xe7, 0x00, 0x77, 0xf7, 0x69, 0xfc, 0x35, 0x4c, 0x22, 0x91, 0x18, 0x01, 0x0d, 0xec, 0x4c, 0x22,
	0x91, 0x90, 0x97, 0xd0, 0x32, 0x63, 0x64, 0x36, 0x32, 0x9a, 0x1a, 0x80, 0x2d, 0x72, 0x02, 0x75,
	0x59, 0x84, 0x8c, 0xcf, 0x68, 0x61, 0x64, 0xd4, 0x64, 0x31, 0x55, 0xa5, 0x57, 0x03, 0x37, 0xa0,
	0x69, 0xbe, 0xf0, 0xbe, 0x40, 0x43, 0x69, 0xb9, 0x7a, 0xa0, 0x5c, 0x92, 0x37, 0x50, 0x51, 0x3b,
	0xa1, 0x8e, 0xe6, 0xe8, 0xd8, 0xb7, 0x6f, 0xe5, 0x6f, 0xc9, 0x9d, 0xec, 0x05, 0x08, 0x22, 0xe7,
	0xe0, 0xe6, 0x8a, 0x02, 0xa5, 0x35, 0x47, 0x9d, 0x0d, 0x1a, 0x99, 0x27, 0x7b, 0x81, 0x9e, 0x5f,
	0xd6, 0xc0, 0xa5, 0x8a, 0xde, 0x7b, 0x0f, 0xf5, 0xcf, 0x2b, 0x99, 0xa5, 0x8c, 0xe3, 0x73, 0xe3,
	0x5e, 0xe6, 0xb9, 0xd5, 0x99, 0x74, 0xc1, 0xd5, 0x62, 0xf7, 0x51, 0xac, 0x2e, 0xbc, 0x35, 0xb4,
	0x6e, 0x32, 0xca, 0x67, 0xd6, 0x28, 0x1f, 0xea, 0xa9, 0x61, 0x31, 0x42, 0xc9, 0xe6, 0x6a, 0xcb,
	0x1f, 0x94, 0x98, 0x3f, 0x9a, 0xb8, 0xe3, 0xd3, 0xc1, 0x6f, 0x3e, 0xfd, 0x74, 0xcc, 0xcd, 0xd6,
	0xa8, 0x0f, 0x70, 0x24, 0x54, 0xcd, 0xf8, 0x22, 0xfc, 0x0b, 0x09, 0x87, 0x16, 0x5c, 0x2e, 0xfd,
	0x1a, 0x3a, 0xca, 0xe9, 0x92, 0x44, 0x16, 0x46, 0x53, 0x3b, 0x8f, 0xd6, 0x37, 0xa6, 0x7b, 0x5b,
	0x90, 0x01, 0x1c, 0x6e, 0x61, 0x74, 0x00, 0x0e, 0x10, 0xf8, 0x9f, 0x28, 0x51, 0x98, 0x82, 0x21,
	0x74, 0x4b, 0x24, 0xe3, 0xd9, 0x4a, 0x3e, 0xb2, 0x9b, 0xd8, 0xd9, 0x54, 0x8d, 0xd0, 0x79, 0x72,
	0x0e, 0x9d, 0xf2, 0x0b, 0x13, 0x1d, 0x17, 0xc1, 0x25, 0xb5, 0x4e, 0x8f, 0xc7, 0x01, 0x50, 0x92,
	0x8e, 0x86, 0x0f, 0x2e, 0xce, 0xcd, 0xbe, 0xbd, 0xcd, 0xbe, 0xdb, 0x4f, 0xa4, 0x4c, 0x47, 0xd8,
	0x3f, 0xa4, 0xe3, 0x02, 0xe0, 0x52, 0x85, 0xf7, 0x2a, 0x4b, 0xe3, 0xe4, 0xc9, 0x7c, 0xf4, 0xa0,
	0x6a, 0x14, 0xeb, 0x80, 0x98, 0x6a, 0xf4, 0xc3, 0x81, 0xf6, 0x58, 0xd1, 0x5f, 0x9b, 0xdf, 0x3a,
	0x99, 0xc2, 0x49, 0x40, 0x17, 0x4c, 0x48, 0x9a, 0xab, 0xe8, 0xb2, 0x7c, 0x19, 0x49, 0x96, 0x72,
	0x71, 0x2d, 0xe7, 0x9c, 0xec, 0xe4, 0xda, 0xe4, 0xea, 0xf4, 0xff, 0xc7, 0x6d, 0x5c, 0x7b, 0xe8,
	0x90, 0x31, 0x1c, 0x59, 0x2a, 0xdc, 0x14, 0x29, 0x76, 0xd7, 0xb7, 0x1c, 0xdd, 0x9d, 0xbe, 0x25,
	0xf9, 0x04, 0x3d, 0x4b, 0xb2, 0xd9, 0x11, 0x99, 0xb6, 0xbe, 0xd8, 0x4c, 0x4e, 0x9f, 0xec, 0x0e,
	0x9d, 0xbb, 0x2a, 0xfe, 0x89, 0xbd, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x70, 0x42, 0x04, 0xbf,
	0xeb, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChainNotifierClient is the client API for ChainNotifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChainNotifierClient interface {
	//
	//RegisterConfirmationsNtfn is a synchronous response-streaming RPC that
	//registers an intent for a client to be notified once a confirmation request
	//has reached its required number of confirmations on-chain.
	//
	//A client can specify whether the confirmation request should be for a
	//particular transaction by its hash or for an output script by specifying a
	//zero hash.
	RegisterConfirmationsNtfn(ctx context.Context, in *ConfRequest, opts ...grpc.CallOption) (ChainNotifier_RegisterConfirmationsNtfnClient, error)
	//
	//RegisterSpendNtfn is a synchronous response-streaming RPC that registers an
	//intent for a client to be notification once a spend request has been spent
	//by a transaction that has confirmed on-chain.
	//
	//A client can specify whether the spend request should be for a particular
	//outpoint  or for an output script by specifying a zero outpoint.
	RegisterSpendNtfn(ctx context.Context, in *SpendRequest, opts ...grpc.CallOption) (ChainNotifier_RegisterSpendNtfnClient, error)
	//
	//RegisterBlockEpochNtfn is a synchronous response-streaming RPC that
	//registers an intent for a client to be notified of blocks in the chain. The
	//stream will return a hash and height tuple of a block for each new/stale
	//block in the chain. It is the client's responsibility to determine whether
	//the tuple returned is for a new or stale block in the chain.
	//
	//A client can also request a historical backlog of blocks from a particular
	//point. This allows clients to be idempotent by ensuring that they do not
	//missing processing a single block within the chain.
	RegisterBlockEpochNtfn(ctx context.Context, in *BlockEpoch, opts ...grpc.CallOption) (ChainNotifier_RegisterBlockEpochNtfnClient, error)
}

type chainNotifierClient struct {
	cc *grpc.ClientConn
}

func NewChainNotifierClient(cc *grpc.ClientConn) ChainNotifierClient {
	return &chainNotifierClient{cc}
}

func (c *chainNotifierClient) RegisterConfirmationsNtfn(ctx context.Context, in *ConfRequest, opts ...grpc.CallOption) (ChainNotifier_RegisterConfirmationsNtfnClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChainNotifier_serviceDesc.Streams[0], "/chainrpc.ChainNotifier/RegisterConfirmationsNtfn", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainNotifierRegisterConfirmationsNtfnClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChainNotifier_RegisterConfirmationsNtfnClient interface {
	Recv() (*ConfEvent, error)
	grpc.ClientStream
}

type chainNotifierRegisterConfirmationsNtfnClient struct {
	grpc.ClientStream
}

func (x *chainNotifierRegisterConfirmationsNtfnClient) Recv() (*ConfEvent, error) {
	m := new(ConfEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chainNotifierClient) RegisterSpendNtfn(ctx context.Context, in *SpendRequest, opts ...grpc.CallOption) (ChainNotifier_RegisterSpendNtfnClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChainNotifier_serviceDesc.Streams[1], "/chainrpc.ChainNotifier/RegisterSpendNtfn", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainNotifierRegisterSpendNtfnClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChainNotifier_RegisterSpendNtfnClient interface {
	Recv() (*SpendEvent, error)
	grpc.ClientStream
}

type chainNotifierRegisterSpendNtfnClient struct {
	grpc.ClientStream
}

func (x *chainNotifierRegisterSpendNtfnClient) Recv() (*SpendEvent, error) {
	m := new(SpendEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chainNotifierClient) RegisterBlockEpochNtfn(ctx context.Context, in *BlockEpoch, opts ...grpc.CallOption) (ChainNotifier_RegisterBlockEpochNtfnClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChainNotifier_serviceDesc.Streams[2], "/chainrpc.ChainNotifier/RegisterBlockEpochNtfn", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainNotifierRegisterBlockEpochNtfnClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChainNotifier_RegisterBlockEpochNtfnClient interface {
	Recv() (*BlockEpoch, error)
	grpc.ClientStream
}

type chainNotifierRegisterBlockEpochNtfnClient struct {
	grpc.ClientStream
}

func (x *chainNotifierRegisterBlockEpochNtfnClient) Recv() (*BlockEpoch, error) {
	m := new(BlockEpoch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChainNotifierServer is the server API for ChainNotifier service.
type ChainNotifierServer interface {
	//
	//RegisterConfirmationsNtfn is a synchronous response-streaming RPC that
	//registers an intent for a client to be notified once a confirmation request
	//has reached its required number of confirmations on-chain.
	//
	//A client can specify whether the confirmation request should be for a
	//particular transaction by its hash or for an output script by specifying a
	//zero hash.
	RegisterConfirmationsNtfn(*ConfRequest, ChainNotifier_RegisterConfirmationsNtfnServer) error
	//
	//RegisterSpendNtfn is a synchronous response-streaming RPC that registers an
	//intent for a client to be notification once a spend request has been spent
	//by a transaction that has confirmed on-chain.
	//
	//A client can specify whether the spend request should be for a particular
	//outpoint  or for an output script by specifying a zero outpoint.
	RegisterSpendNtfn(*SpendRequest, ChainNotifier_RegisterSpendNtfnServer) error
	//
	//RegisterBlockEpochNtfn is a synchronous response-streaming RPC that
	//registers an intent for a client to be notified of blocks in the chain. The
	//stream will return a hash and height tuple of a block for each new/stale
	//block in the chain. It is the client's responsibility to determine whether
	//the tuple returned is for a new or stale block in the chain.
	//
	//A client can also request a historical backlog of blocks from a particular
	//point. This allows clients to be idempotent by ensuring that they do not
	//missing processing a single block within the chain.
	RegisterBlockEpochNtfn(*BlockEpoch, ChainNotifier_RegisterBlockEpochNtfnServer) error
}

// UnimplementedChainNotifierServer can be embedded to have forward compatible implementations.
type UnimplementedChainNotifierServer struct {
}

func (*UnimplementedChainNotifierServer) RegisterConfirmationsNtfn(req *ConfRequest, srv ChainNotifier_RegisterConfirmationsNtfnServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterConfirmationsNtfn not implemented")
}
func (*UnimplementedChainNotifierServer) RegisterSpendNtfn(req *SpendRequest, srv ChainNotifier_RegisterSpendNtfnServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterSpendNtfn not implemented")
}
func (*UnimplementedChainNotifierServer) RegisterBlockEpochNtfn(req *BlockEpoch, srv ChainNotifier_RegisterBlockEpochNtfnServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterBlockEpochNtfn not implemented")
}

func RegisterChainNotifierServer(s *grpc.Server, srv ChainNotifierServer) {
	s.RegisterService(&_ChainNotifier_serviceDesc, srv)
}

func _ChainNotifier_RegisterConfirmationsNtfn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConfRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainNotifierServer).RegisterConfirmationsNtfn(m, &chainNotifierRegisterConfirmationsNtfnServer{stream})
}

type ChainNotifier_RegisterConfirmationsNtfnServer interface {
	Send(*ConfEvent) error
	grpc.ServerStream
}

type chainNotifierRegisterConfirmationsNtfnServer struct {
	grpc.ServerStream
}

func (x *chainNotifierRegisterConfirmationsNtfnServer) Send(m *ConfEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ChainNotifier_RegisterSpendNtfn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SpendRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainNotifierServer).RegisterSpendNtfn(m, &chainNotifierRegisterSpendNtfnServer{stream})
}

type ChainNotifier_RegisterSpendNtfnServer interface {
	Send(*SpendEvent) error
	grpc.ServerStream
}

type chainNotifierRegisterSpendNtfnServer struct {
	grpc.ServerStream
}

func (x *chainNotifierRegisterSpendNtfnServer) Send(m *SpendEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _ChainNotifier_RegisterBlockEpochNtfn_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockEpoch)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainNotifierServer).RegisterBlockEpochNtfn(m, &chainNotifierRegisterBlockEpochNtfnServer{stream})
}

type ChainNotifier_RegisterBlockEpochNtfnServer interface {
	Send(*BlockEpoch) error
	grpc.ServerStream
}

type chainNotifierRegisterBlockEpochNtfnServer struct {
	grpc.ServerStream
}

func (x *chainNotifierRegisterBlockEpochNtfnServer) Send(m *BlockEpoch) error {
	return x.ServerStream.SendMsg(m)
}

var _ChainNotifier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chainrpc.ChainNotifier",
	HandlerType: (*ChainNotifierServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterConfirmationsNtfn",
			Handler:       _ChainNotifier_RegisterConfirmationsNtfn_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RegisterSpendNtfn",
			Handler:       _ChainNotifier_RegisterSpendNtfn_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RegisterBlockEpochNtfn",
			Handler:       _ChainNotifier_RegisterBlockEpochNtfn_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v0.8.2/chainrpc/chainnotifier.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fastchain.proto

package fastchain

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

type Request struct {
	Inner                *Request_Inner        `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
	Dig                  []byte                `protobuf:"bytes,2,opt,name=dig,proto3" json:"dig,omitempty"`
	Sig                  *Request_MsgSignature `protobuf:"bytes,3,opt,name=sig,proto3" json:"sig,omitempty"`
	Outer                []byte                `protobuf:"bytes,4,opt,name=outer,proto3" json:"outer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetInner() *Request_Inner {
	if m != nil {
		return m.Inner
	}
	return nil
}

func (m *Request) GetDig() []byte {
	if m != nil {
		return m.Dig
	}
	return nil
}

func (m *Request) GetSig() *Request_MsgSignature {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *Request) GetOuter() []byte {
	if m != nil {
		return m.Outer
	}
	return nil
}

type Request_Inner struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Seq                  int32    `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	View                 int32    `protobuf:"varint,3,opt,name=view,proto3" json:"view,omitempty"`
	Type                 int32    `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Msg                  []byte   `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp            int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_Inner) Reset()         { *m = Request_Inner{} }
func (m *Request_Inner) String() string { return proto.CompactTextString(m) }
func (*Request_Inner) ProtoMessage()    {}
func (*Request_Inner) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{0, 0}
}
func (m *Request_Inner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_Inner.Unmarshal(m, b)
}
func (m *Request_Inner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_Inner.Marshal(b, m, deterministic)
}
func (dst *Request_Inner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_Inner.Merge(dst, src)
}
func (m *Request_Inner) XXX_Size() int {
	return xxx_messageInfo_Request_Inner.Size(m)
}
func (m *Request_Inner) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_Inner.DiscardUnknown(m)
}

var xxx_messageInfo_Request_Inner proto.InternalMessageInfo

func (m *Request_Inner) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request_Inner) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Request_Inner) GetView() int32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *Request_Inner) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Request_Inner) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Request_Inner) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Request_MsgSignature struct {
	R                    int64    `protobuf:"varint,1,opt,name=r,proto3" json:"r,omitempty"`
	S                    int64    `protobuf:"varint,2,opt,name=s,proto3" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_MsgSignature) Reset()         { *m = Request_MsgSignature{} }
func (m *Request_MsgSignature) String() string { return proto.CompactTextString(m) }
func (*Request_MsgSignature) ProtoMessage()    {}
func (*Request_MsgSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{0, 1}
}
func (m *Request_MsgSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_MsgSignature.Unmarshal(m, b)
}
func (m *Request_MsgSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_MsgSignature.Marshal(b, m, deterministic)
}
func (dst *Request_MsgSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_MsgSignature.Merge(dst, src)
}
func (m *Request_MsgSignature) XXX_Size() int {
	return xxx_messageInfo_Request_MsgSignature.Size(m)
}
func (m *Request_MsgSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_MsgSignature.DiscardUnknown(m)
}

var xxx_messageInfo_Request_MsgSignature proto.InternalMessageInfo

func (m *Request_MsgSignature) GetR() int64 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *Request_MsgSignature) GetS() int64 {
	if m != nil {
		return m.S
	}
	return 0
}

type CheckLeaderReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckLeaderReq) Reset()         { *m = CheckLeaderReq{} }
func (m *CheckLeaderReq) String() string { return proto.CompactTextString(m) }
func (*CheckLeaderReq) ProtoMessage()    {}
func (*CheckLeaderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{1}
}
func (m *CheckLeaderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckLeaderReq.Unmarshal(m, b)
}
func (m *CheckLeaderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckLeaderReq.Marshal(b, m, deterministic)
}
func (dst *CheckLeaderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckLeaderReq.Merge(dst, src)
}
func (m *CheckLeaderReq) XXX_Size() int {
	return xxx_messageInfo_CheckLeaderReq.Size(m)
}
func (m *CheckLeaderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckLeaderReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckLeaderReq proto.InternalMessageInfo

type CheckLeaderResp struct {
	Message              bool     `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckLeaderResp) Reset()         { *m = CheckLeaderResp{} }
func (m *CheckLeaderResp) String() string { return proto.CompactTextString(m) }
func (*CheckLeaderResp) ProtoMessage()    {}
func (*CheckLeaderResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{2}
}
func (m *CheckLeaderResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckLeaderResp.Unmarshal(m, b)
}
func (m *CheckLeaderResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckLeaderResp.Marshal(b, m, deterministic)
}
func (dst *CheckLeaderResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckLeaderResp.Merge(dst, src)
}
func (m *CheckLeaderResp) XXX_Size() int {
	return xxx_messageInfo_CheckLeaderResp.Size(m)
}
func (m *CheckLeaderResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckLeaderResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckLeaderResp proto.InternalMessageInfo

func (m *CheckLeaderResp) GetMessage() bool {
	if m != nil {
		return m.Message
	}
	return false
}

type PbftNode struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Pubkey               string   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Privkey              string   `protobuf:"bytes,3,opt,name=privkey,proto3" json:"privkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbftNode) Reset()         { *m = PbftNode{} }
func (m *PbftNode) String() string { return proto.CompactTextString(m) }
func (*PbftNode) ProtoMessage()    {}
func (*PbftNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{3}
}
func (m *PbftNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftNode.Unmarshal(m, b)
}
func (m *PbftNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftNode.Marshal(b, m, deterministic)
}
func (dst *PbftNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftNode.Merge(dst, src)
}
func (m *PbftNode) XXX_Size() int {
	return xxx_messageInfo_PbftNode.Size(m)
}
func (m *PbftNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftNode.DiscardUnknown(m)
}

var xxx_messageInfo_PbftNode proto.InternalMessageInfo

func (m *PbftNode) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *PbftNode) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *PbftNode) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

type Nodes struct {
	Nodes                []*PbftNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Nodes) Reset()         { *m = Nodes{} }
func (m *Nodes) String() string { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()    {}
func (*Nodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{4}
}
func (m *Nodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nodes.Unmarshal(m, b)
}
func (m *Nodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nodes.Marshal(b, m, deterministic)
}
func (dst *Nodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nodes.Merge(dst, src)
}
func (m *Nodes) XXX_Size() int {
	return xxx_messageInfo_Nodes.Size(m)
}
func (m *Nodes) XXX_DiscardUnknown() {
	xxx_messageInfo_Nodes.DiscardUnknown(m)
}

var xxx_messageInfo_Nodes proto.InternalMessageInfo

func (m *Nodes) GetNodes() []*PbftNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type PbftBlockHeader struct {
	Number               int64    `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
	GasLimit             int64    `protobuf:"varint,2,opt,name=GasLimit,proto3" json:"GasLimit,omitempty"`
	GasUsed              int64    `protobuf:"varint,3,opt,name=GasUsed,proto3" json:"GasUsed,omitempty"`
	Time                 int64    `protobuf:"varint,4,opt,name=Time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbftBlockHeader) Reset()         { *m = PbftBlockHeader{} }
func (m *PbftBlockHeader) String() string { return proto.CompactTextString(m) }
func (*PbftBlockHeader) ProtoMessage()    {}
func (*PbftBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{5}
}
func (m *PbftBlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftBlockHeader.Unmarshal(m, b)
}
func (m *PbftBlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftBlockHeader.Marshal(b, m, deterministic)
}
func (dst *PbftBlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftBlockHeader.Merge(dst, src)
}
func (m *PbftBlockHeader) XXX_Size() int {
	return xxx_messageInfo_PbftBlockHeader.Size(m)
}
func (m *PbftBlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftBlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_PbftBlockHeader proto.InternalMessageInfo

func (m *PbftBlockHeader) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *PbftBlockHeader) GetGasLimit() int64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *PbftBlockHeader) GetGasUsed() int64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *PbftBlockHeader) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type TxnData struct {
	AccountNonce         uint64   `protobuf:"varint,1,opt,name=AccountNonce,proto3" json:"AccountNonce,omitempty"`
	Price                int64    `protobuf:"varint,2,opt,name=Price,proto3" json:"Price,omitempty"`
	GasLimit             int64    `protobuf:"varint,3,opt,name=GasLimit,proto3" json:"GasLimit,omitempty"`
	Recipient            []byte   `protobuf:"bytes,4,opt,name=Recipient,proto3" json:"Recipient,omitempty"`
	Amount               int64    `protobuf:"varint,5,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Payload              []byte   `protobuf:"bytes,6,opt,name=Payload,proto3" json:"Payload,omitempty"`
	V                    int64    `protobuf:"varint,7,opt,name=V,proto3" json:"V,omitempty"`
	R                    int64    `protobuf:"varint,8,opt,name=R,proto3" json:"R,omitempty"`
	S                    int64    `protobuf:"varint,9,opt,name=S,proto3" json:"S,omitempty"`
	Hash                 []byte   `protobuf:"bytes,10,opt,name=Hash,proto3" json:"Hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxnData) Reset()         { *m = TxnData{} }
func (m *TxnData) String() string { return proto.CompactTextString(m) }
func (*TxnData) ProtoMessage()    {}
func (*TxnData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{6}
}
func (m *TxnData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxnData.Unmarshal(m, b)
}
func (m *TxnData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxnData.Marshal(b, m, deterministic)
}
func (dst *TxnData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxnData.Merge(dst, src)
}
func (m *TxnData) XXX_Size() int {
	return xxx_messageInfo_TxnData.Size(m)
}
func (m *TxnData) XXX_DiscardUnknown() {
	xxx_messageInfo_TxnData.DiscardUnknown(m)
}

var xxx_messageInfo_TxnData proto.InternalMessageInfo

func (m *TxnData) GetAccountNonce() uint64 {
	if m != nil {
		return m.AccountNonce
	}
	return 0
}

func (m *TxnData) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TxnData) GetGasLimit() int64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *TxnData) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *TxnData) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TxnData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TxnData) GetV() int64 {
	if m != nil {
		return m.V
	}
	return 0
}

func (m *TxnData) GetR() int64 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *TxnData) GetS() int64 {
	if m != nil {
		return m.S
	}
	return 0
}

func (m *TxnData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Transaction struct {
	Data                 *TxnData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{7}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetData() *TxnData {
	if m != nil {
		return m.Data
	}
	return nil
}

type Transactions struct {
	Txns                 []*Transaction `protobuf:"bytes,1,rep,name=txns,proto3" json:"txns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Transactions) Reset()         { *m = Transactions{} }
func (m *Transactions) String() string { return proto.CompactTextString(m) }
func (*Transactions) ProtoMessage()    {}
func (*Transactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{8}
}
func (m *Transactions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transactions.Unmarshal(m, b)
}
func (m *Transactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transactions.Marshal(b, m, deterministic)
}
func (dst *Transactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transactions.Merge(dst, src)
}
func (m *Transactions) XXX_Size() int {
	return xxx_messageInfo_Transactions.Size(m)
}
func (m *Transactions) XXX_DiscardUnknown() {
	xxx_messageInfo_Transactions.DiscardUnknown(m)
}

var xxx_messageInfo_Transactions proto.InternalMessageInfo

func (m *Transactions) GetTxns() []*Transaction {
	if m != nil {
		return m.Txns
	}
	return nil
}

type PbftBlock struct {
	Header               *PbftBlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Txns                 *Transactions    `protobuf:"bytes,2,opt,name=txns,proto3" json:"txns,omitempty"`
	Signs                []string         `protobuf:"bytes,3,rep,name=signs,proto3" json:"signs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PbftBlock) Reset()         { *m = PbftBlock{} }
func (m *PbftBlock) String() string { return proto.CompactTextString(m) }
func (*PbftBlock) ProtoMessage()    {}
func (*PbftBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{9}
}
func (m *PbftBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftBlock.Unmarshal(m, b)
}
func (m *PbftBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftBlock.Marshal(b, m, deterministic)
}
func (dst *PbftBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftBlock.Merge(dst, src)
}
func (m *PbftBlock) XXX_Size() int {
	return xxx_messageInfo_PbftBlock.Size(m)
}
func (m *PbftBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftBlock.DiscardUnknown(m)
}

var xxx_messageInfo_PbftBlock proto.InternalMessageInfo

func (m *PbftBlock) GetHeader() *PbftBlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PbftBlock) GetTxns() *Transactions {
	if m != nil {
		return m.Txns
	}
	return nil
}

func (m *PbftBlock) GetSigns() []string {
	if m != nil {
		return m.Signs
	}
	return nil
}

type GenericResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResp) Reset()         { *m = GenericResp{} }
func (m *GenericResp) String() string { return proto.CompactTextString(m) }
func (*GenericResp) ProtoMessage()    {}
func (*GenericResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_830dceaa21618eef, []int{10}
}
func (m *GenericResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResp.Unmarshal(m, b)
}
func (m *GenericResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResp.Marshal(b, m, deterministic)
}
func (dst *GenericResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResp.Merge(dst, src)
}
func (m *GenericResp) XXX_Size() int {
	return xxx_messageInfo_GenericResp.Size(m)
}
func (m *GenericResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResp.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResp proto.InternalMessageInfo

func (m *GenericResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "fastchain.Request")
	proto.RegisterType((*Request_Inner)(nil), "fastchain.Request.Inner")
	proto.RegisterType((*Request_MsgSignature)(nil), "fastchain.Request.MsgSignature")
	proto.RegisterType((*CheckLeaderReq)(nil), "fastchain.CheckLeaderReq")
	proto.RegisterType((*CheckLeaderResp)(nil), "fastchain.CheckLeaderResp")
	proto.RegisterType((*PbftNode)(nil), "fastchain.PbftNode")
	proto.RegisterType((*Nodes)(nil), "fastchain.Nodes")
	proto.RegisterType((*PbftBlockHeader)(nil), "fastchain.PbftBlockHeader")
	proto.RegisterType((*TxnData)(nil), "fastchain.TxnData")
	proto.RegisterType((*Transaction)(nil), "fastchain.Transaction")
	proto.RegisterType((*Transactions)(nil), "fastchain.Transactions")
	proto.RegisterType((*PbftBlock)(nil), "fastchain.PbftBlock")
	proto.RegisterType((*GenericResp)(nil), "fastchain.GenericResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FastChainClient is the client API for FastChain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FastChainClient interface {
	// RPC service that responds whether the node is the leader
	CheckLeader(ctx context.Context, in *CheckLeaderReq, opts ...grpc.CallOption) (*CheckLeaderResp, error)
	// Send new transaction to presumed leader node
	NewTxnRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GenericResp, error)
}

type fastChainClient struct {
	cc *grpc.ClientConn
}

func NewFastChainClient(cc *grpc.ClientConn) FastChainClient {
	return &fastChainClient{cc}
}

func (c *fastChainClient) CheckLeader(ctx context.Context, in *CheckLeaderReq, opts ...grpc.CallOption) (*CheckLeaderResp, error) {
	out := new(CheckLeaderResp)
	err := c.cc.Invoke(ctx, "/fastchain.FastChain/CheckLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fastChainClient) NewTxnRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GenericResp, error) {
	out := new(GenericResp)
	err := c.cc.Invoke(ctx, "/fastchain.FastChain/NewTxnRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FastChainServer is the server API for FastChain service.
type FastChainServer interface {
	// RPC service that responds whether the node is the leader
	CheckLeader(context.Context, *CheckLeaderReq) (*CheckLeaderResp, error)
	// Send new transaction to presumed leader node
	NewTxnRequest(context.Context, *Request) (*GenericResp, error)
}

func RegisterFastChainServer(s *grpc.Server, srv FastChainServer) {
	s.RegisterService(&_FastChain_serviceDesc, srv)
}

func _FastChain_CheckLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckLeaderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FastChainServer).CheckLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastchain.FastChain/CheckLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FastChainServer).CheckLeader(ctx, req.(*CheckLeaderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FastChain_NewTxnRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FastChainServer).NewTxnRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastchain.FastChain/NewTxnRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FastChainServer).NewTxnRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _FastChain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fastchain.FastChain",
	HandlerType: (*FastChainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckLeader",
			Handler:    _FastChain_CheckLeader_Handler,
		},
		{
			MethodName: "NewTxnRequest",
			Handler:    _FastChain_NewTxnRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fastchain.proto",
}

func init() { proto.RegisterFile("fastchain.proto", fileDescriptor_fastchain_830dceaa21618eef) }

var fileDescriptor_fastchain_830dceaa21618eef = []byte{
	// 690 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0xc6, 0x71, 0x9c, 0xc4, 0x93, 0xfc, 0x00, 0xed, 0xaf, 0xa2, 0x6e, 0x84, 0x44, 0xe4, 0x43,
	0x95, 0x82, 0x14, 0xa9, 0xa9, 0x7a, 0xa9, 0xd4, 0x03, 0xa5, 0x02, 0x2a, 0xd1, 0x28, 0x5a, 0x52,
	0xee, 0x1b, 0x7b, 0x71, 0x56, 0xe0, 0xb5, 0xf1, 0x6e, 0xf8, 0x73, 0x6c, 0x1f, 0xa1, 0xcf, 0xd4,
	0xd7, 0xe9, 0x3b, 0x54, 0xb3, 0x6b, 0x13, 0x87, 0xd2, 0xdb, 0x7c, 0xb3, 0x33, 0xf3, 0x7d, 0xbb,
	0x33, 0xb3, 0xb0, 0x75, 0xc9, 0x94, 0x8e, 0x16, 0x4c, 0xc8, 0x51, 0x5e, 0x64, 0x3a, 0x23, 0xfe,
	0xa3, 0x23, 0xfc, 0xd5, 0x80, 0x36, 0xe5, 0x37, 0x4b, 0xae, 0x34, 0x19, 0x81, 0x27, 0xa4, 0xe4,
	0x45, 0xe0, 0x0c, 0x9c, 0x61, 0x77, 0x1c, 0x8c, 0x56, 0x79, 0x65, 0xc8, 0xe8, 0x0b, 0x9e, 0x53,
	0x1b, 0x46, 0xb6, 0xc1, 0x8d, 0x45, 0x12, 0x34, 0x06, 0xce, 0xb0, 0x47, 0xd1, 0x24, 0x6f, 0xc1,
	0x55, 0x22, 0x09, 0x5c, 0x93, 0xbf, 0xf7, 0x4c, 0xfe, 0x57, 0x95, 0x9c, 0x8b, 0x44, 0x32, 0xbd,
	0x2c, 0x38, 0xc5, 0x58, 0xf2, 0x02, 0xbc, 0x6c, 0xa9, 0x79, 0x11, 0x34, 0x4d, 0x19, 0x0b, 0xfa,
	0xdf, 0x1d, 0xf0, 0x0c, 0x17, 0xd9, 0x84, 0x86, 0x88, 0x8d, 0x22, 0x8f, 0x36, 0x44, 0x8c, 0xa4,
	0x8a, 0xdf, 0x18, 0x52, 0x8f, 0xa2, 0x49, 0x08, 0x34, 0x6f, 0x05, 0xbf, 0x33, 0xac, 0x1e, 0x35,
	0x36, 0xfa, 0xf4, 0x43, 0xce, 0x4d, 0x51, 0x8f, 0x1a, 0x1b, 0x33, 0x53, 0x95, 0x04, 0x9e, 0x95,
	0x9b, 0xaa, 0x84, 0xec, 0x82, 0xaf, 0x45, 0xca, 0x95, 0x66, 0x69, 0x1e, 0xb4, 0x06, 0xce, 0xd0,
	0xa5, 0x2b, 0x47, 0x7f, 0x1f, 0x7a, 0x75, 0xb9, 0xa4, 0x07, 0x8e, 0x7d, 0x1a, 0x97, 0x3a, 0x05,
	0x22, 0x65, 0x54, 0xb8, 0xd4, 0x51, 0xe1, 0x36, 0x6c, 0x1e, 0x2d, 0x78, 0x74, 0x75, 0xc6, 0x59,
	0xcc, 0x0b, 0xca, 0x6f, 0xc2, 0x03, 0xd8, 0x5a, 0xf3, 0xa8, 0x9c, 0x04, 0xd0, 0x4e, 0xb9, 0x52,
	0x2c, 0xe1, 0xa6, 0x4c, 0x87, 0x56, 0x30, 0x9c, 0x42, 0x67, 0x3a, 0xbf, 0xd4, 0x93, 0x2c, 0xe6,
	0x28, 0x9d, 0xc5, 0xb1, 0x65, 0xf2, 0xa9, 0xb1, 0xc9, 0x0e, 0xb4, 0xf2, 0xe5, 0xfc, 0x8a, 0x3f,
	0x18, 0x46, 0x9f, 0x96, 0x08, 0x2b, 0xe6, 0x85, 0xb8, 0xc5, 0x03, 0xd7, 0x1c, 0x54, 0x30, 0x1c,
	0x83, 0x87, 0xd5, 0x14, 0x79, 0x03, 0x9e, 0x44, 0x23, 0x70, 0x06, 0xee, 0xb0, 0x3b, 0xfe, 0xbf,
	0xd6, 0x94, 0x8a, 0x92, 0xda, 0x88, 0x50, 0xc1, 0x16, 0xba, 0x3e, 0x5d, 0x67, 0xd1, 0xd5, 0xa9,
	0x91, 0x8d, 0xc4, 0x93, 0x65, 0x3a, 0xe7, 0xd5, 0xc5, 0x4b, 0x44, 0xfa, 0xd0, 0x39, 0x61, 0xea,
	0x4c, 0xa4, 0x42, 0x97, 0x8f, 0xf0, 0x88, 0x51, 0xd4, 0x09, 0x53, 0xdf, 0x14, 0x8f, 0x8d, 0x28,
	0x97, 0x56, 0x10, 0xaf, 0x36, 0x13, 0xa9, 0xed, 0x8a, 0x4b, 0x8d, 0x1d, 0xfe, 0x76, 0xa0, 0x3d,
	0xbb, 0x97, 0x9f, 0x99, 0x66, 0x24, 0x84, 0xde, 0x61, 0x14, 0x65, 0x4b, 0xa9, 0x27, 0x99, 0x8c,
	0xec, 0x2b, 0x35, 0xe9, 0x9a, 0x0f, 0xe7, 0x65, 0x5a, 0x88, 0x88, 0x97, 0xb4, 0x16, 0xac, 0xe9,
	0x71, 0x9f, 0xe8, 0xd9, 0x05, 0x9f, 0xf2, 0x48, 0xe4, 0x82, 0x4b, 0x5d, 0x4e, 0xd9, 0xca, 0x81,
	0x37, 0x3c, 0x4c, 0xb1, 0xbc, 0x19, 0x0c, 0x97, 0x96, 0x08, 0x6f, 0x31, 0x65, 0x0f, 0xd7, 0x19,
	0x8b, 0xcd, 0x64, 0xf4, 0x68, 0x05, 0xb1, 0xf3, 0x17, 0x41, 0xdb, 0x76, 0xfe, 0x02, 0x11, 0x0d,
	0x3a, 0x16, 0x51, 0x44, 0xe7, 0x81, 0x6f, 0xd1, 0x39, 0xde, 0xf7, 0x94, 0xa9, 0x45, 0x00, 0xa6,
	0x80, 0xb1, 0xc3, 0xf7, 0xd0, 0x9d, 0x15, 0x4c, 0x2a, 0x16, 0x69, 0x91, 0x49, 0xf2, 0x1a, 0x9a,
	0x31, 0xd3, 0xac, 0x5c, 0x39, 0x52, 0xeb, 0x4e, 0xf9, 0x28, 0xd4, 0x9c, 0x87, 0x1f, 0xa0, 0x57,
	0x4b, 0x53, 0x64, 0x1f, 0x9a, 0xfa, 0x5e, 0x56, 0x5d, 0xdd, 0xa9, 0xe7, 0xad, 0xc2, 0xa8, 0x89,
	0x09, 0x7f, 0x38, 0xe0, 0x3f, 0x36, 0x96, 0x8c, 0xa1, 0xb5, 0x30, 0xcd, 0x2d, 0x39, 0xfb, 0x4f,
	0x26, 0xa2, 0xd6, 0x7e, 0x5a, 0x46, 0x92, 0x83, 0x92, 0xad, 0x61, 0x32, 0x5e, 0x3e, 0xcf, 0xa6,
	0x2c, 0x1d, 0x76, 0x48, 0x89, 0x44, 0xaa, 0xc0, 0x1d, 0xb8, 0x43, 0x9f, 0x5a, 0x10, 0xee, 0x41,
	0xf7, 0x84, 0x4b, 0x5e, 0x88, 0xc8, 0xec, 0x42, 0xb9, 0x8c, 0x76, 0xc8, 0xd1, 0x1c, 0xff, 0x74,
	0xc0, 0x3f, 0x66, 0x4a, 0x1f, 0x61, 0x5d, 0x72, 0x0c, 0xdd, 0xda, 0xfa, 0x90, 0x57, 0x35, 0xca,
	0xf5, 0x45, 0xeb, 0xf7, 0xff, 0x75, 0xa4, 0xf2, 0x70, 0x83, 0x7c, 0x84, 0xff, 0x26, 0xfc, 0x6e,
	0x76, 0x2f, 0xab, 0x4f, 0x8e, 0xfc, 0xfd, 0x2b, 0xf5, 0xeb, 0xcf, 0x57, 0x13, 0x19, 0x6e, 0xcc,
	0x5b, 0xe6, 0xc3, 0x7c, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x47, 0xfb, 0xa4, 0x43, 0x05,
	0x00, 0x00,
}

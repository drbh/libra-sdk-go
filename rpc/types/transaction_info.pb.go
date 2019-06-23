// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/transaction_info.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// `TransactionInfo` is the object we store in the transaction accumulator. It
// consists of the transaction as well as the execution result of this
// transaction. This are later returned to the client so that a client can
// validate the tree
type TransactionInfo struct {
	// Hash of the signed transaction that is stored
	SignedTransactionHash []byte `protobuf:"bytes,1,opt,name=signed_transaction_hash,json=signedTransactionHash,proto3" json:"signed_transaction_hash,omitempty"`
	// The root hash of Sparse Merkle Tree describing the world state at the end
	// of this transaction
	StateRootHash []byte `protobuf:"bytes,2,opt,name=state_root_hash,json=stateRootHash,proto3" json:"state_root_hash,omitempty"`
	// The root hash of Merkle Accumulator storing all events emitted during this
	// transaction.
	EventRootHash []byte `protobuf:"bytes,3,opt,name=event_root_hash,json=eventRootHash,proto3" json:"event_root_hash,omitempty"`
	// The amount of gas used by this transaction.
	GasUsed              uint64   `protobuf:"varint,4,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionInfo) Reset()         { *m = TransactionInfo{} }
func (m *TransactionInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionInfo) ProtoMessage()    {}
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_af679f6db9619bd9, []int{0}
}

func (m *TransactionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionInfo.Unmarshal(m, b)
}
func (m *TransactionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionInfo.Marshal(b, m, deterministic)
}
func (m *TransactionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInfo.Merge(m, src)
}
func (m *TransactionInfo) XXX_Size() int {
	return xxx_messageInfo_TransactionInfo.Size(m)
}
func (m *TransactionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInfo proto.InternalMessageInfo

func (m *TransactionInfo) GetSignedTransactionHash() []byte {
	if m != nil {
		return m.SignedTransactionHash
	}
	return nil
}

func (m *TransactionInfo) GetStateRootHash() []byte {
	if m != nil {
		return m.StateRootHash
	}
	return nil
}

func (m *TransactionInfo) GetEventRootHash() []byte {
	if m != nil {
		return m.EventRootHash
	}
	return nil
}

func (m *TransactionInfo) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func init() {
	proto.RegisterType((*TransactionInfo)(nil), "types.TransactionInfo")
}

func init() { proto.RegisterFile("types/transaction_info.proto", fileDescriptor_af679f6db9619bd9) }

var fileDescriptor_af679f6db9619bd9 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0xcf, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x05, 0x60, 0xaa, 0xd7, 0x1f, 0x82, 0x72, 0xa1, 0x20, 0x56, 0x70, 0x21, 0x2e, 0xc4, 0x4d,
	0x1b, 0x41, 0xf0, 0x01, 0x5c, 0xe9, 0xb6, 0xe8, 0xc6, 0x4d, 0x48, 0xdb, 0x98, 0x0c, 0xc6, 0x4c,
	0xc8, 0x4c, 0x05, 0xdf, 0xcb, 0x07, 0x34, 0xa4, 0xe8, 0xed, 0x72, 0xce, 0xf9, 0x66, 0x71, 0xc4,
	0x25, 0x7f, 0x47, 0x43, 0x92, 0x93, 0x0e, 0xa4, 0x47, 0x06, 0x0c, 0x0a, 0xc2, 0x3b, 0x76, 0x31,
	0x21, 0x63, 0x7d, 0x50, 0xda, 0xeb, 0x9f, 0x4a, 0x6c, 0x5f, 0x76, 0xe2, 0x39, 0x83, 0xfa, 0x41,
	0x9c, 0x13, 0xd8, 0x60, 0x26, 0xb5, 0xfe, 0x75, 0x9a, 0x5c, 0x53, 0x5d, 0x55, 0xb7, 0x27, 0xfd,
	0xd9, 0x52, 0xaf, 0xfe, 0x9e, 0x72, 0x59, 0xdf, 0x88, 0x2d, 0xb1, 0x66, 0xa3, 0x12, 0x22, 0x2f,
	0x7e, 0xaf, 0xf8, 0xd3, 0x12, 0xf7, 0x39, 0xfd, 0x73, 0xe6, 0xcb, 0x04, 0x5e, 0xb9, 0xfd, 0xc5,
	0x95, 0xf8, 0xdf, 0x5d, 0x88, 0x63, 0xab, 0x49, 0xcd, 0x64, 0xa6, 0x66, 0x93, 0xc1, 0xa6, 0x3f,
	0xca, 0xf7, 0x6b, 0x3e, 0x1f, 0xef, 0xde, 0x3a, 0x0b, 0xec, 0xe6, 0xa1, 0x1b, 0xf1, 0x53, 0x46,
	0x07, 0x1e, 0x62, 0xb4, 0xe0, 0xbd, 0x91, 0x1e, 0x86, 0xa4, 0x5b, 0x9a, 0x3e, 0x5a, 0x8b, 0x32,
	0xc5, 0x51, 0x96, 0xa1, 0xc3, 0x61, 0x99, 0x7d, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x13,
	0xdd, 0x29, 0x16, 0x01, 0x00, 0x00,
}
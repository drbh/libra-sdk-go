// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/validator_change.proto

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

// This is used to prove validator changes.  When a validator is changing, it
// triggers an event on /validator_change_account/events/sent.  To tell the
// client about validator changes, we query
// /validator_change_account/events/sent to get all versions that contain
// validator changes after the version that we are trying to update from. For
// each of these versions, the old validator set would have signed the ledger
// info at that version.  The client needs this as well as the event results +
// proof.  The client can then verify that these events were under the current
// tree and that the changes were signed by the old validators (and that the
// events correctly show which validators are the new validators).
//
// This message represents a single validator change event and the proof that
// corresponds to it
type ValidatorChangeEventWithProof struct {
	LedgerInfoWithSigs   *LedgerInfoWithSignatures `protobuf:"bytes,1,opt,name=ledger_info_with_sigs,json=ledgerInfoWithSigs,proto3" json:"ledger_info_with_sigs,omitempty"`
	EventWithProof       *EventWithProof           `protobuf:"bytes,2,opt,name=event_with_proof,json=eventWithProof,proto3" json:"event_with_proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ValidatorChangeEventWithProof) Reset()         { *m = ValidatorChangeEventWithProof{} }
func (m *ValidatorChangeEventWithProof) String() string { return proto.CompactTextString(m) }
func (*ValidatorChangeEventWithProof) ProtoMessage()    {}
func (*ValidatorChangeEventWithProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_366a4adb25281530, []int{0}
}

func (m *ValidatorChangeEventWithProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorChangeEventWithProof.Unmarshal(m, b)
}
func (m *ValidatorChangeEventWithProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorChangeEventWithProof.Marshal(b, m, deterministic)
}
func (m *ValidatorChangeEventWithProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorChangeEventWithProof.Merge(m, src)
}
func (m *ValidatorChangeEventWithProof) XXX_Size() int {
	return xxx_messageInfo_ValidatorChangeEventWithProof.Size(m)
}
func (m *ValidatorChangeEventWithProof) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorChangeEventWithProof.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorChangeEventWithProof proto.InternalMessageInfo

func (m *ValidatorChangeEventWithProof) GetLedgerInfoWithSigs() *LedgerInfoWithSignatures {
	if m != nil {
		return m.LedgerInfoWithSigs
	}
	return nil
}

func (m *ValidatorChangeEventWithProof) GetEventWithProof() *EventWithProof {
	if m != nil {
		return m.EventWithProof
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidatorChangeEventWithProof)(nil), "types.ValidatorChangeEventWithProof")
}

func init() { proto.RegisterFile("types/validator_change.proto", fileDescriptor_366a4adb25281530) }

var fileDescriptor_366a4adb25281530 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0xa0, 0x87, 0x0a, 0x22, 0x81, 0xc5, 0x65, 0x51, 0x14, 0x4f, 0x5e, 0x36, 0x11,
	0x7d, 0x00, 0x41, 0xf1, 0x20, 0x78, 0x90, 0x0a, 0x0a, 0x5e, 0x4a, 0xda, 0xce, 0xa6, 0x83, 0x31,
	0x13, 0x92, 0xec, 0x8a, 0xef, 0xe5, 0x03, 0xda, 0x4e, 0x14, 0xec, 0x1e, 0x27, 0xdf, 0xff, 0x7f,
	0x99, 0x29, 0x8f, 0xd3, 0x97, 0x87, 0xa8, 0x36, 0xda, 0x62, 0xa7, 0x13, 0x85, 0xba, 0xed, 0xb5,
	0x33, 0x20, 0x7d, 0xa0, 0x44, 0x62, 0x97, 0xe9, 0x42, 0xe4, 0x10, 0x6c, 0xc0, 0xa5, 0x98, 0xd1,
	0xe2, 0x28, 0xbf, 0x59, 0xe8, 0x0c, 0x84, 0x1a, 0xdd, 0x8a, 0x32, 0x38, 0xff, 0x2e, 0xca, 0x93,
	0x97, 0x3f, 0xdd, 0x1d, 0xdb, 0xee, 0xc7, 0xe2, 0x2b, 0xa6, 0xfe, 0x29, 0x10, 0xad, 0x44, 0x55,
	0xce, 0xfe, 0xd5, 0xea, 0xcf, 0x01, 0xd4, 0x11, 0x4d, 0x9c, 0x17, 0x67, 0xc5, 0xc5, 0xfe, 0xd5,
	0xa9, 0x64, 0xb5, 0x7c, 0xe4, 0xcc, 0xc3, 0x10, 0x19, 0xab, 0xcf, 0x68, 0x9c, 0x4e, 0xeb, 0x00,
	0xb1, 0x12, 0x76, 0x9b, 0x44, 0x71, 0x53, 0x1e, 0xf2, 0x7a, 0xd9, 0xe6, 0xc7, 0x7f, 0xe6, 0x3b,
	0xac, 0x9b, 0xfd, 0xea, 0xa6, 0x4b, 0x54, 0x07, 0x30, 0x99, 0x6f, 0x2f, 0xdf, 0xa4, 0x19, 0x86,
	0x75, 0x23, 0x5b, 0xfa, 0x50, 0xbe, 0x47, 0x8b, 0xde, 0x1b, 0xb4, 0x16, 0x94, 0xc5, 0x26, 0xe8,
	0x65, 0xec, 0xde, 0x97, 0x86, 0x54, 0xf0, 0xad, 0x62, 0x61, 0xb3, 0xc7, 0xf7, 0x5e, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xc0, 0x49, 0xfb, 0x62, 0x43, 0x01, 0x00, 0x00,
}
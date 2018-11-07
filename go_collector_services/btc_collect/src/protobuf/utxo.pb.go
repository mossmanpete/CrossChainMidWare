// Code generated by protoc-gen-go. DO NOT EDIT.
// source: utxo.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UTXOObject struct {
	Address              *string  `protobuf:"bytes,1,req,name=address" json:"address,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	ScriptPubKey         *string  `protobuf:"bytes,3,req,name=scriptPubKey" json:"scriptPubKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UTXOObject) Reset()         { *m = UTXOObject{} }
func (m *UTXOObject) String() string { return proto.CompactTextString(m) }
func (*UTXOObject) ProtoMessage()    {}
func (*UTXOObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_utxo_a9884cfc0f8a52b8, []int{0}
}
func (m *UTXOObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UTXOObject.Unmarshal(m, b)
}
func (m *UTXOObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UTXOObject.Marshal(b, m, deterministic)
}
func (dst *UTXOObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UTXOObject.Merge(dst, src)
}
func (m *UTXOObject) XXX_Size() int {
	return xxx_messageInfo_UTXOObject.Size(m)
}
func (m *UTXOObject) XXX_DiscardUnknown() {
	xxx_messageInfo_UTXOObject.DiscardUnknown(m)
}

var xxx_messageInfo_UTXOObject proto.InternalMessageInfo

func (m *UTXOObject) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *UTXOObject) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *UTXOObject) GetScriptPubKey() string {
	if m != nil && m.ScriptPubKey != nil {
		return *m.ScriptPubKey
	}
	return ""
}

func init() {
	proto.RegisterType((*UTXOObject)(nil), "tutorial.UTXOObject")
}

func init() { proto.RegisterFile("utxo.proto", fileDescriptor_utxo_a9884cfc0f8a52b8) }

var fileDescriptor_utxo_a9884cfc0f8a52b8 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0xa9, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x29, 0x2d, 0xc9, 0x2f, 0xca, 0x4c, 0xcc,
	0x51, 0x4a, 0xe0, 0xe2, 0x0a, 0x0d, 0x89, 0xf0, 0xf7, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0x11, 0x92,
	0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c,
	0x82, 0x71, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xc0, 0xe2, 0x10,
	0x8e, 0x90, 0x12, 0x17, 0x4f, 0x71, 0x72, 0x51, 0x66, 0x41, 0x49, 0x40, 0x69, 0x92, 0x77, 0x6a,
	0xa5, 0x04, 0x33, 0x58, 0x12, 0x45, 0x0c, 0x10, 0x00, 0x00, 0xff, 0xff, 0x12, 0x47, 0x3d, 0x50,
	0x78, 0x00, 0x00, 0x00,
}

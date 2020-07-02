// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iavl/types.proto

package iavl

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ProofOpValue is a value proof, used internally to encode and decode ValueOp.
type ProofOpValue struct {
	Proof *ProofOpRange `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *ProofOpValue) Reset()         { *m = ProofOpValue{} }
func (m *ProofOpValue) String() string { return proto.CompactTextString(m) }
func (*ProofOpValue) ProtoMessage()    {}
func (*ProofOpValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ef37c124502d49e, []int{0}
}
func (m *ProofOpValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofOpValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofOpValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofOpValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofOpValue.Merge(m, src)
}
func (m *ProofOpValue) XXX_Size() int {
	return m.Size()
}
func (m *ProofOpValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofOpValue.DiscardUnknown(m)
}

var xxx_messageInfo_ProofOpValue proto.InternalMessageInfo

func (m *ProofOpValue) GetProof() *ProofOpRange {
	if m != nil {
		return m.Proof
	}
	return nil
}

// ProofOpValue is a value proof, used internally to encode and decode ValueOp.
type ProofOpAbsence struct {
	Proof *ProofOpRange `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *ProofOpAbsence) Reset()         { *m = ProofOpAbsence{} }
func (m *ProofOpAbsence) String() string { return proto.CompactTextString(m) }
func (*ProofOpAbsence) ProtoMessage()    {}
func (*ProofOpAbsence) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ef37c124502d49e, []int{1}
}
func (m *ProofOpAbsence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofOpAbsence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofOpAbsence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofOpAbsence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofOpAbsence.Merge(m, src)
}
func (m *ProofOpAbsence) XXX_Size() int {
	return m.Size()
}
func (m *ProofOpAbsence) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofOpAbsence.DiscardUnknown(m)
}

var xxx_messageInfo_ProofOpAbsence proto.InternalMessageInfo

func (m *ProofOpAbsence) GetProof() *ProofOpRange {
	if m != nil {
		return m.Proof
	}
	return nil
}

// ProofOpRange is used internally to encode and decode ValueOp and
// AbsenceOp proof operations.
type ProofOpRange struct {
	LeftPath   []*ProofOpInner `protobuf:"bytes,1,rep,name=left_path,json=leftPath,proto3" json:"left_path,omitempty"`
	InnerNodes []*ProofOpPath  `protobuf:"bytes,2,rep,name=inner_nodes,json=innerNodes,proto3" json:"inner_nodes,omitempty"`
	Leaves     []*ProofOpLeaf  `protobuf:"bytes,3,rep,name=leaves,proto3" json:"leaves,omitempty"`
}

func (m *ProofOpRange) Reset()         { *m = ProofOpRange{} }
func (m *ProofOpRange) String() string { return proto.CompactTextString(m) }
func (*ProofOpRange) ProtoMessage()    {}
func (*ProofOpRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ef37c124502d49e, []int{2}
}
func (m *ProofOpRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofOpRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofOpRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofOpRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofOpRange.Merge(m, src)
}
func (m *ProofOpRange) XXX_Size() int {
	return m.Size()
}
func (m *ProofOpRange) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofOpRange.DiscardUnknown(m)
}

var xxx_messageInfo_ProofOpRange proto.InternalMessageInfo

func (m *ProofOpRange) GetLeftPath() []*ProofOpInner {
	if m != nil {
		return m.LeftPath
	}
	return nil
}

func (m *ProofOpRange) GetInnerNodes() []*ProofOpPath {
	if m != nil {
		return m.InnerNodes
	}
	return nil
}

func (m *ProofOpRange) GetLeaves() []*ProofOpLeaf {
	if m != nil {
		return m.Leaves
	}
	return nil
}

// ProofOpPath is used internally to encode and decode leaf
// node paths for ProofOpRange.
type ProofOpPath struct {
	Inners []*ProofOpInner `protobuf:"bytes,1,rep,name=inners,proto3" json:"inners,omitempty"`
}

func (m *ProofOpPath) Reset()         { *m = ProofOpPath{} }
func (m *ProofOpPath) String() string { return proto.CompactTextString(m) }
func (*ProofOpPath) ProtoMessage()    {}
func (*ProofOpPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ef37c124502d49e, []int{3}
}
func (m *ProofOpPath) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofOpPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofOpPath.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofOpPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofOpPath.Merge(m, src)
}
func (m *ProofOpPath) XXX_Size() int {
	return m.Size()
}
func (m *ProofOpPath) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofOpPath.DiscardUnknown(m)
}

var xxx_messageInfo_ProofOpPath proto.InternalMessageInfo

func (m *ProofOpPath) GetInners() []*ProofOpInner {
	if m != nil {
		return m.Inners
	}
	return nil
}

// ProofOpInner is used internally to encode and decode inner nodes
// for ProofOpRange.
type ProofOpInner struct {
	Height  int32  `protobuf:"zigzag32,1,opt,name=height,proto3" json:"height,omitempty"`
	Size_   int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Version int64  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Left    []byte `protobuf:"bytes,4,opt,name=left,proto3" json:"left,omitempty"`
	Right   []byte `protobuf:"bytes,5,opt,name=right,proto3" json:"right,omitempty"`
}

func (m *ProofOpInner) Reset()         { *m = ProofOpInner{} }
func (m *ProofOpInner) String() string { return proto.CompactTextString(m) }
func (*ProofOpInner) ProtoMessage()    {}
func (*ProofOpInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ef37c124502d49e, []int{4}
}
func (m *ProofOpInner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofOpInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofOpInner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofOpInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofOpInner.Merge(m, src)
}
func (m *ProofOpInner) XXX_Size() int {
	return m.Size()
}
func (m *ProofOpInner) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofOpInner.DiscardUnknown(m)
}

var xxx_messageInfo_ProofOpInner proto.InternalMessageInfo

func (m *ProofOpInner) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ProofOpInner) GetSize_() int64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *ProofOpInner) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ProofOpInner) GetLeft() []byte {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *ProofOpInner) GetRight() []byte {
	if m != nil {
		return m.Right
	}
	return nil
}

// ProofOpLeaf is used internally to encode and decode leaf nodes
// for ProofOpRange.
type ProofOpLeaf struct {
	Key       []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ValueHash []byte `protobuf:"bytes,2,opt,name=value_hash,json=valueHash,proto3" json:"value_hash,omitempty"`
	Version   int64  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *ProofOpLeaf) Reset()         { *m = ProofOpLeaf{} }
func (m *ProofOpLeaf) String() string { return proto.CompactTextString(m) }
func (*ProofOpLeaf) ProtoMessage()    {}
func (*ProofOpLeaf) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ef37c124502d49e, []int{5}
}
func (m *ProofOpLeaf) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofOpLeaf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofOpLeaf.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofOpLeaf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofOpLeaf.Merge(m, src)
}
func (m *ProofOpLeaf) XXX_Size() int {
	return m.Size()
}
func (m *ProofOpLeaf) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofOpLeaf.DiscardUnknown(m)
}

var xxx_messageInfo_ProofOpLeaf proto.InternalMessageInfo

func (m *ProofOpLeaf) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ProofOpLeaf) GetValueHash() []byte {
	if m != nil {
		return m.ValueHash
	}
	return nil
}

func (m *ProofOpLeaf) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*ProofOpValue)(nil), "iavl.ProofOpValue")
	proto.RegisterType((*ProofOpAbsence)(nil), "iavl.ProofOpAbsence")
	proto.RegisterType((*ProofOpRange)(nil), "iavl.ProofOpRange")
	proto.RegisterType((*ProofOpPath)(nil), "iavl.ProofOpPath")
	proto.RegisterType((*ProofOpInner)(nil), "iavl.ProofOpInner")
	proto.RegisterType((*ProofOpLeaf)(nil), "iavl.ProofOpLeaf")
}

func init() { proto.RegisterFile("iavl/types.proto", fileDescriptor_7ef37c124502d49e) }

var fileDescriptor_7ef37c124502d49e = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x80, 0x3b, 0x4d, 0x9b, 0x7b, 0x7b, 0x1a, 0x2e, 0xed, 0x70, 0x91, 0xd9, 0x18, 0x42, 0x56,
	0xd1, 0x45, 0x0b, 0x75, 0xa3, 0xee, 0x74, 0xa5, 0x20, 0x5a, 0x66, 0x21, 0xe2, 0xa6, 0x4c, 0xf5,
	0xb4, 0x09, 0x86, 0x24, 0x64, 0x62, 0xa0, 0x2e, 0x7c, 0x06, 0x9f, 0xc0, 0xe7, 0x71, 0xd9, 0xa5,
	0x4b, 0x69, 0x5f, 0x44, 0xe6, 0x34, 0x16, 0xc5, 0x1f, 0x70, 0x77, 0xce, 0xc7, 0x77, 0xfe, 0x92,
	0x81, 0x4e, 0xa4, 0xca, 0xb8, 0x5f, 0xcc, 0x32, 0xd4, 0xbd, 0x2c, 0x4f, 0x8b, 0x94, 0x37, 0x0c,
	0xf1, 0x77, 0xc1, 0x19, 0xe6, 0x69, 0x3a, 0x39, 0xcb, 0xce, 0x55, 0x7c, 0x8b, 0x3c, 0x80, 0x66,
	0x66, 0x72, 0xc1, 0x3c, 0x16, 0xb4, 0x07, 0xbc, 0x67, 0xac, 0x5e, 0xa5, 0x48, 0x95, 0x4c, 0x51,
	0xae, 0x04, 0x7f, 0x1f, 0xfe, 0x55, 0xf8, 0x60, 0xac, 0x31, 0xb9, 0xfa, 0x4d, 0xed, 0x23, 0x5b,
	0x8f, 0x25, 0xce, 0xfb, 0xd0, 0x8a, 0x71, 0x52, 0x8c, 0x32, 0x55, 0x84, 0x82, 0x79, 0xd6, 0xa7,
	0xf2, 0xe3, 0x24, 0xc1, 0x5c, 0xfe, 0x35, 0xd2, 0x50, 0x15, 0x21, 0x1f, 0x40, 0x3b, 0x32, 0x68,
	0x94, 0xa4, 0xd7, 0xa8, 0x45, 0x9d, 0x4a, 0xba, 0x1f, 0x4a, 0x8c, 0x27, 0x81, 0xac, 0x53, 0x23,
	0xf1, 0x2d, 0xb0, 0x63, 0x54, 0x25, 0x6a, 0x61, 0x7d, 0xa1, 0x9f, 0xa0, 0x9a, 0xc8, 0x4a, 0xf0,
	0xf7, 0xa0, 0xfd, 0xae, 0x0b, 0xdf, 0x06, 0x9b, 0xfa, 0xe8, 0x1f, 0x76, 0xab, 0x0c, 0xff, 0x7e,
	0x7d, 0x1a, 0x71, 0xbe, 0x01, 0x76, 0x88, 0xd1, 0x34, 0x2c, 0xe8, 0xb3, 0x74, 0x65, 0x95, 0x71,
	0x0e, 0x0d, 0x1d, 0xdd, 0xa1, 0xa8, 0x7b, 0x2c, 0xb0, 0x24, 0xc5, 0x5c, 0xc0, 0x9f, 0x12, 0x73,
	0x1d, 0xa5, 0x89, 0xb0, 0x08, 0xbf, 0xa5, 0xc6, 0x36, 0xb7, 0x8b, 0x86, 0xc7, 0x02, 0x47, 0x52,
	0xcc, 0xff, 0x43, 0x33, 0xa7, 0xc6, 0x4d, 0x82, 0xab, 0xc4, 0xbf, 0x58, 0xaf, 0x6e, 0x2e, 0xe2,
	0x1d, 0xb0, 0x6e, 0x70, 0x46, 0xb3, 0x1d, 0x69, 0x42, 0xbe, 0x09, 0x50, 0x9a, 0x7f, 0x3d, 0x0a,
	0x95, 0x0e, 0x69, 0xbc, 0x23, 0x5b, 0x44, 0x8e, 0x94, 0x0e, 0xbf, 0xdf, 0xe1, 0xd0, 0x7d, 0x5a,
	0xb8, 0x6c, 0xbe, 0x70, 0xd9, 0xcb, 0xc2, 0x65, 0x0f, 0x4b, 0xb7, 0x36, 0x5f, 0xba, 0xb5, 0xe7,
	0xa5, 0x5b, 0xbb, 0xa4, 0xb7, 0x34, 0xb6, 0xe9, 0x61, 0xed, 0xbc, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0x8c, 0x73, 0x53, 0x6c, 0x02, 0x00, 0x00,
}

func (m *ProofOpValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofOpValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofOpValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProofOpAbsence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofOpAbsence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofOpAbsence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProofOpRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofOpRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofOpRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Leaves) > 0 {
		for iNdEx := len(m.Leaves) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Leaves[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.InnerNodes) > 0 {
		for iNdEx := len(m.InnerNodes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InnerNodes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.LeftPath) > 0 {
		for iNdEx := len(m.LeftPath) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LeftPath[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ProofOpPath) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofOpPath) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofOpPath) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Inners) > 0 {
		for iNdEx := len(m.Inners) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inners[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ProofOpInner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofOpInner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofOpInner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Right) > 0 {
		i -= len(m.Right)
		copy(dAtA[i:], m.Right)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Right)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Left) > 0 {
		i -= len(m.Left)
		copy(dAtA[i:], m.Left)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Left)))
		i--
		dAtA[i] = 0x22
	}
	if m.Version != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if m.Size_ != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64((uint32(m.Height)<<1)^uint32((m.Height>>31))))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProofOpLeaf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofOpLeaf) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofOpLeaf) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ValueHash) > 0 {
		i -= len(m.ValueHash)
		copy(dAtA[i:], m.ValueHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValueHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProofOpValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ProofOpAbsence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ProofOpRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LeftPath) > 0 {
		for _, e := range m.LeftPath {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.InnerNodes) > 0 {
		for _, e := range m.InnerNodes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Leaves) > 0 {
		for _, e := range m.Leaves {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ProofOpPath) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Inners) > 0 {
		for _, e := range m.Inners {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ProofOpInner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sozTypes(uint64(m.Height))
	}
	if m.Size_ != 0 {
		n += 1 + sovTypes(uint64(m.Size_))
	}
	if m.Version != 0 {
		n += 1 + sovTypes(uint64(m.Version))
	}
	l = len(m.Left)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Right)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ProofOpLeaf) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValueHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovTypes(uint64(m.Version))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProofOpValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProofOpValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofOpValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &ProofOpRange{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProofOpAbsence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProofOpAbsence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofOpAbsence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &ProofOpRange{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProofOpRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProofOpRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofOpRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeftPath", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LeftPath = append(m.LeftPath, &ProofOpInner{})
			if err := m.LeftPath[len(m.LeftPath)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InnerNodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InnerNodes = append(m.InnerNodes, &ProofOpPath{})
			if err := m.InnerNodes[len(m.InnerNodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leaves", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Leaves = append(m.Leaves, &ProofOpLeaf{})
			if err := m.Leaves[len(m.Leaves)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProofOpPath) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProofOpPath: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofOpPath: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inners = append(m.Inners, &ProofOpInner{})
			if err := m.Inners[len(m.Inners)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProofOpInner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProofOpInner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofOpInner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.Height = v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Left", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Left = append(m.Left[:0], dAtA[iNdEx:postIndex]...)
			if m.Left == nil {
				m.Left = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Right", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Right = append(m.Right[:0], dAtA[iNdEx:postIndex]...)
			if m.Right == nil {
				m.Right = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProofOpLeaf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProofOpLeaf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofOpLeaf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValueHash = append(m.ValueHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValueHash == nil {
				m.ValueHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)

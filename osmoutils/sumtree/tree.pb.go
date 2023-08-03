// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlin/sumtree/v1beta1/tree.proto

package sumtree

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Node struct {
	Children []*Child `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_31a1c5f55b935f78, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetChildren() []*Child {
	if m != nil {
		return m.Children
	}
	return nil
}

type Child struct {
	Index        []byte                                 `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Accumulation github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=accumulation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"accumulation"`
}

func (m *Child) Reset()         { *m = Child{} }
func (m *Child) String() string { return proto.CompactTextString(m) }
func (*Child) ProtoMessage()    {}
func (*Child) Descriptor() ([]byte, []int) {
	return fileDescriptor_31a1c5f55b935f78, []int{1}
}
func (m *Child) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Child) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Child.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Child) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Child.Merge(m, src)
}
func (m *Child) XXX_Size() int {
	return m.Size()
}
func (m *Child) XXX_DiscardUnknown() {
	xxx_messageInfo_Child.DiscardUnknown(m)
}

var xxx_messageInfo_Child proto.InternalMessageInfo

func (m *Child) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

type Leaf struct {
	Leaf *Child `protobuf:"bytes,1,opt,name=leaf,proto3" json:"leaf,omitempty"`
}

func (m *Leaf) Reset()         { *m = Leaf{} }
func (m *Leaf) String() string { return proto.CompactTextString(m) }
func (*Leaf) ProtoMessage()    {}
func (*Leaf) Descriptor() ([]byte, []int) {
	return fileDescriptor_31a1c5f55b935f78, []int{2}
}
func (m *Leaf) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Leaf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Leaf.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Leaf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Leaf.Merge(m, src)
}
func (m *Leaf) XXX_Size() int {
	return m.Size()
}
func (m *Leaf) XXX_DiscardUnknown() {
	xxx_messageInfo_Leaf.DiscardUnknown(m)
}

var xxx_messageInfo_Leaf proto.InternalMessageInfo

func (m *Leaf) GetLeaf() *Child {
	if m != nil {
		return m.Leaf
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "merlin.store.v1beta1.Node")
	proto.RegisterType((*Child)(nil), "merlin.store.v1beta1.Child")
	proto.RegisterType((*Leaf)(nil), "merlin.store.v1beta1.Leaf")
}

func init() {
	proto.RegisterFile("merlin/sumtree/v1beta1/tree.proto", fileDescriptor_31a1c5f55b935f78)
}

var fileDescriptor_31a1c5f55b935f78 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x18, 0x85, 0xe3, 0x7b, 0x53, 0x04, 0xa6, 0x53, 0x54, 0xa4, 0xa8, 0x42, 0x6e, 0x94, 0x01, 0x65,
	0xa9, 0x4d, 0x60, 0xe9, 0x88, 0xca, 0x84, 0x40, 0x0c, 0x19, 0xd9, 0x1c, 0xc7, 0x4d, 0x2d, 0x92,
	0xb8, 0xc4, 0x0e, 0x82, 0xb7, 0xe0, 0xb1, 0x3a, 0x76, 0x44, 0x0c, 0x15, 0x4a, 0x5e, 0x04, 0xc5,
	0x09, 0x05, 0x24, 0x24, 0xa6, 0xdf, 0xc7, 0xfe, 0x74, 0xce, 0x91, 0x7f, 0xe8, 0x4b, 0x95, 0x4b,
	0x25, 0x14, 0x51, 0x55, 0xae, 0x4b, 0xce, 0xc9, 0x63, 0x18, 0x73, 0x4d, 0x43, 0xd2, 0x0a, 0xbc,
	0x2a, 0xa5, 0x96, 0xce, 0x51, 0xcf, 0x60, 0xa5, 0x65, 0xc9, 0x71, 0x4f, 0x8c, 0x47, 0xa9, 0x4c,
	0xa5, 0x21, 0x48, 0x7b, 0xea, 0xe0, 0x31, 0x62, 0x86, 0x26, 0x31, 0x55, 0x5f, 0x66, 0x4c, 0x8a,
	0xa2, 0x7b, 0xf7, 0x2f, 0xa0, 0x7d, 0x2b, 0x13, 0xee, 0xcc, 0xe0, 0x3e, 0x5b, 0x8a, 0x2c, 0x29,
	0x79, 0xe1, 0x02, 0xef, 0x7f, 0x70, 0x78, 0x76, 0x8c, 0x7f, 0xcd, 0xc1, 0x97, 0x2d, 0x16, 0xed,
	0x68, 0xff, 0x01, 0x0e, 0xcc, 0x95, 0x33, 0x82, 0x03, 0x51, 0x24, 0xfc, 0xc9, 0x05, 0x1e, 0x08,
	0x86, 0x51, 0x27, 0x9c, 0x08, 0x0e, 0x29, 0x63, 0x55, 0x5e, 0x65, 0x54, 0x0b, 0x59, 0xb8, 0xff,
	0x3c, 0x10, 0x1c, 0xcc, 0xf1, 0x7a, 0x3b, 0xb1, 0xde, 0xb6, 0x93, 0x93, 0x54, 0xe8, 0x65, 0x15,
	0x63, 0x26, 0x73, 0xd2, 0x37, 0xed, 0xc6, 0x54, 0x25, 0xf7, 0x44, 0x3f, 0xaf, 0xb8, 0xc2, 0x57,
	0x85, 0x8e, 0x7e, 0x78, 0xf8, 0x33, 0x68, 0xdf, 0x70, 0xba, 0x70, 0x4e, 0xa1, 0x9d, 0x71, 0xba,
	0x30, 0x81, 0x7f, 0x15, 0x36, 0xe4, 0xfc, 0x7a, 0x5d, 0x23, 0xb0, 0xa9, 0x11, 0x78, 0xaf, 0x11,
	0x78, 0x69, 0x90, 0xb5, 0x69, 0x90, 0xf5, 0xda, 0x20, 0xeb, 0x2e, 0xfc, 0xd6, 0xa4, 0xf7, 0x99,
	0x66, 0x34, 0x56, 0x9f, 0xc2, 0xcc, 0x4a, 0x8b, 0x6c, 0xb7, 0x9b, 0x78, 0xcf, 0x7c, 0xe1, 0xf9,
	0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x34, 0x4b, 0x32, 0xb5, 0x01, 0x00, 0x00,
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Children) > 0 {
		for iNdEx := len(m.Children) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Children[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTree(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Child) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Child) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Child) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Accumulation.Size()
		i -= size
		if _, err := m.Accumulation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTree(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTree(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Leaf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Leaf) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Leaf) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Leaf != nil {
		{
			size, err := m.Leaf.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTree(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTree(dAtA []byte, offset int, v uint64) int {
	offset -= sovTree(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Children) > 0 {
		for _, e := range m.Children {
			l = e.Size()
			n += 1 + l + sovTree(uint64(l))
		}
	}
	return n
}

func (m *Child) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTree(uint64(l))
	}
	l = m.Accumulation.Size()
	n += 1 + l + sovTree(uint64(l))
	return n
}

func (m *Leaf) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Leaf != nil {
		l = m.Leaf.Size()
		n += 1 + l + sovTree(uint64(l))
	}
	return n
}

func sovTree(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTree(x uint64) (n int) {
	return sovTree(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTree
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
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
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, &Child{})
			if err := m.Children[len(m.Children)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTree
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
func (m *Child) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTree
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
			return fmt.Errorf("proto: Child: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Child: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
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
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = append(m.Index[:0], dAtA[iNdEx:postIndex]...)
			if m.Index == nil {
				m.Index = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accumulation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Accumulation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTree
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
func (m *Leaf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTree
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
			return fmt.Errorf("proto: Leaf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Leaf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leaf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
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
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Leaf == nil {
				m.Leaf = &Child{}
			}
			if err := m.Leaf.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTree
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
func skipTree(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTree
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
					return 0, ErrIntOverflowTree
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
					return 0, ErrIntOverflowTree
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
				return 0, ErrInvalidLengthTree
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTree
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTree
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTree        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTree          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTree = fmt.Errorf("proto: unexpected end of group")
)

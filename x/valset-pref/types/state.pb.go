// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlin/valset-pref/v1beta1/state.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// ValidatorPreference defines the message structure for
// CreateValidatorSetPreference. It allows a user to set {val_addr, weight} in
// state. If a user does not have a validator set preference list set, and has
// staked, make their preference list default to their current staking
// distribution.
type ValidatorPreference struct {
	// val_oper_address holds the validator address the user wants to delegate
	// funds to.
	ValOperAddress string `protobuf:"bytes,1,opt,name=val_oper_address,json=valOperAddress,proto3" json:"val_oper_address,omitempty" yaml:"val_oper_address"`
	// weight is decimal between 0 and 1, and they all sum to 1.
	Weight github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight"`
}

func (m *ValidatorPreference) Reset()         { *m = ValidatorPreference{} }
func (m *ValidatorPreference) String() string { return proto.CompactTextString(m) }
func (*ValidatorPreference) ProtoMessage()    {}
func (*ValidatorPreference) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3010474a5b89fce, []int{0}
}
func (m *ValidatorPreference) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorPreference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorPreference.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorPreference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorPreference.Merge(m, src)
}
func (m *ValidatorPreference) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorPreference) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorPreference.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorPreference proto.InternalMessageInfo

// ValidatorSetPreferences defines a delegator's validator set preference.
// It contains a list of (validator, percent_allocation) pairs.
// The percent allocation are arranged in decimal notation from 0 to 1 and must
// add up to 1.
type ValidatorSetPreferences struct {
	// preference holds {valAddr, weight} for the user who created it.
	Preferences []ValidatorPreference `protobuf:"bytes,2,rep,name=preferences,proto3" json:"preferences" yaml:"preferences"`
}

func (m *ValidatorSetPreferences) Reset()         { *m = ValidatorSetPreferences{} }
func (m *ValidatorSetPreferences) String() string { return proto.CompactTextString(m) }
func (*ValidatorSetPreferences) ProtoMessage()    {}
func (*ValidatorSetPreferences) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3010474a5b89fce, []int{1}
}
func (m *ValidatorSetPreferences) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSetPreferences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSetPreferences.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSetPreferences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSetPreferences.Merge(m, src)
}
func (m *ValidatorSetPreferences) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSetPreferences) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSetPreferences.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSetPreferences proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ValidatorPreference)(nil), "merlin.valsetpref.v1beta1.ValidatorPreference")
	proto.RegisterType((*ValidatorSetPreferences)(nil), "merlin.valsetpref.v1beta1.ValidatorSetPreferences")
}

func init() {
	proto.RegisterFile("merlin/valset-pref/v1beta1/state.proto", fileDescriptor_d3010474a5b89fce)
}

var fileDescriptor_d3010474a5b89fce = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xeb, 0x40,
	0x18, 0xc5, 0x93, 0x5e, 0x28, 0xdc, 0x14, 0x44, 0xa2, 0xd0, 0x12, 0x25, 0x29, 0x59, 0x68, 0x37,
	0x9d, 0xa1, 0xba, 0x28, 0xb8, 0xb3, 0xa8, 0x5b, 0xa5, 0xa2, 0x0b, 0x37, 0x65, 0x92, 0x7c, 0x4d,
	0x83, 0x93, 0xcc, 0x30, 0x33, 0x46, 0xfb, 0x06, 0x2e, 0x7d, 0x08, 0x1f, 0xa6, 0xcb, 0x2e, 0xc5,
	0x45, 0xd0, 0xf6, 0x0d, 0xfa, 0x04, 0xd2, 0x24, 0xb4, 0x55, 0x74, 0x35, 0xff, 0x7e, 0xe7, 0xe3,
	0x9c, 0x39, 0xc6, 0x21, 0x93, 0x31, 0x93, 0x91, 0xc4, 0x29, 0xa1, 0x12, 0x54, 0x9b, 0x0b, 0x18,
	0xe2, 0xb4, 0xe3, 0x81, 0x22, 0x1d, 0x2c, 0x15, 0x51, 0x80, 0xb8, 0x60, 0x8a, 0x99, 0x56, 0x09,
	0xa2, 0x02, 0x5c, 0x72, 0xa8, 0xe4, 0xac, 0xdd, 0x90, 0x85, 0x2c, 0xc7, 0xf0, 0x72, 0x57, 0x28,
	0xac, 0xfd, 0x90, 0xb1, 0x90, 0x02, 0x26, 0x3c, 0xc2, 0x24, 0x49, 0x98, 0x22, 0x2a, 0x62, 0x89,
	0x2c, 0x5e, 0xdd, 0x57, 0xdd, 0xd8, 0xb9, 0x25, 0x34, 0x0a, 0x88, 0x62, 0xe2, 0x4a, 0xc0, 0x10,
	0x04, 0x24, 0x3e, 0x98, 0xe7, 0xc6, 0x76, 0x4a, 0xe8, 0x80, 0x71, 0x10, 0x03, 0x12, 0x04, 0x02,
	0xa4, 0x6c, 0xe8, 0x4d, 0xbd, 0xf5, 0xbf, 0xb7, 0xb7, 0xc8, 0x9c, 0xfa, 0x98, 0xc4, 0xf4, 0xc4,
	0xfd, 0x49, 0xb8, 0xfd, 0xad, 0x94, 0xd0, 0x4b, 0x0e, 0xe2, 0xb4, 0xb8, 0x30, 0x2f, 0x8c, 0xea,
	0x23, 0x44, 0xe1, 0x48, 0x35, 0x2a, 0xb9, 0x18, 0x4d, 0x32, 0x47, 0x7b, 0xcf, 0x9c, 0x83, 0x30,
	0x52, 0xa3, 0x07, 0x0f, 0xf9, 0x2c, 0xc6, 0x7e, 0x1e, 0xa9, 0x5c, 0xda, 0x32, 0xb8, 0xc7, 0x6a,
	0xcc, 0x41, 0xa2, 0x33, 0xf0, 0xfb, 0xa5, 0xda, 0x7d, 0xd6, 0x8d, 0xfa, 0xca, 0xe6, 0x35, 0xa8,
	0xb5, 0x53, 0x69, 0xc6, 0x46, 0x8d, 0xaf, 0x8f, 0x8d, 0x4a, 0xf3, 0x5f, 0xab, 0x76, 0x84, 0xd1,
	0xdf, 0x1f, 0x85, 0x7e, 0x09, 0xdc, 0xb3, 0x96, 0xce, 0x16, 0x99, 0x63, 0x16, 0xd1, 0x36, 0x26,
	0xba, 0xfd, 0xcd, 0xf9, 0xbd, 0x9b, 0xc9, 0xa7, 0xad, 0x4d, 0x66, 0xb6, 0x3e, 0x9d, 0xd9, 0xfa,
	0xc7, 0xcc, 0xd6, 0x5f, 0xe6, 0xb6, 0x36, 0x9d, 0xdb, 0xda, 0xdb, 0xdc, 0xd6, 0xee, 0xba, 0x1b,
	0xc1, 0x4a, 0x07, 0x6d, 0x4a, 0x3c, 0x89, 0x57, 0x05, 0x77, 0xba, 0xf8, 0xe9, 0x5b, 0xcd, 0x79,
	0x5a, 0xaf, 0x9a, 0xf7, 0x71, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x98, 0x30, 0x4f, 0x65, 0x0a,
	0x02, 0x00, 0x00,
}

func (m *ValidatorPreference) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorPreference) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorPreference) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ValOperAddress) > 0 {
		i -= len(m.ValOperAddress)
		copy(dAtA[i:], m.ValOperAddress)
		i = encodeVarintState(dAtA, i, uint64(len(m.ValOperAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSetPreferences) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSetPreferences) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSetPreferences) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Preferences) > 0 {
		for iNdEx := len(m.Preferences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Preferences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintState(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorPreference) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValOperAddress)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovState(uint64(l))
	return n
}

func (m *ValidatorSetPreferences) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Preferences) > 0 {
		for _, e := range m.Preferences {
			l = e.Size()
			n += 1 + l + sovState(uint64(l))
		}
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorPreference) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ValidatorPreference: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorPreference: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValOperAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValOperAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *ValidatorSetPreferences) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ValidatorSetPreferences: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSetPreferences: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Preferences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Preferences = append(m.Preferences, ValidatorPreference{})
			if err := m.Preferences[len(m.Preferences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)

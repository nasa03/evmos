// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/fees/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// Params defines the fee distribution module params
type Params struct {
	// fee distribution for registered contract withdraw addresses
	ContractDistribution github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=contract_distribution,json=contractDistribution,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"contract_distribution"`
	// enable withawal address
	WithdrawAddrEnabled bool `protobuf:"varint,2,opt,name=withdraw_addr_enabled,json=withdrawAddrEnabled,proto3" json:"withdraw_addr_enabled,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_27418e5be97fbe06, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetWithdrawAddrEnabled() bool {
	if m != nil {
		return m.WithdrawAddrEnabled
	}
	return false
}

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// fee distribution module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// hex addresses of contract used for withdrawing tx fees
	WithdrawAddresses []ContractWithdrawAddress `protobuf:"bytes,2,rep,name=withdraw_addresses,json=withdrawAddresses,proto3" json:"withdraw_addresses"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_27418e5be97fbe06, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetWithdrawAddresses() []ContractWithdrawAddress {
	if m != nil {
		return m.WithdrawAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "evmos.fees.v1.Params")
	proto.RegisterType((*GenesisState)(nil), "evmos.fees.v1.GenesisState")
}

func init() { proto.RegisterFile("evmos/fees/v1/genesis.proto", fileDescriptor_27418e5be97fbe06) }

var fileDescriptor_27418e5be97fbe06 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0xed, 0xf0, 0x5e, 0xc8, 0x7b, 0x45, 0x17, 0x56, 0x48, 0x1a, 0x4c, 0x0a, 0x21, 0x91, 0xb0,
	0x71, 0x26, 0xc0, 0x07, 0x18, 0x11, 0xe3, 0xd6, 0xe0, 0xc2, 0x44, 0x17, 0x64, 0xda, 0xb9, 0x94,
	0x46, 0xe8, 0x90, 0xb9, 0x53, 0xd0, 0xbf, 0xf0, 0x13, 0xf4, 0x6f, 0x58, 0xb2, 0x34, 0x2e, 0x88,
	0x81, 0x1f, 0x31, 0x9d, 0x42, 0x42, 0x5d, 0xcd, 0x9d, 0x7b, 0x4e, 0xce, 0x9c, 0x39, 0xc7, 0x3e,
	0x83, 0xf9, 0x54, 0x22, 0x1b, 0x01, 0x20, 0x9b, 0xb7, 0x59, 0x08, 0x31, 0x60, 0x84, 0x74, 0xa6,
	0xa4, 0x96, 0xce, 0xb1, 0x01, 0x69, 0x0a, 0xd2, 0x79, 0xbb, 0xea, 0xe6, 0xb9, 0x66, 0x6d, 0x88,
	0xd5, 0x72, 0x28, 0x43, 0x69, 0x46, 0x96, 0x4e, 0xbb, 0xad, 0x17, 0x4a, 0x19, 0x4e, 0x80, 0x99,
	0x9b, 0x9f, 0x8c, 0x98, 0x48, 0x14, 0xd7, 0x91, 0x8c, 0x33, 0xbc, 0xf1, 0x41, 0xec, 0xe2, 0x1d,
	0x57, 0x7c, 0x8a, 0x4e, 0x60, 0x57, 0x02, 0x19, 0x6b, 0xc5, 0x03, 0x3d, 0x14, 0x11, 0x6a, 0x15,
	0xf9, 0x49, 0xca, 0x74, 0x49, 0x9d, 0xb4, 0xfe, 0xf7, 0xe8, 0x72, 0x5d, 0xb3, 0xbe, 0xd6, 0xb5,
	0x66, 0x18, 0xe9, 0x71, 0xe2, 0xd3, 0x40, 0x4e, 0x59, 0x20, 0x31, 0x75, 0x93, 0x1d, 0x17, 0x28,
	0x9e, 0x99, 0x7e, 0x9d, 0x01, 0xd2, 0x3e, 0x04, 0x83, 0xf2, 0x5e, 0xac, 0x7f, 0xa0, 0xe5, 0x74,
	0xec, 0xca, 0x22, 0xd2, 0x63, 0xa1, 0xf8, 0x62, 0xc8, 0x85, 0x50, 0x43, 0x88, 0xb9, 0x3f, 0x01,
	0xe1, 0x16, 0xea, 0xa4, 0xf5, 0x6f, 0x70, 0xba, 0x07, 0xaf, 0x84, 0x50, 0x37, 0x19, 0xd4, 0x78,
	0x27, 0xf6, 0xd1, 0x6d, 0x16, 0xca, 0xbd, 0xe6, 0x1a, 0x9c, 0xae, 0x5d, 0x9c, 0x19, 0xcf, 0xc6,
	0x5a, 0xa9, 0x53, 0xa1, 0xb9, 0x90, 0x68, 0xf6, 0xa1, 0xde, 0xdf, 0xd4, 0xf1, 0x60, 0x47, 0x75,
	0x9e, 0x6c, 0x27, 0xf7, 0x32, 0x20, 0x02, 0xba, 0x85, 0xfa, 0x9f, 0x56, 0xa9, 0xd3, 0xfc, 0x25,
	0x70, 0xbd, 0xb3, 0xfe, 0x70, 0xe0, 0x06, 0x70, 0xaf, 0x78, 0xb2, 0xc8, 0xaf, 0x01, 0x7b, 0x97,
	0xcb, 0x8d, 0x47, 0x56, 0x1b, 0x8f, 0x7c, 0x6f, 0x3c, 0xf2, 0xb6, 0xf5, 0xac, 0xd5, 0xd6, 0xb3,
	0x3e, 0xb7, 0x9e, 0xf5, 0x78, 0x7e, 0x10, 0x97, 0x1e, 0x73, 0x85, 0x11, 0xb2, 0xac, 0xc3, 0x97,
	0xac, 0x45, 0x93, 0x98, 0x5f, 0x34, 0x75, 0x74, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x36,
	0x80, 0x6f, 0x0c, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WithdrawAddrEnabled {
		i--
		if m.WithdrawAddrEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.ContractDistribution.Size()
		i -= size
		if _, err := m.ContractDistribution.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawAddresses) > 0 {
		for iNdEx := len(m.WithdrawAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ContractDistribution.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.WithdrawAddrEnabled {
		n += 2
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.WithdrawAddresses) > 0 {
		for _, e := range m.WithdrawAddresses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractDistribution", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ContractDistribution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddrEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.WithdrawAddrEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawAddresses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawAddresses = append(m.WithdrawAddresses, ContractWithdrawAddress{})
			if err := m.WithdrawAddresses[len(m.WithdrawAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
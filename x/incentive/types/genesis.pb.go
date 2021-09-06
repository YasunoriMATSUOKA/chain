// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: incentive/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the incentive module's genesis state.
type GenesisState struct {
	Params               Params                    `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
	CdpAccumulationTimes []GenesisAccumulationTime `protobuf:"bytes,2,rep,name=cdp_accumulation_times,json=cdpAccumulationTimes,proto3" json:"cdp_accumulation_times" yaml:"cdp_accumulation_times"`
	CdpMintingClaims     []CdpMintingClaim         `protobuf:"bytes,3,rep,name=cdp_minting_claims,json=cdpMintingClaims,proto3" json:"cdp_minting_claims" yaml:"cdp_minting_claims"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ea08f29a85d2f6, []int{0}
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

func (m *GenesisState) GetCdpAccumulationTimes() []GenesisAccumulationTime {
	if m != nil {
		return m.CdpAccumulationTimes
	}
	return nil
}

func (m *GenesisState) GetCdpMintingClaims() []CdpMintingClaim {
	if m != nil {
		return m.CdpMintingClaims
	}
	return nil
}

type GenesisAccumulationTime struct {
	CollateralType           string    `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty" yaml:"collateral_type"`
	PreviousAccumulationTime time.Time `protobuf:"bytes,2,opt,name=previous_accumulation_time,json=previousAccumulationTime,proto3,stdtime" json:"previous_accumulation_time" yaml:"previous_accumulation_time"`
}

func (m *GenesisAccumulationTime) Reset()         { *m = GenesisAccumulationTime{} }
func (m *GenesisAccumulationTime) String() string { return proto.CompactTextString(m) }
func (*GenesisAccumulationTime) ProtoMessage()    {}
func (*GenesisAccumulationTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5ea08f29a85d2f6, []int{1}
}
func (m *GenesisAccumulationTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisAccumulationTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisAccumulationTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisAccumulationTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccumulationTime.Merge(m, src)
}
func (m *GenesisAccumulationTime) XXX_Size() int {
	return m.Size()
}
func (m *GenesisAccumulationTime) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccumulationTime.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccumulationTime proto.InternalMessageInfo

func (m *GenesisAccumulationTime) GetCollateralType() string {
	if m != nil {
		return m.CollateralType
	}
	return ""
}

func (m *GenesisAccumulationTime) GetPreviousAccumulationTime() time.Time {
	if m != nil {
		return m.PreviousAccumulationTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "botany.incentive.GenesisState")
	proto.RegisterType((*GenesisAccumulationTime)(nil), "botany.incentive.GenesisAccumulationTime")
}

func init() { proto.RegisterFile("incentive/genesis.proto", fileDescriptor_b5ea08f29a85d2f6) }

var fileDescriptor_b5ea08f29a85d2f6 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x3d, 0x6f, 0xdb, 0x30,
	0x10, 0x35, 0x1d, 0x20, 0x40, 0x95, 0x7e, 0x04, 0x42, 0x9a, 0x28, 0x02, 0x2a, 0xc5, 0x02, 0x8a,
	0xa6, 0x43, 0x49, 0x20, 0xdd, 0xba, 0x45, 0x1e, 0x32, 0x15, 0x28, 0x54, 0x4f, 0x5d, 0x0c, 0x9a,
	0x66, 0x55, 0x16, 0xfc, 0x82, 0x48, 0x05, 0xd1, 0x1f, 0x68, 0xd7, 0xfc, 0xac, 0x8c, 0x19, 0x3a,
	0x74, 0x72, 0x0b, 0x7b, 0xe9, 0x9c, 0x5f, 0x50, 0x98, 0xa2, 0xe3, 0x3a, 0xaa, 0x37, 0x4a, 0xf7,
	0xde, 0xbd, 0x77, 0x77, 0x2f, 0x38, 0x62, 0x92, 0x50, 0x69, 0xd9, 0x25, 0x45, 0x25, 0x95, 0xd4,
	0x30, 0x03, 0x75, 0xa5, 0xac, 0x0a, 0xf7, 0x27, 0xca, 0x62, 0xd9, 0xc0, 0xfb, 0x7a, 0x7c, 0x50,
	0xaa, 0x52, 0xb9, 0x22, 0x5a, 0xbe, 0x5a, 0x5c, 0x9c, 0x96, 0x4a, 0x95, 0x9c, 0x22, 0xf7, 0x35,
	0xa9, 0x3f, 0x23, 0xcb, 0x04, 0x35, 0x16, 0x0b, 0xed, 0x01, 0xc7, 0x6b, 0x85, 0xfb, 0x57, 0x5b,
	0xca, 0x7e, 0xf4, 0x83, 0xc7, 0x17, 0xad, 0xea, 0x47, 0x8b, 0x2d, 0x0d, 0x2f, 0x82, 0x5d, 0x8d,
	0x2b, 0x2c, 0x4c, 0x04, 0x4e, 0xc0, 0xe9, 0xde, 0x59, 0x04, 0x1f, 0xba, 0x80, 0x1f, 0x5c, 0x3d,
	0x7f, 0x7e, 0x33, 0x4b, 0x7b, 0x77, 0xb3, 0xf4, 0x49, 0x83, 0x05, 0x7f, 0x97, 0xb5, 0xac, 0xac,
	0xf0, 0xf4, 0xf0, 0x1b, 0x08, 0x0e, 0xc9, 0x54, 0x8f, 0x31, 0x21, 0xb5, 0xa8, 0x39, 0xb6, 0x4c,
	0xc9, 0xb1, 0x73, 0x16, 0xf5, 0x4f, 0x76, 0x4e, 0xf7, 0xce, 0x5e, 0x77, 0x3b, 0x7b, 0x27, 0xe7,
	0xff, 0x50, 0x46, 0x4c, 0xd0, 0xfc, 0xa5, 0x97, 0x7a, 0xd1, 0x4a, 0xfd, 0xbf, 0x6d, 0x56, 0x1c,
	0x90, 0xa9, 0x7e, 0xc8, 0x35, 0x61, 0x15, 0x84, 0x4b, 0x82, 0x60, 0xd2, 0x32, 0x59, 0x8e, 0x09,
	0xc7, 0x4c, 0x98, 0x68, 0xc7, 0x79, 0x18, 0x74, 0x3d, 0x0c, 0xa7, 0xfa, 0x7d, 0x0b, 0x1d, 0x2e,
	0x91, 0xf9, 0xc0, 0x6b, 0x1f, 0xaf, 0xb5, 0x37, 0x5b, 0x65, 0xc5, 0x3e, 0xd9, 0xe4, 0x98, 0xec,
	0x0f, 0x08, 0x8e, 0xb6, 0x0c, 0x13, 0x0e, 0x83, 0x67, 0x44, 0x71, 0x8e, 0x2d, 0xad, 0x30, 0x1f,
	0xdb, 0x46, 0x53, 0xb7, 0xea, 0x47, 0x79, 0x7c, 0x37, 0x4b, 0x0f, 0xbd, 0xca, 0x26, 0x20, 0x2b,
	0x9e, 0xae, 0xff, 0x8c, 0x1a, 0x4d, 0xc3, 0xef, 0x20, 0x88, 0x75, 0x45, 0x2f, 0x99, 0xaa, 0x4d,
	0x77, 0x17, 0x51, 0xdf, 0xdd, 0x2e, 0x86, 0x6d, 0x32, 0xe0, 0x2a, 0x19, 0x70, 0xb4, 0x4a, 0x46,
	0xfe, 0xc6, 0x8f, 0x35, 0xf0, 0xd7, 0xdb, 0xda, 0x2b, 0xbb, 0xfe, 0x95, 0x82, 0x22, 0x5a, 0x01,
	0x3a, 0xb7, 0x39, 0xbf, 0x99, 0x27, 0xe0, 0x76, 0x9e, 0x80, 0xdf, 0xf3, 0x04, 0x5c, 0x2f, 0x92,
	0xde, 0xed, 0x22, 0xe9, 0xfd, 0x5c, 0x24, 0xbd, 0x4f, 0xaf, 0x4a, 0x66, 0xbf, 0xd4, 0x13, 0x48,
	0x94, 0x40, 0x9c, 0x48, 0x2a, 0xd0, 0x57, 0xdd, 0x5c, 0xa1, 0xab, 0x75, 0x08, 0xd1, 0x72, 0x38,
	0x33, 0xd9, 0x75, 0xfe, 0xde, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x55, 0x2e, 0x87, 0x0a,
	0x03, 0x00, 0x00,
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
	if len(m.CdpMintingClaims) > 0 {
		for iNdEx := len(m.CdpMintingClaims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CdpMintingClaims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.CdpAccumulationTimes) > 0 {
		for iNdEx := len(m.CdpAccumulationTimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CdpAccumulationTimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GenesisAccumulationTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisAccumulationTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisAccumulationTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.PreviousAccumulationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.CdpAccumulationTimes) > 0 {
		for _, e := range m.CdpAccumulationTimes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CdpMintingClaims) > 0 {
		for _, e := range m.CdpMintingClaims {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisAccumulationTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return fmt.Errorf("proto: wrong wireType = %d for field CdpAccumulationTimes", wireType)
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
			m.CdpAccumulationTimes = append(m.CdpAccumulationTimes, GenesisAccumulationTime{})
			if err := m.CdpAccumulationTimes[len(m.CdpAccumulationTimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdpMintingClaims", wireType)
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
			m.CdpMintingClaims = append(m.CdpMintingClaims, CdpMintingClaim{})
			if err := m.CdpMintingClaims[len(m.CdpMintingClaims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisAccumulationTime) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisAccumulationTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisAccumulationTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
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
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAccumulationTime", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.PreviousAccumulationTime, dAtA[iNdEx:postIndex]); err != nil {
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

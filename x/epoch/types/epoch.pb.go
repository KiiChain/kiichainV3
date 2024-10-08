// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: epoch/epoch.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "github.com/gogo/protobuf/types"
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

type Epoch struct {
	GenesisTime           time.Time     `protobuf:"bytes,1,opt,name=genesis_time,json=genesisTime,proto3,stdtime" json:"genesis_time" yaml:"genesis_time"`
	EpochDuration         time.Duration `protobuf:"bytes,2,opt,name=epoch_duration,json=epochDuration,proto3,stdduration" json:"duration,omitempty" yaml:"epoch_duration"`
	CurrentEpoch          uint64        `protobuf:"varint,3,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch" yaml:"current_epoch"`
	CurrentEpochStartTime time.Time     `protobuf:"bytes,4,opt,name=current_epoch_start_time,json=currentEpochStartTime,proto3,stdtime" json:"current_epoch_start_time" yaml:"current_epoch_start_time"`
	CurrentEpochHeight    int64         `protobuf:"varint,5,opt,name=current_epoch_height,json=currentEpochHeight,proto3" json:"current_epoch_height" yaml:"current_epoch_height"`
}

func (m *Epoch) Reset()         { *m = Epoch{} }
func (m *Epoch) String() string { return proto.CompactTextString(m) }
func (*Epoch) ProtoMessage()    {}
func (*Epoch) Descriptor() ([]byte, []int) {
	return fileDescriptor_36a9d1673530db42, []int{0}
}
func (m *Epoch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Epoch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Epoch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Epoch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Epoch.Merge(m, src)
}
func (m *Epoch) XXX_Size() int {
	return m.Size()
}
func (m *Epoch) XXX_DiscardUnknown() {
	xxx_messageInfo_Epoch.DiscardUnknown(m)
}

var xxx_messageInfo_Epoch proto.InternalMessageInfo

func (m *Epoch) GetGenesisTime() time.Time {
	if m != nil {
		return m.GenesisTime
	}
	return time.Time{}
}

func (m *Epoch) GetEpochDuration() time.Duration {
	if m != nil {
		return m.EpochDuration
	}
	return 0
}

func (m *Epoch) GetCurrentEpoch() uint64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func (m *Epoch) GetCurrentEpochStartTime() time.Time {
	if m != nil {
		return m.CurrentEpochStartTime
	}
	return time.Time{}
}

func (m *Epoch) GetCurrentEpochHeight() int64 {
	if m != nil {
		return m.CurrentEpochHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Epoch)(nil), "epoch.Epoch")
}

func init() { proto.RegisterFile("epoch/epoch.proto", fileDescriptor_36a9d1673530db42) }

var fileDescriptor_36a9d1673530db42 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x3f, 0xeb, 0xd3, 0x40,
	0x1c, 0xc6, 0x73, 0xb6, 0x75, 0x48, 0x5b, 0xc1, 0xd8, 0x42, 0xac, 0x90, 0x2b, 0x99, 0x2a, 0x96,
	0x1c, 0xd8, 0xa1, 0xe0, 0x18, 0xff, 0x20, 0x08, 0x0e, 0x55, 0x1c, 0x1c, 0x0c, 0x69, 0x3c, 0x93,
	0xc3, 0x26, 0x17, 0x92, 0x0b, 0x98, 0xcd, 0x97, 0xd0, 0xd1, 0xd1, 0x97, 0xd3, 0xb1, 0xa3, 0xd3,
	0x29, 0xed, 0xd6, 0x31, 0xaf, 0x40, 0xee, 0x2e, 0x81, 0xe4, 0xf7, 0x2b, 0xfc, 0x96, 0x70, 0xf7,
	0x3c, 0xdf, 0xef, 0xf3, 0x49, 0x1e, 0xa2, 0x3f, 0xc4, 0x29, 0x0d, 0x22, 0x24, 0x9f, 0x4e, 0x9a,
	0x51, 0x46, 0x8d, 0x81, 0xbc, 0xcc, 0x26, 0x21, 0x0d, 0xa9, 0x54, 0x90, 0x38, 0x29, 0x73, 0x06,
	0x43, 0x4a, 0xc3, 0x1d, 0x46, 0xf2, 0xb6, 0x2d, 0xbe, 0x21, 0x46, 0x62, 0x9c, 0x33, 0x3f, 0x4e,
	0xeb, 0x01, 0xeb, 0xe6, 0xc0, 0xd7, 0x22, 0xf3, 0x19, 0xa1, 0x89, 0xf2, 0xed, 0xdf, 0x7d, 0x7d,
	0xf0, 0x5a, 0x00, 0x8c, 0x2f, 0xfa, 0x28, 0xc4, 0x09, 0xce, 0x49, 0xee, 0x89, 0x10, 0x13, 0xcc,
	0xc1, 0x62, 0xf8, 0x7c, 0xe6, 0xa8, 0x00, 0xa7, 0x09, 0x70, 0x3e, 0x36, 0x04, 0x17, 0x1e, 0x38,
	0xd4, 0x2a, 0x0e, 0x1f, 0x95, 0x7e, 0xbc, 0x7b, 0x61, 0xb7, 0xb7, 0xed, 0xfd, 0x5f, 0x08, 0x36,
	0xc3, 0x5a, 0x12, 0x2b, 0x46, 0xa9, 0x3f, 0x90, 0x5f, 0xe2, 0x35, 0x6f, 0x60, 0xde, 0x93, 0x84,
	0xc7, 0xb7, 0x08, 0xaf, 0xea, 0x01, 0x77, 0x2d, 0x00, 0x17, 0x0e, 0x8d, 0x66, 0x65, 0x49, 0x63,
	0xc2, 0x70, 0x9c, 0xb2, 0xb2, 0xe2, 0x70, 0xaa, 0xb0, 0xdd, 0x50, 0xfb, 0x97, 0x00, 0x8f, 0xa5,
	0xd8, 0xe4, 0x18, 0xef, 0xf5, 0x71, 0x50, 0x64, 0x19, 0x4e, 0x98, 0x27, 0x0d, 0xb3, 0x37, 0x07,
	0x8b, 0xbe, 0xfb, 0xf4, 0xc2, 0x61, 0xd7, 0xa8, 0x38, 0x9c, 0xa8, 0xd4, 0x8e, 0x6c, 0x6f, 0x46,
	0xf5, 0x5d, 0x55, 0xf5, 0x13, 0xe8, 0x66, 0x67, 0xc0, 0xcb, 0x99, 0x9f, 0x31, 0xd5, 0x5b, 0xff,
	0xce, 0xde, 0x9e, 0xd5, 0xbd, 0xc1, 0x2b, 0xa8, 0x56, 0x92, 0xea, 0x70, 0xda, 0x26, 0x7f, 0x10,
	0xa6, 0x6c, 0x93, 0xe8, 0x93, 0xee, 0x5e, 0x84, 0x49, 0x18, 0x31, 0x73, 0x30, 0x07, 0x8b, 0x9e,
	0xbb, 0xbe, 0x70, 0x78, 0xd5, 0xaf, 0x38, 0x7c, 0x72, 0x8d, 0xaa, 0x5c, 0x7b, 0x63, 0xb4, 0x69,
	0x6f, 0xa5, 0xe8, 0xbe, 0x39, 0x9c, 0x2c, 0x70, 0x3c, 0x59, 0xe0, 0xdf, 0xc9, 0x02, 0xfb, 0xb3,
	0xa5, 0x1d, 0xcf, 0x96, 0xf6, 0xe7, 0x6c, 0x69, 0x9f, 0x97, 0x21, 0x61, 0x51, 0xb1, 0x75, 0x02,
	0x1a, 0xa3, 0x77, 0x84, 0xbc, 0x8c, 0x7c, 0x92, 0xa0, 0xef, 0x84, 0x04, 0xe2, 0xf0, 0x69, 0x85,
	0x7e, 0xa8, 0x1f, 0x19, 0xb1, 0x32, 0xc5, 0xf9, 0xf6, 0xbe, 0xac, 0x62, 0xf5, 0x3f, 0x00, 0x00,
	0xff, 0xff, 0x37, 0x3d, 0xd0, 0x1a, 0xe4, 0x02, 0x00, 0x00,
}

func (m *Epoch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Epoch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Epoch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpochHeight != 0 {
		i = encodeVarintEpoch(dAtA, i, uint64(m.CurrentEpochHeight))
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CurrentEpochStartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CurrentEpochStartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEpoch(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.CurrentEpoch != 0 {
		i = encodeVarintEpoch(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x18
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.EpochDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.EpochDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEpoch(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.GenesisTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.GenesisTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintEpoch(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEpoch(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpoch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Epoch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.GenesisTime)
	n += 1 + l + sovEpoch(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.EpochDuration)
	n += 1 + l + sovEpoch(uint64(l))
	if m.CurrentEpoch != 0 {
		n += 1 + sovEpoch(uint64(m.CurrentEpoch))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CurrentEpochStartTime)
	n += 1 + l + sovEpoch(uint64(l))
	if m.CurrentEpochHeight != 0 {
		n += 1 + sovEpoch(uint64(m.CurrentEpochHeight))
	}
	return n
}

func sovEpoch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpoch(x uint64) (n int) {
	return sovEpoch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Epoch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpoch
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
			return fmt.Errorf("proto: Epoch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Epoch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.GenesisTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.EpochDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CurrentEpochStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochHeight", wireType)
			}
			m.CurrentEpochHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpochHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpoch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpoch
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
func skipEpoch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpoch
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
					return 0, ErrIntOverflowEpoch
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
					return 0, ErrIntOverflowEpoch
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
				return 0, ErrInvalidLengthEpoch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpoch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpoch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpoch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpoch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpoch = fmt.Errorf("proto: unexpected end of group")
)

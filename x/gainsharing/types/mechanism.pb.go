// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vita/gainsharing/mechanism.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Mechanism struct {
	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Metrics       string `protobuf:"bytes,2,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Coefficients  string `protobuf:"bytes,3,opt,name=coefficients,proto3" json:"coefficients,omitempty"`
	ConvergeLimit string `protobuf:"bytes,4,opt,name=convergeLimit,proto3" json:"convergeLimit,omitempty"`
	Slope         string `protobuf:"bytes,5,opt,name=slope,proto3" json:"slope,omitempty"`
	Creator       string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Mechanism) Reset()         { *m = Mechanism{} }
func (m *Mechanism) String() string { return proto.CompactTextString(m) }
func (*Mechanism) ProtoMessage()    {}
func (*Mechanism) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e74e94acf6e87bf, []int{0}
}
func (m *Mechanism) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mechanism) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mechanism.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mechanism) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mechanism.Merge(m, src)
}
func (m *Mechanism) XXX_Size() int {
	return m.Size()
}
func (m *Mechanism) XXX_DiscardUnknown() {
	xxx_messageInfo_Mechanism.DiscardUnknown(m)
}

var xxx_messageInfo_Mechanism proto.InternalMessageInfo

func (m *Mechanism) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Mechanism) GetMetrics() string {
	if m != nil {
		return m.Metrics
	}
	return ""
}

func (m *Mechanism) GetCoefficients() string {
	if m != nil {
		return m.Coefficients
	}
	return ""
}

func (m *Mechanism) GetConvergeLimit() string {
	if m != nil {
		return m.ConvergeLimit
	}
	return ""
}

func (m *Mechanism) GetSlope() string {
	if m != nil {
		return m.Slope
	}
	return ""
}

func (m *Mechanism) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Mechanism)(nil), "vita.gainsharing.Mechanism")
}

func init() { proto.RegisterFile("vita/gainsharing/mechanism.proto", fileDescriptor_2e74e94acf6e87bf) }

var fileDescriptor_2e74e94acf6e87bf = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0xcb, 0x2c, 0x49,
	0xd4, 0x4f, 0x4f, 0xcc, 0xcc, 0x2b, 0xce, 0x48, 0x2c, 0xca, 0xcc, 0x4b, 0xd7, 0xcf, 0x4d, 0x4d,
	0xce, 0x48, 0xcc, 0xcb, 0x2c, 0xce, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x00, 0xa9,
	0xd0, 0x43, 0x52, 0xa1, 0xb4, 0x9e, 0x91, 0x8b, 0xd3, 0x17, 0xa6, 0x4a, 0x88, 0x8f, 0x8b, 0x29,
	0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x88, 0x29, 0x33, 0x45, 0x48, 0x82, 0x8b, 0x3d,
	0x37, 0xb5, 0xa4, 0x28, 0x33, 0xb9, 0x58, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15,
	0x52, 0xe2, 0xe2, 0x49, 0xce, 0x4f, 0x4d, 0x4b, 0xcb, 0x4c, 0xce, 0x4c, 0xcd, 0x2b, 0x29, 0x96,
	0x60, 0x06, 0x4b, 0xa3, 0x88, 0x09, 0xa9, 0x70, 0xf1, 0x26, 0xe7, 0xe7, 0x95, 0xa5, 0x16, 0xa5,
	0xa7, 0xfa, 0x64, 0xe6, 0x66, 0x96, 0x48, 0xb0, 0x80, 0x15, 0xa1, 0x0a, 0x0a, 0x89, 0x70, 0xb1,
	0x16, 0xe7, 0xe4, 0x17, 0xa4, 0x4a, 0xb0, 0x82, 0x65, 0x21, 0x1c, 0x90, 0xcd, 0xc9, 0x45, 0xa9,
	0x89, 0x25, 0xf9, 0x45, 0x12, 0x6c, 0x10, 0x9b, 0xa1, 0x5c, 0x27, 0xa3, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x92, 0x00, 0xfb, 0xbf, 0x02, 0x25, 0x04, 0x4a, 0x2a, 0x0b,
	0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xde, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xd8, 0xea,
	0xff, 0x22, 0x01, 0x00, 0x00,
}

func (m *Mechanism) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mechanism) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mechanism) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMechanism(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Slope) > 0 {
		i -= len(m.Slope)
		copy(dAtA[i:], m.Slope)
		i = encodeVarintMechanism(dAtA, i, uint64(len(m.Slope)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ConvergeLimit) > 0 {
		i -= len(m.ConvergeLimit)
		copy(dAtA[i:], m.ConvergeLimit)
		i = encodeVarintMechanism(dAtA, i, uint64(len(m.ConvergeLimit)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Coefficients) > 0 {
		i -= len(m.Coefficients)
		copy(dAtA[i:], m.Coefficients)
		i = encodeVarintMechanism(dAtA, i, uint64(len(m.Coefficients)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Metrics) > 0 {
		i -= len(m.Metrics)
		copy(dAtA[i:], m.Metrics)
		i = encodeVarintMechanism(dAtA, i, uint64(len(m.Metrics)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintMechanism(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMechanism(dAtA []byte, offset int, v uint64) int {
	offset -= sovMechanism(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mechanism) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMechanism(uint64(m.Id))
	}
	l = len(m.Metrics)
	if l > 0 {
		n += 1 + l + sovMechanism(uint64(l))
	}
	l = len(m.Coefficients)
	if l > 0 {
		n += 1 + l + sovMechanism(uint64(l))
	}
	l = len(m.ConvergeLimit)
	if l > 0 {
		n += 1 + l + sovMechanism(uint64(l))
	}
	l = len(m.Slope)
	if l > 0 {
		n += 1 + l + sovMechanism(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMechanism(uint64(l))
	}
	return n
}

func sovMechanism(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMechanism(x uint64) (n int) {
	return sovMechanism(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mechanism) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMechanism
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
			return fmt.Errorf("proto: Mechanism: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mechanism: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMechanism
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMechanism
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
				return ErrInvalidLengthMechanism
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMechanism
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metrics = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coefficients", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMechanism
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
				return ErrInvalidLengthMechanism
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMechanism
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coefficients = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConvergeLimit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMechanism
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
				return ErrInvalidLengthMechanism
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMechanism
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConvergeLimit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slope", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMechanism
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
				return ErrInvalidLengthMechanism
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMechanism
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Slope = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMechanism
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
				return ErrInvalidLengthMechanism
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMechanism
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMechanism(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMechanism
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
func skipMechanism(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMechanism
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
					return 0, ErrIntOverflowMechanism
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
					return 0, ErrIntOverflowMechanism
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
				return 0, ErrInvalidLengthMechanism
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMechanism
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMechanism
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMechanism        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMechanism          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMechanism = fmt.Errorf("proto: unexpected end of group")
)

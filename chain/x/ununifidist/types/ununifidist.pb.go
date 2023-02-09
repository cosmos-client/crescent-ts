// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ununifidist/ununifidist.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	Active  bool     `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty" yaml:"active"`
	Periods []Period `protobuf:"bytes,2,rep,name=periods,proto3" json:"periods" yaml:"periods"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_57b198cfe07074af, []int{0}
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

func (m *Params) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Params) GetPeriods() []Period {
	if m != nil {
		return m.Periods
	}
	return nil
}

type Period struct {
	Start     time.Time                              `protobuf:"bytes,1,opt,name=start,proto3,stdtime" json:"start" yaml:"start"`
	End       time.Time                              `protobuf:"bytes,2,opt,name=end,proto3,stdtime" json:"end" yaml:"end"`
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation" yaml:"inflation"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_57b198cfe07074af, []int{1}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(m, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetStart() time.Time {
	if m != nil {
		return m.Start
	}
	return time.Time{}
}

func (m *Period) GetEnd() time.Time {
	if m != nil {
		return m.End
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Params)(nil), "ununifi.ununifidist.Params")
	proto.RegisterType((*Period)(nil), "ununifi.ununifidist.Period")
}

func init() { proto.RegisterFile("ununifidist/ununifidist.proto", fileDescriptor_57b198cfe07074af) }

var fileDescriptor_57b198cfe07074af = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x33, 0x2d, 0x56, 0x3b, 0x55, 0xd1, 0x28, 0x12, 0x2a, 0x26, 0x25, 0x07, 0xa9, 0xa0,
	0x33, 0xb4, 0xde, 0x3c, 0x86, 0xd2, 0x83, 0x20, 0x94, 0x60, 0x2f, 0x9e, 0x9c, 0x24, 0xd3, 0x74,
	0xb0, 0x99, 0x09, 0x99, 0x49, 0xb1, 0x57, 0x3f, 0x41, 0x3f, 0x56, 0x8f, 0x3d, 0x8a, 0x87, 0xb8,
	0xb4, 0xdf, 0xa0, 0xe7, 0x3d, 0x2c, 0xc9, 0x24, 0xbb, 0x39, 0x2c, 0xec, 0x29, 0x2f, 0xef, 0xfd,
	0xdf, 0xef, 0xfd, 0xff, 0x21, 0xf0, 0x5d, 0xce, 0x73, 0xce, 0x56, 0x2c, 0x62, 0x52, 0xe1, 0x56,
	0x8d, 0xd2, 0x4c, 0x28, 0x61, 0xbe, 0xaa, 0x5b, 0xa8, 0x35, 0x1a, 0xbe, 0x8e, 0x45, 0x2c, 0xaa,
	0x39, 0x2e, 0x2b, 0x2d, 0x1d, 0x3a, 0xb1, 0x10, 0xf1, 0x86, 0xe2, 0xea, 0x2d, 0xc8, 0x57, 0x58,
	0xb1, 0x84, 0x4a, 0x45, 0x92, 0xb4, 0x16, 0xd8, 0xa1, 0x90, 0x89, 0x90, 0x38, 0x20, 0x92, 0xe2,
	0xed, 0x24, 0xa0, 0x8a, 0x4c, 0x70, 0x28, 0x18, 0xd7, 0x73, 0xf7, 0x0f, 0x80, 0xbd, 0x05, 0xc9,
	0x48, 0x22, 0xcd, 0x0f, 0xb0, 0x47, 0x42, 0xc5, 0xb6, 0xd4, 0x02, 0x23, 0x30, 0x7e, 0xe2, 0xbd,
	0xbc, 0x14, 0xce, 0xb3, 0x1d, 0x49, 0x36, 0x5f, 0x5c, 0xdd, 0x77, 0xfd, 0x5a, 0x60, 0x7e, 0x83,
	0x8f, 0x53, 0x9a, 0x31, 0x11, 0x49, 0xab, 0x33, 0xea, 0x8e, 0x07, 0xd3, 0xb7, 0xe8, 0x1e, 0xcf,
	0x68, 0x51, 0x69, 0xbc, 0x37, 0x87, 0xc2, 0x31, 0x2e, 0x85, 0xf3, 0x5c, 0xc3, 0xea, 0x4d, 0xd7,
	0x6f, 0x18, 0xee, 0x75, 0x69, 0xa2, 0xaa, 0xcd, 0xaf, 0xf0, 0x91, 0x54, 0x24, 0x53, 0x95, 0x87,
	0xc1, 0x74, 0x88, 0x74, 0x40, 0xd4, 0x04, 0x44, 0xdf, 0x9b, 0x80, 0x9e, 0x55, 0x63, 0x9f, 0x6a,
	0x6c, 0xb5, 0xe6, 0xee, 0xff, 0x3b, 0xc0, 0xd7, 0x08, 0x73, 0x06, 0xbb, 0x94, 0x47, 0x56, 0xe7,
	0x41, 0x52, 0x63, 0x10, 0x6a, 0x12, 0xe5, 0x91, 0xe6, 0x94, 0xeb, 0xe6, 0x4f, 0xd8, 0x67, 0x7c,
	0xb5, 0x21, 0x8a, 0x09, 0x6e, 0x0d, 0x46, 0x60, 0xdc, 0xf7, 0xbc, 0x52, 0xff, 0xaf, 0x70, 0xde,
	0xc7, 0x4c, 0xad, 0xf3, 0x00, 0x85, 0x22, 0xc1, 0xf5, 0x77, 0xd6, 0x8f, 0x4f, 0x32, 0xfa, 0x85,
	0xd5, 0x2e, 0xa5, 0x12, 0xcd, 0x68, 0x78, 0x29, 0x9c, 0x17, 0x9a, 0x7c, 0x0b, 0x72, 0xfd, 0x3b,
	0xa8, 0x37, 0x3f, 0x9c, 0x6c, 0x70, 0x3c, 0xd9, 0xe0, 0xea, 0x64, 0x83, 0xfd, 0xd9, 0x36, 0x8e,
	0x67, 0xdb, 0xf8, 0x7b, 0xb6, 0x8d, 0x1f, 0x1f, 0x5b, 0x07, 0x96, 0x7c, 0xc9, 0xd9, 0x9c, 0xe1,
	0x70, 0x4d, 0x18, 0xc7, 0xbf, 0xdb, 0xff, 0x8d, 0x3e, 0x15, 0xf4, 0xaa, 0x68, 0x9f, 0x6f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa9, 0x04, 0xe5, 0xa8, 0x5f, 0x02, 0x00, 0x00,
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
	if len(m.Periods) > 0 {
		for iNdEx := len(m.Periods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Periods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUnunifidist(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Period) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintUnunifidist(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.End, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.End):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintUnunifidist(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Start, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Start):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintUnunifidist(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintUnunifidist(dAtA []byte, offset int, v uint64) int {
	offset -= sovUnunifidist(v)
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
	if m.Active {
		n += 2
	}
	if len(m.Periods) > 0 {
		for _, e := range m.Periods {
			l = e.Size()
			n += 1 + l + sovUnunifidist(uint64(l))
		}
	}
	return n
}

func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)
	n += 1 + l + sovUnunifidist(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.End)
	n += 1 + l + sovUnunifidist(uint64(l))
	l = m.Inflation.Size()
	n += 1 + l + sovUnunifidist(uint64(l))
	return n
}

func sovUnunifidist(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUnunifidist(x uint64) (n int) {
	return sovUnunifidist(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnunifidist
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnunifidist
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
			m.Active = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Periods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnunifidist
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
				return ErrInvalidLengthUnunifidist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnunifidist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Periods = append(m.Periods, Period{})
			if err := m.Periods[len(m.Periods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnunifidist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnunifidist
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
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnunifidist
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnunifidist
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
				return ErrInvalidLengthUnunifidist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnunifidist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Start, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnunifidist
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
				return ErrInvalidLengthUnunifidist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUnunifidist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.End, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnunifidist
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
				return ErrInvalidLengthUnunifidist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnunifidist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUnunifidist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnunifidist
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
func skipUnunifidist(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUnunifidist
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
					return 0, ErrIntOverflowUnunifidist
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
					return 0, ErrIntOverflowUnunifidist
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
				return 0, ErrInvalidLengthUnunifidist
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUnunifidist
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUnunifidist
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUnunifidist        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUnunifidist          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUnunifidist = fmt.Errorf("proto: unexpected end of group")
)

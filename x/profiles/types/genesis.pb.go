// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the profiles module's genesis state.
type GenesisState struct {
	DTagTransferRequests []DTagTransferRequest `protobuf:"bytes,1,rep,name=dtag_transfer_requests,json=dtagTransferRequests,proto3" json:"dtag_transfer_requests" yaml:"dtag_transfer_requests"`
	Relationships        []Relationship        `protobuf:"bytes,2,rep,name=relationships,proto3" json:"relationships" yaml:"users_relationships"`
	Blocks               []UserBlock           `protobuf:"bytes,3,rep,name=blocks,proto3" json:"blocks" yaml:"users_blocks"`
	Params               Params                `protobuf:"bytes,4,opt,name=params,proto3" json:"params" yaml:"params"`
	IBCPortID            string                `protobuf:"bytes,5,opt,name=ibc_port_id,json=ibcPortId,proto3" json:"ibc_port_id,omitempty" yaml:"ibc_port_id"`
	ChainLinks           []ChainLink           `protobuf:"bytes,6,rep,name=chain_links,json=chainLinks,proto3" json:"chain_links"`
	ApplicationLinks     []ApplicationLink     `protobuf:"bytes,7,rep,name=application_links,json=applicationLinks,proto3" json:"application_links"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_24f308c1ccf2d582, []int{0}
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

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.profiles.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/genesis.proto", fileDescriptor_24f308c1ccf2d582)
}

var fileDescriptor_24f308c1ccf2d582 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x1b, 0xb6, 0x15, 0xd5, 0x65, 0x12, 0x84, 0x02, 0x51, 0x11, 0x49, 0x15, 0x34, 0xa9,
	0x12, 0x23, 0x61, 0xe5, 0x36, 0x89, 0x03, 0xe9, 0x24, 0xa8, 0x84, 0xd0, 0x14, 0x86, 0x90, 0xe0,
	0x50, 0x39, 0xa9, 0x97, 0x5a, 0x4d, 0x63, 0x63, 0xbb, 0x88, 0xbd, 0x01, 0x47, 0x1e, 0x01, 0x9e,
	0x84, 0xeb, 0x8e, 0x3b, 0x72, 0x8a, 0x50, 0xfa, 0x06, 0x7d, 0x02, 0x94, 0xd8, 0x65, 0xe9, 0xb4,
	0x28, 0xdc, 0x22, 0xe7, 0xf7, 0xfd, 0xbe, 0x7f, 0x1c, 0x1b, 0xec, 0x4d, 0x10, 0x9f, 0x13, 0xee,
	0x52, 0x46, 0x4e, 0x71, 0x8c, 0xb8, 0xfb, 0xe5, 0x20, 0x40, 0x02, 0x1e, 0xb8, 0x11, 0x4a, 0x10,
	0xc7, 0xdc, 0xa1, 0x8c, 0x08, 0xa2, 0x3f, 0x90, 0x98, 0xb3, 0xc6, 0x1c, 0x85, 0x75, 0x3b, 0x11,
	0x89, 0x48, 0xc1, 0xb8, 0xf9, 0x93, 0xc4, 0xbb, 0x4f, 0xaa, 0xac, 0x73, 0x32, 0x41, 0x31, 0x1f,
	0x53, 0xc8, 0xe0, 0x5c, 0xb9, 0xbb, 0xfb, 0x75, 0xb0, 0x5c, 0x57, 0xf4, 0xa0, 0x86, 0x66, 0x28,
	0x86, 0x02, 0x93, 0x84, 0x4f, 0x31, 0xe5, 0xff, 0x99, 0x99, 0x08, 0x18, 0x8d, 0x19, 0xfa, 0xbc,
	0x40, 0x5c, 0xac, 0x33, 0xcf, 0x6a, 0x32, 0xe1, 0x14, 0xe2, 0x64, 0x1c, 0xe3, 0x64, 0xb6, 0x4e,
	0x38, 0x35, 0x09, 0x48, 0x69, 0x99, 0xb7, 0x7f, 0xed, 0x80, 0x5b, 0xaf, 0xe4, 0x2e, 0xbf, 0x13,
	0x50, 0x20, 0xfd, 0xa7, 0x06, 0xee, 0x17, 0xa3, 0x08, 0x06, 0x13, 0x7e, 0x8a, 0xd8, 0xbf, 0x99,
	0x0c, 0xad, 0xb7, 0xd5, 0x6f, 0x0f, 0xf6, 0x9d, 0x8a, 0xdf, 0xe0, 0x1c, 0x9d, 0xc0, 0xe8, 0x44,
	0xa5, 0x7c, 0x19, 0xf2, 0x5e, 0x9c, 0xa7, 0x56, 0x23, 0x4b, 0xad, 0xce, 0x35, 0x2f, 0xf9, 0x2a,
	0xb5, 0x1e, 0x9d, 0xc1, 0x79, 0x7c, 0x68, 0x5f, 0xdf, 0x68, 0xfb, 0x9d, 0xfc, 0xc5, 0xd5, 0x98,
	0x4e, 0xc0, 0xee, 0xc6, 0x0e, 0x1b, 0x37, 0x8a, 0xc9, 0xf6, 0x2a, 0x27, 0xf3, 0x4b, 0xb4, 0x67,
	0xe7, 0x23, 0xad, 0x52, 0xab, 0x2b, 0xab, 0x17, 0x1c, 0xb1, 0x2b, 0x7f, 0xcc, 0xf6, 0x37, 0xfd,
	0xfa, 0x07, 0xd0, 0x0c, 0x62, 0x12, 0xce, 0xb8, 0xb1, 0x55, 0x34, 0xd9, 0x95, 0x4d, 0xef, 0x39,
	0x62, 0x5e, 0x8e, 0x7a, 0x0f, 0x55, 0xcd, 0xdd, 0x72, 0x8d, 0xb4, 0xd8, 0xbe, 0xd2, 0xe9, 0x6f,
	0x41, 0x53, 0x1e, 0x43, 0x63, 0xbb, 0xa7, 0xf5, 0xdb, 0x03, 0xab, 0x52, 0x7c, 0x5c, 0x60, 0xde,
	0x3d, 0x65, 0xdd, 0x95, 0x56, 0x19, 0xb6, 0x7d, 0x65, 0xd1, 0x87, 0xa0, 0x8d, 0x83, 0x70, 0x4c,
	0x09, 0x13, 0x63, 0x3c, 0x31, 0x76, 0x7a, 0x5a, 0xbf, 0xe5, 0x3d, 0xce, 0x52, 0xab, 0x35, 0xf2,
	0x86, 0xc7, 0x84, 0x89, 0xd1, 0xd1, 0x2a, 0xb5, 0x74, 0x19, 0x2e, 0x91, 0xb6, 0xdf, 0xc2, 0x41,
	0x58, 0x00, 0x13, 0x7d, 0x04, 0xda, 0xa5, 0x83, 0x65, 0x34, 0x6b, 0x3e, 0x79, 0x98, 0xb3, 0x6f,
	0x70, 0x32, 0xf3, 0xb6, 0xf3, 0xe1, 0x7c, 0x10, 0xae, 0x17, 0xb8, 0xfe, 0x09, 0xdc, 0x81, 0x94,
	0xc6, 0x38, 0x2c, 0x36, 0x53, 0x09, 0x6f, 0x16, 0xc2, 0x7e, 0xa5, 0xf0, 0xe5, 0x65, 0xa2, 0xa4,
	0xbd, 0x0d, 0x37, 0x97, 0xf9, 0xe1, 0xf6, 0xb7, 0x1f, 0x56, 0xc3, 0x7b, 0x7d, 0x9e, 0x99, 0xda,
	0x45, 0x66, 0x6a, 0x7f, 0x32, 0x53, 0xfb, 0xbe, 0x34, 0x1b, 0x17, 0x4b, 0xb3, 0xf1, 0x7b, 0x69,
	0x36, 0x3e, 0x3a, 0x11, 0x16, 0xd3, 0x45, 0xe0, 0x84, 0x64, 0xee, 0xca, 0xae, 0xa7, 0x31, 0x0c,
	0xb8, 0x7a, 0x76, 0xbf, 0x5e, 0x5e, 0x12, 0x71, 0x46, 0x11, 0x0f, 0x9a, 0xc5, 0x95, 0x78, 0xfe,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x8b, 0xfc, 0x9e, 0x8f, 0x04, 0x00, 0x00,
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
	if len(m.ApplicationLinks) > 0 {
		for iNdEx := len(m.ApplicationLinks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApplicationLinks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ChainLinks) > 0 {
		for iNdEx := len(m.ChainLinks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainLinks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.IBCPortID) > 0 {
		i -= len(m.IBCPortID)
		copy(dAtA[i:], m.IBCPortID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IBCPortID)))
		i--
		dAtA[i] = 0x2a
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
	dAtA[i] = 0x22
	if len(m.Blocks) > 0 {
		for iNdEx := len(m.Blocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Blocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Relationships) > 0 {
		for iNdEx := len(m.Relationships) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Relationships[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DTagTransferRequests) > 0 {
		for iNdEx := len(m.DTagTransferRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DTagTransferRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	if len(m.DTagTransferRequests) > 0 {
		for _, e := range m.DTagTransferRequests {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Relationships) > 0 {
		for _, e := range m.Relationships {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Blocks) > 0 {
		for _, e := range m.Blocks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.IBCPortID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ChainLinks) > 0 {
		for _, e := range m.ChainLinks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ApplicationLinks) > 0 {
		for _, e := range m.ApplicationLinks {
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
				return fmt.Errorf("proto: wrong wireType = %d for field DTagTransferRequests", wireType)
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
			m.DTagTransferRequests = append(m.DTagTransferRequests, DTagTransferRequest{})
			if err := m.DTagTransferRequests[len(m.DTagTransferRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relationships", wireType)
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
			m.Relationships = append(m.Relationships, Relationship{})
			if err := m.Relationships[len(m.Relationships)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
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
			m.Blocks = append(m.Blocks, UserBlock{})
			if err := m.Blocks[len(m.Blocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IBCPortID", wireType)
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
			m.IBCPortID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainLinks", wireType)
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
			m.ChainLinks = append(m.ChainLinks, ChainLink{})
			if err := m.ChainLinks[len(m.ChainLinks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationLinks", wireType)
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
			m.ApplicationLinks = append(m.ApplicationLinks, ApplicationLink{})
			if err := m.ApplicationLinks[len(m.ApplicationLinks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

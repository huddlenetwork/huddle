// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/subspaces/v3/models_feegrant.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// Grant represents a grant to a user or a group
type Grant struct {
	// Id of the subspace inside which the user was granted the allowance
	SubspaceID uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty" yaml:"subspace_id"`
	// Address of the user that granted the allowance
	Granter string `protobuf:"bytes,2,opt,name=granter,proto3" json:"granter,omitempty" yaml:"granter"`
	// Target to which the allowance has been granted
	Grantee *types.Any `protobuf:"bytes,3,opt,name=grantee,proto3" json:"grantee,omitempty" yaml:"grantee"`
	// Allowance can be any allowance type implementing the FeeAllowanceI
	// interface
	Allowance *types.Any `protobuf:"bytes,4,opt,name=allowance,proto3" json:"allowance,omitempty" yaml:"allowance"`
}

func (m *Grant) Reset()         { *m = Grant{} }
func (m *Grant) String() string { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()    {}
func (*Grant) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b3327e5773ec37e, []int{0}
}
func (m *Grant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Grant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Grant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Grant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grant.Merge(m, src)
}
func (m *Grant) XXX_Size() int {
	return m.Size()
}
func (m *Grant) XXX_DiscardUnknown() {
	xxx_messageInfo_Grant.DiscardUnknown(m)
}

var xxx_messageInfo_Grant proto.InternalMessageInfo

func (m *Grant) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *Grant) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *Grant) GetGrantee() *types.Any {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *Grant) GetAllowance() *types.Any {
	if m != nil {
		return m.Allowance
	}
	return nil
}

// UserGrantee contains the target of a grant about a user
type UserGrantee struct {
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" yaml:"user"`
}

func (m *UserGrantee) Reset()         { *m = UserGrantee{} }
func (m *UserGrantee) String() string { return proto.CompactTextString(m) }
func (*UserGrantee) ProtoMessage()    {}
func (*UserGrantee) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b3327e5773ec37e, []int{1}
}
func (m *UserGrantee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserGrantee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserGrantee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserGrantee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGrantee.Merge(m, src)
}
func (m *UserGrantee) XXX_Size() int {
	return m.Size()
}
func (m *UserGrantee) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGrantee.DiscardUnknown(m)
}

var xxx_messageInfo_UserGrantee proto.InternalMessageInfo

func (m *UserGrantee) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

// GroupGrantee contains the target of a grant about a group
type GroupGrantee struct {
	GroupID uint32 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty" yaml:"group_id"`
}

func (m *GroupGrantee) Reset()         { *m = GroupGrantee{} }
func (m *GroupGrantee) String() string { return proto.CompactTextString(m) }
func (*GroupGrantee) ProtoMessage()    {}
func (*GroupGrantee) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b3327e5773ec37e, []int{2}
}
func (m *GroupGrantee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupGrantee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupGrantee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupGrantee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupGrantee.Merge(m, src)
}
func (m *GroupGrantee) XXX_Size() int {
	return m.Size()
}
func (m *GroupGrantee) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupGrantee.DiscardUnknown(m)
}

var xxx_messageInfo_GroupGrantee proto.InternalMessageInfo

func (m *GroupGrantee) GetGroupID() uint32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func init() {
	proto.RegisterType((*Grant)(nil), "desmos.subspaces.v3.Grant")
	proto.RegisterType((*UserGrantee)(nil), "desmos.subspaces.v3.UserGrantee")
	proto.RegisterType((*GroupGrantee)(nil), "desmos.subspaces.v3.GroupGrantee")
}

func init() {
	proto.RegisterFile("desmos/subspaces/v3/models_feegrant.proto", fileDescriptor_7b3327e5773ec37e)
}

var fileDescriptor_7b3327e5773ec37e = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x25, 0x90, 0xe6, 0x42, 0x29, 0xba, 0x76, 0x08, 0x95, 0xf0, 0x45, 0x07, 0x48,
	0x41, 0xa2, 0x3e, 0xa9, 0x61, 0x21, 0x5b, 0x2d, 0x20, 0x8a, 0x04, 0x4b, 0x10, 0x03, 0x0c, 0x44,
	0x76, 0xfc, 0x6a, 0x2a, 0xd9, 0x3e, 0xcb, 0x67, 0x07, 0xfc, 0x2d, 0x18, 0x19, 0xfb, 0x11, 0x18,
	0xf8, 0x10, 0x88, 0xa9, 0x62, 0x62, 0xb2, 0x90, 0xb3, 0x30, 0xfb, 0x13, 0x20, 0xdf, 0xf9, 0xda,
	0x0a, 0x21, 0xb6, 0x7b, 0xef, 0xff, 0x7f, 0xbf, 0xe7, 0xf7, 0x97, 0xf1, 0x43, 0x1f, 0x64, 0x24,
	0x24, 0x97, 0xb9, 0x27, 0x13, 0x77, 0x05, 0x92, 0xaf, 0x27, 0x3c, 0x12, 0x3e, 0x84, 0x72, 0x79,
	0x02, 0x10, 0xa4, 0x6e, 0x9c, 0xd9, 0x49, 0x2a, 0x32, 0x41, 0xf6, 0xb4, 0xd5, 0xbe, 0xb0, 0xda,
	0xeb, 0xc9, 0xc1, 0x7e, 0x20, 0x02, 0xa1, 0x74, 0xde, 0xbc, 0xb4, 0xf5, 0xe0, 0x4e, 0x20, 0x44,
	0x10, 0x02, 0x57, 0x95, 0x97, 0x9f, 0x70, 0x37, 0x2e, 0x8c, 0xb4, 0x12, 0x0d, 0x65, 0xa9, 0x67,
	0x74, 0xa1, 0x25, 0xf6, 0x65, 0x0b, 0x5f, 0x9f, 0x35, 0x0b, 0xc9, 0x33, 0x3c, 0x30, 0x5b, 0x96,
	0xa7, 0xfe, 0x10, 0x8d, 0xd0, 0xb8, 0xeb, 0xdc, 0xaf, 0x4a, 0x8a, 0x5f, 0xb5, 0xed, 0xf9, 0xd3,
	0xba, 0xa4, 0xa4, 0x70, 0xa3, 0x70, 0xca, 0xae, 0x58, 0xd9, 0x02, 0x9b, 0x6a, 0xee, 0x93, 0x47,
	0xb8, 0xa7, 0x0e, 0x80, 0x74, 0xb8, 0x35, 0x42, 0xe3, 0xbe, 0x43, 0xea, 0x92, 0xde, 0xd2, 0x43,
	0xad, 0xc0, 0x16, 0xc6, 0x42, 0x5e, 0x1a, 0x37, 0x0c, 0xaf, 0x8d, 0xd0, 0x78, 0x70, 0xb4, 0x6f,
	0xeb, 0x33, 0x6c, 0x73, 0x86, 0x7d, 0x1c, 0x17, 0xce, 0xdd, 0xbf, 0x19, 0xc0, 0xbe, 0x7f, 0x3d,
	0xec, 0xcd, 0xf4, 0xdb, 0xe0, 0x80, 0xbc, 0xc1, 0x7d, 0x37, 0x0c, 0xc5, 0x07, 0x37, 0x5e, 0xc1,
	0xb0, 0xfb, 0x1f, 0xe0, 0x83, 0xba, 0xa4, 0xb7, 0x35, 0xf0, 0x62, 0xa0, 0x41, 0xee, 0x3c, 0x07,
	0x38, 0x36, 0x8d, 0xf9, 0xe2, 0x92, 0x36, 0xdd, 0xfe, 0x7c, 0x46, 0xd1, 0xef, 0x33, 0x8a, 0xd8,
	0x0c, 0x0f, 0x5e, 0x4b, 0x48, 0xdb, 0xe5, 0xe4, 0x1e, 0xee, 0xe6, 0x12, 0x52, 0x15, 0x58, 0xdf,
	0xd9, 0xad, 0x4b, 0x3a, 0xd0, 0xe0, 0xa6, 0xcb, 0x16, 0x4a, 0x9c, 0xee, 0x99, 0xe9, 0x1f, 0x97,
	0x9f, 0xcd, 0xde, 0xe1, 0x9b, 0xb3, 0x54, 0xe4, 0x89, 0x21, 0x3d, 0xc1, 0xdb, 0x41, 0x53, 0x9b,
	0xf8, 0x77, 0x1c, 0xab, 0x2a, 0x69, 0x4f, 0x79, 0x54, 0xf6, 0xbb, 0x26, 0x02, 0x6d, 0x52, 0x39,
	0x36, 0x9a, 0xff, 0x4f, 0xbe, 0xf3, 0xe2, 0x5b, 0x65, 0xa1, 0xf3, 0xca, 0x42, 0xbf, 0x2a, 0x0b,
	0x7d, 0xda, 0x58, 0x9d, 0xf3, 0x8d, 0xd5, 0xf9, 0xb9, 0xb1, 0x3a, 0x6f, 0x8f, 0x82, 0xd3, 0xec,
	0x7d, 0xee, 0xd9, 0x2b, 0x11, 0x71, 0xfd, 0x87, 0x1d, 0x86, 0xae, 0x27, 0xdb, 0x37, 0x5f, 0x3f,
	0xe6, 0x1f, 0xaf, 0xfc, 0x9d, 0x59, 0x91, 0x80, 0xf4, 0x6e, 0xa8, 0x00, 0x27, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0xe1, 0x81, 0x55, 0xbe, 0x02, 0x00, 0x00,
}

func (this *Grant) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Grant)
	if !ok {
		that2, ok := that.(Grant)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SubspaceID != that1.SubspaceID {
		return false
	}
	if this.Granter != that1.Granter {
		return false
	}
	if !this.Grantee.Equal(that1.Grantee) {
		return false
	}
	if !this.Allowance.Equal(that1.Allowance) {
		return false
	}
	return true
}
func (this *UserGrantee) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserGrantee)
	if !ok {
		that2, ok := that.(UserGrantee)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.User != that1.User {
		return false
	}
	return true
}
func (this *GroupGrantee) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GroupGrantee)
	if !ok {
		that2, ok := that.(GroupGrantee)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.GroupID != that1.GroupID {
		return false
	}
	return true
}
func (m *Grant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Grant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Grant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModelsFeegrant(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Grantee != nil {
		{
			size, err := m.Grantee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModelsFeegrant(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintModelsFeegrant(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0x12
	}
	if m.SubspaceID != 0 {
		i = encodeVarintModelsFeegrant(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserGrantee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserGrantee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserGrantee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintModelsFeegrant(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GroupGrantee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupGrantee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupGrantee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GroupID != 0 {
		i = encodeVarintModelsFeegrant(dAtA, i, uint64(m.GroupID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsFeegrant(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsFeegrant(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Grant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovModelsFeegrant(uint64(m.SubspaceID))
	}
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovModelsFeegrant(uint64(l))
	}
	if m.Grantee != nil {
		l = m.Grantee.Size()
		n += 1 + l + sovModelsFeegrant(uint64(l))
	}
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovModelsFeegrant(uint64(l))
	}
	return n
}

func (m *UserGrantee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovModelsFeegrant(uint64(l))
	}
	return n
}

func (m *GroupGrantee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupID != 0 {
		n += 1 + sovModelsFeegrant(uint64(m.GroupID))
	}
	return n
}

func sovModelsFeegrant(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsFeegrant(x uint64) (n int) {
	return sovModelsFeegrant(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Grant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsFeegrant
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
			return fmt.Errorf("proto: Grant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Grant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsFeegrant
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
				return ErrInvalidLengthModelsFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsFeegrant
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
				return ErrInvalidLengthModelsFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Grantee == nil {
				m.Grantee = &types.Any{}
			}
			if err := m.Grantee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsFeegrant
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
				return ErrInvalidLengthModelsFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &types.Any{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsFeegrant
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
func (m *UserGrantee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsFeegrant
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
			return fmt.Errorf("proto: UserGrantee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserGrantee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsFeegrant
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
				return ErrInvalidLengthModelsFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsFeegrant
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
func (m *GroupGrantee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsFeegrant
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
			return fmt.Errorf("proto: GroupGrantee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupGrantee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModelsFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsFeegrant
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
func skipModelsFeegrant(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsFeegrant
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
					return 0, ErrIntOverflowModelsFeegrant
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
					return 0, ErrIntOverflowModelsFeegrant
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
				return 0, ErrInvalidLengthModelsFeegrant
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsFeegrant
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsFeegrant
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsFeegrant        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsFeegrant          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsFeegrant = fmt.Errorf("proto: unexpected end of group")
)

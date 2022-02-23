// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/subspaces/v1/genesis.proto

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

// GenesisState contains the data of the genesis state for the subspaces module
type GenesisState struct {
	InitialSubspaceID uint64                  `protobuf:"varint,1,opt,name=initial_subspace_id,json=initialSubspaceId,proto3" json:"initial_subspace_id,omitempty"`
	Subspaces         []GenesisSubspace       `protobuf:"bytes,2,rep,name=subspaces,proto3" json:"subspaces"`
	ACL               []ACLEntry              `protobuf:"bytes,3,rep,name=acl,proto3" json:"acl"`
	UserGroups        []UserGroup             `protobuf:"bytes,4,rep,name=user_groups,json=userGroups,proto3" json:"user_groups"`
	UserGroupsMembers []UserGroupMembersEntry `protobuf:"bytes,5,rep,name=user_groups_members,json=userGroupsMembers,proto3" json:"user_groups_members"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bce8af665337782b, []int{0}
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

func (m *GenesisState) GetInitialSubspaceID() uint64 {
	if m != nil {
		return m.InitialSubspaceID
	}
	return 0
}

func (m *GenesisState) GetSubspaces() []GenesisSubspace {
	if m != nil {
		return m.Subspaces
	}
	return nil
}

func (m *GenesisState) GetACL() []ACLEntry {
	if m != nil {
		return m.ACL
	}
	return nil
}

func (m *GenesisState) GetUserGroups() []UserGroup {
	if m != nil {
		return m.UserGroups
	}
	return nil
}

func (m *GenesisState) GetUserGroupsMembers() []UserGroupMembersEntry {
	if m != nil {
		return m.UserGroupsMembers
	}
	return nil
}

// GenesisSubspace contains the genesis data for a single subspace
type GenesisSubspace struct {
	Subspace       Subspace `protobuf:"bytes,1,opt,name=subspace,proto3" json:"subspace"`
	InitialGroupID uint32   `protobuf:"varint,2,opt,name=initial_group_id,json=initialGroupId,proto3" json:"initial_group_id,omitempty"`
}

func (m *GenesisSubspace) Reset()         { *m = GenesisSubspace{} }
func (m *GenesisSubspace) String() string { return proto.CompactTextString(m) }
func (*GenesisSubspace) ProtoMessage()    {}
func (*GenesisSubspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_bce8af665337782b, []int{1}
}
func (m *GenesisSubspace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisSubspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisSubspace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisSubspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisSubspace.Merge(m, src)
}
func (m *GenesisSubspace) XXX_Size() int {
	return m.Size()
}
func (m *GenesisSubspace) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisSubspace.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisSubspace proto.InternalMessageInfo

func (m *GenesisSubspace) GetSubspace() Subspace {
	if m != nil {
		return m.Subspace
	}
	return Subspace{}
}

func (m *GenesisSubspace) GetInitialGroupID() uint32 {
	if m != nil {
		return m.InitialGroupID
	}
	return 0
}

// ACLEntry represents a single Access Control List entry
type ACLEntry struct {
	SubspaceID  uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	User        string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Permissions uint32 `protobuf:"varint,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *ACLEntry) Reset()         { *m = ACLEntry{} }
func (m *ACLEntry) String() string { return proto.CompactTextString(m) }
func (*ACLEntry) ProtoMessage()    {}
func (*ACLEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bce8af665337782b, []int{2}
}
func (m *ACLEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ACLEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ACLEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ACLEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACLEntry.Merge(m, src)
}
func (m *ACLEntry) XXX_Size() int {
	return m.Size()
}
func (m *ACLEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ACLEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ACLEntry proto.InternalMessageInfo

func (m *ACLEntry) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *ACLEntry) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ACLEntry) GetPermissions() uint32 {
	if m != nil {
		return m.Permissions
	}
	return 0
}

// UserGroupMembersEntry contains all the members of a specific user group
type UserGroupMembersEntry struct {
	SubspaceID uint64   `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	GroupID    uint32   `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Members    []string `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
}

func (m *UserGroupMembersEntry) Reset()         { *m = UserGroupMembersEntry{} }
func (m *UserGroupMembersEntry) String() string { return proto.CompactTextString(m) }
func (*UserGroupMembersEntry) ProtoMessage()    {}
func (*UserGroupMembersEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bce8af665337782b, []int{3}
}
func (m *UserGroupMembersEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserGroupMembersEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserGroupMembersEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserGroupMembersEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroupMembersEntry.Merge(m, src)
}
func (m *UserGroupMembersEntry) XXX_Size() int {
	return m.Size()
}
func (m *UserGroupMembersEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroupMembersEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroupMembersEntry proto.InternalMessageInfo

func (m *UserGroupMembersEntry) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *UserGroupMembersEntry) GetGroupID() uint32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *UserGroupMembersEntry) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.subspaces.v1.GenesisState")
	proto.RegisterType((*GenesisSubspace)(nil), "desmos.subspaces.v1.GenesisSubspace")
	proto.RegisterType((*ACLEntry)(nil), "desmos.subspaces.v1.ACLEntry")
	proto.RegisterType((*UserGroupMembersEntry)(nil), "desmos.subspaces.v1.UserGroupMembersEntry")
}

func init() { proto.RegisterFile("desmos/subspaces/v1/genesis.proto", fileDescriptor_bce8af665337782b) }

var fileDescriptor_bce8af665337782b = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0x9b, 0x6a, 0xdb, 0x17, 0xad, 0x76, 0xea, 0x42, 0x58, 0x30, 0xa9, 0x8b, 0x48,
	0x11, 0x4c, 0xd8, 0x7a, 0x93, 0x05, 0xd9, 0xec, 0x96, 0xb5, 0x50, 0x2f, 0x59, 0xbc, 0x78, 0xa9,
	0x49, 0x33, 0xc4, 0x81, 0xa6, 0x09, 0x99, 0xa4, 0xb8, 0xf8, 0x4f, 0xec, 0xc1, 0x83, 0x78, 0xda,
	0x3f, 0x67, 0x8f, 0x7b, 0xf4, 0x14, 0x24, 0xbd, 0xf8, 0x67, 0x48, 0x26, 0x99, 0x36, 0xd4, 0x28,
	0xec, 0x6d, 0xe6, 0xfd, 0xf8, 0xbc, 0xef, 0xbc, 0xf7, 0x06, 0x9e, 0xb9, 0x84, 0xf9, 0x01, 0x33,
	0x58, 0xe2, 0xb0, 0xd0, 0x9e, 0x13, 0x66, 0xac, 0x8e, 0x0c, 0x8f, 0x2c, 0x09, 0xa3, 0x4c, 0x0f,
	0xa3, 0x20, 0x0e, 0x70, 0xbf, 0x08, 0xd1, 0x37, 0x21, 0xfa, 0xea, 0xe8, 0xe0, 0x89, 0x17, 0x78,
	0x01, 0xf7, 0x1b, 0xf9, 0xa9, 0x08, 0x3d, 0x18, 0xd4, 0xd1, 0xfc, 0xc0, 0x25, 0x8b, 0x12, 0x76,
	0x78, 0x25, 0xc1, 0x83, 0xf3, 0x02, 0x7f, 0x11, 0xdb, 0x31, 0xc1, 0x63, 0xe8, 0xd3, 0x25, 0x8d,
	0xa9, 0xbd, 0x98, 0x89, 0xac, 0x19, 0x75, 0x15, 0x34, 0x40, 0xc3, 0xa6, 0xb9, 0x9f, 0xa5, 0x5a,
	0x6f, 0x52, 0xb8, 0x2f, 0x4a, 0xef, 0xe4, 0xcc, 0xea, 0xd1, 0x1d, 0x93, 0x8b, 0xdf, 0x41, 0x67,
	0x53, 0x54, 0xd9, 0x1b, 0x48, 0x43, 0x79, 0xf4, 0x5c, 0xaf, 0x11, 0xae, 0x8b, 0xe2, 0xa5, 0xcd,
	0x6c, 0xde, 0xa4, 0x5a, 0xc3, 0xda, 0x26, 0xe3, 0x63, 0x90, 0xec, 0xf9, 0x42, 0x91, 0x38, 0xe3,
	0x69, 0x2d, 0xe3, 0xe4, 0x74, 0x3a, 0x5e, 0xc6, 0xd1, 0xa5, 0x29, 0xe7, 0xc9, 0x59, 0xaa, 0x49,
	0x27, 0xa7, 0x53, 0x2b, 0x4f, 0xc3, 0x63, 0x90, 0x13, 0x46, 0xa2, 0x99, 0x17, 0x05, 0x49, 0xc8,
	0x94, 0x26, 0xa7, 0xa8, 0xb5, 0x94, 0x0f, 0x8c, 0x44, 0xe7, 0x79, 0x58, 0xa9, 0x01, 0x12, 0x61,
	0x60, 0xf8, 0x13, 0xf4, 0x2b, 0x98, 0x99, 0x4f, 0x7c, 0x87, 0x44, 0x4c, 0xb9, 0xc7, 0x71, 0x2f,
	0xff, 0x8f, 0x7b, 0x5f, 0x04, 0x17, 0x0a, 0x0b, 0x74, 0x6f, 0x8b, 0x2e, 0xbd, 0x6f, 0xda, 0xdf,
	0xaf, 0x35, 0xf4, 0xfb, 0x5a, 0x43, 0x87, 0x3f, 0x10, 0x3c, 0xda, 0xe9, 0x0a, 0x7e, 0x0b, 0x6d,
	0x01, 0xe7, 0xa3, 0xf8, 0x57, 0x27, 0x76, 0xda, 0xb8, 0x49, 0xc2, 0xc7, 0xf0, 0x58, 0x8c, 0x95,
	0xbf, 0x21, 0x9f, 0xe9, 0xde, 0x00, 0x0d, 0x1f, 0x9a, 0x38, 0x4b, 0xb5, 0x6e, 0x39, 0x53, 0x2e,
	0x69, 0x72, 0x66, 0x75, 0x69, 0xf5, 0xee, 0x56, 0xc4, 0x7d, 0x85, 0xb6, 0xe8, 0x36, 0x36, 0x40,
	0xfe, 0x7b, 0x45, 0xba, 0x59, 0xaa, 0x41, 0x65, 0x37, 0x80, 0x6d, 0x97, 0x02, 0x43, 0x33, 0x7f,
	0x38, 0x2f, 0xdc, 0xb1, 0xf8, 0x19, 0x0f, 0x40, 0x0e, 0x49, 0xe4, 0x53, 0xc6, 0x68, 0xb0, 0x64,
	0x8a, 0x94, 0x6b, 0xb2, 0xaa, 0xa6, 0x4a, 0xf1, 0x6f, 0x08, 0xf6, 0x6b, 0xdb, 0x7a, 0x77, 0x29,
	0x2f, 0xa0, 0xbd, 0xd3, 0x07, 0x39, 0x4b, 0xb5, 0x96, 0x68, 0x40, 0xcb, 0x2b, 0x5e, 0x8e, 0x15,
	0x68, 0x89, 0x61, 0xe7, 0x1b, 0xd8, 0xb1, 0xc4, 0x75, 0x2b, 0xcb, 0x9c, 0xde, 0x64, 0x2a, 0xba,
	0xcd, 0x54, 0xf4, 0x2b, 0x53, 0xd1, 0xd5, 0x5a, 0x6d, 0xdc, 0xae, 0xd5, 0xc6, 0xcf, 0xb5, 0xda,
	0xf8, 0x38, 0xf2, 0x68, 0xfc, 0x39, 0x71, 0xf4, 0x79, 0xe0, 0x1b, 0xc5, 0xb8, 0x5e, 0x2d, 0x6c,
	0x87, 0x95, 0x67, 0x63, 0x35, 0x32, 0xbe, 0x54, 0xfe, 0x66, 0x7c, 0x19, 0x12, 0xe6, 0xdc, 0xe7,
	0x1f, 0xf3, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0xfd, 0x5d, 0x3a, 0x0a, 0x04, 0x00,
	0x00,
}

func (this *GenesisState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
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
	if this.InitialSubspaceID != that1.InitialSubspaceID {
		return false
	}
	if len(this.Subspaces) != len(that1.Subspaces) {
		return false
	}
	for i := range this.Subspaces {
		if !this.Subspaces[i].Equal(&that1.Subspaces[i]) {
			return false
		}
	}
	if len(this.ACL) != len(that1.ACL) {
		return false
	}
	for i := range this.ACL {
		if !this.ACL[i].Equal(&that1.ACL[i]) {
			return false
		}
	}
	if len(this.UserGroups) != len(that1.UserGroups) {
		return false
	}
	for i := range this.UserGroups {
		if !this.UserGroups[i].Equal(&that1.UserGroups[i]) {
			return false
		}
	}
	if len(this.UserGroupsMembers) != len(that1.UserGroupsMembers) {
		return false
	}
	for i := range this.UserGroupsMembers {
		if !this.UserGroupsMembers[i].Equal(&that1.UserGroupsMembers[i]) {
			return false
		}
	}
	return true
}
func (this *GenesisSubspace) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisSubspace)
	if !ok {
		that2, ok := that.(GenesisSubspace)
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
	if !this.Subspace.Equal(&that1.Subspace) {
		return false
	}
	if this.InitialGroupID != that1.InitialGroupID {
		return false
	}
	return true
}
func (this *ACLEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ACLEntry)
	if !ok {
		that2, ok := that.(ACLEntry)
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
	if this.User != that1.User {
		return false
	}
	if this.Permissions != that1.Permissions {
		return false
	}
	return true
}
func (this *UserGroupMembersEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserGroupMembersEntry)
	if !ok {
		that2, ok := that.(UserGroupMembersEntry)
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
	if this.GroupID != that1.GroupID {
		return false
	}
	if len(this.Members) != len(that1.Members) {
		return false
	}
	for i := range this.Members {
		if this.Members[i] != that1.Members[i] {
			return false
		}
	}
	return true
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
	if len(m.UserGroupsMembers) > 0 {
		for iNdEx := len(m.UserGroupsMembers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserGroupsMembers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.UserGroups) > 0 {
		for iNdEx := len(m.UserGroups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserGroups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ACL) > 0 {
		for iNdEx := len(m.ACL) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ACL[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Subspaces) > 0 {
		for iNdEx := len(m.Subspaces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subspaces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.InitialSubspaceID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.InitialSubspaceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GenesisSubspace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisSubspace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisSubspace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InitialGroupID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.InitialGroupID))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Subspace.MarshalToSizedBuffer(dAtA[:i])
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

func (m *ACLEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ACLEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ACLEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Permissions != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Permissions))
		i--
		dAtA[i] = 0x18
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x12
	}
	if m.SubspaceID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserGroupMembersEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserGroupMembersEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserGroupMembersEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Members[iNdEx])
			copy(dAtA[i:], m.Members[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.Members[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.GroupID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.GroupID))
		i--
		dAtA[i] = 0x10
	}
	if m.SubspaceID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
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
	if m.InitialSubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.InitialSubspaceID))
	}
	if len(m.Subspaces) > 0 {
		for _, e := range m.Subspaces {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ACL) > 0 {
		for _, e := range m.ACL {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserGroups) > 0 {
		for _, e := range m.UserGroups {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserGroupsMembers) > 0 {
		for _, e := range m.UserGroupsMembers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisSubspace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Subspace.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.InitialGroupID != 0 {
		n += 1 + sovGenesis(uint64(m.InitialGroupID))
	}
	return n
}

func (m *ACLEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.SubspaceID))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Permissions != 0 {
		n += 1 + sovGenesis(uint64(m.Permissions))
	}
	return n
}

func (m *UserGroupMembersEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.SubspaceID))
	}
	if m.GroupID != 0 {
		n += 1 + sovGenesis(uint64(m.GroupID))
	}
	if len(m.Members) > 0 {
		for _, s := range m.Members {
			l = len(s)
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialSubspaceID", wireType)
			}
			m.InitialSubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialSubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspaces", wireType)
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
			m.Subspaces = append(m.Subspaces, GenesisSubspace{})
			if err := m.Subspaces[len(m.Subspaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ACL", wireType)
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
			m.ACL = append(m.ACL, ACLEntry{})
			if err := m.ACL[len(m.ACL)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserGroups", wireType)
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
			m.UserGroups = append(m.UserGroups, UserGroup{})
			if err := m.UserGroups[len(m.UserGroups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserGroupsMembers", wireType)
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
			m.UserGroupsMembers = append(m.UserGroupsMembers, UserGroupMembersEntry{})
			if err := m.UserGroupsMembers[len(m.UserGroupsMembers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisSubspace) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisSubspace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisSubspace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
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
			if err := m.Subspace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialGroupID", wireType)
			}
			m.InitialGroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialGroupID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *ACLEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ACLEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ACLEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
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
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			m.Permissions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permissions |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *UserGroupMembersEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UserGroupMembersEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserGroupMembersEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
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
			m.Members = append(m.Members, string(dAtA[iNdEx:postIndex]))
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
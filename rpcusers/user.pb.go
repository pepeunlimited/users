// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package rpcusers

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SetProfilePictureParams struct {
	ProfilePictureId     int64    `protobuf:"varint,1,opt,name=profile_picture_id,json=profilePictureId,proto3" json:"profile_picture_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetProfilePictureParams) Reset()         { *m = SetProfilePictureParams{} }
func (m *SetProfilePictureParams) String() string { return proto.CompactTextString(m) }
func (*SetProfilePictureParams) ProtoMessage()    {}
func (*SetProfilePictureParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *SetProfilePictureParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetProfilePictureParams.Unmarshal(m, b)
}
func (m *SetProfilePictureParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetProfilePictureParams.Marshal(b, m, deterministic)
}
func (m *SetProfilePictureParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetProfilePictureParams.Merge(m, src)
}
func (m *SetProfilePictureParams) XXX_Size() int {
	return xxx_messageInfo_SetProfilePictureParams.Size(m)
}
func (m *SetProfilePictureParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SetProfilePictureParams.DiscardUnknown(m)
}

var xxx_messageInfo_SetProfilePictureParams proto.InternalMessageInfo

func (m *SetProfilePictureParams) GetProfilePictureId() int64 {
	if m != nil {
		return m.ProfilePictureId
	}
	return 0
}

type ProfilePicture struct {
	ProfilePictureId     int64    `protobuf:"varint,1,opt,name=profile_picture_id,json=profilePictureId,proto3" json:"profile_picture_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfilePicture) Reset()         { *m = ProfilePicture{} }
func (m *ProfilePicture) String() string { return proto.CompactTextString(m) }
func (*ProfilePicture) ProtoMessage()    {}
func (*ProfilePicture) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *ProfilePicture) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfilePicture.Unmarshal(m, b)
}
func (m *ProfilePicture) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfilePicture.Marshal(b, m, deterministic)
}
func (m *ProfilePicture) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfilePicture.Merge(m, src)
}
func (m *ProfilePicture) XXX_Size() int {
	return xxx_messageInfo_ProfilePicture.Size(m)
}
func (m *ProfilePicture) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfilePicture.DiscardUnknown(m)
}

var xxx_messageInfo_ProfilePicture proto.InternalMessageInfo

func (m *ProfilePicture) GetProfilePictureId() int64 {
	if m != nil {
		return m.ProfilePictureId
	}
	return 0
}

// CreateUser:CreateUserParams
type CreateUserParams struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserParams) Reset()         { *m = CreateUserParams{} }
func (m *CreateUserParams) String() string { return proto.CompactTextString(m) }
func (*CreateUserParams) ProtoMessage()    {}
func (*CreateUserParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *CreateUserParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserParams.Unmarshal(m, b)
}
func (m *CreateUserParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserParams.Marshal(b, m, deterministic)
}
func (m *CreateUserParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserParams.Merge(m, src)
}
func (m *CreateUserParams) XXX_Size() int {
	return xxx_messageInfo_CreateUserParams.Size(m)
}
func (m *CreateUserParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserParams.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserParams proto.InternalMessageInfo

func (m *CreateUserParams) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserParams) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserParams) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type User struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Roles                []string             `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty"`
	ProfilePictureId     *wrappers.Int64Value `protobuf:"bytes,5,opt,name=profile_picture_id,json=profilePictureId,proto3" json:"profile_picture_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *User) GetProfilePictureId() *wrappers.Int64Value {
	if m != nil {
		return m.ProfilePictureId
	}
	return nil
}

func init() {
	proto.RegisterType((*SetProfilePictureParams)(nil), "pepeunlimited.users.SetProfilePictureParams")
	proto.RegisterType((*ProfilePicture)(nil), "pepeunlimited.users.ProfilePicture")
	proto.RegisterType((*CreateUserParams)(nil), "pepeunlimited.users.CreateUserParams")
	proto.RegisterType((*User)(nil), "pepeunlimited.users.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x5d, 0x4b, 0x02, 0x41,
	0x14, 0x65, 0x77, 0xb5, 0x8f, 0x2b, 0x88, 0x4d, 0x52, 0xdb, 0x0a, 0x21, 0x46, 0xe0, 0x83, 0xac,
	0x60, 0xd1, 0x53, 0xf4, 0xd0, 0x07, 0xe2, 0x4b, 0x88, 0x52, 0x0f, 0xbd, 0xd8, 0xe8, 0x5e, 0x65,
	0x60, 0xd7, 0x1d, 0x66, 0x66, 0x93, 0x7e, 0x52, 0xff, 0xa8, 0x9f, 0x13, 0xb3, 0x93, 0x86, 0xeb,
	0x1a, 0xf5, 0x36, 0xf7, 0x9e, 0x7b, 0xce, 0x9c, 0xb9, 0x73, 0x00, 0x12, 0x89, 0xc2, 0xe7, 0x22,
	0x56, 0x31, 0x39, 0xe4, 0xc8, 0x31, 0x99, 0x87, 0x2c, 0x62, 0x0a, 0x03, 0x5f, 0x23, 0xd2, 0xab,
	0xcd, 0xe2, 0x78, 0x16, 0x62, 0x3b, 0x1d, 0x19, 0x27, 0xd3, 0x36, 0x46, 0x5c, 0xbd, 0x1b, 0x86,
	0x77, 0x9a, 0x05, 0x17, 0x82, 0x72, 0x8e, 0x42, 0x1a, 0xbc, 0xd1, 0x85, 0xe3, 0x21, 0xaa, 0xbe,
	0x88, 0xa7, 0x2c, 0xc4, 0x3e, 0x9b, 0xa8, 0x44, 0x60, 0x9f, 0x0a, 0x1a, 0x49, 0xd2, 0x02, 0xc2,
	0x4d, 0x7f, 0xc4, 0x0d, 0x30, 0x62, 0x81, 0x6b, 0xd5, 0xad, 0xa6, 0x33, 0xa8, 0xf0, 0x35, 0x46,
	0x2f, 0x68, 0xdc, 0x40, 0x79, 0x5d, 0xe5, 0x9f, 0xfc, 0x57, 0xa8, 0xdc, 0x09, 0xa4, 0x0a, 0x9f,
	0x24, 0x8a, 0x6f, 0x07, 0x1e, 0xec, 0xe9, 0x27, 0xce, 0x69, 0x84, 0x29, 0x6f, 0x7f, 0xb0, 0xaa,
	0x35, 0xc6, 0xa9, 0x94, 0x8b, 0x58, 0x04, 0xae, 0x6d, 0xb0, 0x65, 0x4d, 0xaa, 0x50, 0xc4, 0x88,
	0xb2, 0xd0, 0x75, 0x52, 0xc0, 0x14, 0x8d, 0x0f, 0x0b, 0x0a, 0x5a, 0x9c, 0x94, 0xc1, 0x5e, 0x19,
	0xb1, 0x59, 0xb0, 0x76, 0x8d, 0x9d, 0xb9, 0x26, 0x57, 0x4a, 0x77, 0x45, 0x1c, 0xa2, 0x74, 0x0b,
	0x75, 0x47, 0x77, 0xd3, 0x82, 0xf4, 0x72, 0x1f, 0x5c, 0xac, 0x5b, 0xcd, 0x52, 0xa7, 0xe6, 0x9b,
	0x8f, 0xf0, 0x97, 0x1f, 0xe1, 0xf7, 0xe6, 0xea, 0xea, 0xf2, 0x99, 0x86, 0x09, 0x6e, 0x6e, 0xa3,
	0xf3, 0x69, 0x43, 0x49, 0x7b, 0x1d, 0xa2, 0x78, 0x63, 0x13, 0x24, 0x8f, 0x00, 0x3f, 0xdb, 0x21,
	0xe7, 0x7e, 0x4e, 0x0e, 0xfc, 0xec, 0xfa, 0xbc, 0x93, 0xdc, 0xb1, 0x54, 0xe1, 0x1a, 0x76, 0xbb,
	0xa8, 0xd2, 0xe3, 0xd1, 0x86, 0xb3, 0x07, 0x9d, 0x9f, 0xdf, 0xd8, 0x53, 0x38, 0xd8, 0x08, 0x0d,
	0x69, 0xe5, 0xce, 0x6f, 0x09, 0x97, 0x77, 0x96, 0x3b, 0x9d, 0x91, 0x1c, 0x42, 0xf5, 0x1e, 0x43,
	0x54, 0x98, 0xe9, 0x6f, 0xb3, 0xfc, 0x17, 0xd1, 0xdb, 0xe2, 0x8b, 0x23, 0xf8, 0x64, 0xbc, 0x93,
	0x72, 0x2f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x7c, 0x71, 0x0b, 0x5f, 0x03, 0x00, 0x00,
}

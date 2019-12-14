// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type SignInParams struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInParams) Reset()         { *m = SignInParams{} }
func (m *SignInParams) String() string { return proto.CompactTextString(m) }
func (*SignInParams) ProtoMessage()    {}
func (*SignInParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *SignInParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInParams.Unmarshal(m, b)
}
func (m *SignInParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInParams.Marshal(b, m, deterministic)
}
func (m *SignInParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInParams.Merge(m, src)
}
func (m *SignInParams) XXX_Size() int {
	return xxx_messageInfo_SignInParams.Size(m)
}
func (m *SignInParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInParams.DiscardUnknown(m)
}

var xxx_messageInfo_SignInParams proto.InternalMessageInfo

func (m *SignInParams) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignInParams) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
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
	return fileDescriptor_116e343673f7ffaf, []int{1}
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

type GetUserParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserParams) Reset()         { *m = GetUserParams{} }
func (m *GetUserParams) String() string { return proto.CompactTextString(m) }
func (*GetUserParams) ProtoMessage()    {}
func (*GetUserParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *GetUserParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserParams.Unmarshal(m, b)
}
func (m *GetUserParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserParams.Marshal(b, m, deterministic)
}
func (m *GetUserParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserParams.Merge(m, src)
}
func (m *GetUserParams) XXX_Size() int {
	return xxx_messageInfo_GetUserParams.Size(m)
}
func (m *GetUserParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserParams proto.InternalMessageInfo

type UpdatePasswordParams struct {
	CurrentPassword      string   `protobuf:"bytes,1,opt,name=current_password,json=currentPassword,proto3" json:"current_password,omitempty"`
	NewPassword          string   `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePasswordParams) Reset()         { *m = UpdatePasswordParams{} }
func (m *UpdatePasswordParams) String() string { return proto.CompactTextString(m) }
func (*UpdatePasswordParams) ProtoMessage()    {}
func (*UpdatePasswordParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UpdatePasswordParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePasswordParams.Unmarshal(m, b)
}
func (m *UpdatePasswordParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePasswordParams.Marshal(b, m, deterministic)
}
func (m *UpdatePasswordParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePasswordParams.Merge(m, src)
}
func (m *UpdatePasswordParams) XXX_Size() int {
	return xxx_messageInfo_UpdatePasswordParams.Size(m)
}
func (m *UpdatePasswordParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePasswordParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePasswordParams proto.InternalMessageInfo

func (m *UpdatePasswordParams) GetCurrentPassword() string {
	if m != nil {
		return m.CurrentPassword
	}
	return ""
}

func (m *UpdatePasswordParams) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type UpdatePasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePasswordResponse) Reset()         { *m = UpdatePasswordResponse{} }
func (m *UpdatePasswordResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePasswordResponse) ProtoMessage()    {}
func (*UpdatePasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UpdatePasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePasswordResponse.Unmarshal(m, b)
}
func (m *UpdatePasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePasswordResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePasswordResponse.Merge(m, src)
}
func (m *UpdatePasswordResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePasswordResponse.Size(m)
}
func (m *UpdatePasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePasswordResponse proto.InternalMessageInfo

type ForgotPasswordParams struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Language             string   `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgotPasswordParams) Reset()         { *m = ForgotPasswordParams{} }
func (m *ForgotPasswordParams) String() string { return proto.CompactTextString(m) }
func (*ForgotPasswordParams) ProtoMessage()    {}
func (*ForgotPasswordParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *ForgotPasswordParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotPasswordParams.Unmarshal(m, b)
}
func (m *ForgotPasswordParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotPasswordParams.Marshal(b, m, deterministic)
}
func (m *ForgotPasswordParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotPasswordParams.Merge(m, src)
}
func (m *ForgotPasswordParams) XXX_Size() int {
	return xxx_messageInfo_ForgotPasswordParams.Size(m)
}
func (m *ForgotPasswordParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotPasswordParams.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotPasswordParams proto.InternalMessageInfo

func (m *ForgotPasswordParams) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ForgotPasswordParams) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role                 string   `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
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

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type ResetPasswordParams struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordParams) Reset()         { *m = ResetPasswordParams{} }
func (m *ResetPasswordParams) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordParams) ProtoMessage()    {}
func (*ResetPasswordParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *ResetPasswordParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordParams.Unmarshal(m, b)
}
func (m *ResetPasswordParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordParams.Marshal(b, m, deterministic)
}
func (m *ResetPasswordParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordParams.Merge(m, src)
}
func (m *ResetPasswordParams) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordParams.Size(m)
}
func (m *ResetPasswordParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordParams.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordParams proto.InternalMessageInfo

func (m *ResetPasswordParams) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ResetPasswordParams) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ResetPasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordResponse) Reset()         { *m = ResetPasswordResponse{} }
func (m *ResetPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordResponse) ProtoMessage()    {}
func (*ResetPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *ResetPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordResponse.Unmarshal(m, b)
}
func (m *ResetPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordResponse.Marshal(b, m, deterministic)
}
func (m *ResetPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordResponse.Merge(m, src)
}
func (m *ResetPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordResponse.Size(m)
}
func (m *ResetPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordResponse proto.InternalMessageInfo

type VerifyPasswordParams struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyPasswordParams) Reset()         { *m = VerifyPasswordParams{} }
func (m *VerifyPasswordParams) String() string { return proto.CompactTextString(m) }
func (*VerifyPasswordParams) ProtoMessage()    {}
func (*VerifyPasswordParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *VerifyPasswordParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyPasswordParams.Unmarshal(m, b)
}
func (m *VerifyPasswordParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyPasswordParams.Marshal(b, m, deterministic)
}
func (m *VerifyPasswordParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyPasswordParams.Merge(m, src)
}
func (m *VerifyPasswordParams) XXX_Size() int {
	return xxx_messageInfo_VerifyPasswordParams.Size(m)
}
func (m *VerifyPasswordParams) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyPasswordParams.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyPasswordParams proto.InternalMessageInfo

func (m *VerifyPasswordParams) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyPasswordResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyPasswordResponse) Reset()         { *m = VerifyPasswordResponse{} }
func (m *VerifyPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyPasswordResponse) ProtoMessage()    {}
func (*VerifyPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *VerifyPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyPasswordResponse.Unmarshal(m, b)
}
func (m *VerifyPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyPasswordResponse.Marshal(b, m, deterministic)
}
func (m *VerifyPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyPasswordResponse.Merge(m, src)
}
func (m *VerifyPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyPasswordResponse.Size(m)
}
func (m *VerifyPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyPasswordResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SignInParams)(nil), "pepeunlimited.users.SignInParams")
	proto.RegisterType((*CreateUserParams)(nil), "pepeunlimited.users.CreateUserParams")
	proto.RegisterType((*GetUserParams)(nil), "pepeunlimited.users.GetUserParams")
	proto.RegisterType((*UpdatePasswordParams)(nil), "pepeunlimited.users.UpdatePasswordParams")
	proto.RegisterType((*UpdatePasswordResponse)(nil), "pepeunlimited.users.UpdatePasswordResponse")
	proto.RegisterType((*ForgotPasswordParams)(nil), "pepeunlimited.users.ForgotPasswordParams")
	proto.RegisterType((*User)(nil), "pepeunlimited.users.User")
	proto.RegisterType((*ResetPasswordParams)(nil), "pepeunlimited.users.ResetPasswordParams")
	proto.RegisterType((*ResetPasswordResponse)(nil), "pepeunlimited.users.ResetPasswordResponse")
	proto.RegisterType((*VerifyPasswordParams)(nil), "pepeunlimited.users.VerifyPasswordParams")
	proto.RegisterType((*VerifyPasswordResponse)(nil), "pepeunlimited.users.VerifyPasswordResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0xa5, 0xfb, 0x51, 0xf5, 0xb6, 0xdd, 0x96, 0xd9, 0x58, 0x63, 0x04, 0xb1, 0x03, 0x42, 0xab,
	0x92, 0x82, 0xfe, 0x03, 0xc5, 0xad, 0xbe, 0x94, 0xb2, 0x4b, 0x7d, 0x10, 0xa1, 0x4e, 0x37, 0xb7,
	0x31, 0x98, 0xcc, 0x0c, 0x33, 0x13, 0x97, 0xbe, 0xfb, 0xc3, 0x25, 0x99, 0x64, 0xeb, 0x84, 0xd9,
	0xed, 0x42, 0xdf, 0x72, 0x73, 0xcf, 0x3d, 0xe7, 0x7e, 0x0e, 0x40, 0xa9, 0x51, 0xc5, 0x52, 0x09,
	0x23, 0xc8, 0x58, 0xa2, 0xc4, 0x92, 0xe7, 0x59, 0x91, 0x19, 0x4c, 0xe2, 0xca, 0xa3, 0xa3, 0x17,
	0xa9, 0x10, 0x69, 0x8e, 0xa7, 0x35, 0xe4, 0xba, 0xbc, 0x39, 0xc5, 0x42, 0x9a, 0x5b, 0x1b, 0x11,
	0xbd, 0xec, 0x3a, 0x17, 0x8a, 0x49, 0x89, 0x4a, 0x5b, 0x3f, 0x9d, 0xc0, 0xee, 0x2c, 0x4b, 0xf9,
	0x57, 0x7e, 0xc1, 0x14, 0x2b, 0x34, 0x89, 0xe0, 0x71, 0xc5, 0xca, 0x59, 0x81, 0xe1, 0xd6, 0xab,
	0xad, 0xe3, 0x27, 0xd3, 0xa5, 0x5d, 0xf9, 0x24, 0xd3, 0x7a, 0x21, 0x54, 0x12, 0xf6, 0xac, 0xaf,
	0xb5, 0xe9, 0x4f, 0x38, 0xf8, 0xa4, 0x90, 0x19, 0xbc, 0xd4, 0xa8, 0x1e, 0xc6, 0x45, 0x02, 0x18,
	0x62, 0xc1, 0xb2, 0x3c, 0xec, 0xd7, 0x0e, 0x6b, 0xd0, 0x7d, 0xd8, 0x3b, 0x43, 0x73, 0x47, 0x4f,
	0x13, 0x08, 0x2e, 0x65, 0xc2, 0x0c, 0x5e, 0x34, 0x81, 0x8d, 0xec, 0x09, 0x1c, 0xcc, 0x4b, 0xa5,
	0x90, 0x9b, 0xab, 0xa5, 0x84, 0x95, 0xdf, 0x6f, 0xfe, 0xb7, 0x01, 0xe4, 0x08, 0x76, 0x39, 0x2e,
	0xae, 0x3a, 0x99, 0xec, 0x70, 0x5c, 0xb4, 0x10, 0x1a, 0xc2, 0xa1, 0xab, 0x32, 0x45, 0x2d, 0x05,
	0xd7, 0x48, 0xcf, 0x21, 0x98, 0x08, 0x95, 0x0a, 0xd3, 0xd1, 0xbf, 0xa7, 0xec, 0x9c, 0xf1, 0xb4,
	0x64, 0x29, 0xb6, 0x65, 0xb7, 0x36, 0xfd, 0x01, 0x83, 0xaa, 0x3a, 0x32, 0x82, 0x5e, 0x66, 0x33,
	0xee, 0x4f, 0x7b, 0x59, 0xe2, 0xf0, 0xf5, 0x3a, 0x7c, 0xde, 0x56, 0x11, 0x02, 0x03, 0x25, 0x72,
	0x0c, 0x07, 0xf5, 0xcf, 0xfa, 0x9b, 0x9e, 0xc1, 0x78, 0x8a, 0x1a, 0xbb, 0xc9, 0x06, 0x30, 0x34,
	0xe2, 0x37, 0xf2, 0x26, 0x53, 0x6b, 0xac, 0x9d, 0xf4, 0x33, 0x78, 0xea, 0x10, 0x2d, 0xfb, 0xf1,
	0x0e, 0x82, 0x6f, 0xa8, 0xb2, 0x9b, 0xdb, 0x4d, 0x24, 0xaa, 0xbe, 0xba, 0xe8, 0x96, 0xe7, 0xfd,
	0xdf, 0x21, 0xec, 0x54, 0x8d, 0x98, 0xa1, 0xfa, 0x93, 0xcd, 0x91, 0x9c, 0x03, 0xdc, 0xad, 0x16,
	0x79, 0x1d, 0x7b, 0x6e, 0x20, 0xee, 0xee, 0x5e, 0xf4, 0xdc, 0x0b, 0xab, 0x19, 0x7e, 0xc1, 0xc8,
	0x9d, 0x28, 0x39, 0xf1, 0x83, 0x3d, 0xcb, 0x15, 0xbd, 0xdd, 0x00, 0xda, 0x56, 0x42, 0x66, 0x30,
	0x72, 0x37, 0x64, 0x85, 0x92, 0x6f, 0x8d, 0xa2, 0xc3, 0xd8, 0x9e, 0x6e, 0xdc, 0x9e, 0x6e, 0xfc,
	0xb9, 0xba, 0x6b, 0x52, 0xc0, 0xd8, 0x36, 0xce, 0x99, 0xc2, 0x0a, 0x66, 0xdf, 0x40, 0x56, 0xd4,
	0xe0, 0x9f, 0x06, 0x41, 0xd8, 0x73, 0x85, 0x8e, 0xbd, 0xd1, 0x9e, 0xdd, 0x8a, 0xde, 0xdc, 0x8f,
	0x5c, 0xca, 0x7c, 0x81, 0x47, 0xcd, 0x75, 0x13, 0xea, 0x0d, 0x73, 0x6e, 0x7f, 0xdd, 0x78, 0x27,
	0xb0, 0x6d, 0x5f, 0x34, 0x72, 0xe4, 0x05, 0xfd, 0xff, 0xdc, 0xad, 0xe1, 0xf9, 0x38, 0xfc, 0xde,
	0x57, 0x72, 0x7e, 0xbd, 0x5d, 0xb7, 0xff, 0xc3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x36,
	0x76, 0x41, 0x87, 0x05, 0x00, 0x00,
}

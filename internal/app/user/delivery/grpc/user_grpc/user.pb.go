// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/app/user/delivery/grpc/user_grpc/user.proto

package user_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9ce0a4e02b3337a, []int{0}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type UserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9ce0a4e02b3337a, []int{1}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type User struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	SecondName           string   `protobuf:"bytes,3,opt,name=SecondName,proto3" json:"SecondName,omitempty"`
	UserName             string   `protobuf:"bytes,4,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=Password,proto3" json:"Password,omitempty"`
	EncryptPassword      string   `protobuf:"bytes,7,opt,name=EncryptPassword,proto3" json:"EncryptPassword,omitempty"`
	UserType             string   `protobuf:"bytes,8,opt,name=UserType,proto3" json:"UserType,omitempty"`
	FreelancerId         int64    `protobuf:"varint,9,opt,name=FreelancerId,proto3" json:"FreelancerId,omitempty"`
	HireManagerId        int64    `protobuf:"varint,10,opt,name=HireManagerId,proto3" json:"HireManagerId,omitempty"`
	CompanyId            int64    `protobuf:"varint,11,opt,name=CompanyId,proto3" json:"CompanyId,omitempty"`
	Avatar               []byte   `protobuf:"bytes,12,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9ce0a4e02b3337a, []int{2}
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

func (m *User) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetSecondName() string {
	if m != nil {
		return m.SecondName
	}
	return ""
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEncryptPassword() string {
	if m != nil {
		return m.EncryptPassword
	}
	return ""
}

func (m *User) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

func (m *User) GetFreelancerId() int64 {
	if m != nil {
		return m.FreelancerId
	}
	return 0
}

func (m *User) GetHireManagerId() int64 {
	if m != nil {
		return m.HireManagerId
	}
	return 0
}

func (m *User) GetCompanyId() int64 {
	if m != nil {
		return m.CompanyId
	}
	return 0
}

func (m *User) GetAvatar() []byte {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func init() {
	proto.RegisterType((*UserID)(nil), "user_grpc.UserID")
	proto.RegisterType((*UserRequest)(nil), "user_grpc.UserRequest")
	proto.RegisterType((*User)(nil), "user_grpc.User")
}

func init() {
	proto.RegisterFile("internal/app/user/delivery/grpc/user_grpc/user.proto", fileDescriptor_b9ce0a4e02b3337a)
}

var fileDescriptor_b9ce0a4e02b3337a = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x5f, 0x8b, 0x9b, 0x40,
	0x14, 0xc5, 0xd1, 0x24, 0x36, 0xde, 0xa4, 0x0d, 0x1d, 0x4a, 0x90, 0x50, 0x4a, 0x90, 0x3e, 0x84,
	0x3e, 0x28, 0xb4, 0x85, 0x3e, 0x96, 0x12, 0x13, 0xe2, 0x43, 0x4b, 0x71, 0xff, 0xbc, 0x2e, 0xb3,
	0x7a, 0x37, 0x08, 0x66, 0x9c, 0xbd, 0x9a, 0x2c, 0x7e, 0x9d, 0xfd, 0x74, 0xfb, 0x31, 0x96, 0x19,
	0xa3, 0xd9, 0x98, 0x37, 0xcf, 0xef, 0x9c, 0x99, 0x7b, 0x9c, 0x19, 0xf8, 0x99, 0x8a, 0x12, 0x49,
	0xf0, 0xcc, 0xe7, 0x52, 0xfa, 0xfb, 0x02, 0xc9, 0x4f, 0x30, 0x4b, 0x0f, 0x48, 0x95, 0xbf, 0x25,
	0x19, 0x6b, 0x74, 0xd7, 0x7e, 0x79, 0x92, 0xf2, 0x32, 0x67, 0x76, 0x4b, 0x5d, 0x07, 0xac, 0x9b,
	0x02, 0x29, 0x0c, 0xd8, 0x07, 0x30, 0xc3, 0xc0, 0x31, 0xe6, 0xc6, 0xa2, 0x17, 0x99, 0x61, 0xe0,
	0xfe, 0x86, 0x91, 0x72, 0x22, 0x7c, 0xdc, 0x63, 0x51, 0xb2, 0x4f, 0x30, 0xc0, 0x1d, 0x4f, 0x33,
	0x9d, 0xb0, 0xa3, 0x5a, 0xb0, 0x19, 0x0c, 0x25, 0x2f, 0x8a, 0xa7, 0x9c, 0x12, 0xc7, 0xd4, 0x46,
	0xab, 0xdd, 0x17, 0x13, 0xfa, 0x6a, 0x87, 0xee, 0xce, 0xec, 0x33, 0xd8, 0xeb, 0x94, 0x8a, 0xf2,
	0x1f, 0xdf, 0xe1, 0x71, 0xd5, 0x09, 0xb0, 0x2f, 0x00, 0x57, 0x18, 0xe7, 0x22, 0xd1, 0x76, 0x4f,
	0xdb, 0x6f, 0x88, 0x1a, 0xa9, 0x76, 0xd5, 0x6e, 0xbf, 0x1e, 0xd9, 0x68, 0x55, 0x72, 0xa5, 0x4b,
	0x0e, 0xea, 0x92, 0xab, 0xa6, 0xe4, 0xff, 0xa6, 0xa4, 0x55, 0xaf, 0x68, 0x34, 0x5b, 0xc0, 0x64,
	0x25, 0x62, 0xaa, 0x64, 0xd9, 0x46, 0xde, 0xe9, 0x48, 0x17, 0x37, 0x73, 0xaf, 0x2b, 0x89, 0xce,
	0xf0, 0x34, 0x57, 0x69, 0xe6, 0xc2, 0x78, 0x4d, 0x88, 0x19, 0x17, 0x31, 0x52, 0x98, 0x38, 0xb6,
	0xfe, 0xd7, 0x33, 0xc6, 0xbe, 0xc2, 0xfb, 0x4d, 0x4a, 0xf8, 0x97, 0x0b, 0xbe, 0xd5, 0x21, 0xd0,
	0xa1, 0x73, 0xa8, 0xce, 0x66, 0x99, 0xef, 0x24, 0x17, 0x55, 0x98, 0x38, 0x23, 0x9d, 0x38, 0x01,
	0x36, 0x05, 0xeb, 0xcf, 0x81, 0x97, 0x9c, 0x9c, 0xf1, 0xdc, 0x58, 0x8c, 0xa3, 0xa3, 0xfa, 0xfe,
	0x6c, 0xd4, 0x97, 0xb5, 0xe1, 0x22, 0xc9, 0x90, 0x98, 0x07, 0xb0, 0x24, 0xe4, 0x25, 0xea, 0xf3,
	0x9f, 0x78, 0xed, 0x7d, 0x7b, 0x0a, 0xcc, 0xba, 0x80, 0xfd, 0x02, 0xb8, 0x45, 0x4a, 0x1f, 0x2a,
	0xad, 0xa6, 0x1d, 0xfb, 0xf8, 0x04, 0x66, 0x1f, 0x3b, 0x3c, 0x0c, 0xd8, 0x37, 0xe8, 0xaf, 0x53,
	0x91, 0xb0, 0x4b, 0xeb, 0x62, 0xc8, 0xbd, 0xa5, 0x1f, 0xdf, 0x8f, 0xd7, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0xd5, 0x9f, 0xb6, 0xb4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserHandlerClient is the client API for UserHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserHandlerClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	VerifyUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserID, error)
	Find(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error)
}

type userHandlerClient struct {
	cc *grpc.ClientConn
}

func NewUserHandlerClient(cc *grpc.ClientConn) UserHandlerClient {
	return &userHandlerClient{cc}
}

func (c *userHandlerClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_grpc.UserHandler/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHandlerClient) VerifyUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/user_grpc.UserHandler/VerifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHandlerClient) Find(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_grpc.UserHandler/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserHandlerServer is the server API for UserHandler service.
type UserHandlerServer interface {
	CreateUser(context.Context, *User) (*User, error)
	VerifyUser(context.Context, *UserRequest) (*UserID, error)
	Find(context.Context, *UserID) (*User, error)
}

// UnimplementedUserHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedUserHandlerServer struct {
}

func (*UnimplementedUserHandlerServer) CreateUser(ctx context.Context, req *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserHandlerServer) VerifyUser(ctx context.Context, req *UserRequest) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUser not implemented")
}
func (*UnimplementedUserHandlerServer) Find(ctx context.Context, req *UserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}

func RegisterUserHandlerServer(s *grpc.Server, srv UserHandlerServer) {
	s.RegisterService(&_UserHandler_serviceDesc, srv)
}

func _UserHandler_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_grpc.UserHandler/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHandler_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_grpc.UserHandler/VerifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).VerifyUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHandler_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHandlerServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_grpc.UserHandler/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHandlerServer).Find(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user_grpc.UserHandler",
	HandlerType: (*UserHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserHandler_CreateUser_Handler,
		},
		{
			MethodName: "VerifyUser",
			Handler:    _UserHandler_VerifyUser_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _UserHandler_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/user/delivery/grpc/user_grpc/user.proto",
}

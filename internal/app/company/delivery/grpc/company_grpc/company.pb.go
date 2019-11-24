// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/app/company/delivery/grpc/company_grpc/company.proto

package company_grpc

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

type Company struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CompanyName          string   `protobuf:"bytes,2,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	Site                 string   `protobuf:"bytes,3,opt,name=Site,proto3" json:"Site,omitempty"`
	TagLine              string   `protobuf:"bytes,4,opt,name=TagLine,proto3" json:"TagLine,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	Country              string   `protobuf:"bytes,6,opt,name=Country,proto3" json:"Country,omitempty"`
	City                 string   `protobuf:"bytes,7,opt,name=City,proto3" json:"City,omitempty"`
	Address              string   `protobuf:"bytes,8,opt,name=Address,proto3" json:"Address,omitempty"`
	Phone                string   `protobuf:"bytes,9,opt,name=Phone,proto3" json:"Phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Company) Reset()         { *m = Company{} }
func (m *Company) String() string { return proto.CompactTextString(m) }
func (*Company) ProtoMessage()    {}
func (*Company) Descriptor() ([]byte, []int) {
	return fileDescriptor_5755ad00c494f094, []int{0}
}

func (m *Company) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Company.Unmarshal(m, b)
}
func (m *Company) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Company.Marshal(b, m, deterministic)
}
func (m *Company) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Company.Merge(m, src)
}
func (m *Company) XXX_Size() int {
	return xxx_messageInfo_Company.Size(m)
}
func (m *Company) XXX_DiscardUnknown() {
	xxx_messageInfo_Company.DiscardUnknown(m)
}

var xxx_messageInfo_Company proto.InternalMessageInfo

func (m *Company) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Company) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *Company) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func (m *Company) GetTagLine() string {
	if m != nil {
		return m.TagLine
	}
	return ""
}

func (m *Company) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Company) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Company) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Company) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Company) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type CompanyID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompanyID) Reset()         { *m = CompanyID{} }
func (m *CompanyID) String() string { return proto.CompactTextString(m) }
func (*CompanyID) ProtoMessage()    {}
func (*CompanyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_5755ad00c494f094, []int{1}
}

func (m *CompanyID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompanyID.Unmarshal(m, b)
}
func (m *CompanyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompanyID.Marshal(b, m, deterministic)
}
func (m *CompanyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompanyID.Merge(m, src)
}
func (m *CompanyID) XXX_Size() int {
	return xxx_messageInfo_CompanyID.Size(m)
}
func (m *CompanyID) XXX_DiscardUnknown() {
	xxx_messageInfo_CompanyID.DiscardUnknown(m)
}

var xxx_messageInfo_CompanyID proto.InternalMessageInfo

func (m *CompanyID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

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
	return fileDescriptor_5755ad00c494f094, []int{2}
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

func init() {
	proto.RegisterType((*Company)(nil), "company_grpc.Company")
	proto.RegisterType((*CompanyID)(nil), "company_grpc.CompanyID")
	proto.RegisterType((*UserID)(nil), "company_grpc.UserID")
}

func init() {
	proto.RegisterFile("internal/app/company/delivery/grpc/company_grpc/company.proto", fileDescriptor_5755ad00c494f094)
}

var fileDescriptor_5755ad00c494f094 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x69, 0xb7, 0xdb, 0xda, 0x51, 0xf7, 0x10, 0x56, 0x0c, 0x7a, 0x29, 0x3d, 0xed, 0xa9,
	0x05, 0x05, 0x4f, 0x7a, 0x90, 0x16, 0xb1, 0x20, 0x22, 0xab, 0x9e, 0x25, 0xb6, 0xc3, 0x1a, 0xe8,
	0x26, 0x21, 0x8d, 0x42, 0x5f, 0xc0, 0x77, 0xf5, 0x2d, 0xa4, 0x4d, 0x0b, 0xf5, 0xcf, 0xde, 0xe6,
	0xfb, 0x7d, 0xf3, 0xcd, 0xb4, 0x13, 0xb8, 0xe2, 0xc2, 0xa0, 0x16, 0xac, 0x4e, 0x99, 0x52, 0x69,
	0x29, 0xb7, 0x8a, 0x89, 0x36, 0xad, 0xb0, 0xe6, 0x1f, 0xa8, 0xdb, 0x74, 0xa3, 0x55, 0x39, 0xd2,
	0x97, 0xa9, 0x48, 0x94, 0x96, 0x46, 0x92, 0x83, 0xa9, 0x17, 0x7f, 0x39, 0x10, 0x64, 0x16, 0x90,
	0x05, 0xb8, 0x45, 0x4e, 0x9d, 0xc8, 0x59, 0xcd, 0xd6, 0x6e, 0x91, 0x93, 0x08, 0xf6, 0x07, 0xeb,
	0x9e, 0x6d, 0x91, 0xba, 0x91, 0xb3, 0x0a, 0xd7, 0x53, 0x44, 0x08, 0x78, 0x8f, 0xdc, 0x20, 0x9d,
	0xf5, 0x56, 0x5f, 0x13, 0x0a, 0xc1, 0x13, 0xdb, 0xdc, 0x71, 0x81, 0xd4, 0xeb, 0xf1, 0x28, 0xbb,
	0x79, 0x39, 0x36, 0xa5, 0xe6, 0xca, 0x70, 0x29, 0xe8, 0xdc, 0xce, 0x9b, 0xa0, 0x2e, 0x9b, 0xc9,
	0x77, 0x61, 0x74, 0x4b, 0x7d, 0x9b, 0x1d, 0x64, 0xb7, 0x29, 0xe3, 0xa6, 0xa5, 0x81, 0xdd, 0xd4,
	0xd5, 0x5d, 0xf7, 0x75, 0x55, 0x69, 0x6c, 0x1a, 0xba, 0x67, 0xbb, 0x07, 0x49, 0x96, 0x30, 0x7f,
	0x78, 0x93, 0x02, 0x69, 0xd8, 0x73, 0x2b, 0xe2, 0x53, 0x08, 0x87, 0x8f, 0x2f, 0xf2, 0xdf, 0x3f,
	0x1b, 0x53, 0xf0, 0x9f, 0x1b, 0xd4, 0x7f, 0x9d, 0xb3, 0x4f, 0x07, 0x16, 0x43, 0xee, 0x96, 0x89,
	0xaa, 0x46, 0x4d, 0x2e, 0xe1, 0x30, 0xd3, 0xc8, 0x0c, 0x8e, 0xa7, 0x5b, 0x26, 0xd3, 0xab, 0x26,
	0x76, 0xd2, 0xc9, 0xd1, 0x4f, 0x3a, 0x36, 0x5f, 0x80, 0x77, 0xc3, 0x45, 0x45, 0x8e, 0xff, 0xb5,
	0x77, 0xe6, 0x5e, 0xfd, 0xfe, 0x01, 0xcf, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x1c, 0xd1,
	0xd7, 0x01, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CompanyHandlerClient is the client API for CompanyHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompanyHandlerClient interface {
	CreateCompany(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Company, error)
	Find(ctx context.Context, in *CompanyID, opts ...grpc.CallOption) (*Company, error)
}

type companyHandlerClient struct {
	cc *grpc.ClientConn
}

func NewCompanyHandlerClient(cc *grpc.ClientConn) CompanyHandlerClient {
	return &companyHandlerClient{cc}
}

func (c *companyHandlerClient) CreateCompany(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/company_grpc.CompanyHandler/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyHandlerClient) Find(ctx context.Context, in *CompanyID, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/company_grpc.CompanyHandler/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyHandlerServer is the server API for CompanyHandler service.
type CompanyHandlerServer interface {
	CreateCompany(context.Context, *UserID) (*Company, error)
	Find(context.Context, *CompanyID) (*Company, error)
}

// UnimplementedCompanyHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedCompanyHandlerServer struct {
}

func (*UnimplementedCompanyHandlerServer) CreateCompany(ctx context.Context, req *UserID) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (*UnimplementedCompanyHandlerServer) Find(ctx context.Context, req *CompanyID) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}

func RegisterCompanyHandlerServer(s *grpc.Server, srv CompanyHandlerServer) {
	s.RegisterService(&_CompanyHandler_serviceDesc, srv)
}

func _CompanyHandler_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyHandlerServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company_grpc.CompanyHandler/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyHandlerServer).CreateCompany(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyHandler_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyHandlerServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company_grpc.CompanyHandler/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyHandlerServer).Find(ctx, req.(*CompanyID))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompanyHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "company_grpc.CompanyHandler",
	HandlerType: (*CompanyHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyHandler_CreateCompany_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _CompanyHandler_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/company/delivery/grpc/company_grpc/company.proto",
}

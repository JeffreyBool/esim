package new

func ProtoBufInit() {
	fc1 := &FileContent{
		FileName: "user.pb.go",
		Dir:      "internal/infra/third_party/protobuf/passport",
		Content: `// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package passport

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

// The request message containing the user's name.
type GetUserByUserNameRequest struct {
	Username             string   {{!}}protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *GetUserByUserNameRequest) Reset()         { *m = GetUserByUserNameRequest{} }
func (m *GetUserByUserNameRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByUserNameRequest) ProtoMessage()    {}
func (*GetUserByUserNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserByUserNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByUserNameRequest.Unmarshal(m, b)
}
func (m *GetUserByUserNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByUserNameRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByUserNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByUserNameRequest.Merge(m, src)
}
func (m *GetUserByUserNameRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByUserNameRequest.Size(m)
}
func (m *GetUserByUserNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByUserNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByUserNameRequest proto.InternalMessageInfo

func (m *GetUserByUserNameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

// The response message containing the greetings
type GrpcUserReply struct {
	Code                 int32    {{!}}protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"{{!}}
	Data                 *Info    {{!}}protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"{{!}}
	Msg                  string   {{!}}protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *GrpcUserReply) Reset()         { *m = GrpcUserReply{} }
func (m *GrpcUserReply) String() string { return proto.CompactTextString(m) }
func (*GrpcUserReply) ProtoMessage()    {}
func (*GrpcUserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GrpcUserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcUserReply.Unmarshal(m, b)
}
func (m *GrpcUserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcUserReply.Marshal(b, m, deterministic)
}
func (m *GrpcUserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcUserReply.Merge(m, src)
}
func (m *GrpcUserReply) XXX_Size() int {
	return xxx_messageInfo_GrpcUserReply.Size(m)
}
func (m *GrpcUserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcUserReply.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcUserReply proto.InternalMessageInfo

func (m *GrpcUserReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GrpcUserReply) GetData() *Info {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GrpcUserReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Info struct {
	UserName             string   {{!}}protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"{{!}}
	PassWord             string   {{!}}protobuf:"bytes,2,opt,name=pass_word,json=passWord,proto3" json:"pass_word,omitempty"{{!}}
	XXX_NoUnkeyedLiteral struct{} {{!}}json:"-"{{!}}
	XXX_unrecognized     []byte   {{!}}json:"-"{{!}}
	XXX_sizecache        int32    {{!}}json:"-"{{!}}
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Info) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserByUserNameRequest)(nil), "passport.GetUserByUserNameRequest")
	proto.RegisterType((*GrpcUserReply)(nil), "passport.GrpcUserReply")
	proto.RegisterType((*Info)(nil), "passport.Info")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xbf, 0x4b, 0xc5, 0x30,
	0x10, 0xb6, 0xbe, 0x2a, 0xed, 0x89, 0xa2, 0xb7, 0x18, 0x9e, 0x4b, 0xc9, 0xd4, 0xa9, 0x43, 0x05,
	0x67, 0x71, 0x29, 0x2e, 0x0e, 0x01, 0x11, 0x17, 0x4b, 0x6c, 0x4e, 0x17, 0xdb, 0xc4, 0x24, 0x45,
	0xfa, 0xdf, 0xcb, 0x45, 0xd4, 0xa2, 0xbc, 0x25, 0x5c, 0xf8, 0xee, 0xfb, 0x75, 0x00, 0x73, 0x20,
	0xdf, 0x38, 0x6f, 0xa3, 0xc5, 0xc2, 0xe9, 0x10, 0x9c, 0xf5, 0x51, 0x5e, 0x81, 0xe8, 0x28, 0xde,
	0x07, 0xf2, 0x37, 0x0b, 0xbf, 0x77, 0x7a, 0x24, 0x45, 0xef, 0x33, 0x85, 0x88, 0x5b, 0x28, 0x98,
	0x33, 0xe9, 0x91, 0x44, 0x56, 0x65, 0x75, 0xa9, 0x7e, 0xfe, 0xf2, 0x11, 0x8e, 0x3b, 0xef, 0x06,
	0xa6, 0x28, 0x72, 0x6f, 0x0b, 0x22, 0xe4, 0x83, 0x35, 0x5f, 0x8b, 0x07, 0x2a, 0xcd, 0x28, 0x21,
	0x37, 0x3a, 0x6a, 0xb1, 0x5f, 0x65, 0xf5, 0x51, 0x7b, 0xd2, 0x7c, 0xbb, 0x36, 0xb7, 0xd3, 0x8b,
	0x55, 0x09, 0xc3, 0x53, 0xd8, 0x8c, 0xe1, 0x55, 0x6c, 0x92, 0x3e, 0x8f, 0xf2, 0x1a, 0x72, 0xc6,
	0xf1, 0x02, 0x4a, 0xb6, 0xeb, 0xff, 0xfa, 0x73, 0x44, 0x06, 0x59, 0xad, 0xff, 0xb0, 0xde, 0x24,
	0xfd, 0x52, 0xa5, 0x52, 0x0f, 0xd6, 0x9b, 0xf6, 0x09, 0x0a, 0x0e, 0x96, 0x54, 0x14, 0x9c, 0xfd,
	0x2b, 0x88, 0xf2, 0x37, 0xca, 0xae, 0xf6, 0xdb, 0xf3, 0xd5, 0xce, 0xba, 0xa9, 0xdc, 0x7b, 0x3e,
	0x4c, 0x57, 0xbc, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x63, 0xd2, 0x84, 0xdb, 0x53, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInfoClient interface {
	GetUserByUserName(ctx context.Context, in *GetUserByUserNameRequest, opts ...grpc.CallOption) (*GrpcUserReply, error)
}

type userInfoClient struct {
	cc *grpc.ClientConn
}

func NewUserInfoClient(cc *grpc.ClientConn) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) GetUserByUserName(ctx context.Context, in *GetUserByUserNameRequest, opts ...grpc.CallOption) (*GrpcUserReply, error) {
	out := new(GrpcUserReply)
	err := c.cc.Invoke(ctx, "/passport.UserInfo/GetUserByUserName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServer is the server API for UserInfo service.
type UserInfoServer interface {
	GetUserByUserName(context.Context, *GetUserByUserNameRequest) (*GrpcUserReply, error)
}

// UnimplementedUserInfoServer can be embedded to have forward compatible implementations.
type UnimplementedUserInfoServer struct {
}

func (*UnimplementedUserInfoServer) GetUserByUserName(ctx context.Context, req *GetUserByUserNameRequest) (*GrpcUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUserName not implemented")
}

func RegisterUserInfoServer(s *grpc.Server, srv UserInfoServer) {
	s.RegisterService(&_UserInfo_serviceDesc, srv)
}

func _UserInfo_GetUserByUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).GetUserByUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passport.UserInfo/GetUserByUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).GetUserByUserName(ctx, req.(*GetUserByUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "passport.UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByUserName",
			Handler:    _UserInfo_GetUserByUserName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

`,
	}

	Files = append(Files, fc1)
}
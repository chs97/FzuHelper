// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	User
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id       int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Stdno    string `protobuf:"bytes,2,opt,name=stdno" json:"stdno,omitempty"`
	Realname string `protobuf:"bytes,3,opt,name=realname" json:"realname,omitempty"`
	Grade    string `protobuf:"bytes,4,opt,name=grade" json:"grade,omitempty"`
	College  string `protobuf:"bytes,5,opt,name=college" json:"college,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetStdno() string {
	if m != nil {
		return m.Stdno
	}
	return ""
}

func (m *User) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *User) GetGrade() string {
	if m != nil {
		return m.Grade
	}
	return ""
}

func (m *User) GetCollege() string {
	if m != nil {
		return m.College
	}
	return ""
}

type CreateRequest struct {
	Stdno    string `protobuf:"bytes,1,opt,name=stdno" json:"stdno,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Realname string `protobuf:"bytes,3,opt,name=realname" json:"realname,omitempty"`
	Grade    string `protobuf:"bytes,4,opt,name=grade" json:"grade,omitempty"`
	College  string `protobuf:"bytes,5,opt,name=college" json:"college,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetStdno() string {
	if m != nil {
		return m.Stdno
	}
	return ""
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateRequest) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *CreateRequest) GetGrade() string {
	if m != nil {
		return m.Grade
	}
	return ""
}

func (m *CreateRequest) GetCollege() string {
	if m != nil {
		return m.College
	}
	return ""
}

type CreateResponse struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadRequest struct {
	Stdno string `protobuf:"bytes,1,opt,name=stdno" json:"stdno,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadRequest) GetStdno() string {
	if m != nil {
		return m.Stdno
	}
	return ""
}

type ReadResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
}

type authClient struct {
	c           client.Client
	serviceName string
}

func NewAuthClient(serviceName string, c client.Client) AuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Auth{hdlr}, opts...))
}

type Auth struct {
	AuthHandler
}

func (h *Auth) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.AuthHandler.Create(ctx, in, out)
}

func (h *Auth) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.AuthHandler.Read(ctx, in, out)
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x9b, 0x35, 0x5b, 0x75, 0xda, 0xae, 0x10, 0x3c, 0xc4, 0x3d, 0x2d, 0x11, 0xa1, 0x22,
	0xe4, 0x50, 0x9f, 0x40, 0x7c, 0x83, 0x80, 0x17, 0x6f, 0xd1, 0x0c, 0x6d, 0x61, 0xdd, 0xac, 0x99,
	0x2c, 0xfa, 0x12, 0xbe, 0xb3, 0x6c, 0xb6, 0xad, 0x6d, 0x0f, 0x9e, 0x7a, 0xfc, 0x26, 0x33, 0xf3,
	0x7f, 0x4c, 0x00, 0x6c, 0x17, 0x57, 0xba, 0x0d, 0x3e, 0x7a, 0xf5, 0x0d, 0xfc, 0x85, 0x30, 0x88,
	0x02, 0xb2, 0xb5, 0x93, 0xac, 0x62, 0xf3, 0xdc, 0x64, 0x6b, 0x27, 0xae, 0x21, 0xa7, 0xe8, 0x1a,
	0x2f, 0xb3, 0x8a, 0xcd, 0x2f, 0xcd, 0x00, 0xa2, 0x84, 0x8b, 0x80, 0xb6, 0x6e, 0xec, 0x07, 0xca,
	0xb3, 0xf4, 0xb0, 0xe3, 0x7e, 0x62, 0x19, 0xac, 0x43, 0xc9, 0x87, 0x89, 0x04, 0x42, 0xc2, 0xf9,
	0xbb, 0xaf, 0x6b, 0x5c, 0xa2, 0xcc, 0x53, 0x7d, 0x8b, 0xea, 0x87, 0xc1, 0xec, 0x39, 0xa0, 0x8d,
	0x68, 0xf0, 0xb3, 0x43, 0x8a, 0x7f, 0x99, 0xec, 0x28, 0xb3, 0xb5, 0x44, 0x5f, 0x3e, 0xb8, 0x8d,
	0xcc, 0x8e, 0x4f, 0xea, 0x53, 0x41, 0xb1, 0xd5, 0xa1, 0xd6, 0x37, 0x84, 0xc7, 0x37, 0x51, 0xb7,
	0x30, 0x31, 0x68, 0xdd, 0xbf, 0xba, 0xea, 0x1e, 0xa6, 0x43, 0xd3, 0x66, 0xc9, 0x0d, 0xf0, 0x8e,
	0x30, 0xa4, 0xa6, 0xc9, 0x22, 0xd7, 0xfd, 0xb5, 0x4d, 0x2a, 0x2d, 0x5e, 0x81, 0x3f, 0x75, 0x71,
	0x25, 0x1e, 0x60, 0x3c, 0x24, 0x8b, 0x42, 0x1f, 0x5c, 0xa4, 0xbc, 0xd2, 0x87, 0x4a, 0x6a, 0x24,
	0xee, 0x80, 0xf7, 0xfb, 0xc5, 0x54, 0xef, 0xb9, 0x94, 0x33, 0xbd, 0x1f, 0xaa, 0x46, 0x6f, 0xe3,
	0xf4, 0xbd, 0x8f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x98, 0x03, 0x2f, 0xca, 0xec, 0x01, 0x00,
	0x00,
}

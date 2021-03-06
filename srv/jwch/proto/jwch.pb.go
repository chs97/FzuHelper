// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jwch.proto

/*
Package jwch is a generated protocol buffer package.

It is generated from these files:
	jwch.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	GetInfoRequest
	GetInfoResponse
	JwtUserLoginRequest
	JwtUserLoginResponse
*/
package jwch

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

type LoginRequest struct {
	Stdno    string `protobuf:"bytes,1,opt,name=stdno" json:"stdno,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetStdno() string {
	if m != nil {
		return m.Stdno
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetInfoRequest struct {
	Stdno    string `protobuf:"bytes,1,opt,name=stdno" json:"stdno,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *GetInfoRequest) Reset()                    { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()               {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetInfoRequest) GetStdno() string {
	if m != nil {
		return m.Stdno
	}
	return ""
}

func (m *GetInfoRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GetInfoResponse struct {
	Stdno    string `protobuf:"bytes,1,opt,name=stdno" json:"stdno,omitempty"`
	Realname string `protobuf:"bytes,2,opt,name=realname" json:"realname,omitempty"`
	Grade    string `protobuf:"bytes,3,opt,name=grade" json:"grade,omitempty"`
	College  string `protobuf:"bytes,4,opt,name=college" json:"college,omitempty"`
}

func (m *GetInfoResponse) Reset()                    { *m = GetInfoResponse{} }
func (m *GetInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetInfoResponse) ProtoMessage()               {}
func (*GetInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetInfoResponse) GetStdno() string {
	if m != nil {
		return m.Stdno
	}
	return ""
}

func (m *GetInfoResponse) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *GetInfoResponse) GetGrade() string {
	if m != nil {
		return m.Grade
	}
	return ""
}

func (m *GetInfoResponse) GetCollege() string {
	if m != nil {
		return m.College
	}
	return ""
}

type JwtUserLoginRequest struct {
	Stdno    string `protobuf:"bytes,1,opt,name=stdno" json:"stdno,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *JwtUserLoginRequest) Reset()                    { *m = JwtUserLoginRequest{} }
func (m *JwtUserLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*JwtUserLoginRequest) ProtoMessage()               {}
func (*JwtUserLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *JwtUserLoginRequest) GetStdno() string {
	if m != nil {
		return m.Stdno
	}
	return ""
}

func (m *JwtUserLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type JwtUserLoginResponse struct {
	Payload string `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *JwtUserLoginResponse) Reset()                    { *m = JwtUserLoginResponse{} }
func (m *JwtUserLoginResponse) String() string            { return proto.CompactTextString(m) }
func (*JwtUserLoginResponse) ProtoMessage()               {}
func (*JwtUserLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *JwtUserLoginResponse) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*GetInfoRequest)(nil), "GetInfoRequest")
	proto.RegisterType((*GetInfoResponse)(nil), "GetInfoResponse")
	proto.RegisterType((*JwtUserLoginRequest)(nil), "JwtUserLoginRequest")
	proto.RegisterType((*JwtUserLoginResponse)(nil), "JwtUserLoginResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Jwch service

type JwchClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Getinfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*GetInfoResponse, error)
	JwtUserLogin(ctx context.Context, in *JwtUserLoginRequest, opts ...client.CallOption) (*JwtUserLoginResponse, error)
}

type jwchClient struct {
	c           client.Client
	serviceName string
}

func NewJwchClient(serviceName string, c client.Client) JwchClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "jwch"
	}
	return &jwchClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *jwchClient) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Jwch.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwchClient) Getinfo(ctx context.Context, in *GetInfoRequest, opts ...client.CallOption) (*GetInfoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Jwch.Getinfo", in)
	out := new(GetInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwchClient) JwtUserLogin(ctx context.Context, in *JwtUserLoginRequest, opts ...client.CallOption) (*JwtUserLoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Jwch.JwtUserLogin", in)
	out := new(JwtUserLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Jwch service

type JwchHandler interface {
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Getinfo(context.Context, *GetInfoRequest, *GetInfoResponse) error
	JwtUserLogin(context.Context, *JwtUserLoginRequest, *JwtUserLoginResponse) error
}

func RegisterJwchHandler(s server.Server, hdlr JwchHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Jwch{hdlr}, opts...))
}

type Jwch struct {
	JwchHandler
}

func (h *Jwch) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.JwchHandler.Login(ctx, in, out)
}

func (h *Jwch) Getinfo(ctx context.Context, in *GetInfoRequest, out *GetInfoResponse) error {
	return h.JwchHandler.Getinfo(ctx, in, out)
}

func (h *Jwch) JwtUserLogin(ctx context.Context, in *JwtUserLoginRequest, out *JwtUserLoginResponse) error {
	return h.JwchHandler.JwtUserLogin(ctx, in, out)
}

func init() { proto.RegisterFile("jwch.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x45, 0x13, 0x6d, 0x8d, 0x3e, 0xda, 0x46, 0xc6, 0x08, 0x43, 0x56, 0x32, 0xab, 0xae, 0x06,
	0xd1, 0xb5, 0x20, 0x6e, 0x82, 0xc5, 0x55, 0xc1, 0x0f, 0x18, 0x93, 0xd7, 0xb4, 0x12, 0xe7, 0xc5,
	0x99, 0x91, 0xe0, 0xd7, 0xf8, 0xab, 0xe2, 0x24, 0x29, 0x8d, 0xd4, 0x8d, 0x5d, 0x9e, 0x84, 0x7b,
	0xb8, 0xef, 0x32, 0x00, 0xaf, 0x4d, 0xbe, 0x96, 0xb5, 0x21, 0x47, 0xe2, 0x1e, 0x26, 0x4f, 0x54,
	0x6e, 0xf4, 0x12, 0xdf, 0x3f, 0xd0, 0x3a, 0x96, 0xc0, 0xd8, 0xba, 0x42, 0x13, 0x0f, 0xaf, 0xc2,
	0xf9, 0xd9, 0xb2, 0x05, 0x96, 0xc2, 0x69, 0xad, 0xac, 0x6d, 0xc8, 0x14, 0xfc, 0xc8, 0xff, 0xd8,
	0xb2, 0x88, 0x61, 0xda, 0x19, 0x6c, 0x4d, 0xda, 0xa2, 0x78, 0x80, 0x59, 0x86, 0xee, 0x51, 0xaf,
	0xe8, 0xff, 0x52, 0x0b, 0xf1, 0xd6, 0xd1, 0x6a, 0xff, 0x96, 0x18, 0x54, 0x95, 0x56, 0x6f, 0xd8,
	0x4b, 0x7a, 0xfe, 0x49, 0x94, 0x46, 0x15, 0xc8, 0x8f, 0xdb, 0x84, 0x07, 0xc6, 0x21, 0xca, 0xa9,
	0xaa, 0xb0, 0x44, 0x3e, 0xf2, 0xdf, 0x7b, 0x14, 0x19, 0x5c, 0x2c, 0x1a, 0xf7, 0x6c, 0xd1, 0x1c,
	0x38, 0xc9, 0x35, 0x24, 0x43, 0x51, 0x77, 0x02, 0x87, 0xa8, 0x56, 0x9f, 0x15, 0xa9, 0xa2, 0x73,
	0xf5, 0x78, 0xf3, 0x15, 0xc2, 0x68, 0xd1, 0xe4, 0x6b, 0x36, 0x87, 0xb1, 0xcf, 0xb0, 0xa9, 0xdc,
	0x2d, 0x91, 0xce, 0xe4, 0x70, 0xe4, 0x80, 0x49, 0x88, 0x32, 0x74, 0x1b, 0xbd, 0x22, 0x16, 0xcb,
	0xe1, 0xe0, 0xe9, 0xb9, 0xfc, 0xb5, 0x9e, 0x08, 0xd8, 0x1d, 0x4c, 0x76, 0x4b, 0xb1, 0x44, 0xee,
	0x39, 0x36, 0xbd, 0x94, 0xfb, 0x9a, 0x8b, 0xe0, 0xe5, 0xc4, 0xbf, 0x97, 0xdb, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb5, 0x3d, 0x97, 0xd0, 0x3d, 0x02, 0x00, 0x00,
}

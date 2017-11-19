// Code generated by protoc-gen-go. DO NOT EDIT.
// source: umbrella.proto

/*
Package umbrella is a generated protocol buffer package.

It is generated from these files:
	umbrella.proto

It has these top-level messages:
	Umbrella
	GetAllUmbRequest
	GetAllumbResponse
	BorrowOneRequest
	BorrowOneResponse
	ReturnRequest
	ReturnResponse
*/
package umbrella

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

type Umbrella struct {
	Number string `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Rented bool   `protobuf:"varint,2,opt,name=rented" json:"rented,omitempty"`
}

func (m *Umbrella) Reset()                    { *m = Umbrella{} }
func (m *Umbrella) String() string            { return proto.CompactTextString(m) }
func (*Umbrella) ProtoMessage()               {}
func (*Umbrella) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Umbrella) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Umbrella) GetRented() bool {
	if m != nil {
		return m.Rented
	}
	return false
}

type GetAllUmbRequest struct {
}

func (m *GetAllUmbRequest) Reset()                    { *m = GetAllUmbRequest{} }
func (m *GetAllUmbRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllUmbRequest) ProtoMessage()               {}
func (*GetAllUmbRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetAllumbResponse struct {
	Umbrellas []*Umbrella `protobuf:"bytes,1,rep,name=umbrellas" json:"umbrellas,omitempty"`
}

func (m *GetAllumbResponse) Reset()                    { *m = GetAllumbResponse{} }
func (m *GetAllumbResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAllumbResponse) ProtoMessage()               {}
func (*GetAllumbResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetAllumbResponse) GetUmbrellas() []*Umbrella {
	if m != nil {
		return m.Umbrellas
	}
	return nil
}

type BorrowOneRequest struct {
	Stdno  string `protobuf:"bytes,1,opt,name=stdno" json:"stdno,omitempty"`
	Place  string `protobuf:"bytes,2,opt,name=place" json:"place,omitempty"`
	Number string `protobuf:"bytes,3,opt,name=number" json:"number,omitempty"`
}

func (m *BorrowOneRequest) Reset()                    { *m = BorrowOneRequest{} }
func (m *BorrowOneRequest) String() string            { return proto.CompactTextString(m) }
func (*BorrowOneRequest) ProtoMessage()               {}
func (*BorrowOneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BorrowOneRequest) GetStdno() string {
	if m != nil {
		return m.Stdno
	}
	return ""
}

func (m *BorrowOneRequest) GetPlace() string {
	if m != nil {
		return m.Place
	}
	return ""
}

func (m *BorrowOneRequest) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type BorrowOneResponse struct {
}

func (m *BorrowOneResponse) Reset()                    { *m = BorrowOneResponse{} }
func (m *BorrowOneResponse) String() string            { return proto.CompactTextString(m) }
func (*BorrowOneResponse) ProtoMessage()               {}
func (*BorrowOneResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ReturnRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReturnRequest) Reset()                    { *m = ReturnRequest{} }
func (m *ReturnRequest) String() string            { return proto.CompactTextString(m) }
func (*ReturnRequest) ProtoMessage()               {}
func (*ReturnRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReturnRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReturnResponse struct {
}

func (m *ReturnResponse) Reset()                    { *m = ReturnResponse{} }
func (m *ReturnResponse) String() string            { return proto.CompactTextString(m) }
func (*ReturnResponse) ProtoMessage()               {}
func (*ReturnResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Umbrella)(nil), "Umbrella")
	proto.RegisterType((*GetAllUmbRequest)(nil), "GetAllUmbRequest")
	proto.RegisterType((*GetAllumbResponse)(nil), "GetAllumbResponse")
	proto.RegisterType((*BorrowOneRequest)(nil), "BorrowOneRequest")
	proto.RegisterType((*BorrowOneResponse)(nil), "BorrowOneResponse")
	proto.RegisterType((*ReturnRequest)(nil), "ReturnRequest")
	proto.RegisterType((*ReturnResponse)(nil), "ReturnResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Lend service

type LendClient interface {
	GetAllUmb(ctx context.Context, in *GetAllUmbRequest, opts ...client.CallOption) (*GetAllumbResponse, error)
	BorrowOne(ctx context.Context, in *BorrowOneRequest, opts ...client.CallOption) (*BorrowOneResponse, error)
	ReturnOne(ctx context.Context, in *ReturnRequest, opts ...client.CallOption) (*ReturnResponse, error)
}

type lendClient struct {
	c           client.Client
	serviceName string
}

func NewLendClient(serviceName string, c client.Client) LendClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "lend"
	}
	return &lendClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *lendClient) GetAllUmb(ctx context.Context, in *GetAllUmbRequest, opts ...client.CallOption) (*GetAllumbResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Lend.GetAllUmb", in)
	out := new(GetAllumbResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lendClient) BorrowOne(ctx context.Context, in *BorrowOneRequest, opts ...client.CallOption) (*BorrowOneResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Lend.BorrowOne", in)
	out := new(BorrowOneResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lendClient) ReturnOne(ctx context.Context, in *ReturnRequest, opts ...client.CallOption) (*ReturnResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Lend.ReturnOne", in)
	out := new(ReturnResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lend service

type LendHandler interface {
	GetAllUmb(context.Context, *GetAllUmbRequest, *GetAllumbResponse) error
	BorrowOne(context.Context, *BorrowOneRequest, *BorrowOneResponse) error
	ReturnOne(context.Context, *ReturnRequest, *ReturnResponse) error
}

func RegisterLendHandler(s server.Server, hdlr LendHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Lend{hdlr}, opts...))
}

type Lend struct {
	LendHandler
}

func (h *Lend) GetAllUmb(ctx context.Context, in *GetAllUmbRequest, out *GetAllumbResponse) error {
	return h.LendHandler.GetAllUmb(ctx, in, out)
}

func (h *Lend) BorrowOne(ctx context.Context, in *BorrowOneRequest, out *BorrowOneResponse) error {
	return h.LendHandler.BorrowOne(ctx, in, out)
}

func (h *Lend) ReturnOne(ctx context.Context, in *ReturnRequest, out *ReturnResponse) error {
	return h.LendHandler.ReturnOne(ctx, in, out)
}

func init() { proto.RegisterFile("umbrella.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4f, 0xbb, 0x40,
	0x10, 0xc5, 0x81, 0xfe, 0xdb, 0x74, 0xe7, 0x1f, 0x11, 0x46, 0x63, 0x08, 0x17, 0xc9, 0x5e, 0xe4,
	0xb4, 0x87, 0xea, 0xc9, 0x78, 0xd1, 0x8b, 0x17, 0x13, 0x93, 0x4d, 0xea, 0xbd, 0xc8, 0x1c, 0x9a,
	0xc0, 0x2e, 0x2e, 0x4b, 0xfc, 0x3e, 0x7e, 0x52, 0x03, 0x0b, 0x95, 0xb6, 0xc7, 0xf7, 0x60, 0xe6,
	0xfd, 0xde, 0x2c, 0x84, 0x5d, 0x5d, 0x18, 0xaa, 0xaa, 0x9d, 0x68, 0x8c, 0xb6, 0x9a, 0x3f, 0xc2,
	0x7a, 0x3b, 0x3a, 0x78, 0x03, 0x2b, 0xd5, 0xd5, 0x05, 0x99, 0xc4, 0xcf, 0xfc, 0x9c, 0xc9, 0x51,
	0xf5, 0xbe, 0x21, 0x65, 0xa9, 0x4c, 0x82, 0xcc, 0xcf, 0xd7, 0x72, 0x54, 0x1c, 0x21, 0x7a, 0x25,
	0xfb, 0x5c, 0x55, 0xdb, 0xba, 0x90, 0xf4, 0xd5, 0x51, 0x6b, 0xf9, 0x13, 0xc4, 0xce, 0xeb, 0x7a,
	0xaf, 0x6d, 0xb4, 0x6a, 0x09, 0xef, 0x80, 0x4d, 0xb1, 0x6d, 0xe2, 0x67, 0x8b, 0xfc, 0xff, 0x86,
	0x89, 0x29, 0x56, 0xfe, 0x7d, 0xe3, 0x1f, 0x10, 0xbd, 0x68, 0x63, 0xf4, 0xf7, 0xbb, 0xa2, 0x71,
	0x23, 0x5e, 0xc3, 0xb2, 0xb5, 0xa5, 0xd2, 0x23, 0x94, 0x13, 0xbd, 0xdb, 0x54, 0xbb, 0x4f, 0x1a,
	0x90, 0x98, 0x74, 0x62, 0xd6, 0x60, 0x31, 0x6f, 0xc0, 0xaf, 0x20, 0x9e, 0xed, 0x75, 0x54, 0xfc,
	0x16, 0x2e, 0x24, 0xd9, 0xce, 0xa8, 0x29, 0x29, 0x84, 0x60, 0x5f, 0x0e, 0x31, 0x4b, 0x19, 0xec,
	0x4b, 0x1e, 0x41, 0x38, 0xfd, 0xe0, 0x46, 0x36, 0x3f, 0x3e, 0xfc, 0x7b, 0x23, 0x55, 0xe2, 0x03,
	0xb0, 0x43, 0x75, 0x8c, 0xc5, 0xe9, 0x19, 0x52, 0x14, 0x67, 0x57, 0xe0, 0x5e, 0x3f, 0x75, 0xc0,
	0xc0, 0x58, 0x9c, 0x56, 0x4d, 0x51, 0x9c, 0x53, 0x7a, 0x28, 0x80, 0x39, 0x8c, 0x7e, 0x2a, 0x14,
	0x47, 0xcc, 0xe9, 0xa5, 0x38, 0x46, 0xe4, 0x5e, 0xb1, 0x1a, 0x5e, 0xf6, 0xfe, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x91, 0x24, 0xe9, 0xcb, 0xeb, 0x01, 0x00, 0x00,
}
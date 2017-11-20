// Code generated by protoc-gen-go. DO NOT EDIT.
// source: umbrella.proto

/*
Package umbrella is a generated protocol buffer package.

It is generated from these files:
	umbrella.proto

It has these top-level messages:
	Umbrella
	Record
	GetAllUmbRequest
	GetAllumbResponse
	BorrowOneRequest
	BorrowOneResponse
	ReturnRequest
	ReturnResponse
	GetRecordsByStdnoRequest
	GetRecordsByStdnoResponse
	GetRecordsByNumRequest
	GetRecordsByNumResponse
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

type Record struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Stdno     string `protobuf:"bytes,2,opt,name=stdno" json:"stdno,omitempty"`
	Place     string `protobuf:"bytes,3,opt,name=place" json:"place,omitempty"`
	Number    string `protobuf:"bytes,4,opt,name=number" json:"number,omitempty"`
	HasReturn bool   `protobuf:"varint,5,opt,name=hasReturn" json:"hasReturn,omitempty"`
	CreateAt  int64  `protobuf:"varint,6,opt,name=createAt" json:"createAt,omitempty"`
	ReturnAt  int64  `protobuf:"varint,7,opt,name=returnAt" json:"returnAt,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Record) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Record) GetStdno() string {
	if m != nil {
		return m.Stdno
	}
	return ""
}

func (m *Record) GetPlace() string {
	if m != nil {
		return m.Place
	}
	return ""
}

func (m *Record) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Record) GetHasReturn() bool {
	if m != nil {
		return m.HasReturn
	}
	return false
}

func (m *Record) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *Record) GetReturnAt() int64 {
	if m != nil {
		return m.ReturnAt
	}
	return 0
}

type GetAllUmbRequest struct {
}

func (m *GetAllUmbRequest) Reset()                    { *m = GetAllUmbRequest{} }
func (m *GetAllUmbRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllUmbRequest) ProtoMessage()               {}
func (*GetAllUmbRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetAllumbResponse struct {
	Umbrellas []*Umbrella `protobuf:"bytes,1,rep,name=umbrellas" json:"umbrellas,omitempty"`
}

func (m *GetAllumbResponse) Reset()                    { *m = GetAllumbResponse{} }
func (m *GetAllumbResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAllumbResponse) ProtoMessage()               {}
func (*GetAllumbResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
func (*BorrowOneRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
func (*BorrowOneResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ReturnRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReturnRequest) Reset()                    { *m = ReturnRequest{} }
func (m *ReturnRequest) String() string            { return proto.CompactTextString(m) }
func (*ReturnRequest) ProtoMessage()               {}
func (*ReturnRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*ReturnResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type GetRecordsByStdnoRequest struct {
	Stdno string `protobuf:"bytes,1,opt,name=stdno" json:"stdno,omitempty"`
}

func (m *GetRecordsByStdnoRequest) Reset()                    { *m = GetRecordsByStdnoRequest{} }
func (m *GetRecordsByStdnoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRecordsByStdnoRequest) ProtoMessage()               {}
func (*GetRecordsByStdnoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetRecordsByStdnoRequest) GetStdno() string {
	if m != nil {
		return m.Stdno
	}
	return ""
}

type GetRecordsByStdnoResponse struct {
	Records []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *GetRecordsByStdnoResponse) Reset()                    { *m = GetRecordsByStdnoResponse{} }
func (m *GetRecordsByStdnoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRecordsByStdnoResponse) ProtoMessage()               {}
func (*GetRecordsByStdnoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetRecordsByStdnoResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type GetRecordsByNumRequest struct {
	Number string `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
}

func (m *GetRecordsByNumRequest) Reset()                    { *m = GetRecordsByNumRequest{} }
func (m *GetRecordsByNumRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRecordsByNumRequest) ProtoMessage()               {}
func (*GetRecordsByNumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetRecordsByNumRequest) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type GetRecordsByNumResponse struct {
	Records []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *GetRecordsByNumResponse) Reset()                    { *m = GetRecordsByNumResponse{} }
func (m *GetRecordsByNumResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRecordsByNumResponse) ProtoMessage()               {}
func (*GetRecordsByNumResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetRecordsByNumResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*Umbrella)(nil), "Umbrella")
	proto.RegisterType((*Record)(nil), "Record")
	proto.RegisterType((*GetAllUmbRequest)(nil), "GetAllUmbRequest")
	proto.RegisterType((*GetAllumbResponse)(nil), "GetAllumbResponse")
	proto.RegisterType((*BorrowOneRequest)(nil), "BorrowOneRequest")
	proto.RegisterType((*BorrowOneResponse)(nil), "BorrowOneResponse")
	proto.RegisterType((*ReturnRequest)(nil), "ReturnRequest")
	proto.RegisterType((*ReturnResponse)(nil), "ReturnResponse")
	proto.RegisterType((*GetRecordsByStdnoRequest)(nil), "GetRecordsByStdnoRequest")
	proto.RegisterType((*GetRecordsByStdnoResponse)(nil), "GetRecordsByStdnoResponse")
	proto.RegisterType((*GetRecordsByNumRequest)(nil), "GetRecordsByNumRequest")
	proto.RegisterType((*GetRecordsByNumResponse)(nil), "GetRecordsByNumResponse")
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
	GetRecordsByNum(ctx context.Context, in *GetRecordsByNumRequest, opts ...client.CallOption) (*GetRecordsByNumResponse, error)
	GetRecordsByStdno(ctx context.Context, in *GetRecordsByStdnoRequest, opts ...client.CallOption) (*GetRecordsByStdnoResponse, error)
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

func (c *lendClient) GetRecordsByNum(ctx context.Context, in *GetRecordsByNumRequest, opts ...client.CallOption) (*GetRecordsByNumResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Lend.GetRecordsByNum", in)
	out := new(GetRecordsByNumResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lendClient) GetRecordsByStdno(ctx context.Context, in *GetRecordsByStdnoRequest, opts ...client.CallOption) (*GetRecordsByStdnoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Lend.GetRecordsByStdno", in)
	out := new(GetRecordsByStdnoResponse)
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
	GetRecordsByNum(context.Context, *GetRecordsByNumRequest, *GetRecordsByNumResponse) error
	GetRecordsByStdno(context.Context, *GetRecordsByStdnoRequest, *GetRecordsByStdnoResponse) error
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

func (h *Lend) GetRecordsByNum(ctx context.Context, in *GetRecordsByNumRequest, out *GetRecordsByNumResponse) error {
	return h.LendHandler.GetRecordsByNum(ctx, in, out)
}

func (h *Lend) GetRecordsByStdno(ctx context.Context, in *GetRecordsByStdnoRequest, out *GetRecordsByStdnoResponse) error {
	return h.LendHandler.GetRecordsByStdno(ctx, in, out)
}

func init() { proto.RegisterFile("umbrella.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0xb4, 0x9d, 0xe6, 0xc3, 0x0f, 0x91, 0x26, 0x0f, 0xd4, 0x6e, 0x2d, 0x24, 0xcc, 0x5e, 0xf0,
	0x69, 0x55, 0x15, 0x4e, 0xa8, 0x42, 0x4a, 0x0f, 0xf4, 0x52, 0x81, 0xb4, 0xa8, 0xdc, 0xed, 0xf8,
	0x49, 0x54, 0xf2, 0x47, 0x58, 0xaf, 0x85, 0xf8, 0x5b, 0x9c, 0xf8, 0x79, 0xc8, 0x5e, 0x3b, 0x71,
	0x1c, 0x1b, 0x71, 0x9c, 0x79, 0x3b, 0x9e, 0x7d, 0x33, 0x2b, 0xc3, 0xb2, 0x4c, 0x23, 0x45, 0x49,
	0x12, 0x8a, 0x9d, 0xca, 0x75, 0xce, 0x3f, 0xc0, 0xe2, 0xb1, 0x61, 0xf0, 0x02, 0x66, 0x59, 0x99,
	0x46, 0xa4, 0x98, 0xed, 0xdb, 0x81, 0x2b, 0x1b, 0x54, 0xf1, 0x8a, 0x32, 0x4d, 0x31, 0x73, 0x7c,
	0x3b, 0x58, 0xc8, 0x06, 0xf1, 0xdf, 0x36, 0xcc, 0x24, 0x6d, 0x73, 0x15, 0xe3, 0x12, 0x9c, 0xa7,
	0xb8, 0x96, 0x4d, 0xa5, 0xf3, 0x14, 0xe3, 0x4b, 0x98, 0x16, 0x3a, 0xce, 0xf2, 0x5a, 0xe1, 0x4a,
	0x03, 0x2a, 0x76, 0x97, 0x84, 0x5b, 0x62, 0x13, 0xc3, 0xd6, 0xa0, 0x63, 0x7b, 0x76, 0x64, 0xfb,
	0x0a, 0xdc, 0xef, 0x61, 0x21, 0x49, 0x97, 0x2a, 0x63, 0xd3, 0xda, 0xf9, 0x40, 0xa0, 0x07, 0x8b,
	0xad, 0xa2, 0x50, 0xd3, 0x46, 0xb3, 0x99, 0x6f, 0x07, 0x13, 0xb9, 0xc7, 0xd5, 0x4c, 0xd5, 0xa7,
	0x36, 0x9a, 0xcd, 0xcd, 0xac, 0xc5, 0x1c, 0x61, 0x75, 0x4f, 0x7a, 0x93, 0x24, 0x8f, 0x69, 0x24,
	0xe9, 0x47, 0x49, 0x85, 0xe6, 0xb7, 0xb0, 0x36, 0x5c, 0x59, 0x71, 0xc5, 0x2e, 0xcf, 0x0a, 0xc2,
	0xb7, 0xe0, 0xb6, 0x59, 0x15, 0xcc, 0xf6, 0x27, 0xc1, 0xb3, 0x1b, 0x57, 0xb4, 0x59, 0xc9, 0xc3,
	0x8c, 0x7f, 0x83, 0xd5, 0x5d, 0xae, 0x54, 0xfe, 0xf3, 0x4b, 0x46, 0xcd, 0x17, 0x0f, 0xfb, 0xdb,
	0x83, 0xfb, 0x3b, 0xc3, 0xfb, 0x4f, 0xba, 0xfb, 0xf3, 0x17, 0xb0, 0xee, 0x7c, 0xd7, 0xdc, 0x8a,
	0xbf, 0x86, 0xe7, 0x26, 0x80, 0xd6, 0xa9, 0x97, 0x3c, 0x5f, 0xc1, 0xb2, 0x3d, 0xd0, 0x48, 0xae,
	0x81, 0xdd, 0x93, 0x36, 0x45, 0x15, 0x77, 0xbf, 0xbe, 0x56, 0x57, 0xf9, 0xe7, 0x3d, 0xf9, 0x47,
	0xb8, 0x1a, 0x50, 0x34, 0xb9, 0xbc, 0x81, 0xb9, 0x32, 0x93, 0x26, 0x95, 0xb9, 0x30, 0x27, 0x65,
	0xcb, 0xf3, 0x6b, 0xb8, 0xe8, 0xea, 0x3f, 0x97, 0x69, 0xeb, 0x37, 0xf2, 0xc4, 0xf8, 0x2d, 0x5c,
	0x9e, 0x28, 0xfe, 0xdb, 0xef, 0xe6, 0x8f, 0x03, 0x67, 0x0f, 0x94, 0xc5, 0xf8, 0x1e, 0xdc, 0x7d,
	0xb9, 0xb8, 0x16, 0xfd, 0xa2, 0x3d, 0x14, 0x27, 0x3d, 0x73, 0xab, 0x52, 0xed, 0x83, 0xc6, 0xb5,
	0xe8, 0x97, 0xe9, 0xa1, 0x38, 0xed, 0xc1, 0x42, 0x01, 0xae, 0x09, 0xba, 0x52, 0x2d, 0xc5, 0x51,
	0x2b, 0xde, 0xb9, 0xe8, 0x95, 0x60, 0xe1, 0x27, 0x38, 0xef, 0xad, 0x88, 0x97, 0x62, 0x38, 0x26,
	0x8f, 0x89, 0x91, 0x34, 0xb8, 0x85, 0x0f, 0xf5, 0x63, 0x3d, 0x2e, 0x07, 0xaf, 0xc4, 0x58, 0xc5,
	0x9e, 0x27, 0x46, 0xbb, 0xe4, 0x56, 0x34, 0xab, 0x7f, 0x03, 0xef, 0xfe, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x49, 0xe6, 0x02, 0x7f, 0x18, 0x04, 0x00, 0x00,
}

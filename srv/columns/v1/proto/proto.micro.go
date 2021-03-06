// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/columns/v1/proto/proto.proto

// import public "google/protobuf/timestamp.proto";

package ProtoColumns

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ProtoColumns service

type ProtoColumnsService interface {
	// 添回
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	// 获取列表
	GetList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResponse, error)
	// 获取单个信息
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	// 修改
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	// 批量删除
	BatchDelete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type protoColumnsService struct {
	c    client.Client
	name string
}

func NewProtoColumnsService(name string, c client.Client) ProtoColumnsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "ProtoColumns"
	}
	return &protoColumnsService{
		c:    c,
		name: name,
	}
}

func (c *protoColumnsService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoColumns.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoColumnsService) GetList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoColumns.GetList", in)
	out := new(GetListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoColumnsService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoColumns.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoColumnsService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoColumns.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoColumnsService) BatchDelete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoColumns.BatchDelete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProtoColumns service

type ProtoColumnsHandler interface {
	// 添回
	Add(context.Context, *AddRequest, *AddResponse) error
	// 获取列表
	GetList(context.Context, *GetListRequest, *GetListResponse) error
	// 获取单个信息
	Get(context.Context, *GetRequest, *GetResponse) error
	// 修改
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	// 批量删除
	BatchDelete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterProtoColumnsHandler(s server.Server, hdlr ProtoColumnsHandler, opts ...server.HandlerOption) error {
	type protoColumns interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		GetList(ctx context.Context, in *GetListRequest, out *GetListResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		BatchDelete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type ProtoColumns struct {
		protoColumns
	}
	h := &protoColumnsHandler{hdlr}
	return s.Handle(s.NewHandler(&ProtoColumns{h}, opts...))
}

type protoColumnsHandler struct {
	ProtoColumnsHandler
}

func (h *protoColumnsHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.ProtoColumnsHandler.Add(ctx, in, out)
}

func (h *protoColumnsHandler) GetList(ctx context.Context, in *GetListRequest, out *GetListResponse) error {
	return h.ProtoColumnsHandler.GetList(ctx, in, out)
}

func (h *protoColumnsHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.ProtoColumnsHandler.Get(ctx, in, out)
}

func (h *protoColumnsHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.ProtoColumnsHandler.Update(ctx, in, out)
}

func (h *protoColumnsHandler) BatchDelete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ProtoColumnsHandler.BatchDelete(ctx, in, out)
}

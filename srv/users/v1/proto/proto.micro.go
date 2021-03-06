// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/users/v1/proto/proto.proto

// import public "google/protobuf/timestamp.proto";

package ProtoUsers

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

// Client API for ProtoUsers service

type ProtoUsersService interface {
	// 添回
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	// 获取列表
	GetList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResponse, error)
	// 获取单个记录
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	// 修改
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	// 批量删除（假删除）
	BatchDelete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type protoUsersService struct {
	c    client.Client
	name string
}

func NewProtoUsersService(name string, c client.Client) ProtoUsersService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "ProtoUsers"
	}
	return &protoUsersService{
		c:    c,
		name: name,
	}
}

func (c *protoUsersService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoUsers.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoUsersService) GetList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoUsers.GetList", in)
	out := new(GetListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoUsersService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoUsers.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoUsersService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoUsers.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoUsersService) BatchDelete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoUsers.BatchDelete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProtoUsers service

type ProtoUsersHandler interface {
	// 添回
	Add(context.Context, *AddRequest, *AddResponse) error
	// 获取列表
	GetList(context.Context, *GetListRequest, *GetListResponse) error
	// 获取单个记录
	Get(context.Context, *GetRequest, *GetResponse) error
	// 修改
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	// 批量删除（假删除）
	BatchDelete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterProtoUsersHandler(s server.Server, hdlr ProtoUsersHandler, opts ...server.HandlerOption) error {
	type protoUsers interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		GetList(ctx context.Context, in *GetListRequest, out *GetListResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		BatchDelete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type ProtoUsers struct {
		protoUsers
	}
	h := &protoUsersHandler{hdlr}
	return s.Handle(s.NewHandler(&ProtoUsers{h}, opts...))
}

type protoUsersHandler struct {
	ProtoUsersHandler
}

func (h *protoUsersHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.ProtoUsersHandler.Add(ctx, in, out)
}

func (h *protoUsersHandler) GetList(ctx context.Context, in *GetListRequest, out *GetListResponse) error {
	return h.ProtoUsersHandler.GetList(ctx, in, out)
}

func (h *protoUsersHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.ProtoUsersHandler.Get(ctx, in, out)
}

func (h *protoUsersHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.ProtoUsersHandler.Update(ctx, in, out)
}

func (h *protoUsersHandler) BatchDelete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ProtoUsersHandler.BatchDelete(ctx, in, out)
}

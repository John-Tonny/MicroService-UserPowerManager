// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/historyuserlogin/v1/proto/proto.proto

package ProtoHistoryUserLogin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/idoall/MicroService-UserPowerManager/srv/users/v1/proto"
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

// Client API for ProtoHistoryUserLogin service

type ProtoHistoryUserLoginService interface {
	// 添回一条记录
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
}

type protoHistoryUserLoginService struct {
	c    client.Client
	name string
}

func NewProtoHistoryUserLoginService(name string, c client.Client) ProtoHistoryUserLoginService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "ProtoHistoryUserLogin"
	}
	return &protoHistoryUserLoginService{
		c:    c,
		name: name,
	}
}

func (c *protoHistoryUserLoginService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "ProtoHistoryUserLogin.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProtoHistoryUserLogin service

type ProtoHistoryUserLoginHandler interface {
	// 添回一条记录
	Add(context.Context, *AddRequest, *AddResponse) error
}

func RegisterProtoHistoryUserLoginHandler(s server.Server, hdlr ProtoHistoryUserLoginHandler, opts ...server.HandlerOption) error {
	type protoHistoryUserLogin interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
	}
	type ProtoHistoryUserLogin struct {
		protoHistoryUserLogin
	}
	h := &protoHistoryUserLoginHandler{hdlr}
	return s.Handle(s.NewHandler(&ProtoHistoryUserLogin{h}, opts...))
}

type protoHistoryUserLoginHandler struct {
	ProtoHistoryUserLoginHandler
}

func (h *protoHistoryUserLoginHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.ProtoHistoryUserLoginHandler.Add(ctx, in, out)
}

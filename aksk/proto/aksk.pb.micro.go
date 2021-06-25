// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: aksk/proto/aksk.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for AkSk service

func NewAkSkEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AkSk service

type AkSkService interface {
	DownloadAkSk(ctx context.Context, in *GetAkSkRequest, opts ...client.CallOption) (*GetAkSkResponse, error)
	GetAkSk(ctx context.Context, in *GetAkSkRequest, opts ...client.CallOption) (*GetAkSkResponse, error)
	CreateAkSk(ctx context.Context, in *AkSkCreateRequest, opts ...client.CallOption) (*AkSkCreateResponse, error)
	DeleteAkSk(ctx context.Context, in *DeleteAkSkRequest, opts ...client.CallOption) (*DeleteAkSkResponse, error)
}

type akSkService struct {
	c    client.Client
	name string
}

func NewAkSkService(name string, c client.Client) AkSkService {
	return &akSkService{
		c:    c,
		name: name,
	}
}

func (c *akSkService) DownloadAkSk(ctx context.Context, in *GetAkSkRequest, opts ...client.CallOption) (*GetAkSkResponse, error) {
	req := c.c.NewRequest(c.name, "AkSk.DownloadAkSk", in)
	out := new(GetAkSkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *akSkService) GetAkSk(ctx context.Context, in *GetAkSkRequest, opts ...client.CallOption) (*GetAkSkResponse, error) {
	req := c.c.NewRequest(c.name, "AkSk.GetAkSk", in)
	out := new(GetAkSkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *akSkService) CreateAkSk(ctx context.Context, in *AkSkCreateRequest, opts ...client.CallOption) (*AkSkCreateResponse, error) {
	req := c.c.NewRequest(c.name, "AkSk.CreateAkSk", in)
	out := new(AkSkCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *akSkService) DeleteAkSk(ctx context.Context, in *DeleteAkSkRequest, opts ...client.CallOption) (*DeleteAkSkResponse, error) {
	req := c.c.NewRequest(c.name, "AkSk.DeleteAkSk", in)
	out := new(DeleteAkSkResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AkSk service

type AkSkHandler interface {
	DownloadAkSk(context.Context, *GetAkSkRequest, *GetAkSkResponse) error
	GetAkSk(context.Context, *GetAkSkRequest, *GetAkSkResponse) error
	CreateAkSk(context.Context, *AkSkCreateRequest, *AkSkCreateResponse) error
	DeleteAkSk(context.Context, *DeleteAkSkRequest, *DeleteAkSkResponse) error
}

func RegisterAkSkHandler(s server.Server, hdlr AkSkHandler, opts ...server.HandlerOption) error {
	type akSk interface {
		DownloadAkSk(ctx context.Context, in *GetAkSkRequest, out *GetAkSkResponse) error
		GetAkSk(ctx context.Context, in *GetAkSkRequest, out *GetAkSkResponse) error
		CreateAkSk(ctx context.Context, in *AkSkCreateRequest, out *AkSkCreateResponse) error
		DeleteAkSk(ctx context.Context, in *DeleteAkSkRequest, out *DeleteAkSkResponse) error
	}
	type AkSk struct {
		akSk
	}
	h := &akSkHandler{hdlr}
	return s.Handle(s.NewHandler(&AkSk{h}, opts...))
}

type akSkHandler struct {
	AkSkHandler
}

func (h *akSkHandler) DownloadAkSk(ctx context.Context, in *GetAkSkRequest, out *GetAkSkResponse) error {
	return h.AkSkHandler.DownloadAkSk(ctx, in, out)
}

func (h *akSkHandler) GetAkSk(ctx context.Context, in *GetAkSkRequest, out *GetAkSkResponse) error {
	return h.AkSkHandler.GetAkSk(ctx, in, out)
}

func (h *akSkHandler) CreateAkSk(ctx context.Context, in *AkSkCreateRequest, out *AkSkCreateResponse) error {
	return h.AkSkHandler.CreateAkSk(ctx, in, out)
}

func (h *akSkHandler) DeleteAkSk(ctx context.Context, in *DeleteAkSkRequest, out *DeleteAkSkResponse) error {
	return h.AkSkHandler.DeleteAkSk(ctx, in, out)
}

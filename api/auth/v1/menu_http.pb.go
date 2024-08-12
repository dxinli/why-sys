// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v4.24.2
// source: auth/v1/menu.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMenuCreateMenu = "/auth.v1.Menu/CreateMenu"
const OperationMenuListMenu = "/auth.v1.Menu/ListMenu"

type MenuHTTPServer interface {
	CreateMenu(context.Context, *MenuDetail) (*CreateMenuResponse, error)
	ListMenu(context.Context, *ListMenuRequest) (*MenuDetail, error)
}

func RegisterMenuHTTPServer(s *http.Server, srv MenuHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/menu", _Menu_CreateMenu0_HTTP_Handler(srv))
	r.POST("/v1/menus", _Menu_ListMenu0_HTTP_Handler(srv))
}

func _Menu_CreateMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MenuDetail
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuCreateMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMenu(ctx, req.(*MenuDetail))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateMenuResponse)
		return ctx.Result(200, reply)
	}
}

func _Menu_ListMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMenuRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuListMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMenu(ctx, req.(*ListMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MenuDetail)
		return ctx.Result(200, reply)
	}
}

type MenuHTTPClient interface {
	CreateMenu(ctx context.Context, req *MenuDetail, opts ...http.CallOption) (rsp *CreateMenuResponse, err error)
	ListMenu(ctx context.Context, req *ListMenuRequest, opts ...http.CallOption) (rsp *MenuDetail, err error)
}

type MenuHTTPClientImpl struct {
	cc *http.Client
}

func NewMenuHTTPClient(client *http.Client) MenuHTTPClient {
	return &MenuHTTPClientImpl{client}
}

func (c *MenuHTTPClientImpl) CreateMenu(ctx context.Context, in *MenuDetail, opts ...http.CallOption) (*CreateMenuResponse, error) {
	var out CreateMenuResponse
	pattern := "/v1/menu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuCreateMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuHTTPClientImpl) ListMenu(ctx context.Context, in *ListMenuRequest, opts ...http.CallOption) (*MenuDetail, error) {
	var out MenuDetail
	pattern := "/v1/menus"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuListMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.2
// source: auth/v1/menu.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Menu_CreateMenu_FullMethodName = "/auth.v1.Menu/CreateMenu"
	Menu_ListMenu_FullMethodName   = "/auth.v1.Menu/ListMenu"
)

// MenuClient is the client API for Menu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuClient interface {
	CreateMenu(ctx context.Context, in *MenuDetail, opts ...grpc.CallOption) (*CreateMenuResponse, error)
	ListMenu(ctx context.Context, in *ListMenuRequest, opts ...grpc.CallOption) (*MenuDetail, error)
}

type menuClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuClient(cc grpc.ClientConnInterface) MenuClient {
	return &menuClient{cc}
}

func (c *menuClient) CreateMenu(ctx context.Context, in *MenuDetail, opts ...grpc.CallOption) (*CreateMenuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMenuResponse)
	err := c.cc.Invoke(ctx, Menu_CreateMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuClient) ListMenu(ctx context.Context, in *ListMenuRequest, opts ...grpc.CallOption) (*MenuDetail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MenuDetail)
	err := c.cc.Invoke(ctx, Menu_ListMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServer is the server API for Menu service.
// All implementations must embed UnimplementedMenuServer
// for forward compatibility.
type MenuServer interface {
	CreateMenu(context.Context, *MenuDetail) (*CreateMenuResponse, error)
	ListMenu(context.Context, *ListMenuRequest) (*MenuDetail, error)
	mustEmbedUnimplementedMenuServer()
}

// UnimplementedMenuServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMenuServer struct{}

func (UnimplementedMenuServer) CreateMenu(context.Context, *MenuDetail) (*CreateMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenu not implemented")
}
func (UnimplementedMenuServer) ListMenu(context.Context, *ListMenuRequest) (*MenuDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMenu not implemented")
}
func (UnimplementedMenuServer) mustEmbedUnimplementedMenuServer() {}
func (UnimplementedMenuServer) testEmbeddedByValue()              {}

// UnsafeMenuServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MenuServer will
// result in compilation errors.
type UnsafeMenuServer interface {
	mustEmbedUnimplementedMenuServer()
}

func RegisterMenuServer(s grpc.ServiceRegistrar, srv MenuServer) {
	// If the following call pancis, it indicates UnimplementedMenuServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Menu_ServiceDesc, srv)
}

func _Menu_CreateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServer).CreateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Menu_CreateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServer).CreateMenu(ctx, req.(*MenuDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _Menu_ListMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServer).ListMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Menu_ListMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServer).ListMenu(ctx, req.(*ListMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Menu_ServiceDesc is the grpc.ServiceDesc for Menu service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Menu_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.Menu",
	HandlerType: (*MenuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMenu",
			Handler:    _Menu_CreateMenu_Handler,
		},
		{
			MethodName: "ListMenu",
			Handler:    _Menu_ListMenu_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/menu.proto",
}

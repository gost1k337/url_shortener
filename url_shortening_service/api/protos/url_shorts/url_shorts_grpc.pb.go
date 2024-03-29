// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: url_shorts.proto

package url_shorts

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UrlShortsClient is the client API for UrlShorts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlShortsClient interface {
	Create(ctx context.Context, in *CreateUrlShortRequest, opts ...grpc.CallOption) (*CreateUrlShortResponse, error)
	Get(ctx context.Context, in *GetUrlShortRequest, opts ...grpc.CallOption) (*GetUrlShortResponse, error)
}

type urlShortsClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlShortsClient(cc grpc.ClientConnInterface) UrlShortsClient {
	return &urlShortsClient{cc}
}

func (c *urlShortsClient) Create(ctx context.Context, in *CreateUrlShortRequest, opts ...grpc.CallOption) (*CreateUrlShortResponse, error) {
	out := new(CreateUrlShortResponse)
	err := c.cc.Invoke(ctx, "/url_shorts.UrlShorts/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortsClient) Get(ctx context.Context, in *GetUrlShortRequest, opts ...grpc.CallOption) (*GetUrlShortResponse, error) {
	out := new(GetUrlShortResponse)
	err := c.cc.Invoke(ctx, "/url_shorts.UrlShorts/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlShortsServer is the server API for UrlShorts service.
// All implementations must embed UnimplementedUrlShortsServer
// for forward compatibility
type UrlShortsServer interface {
	Create(context.Context, *CreateUrlShortRequest) (*CreateUrlShortResponse, error)
	Get(context.Context, *GetUrlShortRequest) (*GetUrlShortResponse, error)
	mustEmbedUnimplementedUrlShortsServer()
}

// UnimplementedUrlShortsServer must be embedded to have forward compatible implementations.
type UnimplementedUrlShortsServer struct {
}

func (UnimplementedUrlShortsServer) Create(context.Context, *CreateUrlShortRequest) (*CreateUrlShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUrlShortsServer) Get(context.Context, *GetUrlShortRequest) (*GetUrlShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUrlShortsServer) mustEmbedUnimplementedUrlShortsServer() {}

// UnsafeUrlShortsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlShortsServer will
// result in compilation errors.
type UnsafeUrlShortsServer interface {
	mustEmbedUnimplementedUrlShortsServer()
}

func RegisterUrlShortsServer(s grpc.ServiceRegistrar, srv UrlShortsServer) {
	s.RegisterService(&UrlShorts_ServiceDesc, srv)
}

func _UrlShorts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUrlShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/url_shorts.UrlShorts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortsServer).Create(ctx, req.(*CreateUrlShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlShorts_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/url_shorts.UrlShorts/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortsServer).Get(ctx, req.(*GetUrlShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlShorts_ServiceDesc is the grpc.ServiceDesc for UrlShorts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlShorts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "url_shorts.UrlShorts",
	HandlerType: (*UrlShortsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UrlShorts_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UrlShorts_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "url_shorts.proto",
}

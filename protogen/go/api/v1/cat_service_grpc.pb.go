// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/v1/cat_service.proto

package v1

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

// CatServiceClient is the client API for CatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatServiceClient interface {
	GetFact(ctx context.Context, in *GetFactRequest, opts ...grpc.CallOption) (*GetFactResponse, error)
	CreateCat(ctx context.Context, in *CreateCatRequest, opts ...grpc.CallOption) (*CreateCatResponse, error)
	ListCats(ctx context.Context, in *ListCatsRequest, opts ...grpc.CallOption) (*ListCatsResponse, error)
}

type catServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatServiceClient(cc grpc.ClientConnInterface) CatServiceClient {
	return &catServiceClient{cc}
}

func (c *catServiceClient) GetFact(ctx context.Context, in *GetFactRequest, opts ...grpc.CallOption) (*GetFactResponse, error) {
	out := new(GetFactResponse)
	err := c.cc.Invoke(ctx, "/api.v1.CatService/GetFact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catServiceClient) CreateCat(ctx context.Context, in *CreateCatRequest, opts ...grpc.CallOption) (*CreateCatResponse, error) {
	out := new(CreateCatResponse)
	err := c.cc.Invoke(ctx, "/api.v1.CatService/CreateCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catServiceClient) ListCats(ctx context.Context, in *ListCatsRequest, opts ...grpc.CallOption) (*ListCatsResponse, error) {
	out := new(ListCatsResponse)
	err := c.cc.Invoke(ctx, "/api.v1.CatService/ListCats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatServiceServer is the server API for CatService service.
// All implementations should embed UnimplementedCatServiceServer
// for forward compatibility
type CatServiceServer interface {
	GetFact(context.Context, *GetFactRequest) (*GetFactResponse, error)
	CreateCat(context.Context, *CreateCatRequest) (*CreateCatResponse, error)
	ListCats(context.Context, *ListCatsRequest) (*ListCatsResponse, error)
}

// UnimplementedCatServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCatServiceServer struct {
}

func (UnimplementedCatServiceServer) GetFact(context.Context, *GetFactRequest) (*GetFactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFact not implemented")
}
func (UnimplementedCatServiceServer) CreateCat(context.Context, *CreateCatRequest) (*CreateCatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCat not implemented")
}
func (UnimplementedCatServiceServer) ListCats(context.Context, *ListCatsRequest) (*ListCatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCats not implemented")
}

// UnsafeCatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatServiceServer will
// result in compilation errors.
type UnsafeCatServiceServer interface {
	mustEmbedUnimplementedCatServiceServer()
}

func RegisterCatServiceServer(s grpc.ServiceRegistrar, srv CatServiceServer) {
	s.RegisterService(&CatService_ServiceDesc, srv)
}

func _CatService_GetFact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).GetFact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.CatService/GetFact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).GetFact(ctx, req.(*GetFactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatService_CreateCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).CreateCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.CatService/CreateCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).CreateCat(ctx, req.(*CreateCatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatService_ListCats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServiceServer).ListCats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.CatService/ListCats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServiceServer).ListCats(ctx, req.(*ListCatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatService_ServiceDesc is the grpc.ServiceDesc for CatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.CatService",
	HandlerType: (*CatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFact",
			Handler:    _CatService_GetFact_Handler,
		},
		{
			MethodName: "CreateCat",
			Handler:    _CatService_CreateCat_Handler,
		},
		{
			MethodName: "ListCats",
			Handler:    _CatService_ListCats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/cat_service.proto",
}

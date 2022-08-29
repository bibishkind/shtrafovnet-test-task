// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: search.proto

package pb

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

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchClient interface {
	GetCompanyByINN(ctx context.Context, in *CompanyByINNRequest, opts ...grpc.CallOption) (*CompanyByINNResponse, error)
}

type searchClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchClient(cc grpc.ClientConnInterface) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) GetCompanyByINN(ctx context.Context, in *CompanyByINNRequest, opts ...grpc.CallOption) (*CompanyByINNResponse, error) {
	out := new(CompanyByINNResponse)
	err := c.cc.Invoke(ctx, "/Search/GetCompanyByINN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServer is the server API for Search service.
// All implementations must embed UnimplementedSearchServer
// for forward compatibility
type SearchServer interface {
	GetCompanyByINN(context.Context, *CompanyByINNRequest) (*CompanyByINNResponse, error)
	mustEmbedUnimplementedSearchServer()
}

// UnimplementedSearchServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServer struct {
}

func (UnimplementedSearchServer) GetCompanyByINN(context.Context, *CompanyByINNRequest) (*CompanyByINNResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyByINN not implemented")
}
func (UnimplementedSearchServer) mustEmbedUnimplementedSearchServer() {}

// UnsafeSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServer will
// result in compilation errors.
type UnsafeSearchServer interface {
	mustEmbedUnimplementedSearchServer()
}

func RegisterSearchServer(s grpc.ServiceRegistrar, srv SearchServer) {
	s.RegisterService(&Search_ServiceDesc, srv)
}

func _Search_GetCompanyByINN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyByINNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).GetCompanyByINN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Search/GetCompanyByINN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).GetCompanyByINN(ctx, req.(*CompanyByINNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Search_ServiceDesc is the grpc.ServiceDesc for Search service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Search_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompanyByINN",
			Handler:    _Search_GetCompanyByINN_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}

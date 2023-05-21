// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: pb/svc/apid/apid.proto

package apid

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

// ApidServiceClient is the client API for ApidService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApidServiceClient interface {
	GetArticleList(ctx context.Context, in *GetArticleListReq, opts ...grpc.CallOption) (*GetArticleListRes, error)
	AddArticle(ctx context.Context, in *AddArticleReq, opts ...grpc.CallOption) (*AddArticleRes, error)
}

type apidServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApidServiceClient(cc grpc.ClientConnInterface) ApidServiceClient {
	return &apidServiceClient{cc}
}

func (c *apidServiceClient) GetArticleList(ctx context.Context, in *GetArticleListReq, opts ...grpc.CallOption) (*GetArticleListRes, error) {
	out := new(GetArticleListRes)
	err := c.cc.Invoke(ctx, "/pb.svc.apid.ApidService/GetArticleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apidServiceClient) AddArticle(ctx context.Context, in *AddArticleReq, opts ...grpc.CallOption) (*AddArticleRes, error) {
	out := new(AddArticleRes)
	err := c.cc.Invoke(ctx, "/pb.svc.apid.ApidService/AddArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApidServiceServer is the server API for ApidService service.
// All implementations must embed UnimplementedApidServiceServer
// for forward compatibility
type ApidServiceServer interface {
	GetArticleList(context.Context, *GetArticleListReq) (*GetArticleListRes, error)
	AddArticle(context.Context, *AddArticleReq) (*AddArticleRes, error)
	mustEmbedUnimplementedApidServiceServer()
}

// UnimplementedApidServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApidServiceServer struct {
}

func (UnimplementedApidServiceServer) GetArticleList(context.Context, *GetArticleListReq) (*GetArticleListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleList not implemented")
}
func (UnimplementedApidServiceServer) AddArticle(context.Context, *AddArticleReq) (*AddArticleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddArticle not implemented")
}
func (UnimplementedApidServiceServer) mustEmbedUnimplementedApidServiceServer() {}

// UnsafeApidServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApidServiceServer will
// result in compilation errors.
type UnsafeApidServiceServer interface {
	mustEmbedUnimplementedApidServiceServer()
}

func RegisterApidServiceServer(s grpc.ServiceRegistrar, srv ApidServiceServer) {
	s.RegisterService(&ApidService_ServiceDesc, srv)
}

func _ApidService_GetArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApidServiceServer).GetArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.svc.apid.ApidService/GetArticleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApidServiceServer).GetArticleList(ctx, req.(*GetArticleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApidService_AddArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApidServiceServer).AddArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.svc.apid.ApidService/AddArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApidServiceServer).AddArticle(ctx, req.(*AddArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ApidService_ServiceDesc is the grpc.ServiceDesc for ApidService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApidService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.svc.apid.ApidService",
	HandlerType: (*ApidServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArticleList",
			Handler:    _ApidService_GetArticleList_Handler,
		},
		{
			MethodName: "AddArticle",
			Handler:    _ApidService_AddArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/svc/apid/apid.proto",
}
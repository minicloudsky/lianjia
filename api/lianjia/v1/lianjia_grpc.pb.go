// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: lianjia/v1/lianjia.proto

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

// LianjiaClient is the client API for Lianjia service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LianjiaClient interface {
	ListErshoufang(ctx context.Context, in *ListErshoufangRequest, opts ...grpc.CallOption) (*ListErshoufangReply, error)
}

type lianjiaClient struct {
	cc grpc.ClientConnInterface
}

func NewLianjiaClient(cc grpc.ClientConnInterface) LianjiaClient {
	return &lianjiaClient{cc}
}

func (c *lianjiaClient) ListErshoufang(ctx context.Context, in *ListErshoufangRequest, opts ...grpc.CallOption) (*ListErshoufangReply, error) {
	out := new(ListErshoufangReply)
	err := c.cc.Invoke(ctx, "/lianjia.v1.Lianjia/ListErshoufang", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LianjiaServer is the server API for Lianjia service.
// All implementations must embed UnimplementedLianjiaServer
// for forward compatibility
type LianjiaServer interface {
	ListErshoufang(context.Context, *ListErshoufangRequest) (*ListErshoufangReply, error)
	mustEmbedUnimplementedLianjiaServer()
}

// UnimplementedLianjiaServer must be embedded to have forward compatible implementations.
type UnimplementedLianjiaServer struct {
}

func (UnimplementedLianjiaServer) ListErshoufang(context.Context, *ListErshoufangRequest) (*ListErshoufangReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListErshoufang not implemented")
}
func (UnimplementedLianjiaServer) mustEmbedUnimplementedLianjiaServer() {}

// UnsafeLianjiaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LianjiaServer will
// result in compilation errors.
type UnsafeLianjiaServer interface {
	mustEmbedUnimplementedLianjiaServer()
}

func RegisterLianjiaServer(s grpc.ServiceRegistrar, srv LianjiaServer) {
	s.RegisterService(&Lianjia_ServiceDesc, srv)
}

func _Lianjia_ListErshoufang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListErshoufangRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LianjiaServer).ListErshoufang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lianjia.v1.Lianjia/ListErshoufang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LianjiaServer).ListErshoufang(ctx, req.(*ListErshoufangRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Lianjia_ServiceDesc is the grpc.ServiceDesc for Lianjia service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lianjia_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lianjia.v1.Lianjia",
	HandlerType: (*LianjiaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListErshoufang",
			Handler:    _Lianjia_ListErshoufang_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lianjia/v1/lianjia.proto",
}
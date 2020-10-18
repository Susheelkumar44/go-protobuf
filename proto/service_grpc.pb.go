// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// StoreDataClient is the client API for StoreData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreDataClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type storeDataClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreDataClient(cc grpc.ClientConnInterface) StoreDataClient {
	return &storeDataClient{cc}
}

func (c *storeDataClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.storeData/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreDataServer is the server API for StoreData service.
// All implementations must embed UnimplementedStoreDataServer
// for forward compatibility
type StoreDataServer interface {
	Add(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedStoreDataServer()
}

// UnimplementedStoreDataServer must be embedded to have forward compatible implementations.
type UnimplementedStoreDataServer struct {
}

func (UnimplementedStoreDataServer) Add(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedStoreDataServer) mustEmbedUnimplementedStoreDataServer() {}

// UnsafeStoreDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreDataServer will
// result in compilation errors.
type UnsafeStoreDataServer interface {
	mustEmbedUnimplementedStoreDataServer()
}

func RegisterStoreDataServer(s *grpc.Server, srv StoreDataServer) {
	s.RegisterService(&_StoreData_serviceDesc, srv)
}

func _StoreData_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreDataServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.storeData/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreDataServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.storeData",
	HandlerType: (*StoreDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _StoreData_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}

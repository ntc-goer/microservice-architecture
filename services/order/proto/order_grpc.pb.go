// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: order.proto

package micro_pj

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	OrderService_Order_FullMethodName = "/proto.OrderService/Order"
	OrderService_Check_FullMethodName = "/proto.OrderService/Check"
	OrderService_Watch_FullMethodName = "/proto.OrderService/Watch"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	Check(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (OrderService_WatchClient, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, OrderService_Order_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Check(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, OrderService_Check_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (OrderService_WatchClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], OrderService_Watch_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceWatchClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_WatchClient interface {
	Recv() (*HealthCheckResponse, error)
	grpc.ClientStream
}

type orderServiceWatchClient struct {
	grpc.ClientStream
}

func (x *orderServiceWatchClient) Recv() (*HealthCheckResponse, error) {
	m := new(HealthCheckResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	Order(context.Context, *OrderRequest) (*OrderResponse, error)
	Check(context.Context, *emptypb.Empty) (*HealthCheckResponse, error)
	Watch(*HealthCheckRequest, OrderService_WatchServer) error
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) Order(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Order not implemented")
}
func (UnimplementedOrderServiceServer) Check(context.Context, *emptypb.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedOrderServiceServer) Watch(*HealthCheckRequest, OrderService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_Order_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Order(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Check(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HealthCheckRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).Watch(m, &orderServiceWatchServer{ServerStream: stream})
}

type OrderService_WatchServer interface {
	Send(*HealthCheckResponse) error
	grpc.ServerStream
}

type orderServiceWatchServer struct {
	grpc.ServerStream
}

func (x *orderServiceWatchServer) Send(m *HealthCheckResponse) error {
	return x.ServerStream.SendMsg(m)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Order",
			Handler:    _OrderService_Order_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _OrderService_Check_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _OrderService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "order.proto",
}

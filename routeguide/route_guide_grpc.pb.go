// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: route_guide.proto

package routeguide

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

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuideClient interface {
	GetSum(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Sum, error)
	GetSumFromPoints(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_GetSumFromPointsClient, error)
	GetYFromFunction(ctx context.Context, in *X, opts ...grpc.CallOption) (RouteGuide_GetYFromFunctionClient, error)
}

type routeGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideClient(cc grpc.ClientConnInterface) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetSum(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Sum, error) {
	out := new(Sum)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/getSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideClient) GetSumFromPoints(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_GetSumFromPointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[0], "/routeguide.RouteGuide/getSumFromPoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideGetSumFromPointsClient{stream}
	return x, nil
}

type RouteGuide_GetSumFromPointsClient interface {
	Send(*Point) error
	CloseAndRecv() (*Sum, error)
	grpc.ClientStream
}

type routeGuideGetSumFromPointsClient struct {
	grpc.ClientStream
}

func (x *routeGuideGetSumFromPointsClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideGetSumFromPointsClient) CloseAndRecv() (*Sum, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Sum)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) GetYFromFunction(ctx context.Context, in *X, opts ...grpc.CallOption) (RouteGuide_GetYFromFunctionClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[1], "/routeguide.RouteGuide/getYFromFunction", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideGetYFromFunctionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_GetYFromFunctionClient interface {
	Recv() (*Y, error)
	grpc.ClientStream
}

type routeGuideGetYFromFunctionClient struct {
	grpc.ClientStream
}

func (x *routeGuideGetYFromFunctionClient) Recv() (*Y, error) {
	m := new(Y)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideServer is the server API for RouteGuide service.
// All implementations must embed UnimplementedRouteGuideServer
// for forward compatibility
type RouteGuideServer interface {
	GetSum(context.Context, *Point) (*Sum, error)
	GetSumFromPoints(RouteGuide_GetSumFromPointsServer) error
	GetYFromFunction(*X, RouteGuide_GetYFromFunctionServer) error
	mustEmbedUnimplementedRouteGuideServer()
}

// UnimplementedRouteGuideServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (UnimplementedRouteGuideServer) GetSum(context.Context, *Point) (*Sum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSum not implemented")
}
func (UnimplementedRouteGuideServer) GetSumFromPoints(RouteGuide_GetSumFromPointsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSumFromPoints not implemented")
}
func (UnimplementedRouteGuideServer) GetYFromFunction(*X, RouteGuide_GetYFromFunctionServer) error {
	return status.Errorf(codes.Unimplemented, "method GetYFromFunction not implemented")
}
func (UnimplementedRouteGuideServer) mustEmbedUnimplementedRouteGuideServer() {}

// UnsafeRouteGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuideServer will
// result in compilation errors.
type UnsafeRouteGuideServer interface {
	mustEmbedUnimplementedRouteGuideServer()
}

func RegisterRouteGuideServer(s grpc.ServiceRegistrar, srv RouteGuideServer) {
	s.RegisterService(&RouteGuide_ServiceDesc, srv)
}

func _RouteGuide_GetSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/getSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetSum(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_GetSumFromPoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).GetSumFromPoints(&routeGuideGetSumFromPointsServer{stream})
}

type RouteGuide_GetSumFromPointsServer interface {
	SendAndClose(*Sum) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideGetSumFromPointsServer struct {
	grpc.ServerStream
}

func (x *routeGuideGetSumFromPointsServer) SendAndClose(m *Sum) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideGetSumFromPointsServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_GetYFromFunction_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(X)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).GetYFromFunction(m, &routeGuideGetYFromFunctionServer{stream})
}

type RouteGuide_GetYFromFunctionServer interface {
	Send(*Y) error
	grpc.ServerStream
}

type routeGuideGetYFromFunctionServer struct {
	grpc.ServerStream
}

func (x *routeGuideGetYFromFunctionServer) Send(m *Y) error {
	return x.ServerStream.SendMsg(m)
}

// RouteGuide_ServiceDesc is the grpc.ServiceDesc for RouteGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSum",
			Handler:    _RouteGuide_GetSum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getSumFromPoints",
			Handler:       _RouteGuide_GetSumFromPoints_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "getYFromFunction",
			Handler:       _RouteGuide_GetYFromFunction_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "route_guide.proto",
}
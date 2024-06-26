// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: flogram.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FlotgServiceClient is the client API for FlotgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlotgServiceClient interface {
	GetChats(ctx context.Context, in *FlotgGetChatsRequest, opts ...grpc.CallOption) (FlotgService_GetChatsClient, error)
	SetMonitoring(ctx context.Context, in *FlotgMonitor, opts ...grpc.CallOption) (*FlotgMonitor, error)
	GetMessages(ctx context.Context, in *FlotgGetMessagesRequest, opts ...grpc.CallOption) (FlotgService_GetMessagesClient, error)
}

type flotgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlotgServiceClient(cc grpc.ClientConnInterface) FlotgServiceClient {
	return &flotgServiceClient{cc}
}

func (c *flotgServiceClient) GetChats(ctx context.Context, in *FlotgGetChatsRequest, opts ...grpc.CallOption) (FlotgService_GetChatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlotgService_ServiceDesc.Streams[0], "/FlotgService/GetChats", opts...)
	if err != nil {
		return nil, err
	}
	x := &flotgServiceGetChatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlotgService_GetChatsClient interface {
	Recv() (*FlotgMonitor, error)
	grpc.ClientStream
}

type flotgServiceGetChatsClient struct {
	grpc.ClientStream
}

func (x *flotgServiceGetChatsClient) Recv() (*FlotgMonitor, error) {
	m := new(FlotgMonitor)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flotgServiceClient) SetMonitoring(ctx context.Context, in *FlotgMonitor, opts ...grpc.CallOption) (*FlotgMonitor, error) {
	out := new(FlotgMonitor)
	err := c.cc.Invoke(ctx, "/FlotgService/SetMonitoring", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flotgServiceClient) GetMessages(ctx context.Context, in *FlotgGetMessagesRequest, opts ...grpc.CallOption) (FlotgService_GetMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlotgService_ServiceDesc.Streams[1], "/FlotgService/GetMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &flotgServiceGetMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FlotgService_GetMessagesClient interface {
	Recv() (*FLO_MESSAGE, error)
	grpc.ClientStream
}

type flotgServiceGetMessagesClient struct {
	grpc.ClientStream
}

func (x *flotgServiceGetMessagesClient) Recv() (*FLO_MESSAGE, error) {
	m := new(FLO_MESSAGE)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FlotgServiceServer is the server API for FlotgService service.
// All implementations must embed UnimplementedFlotgServiceServer
// for forward compatibility
type FlotgServiceServer interface {
	GetChats(*FlotgGetChatsRequest, FlotgService_GetChatsServer) error
	SetMonitoring(context.Context, *FlotgMonitor) (*FlotgMonitor, error)
	GetMessages(*FlotgGetMessagesRequest, FlotgService_GetMessagesServer) error
	mustEmbedUnimplementedFlotgServiceServer()
}

// UnimplementedFlotgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlotgServiceServer struct {
}

func (UnimplementedFlotgServiceServer) GetChats(*FlotgGetChatsRequest, FlotgService_GetChatsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedFlotgServiceServer) SetMonitoring(context.Context, *FlotgMonitor) (*FlotgMonitor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMonitoring not implemented")
}
func (UnimplementedFlotgServiceServer) GetMessages(*FlotgGetMessagesRequest, FlotgService_GetMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedFlotgServiceServer) mustEmbedUnimplementedFlotgServiceServer() {}

// UnsafeFlotgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlotgServiceServer will
// result in compilation errors.
type UnsafeFlotgServiceServer interface {
	mustEmbedUnimplementedFlotgServiceServer()
}

func RegisterFlotgServiceServer(s grpc.ServiceRegistrar, srv FlotgServiceServer) {
	s.RegisterService(&FlotgService_ServiceDesc, srv)
}

func _FlotgService_GetChats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FlotgGetChatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlotgServiceServer).GetChats(m, &flotgServiceGetChatsServer{stream})
}

type FlotgService_GetChatsServer interface {
	Send(*FlotgMonitor) error
	grpc.ServerStream
}

type flotgServiceGetChatsServer struct {
	grpc.ServerStream
}

func (x *flotgServiceGetChatsServer) Send(m *FlotgMonitor) error {
	return x.ServerStream.SendMsg(m)
}

func _FlotgService_SetMonitoring_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlotgMonitor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlotgServiceServer).SetMonitoring(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FlotgService/SetMonitoring",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlotgServiceServer).SetMonitoring(ctx, req.(*FlotgMonitor))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlotgService_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FlotgGetMessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FlotgServiceServer).GetMessages(m, &flotgServiceGetMessagesServer{stream})
}

type FlotgService_GetMessagesServer interface {
	Send(*FLO_MESSAGE) error
	grpc.ServerStream
}

type flotgServiceGetMessagesServer struct {
	grpc.ServerStream
}

func (x *flotgServiceGetMessagesServer) Send(m *FLO_MESSAGE) error {
	return x.ServerStream.SendMsg(m)
}

// FlotgService_ServiceDesc is the grpc.ServiceDesc for FlotgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlotgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FlotgService",
	HandlerType: (*FlotgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMonitoring",
			Handler:    _FlotgService_SetMonitoring_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetChats",
			Handler:       _FlotgService_GetChats_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMessages",
			Handler:       _FlotgService_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "flogram.proto",
}

// FloRssServiceClient is the client API for FloRssService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FloRssServiceClient interface {
	GetFeeds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (FloRssService_GetFeedsClient, error)
	CreateFeed(ctx context.Context, in *FloRssCreateRequest, opts ...grpc.CallOption) (*FloRssFeed, error)
	DeleteFeed(ctx context.Context, in *FloRssFeed, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMessages(ctx context.Context, in *FloRssFeed, opts ...grpc.CallOption) (FloRssService_GetMessagesClient, error)
}

type floRssServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFloRssServiceClient(cc grpc.ClientConnInterface) FloRssServiceClient {
	return &floRssServiceClient{cc}
}

func (c *floRssServiceClient) GetFeeds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (FloRssService_GetFeedsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FloRssService_ServiceDesc.Streams[0], "/FloRssService/GetFeeds", opts...)
	if err != nil {
		return nil, err
	}
	x := &floRssServiceGetFeedsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FloRssService_GetFeedsClient interface {
	Recv() (*FloRssFeed, error)
	grpc.ClientStream
}

type floRssServiceGetFeedsClient struct {
	grpc.ClientStream
}

func (x *floRssServiceGetFeedsClient) Recv() (*FloRssFeed, error) {
	m := new(FloRssFeed)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *floRssServiceClient) CreateFeed(ctx context.Context, in *FloRssCreateRequest, opts ...grpc.CallOption) (*FloRssFeed, error) {
	out := new(FloRssFeed)
	err := c.cc.Invoke(ctx, "/FloRssService/CreateFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floRssServiceClient) DeleteFeed(ctx context.Context, in *FloRssFeed, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/FloRssService/DeleteFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floRssServiceClient) GetMessages(ctx context.Context, in *FloRssFeed, opts ...grpc.CallOption) (FloRssService_GetMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &FloRssService_ServiceDesc.Streams[1], "/FloRssService/GetMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &floRssServiceGetMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FloRssService_GetMessagesClient interface {
	Recv() (*FLO_MESSAGE, error)
	grpc.ClientStream
}

type floRssServiceGetMessagesClient struct {
	grpc.ClientStream
}

func (x *floRssServiceGetMessagesClient) Recv() (*FLO_MESSAGE, error) {
	m := new(FLO_MESSAGE)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FloRssServiceServer is the server API for FloRssService service.
// All implementations must embed UnimplementedFloRssServiceServer
// for forward compatibility
type FloRssServiceServer interface {
	GetFeeds(*emptypb.Empty, FloRssService_GetFeedsServer) error
	CreateFeed(context.Context, *FloRssCreateRequest) (*FloRssFeed, error)
	DeleteFeed(context.Context, *FloRssFeed) (*emptypb.Empty, error)
	GetMessages(*FloRssFeed, FloRssService_GetMessagesServer) error
	mustEmbedUnimplementedFloRssServiceServer()
}

// UnimplementedFloRssServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFloRssServiceServer struct {
}

func (UnimplementedFloRssServiceServer) GetFeeds(*emptypb.Empty, FloRssService_GetFeedsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFeeds not implemented")
}
func (UnimplementedFloRssServiceServer) CreateFeed(context.Context, *FloRssCreateRequest) (*FloRssFeed, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeed not implemented")
}
func (UnimplementedFloRssServiceServer) DeleteFeed(context.Context, *FloRssFeed) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeed not implemented")
}
func (UnimplementedFloRssServiceServer) GetMessages(*FloRssFeed, FloRssService_GetMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedFloRssServiceServer) mustEmbedUnimplementedFloRssServiceServer() {}

// UnsafeFloRssServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FloRssServiceServer will
// result in compilation errors.
type UnsafeFloRssServiceServer interface {
	mustEmbedUnimplementedFloRssServiceServer()
}

func RegisterFloRssServiceServer(s grpc.ServiceRegistrar, srv FloRssServiceServer) {
	s.RegisterService(&FloRssService_ServiceDesc, srv)
}

func _FloRssService_GetFeeds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FloRssServiceServer).GetFeeds(m, &floRssServiceGetFeedsServer{stream})
}

type FloRssService_GetFeedsServer interface {
	Send(*FloRssFeed) error
	grpc.ServerStream
}

type floRssServiceGetFeedsServer struct {
	grpc.ServerStream
}

func (x *floRssServiceGetFeedsServer) Send(m *FloRssFeed) error {
	return x.ServerStream.SendMsg(m)
}

func _FloRssService_CreateFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FloRssCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloRssServiceServer).CreateFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FloRssService/CreateFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloRssServiceServer).CreateFeed(ctx, req.(*FloRssCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloRssService_DeleteFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FloRssFeed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloRssServiceServer).DeleteFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FloRssService/DeleteFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloRssServiceServer).DeleteFeed(ctx, req.(*FloRssFeed))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloRssService_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FloRssFeed)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FloRssServiceServer).GetMessages(m, &floRssServiceGetMessagesServer{stream})
}

type FloRssService_GetMessagesServer interface {
	Send(*FLO_MESSAGE) error
	grpc.ServerStream
}

type floRssServiceGetMessagesServer struct {
	grpc.ServerStream
}

func (x *floRssServiceGetMessagesServer) Send(m *FLO_MESSAGE) error {
	return x.ServerStream.SendMsg(m)
}

// FloRssService_ServiceDesc is the grpc.ServiceDesc for FloRssService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FloRssService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FloRssService",
	HandlerType: (*FloRssServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeed",
			Handler:    _FloRssService_CreateFeed_Handler,
		},
		{
			MethodName: "DeleteFeed",
			Handler:    _FloRssService_DeleteFeed_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFeeds",
			Handler:       _FloRssService_GetFeeds_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMessages",
			Handler:       _FloRssService_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "flogram.proto",
}

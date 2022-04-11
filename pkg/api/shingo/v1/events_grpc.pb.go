// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: events/v1/events.proto

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

// EventsServiceClient is the client API for EventsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventsServiceClient interface {
	// Emits an event for a topic
	// FIXME: This should be handled on a single stream, will implement once grpc web handled bi-di streaming
	Emit(ctx context.Context, in *EmitRequest, opts ...grpc.CallOption) (*EmitResponse, error)
	// Subscribes to events on a topic
	// FIXME: deprecate once bi-di streaming is available for gRPC web
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (EventsService_SubscribeClient, error)
	// Unsubscribe from the topic
	// FIXME: This should be handled on a single stream, will implement once grpc web handled bi-di streaming
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
}

type eventsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsServiceClient(cc grpc.ClientConnInterface) EventsServiceClient {
	return &eventsServiceClient{cc}
}

func (c *eventsServiceClient) Emit(ctx context.Context, in *EmitRequest, opts ...grpc.CallOption) (*EmitResponse, error) {
	out := new(EmitResponse)
	err := c.cc.Invoke(ctx, "/shingo.events.v1.EventsService/Emit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (EventsService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventsService_ServiceDesc.Streams[0], "/shingo.events.v1.EventsService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventsService_SubscribeClient interface {
	Recv() (*SubscriptionEvent, error)
	grpc.ClientStream
}

type eventsServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *eventsServiceSubscribeClient) Recv() (*SubscriptionEvent, error) {
	m := new(SubscriptionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventsServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, "/shingo.events.v1.EventsService/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServiceServer is the server API for EventsService service.
// All implementations must embed UnimplementedEventsServiceServer
// for forward compatibility
type EventsServiceServer interface {
	// Emits an event for a topic
	// FIXME: This should be handled on a single stream, will implement once grpc web handled bi-di streaming
	Emit(context.Context, *EmitRequest) (*EmitResponse, error)
	// Subscribes to events on a topic
	// FIXME: deprecate once bi-di streaming is available for gRPC web
	Subscribe(*SubscribeRequest, EventsService_SubscribeServer) error
	// Unsubscribe from the topic
	// FIXME: This should be handled on a single stream, will implement once grpc web handled bi-di streaming
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
	mustEmbedUnimplementedEventsServiceServer()
}

// UnimplementedEventsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventsServiceServer struct {
}

func (UnimplementedEventsServiceServer) Emit(context.Context, *EmitRequest) (*EmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Emit not implemented")
}
func (UnimplementedEventsServiceServer) Subscribe(*SubscribeRequest, EventsService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedEventsServiceServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedEventsServiceServer) mustEmbedUnimplementedEventsServiceServer() {}

// UnsafeEventsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventsServiceServer will
// result in compilation errors.
type UnsafeEventsServiceServer interface {
	mustEmbedUnimplementedEventsServiceServer()
}

func RegisterEventsServiceServer(s grpc.ServiceRegistrar, srv EventsServiceServer) {
	s.RegisterService(&EventsService_ServiceDesc, srv)
}

func _EventsService_Emit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).Emit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shingo.events.v1.EventsService/Emit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).Emit(ctx, req.(*EmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsServiceServer).Subscribe(m, &eventsServiceSubscribeServer{stream})
}

type EventsService_SubscribeServer interface {
	Send(*SubscriptionEvent) error
	grpc.ServerStream
}

type eventsServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *eventsServiceSubscribeServer) Send(m *SubscriptionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _EventsService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shingo.events.v1.EventsService/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventsService_ServiceDesc is the grpc.ServiceDesc for EventsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shingo.events.v1.EventsService",
	HandlerType: (*EventsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Emit",
			Handler:    _EventsService_Emit_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _EventsService_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _EventsService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "events/v1/events.proto",
}

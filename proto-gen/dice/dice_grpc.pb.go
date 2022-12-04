// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protoc/dice.proto

package dice

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

// TutorialClient is the client API for Tutorial service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TutorialClient interface {
	// Ping rpc
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
	// RollDice method
	RollDice(ctx context.Context, in *RollDiceRequest, opts ...grpc.CallOption) (*RollDiceResponse, error)
}

type tutorialClient struct {
	cc grpc.ClientConnInterface
}

func NewTutorialClient(cc grpc.ClientConnInterface) TutorialClient {
	return &tutorialClient{cc}
}

func (c *tutorialClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/Tutorial/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tutorialClient) RollDice(ctx context.Context, in *RollDiceRequest, opts ...grpc.CallOption) (*RollDiceResponse, error) {
	out := new(RollDiceResponse)
	err := c.cc.Invoke(ctx, "/Tutorial/RollDice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TutorialServer is the server API for Tutorial service.
// All implementations must embed UnimplementedTutorialServer
// for forward compatibility
type TutorialServer interface {
	// Ping rpc
	Ping(context.Context, *Empty) (*Pong, error)
	// RollDice method
	RollDice(context.Context, *RollDiceRequest) (*RollDiceResponse, error)
	mustEmbedUnimplementedTutorialServer()
}

// UnimplementedTutorialServer must be embedded to have forward compatible implementations.
type UnimplementedTutorialServer struct {
}

func (UnimplementedTutorialServer) Ping(context.Context, *Empty) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTutorialServer) RollDice(context.Context, *RollDiceRequest) (*RollDiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollDice not implemented")
}
func (UnimplementedTutorialServer) mustEmbedUnimplementedTutorialServer() {}

// UnsafeTutorialServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TutorialServer will
// result in compilation errors.
type UnsafeTutorialServer interface {
	mustEmbedUnimplementedTutorialServer()
}

func RegisterTutorialServer(s grpc.ServiceRegistrar, srv TutorialServer) {
	s.RegisterService(&Tutorial_ServiceDesc, srv)
}

func _Tutorial_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tutorial/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tutorial_RollDice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollDiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TutorialServer).RollDice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Tutorial/RollDice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TutorialServer).RollDice(ctx, req.(*RollDiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tutorial_ServiceDesc is the grpc.ServiceDesc for Tutorial service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tutorial_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Tutorial",
	HandlerType: (*TutorialServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Tutorial_Ping_Handler,
		},
		{
			MethodName: "RollDice",
			Handler:    _Tutorial_RollDice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoc/dice.proto",
}

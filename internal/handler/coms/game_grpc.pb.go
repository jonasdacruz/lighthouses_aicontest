// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: game.proto

package coms

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

const (
	GameService_Join_FullMethodName         = "/GameService/Join"
	GameService_InitialState_FullMethodName = "/GameService/InitialState"
	GameService_Turn_FullMethodName         = "/GameService/Turn"
)

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	Join(ctx context.Context, in *NewPlayer, opts ...grpc.CallOption) (*PlayerID, error)
	InitialState(ctx context.Context, in *PlayerID, opts ...grpc.CallOption) (*NewPlayerInitialState, error)
	Turn(ctx context.Context, in *NewTurn, opts ...grpc.CallOption) (*NewAction, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) Join(ctx context.Context, in *NewPlayer, opts ...grpc.CallOption) (*PlayerID, error) {
	out := new(PlayerID)
	err := c.cc.Invoke(ctx, GameService_Join_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) InitialState(ctx context.Context, in *PlayerID, opts ...grpc.CallOption) (*NewPlayerInitialState, error) {
	out := new(NewPlayerInitialState)
	err := c.cc.Invoke(ctx, GameService_InitialState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) Turn(ctx context.Context, in *NewTurn, opts ...grpc.CallOption) (*NewAction, error) {
	out := new(NewAction)
	err := c.cc.Invoke(ctx, GameService_Turn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations should embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	Join(context.Context, *NewPlayer) (*PlayerID, error)
	InitialState(context.Context, *PlayerID) (*NewPlayerInitialState, error)
	Turn(context.Context, *NewTurn) (*NewAction, error)
}

// UnimplementedGameServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) Join(context.Context, *NewPlayer) (*PlayerID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedGameServiceServer) InitialState(context.Context, *PlayerID) (*NewPlayerInitialState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitialState not implemented")
}
func (UnimplementedGameServiceServer) Turn(context.Context, *NewTurn) (*NewAction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Turn not implemented")
}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPlayer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_Join_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).Join(ctx, req.(*NewPlayer))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_InitialState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).InitialState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_InitialState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).InitialState(ctx, req.(*PlayerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_Turn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTurn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).Turn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GameService_Turn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).Turn(ctx, req.(*NewTurn))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _GameService_Join_Handler,
		},
		{
			MethodName: "InitialState",
			Handler:    _GameService_InitialState_Handler,
		},
		{
			MethodName: "Turn",
			Handler:    _GameService_Turn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}

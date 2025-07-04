// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: character/service/character_progress.proto

package character_progress

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CharacterProgressService_GainExp_FullMethodName      = "/character_progress.v1.CharacterProgressService/GainExp"
	CharacterProgressService_RemoveExp_FullMethodName    = "/character_progress.v1.CharacterProgressService/RemoveExp"
	CharacterProgressService_LvlUp_FullMethodName        = "/character_progress.v1.CharacterProgressService/LvlUp"
	CharacterProgressService_LvlDown_FullMethodName      = "/character_progress.v1.CharacterProgressService/LvlDown"
	CharacterProgressService_CanLvlUp_FullMethodName     = "/character_progress.v1.CharacterProgressService/CanLvlUp"
	CharacterProgressService_CanLvlDown_FullMethodName   = "/character_progress.v1.CharacterProgressService/CanLvlDown"
	CharacterProgressService_CurrentLvl_FullMethodName   = "/character_progress.v1.CharacterProgressService/CurrentLvl"
	CharacterProgressService_ExpToNextLvl_FullMethodName = "/character_progress.v1.CharacterProgressService/ExpToNextLvl"
	CharacterProgressService_SetLvl_FullMethodName       = "/character_progress.v1.CharacterProgressService/SetLvl"
	CharacterProgressService_GetLvlState_FullMethodName  = "/character_progress.v1.CharacterProgressService/GetLvlState"
)

// CharacterProgressServiceClient is the client API for CharacterProgressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharacterProgressServiceClient interface {
	// GainExp adds experience points to character
	GainExp(ctx context.Context, in *ModifyExpRequest, opts ...grpc.CallOption) (*ExpResponse, error)
	// RemoveExp subtracts experience points from character
	RemoveExp(ctx context.Context, in *ModifyExpRequest, opts ...grpc.CallOption) (*ExpResponse, error)
	// LevelUp increases character level by 1
	LvlUp(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*ExpResponse, error)
	// LevelDown decreases character level by 1
	LvlDown(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*ExpResponse, error)
	// CanLvlUp checks if character can level up
	CanLvlUp(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*CanLvlResponse, error)
	// CanLvlDown checks if character can level down
	CanLvlDown(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*CanLvlResponse, error)
	// CurrentLvl returns character's current level
	CurrentLvl(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*CurrentLvlResponse, error)
	// ExpToNextLvl returns experience needed for next level
	ExpToNextLvl(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*ExpToNextLvlResponse, error)
	// SetLevel directly sets character level
	SetLvl(ctx context.Context, in *SetLevelRequest, opts ...grpc.CallOption) (*ExpResponse, error)
	// Obtain the current lvl state
	GetLvlState(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*ExpResponse, error)
}

type characterProgressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCharacterProgressServiceClient(cc grpc.ClientConnInterface) CharacterProgressServiceClient {
	return &characterProgressServiceClient{cc}
}

func (c *characterProgressServiceClient) GainExp(ctx context.Context, in *ModifyExpRequest, opts ...grpc.CallOption) (*ExpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpResponse)
	err := c.cc.Invoke(ctx, CharacterProgressService_GainExp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterProgressServiceClient) RemoveExp(ctx context.Context, in *ModifyExpRequest, opts ...grpc.CallOption) (*ExpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpResponse)
	err := c.cc.Invoke(ctx, CharacterProgressService_RemoveExp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterProgressServiceClient) LvlUp(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*ExpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpResponse)
	err := c.cc.Invoke(ctx, CharacterProgressService_LvlUp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterProgressServiceClient) LvlDown(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*ExpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpResponse)
	err := c.cc.Invoke(ctx, CharacterProgressService_LvlDown_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterProgressServiceClient) CanLvlUp(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*CanLvlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CanLvlResponse)
	err := c.cc.Invoke(ctx, CharacterProgressService_CanLvlUp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterProgressServiceClient) CanLvlDown(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*CanLvlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CanLvlResponse)
	err := c.cc.Invoke(ctx, CharacterProgressService_CanLvlDown_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterProgressServiceClient) CurrentLvl(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*CurrentLvlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CurrentLvlResponse)
	err := c.cc.Invoke(ctx, CharacterProgressService_CurrentLvl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterProgressServiceClient) ExpToNextLvl(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*ExpToNextLvlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpToNextLvlResponse)
	err := c.cc.Invoke(ctx, CharacterProgressService_ExpToNextLvl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterProgressServiceClient) SetLvl(ctx context.Context, in *SetLevelRequest, opts ...grpc.CallOption) (*ExpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpResponse)
	err := c.cc.Invoke(ctx, CharacterProgressService_SetLvl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterProgressServiceClient) GetLvlState(ctx context.Context, in *CharacterID, opts ...grpc.CallOption) (*ExpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpResponse)
	err := c.cc.Invoke(ctx, CharacterProgressService_GetLvlState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharacterProgressServiceServer is the server API for CharacterProgressService service.
// All implementations must embed UnimplementedCharacterProgressServiceServer
// for forward compatibility.
type CharacterProgressServiceServer interface {
	// GainExp adds experience points to character
	GainExp(context.Context, *ModifyExpRequest) (*ExpResponse, error)
	// RemoveExp subtracts experience points from character
	RemoveExp(context.Context, *ModifyExpRequest) (*ExpResponse, error)
	// LevelUp increases character level by 1
	LvlUp(context.Context, *CharacterID) (*ExpResponse, error)
	// LevelDown decreases character level by 1
	LvlDown(context.Context, *CharacterID) (*ExpResponse, error)
	// CanLvlUp checks if character can level up
	CanLvlUp(context.Context, *CharacterID) (*CanLvlResponse, error)
	// CanLvlDown checks if character can level down
	CanLvlDown(context.Context, *CharacterID) (*CanLvlResponse, error)
	// CurrentLvl returns character's current level
	CurrentLvl(context.Context, *CharacterID) (*CurrentLvlResponse, error)
	// ExpToNextLvl returns experience needed for next level
	ExpToNextLvl(context.Context, *CharacterID) (*ExpToNextLvlResponse, error)
	// SetLevel directly sets character level
	SetLvl(context.Context, *SetLevelRequest) (*ExpResponse, error)
	// Obtain the current lvl state
	GetLvlState(context.Context, *CharacterID) (*ExpResponse, error)
	mustEmbedUnimplementedCharacterProgressServiceServer()
}

// UnimplementedCharacterProgressServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCharacterProgressServiceServer struct{}

func (UnimplementedCharacterProgressServiceServer) GainExp(context.Context, *ModifyExpRequest) (*ExpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GainExp not implemented")
}
func (UnimplementedCharacterProgressServiceServer) RemoveExp(context.Context, *ModifyExpRequest) (*ExpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveExp not implemented")
}
func (UnimplementedCharacterProgressServiceServer) LvlUp(context.Context, *CharacterID) (*ExpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LvlUp not implemented")
}
func (UnimplementedCharacterProgressServiceServer) LvlDown(context.Context, *CharacterID) (*ExpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LvlDown not implemented")
}
func (UnimplementedCharacterProgressServiceServer) CanLvlUp(context.Context, *CharacterID) (*CanLvlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanLvlUp not implemented")
}
func (UnimplementedCharacterProgressServiceServer) CanLvlDown(context.Context, *CharacterID) (*CanLvlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanLvlDown not implemented")
}
func (UnimplementedCharacterProgressServiceServer) CurrentLvl(context.Context, *CharacterID) (*CurrentLvlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentLvl not implemented")
}
func (UnimplementedCharacterProgressServiceServer) ExpToNextLvl(context.Context, *CharacterID) (*ExpToNextLvlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpToNextLvl not implemented")
}
func (UnimplementedCharacterProgressServiceServer) SetLvl(context.Context, *SetLevelRequest) (*ExpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLvl not implemented")
}
func (UnimplementedCharacterProgressServiceServer) GetLvlState(context.Context, *CharacterID) (*ExpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLvlState not implemented")
}
func (UnimplementedCharacterProgressServiceServer) mustEmbedUnimplementedCharacterProgressServiceServer() {
}
func (UnimplementedCharacterProgressServiceServer) testEmbeddedByValue() {}

// UnsafeCharacterProgressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharacterProgressServiceServer will
// result in compilation errors.
type UnsafeCharacterProgressServiceServer interface {
	mustEmbedUnimplementedCharacterProgressServiceServer()
}

func RegisterCharacterProgressServiceServer(s grpc.ServiceRegistrar, srv CharacterProgressServiceServer) {
	// If the following call pancis, it indicates UnimplementedCharacterProgressServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CharacterProgressService_ServiceDesc, srv)
}

func _CharacterProgressService_GainExp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyExpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterProgressServiceServer).GainExp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterProgressService_GainExp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterProgressServiceServer).GainExp(ctx, req.(*ModifyExpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterProgressService_RemoveExp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyExpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterProgressServiceServer).RemoveExp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterProgressService_RemoveExp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterProgressServiceServer).RemoveExp(ctx, req.(*ModifyExpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterProgressService_LvlUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CharacterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterProgressServiceServer).LvlUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterProgressService_LvlUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterProgressServiceServer).LvlUp(ctx, req.(*CharacterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterProgressService_LvlDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CharacterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterProgressServiceServer).LvlDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterProgressService_LvlDown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterProgressServiceServer).LvlDown(ctx, req.(*CharacterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterProgressService_CanLvlUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CharacterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterProgressServiceServer).CanLvlUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterProgressService_CanLvlUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterProgressServiceServer).CanLvlUp(ctx, req.(*CharacterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterProgressService_CanLvlDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CharacterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterProgressServiceServer).CanLvlDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterProgressService_CanLvlDown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterProgressServiceServer).CanLvlDown(ctx, req.(*CharacterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterProgressService_CurrentLvl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CharacterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterProgressServiceServer).CurrentLvl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterProgressService_CurrentLvl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterProgressServiceServer).CurrentLvl(ctx, req.(*CharacterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterProgressService_ExpToNextLvl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CharacterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterProgressServiceServer).ExpToNextLvl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterProgressService_ExpToNextLvl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterProgressServiceServer).ExpToNextLvl(ctx, req.(*CharacterID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterProgressService_SetLvl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterProgressServiceServer).SetLvl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterProgressService_SetLvl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterProgressServiceServer).SetLvl(ctx, req.(*SetLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterProgressService_GetLvlState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CharacterID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterProgressServiceServer).GetLvlState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterProgressService_GetLvlState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterProgressServiceServer).GetLvlState(ctx, req.(*CharacterID))
	}
	return interceptor(ctx, in, info, handler)
}

// CharacterProgressService_ServiceDesc is the grpc.ServiceDesc for CharacterProgressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CharacterProgressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "character_progress.v1.CharacterProgressService",
	HandlerType: (*CharacterProgressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GainExp",
			Handler:    _CharacterProgressService_GainExp_Handler,
		},
		{
			MethodName: "RemoveExp",
			Handler:    _CharacterProgressService_RemoveExp_Handler,
		},
		{
			MethodName: "LvlUp",
			Handler:    _CharacterProgressService_LvlUp_Handler,
		},
		{
			MethodName: "LvlDown",
			Handler:    _CharacterProgressService_LvlDown_Handler,
		},
		{
			MethodName: "CanLvlUp",
			Handler:    _CharacterProgressService_CanLvlUp_Handler,
		},
		{
			MethodName: "CanLvlDown",
			Handler:    _CharacterProgressService_CanLvlDown_Handler,
		},
		{
			MethodName: "CurrentLvl",
			Handler:    _CharacterProgressService_CurrentLvl_Handler,
		},
		{
			MethodName: "ExpToNextLvl",
			Handler:    _CharacterProgressService_ExpToNextLvl_Handler,
		},
		{
			MethodName: "SetLvl",
			Handler:    _CharacterProgressService_SetLvl_Handler,
		},
		{
			MethodName: "GetLvlState",
			Handler:    _CharacterProgressService_GetLvlState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "character/service/character_progress.proto",
}

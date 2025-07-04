// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: sso/user_info.proto

package ssov1

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
	UserInfo_GetUserById_FullMethodName    = "/userinfo.UserInfo/GetUserById"
	UserInfo_GetUserByEmail_FullMethodName = "/userinfo.UserInfo/GetUserByEmail"
	UserInfo_UpdateUser_FullMethodName     = "/userinfo.UserInfo/UpdateUser"
	UserInfo_DeleteUser_FullMethodName     = "/userinfo.UserInfo/DeleteUser"
)

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoClient interface {
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*GetUserByEmailResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoClient(cc grpc.ClientConnInterface) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, UserInfo_GetUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) GetUserByEmail(ctx context.Context, in *GetUserByEmailRequest, opts ...grpc.CallOption) (*GetUserByEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserByEmailResponse)
	err := c.cc.Invoke(ctx, UserInfo_GetUserByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, UserInfo_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserInfo_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServer is the server API for UserInfo service.
// All implementations must embed UnimplementedUserInfoServer
// for forward compatibility.
type UserInfoServer interface {
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	GetUserByEmail(context.Context, *GetUserByEmailRequest) (*GetUserByEmailResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedUserInfoServer()
}

// UnimplementedUserInfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserInfoServer struct{}

func (UnimplementedUserInfoServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserInfoServer) GetUserByEmail(context.Context, *GetUserByEmailRequest) (*GetUserByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUserInfoServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserInfoServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserInfoServer) mustEmbedUnimplementedUserInfoServer() {}
func (UnimplementedUserInfoServer) testEmbeddedByValue()                  {}

// UnsafeUserInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServer will
// result in compilation errors.
type UnsafeUserInfoServer interface {
	mustEmbedUnimplementedUserInfoServer()
}

func RegisterUserInfoServer(s grpc.ServiceRegistrar, srv UserInfoServer) {
	// If the following call pancis, it indicates UnimplementedUserInfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserInfo_ServiceDesc, srv)
}

func _UserInfo_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfo_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfo_GetUserByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).GetUserByEmail(ctx, req.(*GetUserByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfo_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfo_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfo_ServiceDesc is the grpc.ServiceDesc for UserInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userinfo.UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserInfo_GetUserById_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UserInfo_GetUserByEmail_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserInfo_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserInfo_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso/user_info.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.1
// source: helloworld/v1/user.proto

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

const (
	UserService_UserRegister_FullMethodName = "/user.v1.UserService/UserRegister"
	UserService_UserList_FullMethodName     = "/user.v1.UserService/UserList"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// user register
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	UserList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserListResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, UserService_UserRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, UserService_UserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// user register
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	UserList(context.Context, *Empty) (*UserListResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserServiceServer) UserList(context.Context, *Empty) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _UserService_UserRegister_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _UserService_UserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/user.proto",
}
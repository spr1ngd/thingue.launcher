// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: apis/v1/server_instance_api.proto

package apisv1

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

const (
	ServerInstanceService_AddInstance_FullMethodName        = "/apis.v1.ServerInstanceService/AddInstance"
	ServerInstanceService_DeleteInstance_FullMethodName     = "/apis.v1.ServerInstanceService/DeleteInstance"
	ServerInstanceService_UpdateConfig_FullMethodName       = "/apis.v1.ServerInstanceService/UpdateConfig"
	ServerInstanceService_GetStreamerId_FullMethodName      = "/apis.v1.ServerInstanceService/GetStreamerId"
	ServerInstanceService_UpdateProcessState_FullMethodName = "/apis.v1.ServerInstanceService/UpdateProcessState"
	ServerInstanceService_UpdateRestarting_FullMethodName   = "/apis.v1.ServerInstanceService/UpdateRestarting"
	ServerInstanceService_ClearPakState_FullMethodName      = "/apis.v1.ServerInstanceService/ClearPakState"
)

// ServerInstanceServiceClient is the client API for ServerInstanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerInstanceServiceClient interface {
	AddInstance(ctx context.Context, in *AddInstanceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetStreamerId(ctx context.Context, in *GetStreamerIdRequest, opts ...grpc.CallOption) (*GetStreamerIdResponse, error)
	UpdateProcessState(ctx context.Context, in *UpdateProcessStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateRestarting(ctx context.Context, in *UpdateRestartingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ClearPakState(ctx context.Context, in *ClearPakStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type serverInstanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerInstanceServiceClient(cc grpc.ClientConnInterface) ServerInstanceServiceClient {
	return &serverInstanceServiceClient{cc}
}

func (c *serverInstanceServiceClient) AddInstance(ctx context.Context, in *AddInstanceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ServerInstanceService_AddInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverInstanceServiceClient) DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ServerInstanceService_DeleteInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverInstanceServiceClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ServerInstanceService_UpdateConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverInstanceServiceClient) GetStreamerId(ctx context.Context, in *GetStreamerIdRequest, opts ...grpc.CallOption) (*GetStreamerIdResponse, error) {
	out := new(GetStreamerIdResponse)
	err := c.cc.Invoke(ctx, ServerInstanceService_GetStreamerId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverInstanceServiceClient) UpdateProcessState(ctx context.Context, in *UpdateProcessStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ServerInstanceService_UpdateProcessState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverInstanceServiceClient) UpdateRestarting(ctx context.Context, in *UpdateRestartingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ServerInstanceService_UpdateRestarting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverInstanceServiceClient) ClearPakState(ctx context.Context, in *ClearPakStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ServerInstanceService_ClearPakState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerInstanceServiceServer is the server API for ServerInstanceService service.
// All implementations must embed UnimplementedServerInstanceServiceServer
// for forward compatibility
type ServerInstanceServiceServer interface {
	AddInstance(context.Context, *AddInstanceRequest) (*emptypb.Empty, error)
	DeleteInstance(context.Context, *DeleteInstanceRequest) (*emptypb.Empty, error)
	UpdateConfig(context.Context, *UpdateConfigRequest) (*emptypb.Empty, error)
	GetStreamerId(context.Context, *GetStreamerIdRequest) (*GetStreamerIdResponse, error)
	UpdateProcessState(context.Context, *UpdateProcessStateRequest) (*emptypb.Empty, error)
	UpdateRestarting(context.Context, *UpdateRestartingRequest) (*emptypb.Empty, error)
	ClearPakState(context.Context, *ClearPakStateRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedServerInstanceServiceServer()
}

// UnimplementedServerInstanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerInstanceServiceServer struct {
}

func (UnimplementedServerInstanceServiceServer) AddInstance(context.Context, *AddInstanceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInstance not implemented")
}
func (UnimplementedServerInstanceServiceServer) DeleteInstance(context.Context, *DeleteInstanceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstance not implemented")
}
func (UnimplementedServerInstanceServiceServer) UpdateConfig(context.Context, *UpdateConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedServerInstanceServiceServer) GetStreamerId(context.Context, *GetStreamerIdRequest) (*GetStreamerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamerId not implemented")
}
func (UnimplementedServerInstanceServiceServer) UpdateProcessState(context.Context, *UpdateProcessStateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProcessState not implemented")
}
func (UnimplementedServerInstanceServiceServer) UpdateRestarting(context.Context, *UpdateRestartingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestarting not implemented")
}
func (UnimplementedServerInstanceServiceServer) ClearPakState(context.Context, *ClearPakStateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearPakState not implemented")
}
func (UnimplementedServerInstanceServiceServer) mustEmbedUnimplementedServerInstanceServiceServer() {}

// UnsafeServerInstanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerInstanceServiceServer will
// result in compilation errors.
type UnsafeServerInstanceServiceServer interface {
	mustEmbedUnimplementedServerInstanceServiceServer()
}

func RegisterServerInstanceServiceServer(s grpc.ServiceRegistrar, srv ServerInstanceServiceServer) {
	s.RegisterService(&ServerInstanceService_ServiceDesc, srv)
}

func _ServerInstanceService_AddInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInstanceServiceServer).AddInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerInstanceService_AddInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInstanceServiceServer).AddInstance(ctx, req.(*AddInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerInstanceService_DeleteInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInstanceServiceServer).DeleteInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerInstanceService_DeleteInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInstanceServiceServer).DeleteInstance(ctx, req.(*DeleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerInstanceService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInstanceServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerInstanceService_UpdateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInstanceServiceServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerInstanceService_GetStreamerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInstanceServiceServer).GetStreamerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerInstanceService_GetStreamerId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInstanceServiceServer).GetStreamerId(ctx, req.(*GetStreamerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerInstanceService_UpdateProcessState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProcessStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInstanceServiceServer).UpdateProcessState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerInstanceService_UpdateProcessState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInstanceServiceServer).UpdateProcessState(ctx, req.(*UpdateProcessStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerInstanceService_UpdateRestarting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRestartingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInstanceServiceServer).UpdateRestarting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerInstanceService_UpdateRestarting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInstanceServiceServer).UpdateRestarting(ctx, req.(*UpdateRestartingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerInstanceService_ClearPakState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearPakStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInstanceServiceServer).ClearPakState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerInstanceService_ClearPakState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInstanceServiceServer).ClearPakState(ctx, req.(*ClearPakStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerInstanceService_ServiceDesc is the grpc.ServiceDesc for ServerInstanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerInstanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apis.v1.ServerInstanceService",
	HandlerType: (*ServerInstanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddInstance",
			Handler:    _ServerInstanceService_AddInstance_Handler,
		},
		{
			MethodName: "DeleteInstance",
			Handler:    _ServerInstanceService_DeleteInstance_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _ServerInstanceService_UpdateConfig_Handler,
		},
		{
			MethodName: "GetStreamerId",
			Handler:    _ServerInstanceService_GetStreamerId_Handler,
		},
		{
			MethodName: "UpdateProcessState",
			Handler:    _ServerInstanceService_UpdateProcessState_Handler,
		},
		{
			MethodName: "UpdateRestarting",
			Handler:    _ServerInstanceService_UpdateRestarting_Handler,
		},
		{
			MethodName: "ClearPakState",
			Handler:    _ServerInstanceService_ClearPakState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/v1/server_instance_api.proto",
}

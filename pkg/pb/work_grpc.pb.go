// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// WorkManagerClient is the client API for WorkManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkManagerClient interface {
	CreateWork(ctx context.Context, in *CreateWorkRequest, opts ...grpc.CallOption) (*Work, error)
	GetWork(ctx context.Context, in *GetWorkRequest, opts ...grpc.CallOption) (*Work, error)
	UpdateWork(ctx context.Context, in *UpdateWorkRequest, opts ...grpc.CallOption) (*Work, error)
	DeleteWork(ctx context.Context, in *DeleteWorkRequest, opts ...grpc.CallOption) (*Work, error)
	GetWorks(ctx context.Context, in *GetWorksRequest, opts ...grpc.CallOption) (*GetWorksReply, error)
}

type workManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkManagerClient(cc grpc.ClientConnInterface) WorkManagerClient {
	return &workManagerClient{cc}
}

func (c *workManagerClient) CreateWork(ctx context.Context, in *CreateWorkRequest, opts ...grpc.CallOption) (*Work, error) {
	out := new(Work)
	err := c.cc.Invoke(ctx, "/pb.WorkManager/CreateWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workManagerClient) GetWork(ctx context.Context, in *GetWorkRequest, opts ...grpc.CallOption) (*Work, error) {
	out := new(Work)
	err := c.cc.Invoke(ctx, "/pb.WorkManager/GetWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workManagerClient) UpdateWork(ctx context.Context, in *UpdateWorkRequest, opts ...grpc.CallOption) (*Work, error) {
	out := new(Work)
	err := c.cc.Invoke(ctx, "/pb.WorkManager/UpdateWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workManagerClient) DeleteWork(ctx context.Context, in *DeleteWorkRequest, opts ...grpc.CallOption) (*Work, error) {
	out := new(Work)
	err := c.cc.Invoke(ctx, "/pb.WorkManager/DeleteWork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workManagerClient) GetWorks(ctx context.Context, in *GetWorksRequest, opts ...grpc.CallOption) (*GetWorksReply, error) {
	out := new(GetWorksReply)
	err := c.cc.Invoke(ctx, "/pb.WorkManager/GetWorks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkManagerServer is the server API for WorkManager service.
// All implementations must embed UnimplementedWorkManagerServer
// for forward compatibility
type WorkManagerServer interface {
	CreateWork(context.Context, *CreateWorkRequest) (*Work, error)
	GetWork(context.Context, *GetWorkRequest) (*Work, error)
	UpdateWork(context.Context, *UpdateWorkRequest) (*Work, error)
	DeleteWork(context.Context, *DeleteWorkRequest) (*Work, error)
	GetWorks(context.Context, *GetWorksRequest) (*GetWorksReply, error)
	mustEmbedUnimplementedWorkManagerServer()
}

// UnimplementedWorkManagerServer must be embedded to have forward compatible implementations.
type UnimplementedWorkManagerServer struct {
}

func (UnimplementedWorkManagerServer) CreateWork(context.Context, *CreateWorkRequest) (*Work, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWork not implemented")
}
func (UnimplementedWorkManagerServer) GetWork(context.Context, *GetWorkRequest) (*Work, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWork not implemented")
}
func (UnimplementedWorkManagerServer) UpdateWork(context.Context, *UpdateWorkRequest) (*Work, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWork not implemented")
}
func (UnimplementedWorkManagerServer) DeleteWork(context.Context, *DeleteWorkRequest) (*Work, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWork not implemented")
}
func (UnimplementedWorkManagerServer) GetWorks(context.Context, *GetWorksRequest) (*GetWorksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorks not implemented")
}
func (UnimplementedWorkManagerServer) mustEmbedUnimplementedWorkManagerServer() {}

// UnsafeWorkManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkManagerServer will
// result in compilation errors.
type UnsafeWorkManagerServer interface {
	mustEmbedUnimplementedWorkManagerServer()
}

func RegisterWorkManagerServer(s grpc.ServiceRegistrar, srv WorkManagerServer) {
	s.RegisterService(&WorkManager_ServiceDesc, srv)
}

func _WorkManager_CreateWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkManagerServer).CreateWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkManager/CreateWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkManagerServer).CreateWork(ctx, req.(*CreateWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkManager_GetWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkManagerServer).GetWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkManager/GetWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkManagerServer).GetWork(ctx, req.(*GetWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkManager_UpdateWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkManagerServer).UpdateWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkManager/UpdateWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkManagerServer).UpdateWork(ctx, req.(*UpdateWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkManager_DeleteWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkManagerServer).DeleteWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkManager/DeleteWork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkManagerServer).DeleteWork(ctx, req.(*DeleteWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkManager_GetWorks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkManagerServer).GetWorks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WorkManager/GetWorks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkManagerServer).GetWorks(ctx, req.(*GetWorksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkManager_ServiceDesc is the grpc.ServiceDesc for WorkManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.WorkManager",
	HandlerType: (*WorkManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWork",
			Handler:    _WorkManager_CreateWork_Handler,
		},
		{
			MethodName: "GetWork",
			Handler:    _WorkManager_GetWork_Handler,
		},
		{
			MethodName: "UpdateWork",
			Handler:    _WorkManager_UpdateWork_Handler,
		},
		{
			MethodName: "DeleteWork",
			Handler:    _WorkManager_DeleteWork_Handler,
		},
		{
			MethodName: "GetWorks",
			Handler:    _WorkManager_GetWorks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "work.proto",
}

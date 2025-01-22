// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: provider/plugin/common/plugin.proto

package common

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
	ProviderPlugin_SetLogger_FullMethodName   = "/common.ProviderPlugin/SetLogger"
	ProviderPlugin_Configure_FullMethodName   = "/common.ProviderPlugin/Configure"
	ProviderPlugin_SetDataDir_FullMethodName  = "/common.ProviderPlugin/SetDataDir"
	ProviderPlugin_SetCacheDir_FullMethodName = "/common.ProviderPlugin/SetCacheDir"
	ProviderPlugin_Invalidate_FullMethodName  = "/common.ProviderPlugin/Invalidate"
	ProviderPlugin_Keep_FullMethodName        = "/common.ProviderPlugin/Keep"
	ProviderPlugin_Load_FullMethodName        = "/common.ProviderPlugin/Load"
)

// ProviderPluginClient is the client API for ProviderPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderPluginClient interface {
	SetLogger(ctx context.Context, in *SetLoggerRequest, opts ...grpc.CallOption) (*SetLoggerResponse, error)
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error)
	SetDataDir(ctx context.Context, in *SetDataDirRequest, opts ...grpc.CallOption) (*SetDataDirResponse, error)
	SetCacheDir(ctx context.Context, in *SetCacheDirRequest, opts ...grpc.CallOption) (*Empty, error)
	Invalidate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Keep(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error)
}

type providerPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderPluginClient(cc grpc.ClientConnInterface) ProviderPluginClient {
	return &providerPluginClient{cc}
}

func (c *providerPluginClient) SetLogger(ctx context.Context, in *SetLoggerRequest, opts ...grpc.CallOption) (*SetLoggerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetLoggerResponse)
	err := c.cc.Invoke(ctx, ProviderPlugin_SetLogger_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfigureResponse)
	err := c.cc.Invoke(ctx, ProviderPlugin_Configure_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) SetDataDir(ctx context.Context, in *SetDataDirRequest, opts ...grpc.CallOption) (*SetDataDirResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetDataDirResponse)
	err := c.cc.Invoke(ctx, ProviderPlugin_SetDataDir_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) SetCacheDir(ctx context.Context, in *SetCacheDirRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ProviderPlugin_SetCacheDir_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) Invalidate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ProviderPlugin_Invalidate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) Keep(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, ProviderPlugin_Keep_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerPluginClient) Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadResponse)
	err := c.cc.Invoke(ctx, ProviderPlugin_Load_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderPluginServer is the server API for ProviderPlugin service.
// All implementations must embed UnimplementedProviderPluginServer
// for forward compatibility.
type ProviderPluginServer interface {
	SetLogger(context.Context, *SetLoggerRequest) (*SetLoggerResponse, error)
	Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error)
	SetDataDir(context.Context, *SetDataDirRequest) (*SetDataDirResponse, error)
	SetCacheDir(context.Context, *SetCacheDirRequest) (*Empty, error)
	Invalidate(context.Context, *Empty) (*Empty, error)
	Keep(context.Context, *Empty) (*Empty, error)
	Load(context.Context, *LoadRequest) (*LoadResponse, error)
	mustEmbedUnimplementedProviderPluginServer()
}

// UnimplementedProviderPluginServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProviderPluginServer struct{}

func (UnimplementedProviderPluginServer) SetLogger(context.Context, *SetLoggerRequest) (*SetLoggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLogger not implemented")
}
func (UnimplementedProviderPluginServer) Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedProviderPluginServer) SetDataDir(context.Context, *SetDataDirRequest) (*SetDataDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDataDir not implemented")
}
func (UnimplementedProviderPluginServer) SetCacheDir(context.Context, *SetCacheDirRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCacheDir not implemented")
}
func (UnimplementedProviderPluginServer) Invalidate(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invalidate not implemented")
}
func (UnimplementedProviderPluginServer) Keep(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Keep not implemented")
}
func (UnimplementedProviderPluginServer) Load(context.Context, *LoadRequest) (*LoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedProviderPluginServer) mustEmbedUnimplementedProviderPluginServer() {}
func (UnimplementedProviderPluginServer) testEmbeddedByValue()                        {}

// UnsafeProviderPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderPluginServer will
// result in compilation errors.
type UnsafeProviderPluginServer interface {
	mustEmbedUnimplementedProviderPluginServer()
}

func RegisterProviderPluginServer(s grpc.ServiceRegistrar, srv ProviderPluginServer) {
	// If the following call pancis, it indicates UnimplementedProviderPluginServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProviderPlugin_ServiceDesc, srv)
}

func _ProviderPlugin_SetLogger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLoggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).SetLogger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderPlugin_SetLogger_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).SetLogger(ctx, req.(*SetLoggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderPlugin_Configure_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_SetDataDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDataDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).SetDataDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderPlugin_SetDataDir_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).SetDataDir(ctx, req.(*SetDataDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_SetCacheDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCacheDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).SetCacheDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderPlugin_SetCacheDir_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).SetCacheDir(ctx, req.(*SetCacheDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_Invalidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).Invalidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderPlugin_Invalidate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).Invalidate(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_Keep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).Keep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderPlugin_Keep_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).Keep(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderPlugin_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderPluginServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderPlugin_Load_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderPluginServer).Load(ctx, req.(*LoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProviderPlugin_ServiceDesc is the grpc.ServiceDesc for ProviderPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProviderPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.ProviderPlugin",
	HandlerType: (*ProviderPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLogger",
			Handler:    _ProviderPlugin_SetLogger_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _ProviderPlugin_Configure_Handler,
		},
		{
			MethodName: "SetDataDir",
			Handler:    _ProviderPlugin_SetDataDir_Handler,
		},
		{
			MethodName: "SetCacheDir",
			Handler:    _ProviderPlugin_SetCacheDir_Handler,
		},
		{
			MethodName: "Invalidate",
			Handler:    _ProviderPlugin_Invalidate_Handler,
		},
		{
			MethodName: "Keep",
			Handler:    _ProviderPlugin_Keep_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _ProviderPlugin_Load_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider/plugin/common/plugin.proto",
}

const (
	Logger_LoadingMessage_FullMethodName = "/common.Logger/LoadingMessage"
	Logger_EmitLogMessage_FullMethodName = "/common.Logger/EmitLogMessage"
)

// LoggerClient is the client API for Logger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggerClient interface {
	LoadingMessage(ctx context.Context, in *LoadingMessageRequest, opts ...grpc.CallOption) (*Empty, error)
	EmitLogMessage(ctx context.Context, in *EmitLogMessageRequest, opts ...grpc.CallOption) (*Empty, error)
}

type loggerClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggerClient(cc grpc.ClientConnInterface) LoggerClient {
	return &loggerClient{cc}
}

func (c *loggerClient) LoadingMessage(ctx context.Context, in *LoadingMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Logger_LoadingMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggerClient) EmitLogMessage(ctx context.Context, in *EmitLogMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Logger_EmitLogMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServer is the server API for Logger service.
// All implementations must embed UnimplementedLoggerServer
// for forward compatibility.
type LoggerServer interface {
	LoadingMessage(context.Context, *LoadingMessageRequest) (*Empty, error)
	EmitLogMessage(context.Context, *EmitLogMessageRequest) (*Empty, error)
	mustEmbedUnimplementedLoggerServer()
}

// UnimplementedLoggerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLoggerServer struct{}

func (UnimplementedLoggerServer) LoadingMessage(context.Context, *LoadingMessageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadingMessage not implemented")
}
func (UnimplementedLoggerServer) EmitLogMessage(context.Context, *EmitLogMessageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitLogMessage not implemented")
}
func (UnimplementedLoggerServer) mustEmbedUnimplementedLoggerServer() {}
func (UnimplementedLoggerServer) testEmbeddedByValue()                {}

// UnsafeLoggerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggerServer will
// result in compilation errors.
type UnsafeLoggerServer interface {
	mustEmbedUnimplementedLoggerServer()
}

func RegisterLoggerServer(s grpc.ServiceRegistrar, srv LoggerServer) {
	// If the following call pancis, it indicates UnimplementedLoggerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Logger_ServiceDesc, srv)
}

func _Logger_LoadingMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadingMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).LoadingMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logger_LoadingMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).LoadingMessage(ctx, req.(*LoadingMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logger_EmitLogMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitLogMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).EmitLogMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logger_EmitLogMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).EmitLogMessage(ctx, req.(*EmitLogMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Logger_ServiceDesc is the grpc.ServiceDesc for Logger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Logger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "common.Logger",
	HandlerType: (*LoggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadingMessage",
			Handler:    _Logger_LoadingMessage_Handler,
		},
		{
			MethodName: "EmitLogMessage",
			Handler:    _Logger_EmitLogMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider/plugin/common/plugin.proto",
}

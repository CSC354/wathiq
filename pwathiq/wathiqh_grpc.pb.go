// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: protobufs/wathiq/wathiqh.proto

package pwathiq

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

// WathiqClient is the client API for Wathiq service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WathiqClient interface {
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	GetToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type wathiqClient struct {
	cc grpc.ClientConnInterface
}

func NewWathiqClient(cc grpc.ClientConnInterface) WathiqClient {
	return &wathiqClient{cc}
}

func (c *wathiqClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/wathiq.Wathiq/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wathiqClient) GetToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/wathiq.Wathiq/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WathiqServer is the server API for Wathiq service.
// All implementations must embed UnimplementedWathiqServer
// for forward compatibility
type WathiqServer interface {
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	GetToken(context.Context, *TokenRequest) (*TokenResponse, error)
	mustEmbedUnimplementedWathiqServer()
}

// UnimplementedWathiqServer must be embedded to have forward compatible implementations.
type UnimplementedWathiqServer struct {
}

func (UnimplementedWathiqServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedWathiqServer) GetToken(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedWathiqServer) mustEmbedUnimplementedWathiqServer() {}

// UnsafeWathiqServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WathiqServer will
// result in compilation errors.
type UnsafeWathiqServer interface {
	mustEmbedUnimplementedWathiqServer()
}

func RegisterWathiqServer(s grpc.ServiceRegistrar, srv WathiqServer) {
	s.RegisterService(&Wathiq_ServiceDesc, srv)
}

func _Wathiq_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WathiqServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wathiq.Wathiq/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WathiqServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wathiq_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WathiqServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wathiq.Wathiq/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WathiqServer).GetToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wathiq_ServiceDesc is the grpc.ServiceDesc for Wathiq service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wathiq_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wathiq.Wathiq",
	HandlerType: (*WathiqServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _Wathiq_Validate_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _Wathiq_GetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobufs/wathiq/wathiqh.proto",
}

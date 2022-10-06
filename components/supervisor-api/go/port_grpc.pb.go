// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: port.proto

package api

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

// PortServiceClient is the client API for PortService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortServiceClient interface {
	// Tunnel notifies clients to install listeners on remote machines.
	// After that such clients should call EstablishTunnel to forward incoming connections.
	Tunnel(ctx context.Context, in *TunnelPortRequest, opts ...grpc.CallOption) (*TunnelPortResponse, error)
	// CloseTunnel notifies clients to remove listeners on remote machines.
	CloseTunnel(ctx context.Context, in *CloseTunnelRequest, opts ...grpc.CallOption) (*CloseTunnelResponse, error)
	// EstablishTunnel actually establishes the tunnel for an incoming connection on a remote machine.
	EstablishTunnel(ctx context.Context, opts ...grpc.CallOption) (PortService_EstablishTunnelClient, error)
	// AutoTunnel controls enablement of auto tunneling
	AutoTunnel(ctx context.Context, in *AutoTunnelRequest, opts ...grpc.CallOption) (*AutoTunnelResponse, error)
	// RetryAutoExpose retries auto exposing the give port
	RetryAutoExpose(ctx context.Context, in *RetryAutoExposeRequest, opts ...grpc.CallOption) (*RetryAutoExposeResponse, error)
}

type portServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortServiceClient(cc grpc.ClientConnInterface) PortServiceClient {
	return &portServiceClient{cc}
}

func (c *portServiceClient) Tunnel(ctx context.Context, in *TunnelPortRequest, opts ...grpc.CallOption) (*TunnelPortResponse, error) {
	out := new(TunnelPortResponse)
	err := c.cc.Invoke(ctx, "/supervisor.PortService/Tunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portServiceClient) CloseTunnel(ctx context.Context, in *CloseTunnelRequest, opts ...grpc.CallOption) (*CloseTunnelResponse, error) {
	out := new(CloseTunnelResponse)
	err := c.cc.Invoke(ctx, "/supervisor.PortService/CloseTunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portServiceClient) EstablishTunnel(ctx context.Context, opts ...grpc.CallOption) (PortService_EstablishTunnelClient, error) {
	stream, err := c.cc.NewStream(ctx, &PortService_ServiceDesc.Streams[0], "/supervisor.PortService/EstablishTunnel", opts...)
	if err != nil {
		return nil, err
	}
	x := &portServiceEstablishTunnelClient{stream}
	return x, nil
}

type PortService_EstablishTunnelClient interface {
	Send(*EstablishTunnelRequest) error
	Recv() (*EstablishTunnelResponse, error)
	grpc.ClientStream
}

type portServiceEstablishTunnelClient struct {
	grpc.ClientStream
}

func (x *portServiceEstablishTunnelClient) Send(m *EstablishTunnelRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *portServiceEstablishTunnelClient) Recv() (*EstablishTunnelResponse, error) {
	m := new(EstablishTunnelResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *portServiceClient) AutoTunnel(ctx context.Context, in *AutoTunnelRequest, opts ...grpc.CallOption) (*AutoTunnelResponse, error) {
	out := new(AutoTunnelResponse)
	err := c.cc.Invoke(ctx, "/supervisor.PortService/AutoTunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portServiceClient) RetryAutoExpose(ctx context.Context, in *RetryAutoExposeRequest, opts ...grpc.CallOption) (*RetryAutoExposeResponse, error) {
	out := new(RetryAutoExposeResponse)
	err := c.cc.Invoke(ctx, "/supervisor.PortService/RetryAutoExpose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortServiceServer is the server API for PortService service.
// All implementations must embed UnimplementedPortServiceServer
// for forward compatibility
type PortServiceServer interface {
	// Tunnel notifies clients to install listeners on remote machines.
	// After that such clients should call EstablishTunnel to forward incoming connections.
	Tunnel(context.Context, *TunnelPortRequest) (*TunnelPortResponse, error)
	// CloseTunnel notifies clients to remove listeners on remote machines.
	CloseTunnel(context.Context, *CloseTunnelRequest) (*CloseTunnelResponse, error)
	// EstablishTunnel actually establishes the tunnel for an incoming connection on a remote machine.
	EstablishTunnel(PortService_EstablishTunnelServer) error
	// AutoTunnel controls enablement of auto tunneling
	AutoTunnel(context.Context, *AutoTunnelRequest) (*AutoTunnelResponse, error)
	// RetryAutoExpose retries auto exposing the give port
	RetryAutoExpose(context.Context, *RetryAutoExposeRequest) (*RetryAutoExposeResponse, error)
	mustEmbedUnimplementedPortServiceServer()
}

// UnimplementedPortServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortServiceServer struct {
}

func (UnimplementedPortServiceServer) Tunnel(context.Context, *TunnelPortRequest) (*TunnelPortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tunnel not implemented")
}
func (UnimplementedPortServiceServer) CloseTunnel(context.Context, *CloseTunnelRequest) (*CloseTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTunnel not implemented")
}
func (UnimplementedPortServiceServer) EstablishTunnel(PortService_EstablishTunnelServer) error {
	return status.Errorf(codes.Unimplemented, "method EstablishTunnel not implemented")
}
func (UnimplementedPortServiceServer) AutoTunnel(context.Context, *AutoTunnelRequest) (*AutoTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AutoTunnel not implemented")
}
func (UnimplementedPortServiceServer) RetryAutoExpose(context.Context, *RetryAutoExposeRequest) (*RetryAutoExposeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetryAutoExpose not implemented")
}
func (UnimplementedPortServiceServer) mustEmbedUnimplementedPortServiceServer() {}

// UnsafePortServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortServiceServer will
// result in compilation errors.
type UnsafePortServiceServer interface {
	mustEmbedUnimplementedPortServiceServer()
}

func RegisterPortServiceServer(s grpc.ServiceRegistrar, srv PortServiceServer) {
	s.RegisterService(&PortService_ServiceDesc, srv)
}

func _PortService_Tunnel_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(TunnelPortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortServiceServer).Tunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.PortService/Tunnel",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(PortServiceServer).Tunnel(ctx, req.(*TunnelPortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortService_CloseTunnel_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(CloseTunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortServiceServer).CloseTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.PortService/CloseTunnel",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(PortServiceServer).CloseTunnel(ctx, req.(*CloseTunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortService_EstablishTunnel_Handler(srv any, stream grpc.ServerStream) error {
	return srv.(PortServiceServer).EstablishTunnel(&portServiceEstablishTunnelServer{stream})
}

type PortService_EstablishTunnelServer interface {
	Send(*EstablishTunnelResponse) error
	Recv() (*EstablishTunnelRequest, error)
	grpc.ServerStream
}

type portServiceEstablishTunnelServer struct {
	grpc.ServerStream
}

func (x *portServiceEstablishTunnelServer) Send(m *EstablishTunnelResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *portServiceEstablishTunnelServer) Recv() (*EstablishTunnelRequest, error) {
	m := new(EstablishTunnelRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PortService_AutoTunnel_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(AutoTunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortServiceServer).AutoTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.PortService/AutoTunnel",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(PortServiceServer).AutoTunnel(ctx, req.(*AutoTunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortService_RetryAutoExpose_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(RetryAutoExposeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortServiceServer).RetryAutoExpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.PortService/RetryAutoExpose",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(PortServiceServer).RetryAutoExpose(ctx, req.(*RetryAutoExposeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PortService_ServiceDesc is the grpc.ServiceDesc for PortService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.PortService",
	HandlerType: (*PortServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tunnel",
			Handler:    _PortService_Tunnel_Handler,
		},
		{
			MethodName: "CloseTunnel",
			Handler:    _PortService_CloseTunnel_Handler,
		},
		{
			MethodName: "AutoTunnel",
			Handler:    _PortService_AutoTunnel_Handler,
		},
		{
			MethodName: "RetryAutoExpose",
			Handler:    _PortService_RetryAutoExpose_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EstablishTunnel",
			Handler:       _PortService_EstablishTunnel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "port.proto",
}

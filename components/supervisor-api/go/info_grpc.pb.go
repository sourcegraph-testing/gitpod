// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: info.proto

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

// InfoServiceClient is the client API for InfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoServiceClient interface {
	WorkspaceInfo(ctx context.Context, in *WorkspaceInfoRequest, opts ...grpc.CallOption) (*WorkspaceInfoResponse, error)
}

type infoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoServiceClient(cc grpc.ClientConnInterface) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) WorkspaceInfo(ctx context.Context, in *WorkspaceInfoRequest, opts ...grpc.CallOption) (*WorkspaceInfoResponse, error) {
	out := new(WorkspaceInfoResponse)
	err := c.cc.Invoke(ctx, "/supervisor.InfoService/WorkspaceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServiceServer is the server API for InfoService service.
// All implementations must embed UnimplementedInfoServiceServer
// for forward compatibility
type InfoServiceServer interface {
	WorkspaceInfo(context.Context, *WorkspaceInfoRequest) (*WorkspaceInfoResponse, error)
	mustEmbedUnimplementedInfoServiceServer()
}

// UnimplementedInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInfoServiceServer struct {
}

func (UnimplementedInfoServiceServer) WorkspaceInfo(context.Context, *WorkspaceInfoRequest) (*WorkspaceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkspaceInfo not implemented")
}
func (UnimplementedInfoServiceServer) mustEmbedUnimplementedInfoServiceServer() {}

// UnsafeInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoServiceServer will
// result in compilation errors.
type UnsafeInfoServiceServer interface {
	mustEmbedUnimplementedInfoServiceServer()
}

func RegisterInfoServiceServer(s grpc.ServiceRegistrar, srv InfoServiceServer) {
	s.RegisterService(&InfoService_ServiceDesc, srv)
}

func _InfoService_WorkspaceInfo_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(WorkspaceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).WorkspaceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.InfoService/WorkspaceInfo",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(InfoServiceServer).WorkspaceInfo(ctx, req.(*WorkspaceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InfoService_ServiceDesc is the grpc.ServiceDesc for InfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WorkspaceInfo",
			Handler:    _InfoService_WorkspaceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info.proto",
}

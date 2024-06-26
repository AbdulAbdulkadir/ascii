// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: service.proto

package proto

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

// AsciiServiceClient is the client API for AsciiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AsciiServiceClient interface {
	// Displays a random ascii art retrieved from a database
	DisplayAscii(ctx context.Context, in *DisplayRequest, opts ...grpc.CallOption) (*DisplayResponse, error)
	// Uploads a clients ascii art to the database
	UploadAscii(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
}

type asciiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAsciiServiceClient(cc grpc.ClientConnInterface) AsciiServiceClient {
	return &asciiServiceClient{cc}
}

func (c *asciiServiceClient) DisplayAscii(ctx context.Context, in *DisplayRequest, opts ...grpc.CallOption) (*DisplayResponse, error) {
	out := new(DisplayResponse)
	err := c.cc.Invoke(ctx, "/proto.AsciiService/DisplayAscii", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asciiServiceClient) UploadAscii(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, "/proto.AsciiService/UploadAscii", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AsciiServiceServer is the server API for AsciiService service.
// All implementations must embed UnimplementedAsciiServiceServer
// for forward compatibility
type AsciiServiceServer interface {
	// Displays a random ascii art retrieved from a database
	DisplayAscii(context.Context, *DisplayRequest) (*DisplayResponse, error)
	// Uploads a clients ascii art to the database
	UploadAscii(context.Context, *UploadRequest) (*UploadResponse, error)
	mustEmbedUnimplementedAsciiServiceServer()
}

// UnimplementedAsciiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAsciiServiceServer struct {
}

func (UnimplementedAsciiServiceServer) DisplayAscii(context.Context, *DisplayRequest) (*DisplayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayAscii not implemented")
}
func (UnimplementedAsciiServiceServer) UploadAscii(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAscii not implemented")
}
func (UnimplementedAsciiServiceServer) mustEmbedUnimplementedAsciiServiceServer() {}

// UnsafeAsciiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AsciiServiceServer will
// result in compilation errors.
type UnsafeAsciiServiceServer interface {
	mustEmbedUnimplementedAsciiServiceServer()
}

func RegisterAsciiServiceServer(s grpc.ServiceRegistrar, srv AsciiServiceServer) {
	s.RegisterService(&AsciiService_ServiceDesc, srv)
}

func _AsciiService_DisplayAscii_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsciiServiceServer).DisplayAscii(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AsciiService/DisplayAscii",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsciiServiceServer).DisplayAscii(ctx, req.(*DisplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsciiService_UploadAscii_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsciiServiceServer).UploadAscii(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AsciiService/UploadAscii",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsciiServiceServer).UploadAscii(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AsciiService_ServiceDesc is the grpc.ServiceDesc for AsciiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AsciiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AsciiService",
	HandlerType: (*AsciiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DisplayAscii",
			Handler:    _AsciiService_DisplayAscii_Handler,
		},
		{
			MethodName: "UploadAscii",
			Handler:    _AsciiService_UploadAscii_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

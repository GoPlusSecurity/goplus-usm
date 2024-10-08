// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: detect/service/v1/detect.proto

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
	Detect_DetectTx_FullMethodName = "/detect.service.v1.Detect/DetectTx"
)

// DetectClient is the client API for Detect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DetectClient interface {
	// tx detect
	DetectTx(ctx context.Context, in *DetectTxRequest, opts ...grpc.CallOption) (*DetectTxResponse, error)
}

type detectClient struct {
	cc grpc.ClientConnInterface
}

func NewDetectClient(cc grpc.ClientConnInterface) DetectClient {
	return &detectClient{cc}
}

func (c *detectClient) DetectTx(ctx context.Context, in *DetectTxRequest, opts ...grpc.CallOption) (*DetectTxResponse, error) {
	out := new(DetectTxResponse)
	err := c.cc.Invoke(ctx, Detect_DetectTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetectServer is the server API for Detect service.
// All implementations must embed UnimplementedDetectServer
// for forward compatibility
type DetectServer interface {
	// tx detect
	DetectTx(context.Context, *DetectTxRequest) (*DetectTxResponse, error)
	mustEmbedUnimplementedDetectServer()
}

// UnimplementedDetectServer must be embedded to have forward compatible implementations.
type UnimplementedDetectServer struct {
}

func (UnimplementedDetectServer) DetectTx(context.Context, *DetectTxRequest) (*DetectTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectTx not implemented")
}
func (UnimplementedDetectServer) mustEmbedUnimplementedDetectServer() {}

// UnsafeDetectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DetectServer will
// result in compilation errors.
type UnsafeDetectServer interface {
	mustEmbedUnimplementedDetectServer()
}

func RegisterDetectServer(s grpc.ServiceRegistrar, srv DetectServer) {
	s.RegisterService(&Detect_ServiceDesc, srv)
}

func _Detect_DetectTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectServer).DetectTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Detect_DetectTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectServer).DetectTx(ctx, req.(*DetectTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Detect_ServiceDesc is the grpc.ServiceDesc for Detect service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Detect_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "detect.service.v1.Detect",
	HandlerType: (*DetectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DetectTx",
			Handler:    _Detect_DetectTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/detect.proto",
}

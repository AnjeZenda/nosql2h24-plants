// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: data/v1/data.proto

package data_v1

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
	DataAPI_ExportDBV1_FullMethodName = "/data.v1.DataAPI/ExportDBV1"
	DataAPI_ImportDBV1_FullMethodName = "/data.v1.DataAPI/ImportDBV1"
)

// DataAPIClient is the client API for DataAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataAPIClient interface {
	ExportDBV1(ctx context.Context, in *ExportDBV1Request, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ExportDBV1Response], error)
	ImportDBV1(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ImportDBV1Request, ImportDBV1Response], error)
}

type dataAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDataAPIClient(cc grpc.ClientConnInterface) DataAPIClient {
	return &dataAPIClient{cc}
}

func (c *dataAPIClient) ExportDBV1(ctx context.Context, in *ExportDBV1Request, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ExportDBV1Response], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DataAPI_ServiceDesc.Streams[0], DataAPI_ExportDBV1_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ExportDBV1Request, ExportDBV1Response]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DataAPI_ExportDBV1Client = grpc.ServerStreamingClient[ExportDBV1Response]

func (c *dataAPIClient) ImportDBV1(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ImportDBV1Request, ImportDBV1Response], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DataAPI_ServiceDesc.Streams[1], DataAPI_ImportDBV1_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ImportDBV1Request, ImportDBV1Response]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DataAPI_ImportDBV1Client = grpc.ClientStreamingClient[ImportDBV1Request, ImportDBV1Response]

// DataAPIServer is the server API for DataAPI service.
// All implementations must embed UnimplementedDataAPIServer
// for forward compatibility.
type DataAPIServer interface {
	ExportDBV1(*ExportDBV1Request, grpc.ServerStreamingServer[ExportDBV1Response]) error
	ImportDBV1(grpc.ClientStreamingServer[ImportDBV1Request, ImportDBV1Response]) error
	mustEmbedUnimplementedDataAPIServer()
}

// UnimplementedDataAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataAPIServer struct{}

func (UnimplementedDataAPIServer) ExportDBV1(*ExportDBV1Request, grpc.ServerStreamingServer[ExportDBV1Response]) error {
	return status.Errorf(codes.Unimplemented, "method ExportDBV1 not implemented")
}
func (UnimplementedDataAPIServer) ImportDBV1(grpc.ClientStreamingServer[ImportDBV1Request, ImportDBV1Response]) error {
	return status.Errorf(codes.Unimplemented, "method ImportDBV1 not implemented")
}
func (UnimplementedDataAPIServer) mustEmbedUnimplementedDataAPIServer() {}
func (UnimplementedDataAPIServer) testEmbeddedByValue()                 {}

// UnsafeDataAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataAPIServer will
// result in compilation errors.
type UnsafeDataAPIServer interface {
	mustEmbedUnimplementedDataAPIServer()
}

func RegisterDataAPIServer(s grpc.ServiceRegistrar, srv DataAPIServer) {
	// If the following call pancis, it indicates UnimplementedDataAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataAPI_ServiceDesc, srv)
}

func _DataAPI_ExportDBV1_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExportDBV1Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataAPIServer).ExportDBV1(m, &grpc.GenericServerStream[ExportDBV1Request, ExportDBV1Response]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DataAPI_ExportDBV1Server = grpc.ServerStreamingServer[ExportDBV1Response]

func _DataAPI_ImportDBV1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataAPIServer).ImportDBV1(&grpc.GenericServerStream[ImportDBV1Request, ImportDBV1Response]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DataAPI_ImportDBV1Server = grpc.ClientStreamingServer[ImportDBV1Request, ImportDBV1Response]

// DataAPI_ServiceDesc is the grpc.ServiceDesc for DataAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data.v1.DataAPI",
	HandlerType: (*DataAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExportDBV1",
			Handler:       _DataAPI_ExportDBV1_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ImportDBV1",
			Handler:       _DataAPI_ImportDBV1_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "data/v1/data.proto",
}

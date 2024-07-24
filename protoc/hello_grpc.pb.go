// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: hello.proto

package __

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

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleClient interface {
	// Bi-directional streaming
	ServerReply(ctx context.Context, opts ...grpc.CallOption) (Example_ServerReplyClient, error)
}

type exampleClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleClient(cc grpc.ClientConnInterface) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) ServerReply(ctx context.Context, opts ...grpc.CallOption) (Example_ServerReplyClient, error) {
	stream, err := c.cc.NewStream(ctx, &Example_ServiceDesc.Streams[0], "/Example/ServerReply", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServerReplyClient{stream}
	return x, nil
}

type Example_ServerReplyClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type exampleServerReplyClient struct {
	grpc.ClientStream
}

func (x *exampleServerReplyClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exampleServerReplyClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExampleServer is the server API for Example service.
// All implementations must embed UnimplementedExampleServer
// for forward compatibility
type ExampleServer interface {
	// Bi-directional streaming
	ServerReply(Example_ServerReplyServer) error
	mustEmbedUnimplementedExampleServer()
}

// UnimplementedExampleServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (UnimplementedExampleServer) ServerReply(Example_ServerReplyServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerReply not implemented")
}
func (UnimplementedExampleServer) mustEmbedUnimplementedExampleServer() {}

// UnsafeExampleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServer will
// result in compilation errors.
type UnsafeExampleServer interface {
	mustEmbedUnimplementedExampleServer()
}

func RegisterExampleServer(s grpc.ServiceRegistrar, srv ExampleServer) {
	s.RegisterService(&Example_ServiceDesc, srv)
}

func _Example_ServerReply_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServer).ServerReply(&exampleServerReplyServer{stream})
}

type Example_ServerReplyServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type exampleServerReplyServer struct {
	grpc.ServerStream
}

func (x *exampleServerReplyServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exampleServerReplyServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Example_ServiceDesc is the grpc.ServiceDesc for Example service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Example_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Example",
	HandlerType: (*ExampleServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerReply",
			Handler:       _Example_ServerReply_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}

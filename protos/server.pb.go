// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.0
// source: protos/server.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protos_server_proto protoreflect.FileDescriptor

var file_protos_server_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x14, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x6c, 0x0a, 0x0a, 0x47, 0x72,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_server_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: protos.Request
	(*Response)(nil), // 1: protos.Response
}
var file_protos_server_proto_depIdxs = []int32{
	0, // 0: protos.GrpcServer.Send:input_type -> protos.Request
	0, // 1: protos.GrpcServer.SendStream:input_type -> protos.Request
	1, // 2: protos.GrpcServer.Send:output_type -> protos.Response
	1, // 3: protos.GrpcServer.SendStream:output_type -> protos.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_server_proto_init() }
func file_protos_server_proto_init() {
	if File_protos_server_proto != nil {
		return
	}
	file_protos_request_proto_init()
	file_protos_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_server_proto_goTypes,
		DependencyIndexes: file_protos_server_proto_depIdxs,
	}.Build()
	File_protos_server_proto = out.File
	file_protos_server_proto_rawDesc = nil
	file_protos_server_proto_goTypes = nil
	file_protos_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GrpcServerClient is the client API for GrpcServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcServerClient interface {
	Send(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	SendStream(ctx context.Context, opts ...grpc.CallOption) (GrpcServer_SendStreamClient, error)
}

type grpcServerClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcServerClient(cc grpc.ClientConnInterface) GrpcServerClient {
	return &grpcServerClient{cc}
}

func (c *grpcServerClient) Send(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.GrpcServer/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcServerClient) SendStream(ctx context.Context, opts ...grpc.CallOption) (GrpcServer_SendStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GrpcServer_serviceDesc.Streams[0], "/protos.GrpcServer/SendStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcServerSendStreamClient{stream}
	return x, nil
}

type GrpcServer_SendStreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type grpcServerSendStreamClient struct {
	grpc.ClientStream
}

func (x *grpcServerSendStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *grpcServerSendStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GrpcServerServer is the server API for GrpcServer service.
type GrpcServerServer interface {
	Send(context.Context, *Request) (*Response, error)
	SendStream(GrpcServer_SendStreamServer) error
}

// UnimplementedGrpcServerServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcServerServer struct {
}

func (*UnimplementedGrpcServerServer) Send(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedGrpcServerServer) SendStream(GrpcServer_SendStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SendStream not implemented")
}

func RegisterGrpcServerServer(s *grpc.Server, srv GrpcServerServer) {
	s.RegisterService(&_GrpcServer_serviceDesc, srv)
}

func _GrpcServer_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServerServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GrpcServer/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServerServer).Send(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcServer_SendStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GrpcServerServer).SendStream(&grpcServerSendStreamServer{stream})
}

type GrpcServer_SendStreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type grpcServerSendStreamServer struct {
	grpc.ServerStream
}

func (x *grpcServerSendStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *grpcServerSendStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GrpcServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.GrpcServer",
	HandlerType: (*GrpcServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _GrpcServer_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendStream",
			Handler:       _GrpcServer_SendStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/server.proto",
}

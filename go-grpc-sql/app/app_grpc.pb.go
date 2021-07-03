// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package app

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

// StudentsClient is the client API for Students service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentsClient interface {
	// Sends a greeting
	GetStudents(ctx context.Context, in *StudentRequest, opts ...grpc.CallOption) (*StudentResp, error)
}

type studentsClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentsClient(cc grpc.ClientConnInterface) StudentsClient {
	return &studentsClient{cc}
}

func (c *studentsClient) GetStudents(ctx context.Context, in *StudentRequest, opts ...grpc.CallOption) (*StudentResp, error) {
	out := new(StudentResp)
	err := c.cc.Invoke(ctx, "/app.Students/GetStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentsServer is the server API for Students service.
// All implementations must embed UnimplementedStudentsServer
// for forward compatibility
type StudentsServer interface {
	// Sends a greeting
	GetStudents(context.Context, *StudentRequest) (*StudentResp, error)
	mustEmbedUnimplementedStudentsServer()
}

// UnimplementedStudentsServer must be embedded to have forward compatible implementations.
type UnimplementedStudentsServer struct {
}

func (UnimplementedStudentsServer) GetStudents(context.Context, *StudentRequest) (*StudentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudents not implemented")
}
func (UnimplementedStudentsServer) mustEmbedUnimplementedStudentsServer() {}

// UnsafeStudentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentsServer will
// result in compilation errors.
type UnsafeStudentsServer interface {
	mustEmbedUnimplementedStudentsServer()
}

func RegisterStudentsServer(s grpc.ServiceRegistrar, srv StudentsServer) {
	s.RegisterService(&Students_ServiceDesc, srv)
}

func _Students_GetStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentsServer).GetStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.Students/GetStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentsServer).GetStudents(ctx, req.(*StudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Students_ServiceDesc is the grpc.ServiceDesc for Students service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Students_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.Students",
	HandlerType: (*StudentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudents",
			Handler:    _Students_GetStudents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/app.proto",
}
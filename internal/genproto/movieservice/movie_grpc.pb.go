// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: protos/movie-service/movie.proto

package movieservice

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

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieServiceClient interface {
	AddMovie(ctx context.Context, in *AddMovieReq, opts ...grpc.CallOption) (*AddMovieRes, error)
	GetMovieById(ctx context.Context, in *GetMovieByIdReq, opts ...grpc.CallOption) (*GetMovieByIdRes, error)
	UpdateMovie(ctx context.Context, in *UpdateMovieReq, opts ...grpc.CallOption) (*UpdateMovieRes, error)
	DeleteMovie(ctx context.Context, in *DeleteMovieReq, opts ...grpc.CallOption) (*DeleteMovieRes, error)
	RemoveMovie(ctx context.Context, in *RemoveMovieReq, opts ...grpc.CallOption) (*RemoveMovieRes, error)
	GetAllMovies(ctx context.Context, in *GetAllMoviesReq, opts ...grpc.CallOption) (*GetAllMoviesRes, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) AddMovie(ctx context.Context, in *AddMovieReq, opts ...grpc.CallOption) (*AddMovieRes, error) {
	out := new(AddMovieRes)
	err := c.cc.Invoke(ctx, "/movieservice.MovieService/AddMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) GetMovieById(ctx context.Context, in *GetMovieByIdReq, opts ...grpc.CallOption) (*GetMovieByIdRes, error) {
	out := new(GetMovieByIdRes)
	err := c.cc.Invoke(ctx, "/movieservice.MovieService/GetMovieById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) UpdateMovie(ctx context.Context, in *UpdateMovieReq, opts ...grpc.CallOption) (*UpdateMovieRes, error) {
	out := new(UpdateMovieRes)
	err := c.cc.Invoke(ctx, "/movieservice.MovieService/UpdateMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) DeleteMovie(ctx context.Context, in *DeleteMovieReq, opts ...grpc.CallOption) (*DeleteMovieRes, error) {
	out := new(DeleteMovieRes)
	err := c.cc.Invoke(ctx, "/movieservice.MovieService/DeleteMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) RemoveMovie(ctx context.Context, in *RemoveMovieReq, opts ...grpc.CallOption) (*RemoveMovieRes, error) {
	out := new(RemoveMovieRes)
	err := c.cc.Invoke(ctx, "/movieservice.MovieService/RemoveMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) GetAllMovies(ctx context.Context, in *GetAllMoviesReq, opts ...grpc.CallOption) (*GetAllMoviesRes, error) {
	out := new(GetAllMoviesRes)
	err := c.cc.Invoke(ctx, "/movieservice.MovieService/GetAllMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
// All implementations must embed UnimplementedMovieServiceServer
// for forward compatibility
type MovieServiceServer interface {
	AddMovie(context.Context, *AddMovieReq) (*AddMovieRes, error)
	GetMovieById(context.Context, *GetMovieByIdReq) (*GetMovieByIdRes, error)
	UpdateMovie(context.Context, *UpdateMovieReq) (*UpdateMovieRes, error)
	DeleteMovie(context.Context, *DeleteMovieReq) (*DeleteMovieRes, error)
	RemoveMovie(context.Context, *RemoveMovieReq) (*RemoveMovieRes, error)
	GetAllMovies(context.Context, *GetAllMoviesReq) (*GetAllMoviesRes, error)
	mustEmbedUnimplementedMovieServiceServer()
}

// UnimplementedMovieServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (UnimplementedMovieServiceServer) AddMovie(context.Context, *AddMovieReq) (*AddMovieRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovie not implemented")
}
func (UnimplementedMovieServiceServer) GetMovieById(context.Context, *GetMovieByIdReq) (*GetMovieByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieById not implemented")
}
func (UnimplementedMovieServiceServer) UpdateMovie(context.Context, *UpdateMovieReq) (*UpdateMovieRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovie not implemented")
}
func (UnimplementedMovieServiceServer) DeleteMovie(context.Context, *DeleteMovieReq) (*DeleteMovieRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}
func (UnimplementedMovieServiceServer) RemoveMovie(context.Context, *RemoveMovieReq) (*RemoveMovieRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMovie not implemented")
}
func (UnimplementedMovieServiceServer) GetAllMovies(context.Context, *GetAllMoviesReq) (*GetAllMoviesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMovies not implemented")
}
func (UnimplementedMovieServiceServer) mustEmbedUnimplementedMovieServiceServer() {}

// UnsafeMovieServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieServiceServer will
// result in compilation errors.
type UnsafeMovieServiceServer interface {
	mustEmbedUnimplementedMovieServiceServer()
}

func RegisterMovieServiceServer(s grpc.ServiceRegistrar, srv MovieServiceServer) {
	s.RegisterService(&MovieService_ServiceDesc, srv)
}

func _MovieService_AddMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMovieReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).AddMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movieservice.MovieService/AddMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).AddMovie(ctx, req.(*AddMovieReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_GetMovieById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovieById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movieservice.MovieService/GetMovieById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovieById(ctx, req.(*GetMovieByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_UpdateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMovieReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).UpdateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movieservice.MovieService/UpdateMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).UpdateMovie(ctx, req.(*UpdateMovieReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMovieReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movieservice.MovieService/DeleteMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).DeleteMovie(ctx, req.(*DeleteMovieReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_RemoveMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMovieReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).RemoveMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movieservice.MovieService/RemoveMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).RemoveMovie(ctx, req.(*RemoveMovieReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_GetAllMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMoviesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetAllMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movieservice.MovieService/GetAllMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetAllMovies(ctx, req.(*GetAllMoviesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieService_ServiceDesc is the grpc.ServiceDesc for MovieService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "movieservice.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMovie",
			Handler:    _MovieService_AddMovie_Handler,
		},
		{
			MethodName: "GetMovieById",
			Handler:    _MovieService_GetMovieById_Handler,
		},
		{
			MethodName: "UpdateMovie",
			Handler:    _MovieService_UpdateMovie_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _MovieService_DeleteMovie_Handler,
		},
		{
			MethodName: "RemoveMovie",
			Handler:    _MovieService_RemoveMovie_Handler,
		},
		{
			MethodName: "GetAllMovies",
			Handler:    _MovieService_GetAllMovies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/movie-service/movie.proto",
}

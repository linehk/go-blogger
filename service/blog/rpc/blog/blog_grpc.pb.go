// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: service/blog/rpc/blog.proto

package blog

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
	BlogService_Get_FullMethodName              = "/blog.BlogService/Get"
	BlogService_GetByUrl_FullMethodName         = "/blog.BlogService/GetByUrl"
	BlogService_ListByUser_FullMethodName       = "/blog.BlogService/ListByUser"
	BlogService_GetBlogUserInfos_FullMethodName = "/blog.BlogService/GetBlogUserInfos"
	BlogService_GetPageViews_FullMethodName     = "/blog.BlogService/GetPageViews"
)

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Blog, error)
	GetByUrl(ctx context.Context, in *GetByUrlReq, opts ...grpc.CallOption) (*Blog, error)
	ListByUser(ctx context.Context, in *ListByUserReq, opts ...grpc.CallOption) (*ListByUserResp, error)
	GetBlogUserInfos(ctx context.Context, in *GetBlogUserInfosReq, opts ...grpc.CallOption) (*BlogUserInfos, error)
	GetPageViews(ctx context.Context, in *GetPageViewsReq, opts ...grpc.CallOption) (*PageViews, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Blog, error) {
	out := new(Blog)
	err := c.cc.Invoke(ctx, BlogService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetByUrl(ctx context.Context, in *GetByUrlReq, opts ...grpc.CallOption) (*Blog, error) {
	out := new(Blog)
	err := c.cc.Invoke(ctx, BlogService_GetByUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListByUser(ctx context.Context, in *ListByUserReq, opts ...grpc.CallOption) (*ListByUserResp, error) {
	out := new(ListByUserResp)
	err := c.cc.Invoke(ctx, BlogService_ListByUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetBlogUserInfos(ctx context.Context, in *GetBlogUserInfosReq, opts ...grpc.CallOption) (*BlogUserInfos, error) {
	out := new(BlogUserInfos)
	err := c.cc.Invoke(ctx, BlogService_GetBlogUserInfos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetPageViews(ctx context.Context, in *GetPageViewsReq, opts ...grpc.CallOption) (*PageViews, error) {
	out := new(PageViews)
	err := c.cc.Invoke(ctx, BlogService_GetPageViews_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	Get(context.Context, *GetReq) (*Blog, error)
	GetByUrl(context.Context, *GetByUrlReq) (*Blog, error)
	ListByUser(context.Context, *ListByUserReq) (*ListByUserResp, error)
	GetBlogUserInfos(context.Context, *GetBlogUserInfosReq) (*BlogUserInfos, error)
	GetPageViews(context.Context, *GetPageViewsReq) (*PageViews, error)
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) Get(context.Context, *GetReq) (*Blog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBlogServiceServer) GetByUrl(context.Context, *GetByUrlReq) (*Blog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUrl not implemented")
}
func (UnimplementedBlogServiceServer) ListByUser(context.Context, *ListByUserReq) (*ListByUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByUser not implemented")
}
func (UnimplementedBlogServiceServer) GetBlogUserInfos(context.Context, *GetBlogUserInfosReq) (*BlogUserInfos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlogUserInfos not implemented")
}
func (UnimplementedBlogServiceServer) GetPageViews(context.Context, *GetPageViewsReq) (*PageViews, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPageViews not implemented")
}
func (UnimplementedBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetByUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUrlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetByUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetByUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetByUrl(ctx, req.(*GetByUrlReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ListByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_ListByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ListByUser(ctx, req.(*ListByUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetBlogUserInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlogUserInfosReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetBlogUserInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetBlogUserInfos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetBlogUserInfos(ctx, req.(*GetBlogUserInfosReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetPageViews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPageViewsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetPageViews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetPageViews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetPageViews(ctx, req.(*GetPageViewsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BlogService_Get_Handler,
		},
		{
			MethodName: "GetByUrl",
			Handler:    _BlogService_GetByUrl_Handler,
		},
		{
			MethodName: "ListByUser",
			Handler:    _BlogService_ListByUser_Handler,
		},
		{
			MethodName: "GetBlogUserInfos",
			Handler:    _BlogService_GetBlogUserInfos_Handler,
		},
		{
			MethodName: "GetPageViews",
			Handler:    _BlogService_GetPageViews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/blog/rpc/blog.proto",
}

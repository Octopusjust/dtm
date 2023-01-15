// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: client/dtmgrpc/dtmgpb/dtmgimp.proto

package dtmgpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DtmClient is the client API for Dtm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DtmClient interface {
	NewGid(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DtmGidReply, error)
	Submit(ctx context.Context, in *DtmRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Prepare(ctx context.Context, in *DtmRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Abort(ctx context.Context, in *DtmRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RegisterBranch(ctx context.Context, in *DtmBranchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PrepareWorkflow(ctx context.Context, in *DtmRequest, opts ...grpc.CallOption) (*DtmProgressesReply, error)
	Subscribe(ctx context.Context, in *DtmTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Unsubscribe(ctx context.Context, in *DtmTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteTopic(ctx context.Context, in *DtmTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dtmClient struct {
	cc grpc.ClientConnInterface
}

func NewDtmClient(cc grpc.ClientConnInterface) DtmClient {
	return &dtmClient{cc}
}

func (c *dtmClient) NewGid(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DtmGidReply, error) {
	out := new(DtmGidReply)
	err := c.cc.Invoke(ctx, "/dtmgimp.Dtm/NewGid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dtmClient) Submit(ctx context.Context, in *DtmRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dtmgimp.Dtm/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dtmClient) Prepare(ctx context.Context, in *DtmRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dtmgimp.Dtm/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dtmClient) Abort(ctx context.Context, in *DtmRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dtmgimp.Dtm/Abort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dtmClient) RegisterBranch(ctx context.Context, in *DtmBranchRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dtmgimp.Dtm/RegisterBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dtmClient) PrepareWorkflow(ctx context.Context, in *DtmRequest, opts ...grpc.CallOption) (*DtmProgressesReply, error) {
	out := new(DtmProgressesReply)
	err := c.cc.Invoke(ctx, "/dtmgimp.Dtm/PrepareWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dtmClient) Subscribe(ctx context.Context, in *DtmTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dtmgimp.Dtm/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dtmClient) Unsubscribe(ctx context.Context, in *DtmTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dtmgimp.Dtm/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dtmClient) DeleteTopic(ctx context.Context, in *DtmTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dtmgimp.Dtm/DeleteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DtmServer is the server API for Dtm service.
// All implementations must embed UnimplementedDtmServer
// for forward compatibility
type DtmServer interface {
	NewGid(context.Context, *emptypb.Empty) (*DtmGidReply, error)
	Submit(context.Context, *DtmRequest) (*emptypb.Empty, error)
	Prepare(context.Context, *DtmRequest) (*emptypb.Empty, error)
	Abort(context.Context, *DtmRequest) (*emptypb.Empty, error)
	RegisterBranch(context.Context, *DtmBranchRequest) (*emptypb.Empty, error)
	PrepareWorkflow(context.Context, *DtmRequest) (*DtmProgressesReply, error)
	Subscribe(context.Context, *DtmTopicRequest) (*emptypb.Empty, error)
	Unsubscribe(context.Context, *DtmTopicRequest) (*emptypb.Empty, error)
	DeleteTopic(context.Context, *DtmTopicRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDtmServer()
}

// UnimplementedDtmServer must be embedded to have forward compatible implementations.
type UnimplementedDtmServer struct {
}

func (UnimplementedDtmServer) NewGid(context.Context, *emptypb.Empty) (*DtmGidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewGid not implemented")
}
func (UnimplementedDtmServer) Submit(context.Context, *DtmRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedDtmServer) Prepare(context.Context, *DtmRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (UnimplementedDtmServer) Abort(context.Context, *DtmRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Abort not implemented")
}
func (UnimplementedDtmServer) RegisterBranch(context.Context, *DtmBranchRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterBranch not implemented")
}
func (UnimplementedDtmServer) PrepareWorkflow(context.Context, *DtmRequest) (*DtmProgressesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareWorkflow not implemented")
}
func (UnimplementedDtmServer) Subscribe(context.Context, *DtmTopicRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedDtmServer) Unsubscribe(context.Context, *DtmTopicRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedDtmServer) DeleteTopic(context.Context, *DtmTopicRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedDtmServer) mustEmbedUnimplementedDtmServer() {}

// UnsafeDtmServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DtmServer will
// result in compilation errors.
type UnsafeDtmServer interface {
	mustEmbedUnimplementedDtmServer()
}

func RegisterDtmServer(s grpc.ServiceRegistrar, srv DtmServer) {
	s.RegisterService(&Dtm_ServiceDesc, srv)
}

func _Dtm_NewGid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DtmServer).NewGid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtmgimp.Dtm/NewGid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DtmServer).NewGid(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dtm_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DtmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DtmServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtmgimp.Dtm/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DtmServer).Submit(ctx, req.(*DtmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dtm_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DtmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DtmServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtmgimp.Dtm/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DtmServer).Prepare(ctx, req.(*DtmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dtm_Abort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DtmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DtmServer).Abort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtmgimp.Dtm/Abort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DtmServer).Abort(ctx, req.(*DtmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dtm_RegisterBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DtmBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DtmServer).RegisterBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtmgimp.Dtm/RegisterBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DtmServer).RegisterBranch(ctx, req.(*DtmBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dtm_PrepareWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DtmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DtmServer).PrepareWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtmgimp.Dtm/PrepareWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DtmServer).PrepareWorkflow(ctx, req.(*DtmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dtm_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DtmTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DtmServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtmgimp.Dtm/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DtmServer).Subscribe(ctx, req.(*DtmTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dtm_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DtmTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DtmServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtmgimp.Dtm/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DtmServer).Unsubscribe(ctx, req.(*DtmTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dtm_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DtmTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DtmServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtmgimp.Dtm/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DtmServer).DeleteTopic(ctx, req.(*DtmTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dtm_ServiceDesc is the grpc.ServiceDesc for Dtm service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dtm_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dtmgimp.Dtm",
	HandlerType: (*DtmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewGid",
			Handler:    _Dtm_NewGid_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _Dtm_Submit_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _Dtm_Prepare_Handler,
		},
		{
			MethodName: "Abort",
			Handler:    _Dtm_Abort_Handler,
		},
		{
			MethodName: "RegisterBranch",
			Handler:    _Dtm_RegisterBranch_Handler,
		},
		{
			MethodName: "PrepareWorkflow",
			Handler:    _Dtm_PrepareWorkflow_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _Dtm_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _Dtm_Unsubscribe_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _Dtm_DeleteTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/dtmgrpc/dtmgpb/dtmgimp.proto",
}

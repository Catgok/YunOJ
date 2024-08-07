// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: judge.proto

package judge

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
	JudgeService_Judge_FullMethodName                        = "/judge.JudgeService/Judge"
	JudgeService_OnlineJudge_FullMethodName                  = "/judge.JudgeService/OnlineJudge"
	JudgeService_AddJudgeCases_FullMethodName                = "/judge.JudgeService/AddJudgeCases"
	JudgeService_GetJudgeCasePathsByProblemId_FullMethodName = "/judge.JudgeService/GetJudgeCasePathsByProblemId"
	JudgeService_DeleteJudgeCaseByProblemId_FullMethodName   = "/judge.JudgeService/DeleteJudgeCaseByProblemId"
)

// JudgeServiceClient is the client API for JudgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JudgeServiceClient interface {
	Judge(ctx context.Context, in *JudgeRequest, opts ...grpc.CallOption) (*JudgeResponse, error)
	OnlineJudge(ctx context.Context, in *OnlineJudgeRequest, opts ...grpc.CallOption) (*OnlineJudgeResponse, error)
	AddJudgeCases(ctx context.Context, in *AddJudgeCasesRequest, opts ...grpc.CallOption) (*AddJudgeCasesResponse, error)
	GetJudgeCasePathsByProblemId(ctx context.Context, in *GetJudgeCasePathsRequest, opts ...grpc.CallOption) (*GetJudgeCasePathsResponse, error)
	DeleteJudgeCaseByProblemId(ctx context.Context, in *DeleteJudgeCaseByProblemIdRequest, opts ...grpc.CallOption) (*DeleteJudgeCaseByProblemIdResponse, error)
}

type judgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJudgeServiceClient(cc grpc.ClientConnInterface) JudgeServiceClient {
	return &judgeServiceClient{cc}
}

func (c *judgeServiceClient) Judge(ctx context.Context, in *JudgeRequest, opts ...grpc.CallOption) (*JudgeResponse, error) {
	out := new(JudgeResponse)
	err := c.cc.Invoke(ctx, JudgeService_Judge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeServiceClient) OnlineJudge(ctx context.Context, in *OnlineJudgeRequest, opts ...grpc.CallOption) (*OnlineJudgeResponse, error) {
	out := new(OnlineJudgeResponse)
	err := c.cc.Invoke(ctx, JudgeService_OnlineJudge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeServiceClient) AddJudgeCases(ctx context.Context, in *AddJudgeCasesRequest, opts ...grpc.CallOption) (*AddJudgeCasesResponse, error) {
	out := new(AddJudgeCasesResponse)
	err := c.cc.Invoke(ctx, JudgeService_AddJudgeCases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeServiceClient) GetJudgeCasePathsByProblemId(ctx context.Context, in *GetJudgeCasePathsRequest, opts ...grpc.CallOption) (*GetJudgeCasePathsResponse, error) {
	out := new(GetJudgeCasePathsResponse)
	err := c.cc.Invoke(ctx, JudgeService_GetJudgeCasePathsByProblemId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeServiceClient) DeleteJudgeCaseByProblemId(ctx context.Context, in *DeleteJudgeCaseByProblemIdRequest, opts ...grpc.CallOption) (*DeleteJudgeCaseByProblemIdResponse, error) {
	out := new(DeleteJudgeCaseByProblemIdResponse)
	err := c.cc.Invoke(ctx, JudgeService_DeleteJudgeCaseByProblemId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgeServiceServer is the server API for JudgeService service.
// All implementations must embed UnimplementedJudgeServiceServer
// for forward compatibility
type JudgeServiceServer interface {
	Judge(context.Context, *JudgeRequest) (*JudgeResponse, error)
	OnlineJudge(context.Context, *OnlineJudgeRequest) (*OnlineJudgeResponse, error)
	AddJudgeCases(context.Context, *AddJudgeCasesRequest) (*AddJudgeCasesResponse, error)
	GetJudgeCasePathsByProblemId(context.Context, *GetJudgeCasePathsRequest) (*GetJudgeCasePathsResponse, error)
	DeleteJudgeCaseByProblemId(context.Context, *DeleteJudgeCaseByProblemIdRequest) (*DeleteJudgeCaseByProblemIdResponse, error)
	mustEmbedUnimplementedJudgeServiceServer()
}

// UnimplementedJudgeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJudgeServiceServer struct {
}

func (UnimplementedJudgeServiceServer) Judge(context.Context, *JudgeRequest) (*JudgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Judge not implemented")
}
func (UnimplementedJudgeServiceServer) OnlineJudge(context.Context, *OnlineJudgeRequest) (*OnlineJudgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineJudge not implemented")
}
func (UnimplementedJudgeServiceServer) AddJudgeCases(context.Context, *AddJudgeCasesRequest) (*AddJudgeCasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddJudgeCases not implemented")
}
func (UnimplementedJudgeServiceServer) GetJudgeCasePathsByProblemId(context.Context, *GetJudgeCasePathsRequest) (*GetJudgeCasePathsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJudgeCasePathsByProblemId not implemented")
}
func (UnimplementedJudgeServiceServer) DeleteJudgeCaseByProblemId(context.Context, *DeleteJudgeCaseByProblemIdRequest) (*DeleteJudgeCaseByProblemIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJudgeCaseByProblemId not implemented")
}
func (UnimplementedJudgeServiceServer) mustEmbedUnimplementedJudgeServiceServer() {}

// UnsafeJudgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JudgeServiceServer will
// result in compilation errors.
type UnsafeJudgeServiceServer interface {
	mustEmbedUnimplementedJudgeServiceServer()
}

func RegisterJudgeServiceServer(s grpc.ServiceRegistrar, srv JudgeServiceServer) {
	s.RegisterService(&JudgeService_ServiceDesc, srv)
}

func _JudgeService_Judge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JudgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).Judge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JudgeService_Judge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).Judge(ctx, req.(*JudgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JudgeService_OnlineJudge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineJudgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).OnlineJudge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JudgeService_OnlineJudge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).OnlineJudge(ctx, req.(*OnlineJudgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JudgeService_AddJudgeCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddJudgeCasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).AddJudgeCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JudgeService_AddJudgeCases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).AddJudgeCases(ctx, req.(*AddJudgeCasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JudgeService_GetJudgeCasePathsByProblemId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJudgeCasePathsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).GetJudgeCasePathsByProblemId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JudgeService_GetJudgeCasePathsByProblemId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).GetJudgeCasePathsByProblemId(ctx, req.(*GetJudgeCasePathsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JudgeService_DeleteJudgeCaseByProblemId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJudgeCaseByProblemIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServiceServer).DeleteJudgeCaseByProblemId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JudgeService_DeleteJudgeCaseByProblemId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServiceServer).DeleteJudgeCaseByProblemId(ctx, req.(*DeleteJudgeCaseByProblemIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JudgeService_ServiceDesc is the grpc.ServiceDesc for JudgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JudgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "judge.JudgeService",
	HandlerType: (*JudgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Judge",
			Handler:    _JudgeService_Judge_Handler,
		},
		{
			MethodName: "OnlineJudge",
			Handler:    _JudgeService_OnlineJudge_Handler,
		},
		{
			MethodName: "AddJudgeCases",
			Handler:    _JudgeService_AddJudgeCases_Handler,
		},
		{
			MethodName: "GetJudgeCasePathsByProblemId",
			Handler:    _JudgeService_GetJudgeCasePathsByProblemId_Handler,
		},
		{
			MethodName: "DeleteJudgeCaseByProblemId",
			Handler:    _JudgeService_DeleteJudgeCaseByProblemId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "judge.proto",
}

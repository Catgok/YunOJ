// Code generated by goctl. DO NOT EDIT.
// Source: problem.proto

package problemservice

import (
	"context"

	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateProblemRequest                  = problem.CreateProblemRequest
	CreateProblemResponse                 = problem.CreateProblemResponse
	CreateSubmitRequest                   = problem.CreateSubmitRequest
	CreateSubmitResponse                  = problem.CreateSubmitResponse
	DeleteProblemRequest                  = problem.DeleteProblemRequest
	DeleteProblemResponse                 = problem.DeleteProblemResponse
	GetProblemByIdRequest                 = problem.GetProblemByIdRequest
	GetProblemByIdResponse                = problem.GetProblemByIdResponse
	GetProblemTitleByIdsRequest           = problem.GetProblemTitleByIdsRequest
	GetProblemTitleByIdsResponse          = problem.GetProblemTitleByIdsResponse
	GetProblemsByPageRequest              = problem.GetProblemsByPageRequest
	GetProblemsByPageResponse             = problem.GetProblemsByPageResponse
	GetRecentProblemsRequest              = problem.GetRecentProblemsRequest
	GetRecentProblemsResponse             = problem.GetRecentProblemsResponse
	GetSubmitByIdRequest                  = problem.GetSubmitByIdRequest
	GetSubmitByIdResponse                 = problem.GetSubmitByIdResponse
	GetSubmitByUserIdAndProblemIdRequest  = problem.GetSubmitByUserIdAndProblemIdRequest
	GetSubmitByUserIdAndProblemIdResponse = problem.GetSubmitByUserIdAndProblemIdResponse
	Problem                               = problem.Problem
	ProblemTitleInfo                      = problem.ProblemTitleInfo
	Submit                                = problem.Submit
	UpdateProblemRequest                  = problem.UpdateProblemRequest
	UpdateProblemResponse                 = problem.UpdateProblemResponse
	UpdateSubmitRequest                   = problem.UpdateSubmitRequest
	UpdateSubmitResponse                  = problem.UpdateSubmitResponse

	ProblemService interface {
		GetProblemById(ctx context.Context, in *GetProblemByIdRequest, opts ...grpc.CallOption) (*GetProblemByIdResponse, error)
		GetProblemTitleByIds(ctx context.Context, in *GetProblemTitleByIdsRequest, opts ...grpc.CallOption) (*GetProblemTitleByIdsResponse, error)
		GetRecentProblems(ctx context.Context, in *GetRecentProblemsRequest, opts ...grpc.CallOption) (*GetRecentProblemsResponse, error)
		GetProblemsByPage(ctx context.Context, in *GetProblemsByPageRequest, opts ...grpc.CallOption) (*GetProblemsByPageResponse, error)
		CreateProblem(ctx context.Context, in *CreateProblemRequest, opts ...grpc.CallOption) (*CreateProblemResponse, error)
		UpdateProblem(ctx context.Context, in *UpdateProblemRequest, opts ...grpc.CallOption) (*UpdateProblemResponse, error)
		DeleteProblem(ctx context.Context, in *DeleteProblemRequest, opts ...grpc.CallOption) (*DeleteProblemResponse, error)
		CreateSubmit(ctx context.Context, in *CreateSubmitRequest, opts ...grpc.CallOption) (*CreateSubmitResponse, error)
		GetSubmitById(ctx context.Context, in *GetSubmitByIdRequest, opts ...grpc.CallOption) (*GetSubmitByIdResponse, error)
		GetSubmitByUserIdAndProblemId(ctx context.Context, in *GetSubmitByUserIdAndProblemIdRequest, opts ...grpc.CallOption) (*GetSubmitByUserIdAndProblemIdResponse, error)
		UpdateSubmit(ctx context.Context, in *UpdateSubmitRequest, opts ...grpc.CallOption) (*UpdateSubmitResponse, error)
	}

	defaultProblemService struct {
		cli zrpc.Client
	}
)

func NewProblemService(cli zrpc.Client) ProblemService {
	return &defaultProblemService{
		cli: cli,
	}
}

func (m *defaultProblemService) GetProblemById(ctx context.Context, in *GetProblemByIdRequest, opts ...grpc.CallOption) (*GetProblemByIdResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.GetProblemById(ctx, in, opts...)
}

func (m *defaultProblemService) GetProblemTitleByIds(ctx context.Context, in *GetProblemTitleByIdsRequest, opts ...grpc.CallOption) (*GetProblemTitleByIdsResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.GetProblemTitleByIds(ctx, in, opts...)
}

func (m *defaultProblemService) GetRecentProblems(ctx context.Context, in *GetRecentProblemsRequest, opts ...grpc.CallOption) (*GetRecentProblemsResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.GetRecentProblems(ctx, in, opts...)
}

func (m *defaultProblemService) GetProblemsByPage(ctx context.Context, in *GetProblemsByPageRequest, opts ...grpc.CallOption) (*GetProblemsByPageResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.GetProblemsByPage(ctx, in, opts...)
}

func (m *defaultProblemService) CreateProblem(ctx context.Context, in *CreateProblemRequest, opts ...grpc.CallOption) (*CreateProblemResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.CreateProblem(ctx, in, opts...)
}

func (m *defaultProblemService) UpdateProblem(ctx context.Context, in *UpdateProblemRequest, opts ...grpc.CallOption) (*UpdateProblemResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.UpdateProblem(ctx, in, opts...)
}

func (m *defaultProblemService) DeleteProblem(ctx context.Context, in *DeleteProblemRequest, opts ...grpc.CallOption) (*DeleteProblemResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.DeleteProblem(ctx, in, opts...)
}

func (m *defaultProblemService) CreateSubmit(ctx context.Context, in *CreateSubmitRequest, opts ...grpc.CallOption) (*CreateSubmitResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.CreateSubmit(ctx, in, opts...)
}

func (m *defaultProblemService) GetSubmitById(ctx context.Context, in *GetSubmitByIdRequest, opts ...grpc.CallOption) (*GetSubmitByIdResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.GetSubmitById(ctx, in, opts...)
}

func (m *defaultProblemService) GetSubmitByUserIdAndProblemId(ctx context.Context, in *GetSubmitByUserIdAndProblemIdRequest, opts ...grpc.CallOption) (*GetSubmitByUserIdAndProblemIdResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.GetSubmitByUserIdAndProblemId(ctx, in, opts...)
}

func (m *defaultProblemService) UpdateSubmit(ctx context.Context, in *UpdateSubmitRequest, opts ...grpc.CallOption) (*UpdateSubmitResponse, error) {
	client := problem.NewProblemServiceClient(m.cli.Conn())
	return client.UpdateSubmit(ctx, in, opts...)
}

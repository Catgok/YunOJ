// Code generated by goctl. DO NOT EDIT.
// Source: contest.proto

package contestservice

import (
	"context"

	"YunOJ/services/contest/rpc/contest"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddProblemToContestRequest   = contest.AddProblemToContestRequest
	AddProblemToContestResponse  = contest.AddProblemToContestResponse
	Contest                      = contest.Contest
	ContestRankInfo              = contest.ContestRankInfo
	CreateContestRequest         = contest.CreateContestRequest
	CreateContestResponse        = contest.CreateContestResponse
	GetContestListByPageRequest  = contest.GetContestListByPageRequest
	GetContestListByPageResponse = contest.GetContestListByPageResponse
	GetContestProblemsRequest    = contest.GetContestProblemsRequest
	GetContestProblemsResponse   = contest.GetContestProblemsResponse
	GetContestRankingRequest     = contest.GetContestRankingRequest
	GetContestRankingResponse    = contest.GetContestRankingResponse
	Problem                      = contest.Problem
	SignUpContestRequest         = contest.SignUpContestRequest
	SignUpContestResponse        = contest.SignUpContestResponse
	SubmitAnswerRequest          = contest.SubmitAnswerRequest
	SubmitAnswerResponse         = contest.SubmitAnswerResponse
	UpdateContestRequest         = contest.UpdateContestRequest
	UpdateContestResponse        = contest.UpdateContestResponse

	ContestService interface {
		CreateContest(ctx context.Context, in *CreateContestRequest, opts ...grpc.CallOption) (*CreateContestResponse, error)
		UpdateContest(ctx context.Context, in *UpdateContestRequest, opts ...grpc.CallOption) (*UpdateContestResponse, error)
		GetContestListByPage(ctx context.Context, in *GetContestListByPageRequest, opts ...grpc.CallOption) (*GetContestListByPageResponse, error)
		AddProblemToContest(ctx context.Context, in *AddProblemToContestRequest, opts ...grpc.CallOption) (*AddProblemToContestResponse, error)
		GetContestProblems(ctx context.Context, in *GetContestProblemsRequest, opts ...grpc.CallOption) (*GetContestProblemsResponse, error)
		SubmitAnswer(ctx context.Context, in *SubmitAnswerRequest, opts ...grpc.CallOption) (*SubmitAnswerResponse, error)
		GetContestRanking(ctx context.Context, in *GetContestRankingRequest, opts ...grpc.CallOption) (*GetContestRankingResponse, error)
		SignUpContest(ctx context.Context, in *SignUpContestRequest, opts ...grpc.CallOption) (*SignUpContestResponse, error)
	}

	defaultContestService struct {
		cli zrpc.Client
	}
)

func NewContestService(cli zrpc.Client) ContestService {
	return &defaultContestService{
		cli: cli,
	}
}

func (m *defaultContestService) CreateContest(ctx context.Context, in *CreateContestRequest, opts ...grpc.CallOption) (*CreateContestResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.CreateContest(ctx, in, opts...)
}

func (m *defaultContestService) UpdateContest(ctx context.Context, in *UpdateContestRequest, opts ...grpc.CallOption) (*UpdateContestResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.UpdateContest(ctx, in, opts...)
}

func (m *defaultContestService) GetContestListByPage(ctx context.Context, in *GetContestListByPageRequest, opts ...grpc.CallOption) (*GetContestListByPageResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.GetContestListByPage(ctx, in, opts...)
}

func (m *defaultContestService) AddProblemToContest(ctx context.Context, in *AddProblemToContestRequest, opts ...grpc.CallOption) (*AddProblemToContestResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.AddProblemToContest(ctx, in, opts...)
}

func (m *defaultContestService) GetContestProblems(ctx context.Context, in *GetContestProblemsRequest, opts ...grpc.CallOption) (*GetContestProblemsResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.GetContestProblems(ctx, in, opts...)
}

func (m *defaultContestService) SubmitAnswer(ctx context.Context, in *SubmitAnswerRequest, opts ...grpc.CallOption) (*SubmitAnswerResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.SubmitAnswer(ctx, in, opts...)
}

func (m *defaultContestService) GetContestRanking(ctx context.Context, in *GetContestRankingRequest, opts ...grpc.CallOption) (*GetContestRankingResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.GetContestRanking(ctx, in, opts...)
}

func (m *defaultContestService) SignUpContest(ctx context.Context, in *SignUpContestRequest, opts ...grpc.CallOption) (*SignUpContestResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.SignUpContest(ctx, in, opts...)
}

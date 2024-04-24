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
	ContestBaseInfo              = contest.ContestBaseInfo
	ContestRankInfo              = contest.ContestRankInfo
	CreateContestRequest         = contest.CreateContestRequest
	CreateContestResponse        = contest.CreateContestResponse
	GetContestByIdRequest        = contest.GetContestByIdRequest
	GetContestByIdResponse       = contest.GetContestByIdResponse
	GetContestListByPageRequest  = contest.GetContestListByPageRequest
	GetContestListByPageResponse = contest.GetContestListByPageResponse
	GetContestProblemsRequest    = contest.GetContestProblemsRequest
	GetContestProblemsResponse   = contest.GetContestProblemsResponse
	GetContestRankRequest        = contest.GetContestRankRequest
	GetContestRankResponse       = contest.GetContestRankResponse
	GetRecentContestsRequest     = contest.GetRecentContestsRequest
	GetRecentContestsResponse    = contest.GetRecentContestsResponse
	GetSignUpContestsRequest     = contest.GetSignUpContestsRequest
	GetSignUpContestsResponse    = contest.GetSignUpContestsResponse
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
		GetContestById(ctx context.Context, in *GetContestByIdRequest, opts ...grpc.CallOption) (*GetContestByIdResponse, error)
		GetRecentContests(ctx context.Context, in *GetRecentContestsRequest, opts ...grpc.CallOption) (*GetRecentContestsResponse, error)
		SignUpContest(ctx context.Context, in *SignUpContestRequest, opts ...grpc.CallOption) (*SignUpContestResponse, error)
		GetSignUpContests(ctx context.Context, in *GetSignUpContestsRequest, opts ...grpc.CallOption) (*GetSignUpContestsResponse, error)
		SubmitAnswer(ctx context.Context, in *SubmitAnswerRequest, opts ...grpc.CallOption) (*SubmitAnswerResponse, error)
		GetContestRank(ctx context.Context, in *GetContestRankRequest, opts ...grpc.CallOption) (*GetContestRankResponse, error)
		AddProblemToContest(ctx context.Context, in *AddProblemToContestRequest, opts ...grpc.CallOption) (*AddProblemToContestResponse, error)
		GetContestProblems(ctx context.Context, in *GetContestProblemsRequest, opts ...grpc.CallOption) (*GetContestProblemsResponse, error)
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

func (m *defaultContestService) GetContestById(ctx context.Context, in *GetContestByIdRequest, opts ...grpc.CallOption) (*GetContestByIdResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.GetContestById(ctx, in, opts...)
}

func (m *defaultContestService) GetRecentContests(ctx context.Context, in *GetRecentContestsRequest, opts ...grpc.CallOption) (*GetRecentContestsResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.GetRecentContests(ctx, in, opts...)
}

func (m *defaultContestService) SignUpContest(ctx context.Context, in *SignUpContestRequest, opts ...grpc.CallOption) (*SignUpContestResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.SignUpContest(ctx, in, opts...)
}

func (m *defaultContestService) GetSignUpContests(ctx context.Context, in *GetSignUpContestsRequest, opts ...grpc.CallOption) (*GetSignUpContestsResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.GetSignUpContests(ctx, in, opts...)
}

func (m *defaultContestService) SubmitAnswer(ctx context.Context, in *SubmitAnswerRequest, opts ...grpc.CallOption) (*SubmitAnswerResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.SubmitAnswer(ctx, in, opts...)
}

func (m *defaultContestService) GetContestRank(ctx context.Context, in *GetContestRankRequest, opts ...grpc.CallOption) (*GetContestRankResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.GetContestRank(ctx, in, opts...)
}

func (m *defaultContestService) AddProblemToContest(ctx context.Context, in *AddProblemToContestRequest, opts ...grpc.CallOption) (*AddProblemToContestResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.AddProblemToContest(ctx, in, opts...)
}

func (m *defaultContestService) GetContestProblems(ctx context.Context, in *GetContestProblemsRequest, opts ...grpc.CallOption) (*GetContestProblemsResponse, error) {
	client := contest.NewContestServiceClient(m.cli.Conn())
	return client.GetContestProblems(ctx, in, opts...)
}

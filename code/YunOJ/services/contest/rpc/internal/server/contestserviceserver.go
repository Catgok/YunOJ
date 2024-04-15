// Code generated by goctl. DO NOT EDIT.
// Source: contest.proto

package server

import (
	"context"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/logic"
	"YunOJ/services/contest/rpc/internal/svc"
)

type ContestServiceServer struct {
	svcCtx *svc.ServiceContext
	contest.UnimplementedContestServiceServer
}

func NewContestServiceServer(svcCtx *svc.ServiceContext) *ContestServiceServer {
	return &ContestServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ContestServiceServer) CreateContest(ctx context.Context, in *contest.CreateContestRequest) (*contest.CreateContestResponse, error) {
	l := logic.NewCreateContestLogic(ctx, s.svcCtx)
	return l.CreateContest(in)
}

func (s *ContestServiceServer) UpdateContest(ctx context.Context, in *contest.UpdateContestRequest) (*contest.UpdateContestResponse, error) {
	l := logic.NewUpdateContestLogic(ctx, s.svcCtx)
	return l.UpdateContest(in)
}

func (s *ContestServiceServer) GetContestListByPage(ctx context.Context, in *contest.GetContestListByPageRequest) (*contest.GetContestListByPageResponse, error) {
	l := logic.NewGetContestListByPageLogic(ctx, s.svcCtx)
	return l.GetContestListByPage(in)
}

func (s *ContestServiceServer) GetContestById(ctx context.Context, in *contest.GetContestByIdRequest) (*contest.GetContestByIdResponse, error) {
	l := logic.NewGetContestByIdLogic(ctx, s.svcCtx)
	return l.GetContestById(in)
}

func (s *ContestServiceServer) SignUpContest(ctx context.Context, in *contest.SignUpContestRequest) (*contest.SignUpContestResponse, error) {
	l := logic.NewSignUpContestLogic(ctx, s.svcCtx)
	return l.SignUpContest(in)
}

func (s *ContestServiceServer) GetSignUpContests(ctx context.Context, in *contest.GetSignUpContestsRequest) (*contest.GetSignUpContestsResponse, error) {
	l := logic.NewGetSignUpContestsLogic(ctx, s.svcCtx)
	return l.GetSignUpContests(in)
}

func (s *ContestServiceServer) SubmitAnswer(ctx context.Context, in *contest.SubmitAnswerRequest) (*contest.SubmitAnswerResponse, error) {
	l := logic.NewSubmitAnswerLogic(ctx, s.svcCtx)
	return l.SubmitAnswer(in)
}

func (s *ContestServiceServer) GetContestRank(ctx context.Context, in *contest.GetContestRankRequest) (*contest.GetContestRankResponse, error) {
	l := logic.NewGetContestRankLogic(ctx, s.svcCtx)
	return l.GetContestRank(in)
}

func (s *ContestServiceServer) AddProblemToContest(ctx context.Context, in *contest.AddProblemToContestRequest) (*contest.AddProblemToContestResponse, error) {
	l := logic.NewAddProblemToContestLogic(ctx, s.svcCtx)
	return l.AddProblemToContest(in)
}

func (s *ContestServiceServer) GetContestProblems(ctx context.Context, in *contest.GetContestProblemsRequest) (*contest.GetContestProblemsResponse, error) {
	l := logic.NewGetContestProblemsLogic(ctx, s.svcCtx)
	return l.GetContestProblems(in)
}

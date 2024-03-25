// Code generated by goctl. DO NOT EDIT.
// Source: problem.proto

package server

import (
	problem2 "YunOJ/services/problem/rpc/internal/logic/problem"
	"YunOJ/services/problem/rpc/internal/logic/submit"
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"
)

type ProblemServiceServer struct {
	svcCtx *svc.ServiceContext
	problem.UnimplementedProblemServiceServer
}

func NewProblemServiceServer(svcCtx *svc.ServiceContext) *ProblemServiceServer {
	return &ProblemServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ProblemServiceServer) GetProblemById(ctx context.Context, in *problem.GetProblemByIdRequest) (*problem.GetProblemByIdResponse, error) {
	l := problem2.NewGetProblemByIdLogic(ctx, s.svcCtx)
	return l.GetProblemById(in)
}

func (s *ProblemServiceServer) GetProblemsByPage(ctx context.Context, in *problem.GetProblemsByPageRequest) (*problem.GetProblemsByPageResponse, error) {
	l := problem2.NewGetProblemsByPageLogic(ctx, s.svcCtx)
	return l.GetProblemsByPage(in)
}

func (s *ProblemServiceServer) CreateProblem(ctx context.Context, in *problem.CreateProblemRequest) (*problem.CreateProblemResponse, error) {
	l := problem2.NewCreateProblemLogic(ctx, s.svcCtx)
	return l.CreateProblem(in)
}

func (s *ProblemServiceServer) UpdateProblem(ctx context.Context, in *problem.UpdateProblemRequest) (*problem.UpdateProblemResponse, error) {
	l := problem2.NewUpdateProblemLogic(ctx, s.svcCtx)
	return l.UpdateProblem(in)
}

func (s *ProblemServiceServer) DeleteProblem(ctx context.Context, in *problem.DeleteProblemRequest) (*problem.DeleteProblemResponse, error) {
	l := problem2.NewDeleteProblemLogic(ctx, s.svcCtx)
	return l.DeleteProblem(in)
}

func (s *ProblemServiceServer) CreateSubmit(ctx context.Context, in *problem.CreateSubmitRequest) (*problem.CreateSubmitResponse, error) {
	l := submit.NewCreateSubmitLogic(ctx, s.svcCtx)
	return l.CreateSubmit(in)
}

func (s *ProblemServiceServer) GetSubmitByUserIdAndProblemId(ctx context.Context, in *problem.GetSubmitByUserIdAndProblemIdRequest) (*problem.GetSubmitByUserIdAndProblemIdResponse, error) {
	l := submit.NewGetSubmitByUserIdAndProblemIdLogic(ctx, s.svcCtx)
	return l.GetSubmitByUserIdAndProblemId(in)
}

func (s *ProblemServiceServer) UpdateSubmit(ctx context.Context, in *problem.UpdateSubmitRequest) (*problem.UpdateSubmitResponse, error) {
	l := submit.NewUpdateSubmitLogic(ctx, s.svcCtx)
	return l.UpdateSubmit(in)
}

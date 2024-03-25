package problem

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemLogic {
	return &UpdateProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProblemLogic) UpdateProblem(in *problem.UpdateProblemRequest) (*problem.UpdateProblemResponse, error) {
	resp := &problem.UpdateProblemResponse{
		Code:    0,
		Message: "success",
	}
	res, err := l.svcCtx.ProblemModel.FindOne(l.ctx, in.Problem.ProblemId)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		resp.Success = false
		return resp, nil
	}
	res.Title = in.Problem.Title
	res.TimeLimit = in.Problem.TimeLimit
	res.MemoryLimit = in.Problem.MemoryLimit
	res.Description = in.Problem.Description
	res.HardLevel = in.Problem.HardLevel

	err = l.svcCtx.ProblemModel.Update(l.ctx, res)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		resp.Success = false
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

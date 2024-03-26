package problem

import (
	"YunOJ/services/problem/rpc/problem"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemLogic {
	return &UpdateProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProblemLogic) UpdateProblem(req *types.UpdateProblemRequest) (resp *types.UpdateProblemResponse, err error) {
	resp = &types.UpdateProblemResponse{}
	res, err := l.svcCtx.ProblemRpc.UpdateProblem(l.ctx, &problem.UpdateProblemRequest{
		Problem: &problem.Problem{
			ProblemId:   req.Problem.ProblemId,
			Title:       req.Problem.Title,
			TimeLimit:   req.Problem.TimeLimit,
			MemoryLimit: req.Problem.MemoryLimit,
			Description: req.Problem.Description,
			HardLevel:   req.Problem.ProblemId,
		},
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = res.GetSuccess()
	return resp, nil
}

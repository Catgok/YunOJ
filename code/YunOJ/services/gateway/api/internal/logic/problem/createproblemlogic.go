package problem

import (
	"YunOJ/services/problem/rpc/problem"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProblemLogic {
	return &CreateProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProblemLogic) CreateProblem(req *types.CreateProblemRequest) (resp *types.CreateProblemResponse, err error) {
	resp = &types.CreateProblemResponse{}
	res, err := l.svcCtx.ProblemRpc.CreateProblem(l.ctx, &problem.CreateProblemRequest{
		Problem: &problem.Problem{
			Title:       req.Title,
			TimeLimit:   req.TimeLimit,
			MemoryLimit: req.MemoryLimit,
			Description: req.Description,
			HardLevel:   req.HardLevel,
		},
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	if res.Code != 0 {
		resp.Code, resp.Message = 500, "Internal Server Error"
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = res.GetProblemId()
	return resp, nil
}

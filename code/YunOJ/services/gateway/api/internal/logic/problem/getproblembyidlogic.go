package problem

import (
	"YunOJ/services/problem/rpc/problem"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProblemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemByIdLogic {
	return &GetProblemByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProblemByIdLogic) GetProblemById(req *types.GetProblemByIdRequest) (resp *types.GetProblemByIdResponse, err error) {
	resp = &types.GetProblemByIdResponse{}
	res, err := l.svcCtx.ProblemRpc.GetProblemById(l.ctx, &problem.GetProblemByIdRequest{
		ProblemId: req.ProblemId,
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
	problemData := types.Problem{
		ProblemId:   res.GetProblem().GetProblemId(),
		Title:       res.GetProblem().GetTitle(),
		TimeLimit:   res.GetProblem().GetTimeLimit(),
		MemoryLimit: res.GetProblem().GetMemoryLimit(),
		Description: res.GetProblem().GetDescription(),
		HardLevel:   res.GetProblem().GetHardLevel(),
		SubmitCount: res.GetProblem().GetSubmitCount(),
		PassCount:   res.GetProblem().GetPassCount(),
	}
	resp.Data = problemData
	return resp, nil
}

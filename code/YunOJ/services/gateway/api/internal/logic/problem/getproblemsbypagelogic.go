package problem

import (
	"YunOJ/services/problem/rpc/problem"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemsByPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProblemsByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemsByPageLogic {
	return &GetProblemsByPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProblemsByPageLogic) GetProblemsByPage(req *types.GetProblemsByPageRequest) (resp *types.GetProblemsByPageResponse, err error) {
	resp = &types.GetProblemsByPageResponse{}
	res, err := l.svcCtx.ProblemRpc.GetProblemsByPage(l.ctx, &problem.GetProblemsByPageRequest{
		PageSize:   req.PageSize,
		PageNumber: req.PageNo,
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
	var data []types.Problem
	for _, v := range res.GetProblems() {
		data = append(data, types.Problem{
			ProblemId:   v.GetProblemId(),
			Title:       v.GetTitle(),
			TimeLimit:   v.GetTimeLimit(),
			MemoryLimit: v.GetMemoryLimit(),
			Description: v.GetDescription(),
			HardLevel:   v.GetHardLevel(),
			SubmitCount: v.GetSubmitCount(),
			PassCount:   v.GetPassCount(),
		})
	}
	resp.Data = data
	return resp, nil
}

package problem

import (
	"YunOJ/services/problem/rpc/problem"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecentProblemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecentProblemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecentProblemsLogic {
	return &GetRecentProblemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecentProblemsLogic) GetRecentProblems() (resp *types.GetRecentProblemsResponse, err error) {
	resp = &types.GetRecentProblemsResponse{}
	res, err := l.svcCtx.ProblemRpc.GetRecentProblems(l.ctx, &problem.GetRecentProblemsRequest{})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	if res.Code != 0 {
		resp.Code, resp.Message = 500, "Internal Server Error"
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	var data []types.ProblemTitleInfo
	for _, problemTitleInfo := range res.GetProblemTitleInfo() {
		data = append(data, types.ProblemTitleInfo{
			ProblemId: problemTitleInfo.GetProblemId(),
			Title:     problemTitleInfo.GetTitle(),
		})
	}
	resp.Data = data
	return resp, nil
}

package problem

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecentProblemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecentProblemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecentProblemsLogic {
	return &GetRecentProblemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecentProblemsLogic) GetRecentProblems(in *problem.GetRecentProblemsRequest) (*problem.GetRecentProblemsResponse, error) {
	resp := &problem.GetRecentProblemsResponse{
		Code:    0,
		Message: "success",
	}

	res, err := l.svcCtx.ProblemModel.FindRecentProblems(l.ctx, 10)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	var data []*problem.ProblemTitleInfo
	for _, problemTitleInfo := range res {
		data = append(data, &problem.ProblemTitleInfo{
			ProblemId: problemTitleInfo.ProblemId,
			Title:     problemTitleInfo.Title,
		})
	}
	resp.ProblemTitleInfo = data
	return resp, nil
}

package logic

import (
	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestProblemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContestProblemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestProblemsLogic {
	return &GetContestProblemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContestProblemsLogic) GetContestProblems(in *contest.GetContestProblemsRequest) (*contest.GetContestProblemsResponse, error) {
	resp := &contest.GetContestProblemsResponse{
		Code:    0,
		Message: "success",
	}
	problemIds, err := l.svcCtx.ContestProblemInfoModel.FindByContestId(l.ctx, in.ContestId)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	problemTitleInfosResp, err := l.svcCtx.ProblemRpc.GetProblemTitleByIds(l.ctx, &problem.GetProblemTitleByIdsRequest{ProblemIds: problemIds})
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	var data []*contest.Problem
	for _, v := range problemTitleInfosResp.ProblemTitleInfos {
		data = append(data, &contest.Problem{
			ProblemId: v.GetProblemIds(),
			Title:     v.GetTitle(),
		})

	}
	resp.Problems = data
	return resp, nil

}

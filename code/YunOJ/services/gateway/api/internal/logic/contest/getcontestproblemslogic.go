package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestProblemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContestProblemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestProblemsLogic {
	return &GetContestProblemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContestProblemsLogic) GetContestProblems(req *types.GetContestProblemsRequest) (resp *types.GetContestProblemsResponse, err error) {
	resp = &types.GetContestProblemsResponse{}

	res, err := l.svcCtx.ContestRpc.GetContestProblems(l.ctx, &contest.GetContestProblemsRequest{
		ContestId: req.ContestId,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	var data []types.ContestProblemInfo
	for _, problemInfo := range res.GetProblems() {
		data = append(data, types.ContestProblemInfo{
			ProblemId: problemInfo.GetProblemId(),
			Title:     problemInfo.GetTitle(),
		})
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = data
	return resp, nil
}

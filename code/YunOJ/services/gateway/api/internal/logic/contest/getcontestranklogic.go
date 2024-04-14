package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContestRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestRankLogic {
	return &GetContestRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContestRankLogic) GetContestRank(req *types.GetContestRankRequest) (resp *types.GetContestRankResponse, err error) {
	resp = &types.GetContestRankResponse{}

	res, err := l.svcCtx.ContestRpc.GetContestRank(l.ctx, &contest.GetContestRankRequest{
		ContestId: req.ContestId,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	var data []types.ContestRankInfo
	for _, item := range res.GetRankInfo() {
		data = append(data, types.ContestRankInfo{
			UserId:        item.GetUserId(),
			ProblemId:     item.GetProblemId(),
			TryTimes:      item.GetTryTimes(),
			SubmitTimes:   item.GetSubmitTime(),
			FirstPassTime: item.GetFirstPassTime(),
		})
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = data
	return resp, nil
}

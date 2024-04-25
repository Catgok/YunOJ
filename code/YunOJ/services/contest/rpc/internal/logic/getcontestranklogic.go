package logic

import (
	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContestRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestRankLogic {
	return &GetContestRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContestRankLogic) GetContestRank(in *contest.GetContestRankRequest) (*contest.GetContestRankResponse, error) {
	resp := &contest.GetContestRankResponse{
		Code:    0,
		Message: "success",
	}
	data, err := l.svcCtx.ContestRankInfoModel.FindByContestId(l.ctx, in.GetContestId())
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	var rankInfoList []*contest.ContestRankInfo
	for _, contestRankInfo := range data {
		rankInfoList = append(rankInfoList, &contest.ContestRankInfo{
			UserId:        contestRankInfo.UserId,
			UserName:      contestRankInfo.UserName,
			ProblemId:     contestRankInfo.ProblemId,
			TryTimes:      contestRankInfo.TryTimes,
			SubmitTime:    contestRankInfo.SubmitTimes,
			IsPass:        contestRankInfo.IsPass,
			FirstPassTime: contestRankInfo.FirstPassTime.Int64,
		})
	}
	resp.RankInfo = rankInfoList
	return resp, nil
}

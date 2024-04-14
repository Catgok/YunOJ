package logic

import (
	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestRankingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContestRankingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestRankingLogic {
	return &GetContestRankingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContestRankingLogic) GetContestRanking(in *contest.GetContestRankingRequest) (*contest.GetContestRankingResponse, error) {
	resp := &contest.GetContestRankingResponse{
		Code:    0,
		Message: "success",
	}
	data, err := l.svcCtx.ContestRankInfoModel.FindByContestId(l.ctx, in.GetContestId())
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	var rankInfoList []*contest.ContestRankInfo
	for _, v := range data {
		rankInfoList = append(rankInfoList, &contest.ContestRankInfo{
			UserId:        v.UserId,
			ProblemId:     v.ProblemId,
			TryTimes:      v.TryTimes,
			FirstPassTime: timeToTimestamp(v.FirstPassTime),
		})
	}
	resp.RankInfo = rankInfoList
	return resp, nil
}

func timeToTimestamp(v sql.NullTime) int64 {
	if v.Valid {
		tm := v.Time
		return tm.Unix()
	}
	return 0
}

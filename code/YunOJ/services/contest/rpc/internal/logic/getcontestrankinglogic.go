package logic

import (
	"context"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"

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
	// todo: add your logic here and delete this line

	return &contest.GetContestRankingResponse{}, nil
}

package logic

import (
	"context"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestListByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContestListByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestListByPageLogic {
	return &GetContestListByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContestListByPageLogic) GetContestListByPage(in *contest.GetContestListByPageRequest) (*contest.GetContestListByPageResponse, error) {
	// todo: add your logic here and delete this line

	return &contest.GetContestListByPageResponse{}, nil
}

package logic

import (
	"context"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecentContestsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecentContestsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecentContestsLogic {
	return &GetRecentContestsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecentContestsLogic) GetRecentContests(in *contest.GetRecentContestsRequest) (*contest.GetRecentContestsResponse, error) {
	resp := &contest.GetRecentContestsResponse{
		Code:    0,
		Message: "success",
	}

	res, err := l.svcCtx.ContestInfoModel.FindRecentContests(l.ctx, 10)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	var data []*contest.ContestBaseInfo
	for _, contestBaseInfo := range res {
		data = append(data, &contest.ContestBaseInfo{
			Id:   contestBaseInfo.ContestId,
			Name: contestBaseInfo.ContestName,
		})
	}
	resp.ContestBaseInfos = data
	return resp, nil
}

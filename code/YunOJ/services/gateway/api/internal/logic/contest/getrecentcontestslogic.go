package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecentContestsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecentContestsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecentContestsLogic {
	return &GetRecentContestsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecentContestsLogic) GetRecentContests() (resp *types.GetRecentContestsResponse, err error) {
	resp = &types.GetRecentContestsResponse{}
	res, err := l.svcCtx.ContestRpc.GetRecentContests(l.ctx, &contest.GetRecentContestsRequest{})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	if res.Code != 0 {
		resp.Code, resp.Message = 500, "Internal Server Error"
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	var data []types.ContestBaseInfo
	for _, contestBaseInfo := range res.GetContestBaseInfos() {
		data = append(data, types.ContestBaseInfo{
			Id:   contestBaseInfo.GetId(),
			Name: contestBaseInfo.GetName(),
		})
	}
	resp.Data = data
	return resp, nil
}

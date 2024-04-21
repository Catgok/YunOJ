package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestsByPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContestsByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestsByPageLogic {
	return &GetContestsByPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContestsByPageLogic) GetContestsByPage(req *types.GetContestsByPageRequest) (resp *types.GetContestsByPageResponse, err error) {
	resp = &types.GetContestsByPageResponse{}

	res, err := l.svcCtx.ContestRpc.GetContestListByPage(l.ctx, &contest.GetContestListByPageRequest{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	var data []types.Contest
	for _, contestInfo := range res.GetContests() {
		data = append(data, types.Contest{
			Id:          contestInfo.GetId(),
			Name:        contestInfo.GetName(),
			Description: contestInfo.GetDescription(),
			StartTime:   contestInfo.GetStartTime(),
			EndTime:     contestInfo.GetEndTime(),
		})
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Total = res.GetTotal()
	resp.Data = data
	return resp, nil
}

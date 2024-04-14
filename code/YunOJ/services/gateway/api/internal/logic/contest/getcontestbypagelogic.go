package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestByPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContestByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestByPageLogic {
	return &GetContestByPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContestByPageLogic) GetContestByPage(req *types.GetContestByPageRequest) (resp *types.GetContestByPageResponse, err error) {
	resp = &types.GetContestByPageResponse{}

	res, err := l.svcCtx.ContestRpc.GetContestListByPage(l.ctx, &contest.GetContestListByPageRequest{
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	var data []types.Contest
	for _, item := range res.GetContests() {
		data = append(data, types.Contest{
			Id:          item.GetId(),
			Name:        item.GetName(),
			Description: item.GetDescription(),
			StartTime:   item.GetStartTime(),
			EndTime:     item.GetEndTime(),
		})
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = data
	return resp, nil
}

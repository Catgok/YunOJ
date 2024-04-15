package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContestByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestByIdLogic {
	return &GetContestByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContestByIdLogic) GetContestById(req *types.GetContestByIdRequest) (resp *types.GetContestByIdResponse, err error) {
	resp = &types.GetContestByIdResponse{}

	res, err := l.svcCtx.ContestRpc.GetContestById(l.ctx, &contest.GetContestByIdRequest{
		ContestId: req.ContestId,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	var data types.Contest
	data = types.Contest{
		Id:          res.GetContest().GetId(),
		Name:        res.GetContest().GetName(),
		Description: res.GetContest().GetDescription(),
		StartTime:   res.GetContest().GetStartTime(),
		EndTime:     res.GetContest().GetEndTime(),
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = data
	return resp, nil
}

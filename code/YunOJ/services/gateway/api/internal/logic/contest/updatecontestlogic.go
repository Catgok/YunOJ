package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContestLogic {
	return &UpdateContestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContestLogic) UpdateContest(req *types.UpdateContestRequest) (resp *types.UpdateContestResponse, err error) {
	resp = &types.UpdateContestResponse{}

	userType := l.ctx.Value("user_type").(int64)
	if userType != 1 {
		resp.Code, resp.Message = 403, "Permission denied"
		return resp, nil
	}
	res, err := l.svcCtx.ContestRpc.UpdateContest(l.ctx, &contest.UpdateContestRequest{
		Contest: &contest.Contest{
			Id:          req.Contest.Id,
			Name:        req.Contest.Name,
			Description: req.Contest.Description,
			StartTime:   req.Contest.StartTime,
			EndTime:     req.Contest.EndTime,
		},
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = true
	return resp, nil
}

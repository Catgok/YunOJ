package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContestLogic {
	return &CreateContestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContestLogic) CreateContest(req *types.CreateContestRequest) (resp *types.CreateContestResponse, err error) {
	resp = &types.CreateContestResponse{}

	userType := l.ctx.Value("user_type").(int64)
	if userType != 1 {
		resp.Code, resp.Message = 403, "Permission denied"
		return resp, nil
	}
	res, err := l.svcCtx.ContestRpc.CreateContest(l.ctx, &contest.CreateContestRequest{
		Contest: &contest.Contest{
			Name:        req.Name,
			Description: req.Description,
			StartTime:   req.StartTime,
			EndTime:     req.EndTime,
		},
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = res.GetContestId()
	return resp, nil
}

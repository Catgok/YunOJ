package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProblemLogic {
	return &AddProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProblemLogic) AddProblem(req *types.AddProblemRequest) (resp *types.AddProblemResponse, err error) {
	resp = &types.AddProblemResponse{}

	userType := l.ctx.Value("user_type").(int64)
	if userType != 1 {
		resp.Code, resp.Message = 403, "Permission denied"
		return resp, nil
	}
	res, err := l.svcCtx.ContestRpc.AddProblemToContest(l.ctx, &contest.AddProblemToContestRequest{
		ContestId:  req.ContestId,
		ProblemIds: req.ProblemIds,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = true
	return resp, nil
}

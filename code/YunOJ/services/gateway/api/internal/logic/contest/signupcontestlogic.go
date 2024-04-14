package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpContestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpContestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpContestLogic {
	return &SignUpContestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpContestLogic) SignUpContest(req *types.SignUpContestRequest) (resp *types.SignUpContestResponse, err error) {
	resp = &types.SignUpContestResponse{}
	userId := l.ctx.Value("user_id").(int64)
	res, err := l.svcCtx.ContestRpc.SignUpContest(l.ctx, &contest.SignUpContestRequest{
		ContestId: req.ContestId,
		UserId:    userId,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = true
	return resp, nil
}

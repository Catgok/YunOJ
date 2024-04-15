package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSignUpContestsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSignUpContestsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSignUpContestsLogic {
	return &GetSignUpContestsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSignUpContestsLogic) GetSignUpContests() (resp *types.GetSignUpContestsResponse, err error) {
	resp = &types.GetSignUpContestsResponse{}
	userId := l.ctx.Value("user_id").(int64)
	res, err := l.svcCtx.ContestRpc.GetSignUpContests(l.ctx, &contest.GetSignUpContestsRequest{
		UserId: userId,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = res.GetContestIds()
	return resp, nil
}

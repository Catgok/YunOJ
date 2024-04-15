package logic

import (
	"context"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSignUpContestsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSignUpContestsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSignUpContestsLogic {
	return &GetSignUpContestsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSignUpContestsLogic) GetSignUpContests(in *contest.GetSignUpContestsRequest) (*contest.GetSignUpContestsResponse, error) {
	resp := &contest.GetSignUpContestsResponse{
		Code:    0,
		Message: "success",
	}
	contestIds, err := l.svcCtx.ContestUserInfoModel.FindContestIdsByUserId(l.ctx, in.UserId)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	resp.ContestIds = contestIds
	return resp, nil
}

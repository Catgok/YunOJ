package logic

import (
	"YunOJ/services/contest/model"
	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpContestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpContestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpContestLogic {
	return &SignUpContestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUpContestLogic) SignUpContest(in *contest.SignUpContestRequest) (*contest.SignUpContestResponse, error) {
	resp := &contest.SignUpContestResponse{
		Code:    0,
		Message: "success",
	}
	data := &model.ContestUserInfo{
		UserId:     in.GetUserId(),
		ContestId:  in.GetContestId(),
		JoinStatus: 1,
	}
	_, err := l.svcCtx.ContestUserInfoModel.Insert(l.ctx, data)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

package logic

import (
	"YunOJ/services/contest/model"
	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"
	"context"
	"time"

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
	contestInfo, err := l.svcCtx.ContestInfoModel.FindOne(l.ctx, in.GetContestId())
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	currentTime := time.Now()
	if currentTime.After(contestInfo.EndTime) || currentTime.Before(contestInfo.StartTime) {
		resp.Code, resp.Message = 10301, "The contest is not in progress"
		return resp, nil
	}
	data := &model.ContestUserInfo{
		UserId:     in.GetUserId(),
		ContestId:  in.GetContestId(),
		JoinStatus: 1,
	}
	_, err = l.svcCtx.ContestUserInfoModel.Insert(l.ctx, data)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

package logic

import (
	"YunOJ/services/contest/model"
	"YunOJ/services/problem/rpc/problem"
	"context"
	"time"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitAnswerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitAnswerLogic {
	return &SubmitAnswerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitAnswerLogic) SubmitAnswer(in *contest.SubmitAnswerRequest) (*contest.SubmitAnswerResponse, error) {
	resp := &contest.SubmitAnswerResponse{
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
	createSubmitResp, err := l.svcCtx.ProblemRpc.CreateSubmit(l.ctx, &problem.CreateSubmitRequest{
		UserId:    in.GetUserId(),
		ProblemId: in.GetProblemId(),
		Code:      in.GetCode(),
		Language:  in.GetLanguage(),
	})
	if err != nil {
		resp.Code, resp.Message = 6001, err.Error()
		return resp, nil
	}
	contestSubmitInfo := model.ContestSubmitInfo{
		ContestId: in.GetContestId(),
		SubmitId:  createSubmitResp.GetSubmitId(),
	}
	_, err = l.svcCtx.ContestSubmitInfoModel.Insert(l.ctx, &contestSubmitInfo)
	resp.SubmitId = createSubmitResp.GetSubmitId()
	return resp, nil
}

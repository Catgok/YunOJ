package mq

import (
	"YunOJ/services/contest/model"
	"YunOJ/services/contest/rpc/internal/svc"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

type SubmitChange struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitChange(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitChange {
	return &SubmitChange{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitChange) Consume(key, val string) error {
	var submitChangeMessage userSubmit
	err := json.Unmarshal([]byte(val), &submitChangeMessage)
	if err != nil {
		return err
	}
	submitId, problemId, userId, submitTime := submitChangeMessage.SubmitId, submitChangeMessage.ProblemId, submitChangeMessage.UserId, submitChangeMessage.CreateTime
	// 1. 是否为比赛提交
	contestSubmitInfo, err := l.svcCtx.ContestSubmitInfoModel.FindOneBySubmitId(l.ctx, submitId)
	if err != nil || contestSubmitInfo == nil {
		return nil
	}
	contestId := contestSubmitInfo.ContestId
	// 2. 是否为比赛题目
	contestProblemInfo, err := l.svcCtx.ContestProblemInfoModel.FindOneByContestIdProblemId(l.ctx, contestId, problemId)
	if err != nil || contestProblemInfo == nil {
		return nil
	}
	// 3. 是否为比赛用户
	contestUserInfo, err := l.svcCtx.ContestUserInfoModel.FindOneByContestIdUserId(l.ctx, contestId, userId)
	if err != nil || contestUserInfo == nil {
		return nil
	}

	contestRankInfo, err := l.svcCtx.ContestRankInfoModel.FindOneByContestIdUserIdProblemId(l.ctx, contestId, userId, problemId)
	if errors.Is(err, model.ErrNotFound) || contestRankInfo == nil {
		contestRankInfo = &model.ContestRankInfo{
			ContestId:     contestId,
			UserId:        userId,
			ProblemId:     problemId,
			TryTimes:      0,
			IsPass:        false,
			FirstPassTime: sql.NullTime{},
		}
	}
	// 非AC
	if submitChangeMessage.Result != 0 {
		contestRankInfo.SubmitTimes++
		contestRankInfo.TryTimes++
		_ = l.svcCtx.ContestRankInfoModel.Update(l.ctx, contestRankInfo)
	} else {
		contestRankInfo.SubmitTimes++
		contestRankInfo.IsPass = true
		contestRankInfo.FirstPassTime = sql.NullTime{
			Time:  submitTime,
			Valid: true,
		}
		_ = l.svcCtx.ContestRankInfoModel.Update(l.ctx, contestRankInfo)
	}
	return nil
}

type userSubmit struct {
	SubmitId   int64     `json:"submit_id"`
	UserId     int64     `json:"user_id"`
	ProblemId  int64     `json:"problem_id"`
	Code       string    `json:"code"`
	Status     int64     `json:"status"`
	Language   int64     `json:"language"`
	Result     int64     `json:"result"`
	Time       int64     `json:"time"`
	Memory     int64     `json:"memory"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

package mq

import (
	"YunOJ/services/judge/rpc/internal/logic"
	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/judge"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type ProblemSubmit struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProblemSubmit(ctx context.Context, svcCtx *svc.ServiceContext) *ProblemSubmit {
	return &ProblemSubmit{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProblemSubmit) Consume(key, val string) error {
	logx.Infof("Consume key :%s , val :%s", key, val)
	var data ProblemSubmitMessageBody
	err := json.Unmarshal([]byte(val), &data)
	if err != nil {
		// todo
	}
	judgeLogic := logic.NewJudgeLogic(l.ctx, l.svcCtx)
	_, err = judgeLogic.Judge(&judge.JudgeRequest{
		SubmitId:  data.SubmitId,
		UserId:    data.UserId,
		ProblemId: data.ProblemId,
		Code:      data.Code,
		Language:  data.Language,
	})
	if err != nil {
		return err
	}
	return nil
}

type ProblemSubmitMessageBody struct {
	SubmitId  int64  `json:"submitId"`
	ProblemId int64  `json:"problemId"`
	UserId    int64  `json:"userId"`
	Language  int64  `json:"language"`
	Code      string `json:"code"`
}

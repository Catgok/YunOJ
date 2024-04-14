package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &contest.SubmitAnswerResponse{}, nil
}

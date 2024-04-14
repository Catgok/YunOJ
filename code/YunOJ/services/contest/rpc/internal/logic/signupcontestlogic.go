package logic

import (
	"context"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"

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
	// todo: add your logic here and delete this line

	return &contest.SignUpContestResponse{}, nil
}

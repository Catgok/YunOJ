package logic

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubmitLogic {
	return &CreateSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSubmitLogic) CreateSubmit(in *problem.CreateSubmitRequest) (*problem.CreateSubmitResponse, error) {
	// todo: add your logic here and delete this line

	return &problem.CreateSubmitResponse{}, nil
}

package logic

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProblemLogic {
	return &CreateProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProblemLogic) CreateProblem(in *problem.CreateProblemRequest) (*problem.CreateProblemResponse, error) {
	// todo: add your logic here and delete this line

	return &problem.CreateProblemResponse{}, nil
}

package logic

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemLogic {
	return &UpdateProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProblemLogic) UpdateProblem(in *problem.UpdateProblemRequest) (*problem.UpdateProblemResponse, error) {
	// todo: add your logic here and delete this line

	return &problem.UpdateProblemResponse{}, nil
}

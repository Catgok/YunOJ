package logic

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProblemLogic {
	return &DeleteProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProblemLogic) DeleteProblem(in *problem.DeleteProblemRequest) (*problem.DeleteProblemResponse, error) {
	// todo: add your logic here and delete this line

	return &problem.DeleteProblemResponse{}, nil
}

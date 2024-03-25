package logic

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemByIdLogic {
	return &GetProblemByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemByIdLogic) GetProblemById(in *problem.GetProblemByIdRequest) (*problem.GetProblemByIdResponse, error) {
	// todo: add your logic here and delete this line

	return &problem.GetProblemByIdResponse{}, nil
}

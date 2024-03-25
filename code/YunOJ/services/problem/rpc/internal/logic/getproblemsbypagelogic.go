package logic

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemsByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemsByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemsByPageLogic {
	return &GetProblemsByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemsByPageLogic) GetProblemsByPage(in *problem.GetProblemsByPageRequest) (*problem.GetProblemsByPageResponse, error) {
	// todo: add your logic here and delete this line

	return &problem.GetProblemsByPageResponse{}, nil
}

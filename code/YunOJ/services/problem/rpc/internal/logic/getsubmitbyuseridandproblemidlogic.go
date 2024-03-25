package logic

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubmitByUserIdAndProblemIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubmitByUserIdAndProblemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubmitByUserIdAndProblemIdLogic {
	return &GetSubmitByUserIdAndProblemIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubmitByUserIdAndProblemIdLogic) GetSubmitByUserIdAndProblemId(in *problem.GetSubmitByUserIdAndProblemIdRequest) (*problem.GetSubmitByUserIdAndProblemIdResponse, error) {
	// todo: add your logic here and delete this line

	return &problem.GetSubmitByUserIdAndProblemIdResponse{}, nil
}

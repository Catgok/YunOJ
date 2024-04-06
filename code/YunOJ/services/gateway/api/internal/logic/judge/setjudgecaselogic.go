package judge

import (
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetJudgeCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetJudgeCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetJudgeCaseLogic {
	return &SetJudgeCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetJudgeCaseLogic) SetJudgeCase(req *types.SetJudgeCaseRequest) (resp *types.SetJudgeCaseResponse, err error) {
	resp = &types.SetJudgeCaseResponse{}
	/*	todo
		1.删除problem下所有case
		2.新增所有case
	*/
	return resp, nil
}

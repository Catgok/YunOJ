package logic

import (
	"context"

	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteJudgeCaseByCaseIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteJudgeCaseByCaseIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteJudgeCaseByCaseIdLogic {
	return &DeleteJudgeCaseByCaseIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteJudgeCaseByCaseIdLogic) DeleteJudgeCaseByCaseId(in *judge.DeleteJudgeCaseByCaseIdRequest) (*judge.DeleteJudgeCaseByCaseIdResponse, error) {
	resp := &judge.DeleteJudgeCaseByCaseIdResponse{
		Code:    0,
		Message: "success",
	}
	err := l.svcCtx.JudgeCaseModel.Delete(l.ctx, in.CaseId)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	// todo oss输入输出
	return resp, nil
}

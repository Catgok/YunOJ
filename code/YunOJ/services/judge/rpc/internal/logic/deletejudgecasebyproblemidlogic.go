package logic

import (
	"context"

	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteJudgeCaseByProblemIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteJudgeCaseByProblemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteJudgeCaseByProblemIdLogic {
	return &DeleteJudgeCaseByProblemIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteJudgeCaseByProblemIdLogic) DeleteJudgeCaseByProblemId(in *judge.DeleteJudgeCaseByProblemIdRequest) (*judge.DeleteJudgeCaseByProblemIdResponse, error) {
	resp := &judge.DeleteJudgeCaseByProblemIdResponse{
		Code:    0,
		Message: "success",
	}
	err := l.svcCtx.OssConfig.Delete(in.ProblemId)
	if err != nil {
		resp.Code, resp.Message = 6001, err.Error()
		return resp, nil
	}
	return resp, nil
}

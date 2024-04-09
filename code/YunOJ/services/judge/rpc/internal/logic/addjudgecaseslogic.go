package logic

import (
	"context"

	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddJudgeCasesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddJudgeCasesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddJudgeCasesLogic {
	return &AddJudgeCasesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddJudgeCasesLogic) AddJudgeCases(in *judge.AddJudgeCasesRequest) (*judge.AddJudgeCasesResponse, error) {
	resp := &judge.AddJudgeCasesResponse{
		Code:    0,
		Message: "success",
	}
	err := l.svcCtx.OssConfig.Upload(in.Cases, in.ProblemId)
	if err != nil {
		resp.Code, resp.Message = 6001, err.Error()
		return resp, nil
	}
	return resp, nil
}

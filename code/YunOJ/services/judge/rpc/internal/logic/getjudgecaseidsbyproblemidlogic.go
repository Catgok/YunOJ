package logic

import (
	"context"

	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJudgeCaseIdsByProblemIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJudgeCaseIdsByProblemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJudgeCaseIdsByProblemIdLogic {
	return &GetJudgeCaseIdsByProblemIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetJudgeCaseIdsByProblemIdLogic) GetJudgeCaseIdsByProblemId(in *judge.GetJudgeCaseIdsByProblemIdRequest) (*judge.GetJudgeCaseIdsByProblemIdResponse, error) {
	resp := &judge.GetJudgeCaseIdsByProblemIdResponse{
		Code:    0,
		Message: "success",
	}
	// todo 写sql
	//res, err := l.svcCtx.JudgeCaseModel.FindByProblemId(l.ctx, in.ProblemId)
	//if err != nil {
	//	resp.Code, resp.Message = 5003, err.Error()
	//	return resp, nil
	//}
	//caseIds := make([]int64, 0)
	//for _, v := range res {
	//	caseIds = append(caseIds, v.Id)
	//}
	return resp, nil
}

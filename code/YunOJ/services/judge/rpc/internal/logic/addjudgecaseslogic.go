package logic

import (
	"YunOJ/services/judge/model"
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

	newJudgeCases := make([]model.JudgeCase, 0)
	for _, v := range in.Cases {
		newJudgeCases = append(newJudgeCases, model.JudgeCase{
			ProblemId: in.ProblemId,
			Input:     v.Input,
			Output:    v.Output,
		})
	}
	//res, err := l.svcCtx.JudgeCaseModel.Inserts(l.ctx, newJudgeCases)
	//if err != nil {
	//	resp.Code, resp.Message = 5003, err.Error()
	//	return resp, nil
	//}
	// todo oss处理输入输出
	caseIds := make([]int64, 0)

	resp.CaseIds = caseIds
	return resp, nil
}

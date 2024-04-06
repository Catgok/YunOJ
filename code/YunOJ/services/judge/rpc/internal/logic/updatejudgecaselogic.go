package logic

import (
	"YunOJ/services/judge/model"
	"context"

	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateJudgeCaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateJudgeCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateJudgeCaseLogic {
	return &UpdateJudgeCaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateJudgeCaseLogic) UpdateJudgeCase(in *judge.UpdateJudgeCaseRequest) (*judge.UpdateJudgeCaseResponse, error) {
	resp := &judge.UpdateJudgeCaseResponse{
		Code:    0,
		Message: "success",
	}
	judgeCase := &model.JudgeCase{
		CaseId: in.CaseId,
		Input:  in.Input,
		Output: in.Output,
	}
	err := l.svcCtx.JudgeCaseModel.Update(l.ctx, judgeCase)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	// todo oss更新输入输出
	return resp, nil
}

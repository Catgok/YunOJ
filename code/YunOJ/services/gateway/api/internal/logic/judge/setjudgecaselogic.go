package judge

import (
	"YunOJ/services/judge/rpc/judge"
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

	userType := l.ctx.Value("user_type").(int64)
	if userType != 1 {
		resp.Code, resp.Message = 403, "Permission denied"
		return resp, nil
	}
	_, err = l.svcCtx.JudgeRpc.DeleteJudgeCaseByProblemId(l.ctx, &judge.DeleteJudgeCaseByProblemIdRequest{
		ProblemId: req.ProblemId,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}

	var judgeCase []*judge.JudgeCase
	for _, v := range req.Cases {
		judgeCase = append(judgeCase, &judge.JudgeCase{
			Input:  v.Input,
			Output: v.Output,
		})
	}
	res, err := l.svcCtx.JudgeRpc.AddJudgeCases(l.ctx, &judge.AddJudgeCasesRequest{
		ProblemId: req.ProblemId,
		Cases:     judgeCase,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.Code, res.Message
	return resp, nil
}

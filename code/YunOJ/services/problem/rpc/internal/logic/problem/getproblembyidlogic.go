package problem

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemByIdLogic {
	return &GetProblemByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemByIdLogic) GetProblemById(in *problem.GetProblemByIdRequest) (*problem.GetProblemByIdResponse, error) {
	resp := &problem.GetProblemByIdResponse{
		Code:    0,
		Message: "success",
	}

	res, err := l.svcCtx.ProblemModel.FindOne(l.ctx, in.ProblemId)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	if res.IsDelete == 1 {
		resp.Code, resp.Message = 3001, "题目不存在"
		return resp, nil
	}

	resp.Problem = &problem.Problem{
		ProblemId:   res.ProblemId,
		Title:       res.Title,
		TimeLimit:   res.TimeLimit,
		MemoryLimit: res.MemoryLimit,
		Description: res.Description,
		HardLevel:   res.HardLevel,
		SubmitCount: res.SubmitCount,
		PassCount:   res.PassCount,
	}
	return resp, nil
}

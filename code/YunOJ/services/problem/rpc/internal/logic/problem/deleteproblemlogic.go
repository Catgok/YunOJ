package problem

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProblemLogic {
	return &DeleteProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProblemLogic) DeleteProblem(in *problem.DeleteProblemRequest) (*problem.DeleteProblemResponse, error) {
	resp := &problem.DeleteProblemResponse{
		Code:    0,
		Message: "success",
	}

	res, err := l.svcCtx.ProblemModel.FindOne(l.ctx, in.ProblemId)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		resp.Success = false
		return resp, nil
	}
	res.IsDelete = 1
	err = l.svcCtx.ProblemModel.Update(l.ctx, res)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		resp.Success = false
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

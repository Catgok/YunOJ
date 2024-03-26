package problem

import (
	"YunOJ/services/problem/rpc/problem"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProblemLogic {
	return &DeleteProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProblemLogic) DeleteProblem(req *types.DeleteProblemRequest) (resp *types.DeleteProblemResponse, err error) {
	resp = &types.DeleteProblemResponse{}
	res, err := l.svcCtx.ProblemRpc.DeleteProblem(l.ctx, &problem.DeleteProblemRequest{
		ProblemId: req.ProblemId,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = res.GetSuccess()
	return resp, nil
}

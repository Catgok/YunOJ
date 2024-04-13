package problem

import (
	"YunOJ/services/problem/rpc/problem"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProblemLogic {
	return &UpdateProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProblemLogic) UpdateProblem(req *types.UpdateProblemRequest) (resp *types.UpdateProblemResponse, err error) {
	resp = &types.UpdateProblemResponse{}

	userType := l.ctx.Value("user_type").(int64)
	if userType != 1 {
		resp.Code, resp.Message = 403, "Permission denied"
		return resp, nil
	}
	res, err := l.svcCtx.ProblemRpc.UpdateProblem(l.ctx, &problem.UpdateProblemRequest{
		Problem: &problem.Problem{
			ProblemId:   req.ProblemId,
			Title:       req.Title,
			TimeLimit:   req.TimeLimit,
			MemoryLimit: req.MemoryLimit,
			Description: req.Description,
			HardLevel:   req.HardLevel,
		},
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	if res.Code != 0 {
		resp.Code, resp.Message = 500, "Internal Server Error"
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = res.GetSuccess()
	return resp, nil
}

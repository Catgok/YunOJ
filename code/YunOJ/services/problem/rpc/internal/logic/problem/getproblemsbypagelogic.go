package problem

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemsByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemsByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemsByPageLogic {
	return &GetProblemsByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemsByPageLogic) GetProblemsByPage(in *problem.GetProblemsByPageRequest) (*problem.GetProblemsByPageResponse, error) {
	resp := &problem.GetProblemsByPageResponse{
		Code:    0,
		Message: "success",
	}
	limit, offset := in.PageSize, (in.PageNumber-1)*in.PageSize
	problems, err := l.svcCtx.ProblemModel.FindByPage(l.ctx, offset, limit)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}

	resp.Problems = []*problem.Problem{}
	for _, p := range problems {
		resp.Problems = append(resp.Problems, &problem.Problem{
			ProblemId:   p.ProblemId,
			Title:       p.Title,
			HardLevel:   p.HardLevel,
			Description: p.Description,
			SubmitCount: p.SubmitCount,
			PassCount:   p.PassCount,
		})
	}
	return resp, nil
}

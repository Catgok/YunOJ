package problem

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProblemTitleByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProblemTitleByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProblemTitleByIdsLogic {
	return &GetProblemTitleByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProblemTitleByIdsLogic) GetProblemTitleByIds(in *problem.GetProblemTitleByIdsRequest) (*problem.GetProblemTitleByIdsResponse, error) {
	resp := &problem.GetProblemTitleByIdsResponse{
		Code:    0,
		Message: "success",
	}

	res, err := l.svcCtx.ProblemModel.FindTitlesByIds(l.ctx, in.ProblemIds)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	var data []*problem.ProblemTitleInfo
	for _, v := range res {
		data = append(data, &problem.ProblemTitleInfo{
			ProblemIds: v.ProblemId,
			Title:      v.Title,
		})
	}
	resp.ProblemTitleInfos = data
	return resp, nil
}

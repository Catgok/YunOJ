package logic

import (
	"context"

	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJudgeCasePathsByProblemIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJudgeCasePathsByProblemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJudgeCasePathsByProblemIdLogic {
	return &GetJudgeCasePathsByProblemIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetJudgeCasePathsByProblemIdLogic) GetJudgeCasePathsByProblemId(in *judge.GetJudgeCasePathsRequest) (*judge.GetJudgeCasePathsResponse, error) {
	resp := &judge.GetJudgeCasePathsResponse{
		Code:    0,
		Message: "success",
	}
	paths, err := l.svcCtx.OssConfig.GetDirectoryFilesByProblemId(in.ProblemId)
	if err != nil {
		resp.Code, resp.Message = 6001, err.Error()
		return resp, nil
	}
	resp.Paths = paths
	return resp, nil
}

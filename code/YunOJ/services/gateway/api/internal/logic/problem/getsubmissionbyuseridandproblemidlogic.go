package problem

import (
	"YunOJ/services/problem/rpc/problem"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubmissionByUserIdAndProblemIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubmissionByUserIdAndProblemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubmissionByUserIdAndProblemIdLogic {
	return &GetSubmissionByUserIdAndProblemIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubmissionByUserIdAndProblemIdLogic) GetSubmissionByUserIdAndProblemId(req *types.GetSubmissionByProblemIdRequest) (resp *types.GetSubmissionByProblemIdResponse, err error) {
	resp = &types.GetSubmissionByProblemIdResponse{}
	res, err := l.svcCtx.ProblemRpc.GetSubmitByUserIdAndProblemId(l.ctx, &problem.GetSubmitByUserIdAndProblemIdRequest{
		ProblemId: req.ProblemId,
		UserId:    req.UserId,
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
	var data []types.Submit
	for _, v := range res.GetSubmits() {
		data = append(data, types.Submit{
			SubmitId:  v.GetSubmitId(),
			UserId:    v.GetUserId(),
			ProblemId: v.GetProblemId(),
			Code:      v.GetCode(),
			Status:    v.GetStatus(),
			Language:  v.GetLanguage(),
			Result:    v.GetResult(),
			Time:      v.GetTime(),
			Memory:    v.GetMemory(),
		})
	}
	resp.Data = data
	return resp, nil
}

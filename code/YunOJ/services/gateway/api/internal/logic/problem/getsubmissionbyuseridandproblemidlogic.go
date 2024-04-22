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
	userId := l.ctx.Value("user_id").(int64)
	res, err := l.svcCtx.ProblemRpc.GetSubmitByUserIdAndProblemId(l.ctx, &problem.GetSubmitByUserIdAndProblemIdRequest{
		ProblemId: req.ProblemId,
		UserId:    userId,
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
	for _, submitInfo := range res.GetSubmits() {
		data = append(data, types.Submit{
			SubmitId:   submitInfo.GetSubmitId(),
			UserId:     submitInfo.GetUserId(),
			ProblemId:  submitInfo.GetProblemId(),
			Code:       submitInfo.GetCode(),
			Status:     submitInfo.GetStatus(),
			Language:   submitInfo.GetLanguage(),
			Result:     submitInfo.GetResult(),
			Time:       submitInfo.GetTime(),
			Memory:     submitInfo.GetMemory(),
			CreateTime: submitInfo.GetCreateTime(),
		})
	}
	resp.Data = data
	return resp, nil
}

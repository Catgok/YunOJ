package problem

import (
	"YunOJ/services/problem/rpc/problem"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubmitByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubmitByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubmitByIdLogic {
	return &GetSubmitByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubmitByIdLogic) GetSubmitById(req *types.GetSubmitByIdRequest) (resp *types.GetSubmitByIdResponse, err error) {
	resp = &types.GetSubmitByIdResponse{}
	res, err := l.svcCtx.ProblemRpc.GetSubmitById(l.ctx, &problem.GetSubmitByIdRequest{
		SubmitId: req.SubmitId,
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
	data := types.Submit{
		SubmitId:  res.GetSubmit().GetSubmitId(),
		UserId:    res.GetSubmit().GetUserId(),
		ProblemId: res.GetSubmit().GetProblemId(),
		Code:      res.GetSubmit().GetCode(),
		Status:    res.GetSubmit().GetStatus(),
		Language:  res.GetSubmit().GetLanguage(),
		Result:    res.GetSubmit().GetResult(),
		Time:      res.GetSubmit().GetTime(),
		Memory:    res.GetSubmit().GetMemory(),
	}
	resp.Data = data
	return resp, nil
}

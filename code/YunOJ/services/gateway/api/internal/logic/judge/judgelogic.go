package judge

import (
	"YunOJ/services/judge/rpc/judge"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJudgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeLogic {
	return &JudgeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JudgeLogic) Judge(req *types.JudgeRequest) (resp *types.JudgeResponse, err error) {
	resp = &types.JudgeResponse{}
	res, err := l.svcCtx.JudgeRpc.Judge(l.ctx, &judge.JudgeRequest{
		UserId:    req.UserId,
		ProblemId: req.ProblemId,
		Code:      req.Code,
		Language:  req.Language,
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
	return resp, nil
}

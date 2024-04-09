package judge

import (
	"YunOJ/services/judge/rpc/judge"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnlineJudgeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnlineJudgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnlineJudgeLogic {
	return &OnlineJudgeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnlineJudgeLogic) OnlineJudge(req *types.OnlineJudgeRequest) (resp *types.OnlineJudgeResponse, err error) {
	resp = &types.OnlineJudgeResponse{}
	res, err := l.svcCtx.JudgeRpc.OnlineJudge(l.ctx, &judge.OnlineJudgeRequest{
		Code:     req.Code,
		Input:    req.Input,
		Language: req.Language,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.Code, res.Message
	resp.Data = res.Output
	return
}

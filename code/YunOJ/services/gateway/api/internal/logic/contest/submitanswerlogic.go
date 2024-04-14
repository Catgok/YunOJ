package contest

import (
	"YunOJ/services/contest/rpc/contest"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitAnswerLogic {
	return &SubmitAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitAnswerLogic) SubmitAnswer(req *types.SubmitAnswerRequest) (resp *types.SubmitAnswerResponse, err error) {
	resp = &types.SubmitAnswerResponse{}
	userId := l.ctx.Value("user_id").(int64)
	res, err := l.svcCtx.ContestRpc.SubmitAnswer(l.ctx, &contest.SubmitAnswerRequest{
		ContestId: req.ContestId,
		UserId:    userId,
		ProblemId: req.ProblemId,
		Code:      req.Code,
		Language:  req.Language,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = res.SubmitId
	return resp, nil
}

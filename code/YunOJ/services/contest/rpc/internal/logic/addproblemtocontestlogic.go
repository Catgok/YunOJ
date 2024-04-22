package logic

import (
	"YunOJ/services/contest/model"
	"context"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProblemToContestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProblemToContestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProblemToContestLogic {
	return &AddProblemToContestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProblemToContestLogic) AddProblemToContest(in *contest.AddProblemToContestRequest) (*contest.AddProblemToContestResponse, error) {
	resp := &contest.AddProblemToContestResponse{
		Code:    0,
		Message: "success",
	}
	var data []*model.ContestProblemInfo
	for _, problemId := range in.ProblemIds {
		data = append(data, &model.ContestProblemInfo{
			ContestId: in.ContestId,
			ProblemId: problemId,
		})
	}
	_, err := l.svcCtx.ContestProblemInfoModel.InsertBatch(l.ctx, data)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

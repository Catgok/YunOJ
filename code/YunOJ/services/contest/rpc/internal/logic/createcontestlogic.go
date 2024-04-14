package logic

import (
	"YunOJ/services/contest/model"
	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateContestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContestLogic {
	return &CreateContestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateContestLogic) CreateContest(in *contest.CreateContestRequest) (*contest.CreateContestResponse, error) {
	resp := &contest.CreateContestResponse{
		Code:    0,
		Message: "success",
	}

	data := &model.ContestInfo{
		ContestName: in.GetContest().GetName(),
		Description: in.GetContest().GetDescription(),
		StartTime:   time.Unix(in.GetContest().GetStartTime(), 0),
		EndTime:     time.Unix(in.GetContest().GetEndTime(), 0),
	}
	_, err := l.svcCtx.ContestInfoModel.Insert(l.ctx, data)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

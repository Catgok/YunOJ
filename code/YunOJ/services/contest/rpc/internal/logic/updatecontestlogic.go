package logic

import (
	"YunOJ/services/contest/model"
	"context"
	"time"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateContestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateContestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContestLogic {
	return &UpdateContestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateContestLogic) UpdateContest(in *contest.UpdateContestRequest) (*contest.UpdateContestResponse, error) {
	resp := &contest.UpdateContestResponse{
		Code:    0,
		Message: "success",
	}

	data := &model.ContestInfo{
		ContestId:   in.GetContest().GetId(),
		ContestName: in.GetContest().GetName(),
		Description: in.GetContest().GetDescription(),
		StartTime:   time.Unix(in.GetContest().GetStartTime(), 0),
		EndTime:     time.Unix(in.GetContest().GetEndTime(), 0),
	}
	err := l.svcCtx.ContestInfoModel.Update(l.ctx, data)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

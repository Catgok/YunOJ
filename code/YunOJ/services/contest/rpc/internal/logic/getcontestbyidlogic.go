package logic

import (
	"context"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContestByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestByIdLogic {
	return &GetContestByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContestByIdLogic) GetContestById(in *contest.GetContestByIdRequest) (*contest.GetContestByIdResponse, error) {
	resp := &contest.GetContestByIdResponse{
		Code:    0,
		Message: "success",
	}
	res, err := l.svcCtx.ContestInfoModel.FindOne(l.ctx, in.GetContestId())
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	var data *contest.Contest
	data = &contest.Contest{
		Id:          res.ContestId,
		Name:        res.ContestName,
		Description: res.Description,
		StartTime:   res.StartTime.Unix(),
		EndTime:     res.EndTime.Unix(),
	}
	resp.Contest = data
	return resp, nil
}

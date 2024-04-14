package logic

import (
	"context"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContestListByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContestListByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContestListByPageLogic {
	return &GetContestListByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContestListByPageLogic) GetContestListByPage(in *contest.GetContestListByPageRequest) (*contest.GetContestListByPageResponse, error) {
	resp := &contest.GetContestListByPageResponse{
		Code:    0,
		Message: "success",
	}
	total, err := l.svcCtx.ContestInfoModel.Count(l.ctx)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	resp.Total = total

	limit, offset := in.PageSize, (in.PageNo-1)*in.PageSize
	contests, err := l.svcCtx.ContestInfoModel.FindByPage(l.ctx, offset, limit)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	resp.Contests = []*contest.Contest{}
	for _, c := range contests {
		resp.Contests = append(resp.Contests, &contest.Contest{
			Id:          c.ContestId,
			Name:        c.ContestName,
			Description: c.Description,
			StartTime:   c.StartTime.Unix(),
			EndTime:     c.EndTime.Unix(),
		})
	}
	return resp, nil
}

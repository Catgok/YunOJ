package submit

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubmitLogic {
	return &UpdateSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSubmitLogic) UpdateSubmit(in *problem.UpdateSubmitRequest) (*problem.UpdateSubmitResponse, error) {
	resp := &problem.UpdateSubmitResponse{
		Code:    0,
		Message: "success",
	}
	res, err := l.svcCtx.UserSubmitModel.FindOne(l.ctx, in.Submit.SubmitId)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		resp.Success = false
		return resp, nil
	}

	res.Status = in.Submit.Status
	res.Result = in.Submit.Result
	res.Time = in.Submit.Time
	res.Memory = in.Submit.Memory
	err = l.svcCtx.UserSubmitModel.Update(l.ctx, res)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		resp.Success = false
		return resp, nil
	}
	resp.Success = true
	return resp, nil
}

package submit

import (
	"context"
	"encoding/json"

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

	res.Status = in.Submit.GetStatus()
	res.Result = in.Submit.GetResult()
	res.Time = in.Submit.GetTime()
	res.Memory = in.Submit.GetMemory()
	if in.Submit.GetResult() == 0 {
		err = l.svcCtx.ProblemModel.AddPassCount(l.ctx, res.ProblemId)
		if err != nil {
			logx.Errorf("AddPassCount Error , err :%v", err)
		}
	}
	err = l.svcCtx.UserSubmitModel.Update(l.ctx, res)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		resp.Success = false
		return resp, nil
	}
	resp.Success = true
	message, _ := json.Marshal(res)
	if l.svcCtx.SubmitChangeNoticer.Push(string(message)) != nil {
		logx.Errorf("SubmitChangeNoticer Push Error , err :%v", err)
	}
	return resp, nil
}

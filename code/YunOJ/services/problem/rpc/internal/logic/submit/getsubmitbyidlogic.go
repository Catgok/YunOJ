package submit

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubmitByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubmitByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubmitByIdLogic {
	return &GetSubmitByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubmitByIdLogic) GetSubmitById(in *problem.GetSubmitByIdRequest) (*problem.GetSubmitByIdResponse, error) {
	resp := &problem.GetSubmitByIdResponse{
		Code:    0,
		Message: "success",
	}

	submit, err := l.svcCtx.UserSubmitModel.FindOne(l.ctx, in.SubmitId)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}

	data := &problem.Submit{
		SubmitId:  submit.SubmitId,
		ProblemId: submit.ProblemId,
		UserId:    submit.UserId,
		Language:  submit.Language,
		Code:      submit.Code,
		Status:    submit.Status,
		Result:    submit.Result,
		Time:      submit.Time,
		Memory:    submit.Memory,
	}
	resp.Submit = data
	return resp, nil
}

package submit

import (
	"YunOJ/services/problem/model"
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubmitLogic {
	return &CreateSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSubmitLogic) CreateSubmit(in *problem.CreateSubmitRequest) (*problem.CreateSubmitResponse, error) {
	resp := &problem.CreateSubmitResponse{
		Code:    0,
		Message: "success",
	}

	newUserSubmit := model.NewUserSubmit()
	newUserSubmit.ProblemId = in.ProblemId
	newUserSubmit.UserId = in.UserId
	newUserSubmit.Language = in.Language
	newUserSubmit.Code = in.Code

	res, err := l.svcCtx.UserSubmitModel.Insert(l.ctx, newUserSubmit)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	ld, err := res.LastInsertId()
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	resp.SubmitId = ld
	return resp, nil
}

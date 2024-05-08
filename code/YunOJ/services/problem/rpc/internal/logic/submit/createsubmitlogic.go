package submit

import (
	"YunOJ/services/problem/model"
	"context"
	"encoding/json"

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
	id, err := res.LastInsertId()
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	err = l.svcCtx.ProblemModel.AddSubmitCount(l.ctx, in.ProblemId)
	if err != nil {
		logx.Errorf("AddSubmitCount Error , err :%v", err)
	}
	problemSubmitMessageBody := ProblemSubmitMessageBody{
		SubmitId:  id,
		ProblemId: in.ProblemId,
		UserId:    in.UserId,
		Language:  in.Language,
		Code:      in.Code,
	}
	message, _ := json.Marshal(problemSubmitMessageBody)
	if l.svcCtx.JudgePusher.Push(string(message)) != nil {
		logx.Errorf("JudgePusher Push Error , err :%v", err)
	}
	resp.SubmitId = id
	return resp, nil
}

type ProblemSubmitMessageBody struct {
	SubmitId  int64  `json:"submitId"`
	ProblemId int64  `json:"problemId"`
	UserId    int64  `json:"userId"`
	Language  int64  `json:"language"`
	Code      string `json:"code"`
}

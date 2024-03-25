package problem

import (
	"YunOJ/services/problem/model"
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProblemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProblemLogic {
	return &CreateProblemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProblemLogic) CreateProblem(in *problem.CreateProblemRequest) (*problem.CreateProblemResponse, error) {
	resp := &problem.CreateProblemResponse{
		Code:    0,
		Message: "success",
	}

	newProblem := model.NewProblem()
	newProblem.Title = in.Problem.Title
	newProblem.TimeLimit = in.Problem.TimeLimit
	newProblem.MemoryLimit = in.Problem.MemoryLimit
	newProblem.Description = in.Problem.Description
	newProblem.HardLevel = in.Problem.HardLevel

	res, err := l.svcCtx.ProblemModel.Insert(l.ctx, newProblem)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	lastId, err := res.LastInsertId()
	newProblem, err = l.svcCtx.ProblemModel.FindOne(l.ctx, lastId)

	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}
	resp.ProblemId = newProblem.ProblemId
	//resp.Problem = &problem.Problem{
	//	ProblemId:   newProblem.ProblemId,
	//	Title:       newProblem.Title,
	//	TimeLimit:   newProblem.TimeLimit,
	//	MemoryLimit: newProblem.MemoryLimit,
	//	Description: newProblem.Description,
	//	HardLevel:   newProblem.HardLevel,
	//}
	return resp, nil
}

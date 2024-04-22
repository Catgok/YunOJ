package submit

import (
	"context"

	"YunOJ/services/problem/rpc/internal/svc"
	"YunOJ/services/problem/rpc/problem"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubmitByUserIdAndProblemIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubmitByUserIdAndProblemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubmitByUserIdAndProblemIdLogic {
	return &GetSubmitByUserIdAndProblemIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubmitByUserIdAndProblemIdLogic) GetSubmitByUserIdAndProblemId(in *problem.GetSubmitByUserIdAndProblemIdRequest) (*problem.GetSubmitByUserIdAndProblemIdResponse, error) {
	resp := &problem.GetSubmitByUserIdAndProblemIdResponse{
		Code:    0,
		Message: "success",
	}

	res, err := l.svcCtx.UserSubmitModel.FindByUserIdAndProblemId(l.ctx, in.UserId, in.ProblemId)
	if err != nil {
		resp.Code, resp.Message = 5003, err.Error()
		return resp, nil
	}

	resp.Submits = []*problem.Submit{}
	for _, submit := range res {
		resp.Submits = append(resp.Submits, &problem.Submit{
			SubmitId:   submit.SubmitId,
			ProblemId:  submit.ProblemId,
			UserId:     submit.UserId,
			Language:   submit.Language,
			Code:       submit.Code,
			Status:     submit.Status,
			Result:     submit.Result,
			Time:       submit.Time,
			Memory:     submit.Memory,
			CreateTime: submit.CreateTime.Unix(),
		})
	}
	return resp, nil
}

package logic

import (
	"context"

	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeLogic {
	return &JudgeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JudgeLogic) Judge(in *judge.JudgeRequest) (*judge.JudgeResponse, error) {
	resp := &judge.JudgeResponse{
		Code:    0,
		Message: "success",
	}
	// language,problemId,code := in.Language,in.ProblemId,in.Code
	/* 	todo
	1、查询时间限制 调用 problem.rpc 服务
	2、获取input、output case, 先查库获取url,再oss获取内容
	3、for 每一个case go-judge api评测, 对比输出结果
	4、调用problem.rpc更新评测结果
	*/
	return resp, nil
}

package logic

import (
	"context"

	"YunOJ/services/test/rpc/internal/svc"
	"YunOJ/services/test/rpc/test"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneLogic) GetOne(in *test.GetOneReq) (*test.GetOneResp, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TestModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &test.GetOneResp{
		Info: res.Info,
	}, nil
}

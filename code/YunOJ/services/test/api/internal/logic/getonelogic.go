package logic

import (
	"YunOJ/services/test/rpc/testclient"
	"context"

	"YunOJ/services/test/api/internal/svc"
	"YunOJ/services/test/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOneLogic) GetOne(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TestRpc.GetOne(l.ctx, &testclient.GetOneReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.Response{
		Info: res.Info,
	}, nil
}

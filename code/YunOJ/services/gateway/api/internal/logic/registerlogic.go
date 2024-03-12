package logic

import (
	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"
	"YunOJ/services/user/rpc/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		return resp, nil
	}
	resp = &types.RegisterResponse{
		Code:    res.GetCode(),
		Message: res.GetMessage(),
		Data:    res.GetSuccess(),
	}
	return resp, nil
}

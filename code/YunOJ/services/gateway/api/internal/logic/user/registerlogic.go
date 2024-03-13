package user

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
	resp = &types.RegisterResponse{}
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = res.GetSuccess()
	return resp, nil
}

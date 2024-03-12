package logic

import (
	"YunOJ/services/user/rpc/user"
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.LoginByUserKey(l.ctx, &user.LoginByUserKeyRequest{
		UserKey:  req.UserKey,
		Password: req.Password,
	})
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		return resp, nil
	}

	userInfo := types.User{
		UserId:   res.GetUser().GetUserid(),
		Username: res.GetUser().GetUsername(),
		Email:    res.GetUser().GetEmail(),
		Phone:    res.GetUser().GetPhone(),
		UserType: res.GetUser().GetUserType(),
		Avatar:   res.GetUser().GetAvatar(),
		Status:   res.GetUser().GetStatus(),
	}
	resp = &types.LoginResponse{
		Code:    res.GetCode(),
		Message: res.GetMessage(),
		User:    userInfo,
	}

	return resp, nil
}

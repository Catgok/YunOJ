package user

import (
	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"
	"YunOJ/services/user/rpc/user"
	"context"

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
	resp = &types.LoginResponse{}
	if req.UserKey == "" {
		resp.Code, resp.Message = 400, "入参非法"
		return resp, nil
	}

	res, err := l.svcCtx.UserRpc.LoginByUserKey(l.ctx, &user.LoginByUserKeyRequest{
		UserKey:  req.UserKey,
		Password: req.Password,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
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
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = userInfo

	return resp, nil
}

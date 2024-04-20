package logic

import (
	"YunOJ/services/user/model"
	"YunOJ/services/user/rpc/internal/svc"
	"YunOJ/services/user/rpc/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByUserKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByUserKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByUserKeyLogic {
	return &LoginByUserKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByUserKeyLogic) LoginByUserKey(in *user.LoginByUserKeyRequest) (*user.LoginByUserKeyResponse, error) {
	resp := &user.LoginByUserKeyResponse{
		Code:    0,
		Message: "success",
	}

	// 查询用户是否存在
	res, err := l.loginByUserKey(in)
	if err != nil {
		resp.Code, resp.Message = 1001, "用户不存在"
		return resp, nil
	}
	if in.Password != res.Password {
		resp.Code, resp.Message = 4001, "密码错误"
		return resp, nil
	}
	resp.User = &user.User{
		Userid:   res.Userid,
		Username: res.Username,
		Phone:    res.Phone,
		UserType: res.UserType,
	}
	return resp, nil
}

func (l *LoginByUserKeyLogic) loginByUserKey(in *user.LoginByUserKeyRequest) (*model.UserInfo, error) {
	res, err := l.svcCtx.UserInfoModel.FindOneByPhone(l.ctx, in.UserKey)
	if err == nil {
		return res, nil
	}
	res, err = l.svcCtx.UserInfoModel.FindOneByUsername(l.ctx, in.UserKey)
	if err == nil {
		return res, nil
	}
	return nil, err
}

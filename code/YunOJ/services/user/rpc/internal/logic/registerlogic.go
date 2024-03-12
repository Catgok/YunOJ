package logic

import (
	"YunOJ/services/user/model"
	"context"

	"YunOJ/services/user/rpc/internal/svc"
	"YunOJ/services/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	resp := &user.RegisterResponse{
		Code:    0,
		Message: "success",
		Success: true,
	}
	_, err := l.svcCtx.UserInfoModel.FindOneByPhone(l.ctx, in.Phone)
	if err == nil {
		resp.Code = 1002
		resp.Message = "手机号已注册"
		return resp, nil
	}
	_, err = l.svcCtx.UserInfoModel.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		resp.Code = 1003
		resp.Message = "用户名已存在"
		return resp, nil
	}
	newUserInfo := &model.UserInfo{
		Username: in.GetUsername(),
		Phone:    in.GetPhone(),
		Password: in.GetPassword(),
		Avatar:   "empty_avatar",
	}
	_, err = l.svcCtx.UserInfoModel.Insert(l.ctx, newUserInfo)
	if err != nil {
		resp.Code = 2001
		resp.Message = err.Error()
		resp.Success = false
		return resp, nil
	}
	return resp, nil
}

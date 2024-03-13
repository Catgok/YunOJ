package logic

import (
	"context"

	"YunOJ/services/user/rpc/internal/svc"
	"YunOJ/services/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIdLogic {
	return &GetUserInfoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByIdLogic) GetUserInfoById(in *user.GetUserInfoByIdRequest) (*user.GetUserInfoByIdResponse, error) {
	resp := &user.GetUserInfoByIdResponse{
		Code:    0,
		Message: "success",
	}
	res, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		resp.Code = 1001
		resp.Message = "用户不存在"
		return resp, nil
	}
	resp.User = &user.User{
		Userid:   res.Userid,
		Username: res.Username,
		Email:    res.Email,
		Phone:    res.Phone,
		UserType: res.UserType,
		Avatar:   res.Avatar,
		Status:   res.Status,
	}
	return resp, nil
}

package logic

import (
	"YunOJ/services/user/model"
	"context"
	"errors"
	"google.golang.org/grpc/status"

	"YunOJ/services/user/rpc/internal/svc"
	"YunOJ/services/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPhoneLogic {
	return &LoginByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByPhoneLogic) LoginByPhone(in *user.LoginByPhoneRequest) (*user.LoginByPhoneResponse, error) {
	// todo: add your logic here and delete this line
	// 查询用户是否存在
	res, err := l.svcCtx.UserInfoModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &user.LoginByPhoneResponse{
		Code:    200,
		Message: "success",
		User: &user.User{
			Userid: res.Userid,
		},
	}, nil
}

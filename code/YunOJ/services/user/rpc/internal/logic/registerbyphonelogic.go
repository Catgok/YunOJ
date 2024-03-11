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

type RegisterByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByPhoneLogic {
	return &RegisterByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterByPhoneLogic) RegisterByPhone(in *user.RegisterByPhoneRequest) (*user.RegisterByPhoneResponse, error) {
	// todo: add your logic here and delete this line
	// 判断手机号是否已经注册
	_, err := l.svcCtx.UserInfoModel.FindOneByPhone(l.ctx, in.Phone)
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}

	if errors.Is(err, model.ErrNotFound) {

		newUser := model.UserInfo{
			Phone:    in.Phone,
			Password: in.Password,
		}

		res, err := l.svcCtx.UserInfoModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.Userid, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.RegisterByPhoneResponse{
			Code:    200,
			Message: "success",
			User: &user.User{
				Userid: newUser.Userid,
			},
		}, nil

	}

	return nil, status.Error(500, err.Error())

}

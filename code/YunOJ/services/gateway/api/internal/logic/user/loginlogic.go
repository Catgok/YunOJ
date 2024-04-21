package user

import (
	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"
	"YunOJ/services/user/rpc/user"
	"context"
	"github.com/golang-jwt/jwt/v4"
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
	if res.Code != 0 {
		resp.Code, resp.Message = 500, "Internal Server Error"
		return resp, nil
	}

	userInfo := types.User{
		UserId:   res.GetUser().GetUserid(),
		Username: res.GetUser().GetUsername(),
		Phone:    res.GetUser().GetPhone(),
		UserType: res.GetUser().GetUserType(),
	}
	resp.Code, resp.Message = res.GetCode(), res.GetMessage()
	resp.Data = userInfo
	resp.Utoken, err = l.generateUToken(&userInfo)
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	return resp, nil
}

func (l *LoginLogic) generateUToken(user *types.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	nowTime := jwt.TimeFunc()
	claims["user_id"] = user.UserId
	claims["username"] = user.Username
	claims["user_type"] = user.UserType
	claims["iat"] = nowTime.Unix()
	claims["exp"] = nowTime.Unix() + l.svcCtx.Config.Auth.AccessExpire
	signedToken, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

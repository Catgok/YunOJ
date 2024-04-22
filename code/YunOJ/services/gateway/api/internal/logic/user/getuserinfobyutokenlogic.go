package user

import (
	"context"

	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByUTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByUTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByUTokenLogic {
	return &GetUserInfoByUTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByUTokenLogic) GetUserInfoByUToken(req *types.GetUserInfoByUTokenRequest) (resp *types.GetUserInfoByUTokenResponse, err error) {
	resp = &types.GetUserInfoByUTokenResponse{}
	userType := l.ctx.Value("user_type").(int64)
	if userType != 1 {
		resp.Code, resp.Message = 403, "Permission denied"
		return resp, nil
	}
	userId := l.ctx.Value("user_id").(int64)
	userInfo := types.User{
		UserId:   userId,
		UserType: userType,
	}
	resp.Data = userInfo
	resp.Code, resp.Message = 0, "success"
	return
}

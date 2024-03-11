package svc

import (
	"YunOJ/services/user/model"
	"YunOJ/services/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserInfoModel model.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserInfoModel: model.NewUserInfoModel(conn, c.CacheRedis),
	}
}

package svc

import (
	"YunOJ/services/user/model"
	"YunOJ/services/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ServiceContext struct {
	Config        config.Config
	UserInfoModel model.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	expiryConf := cache.WithExpiry(60 * time.Second)
	return &ServiceContext{
		Config:        c,
		UserInfoModel: model.NewUserInfoModel(conn, c.CacheRedis, expiryConf),
	}
}

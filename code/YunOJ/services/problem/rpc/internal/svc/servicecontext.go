package svc

import (
	"YunOJ/services/problem/model"
	"YunOJ/services/problem/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"time"
)

type ServiceContext struct {
	Config          config.Config
	ProblemModel    model.ProblemModel
	UserSubmitModel model.UserSubmitModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	expiryConf := cache.WithExpiry(time.Second)
	return &ServiceContext{
		Config:          c,
		ProblemModel:    model.NewProblemModel(conn, c.CacheRedis, expiryConf),
		UserSubmitModel: model.NewUserSubmitModel(conn, c.CacheRedis, expiryConf),
	}
}
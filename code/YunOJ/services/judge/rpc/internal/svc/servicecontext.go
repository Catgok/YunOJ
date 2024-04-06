package svc

import (
	"YunOJ/services/judge/model"
	"YunOJ/services/judge/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ServiceContext struct {
	Config         config.Config
	JudgeCaseModel model.JudgeCaseModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	expiryConf := cache.WithExpiry(time.Second)
	return &ServiceContext{
		Config:         c,
		JudgeCaseModel: model.NewJudgeCaseModel(conn, c.CacheRedis, expiryConf),
	}
}

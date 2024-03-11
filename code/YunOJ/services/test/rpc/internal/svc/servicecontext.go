package svc

import (
	"YunOJ/services/test/model"
	"YunOJ/services/test/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	TestModel model.TestModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		TestModel: model.NewTestModel(conn, c.CacheRedis),
	}
}

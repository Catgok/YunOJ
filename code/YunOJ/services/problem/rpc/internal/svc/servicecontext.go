package svc

import (
	"YunOJ/services/problem/model"
	"YunOJ/services/problem/rpc/internal/config"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"time"
)

type ServiceContext struct {
	Config              config.Config
	JudgePusher         *kq.Pusher
	SubmitChangeNoticer *kq.Pusher
	ProblemModel        model.ProblemModel
	UserSubmitModel     model.UserSubmitModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	//expiryConf := cachet.WithExpiry(60 * time.Second)
	fastExpiryConf := cache.WithExpiry(1 * time.Second)
	return &ServiceContext{
		Config:              c,
		JudgePusher:         kq.NewPusher(c.JudgePusherConf.Brokers, c.JudgePusherConf.Topic),
		SubmitChangeNoticer: kq.NewPusher(c.SubmitChangeNoticerConf.Brokers, c.SubmitChangeNoticerConf.Topic),
		ProblemModel:        model.NewProblemModel(conn, c.CacheRedis, fastExpiryConf),
		UserSubmitModel:     model.NewUserSubmitModel(conn, c.CacheRedis, fastExpiryConf),
	}
}

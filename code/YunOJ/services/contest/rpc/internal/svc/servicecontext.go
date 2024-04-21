package svc

import (
	"YunOJ/services/contest/model"
	"YunOJ/services/contest/rpc/internal/config"
	"YunOJ/services/problem/rpc/problemservice"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type ServiceContext struct {
	Config     config.Config
	ProblemRpc problemservice.ProblemService

	ContestInfoModel        model.ContestInfoModel
	ContestProblemInfoModel model.ContestProblemInfoModel
	ContestUserInfoModel    model.ContestUserInfoModel
	ContestRankInfoModel    model.ContestRankInfoModel
	ContestSubmitInfoModel  model.ContestSubmitInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	expiryConf := cache.WithExpiry(60 * time.Second)
	fastExpiryConf := cache.WithExpiry(3 * time.Second)
	return &ServiceContext{
		Config:     c,
		ProblemRpc: problemservice.NewProblemService(zrpc.MustNewClient(c.ProblemRpc)),

		ContestInfoModel:        model.NewContestInfoModel(conn, c.CacheRedis, fastExpiryConf),
		ContestProblemInfoModel: model.NewContestProblemInfoModel(conn, c.CacheRedis, fastExpiryConf),
		ContestUserInfoModel:    model.NewContestUserInfoModel(conn, c.CacheRedis, fastExpiryConf),
		ContestRankInfoModel:    model.NewContestRankInfoModel(conn, c.CacheRedis, expiryConf),
		ContestSubmitInfoModel:  model.NewContestSubmitInfoModel(conn, c.CacheRedis, fastExpiryConf),
	}
}

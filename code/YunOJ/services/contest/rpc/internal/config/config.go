package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	ProblemRpc               zrpc.RpcClientConf
	UserRpc                  zrpc.RpcClientConf
	CacheRedis               cache.CacheConf
	SubmitChangeConsumerConf kq.KqConf
}

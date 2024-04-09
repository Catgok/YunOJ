package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
	CacheRedis cache.CacheConf
}

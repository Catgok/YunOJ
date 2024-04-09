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
	JudgePusherConf struct {
		Brokers []string
		Topic   string
	}
	SubmitChangeNoticerConf struct {
		Brokers []string
		Topic   string
	}
	CacheRedis cache.CacheConf
}

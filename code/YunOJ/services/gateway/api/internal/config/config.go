package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc    zrpc.RpcClientConf
	ProblemRpc zrpc.RpcClientConf
	JudgeRpc   zrpc.RpcClientConf
	ContestRpc zrpc.RpcClientConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
}

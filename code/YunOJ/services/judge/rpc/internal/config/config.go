package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	ProblemRpc zrpc.RpcClientConf
	OssConfig  struct {
		AccessKeyID     string
		AccessKeySecret string
		Endpoint        string
		BucketName      string
		RootPath        string
	}
	KqConsumerConf kq.KqConf
}

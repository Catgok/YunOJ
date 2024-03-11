package svc

import (
	"YunOJ/services/test/api/internal/config"
	"YunOJ/services/test/rpc/testclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	TestRpc testclient.Test
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		TestRpc: testclient.NewTest(zrpc.MustNewClient(c.TestRpc)),
	}
}

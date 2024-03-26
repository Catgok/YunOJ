package svc

import (
	"YunOJ/services/gateway/api/internal/config"
	"YunOJ/services/problem/rpc/problemservice"
	"YunOJ/services/user/rpc/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    userservice.UserService
	ProblemRpc problemservice.ProblemService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		ProblemRpc: problemservice.NewProblemService(zrpc.MustNewClient(c.UserRpc)),
	}
}

package svc

import (
	"YunOJ/services/judge/rpc/internal/config"
	"YunOJ/services/judge/rpc/internal/utils"
	"YunOJ/services/problem/rpc/problemservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProblemRpc problemservice.ProblemService
	OssConfig  *utils.OSSUtil
}

func NewServiceContext(c config.Config) *ServiceContext {
	ossConfig, _ := utils.NewOSSUtil(c.OssConfig.AccessKeyID, c.OssConfig.AccessKeySecret, c.OssConfig.Endpoint, c.OssConfig.BucketName, c.OssConfig.RootPath)
	return &ServiceContext{
		Config:     c,
		ProblemRpc: problemservice.NewProblemService(zrpc.MustNewClient(c.ProblemRpc)),
		OssConfig:  ossConfig,
	}
}

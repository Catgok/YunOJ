package mq

import (
	"YunOJ/services/judge/rpc/internal/config"
	"YunOJ/services/judge/rpc/internal/svc"
	"context"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.JudgeConsumerConf, NewProblemSubmit(ctx, svcContext)),
		//.....
	}

}

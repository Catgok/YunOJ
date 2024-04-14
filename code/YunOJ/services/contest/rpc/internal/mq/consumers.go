package mq

import (
	"YunOJ/services/contest/rpc/internal/config"
	"YunOJ/services/contest/rpc/internal/svc"
	"context"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.SubmitChangeConsumerConf, NewSubmitChange(ctx, svcContext)),
		//.....
	}

}

package main

import (
	"YunOJ/services/judge/rpc/internal/mq"
	"context"
	"flag"
	"fmt"

	"YunOJ/services/judge/rpc/internal/config"
	"YunOJ/services/judge/rpc/internal/server"
	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/judge"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/judge.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		judge.RegisterJudgeServiceServer(grpcServer, server.NewJudgeServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	serviceGroup := service.NewServiceGroup()
	for _, mqConsumer := range mq.Consumers(c, context.Background(), ctx) {
		serviceGroup.Add(mqConsumer)
	}
	go serviceGroup.Start()
	defer serviceGroup.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

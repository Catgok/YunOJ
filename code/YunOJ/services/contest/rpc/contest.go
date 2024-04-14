package main

import (
	"YunOJ/services/contest/rpc/internal/mq"
	"context"
	"flag"
	"fmt"

	"YunOJ/services/contest/rpc/contest"
	"YunOJ/services/contest/rpc/internal/config"
	"YunOJ/services/contest/rpc/internal/server"
	"YunOJ/services/contest/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/contest.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		contest.RegisterContestServiceServer(grpcServer, server.NewContestServiceServer(ctx))

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

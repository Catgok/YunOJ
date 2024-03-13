package main

import (
	"YunOJ/services/gateway/api/internal/config"
	"YunOJ/services/gateway/api/internal/handler"
	"YunOJ/services/gateway/api/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

// go run gateway.go  -f ./etc/gateway.yaml
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//time.Sleep(5 * time.Second)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

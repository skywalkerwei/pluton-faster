package main

import (
	"flag"
	"fmt"

	"github.com/skywalkerwei/pluton-faster/service/order/api/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/order/api/internal/handler"
	"github.com/skywalkerwei/pluton-faster/service/order/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

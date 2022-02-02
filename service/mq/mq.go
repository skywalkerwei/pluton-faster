package main

import (
	"flag"
	"github.com/skywalkerwei/pluton-faster/service/mq/internal/listen"
	"github.com/zeromicro/go-zero/core/service"

	"github.com/skywalkerwei/pluton-faster/service/mq/internal/config"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/mq.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()
}

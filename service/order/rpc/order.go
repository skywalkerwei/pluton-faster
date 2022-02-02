package main

import (
	"flag"
	"fmt"
<<<<<<< HEAD:service/rpc/order/order.go
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/order"
=======
	"github.com/skywalkerwei/pluton-faster/service/order/rpc/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/order/rpc/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/order/rpc/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/order/rpc/order"
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/order/rpc/order.go

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewOrderServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order.RegisterOrderServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
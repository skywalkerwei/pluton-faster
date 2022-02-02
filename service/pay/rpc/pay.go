package main

import (
	"flag"
	"fmt"

<<<<<<< HEAD:service/rpc/pay/pay.go
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/pay"
=======
	"github.com/skywalkerwei/pluton-faster/service/pay/rpc/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/pay/rpc/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/pay/rpc/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/pay/rpc/pay"
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/pay/rpc/pay.go

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/pay.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewPayServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pay.RegisterPayServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

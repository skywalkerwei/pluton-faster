package main

import (
	"flag"
	"fmt"

<<<<<<< HEAD:service/rpc/product/product.go
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/product"
=======
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/product"
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/product/rpc/product.go

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewProductServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		product.RegisterProductServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

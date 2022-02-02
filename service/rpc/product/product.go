package main

import (
	"flag"
	"fmt"

	"github.com/skywalkerwei/pluton-faster/service/rpc/product/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/product"

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

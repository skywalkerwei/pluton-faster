package main

import (
	"flag"
	"fmt"

	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/identity"
	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/identity.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewIdentityServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		identity.RegisterIdentityServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

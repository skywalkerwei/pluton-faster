package main

import (
	"flag"
	"fmt"
	"github.com/skywalkerwei/pluton-faster/common/interceptor/rpcserver"

	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/pay"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
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
	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

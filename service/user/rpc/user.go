package main

import (
	"flag"
	"fmt"

<<<<<<< HEAD:service/rpc/user/user.go
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/user"
=======
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/internal/server"
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/user"
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/user/rpc/user.go

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

package main

import (
	"flag"
	"fmt"
	"github.com/skywalkerwei/pluton-faster/common/resultx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/api/product/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/api/product/internal/handler"
	"github.com/skywalkerwei/pluton-faster/service/api/product/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *resultx.CodeError:
			return http.StatusOK, e.DataInfo()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

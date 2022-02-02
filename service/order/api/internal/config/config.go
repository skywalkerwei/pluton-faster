package config

import (
<<<<<<< HEAD:service/api/wxapp/internal/config/config.go
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
=======
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/order/api/internal/config/config.go
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	OrderRpc   zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
}

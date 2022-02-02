package config

<<<<<<< HEAD:service/api/admin/internal/config/config.go
import "github.com/tal-tech/go-zero/rest"
=======
import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/pay/api/internal/config/config.go

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	PayRpc zrpc.RpcClientConf
}

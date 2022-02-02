package config

<<<<<<< HEAD:service/api/admin/internal/config/config.go
import "github.com/tal-tech/go-zero/rest"
=======
import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/product/api/internal/config/config.go

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	ProductRpc zrpc.RpcClientConf
}

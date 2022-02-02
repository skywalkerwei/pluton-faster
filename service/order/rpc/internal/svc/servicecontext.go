package svc

import (
<<<<<<< HEAD:service/rpc/order/internal/svc/servicecontext.go
	"github.com/skywalkerwei/pluton-faster/service/model"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/productclient"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/userclient"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
<<<<<<< HEAD:service/rpc/order/internal/svc/servicecontext.go
=======
=======
	"github.com/skywalkerwei/pluton-faster/service/order/model"
	"github.com/skywalkerwei/pluton-faster/service/order/rpc/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/productclient"
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/order/rpc/internal/svc/servicecontext.go
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/order/rpc/internal/svc/servicecontext.go
)

type ServiceContext struct {
	Config config.Config

	OrderModel model.OrderModel

	UserRpc    userclient.User
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}

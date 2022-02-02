package svc

import (
<<<<<<< HEAD:service/rpc/cart/internal/svc/servicecontext.go
	"github.com/skywalkerwei/pluton-faster/service/model"
	"github.com/skywalkerwei/pluton-faster/service/rpc/cart/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/productclient"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
=======
	"github.com/skywalkerwei/pluton-faster/service/cart/model"
	"github.com/skywalkerwei/pluton-faster/service/cart/rpc/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/productclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/cart/rpc/internal/svc/servicecontext.go
)

type ServiceContext struct {
	Config     config.Config
	CartModel  model.CartsModel
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		CartModel:  model.NewCartsModel(conn, c.CacheRedis),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}

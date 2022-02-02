package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/cart/model"
	"github.com/skywalkerwei/pluton-faster/service/cart/rpc/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/productclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
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

package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/model"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/productclient"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/userclient"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
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

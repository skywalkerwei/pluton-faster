package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/order/api/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/order/rpc/orderclient"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/productclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   orderclient.Order
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}

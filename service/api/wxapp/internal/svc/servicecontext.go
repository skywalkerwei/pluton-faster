package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/cart/cartclient"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/orderclient"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/payclient"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/productclient"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/userclient"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     userclient.User
	CartRpc     cartclient.Cart
	PayRpc      payclient.Pay
	ProductRpc  productclient.Product
	OrderRpc    orderclient.Order
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CartRpc:     cartclient.NewCart(zrpc.MustNewClient(c.CartRpc)),
		PayRpc:      payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
		OrderRpc:    orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		RedisClient: redis.New(c.RedisConf.Host),
	}
}

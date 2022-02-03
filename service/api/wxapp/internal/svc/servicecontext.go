package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	//UserRpc     userclient.User
	//CartRpc     cartclient.Cart
	//PayRpc      payclient.Pay
	//ProductRpc  productclient.Product
	//OrderRpc    orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.New(c.RedisConf.Host),
		//UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		//CartRpc:     cartclient.NewCart(zrpc.MustNewClient(c.CartRpc)),
		//PayRpc:      payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
		//OrderRpc:    orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		//ProductRpc:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}

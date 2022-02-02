package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/cart/rpc/cartclient"
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/userclient"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     userclient.User
	CartRpc     cartclient.Cart
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CartRpc:     cartclient.NewCart(zrpc.MustNewClient(c.CartRpc)),
		RedisClient: redis.New(c.RedisConf.Host),
	}
}

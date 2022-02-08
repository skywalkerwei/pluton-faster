package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/api/identity/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/identityclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	IdentityRpc identityclient.Identity
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.RedisConf.Host, func(r *redis.Redis) {
			r.Pass = c.RedisConf.Pass
			r.Type = c.RedisConf.Type
		}),
		IdentityRpc: identityclient.NewIdentity(zrpc.MustNewClient(c.IdentityRpc)),
	}
}

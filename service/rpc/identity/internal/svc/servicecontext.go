package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.RedisConf.Host, func(r *redis.Redis) {
			r.Pass = c.RedisConf.Pass
			r.Type = c.RedisConf.Type
		}),
	}
}

//func WithRedisConf(c config.Config) redis.Option {
//	return func(r *redis.Redis) {
//		r.Pass = c.RedisConf.Pass
//	}
//}

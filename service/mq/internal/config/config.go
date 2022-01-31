package config

import (
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type Config struct {
	service.ServiceConf
	Redis redis.RedisConf
}

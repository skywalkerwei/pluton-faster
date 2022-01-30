package config

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Redis redis.RedisConf
}

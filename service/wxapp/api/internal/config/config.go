package config

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	WxMiniConf WxMiniConf
	WxPayConf  WxPayConf
	RedisConf  redis.RedisConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	CartRpc zrpc.RpcClientConf
	UserRpc zrpc.RpcClientConf
}

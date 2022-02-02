package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
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
	CartRpc    zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
	OrderRpc   zrpc.RpcClientConf
	PayRpc     zrpc.RpcClientConf
}

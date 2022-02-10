package config

import (
	"github.com/skywalkerwei/pluton-faster/common/pay"
	"github.com/skywalkerwei/pluton-faster/common/sms"
	"github.com/skywalkerwei/pluton-faster/common/wechat"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	WxMiniConf wechat.WxMiniConf
	WxPayConf  pay.WxConf
	SmsConf    sms.Conf
	RedisConf  redis.RedisConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	//CartRpc    zrpc.RpcClientConf
	//UserRpc    zrpc.RpcClientConf
	//ProductRpc zrpc.RpcClientConf
	//OrderRpc   zrpc.RpcClientConf
	//PayRpc     zrpc.RpcClientConf
}

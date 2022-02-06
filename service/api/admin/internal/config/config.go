package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	SysRpc     zrpc.RpcClientConf
	CartRpc    zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
	OrderRpc   zrpc.RpcClientConf
	PayRpc     zrpc.RpcClientConf
}

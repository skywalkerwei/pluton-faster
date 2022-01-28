package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/userclient"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/config"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}

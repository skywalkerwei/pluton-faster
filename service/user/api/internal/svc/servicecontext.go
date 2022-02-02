package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/user/api/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

//type ServiceContext struct {
//	Config config.Config
//}
//
//func NewServiceContext(c config.Config) *ServiceContext {
//	return &ServiceContext{
//		Config: c,
//	}
//}

type ServiceContext struct {
	Config config.Config

	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}

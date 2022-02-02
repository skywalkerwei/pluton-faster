package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/api/pay/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/payclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PayRpc payclient.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}

package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/pay/api/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/pay/rpc/payclient"
	"github.com/tal-tech/go-zero/zrpc"
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

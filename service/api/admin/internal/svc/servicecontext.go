package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/api/admin/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	SysRpc sysclient.Sys

	//UserRpc    userclient.User
	//CartRpc    cartclient.Cart
	//PayRpc     payclient.Pay
	//ProductRpc productclient.Product
	//OrderRpc   orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SysRpc: sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),

		//UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		//CartRpc:     cartclient.NewCart(zrpc.MustNewClient(c.CartRpc)),
		//PayRpc:      payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
		//OrderRpc:    orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		//ProductRpc:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}

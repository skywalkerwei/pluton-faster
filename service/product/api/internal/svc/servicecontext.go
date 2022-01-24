package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/product/api/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/productclient"
	"github.com/tal-tech/go-zero/zrpc"
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

	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}

package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/orderclient"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/model"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.PayModel
	UserRpc  userclient.User
	OrderRpc orderclient.Order
}

//func NewServiceContext(c config.Config) *ServiceContext {
//	return &ServiceContext{
//		Config: c,
//	}
//}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}

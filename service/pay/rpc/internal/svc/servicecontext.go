package svc

import (
<<<<<<< HEAD:service/rpc/pay/internal/svc/servicecontext.go
	"github.com/skywalkerwei/pluton-faster/service/model"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/orderclient"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/userclient"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
=======
	"github.com/skywalkerwei/pluton-faster/service/order/rpc/orderclient"
	"github.com/skywalkerwei/pluton-faster/service/pay/model"
	"github.com/skywalkerwei/pluton-faster/service/pay/rpc/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/pay/rpc/internal/svc/servicecontext.go
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

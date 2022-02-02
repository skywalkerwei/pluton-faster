package svc

import (
<<<<<<< HEAD:service/rpc/user/internal/svc/servicecontext.go
	"github.com/skywalkerwei/pluton-faster/service/model"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/internal/config"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
=======
	"github.com/skywalkerwei/pluton-faster/service/user/model"
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/user/rpc/internal/svc/servicecontext.go
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

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}

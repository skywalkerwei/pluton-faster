package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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

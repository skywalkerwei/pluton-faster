package svc

import (
<<<<<<< HEAD:service/rpc/product/internal/svc/servicecontext.go
	"github.com/skywalkerwei/pluton-faster/service/model"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/internal/config"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
=======
	"github.com/skywalkerwei/pluton-faster/service/product/model"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/product/rpc/internal/svc/servicecontext.go
)

type ServiceContext struct {
	Config config.Config

	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn, c.CacheRedis),
	}
}

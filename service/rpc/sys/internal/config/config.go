package config

import (
	"github.com/skywalkerwei/pluton-faster/common/gormV2"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	//Mysql struct {
	//	DataSource string
	//}

	GormConf   gormV2.Conf
	CacheRedis cache.CacheConf
}

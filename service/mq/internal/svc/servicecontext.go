package svc

import (
	"github.com/skywalkerwei/pluton-faster/service/mq/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

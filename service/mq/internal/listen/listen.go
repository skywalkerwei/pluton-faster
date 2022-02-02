package listen

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/mq/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/mq/internal/svc"

	"github.com/tal-tech/go-zero/core/service"
)

//返回所有消费者
func Mqs(c config.Config) []service.Service {

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	//kq ：消息队列.
	//services = append(services, KqMqs(c, ctx, svcContext)...)
	//asynq ： 延迟队列、定时任务
	services = append(services, AsynqMqs(c, ctx, svcContext)...)

	return services
}

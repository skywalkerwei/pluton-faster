package listen

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/mq/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/mq/internal/mqs/deferMq"
	"github.com/skywalkerwei/pluton-faster/service/mq/internal/svc"

	"github.com/tal-tech/go-zero/core/service"
)

//asynq
func AsynqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//监听消费流水状态变更
		//.....
		deferMq.NewAsynqTask(ctx, svcContext),
		deferMq.NewSchedulerTask(ctx, svcContext),
	}

}

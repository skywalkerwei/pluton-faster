package listen

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/mq/internal/config"
	"github.com/skywalkerwei/pluton-faster/service/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

//kq
//消息队列
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//监听消费流水状态变更.
		//kq.MustNewQueue(c.SendWxMiniTplMessageConf, wxMessage.NewSendWxMiniSubMessageMq(ctx, svcContext)),
		//.....
	}

}

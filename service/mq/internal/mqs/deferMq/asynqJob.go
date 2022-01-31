package deferMq

import (
	"context"
	"fmt"
	"github.com/skywalkerwei/pluton-faster/service/mq/internal/svc"
	"log"

	"github.com/hibiken/asynq"
)

/**
监听关闭订单
*/
type AsynqJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsynqJob(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqJob {
	return &AsynqJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AsynqJob) Start() {

	fmt.Println("AsynqJob start ")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: l.svcCtx.Config.Redis.Host, Password: l.svcCtx.Config.Redis.Pass},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()

	mux.HandleFunc("msg", l.testJob)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func (l *AsynqJob) Stop() {
	fmt.Println("AsynqJob stop")
}

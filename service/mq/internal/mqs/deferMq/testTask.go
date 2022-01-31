package deferMq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"time"
)

type TestTaskPayload struct {
	Sn string
}

func (l *AsynqJob) testJob(ctx context.Context, t *asynq.Task) error {

	fmt.Println("------TestJob start------")

	var p TestTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(errors.New("解析asynq task payload err"), "closeHomestayOrderStateMqHandler payload err:%v, payLoad:%+v", err, t.Payload())
	}

	fmt.Println(fmt.Printf("{sn: \"%s\"}", p.Sn+time.Now().String()))

	return nil

}

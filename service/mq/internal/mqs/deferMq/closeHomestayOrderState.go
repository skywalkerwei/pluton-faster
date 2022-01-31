package deferMq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

type HomestayOrderCloseTaskPayload struct {
	Sn string
}

func (l *AsynqTask) closeHomestayOrderStateMqHandler(ctx context.Context, t *asynq.Task) error {

	fmt.Println("------HandleMsg start------")

	var p HomestayOrderCloseTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(errors.New("解析asynq task payload err"), "closeHomestayOrderStateMqHandler payload err:%v, payLoad:%+v", err, t.Payload())
	}

	fmt.Println(fmt.Printf("{sn: \"%s\"}", p.Sn))
	return nil

}

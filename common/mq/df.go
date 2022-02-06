package mq

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

//DfSend 立即执行
func DfSend(redsHost string, typeName string, v interface{}) error {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redsHost})
	defer func(client *asynq.Client) {
		_ = client.Close()
	}(client)
	payload, err := json.Marshal(v)
	if err != nil {
		return err
	}
	task := asynq.NewTask(typeName, payload)
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	return err
}

//DfSendAfter 延迟执行
func DfSendAfter(redsHost string, typeName string, v interface{}, t time.Duration) error {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redsHost})
	defer func(client *asynq.Client) {
		_ = client.Close()
	}(client)
	payload, err := json.Marshal(v)
	if err != nil {
		return err
	}
	task := asynq.NewTask(typeName, payload)
	//5*time.Second
	info, err := client.Enqueue(task, asynq.ProcessIn(t))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	return err
}

//DfSendAt 定时执行
func DfSendAt(redsHost string, typeName string, v interface{}, t time.Time) error {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redsHost})
	defer func(client *asynq.Client) {
		_ = client.Close()
	}(client)
	payload, err := json.Marshal(v)
	if err != nil {
		return err
	}
	task := asynq.NewTask(typeName, payload)
	//5*time.Second
	info, err := client.Enqueue(task, asynq.ProcessAt(t))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	return err
}

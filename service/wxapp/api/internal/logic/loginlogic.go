package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/skywalkerwei/pluton-faster/common/jwtx"
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/userclient"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HomestayOrderCloseTaskPayload struct {
	Sn string
}

func (l *LoginLogic) Login(req types.LoginRequest) (resp *types.LoginResponse, err error) {

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: l.svcCtx.Config.RedisConf.Host})
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	payload, err := json.Marshal(HomestayOrderCloseTaskPayload{Sn: "sb001"})
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask("msg", payload)
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	res, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}

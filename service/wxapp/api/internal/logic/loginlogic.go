package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/skywalkerwei/pluton-faster/common/jwtx"
<<<<<<< HEAD:service/api/wxapp/internal/logic/loginlogic.go
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"
	"github.com/skywalkerwei/pluton-faster/service/rpc/user/userclient"
	"github.com/tal-tech/go-zero/core/logx"
<<<<<<< HEAD:service/api/wxapp/internal/logic/loginlogic.go
=======
=======
	"github.com/skywalkerwei/pluton-faster/service/user/rpc/userclient"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/wxapp/api/internal/logic/loginlogic.go
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/wxapp/api/internal/logic/loginlogic.go
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

package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/common/jwtx"
	"github.com/skywalkerwei/pluton-faster/service/api/admin/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/admin/internal/types"
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LoginLogic) Login(req types.LoginRequest) (resp *types.LoginResponse, err error) {

	res, err := l.svcCtx.SysRpc.Login(l.ctx, &sysclient.LoginReq{
		UserName: req.UserName,
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
		Uid:          res.Id,
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil

}

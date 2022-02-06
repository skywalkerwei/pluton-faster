package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/model"
	"time"

	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogAddLogic {
	return &LoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogAddLogic) LoginLogAdd(in *sys.LoginLogAddReq) (*sys.LoginLogAddResp, error) {
	_, err := l.svcCtx.LoginLogModel.Insert(&model.SysLoginLog{
		UserName:  in.UserName,
		Status:    in.Status,
		Ip:        in.Ip,
		CreatedAt: time.Now(),
		CreatedBy: in.CreateBy,
		UpdatedBy: in.CreateBy,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &sys.LoginLogAddResp{}, nil
}

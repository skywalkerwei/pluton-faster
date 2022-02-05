package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigUpdateLogic {
	return &ConfigUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigUpdateLogic) ConfigUpdate(in *sys.ConfigUpdateReq) (*sys.ConfigUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sys.ConfigUpdateResp{}, nil
}

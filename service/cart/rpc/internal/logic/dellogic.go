package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/cart/rpc/cart"
	"github.com/skywalkerwei/pluton-faster/service/cart/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelLogic) Del(in *cart.DelReq) (*cart.DelResp, error) {
	// todo: add your logic here and delete this line

	return &cart.DelResp{}, nil
}

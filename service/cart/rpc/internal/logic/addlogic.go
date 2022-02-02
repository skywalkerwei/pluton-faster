package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/cart/rpc/cart"
	"github.com/skywalkerwei/pluton-faster/service/cart/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *cart.AddReq) (*cart.AddResp, error) {
	// todo: add your logic here and delete this line

	return &cart.AddResp{}, nil
}

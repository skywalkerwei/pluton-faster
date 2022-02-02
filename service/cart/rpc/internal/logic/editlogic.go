package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/cart/rpc/cart"
	"github.com/skywalkerwei/pluton-faster/service/cart/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type EditLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditLogic) Edit(in *cart.EditReq) (*cart.EditResp, error) {
	// todo: add your logic here and delete this line

	return &cart.EditResp{}, nil
}

package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CartsDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartsDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartsDelLogic {
	return CartsDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartsDelLogic) CartsDel(req types.CartsDelRequest) (resp *types.CartsDelResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

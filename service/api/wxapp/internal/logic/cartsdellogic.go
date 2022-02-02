package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
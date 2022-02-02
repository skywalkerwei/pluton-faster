package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartsListLogic {
	return CartsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartsListLogic) CartsList() (resp *types.CartsListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

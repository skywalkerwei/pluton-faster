package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CartsAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartsAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartsAddLogic {
	return CartsAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartsAddLogic) CartsAdd(req types.CartsAddRequest) (resp *types.CartsAddResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

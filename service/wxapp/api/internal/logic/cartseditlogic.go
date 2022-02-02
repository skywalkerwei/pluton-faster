package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartsEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartsEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartsEditLogic {
	return CartsEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartsEditLogic) CartsEdit(req types.CartsEditRequest) (resp *types.CartsEditResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/orderclient"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderRemoveLogic {
	return OrderRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderRemoveLogic) OrderRemove(req types.OrderRemoveRequest) (resp *types.OrderRemoveResponse, err error) {
	_, err = l.svcCtx.OrderRpc.Remove(l.ctx, &orderclient.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.OrderRemoveResponse{}, nil
}

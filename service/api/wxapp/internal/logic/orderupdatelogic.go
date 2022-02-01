package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/orderclient"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderUpdateLogic {
	return OrderUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderUpdateLogic) OrderUpdate(req types.OrderUpdateRequest) (resp *types.OrderUpdateResponse, err error) {
	_, err = l.svcCtx.OrderRpc.Update(l.ctx, &orderclient.UpdateRequest{
		Id:     req.Id,
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.OrderUpdateResponse{}, nil
}

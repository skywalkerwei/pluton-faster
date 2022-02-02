package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/rpc/order/orderclient"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderListLogic {
	return OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req types.OrderListRequest) (resp []*types.OrderListResponse, err error) {
	res, err := l.svcCtx.OrderRpc.List(l.ctx, &orderclient.ListRequest{
		Uid: req.Uid,
	})
	if err != nil {
		return nil, err
	}

	orderList := make([]*types.OrderListResponse, 0)
	for _, item := range res.Data {
		orderList = append(orderList, &types.OrderListResponse{
			Id:     item.Id,
			Uid:    item.Uid,
			Pid:    item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}

	return orderList, nil
}

package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"github.com/skywalkerwei/pluton-faster/service/rpc/cart/cart"
	"github.com/skywalkerwei/pluton-faster/service/rpc/cart/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *cart.ListReq) (*cart.ListResp, error) {
	list, err := l.svcCtx.CartModel.FindAllByUid(in.Uid)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	cartItems := make([]*cart.CartItem, 0)
	for _, item := range list {
		cartItems = append(cartItems, &cart.CartItem{
			Id:   item.Id,
			Uid:  item.Uid,
			Gid:  item.Gid,
			Num:  item.Num,
			Name: item.Name,
		})
	}

	return &cart.ListResp{List: cartItems}, nil
}

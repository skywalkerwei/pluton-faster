package logic

import (
	"context"
	"encoding/json"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"
	"github.com/skywalkerwei/pluton-faster/service/rpc/cart/cartclient"

	"github.com/tal-tech/go-zero/core/logx"
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
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.CartRpc.List(l.ctx, &cartclient.ListReq{
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	cartsItems := make([]types.CartsItem, 0)
	for _, item := range res.List {
		cartsItems = append(cartsItems, types.CartsItem{
			Id:        item.Id,
			GoodsID:   item.Gid,
			GoodsName: item.Name,
			Num:       item.Num,
		})
	}

	return &types.CartsListResponse{List: cartsItems}, nil

}

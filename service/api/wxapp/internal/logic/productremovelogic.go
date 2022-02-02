package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/product"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductRemoveLogic {
	return ProductRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductRemoveLogic) ProductRemove(req types.ProductRemoveRequest) (resp *types.ProductRemoveResponse, err error) {
	_, err = l.svcCtx.ProductRpc.Remove(l.ctx, &product.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.ProductRemoveResponse{}, nil
}

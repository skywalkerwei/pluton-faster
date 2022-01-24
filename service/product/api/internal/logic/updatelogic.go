package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/product"

	"github.com/skywalkerwei/pluton-faster/service/product/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/product/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id:     req.Id,
		Name:   req.Name,
		Desc:   req.Desc,
		Stock:  req.Stock,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil
}

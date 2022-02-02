package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"github.com/skywalkerwei/pluton-faster/service/product/model"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *product.RemoveRequest) (*product.RemoveResponse, error) {
	// 查询产品是否存在
	res, err := l.svcCtx.ProductModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.ProductModel.Delete(res.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.RemoveResponse{}, nil
}

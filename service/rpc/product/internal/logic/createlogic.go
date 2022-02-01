package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/model"
	"google.golang.org/grpc/status"

	"github.com/skywalkerwei/pluton-faster/service/rpc/product/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/rpc/product/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	newProduct := model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
		Status: in.Status,
	}

	res, err := l.svcCtx.ProductModel.Insert(&newProduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newProduct.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.CreateResponse{
		Id: newProduct.Id,
	}, nil
}

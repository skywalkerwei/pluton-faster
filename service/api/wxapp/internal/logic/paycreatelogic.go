package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/pay"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) PayCreateLogic {
	return PayCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayCreateLogic) PayCreate(req types.PayCreateRequest) (resp *types.PayCreateResponse, err error) {
	res, err := l.svcCtx.PayRpc.Create(l.ctx, &pay.CreateRequest{
		Uid:    req.Uid,
		Oid:    req.Oid,
		Amount: req.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &types.PayCreateResponse{
		Id: res.Id,
	}, nil
}

package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/pay"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) PayDetailLogic {
	return PayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayDetailLogic) PayDetail(req types.PayDetailRequest) (resp *types.PayDetailResponse, err error) {
	res, err := l.svcCtx.PayRpc.Detail(l.ctx, &pay.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.PayDetailResponse{
		Id:     req.Id,
		Uid:    res.Uid,
		Oid:    res.Oid,
		Amount: res.Amount,
		Source: res.Source,
		Status: res.Status,
	}, nil
}

package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/rpc/pay/pay"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PayCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) PayCallbackLogic {
	return PayCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayCallbackLogic) PayCallback(req types.PayCallbackRequest) (resp *types.PayCallbackResponse, err error) {
	_, err = l.svcCtx.PayRpc.Callback(l.ctx, &pay.CallbackRequest{
		Id:     req.Id,
		Uid:    req.Uid,
		Oid:    req.Oid,
		Amount: req.Amount,
		Source: req.Source,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.PayCallbackResponse{}, nil
}

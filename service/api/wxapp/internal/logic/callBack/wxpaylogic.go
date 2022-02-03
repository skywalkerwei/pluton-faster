package callBack

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) WxPayLogic {
	return WxPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxPayLogic) WxPay(req types.WxPayCallBackReq) (resp *types.WxPayCallBackResp, err error) {
	// todo: add your logic here and delete this line

	return
}

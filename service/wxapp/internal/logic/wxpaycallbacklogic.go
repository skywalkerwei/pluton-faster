package logic

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxPayCallBackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxPayCallBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) WxPayCallBackLogic {
	return WxPayCallBackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxPayCallBackLogic) WxPayCallBack(req types.WxPayCallBackReq) (resp *types.WxPayCallBackResp, err error) {
	// todo: add your logic here and delete this line

	return
}

package logic

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/types"
	"net/http"

<<<<<<< HEAD:service/api/wxapp/internal/logic/wxpaycallbacklogic.go
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
=======
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/wxapp/api/internal/logic/wxpaycallbacklogic.go
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

//func (l *WxPayCallBackLogic) WxPayCallBack(req types.WxPayCallBackReq) (resp *types.WxPayCallBackResp, err error) {
//	// todo: add your logic here and delete this line
//
//	return
//}

func (l *WxPayCallBackLogic) WxPayCallBack(rw http.ResponseWriter, req *http.Request) (resp *types.WxPayCallBackResp, err error) {
	returnCode := "SUCCESS"
	return &types.WxPayCallBackResp{
		ReturnCode: returnCode,
	}, err

}

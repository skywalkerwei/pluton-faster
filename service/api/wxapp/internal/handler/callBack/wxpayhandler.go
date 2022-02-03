package callBack

import (
	"fmt"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic/callBack"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

func WxPayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.WxPayCallBackReq
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		//
		//l := callBack.NewWxPayLogic(r.Context(), svcCtx)
		//resp, err := l.WxPay(req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
		l := callBack.NewWxPayLogic(r.Context(), svcCtx)
		resp, err := l.WxPay(w, r)

		if err != nil {
			logx.WithContext(r.Context()).Errorf("【API-ERR】CallBack WxPayHandler : %+v ", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		//logx.Infof("ReturnCode : %s ", resp.ReturnCode)
		fmt.Fprint(w.(http.ResponseWriter), resp.ReturnCode)
	}
}

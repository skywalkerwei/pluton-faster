package handler

import (
	"net/http"

<<<<<<< HEAD:service/api/wxapp/internal/handler/wxpaycallbackhandler.go
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
=======
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/wxapp/api/internal/handler/wxpaycallbackhandler.go
)

func WxPayCallBackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.WxPayCallBackReq
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}

		l := logic.NewWxPayCallBackLogic(r.Context(), svcCtx)
		resp, err := l.WxPayCallBack(w, r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

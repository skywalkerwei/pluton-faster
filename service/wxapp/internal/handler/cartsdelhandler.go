package handler

import (
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/wxapp/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CartsDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartsDelRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCartsDelLogic(r.Context(), svcCtx)
		resp, err := l.CartsDel(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

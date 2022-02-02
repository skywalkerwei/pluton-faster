package handler

import (
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/wxapp/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CartsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCartsListLogic(r.Context(), svcCtx)
		resp, err := l.CartsList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

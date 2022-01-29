package handler

import (
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func CartsEditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartsEditRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCartsEditLogic(r.Context(), svcCtx)
		resp, err := l.CartsEdit(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
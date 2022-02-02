package handler

import (
	"net/http"

<<<<<<< HEAD:service/api/wxapp/internal/handler/cartsaddhandler.go
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
=======
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/wxapp/api/internal/handler/cartsaddhandler.go
)

func CartsAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartsAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCartsAddLogic(r.Context(), svcCtx)
		resp, err := l.CartsAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

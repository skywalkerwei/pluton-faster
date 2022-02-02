package handler

import (
	"net/http"

<<<<<<< HEAD:service/api/wxapp/internal/handler/cartsedithandler.go
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
<<<<<<< HEAD:service/api/wxapp/internal/handler/cartsedithandler.go
=======
=======
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/wxapp/api/internal/handler/cartsedithandler.go
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/wxapp/api/internal/handler/cartsedithandler.go
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

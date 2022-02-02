package handler

import (
	"github.com/skywalkerwei/pluton-faster/common/resultx"
	"google.golang.org/grpc/status"
	"net/http"

<<<<<<< HEAD:service/api/wxapp/internal/handler/cartslisthandler.go
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
=======
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/wxapp/api/internal/handler/cartslisthandler.go
)

func CartsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCartsListLogic(r.Context(), svcCtx)
		resp, err := l.CartsList()
		if err != nil {
			r := status.Convert(err)
			httpx.Error(w, resultx.Error(int(r.Code()), r.Message(), nil))
		} else {
			httpx.OkJson(w, resultx.Success(resp))
		}
	}
}

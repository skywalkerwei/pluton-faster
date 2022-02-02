package handler

import (
	"fmt"
	"github.com/skywalkerwei/pluton-faster/common/resultx"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/pay/api/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/pay/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/pay/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CallbackRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, resultx.NewDefault(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewCallbackLogic(r.Context(), svcCtx)
		resp, err := l.Callback(req)
		if err != nil {
			httpx.Error(w, resultx.NewDefault(status.Convert(err).Message(), ""))
		} else {
			httpx.OkJson(w, resultx.NewCode(200, "ok", resp))
		}
	}
}

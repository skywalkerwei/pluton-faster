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

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, resultx.NewDefault(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(req)
		if err != nil {
			r := status.Convert(err)
			httpx.Error(w, resultx.NewDefault(r.Message(), ""))
		} else {
			httpx.OkJson(w, resultx.NewCode(200, "success", resp))
		}
	}
}

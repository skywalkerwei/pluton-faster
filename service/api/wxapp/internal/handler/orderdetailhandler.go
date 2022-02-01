package handler

import (
	"fmt"
	"github.com/skywalkerwei/pluton-faster/common/resultx"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func OrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, resultx.Error(400, fmt.Sprintf("%v", err), nil))
			return
		}

		l := logic.NewOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.OrderDetail(req)
		if err != nil {
			r := status.Convert(err)
			httpx.Error(w, resultx.Error(int(r.Code()), r.Message(), nil))
		} else {
			httpx.OkJson(w, resultx.Success(resp))
		}
	}
}

package handler

import (
	"fmt"
	"github.com/skywalkerwei/pluton-faster/common/errorx"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/pay/api/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/pay/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/pay/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(req)
		if err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

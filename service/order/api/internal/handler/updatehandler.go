package handler

import (
	"fmt"
	"github.com/skywalkerwei/pluton-faster/common/errorx"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/order/api/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/order/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/order/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(req)
		if err != nil {
			//httpx.Error(w, err)
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

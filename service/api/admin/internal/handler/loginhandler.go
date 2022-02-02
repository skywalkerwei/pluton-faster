package handler

import (
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/api/admin/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/api/admin/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

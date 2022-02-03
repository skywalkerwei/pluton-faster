package auth

import (
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic/auth"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DecryptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DecryptRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := auth.NewDecryptLogic(r.Context(), svcCtx)
		resp, err := l.Decrypt(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

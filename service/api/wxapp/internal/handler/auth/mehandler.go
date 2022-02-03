package auth

import (
	"github.com/skywalkerwei/pluton-faster/common/result"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic/auth"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
)

func MeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewMeLogic(r.Context(), svcCtx)
		resp, err := l.Me()
		result.HttpResult(r, w, resp, err)
	}
}

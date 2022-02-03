package auth

import (
	"github.com/skywalkerwei/pluton-faster/common/result"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic/auth"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewCodeLogic(r.Context(), svcCtx)
		resp, err := l.Code(req)
		result.HttpResult(r, w, resp, err)
	}
}

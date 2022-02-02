package handler

import (
	"github.com/skywalkerwei/pluton-faster/common/result"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, resultx.Error(400, fmt.Sprintf("%v", err), nil))
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(req)
		result.HttpResult(r, w, resp, err)
		//if err != nil {
		//	r := status.Convert(err)
		//	httpx.Error(w, resultx.Error(int(r.Code()), r.Message(), nil))
		//} else {
		//	httpx.OkJson(w, resultx.Success(resp))
		//}
	}
}

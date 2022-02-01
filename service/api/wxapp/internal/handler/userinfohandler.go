package handler

import (
	"github.com/skywalkerwei/pluton-faster/common/resultx"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/api/wxapp/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			r := status.Convert(err)
			httpx.Error(w, resultx.Error(int(r.Code()), r.Message(), nil))
		} else {
			httpx.OkJson(w, resultx.Success(resp))
		}
	}
}

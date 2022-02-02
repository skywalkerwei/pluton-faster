package handler

import (
	"net/http"

<<<<<<< HEAD:service/api/admin/internal/handler/userinfohandler.go
	"github.com/skywalkerwei/pluton-faster/service/api/admin/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/api/admin/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
<<<<<<< HEAD:service/api/admin/internal/handler/userinfohandler.go
=======
=======
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/wxapp/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
>>>>>>> c3442f2aed4419e018f66e1d844e317db6477f8c:service/wxapp/api/internal/handler/userinfohandler.go
>>>>>>> 8002f90cd53259a53a26dc23ec27bdbcf3ea6389:service/wxapp/api/internal/handler/userinfohandler.go
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
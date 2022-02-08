package verify

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/skywalkerwei/pluton-faster/common/result"
	"github.com/skywalkerwei/pluton-faster/common/xerr"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/api/identity/internal/logic/verify"
	"github.com/skywalkerwei/pluton-faster/service/api/identity/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/api/identity/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var ErrTokenExpireError = xerr.NewErrCode(xerr.TOKEN_EXPIRE_ERROR)

func TokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := verify.NewTokenLogic(r.Context(), svcCtx)
		resp, err := l.Token(req, r)
		if err == nil && (resp == nil || !resp.Ok) {
			err = errors.Wrapf(ErrTokenExpireError, "jwtAuthHandler JWT Auth no err , userId is zero , req:%+v,resp:%+v", req, resp)
		}

		XUser := "0" //ide会有警告，实际resp不会为nil
		if resp != nil {
			XUser = fmt.Sprintf("%d", resp.UserId)
		}
		w.Header().Set("x-user", XUser)

		result.AuthHttpResult(r, w, resp, err)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
	}
}

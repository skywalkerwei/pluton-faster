package handler

import (
	"fmt"
	"github.com/skywalkerwei/pluton-faster/common/resultx"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/skywalkerwei/pluton-faster/service/user/api/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/user/api/internal/svc"
	"github.com/skywalkerwei/pluton-faster/service/user/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.Error(w, err)
			httpx.Error(w, resultx.NewDefault(fmt.Sprintf("%v", err), ""))
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(req)
		if err != nil {
			r := status.Convert(err)
			httpx.Error(w, resultx.NewDefault(r.Message(), ""))
		} else {
			httpx.OkJson(w, resultx.NewCode(200, "success", resp))
		}
		//if err != nil {
		//	//returnCode := 500
		//	//switch e := err.(type) {
		//	//case *errorx.CodeError:
		//	//	returnCode = e.DataInfo().Code
		//	//default:
		//	//	returnCode = 500
		//	//}
		//	//httpx.Error(w, errorx.NewCodeError(returnCode, fmt.Sprintf("%v", err), ""))
		//	httpx.Error(w, errorx.NewDefaultError(fmt.Sprintf("%v", err), ""))
		//} else {
		//	httpx.OkJson(w, resp)
		//}
	}
}

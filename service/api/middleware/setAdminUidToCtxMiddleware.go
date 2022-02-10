package middleware

import (
	"context"
	"github.com/skywalkerwei/pluton-faster/common/jwtx"
	"net/http"
	"strconv"
)

type SetAdminUidToCtxMiddleware struct {
}

func NewSetAdminUidToCtxMiddleware() *SetUidToCtxMiddleware {
	return &SetUidToCtxMiddleware{}
}

func (m *SetAdminUidToCtxMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := strconv.ParseInt(r.Header.Get("X-User"), 10, 64)
		ctx := r.Context()
		ctx = context.WithValue(ctx, jwtx.CtxKeyJwtAdminId, userId)

		next(w, r.WithContext(ctx))
	}
}

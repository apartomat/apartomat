package graphql

import (
	"context"
	apartomat "github.com/ztsu/apartomat/internal"
	"net/http"
	"strings"
)

const userCtxKey = "user"

type UserCtx struct {
	Email string
}

func ContextWithAuthToken(ver *apartomat.CheckAuthToken, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _ := ver.Do(strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
		if token != nil {
			ctx := context.WithValue(r.Context(), userCtxKey, &UserCtx{Email: token.Subject})
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}

func UserFromCtx(ctx context.Context) *UserCtx {
	user, _ := ctx.Value(userCtxKey).(*UserCtx)
	return user
}

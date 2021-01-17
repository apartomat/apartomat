package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	apartomat "github.com/ztsu/apartomat/internal"
	"net/http"
	"strings"
)

func Handler(ch *apartomat.CheckAuthToken, stores *Stores, resolver ResolverRoot) http.Handler {
	return WithUserHandler(
		ch,
		ContextWithDataLoaders(
			stores,
			handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver})),
		),
	)
}

func WithUserHandler(ver *apartomat.CheckAuthToken, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _ := ver.Do(strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
		if token != nil {
			userCtx := &apartomat.UserCtx{Email: token.Subject}
			r = r.WithContext(apartomat.WithUserCtx(r.Context(), userCtx))
		}

		next.ServeHTTP(w, r)
	})
}

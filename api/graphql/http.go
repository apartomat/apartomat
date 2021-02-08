package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	apartomat "github.com/ztsu/apartomat/internal"
	"log"
	"net/http"
	"strconv"
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
			id, err := strconv.Atoi(token.Get("userId"))
			if err != nil {
				log.Printf("token has not user id: %s", err)

				w.WriteHeader(http.StatusBadRequest)
				return
			}

			userCtx := &apartomat.UserCtx{ID: id, Email: token.Subject}
			r = r.WithContext(apartomat.WithUserCtx(r.Context(), userCtx))
		}

		next.ServeHTTP(w, r)
	})
}

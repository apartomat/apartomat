package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/token"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CheckAuthTokenFn func(str string) (*token.AuthToken, error)

func Handler(ch CheckAuthTokenFn, loaders *apartomat.DataLoaders, resolver ResolverRoot) http.Handler {
	return CorsHandler(
		WithDataLoadersHandler(
			loaders,
			WithUserHandler(
				ch,
				handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver})),
			),
		),
	)
}

func WithUserHandler(checkAuthToken CheckAuthTokenFn, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, _ := checkAuthToken(strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
		if t != nil {
			id, err := strconv.Atoi(t.Get("userId"))
			if err != nil {
				log.Printf("token has not user id: %s", err)

				w.WriteHeader(http.StatusBadRequest)
				return
			}

			userCtx := &apartomat.UserCtx{ID: id, Email: t.Subject}
			r = r.WithContext(apartomat.WithUserCtx(r.Context(), userCtx))
		}

		next.ServeHTTP(w, r)
	})
}

func WithDataLoadersHandler(loaders *apartomat.DataLoaders, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(apartomat.WithDataLoadersCtx(r.Context(), loaders))
		next.ServeHTTP(w, r)
	})
}

func CorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		next.ServeHTTP(w, r)
	})
}

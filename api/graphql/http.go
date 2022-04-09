package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	apartomat "github.com/apartomat/apartomat/internal"
	"net/http"
	"strings"
)

type CheckAuthTokenFn func(str string) (apartomat.AuthToken, error)

func Handler(
	ch CheckAuthTokenFn,
	loaders *apartomat.DataLoaders,
	resolver ResolverRoot,
	complexityLimit int,
) http.Handler {
	gh := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver}))
	gh.Use(extension.FixedComplexityLimit(complexityLimit))

	return CorsHandler(
		WithDataLoadersHandler(
			loaders,
			WithUserHandler(ch, gh),
		),
	)
}

func WithUserHandler(checkAuthToken CheckAuthTokenFn, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			header = r.Header.Get("Authorization")
		)

		if header == "" {
			header = r.Header.Get("X-Authorization")
		}

		t, _ := checkAuthToken(strings.TrimPrefix(header, "Bearer "))
		if t != nil {
			userCtx := &apartomat.UserCtx{ID: t.UserID()}
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

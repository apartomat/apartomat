package graphql

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/apartomat/apartomat/internal/auth"
	"github.com/apartomat/apartomat/internal/dataloader"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"time"
)

type CheckAuthTokenFn func(str string) (auth.AuthToken, error)

func Handler(
	ch CheckAuthTokenFn,
	loaders *dataloader.DataLoaders,
	resolver ResolverRoot,
	complexityLimit int,
) http.Handler {
	var (
		gh = handler.New(NewExecutableSchema(Config{Resolvers: resolver}))
	)

	gh.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, payload transport.InitPayload) (context.Context, error) {
			if t, _ := ch(payload.Authorization()); t != nil {
				return auth.WithUserCtx(ctx, &auth.UserCtx{ID: t.UserID()}), nil
			}

			return ctx, nil
		},
	})

	gh.AddTransport(transport.Options{})
	gh.AddTransport(transport.GET{})
	gh.AddTransport(transport.POST{})
	gh.AddTransport(transport.MultipartForm{})

	gh.SetQueryCache(lru.New(1000))

	gh.Use(extension.Introspection{})

	gh.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

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
			userCtx := &auth.UserCtx{ID: t.UserID()}
			r = r.WithContext(auth.WithUserCtx(r.Context(), userCtx))
		}

		next.ServeHTTP(w, r)
	})
}

func WithDataLoadersHandler(loaders *dataloader.DataLoaders, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(dataloader.WithDataLoadersCtx(r.Context(), loaders))
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

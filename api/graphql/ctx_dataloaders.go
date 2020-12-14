package graphql

import (
	"context"
	"net/http"
)

const dataLoadersKey = "dataloaders"

type loaders struct{}

type Stores struct{}

func ContextWithDataLoaders(services *Stores, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), dataLoadersKey, &loaders{})

		r = r.WithContext(ctx)

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		next.ServeHTTP(w, r)
	})
}

func DataLoaders(ctx context.Context) *loaders {
	return ctx.Value(dataLoadersKey).(*loaders)
}

package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	apartomat "github.com/ztsu/apartomat/internal"
	"net/http"
)

func Handler(ch *apartomat.CheckAuthToken, stores *Stores, resolver ResolverRoot) http.Handler {
	return ContextWithAuthToken(
		ch,
		ContextWithDataLoaders(
			stores,
			handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver})),
		),
	)
}

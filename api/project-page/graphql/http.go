package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"net/http"
)

func Handler(
	resolver ResolverRoot,
) http.Handler {
	var (
		gh = handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver}))
	)

	return gh
}

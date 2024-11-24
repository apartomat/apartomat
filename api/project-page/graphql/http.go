package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"net/http"
)

func Handler(resolver ResolverRoot) http.Handler {
	return handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: resolver}))
}

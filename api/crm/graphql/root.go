package graphql

import (
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/uptrace/bun"
)

type rootResolver struct {
	db       *bun.DB
	useCases *apartomat.Apartomat
}

func NewRootResolver(db *bun.DB, useCases *apartomat.Apartomat) ResolverRoot {
	return &rootResolver{db: db, useCases: useCases}
}

func notFound() (NotFound, error) {
	return NotFound{Message: "not found"}, nil
}

func forbidden() (Forbidden, error) {
	return Forbidden{Message: "forbidden"}, nil
}

func serverError() (ServerError, error) {
	return ServerError{Message: "server error"}, nil
}

func notImplementedYetError() (ServerError, error) {
	return ServerError{Message: "not implemented yet"}, nil
}

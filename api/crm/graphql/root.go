package graphql

import (
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/uptrace/bun"
)

type rootResolver struct {
	db       *bun.DB
	useCases *crm.CRM
}

func NewRootResolver(db *bun.DB, useCases *crm.CRM) ResolverRoot {
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

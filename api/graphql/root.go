package graphql

import (
	apartomat "github.com/apartomat/apartomat/internal"
)

type rootResolver struct {
	useCases *apartomat.Apartomat
}

func NewRootResolver(useCases *apartomat.Apartomat) ResolverRoot {
	return &rootResolver{useCases: useCases}
}

func notFound() (NotFound, error) {
	return NotFound{Message: "not found"}, nil
}

func forbidden() (Forbidden, error) {
	return Forbidden{Message: "not found"}, nil
}

func serverError() (ServerError, error) {
	return ServerError{Message: "server error"}, nil
}

func notImplementedYetError() (ServerError, error) {
	return ServerError{Message: "not implemented yet"}, nil
}

package graphql

import (
	"github.com/apartomat/apartomat/internal/project-page"
)

type rootResolver struct {
	projectPage *project_page.Service
}

func NewRootResolver(service *project_page.Service) ResolverRoot {
	return &rootResolver{service}
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

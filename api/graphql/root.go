package graphql

import (
	"context"
	apartomat "github.com/ztsu/apartomat/internal"
)

type UseCases struct {
	LoginByEmail            *apartomat.LoginByEmail
	ConfirmLogin            *apartomat.ConfirmLogin
	CheckAuthToken          *apartomat.CheckAuthToken
	GetUserProfile          *apartomat.GetUserProfile
	GetDefaultWorkspace     *apartomat.GetDefaultWorkspace
	GetWorkspace            *apartomat.GetWorkspace
	GetWorkspaceUsers       *apartomat.GetWorkspaceUsers
	GetWorkspaceUserProfile *apartomat.GetWorkspaceUserProfile
	GetWorkspaceProjects    *apartomat.GetWorkspaceProjects
}

type rootResolver struct {
	useCases *UseCases
}

func NewRootResolver(useCases *UseCases) ResolverRoot {
	return &rootResolver{useCases: useCases}
}

func (r *rootResolver) Query() QueryResolver { return &queryResolver{r} }

func (r *rootResolver) Mutation() MutationResolver { return &mutationResolver{r} }

type queryResolver struct {
	*rootResolver
}

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	return "", nil
}

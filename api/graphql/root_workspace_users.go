package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"log"
)

func (r *rootResolver) WorkspaceUsers() WorkspaceUsersResolver {
	return &workspaceUsersResolver{r}
}

type workspaceUsersResolver struct {
	*rootResolver
}

func (r *workspaceUsersResolver) List(
	ctx context.Context,
	obj *WorkspaceUsers,
	filter WorkspaceUsersFilter,
	limit int,
) (WorkspaceUsersListResult, error) {
	var (
		workspace *Workspace
	)

	if w, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Workspace); !ok {
		log.Printf("can't resolve project contacts: %s", errors.New("unknown workspace"))

		return serverError()
	} else {
		workspace = w
	}

	users, err := r.useCases.GetWorkspaceUsers(ctx, workspace.ID, limit, 0)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't resolve workspace (id=%s) users: %s", workspace.ID, err)

		return serverError()
	}

	return &WorkspaceUsersList{Items: workspaceUsersToGraphQL(users)}, nil
}

func (r *workspaceUsersResolver) Total(
	ctx context.Context,
	obj *WorkspaceUsers,
	filter WorkspaceUsersFilter,
) (WorkspaceUsersTotalResult, error) {
	return nil, errors.New("not implemented yet")
}

func workspaceUsersToGraphQL(users []*store.WorkspaceUser) []*WorkspaceUser {
	result := make([]*WorkspaceUser, 0, len(users))

	for _, u := range users {
		result = append(result, workspaceUserToGraphQL(u))
	}

	return result
}

func workspaceUserToGraphQL(wu *store.WorkspaceUser) *WorkspaceUser {
	return &WorkspaceUser{
		ID:        wu.UserID,
		Workspace: &ID{ID: wu.WorkspaceID},
		Role:      workspaceUserRoleToGraphQL(wu.Role),
	}
}

func workspaceUserRoleToGraphQL(role store.WorkspaceUserRole) WorkspaceUserRole {
	switch role {
	case store.WorkspaceUserRoleAdmin:
		return WorkspaceUserRoleAdmin
	case store.WorkspaceUserRoleUser:
		return WorkspaceUserRoleUser
	}

	return ""
}

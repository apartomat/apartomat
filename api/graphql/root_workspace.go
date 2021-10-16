package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"log"
)

type workspaceResolver struct {
	*rootResolver
}

func (r *rootResolver) Workspace() WorkspaceResolver { return &workspaceResolver{r} }

func (r *workspaceResolver) Users(ctx context.Context, obj *Workspace) (WorkspaceUsersResult, error) {
	users, err := r.useCases.GetWorkspaceUsers.Do(ctx, obj.ID, 5, 0)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't resolve workspace (id=%d) users: %s", obj.ID, err)

		return ServerError{Message: "internal server error"}, nil
	}

	return &WorkspaceUsers{Items: workspaceUsersToGraphQL(users)}, nil
}

func (r *workspaceResolver) Projects(ctx context.Context, obj *Workspace) (*WorkspaceProjects, error) {
	return &WorkspaceProjects{Workspace: &ID{ID: obj.ID}}, nil
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

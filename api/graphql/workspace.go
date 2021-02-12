package graphql

import (
	"context"
	"github.com/pkg/errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"log"
)

func (r *rootResolver) Workspace() WorkspaceResolver {
	return &workspaceResolver{r}

}

func (r *queryResolver) Workspace(ctx context.Context, id int) (WorkspaceResult, error) {
	ws, err := r.useCases.GetWorkspace.Do(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't resolve workspace (id=%d): %s", id, err)

		return ServerError{}, nil
	}

	return Workspace{ID: ws.ID, Name: ws.Name}, nil

}

type workspaceResolver struct {
	*rootResolver
}

func (r *workspaceResolver) Users(ctx context.Context, obj *Workspace) (WorkspaceUsersResult, error) {
	users, err := r.useCases.GetWorkspaceUsers.Do(ctx, obj.ID, 5, 0)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't resolve workspace (id=%d) users: %s", obj.ID, err)

		return ServerError{Message: "internal server error"}, nil
	}

	return &WorkspaceUsers{Items: workspacUsersToGraphQL(users)}, nil
}

func (r *workspaceResolver) Projects(ctx context.Context, obj *Workspace) (*WorkspaceProjects, error) {
	return &WorkspaceProjects{Workspace: &ID{ID: obj.ID}}, nil
}

func workspaceUserToGraphQL(user *store.WorkspaceUser) *WorkspaceUser {
	return &WorkspaceUser{ID: user.ID, Role: workspaceUserRoleToGraphQL(user.Role)}
}

func workspacUsersToGraphQL(users []*store.WorkspaceUser) []*WorkspaceUser {
	result := make([]*WorkspaceUser, 0, len(users))

	for _, u := range users {
		result = append(result, workspaceUserToGraphQL(u))
	}

	return result
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

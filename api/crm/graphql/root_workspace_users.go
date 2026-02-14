package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/workspaceusers"
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
		slog.ErrorContext(ctx, "can't resolve project contacts", slog.String("err", "unknown workspace"))

		return serverError()
	} else {
		workspace = w
	}

	users, err := r.crm.GetWorkspaceUsers(ctx, workspace.ID, limit, 0)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(ctx, "can't resolve workspace users", slog.String("workspace", workspace.ID), slog.Any("err", err))

		return serverError()
	}

	return &WorkspaceUsersList{Items: workspaceUsersToGraphQL(users)}, nil
}

func (r *workspaceUsersResolver) Total(
	ctx context.Context,
	obj *WorkspaceUsers,
	filter WorkspaceUsersFilter,
) (WorkspaceUsersTotalResult, error) {
	return notImplementedYetError()
}

func workspaceUsersToGraphQL(users []*workspaceusers.WorkspaceUser) []*WorkspaceUser {
	result := make([]*WorkspaceUser, 0, len(users))

	for _, u := range users {
		result = append(result, workspaceUserToGraphQL(u))
	}

	return result
}

func workspaceUserToGraphQL(wu *workspaceusers.WorkspaceUser) *WorkspaceUser {
	return &WorkspaceUser{
		ID:        wu.UserID,
		Workspace: &ID{ID: wu.WorkspaceID},
		Role:      workspaceUserRoleToGraphQL(wu.Role),
	}
}

func workspaceUserRoleToGraphQL(role workspaceusers.WorkspaceUserRole) WorkspaceUserRole {
	switch role {
	case workspaceusers.WorkspaceUserRoleAdmin:
		return WorkspaceUserRoleAdmin
	case workspaceusers.WorkspaceUserRoleUser:
		return WorkspaceUserRoleUser
	}

	return ""
}

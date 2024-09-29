package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/projects"
)

type workspaceProjectsResolver struct {
	*rootResolver
}

func (r *rootResolver) WorkspaceProjects() WorkspaceProjectsResolver {
	return &workspaceProjectsResolver{r}
}

func (r *workspaceProjectsResolver) List(
	ctx context.Context,
	obj *WorkspaceProjects,
	filter WorkspaceProjectsFilter,
	limit int,
) (WorkspaceProjectsListResult, error) {
	workspace, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Workspace)

	if !ok {
		slog.ErrorContext(ctx, "can't resolve workspace projects", slog.String("err", "unknown workspace"))

		return serverError()
	}

	items, err := r.useCases.GetWorkspaceProjects(
		ctx,
		workspace.ID,
		toProjectStatuses(filter.Status),
		limit,
		0,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(
			ctx,
			"can't resolve workspace projects",
			slog.String("workspace", workspace.ID),
			slog.Any("err", err),
		)

		return serverError()
	}

	return WorkspaceProjectsList{Items: projectsToGraphQL(items)}, nil
}

func (r *workspaceProjectsResolver) Total(
	ctx context.Context,
	obj *WorkspaceProjects,
	filter WorkspaceProjectsFilter,
) (WorkspaceProjectsTotalResult, error) {
	return notImplementedYetError() // TODO
}

func projectsToGraphQL(projects []*projects.Project) []*Project {
	result := make([]*Project, 0, len(projects))

	for _, p := range projects {
		result = append(result, projectToGraphQL(p))
	}

	return result
}

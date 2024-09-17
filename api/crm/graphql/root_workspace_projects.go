package graphql

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/projects"
	"go.uber.org/zap"
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
		r.logger.Error("can't resolve workspace projects", zap.Error(errors.New("unknown workspace")))

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

		r.logger.Error(
			"can't resolve workspace projects",
			zap.String("workspace", workspace.ID),
			zap.Error(err),
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

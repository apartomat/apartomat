package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"log"
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
		log.Printf("can't resolve workspace projects: %s", errors.New("unknown workspace"))

		return serverError()
	}

	projects, err := r.useCases.GetWorkspaceProjects(
		ctx,
		workspace.ID,
		apartomat.GetWorkspaceProjectsFilter{
			Status: store.ProjectStatusExpr{
				Eq: toProjectStatuses(filter.Status),
			},
		},
		limit,
		0,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't resolve workspace (id=%s) projects: %s", workspace.ID, err)

		return ServerError{Message: "internal server error"}, nil
	}

	return WorkspaceProjectsList{Items: projectsToGraphQL(projects)}, nil
}

func (r *workspaceProjectsResolver) Total(
	ctx context.Context,
	obj *WorkspaceProjects,
	filter WorkspaceProjectsFilter,
) (WorkspaceProjectsTotalResult, error) {
	return notImplementedYetError() // TODO
}

func projectsToGraphQL(projects []*store.Project) []*Project {
	result := make([]*Project, 0, len(projects))

	for _, p := range projects {
		result = append(result, projectToGraphQL(p))
	}

	return result
}

package graphql

import (
	"context"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
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
	projects, err := r.useCases.GetWorkspaceProjects(
		ctx,
		obj.Workspace.ID,
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

		log.Printf("can't resolve workspace (id=%d) projects: %s", obj.Workspace.ID, err)

		return ServerError{Message: "internal server error"}, nil
	}

	return WorkspaceProjectsList{Items: projectsToGraphQLWorkspaceProjects(projects)}, nil
}

func (r *workspaceProjectsResolver) Total(
	ctx context.Context,
	obj *WorkspaceProjects,
	filter WorkspaceProjectsFilter,
) (WorkspaceProjectsTotalResult, error) {
	return notImplementedYetError() // TODO
}

func projectToGraphQLWorkspaceProject(project *store.Project) *WorkspaceProject {
	wp := &WorkspaceProject{
		ID:      project.ID,
		Name:    project.Name,
		Status:  projectStatusToGraphQL(project.Status),
		StartAt: project.StartAt,
		EndAt:   project.EndAt,
	}

	return wp
}

func projectsToGraphQLWorkspaceProjects(projects []*store.Project) []*WorkspaceProject {
	result := make([]*WorkspaceProject, 0, len(projects))

	for _, u := range projects {
		result = append(result, projectToGraphQLWorkspaceProject(u))
	}

	return result
}

func pstring(str string) *string {
	return &str
}

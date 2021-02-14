package graphql

import (
	"context"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/pkg/errors"
	"log"
)

func (r *rootResolver) WorkspaceProjects() WorkspaceProjectsResolver {
	return &workspaceProjectsResolver{r}
}

type workspaceProjectsResolver struct {
	*rootResolver
}

func (r *workspaceProjectsResolver) List(ctx context.Context, obj *WorkspaceProjects) (WorkspaceProjectsListResult, error) {
	projects, err := r.useCases.GetWorkspaceProjects.Do(ctx, obj.Workspace.ID, 10, 0)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		log.Printf("can't resolve workspace (id=%d) projects: %s", obj.Workspace.ID, err)

		return ServerError{Message: "internal server error"}, nil
	}

	return WorkspaceProjectsList{Items: projectsToGraphQL(projects)}, nil
}

func (r *workspaceProjectsResolver) Total(ctx context.Context, obj *WorkspaceProjects) (WorkspaceProjectsTotalResult, error) {
	// todo
	return ServerError{Message: "not implemented yet"}, nil
}

func projectToGraphQL(project *store.Project) *WorkspaceProject {
	return &WorkspaceProject{ID: project.ID, Name: project.Name}
}

func projectsToGraphQL(projects []*store.Project) []*WorkspaceProject {
	result := make([]*WorkspaceProject, 0, len(projects))

	for _, u := range projects {
		result = append(result, projectToGraphQL(u))
	}

	return result
}

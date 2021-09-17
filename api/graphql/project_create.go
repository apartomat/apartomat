package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input CreateProjectInput) (CreateProjectResult, error) {
	project, err := r.useCases.CreateProject.Do(ctx, input.WorkspaceID, input.Title)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		log.Printf("can't create project in workspace (id=%d): %s", input.WorkspaceID, err)

		return serverError()
	}

	return projectToGraphQL(project), nil
}

package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input CreateProjectInput) (CreateProjectResult, error) {
	project, err := r.useCases.CreateProject(ctx, input.WorkspaceID, input.Name, input.StartAt, input.EndAt)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		r.logger.Error(
			"can't create project in workspace",
			zap.String("workspace", input.WorkspaceID),
			zap.Error(err),
		)

		return serverError()
	}

	return ProjectCreated{Project: projectToGraphQL(project)}, nil
}

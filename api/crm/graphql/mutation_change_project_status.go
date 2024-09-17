package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

func (r *mutationResolver) ChangeProjectStatus(
	ctx context.Context,
	projectID string,
	status ProjectStatus,
) (ChangeProjectStatusResult, error) {
	project, err := r.useCases.ChangeProjectStatus(
		ctx,
		projectID,
		toProjectStatus(status),
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		r.logger.Error("can't change project status", zap.Error(err))

		return serverError()
	}

	return ProjectStatusChanged{Project: projectToGraphQL(project)}, nil
}

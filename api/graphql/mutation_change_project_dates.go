package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

func (r *mutationResolver) ChangeProjectDates(
	ctx context.Context,
	projectID string,
	input ChangeProjectDatesInput,
) (ChangeProjectDatesResult, error) {
	project, err := r.useCases.ChangeProjectDates(
		ctx,
		projectID,
		input.StartAt,
		input.EndAt,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		r.logger.Error("can't change project dates", zap.String("project", projectID), zap.Error(err))

		return serverError()
	}

	return ProjectDatesChanged{Project: projectToGraphQL(project)}, nil
}

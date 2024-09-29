package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
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

		slog.ErrorContext(
			ctx,
			"can't change project dates",
			slog.String("project", projectID),
			slog.Any("err", err),
		)

		return serverError()
	}

	return ProjectDatesChanged{Project: projectToGraphQL(project)}, nil
}

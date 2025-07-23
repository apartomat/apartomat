package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) ChangeProjectDates(
	ctx context.Context,
	projectID string,
	input ChangeProjectDatesInput,
) (ChangeProjectDatesResult, error) {
	project, err := r.crm.ChangeProjectDates(
		ctx,
		projectID,
		input.StartAt,
		input.EndAt,
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
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

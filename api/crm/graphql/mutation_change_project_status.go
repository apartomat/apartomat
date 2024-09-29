package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
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

		slog.ErrorContext(ctx, "can't change project status", slog.Any("err", err))

		return serverError()
	}

	return ProjectStatusChanged{Project: projectToGraphQL(project)}, nil
}

package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) ChangeProjectStatus(
	ctx context.Context,
	projectID string,
	status ProjectStatus,
) (ChangeProjectStatusResult, error) {
	project, err := r.crm.ChangeProjectStatus(
		ctx,
		projectID,
		toProjectStatus(status),
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't change project status", slog.Any("err", err))

		return serverError()
	}

	return ProjectStatusChanged{Project: projectToGraphQL(project)}, nil
}

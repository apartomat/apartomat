package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *queryResolver) Project(ctx context.Context, id string) (ProjectResult, error) {
	project, err := r.crm.GetProject(ctx, id)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't resolve project", slog.String("project", id), slog.Any("err", err))

		return serverError()
	}

	return projectToGraphQL(project), nil
}

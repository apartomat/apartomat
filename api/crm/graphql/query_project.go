package graphql

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
)

func (r *queryResolver) Project(ctx context.Context, id string) (ProjectResult, error) {
	project, err := r.useCases.GetProject(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't resolve project", slog.String("project", id), slog.Any("err", err))

		return nil, fmt.Errorf("can't resolve project (id=%s): %w", id, err)
	}

	return projectToGraphQL(project), nil
}

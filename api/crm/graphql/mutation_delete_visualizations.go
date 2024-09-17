package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

func (r *mutationResolver) DeleteVisualizations(ctx context.Context, id []string) (DeleteVisualizationsResult, error) {
	res, err := r.useCases.DeleteVisualizations(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		r.logger.Error("can't delete visualizations", zap.Error(err))

		return serverError()
	}

	if len(res) != len(id) {
		return SomeVisualizationsDeleted{Visualizations: visualizationsToGraphQL(res)}, nil
	}

	return VisualizationsDeleted{Visualizations: visualizationsToGraphQL(res)}, nil
}

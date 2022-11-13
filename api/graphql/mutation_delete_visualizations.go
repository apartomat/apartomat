package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) DeleteVisualizations(ctx context.Context, id []string) (DeleteVisualizationsResult, error) {
	res, err := r.useCases.DeleteVisualizations(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't delete visualizations: %s", err)

		return nil, err
	}

	if len(res) != len(id) {
		return SomeVisualizationsDeleted{Visualizations: visualizationsToGraphQL(res)}, nil
	}

	return VisualizationsDeleted{Visualizations: visualizationsToGraphQL(res)}, nil
}

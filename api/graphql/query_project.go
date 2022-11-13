package graphql

import (
	"context"
	"errors"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
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

		log.Printf("can't resolve project (id=%s): %s", id, err)

		return nil, fmt.Errorf("can't resolve project (id=%s): %w", id, err)
	}

	return projectToGraphQL(project), nil
}

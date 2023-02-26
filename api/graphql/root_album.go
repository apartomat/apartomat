package graphql

import (
	"context"
	"errors"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *rootResolver) Album() AlbumResolver { return &albumResolver{r} }

type albumResolver struct {
	*rootResolver
}

func (r *albumResolver) Pages(ctx context.Context, obj *Album) (AlbumPagesResult, error) {
	return obj.Pages, nil
}

func (r *albumResolver) Project(ctx context.Context, obj *Album) (AlbumProjectResult, error) {
	var (
		gp *Project
	)

	if pr, ok := obj.Project.(Project); ok {
		gp = &pr
	}

	if gp == nil {
		return serverError()
	}

	project, err := r.useCases.GetProject(ctx, gp.ID)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		log.Printf("can't resolve project (id=%s): %s", gp.ID, err)

		return nil, fmt.Errorf("can't resolve project (id=%s): %w", gp.ID, err)
	}

	return projectToGraphQL(project), nil
}

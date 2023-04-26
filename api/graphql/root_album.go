package graphql

import (
	"context"
	"errors"
	"fmt"
	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
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

		r.logger.Error("can't resolve project", zap.String("project", gp.ID), zap.Error(err))

		return nil, fmt.Errorf("can't resolve project (id=%s): %w", gp.ID, err)
	}

	return projectToGraphQL(project), nil
}

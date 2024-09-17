package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

func (r *queryResolver) Album(ctx context.Context, id string) (AlbumResult, error) {
	album, err := r.useCases.GetAlbum(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		r.logger.Error("can't resolve project", zap.String("project", id), zap.Error(err))

		return serverError()
	}

	return albumToGraphQL(album), nil
}

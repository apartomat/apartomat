package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
)

func (r *mutationResolver) DeleteAlbum(ctx context.Context, id string) (DeleteAlbumResult, error) {
	album, err := r.useCases.DeleteAlbum(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		r.logger.Error("can't delete album", zap.Error(err))

		return serverError()
	}

	return AlbumDeleted{Album: albumToGraphQL(album)}, nil
}

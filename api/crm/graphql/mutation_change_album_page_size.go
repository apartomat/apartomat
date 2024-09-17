package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/albums"
	"go.uber.org/zap"
)

func (r *mutationResolver) ChangeAlbumPageSize(
	ctx context.Context,
	albumID string,
	size PageSize,
) (ChangeAlbumPageSizeResult, error) {
	album, err := r.useCases.ChangeAlbumPageSize(ctx, albumID, albums.PageSize(size))

	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		r.logger.Error("can't change album page size", zap.Error(err))

		return serverError()
	}

	return AlbumPageSizeChanged{Album: albumToGraphQL(album)}, nil
}

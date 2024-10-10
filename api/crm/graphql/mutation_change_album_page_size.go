package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/albums"
)

func (r *mutationResolver) ChangeAlbumPageSize(
	ctx context.Context,
	albumID string,
	size PageSize,
) (ChangeAlbumPageSizeResult, error) {
	album, err := r.useCases.ChangeAlbumPageSize(ctx, albumID, albums.PageSize(size))

	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't change album page size", slog.Any("err", err))

		return serverError()
	}

	return AlbumPageSizeChanged{Album: albumToGraphQL(album)}, nil
}

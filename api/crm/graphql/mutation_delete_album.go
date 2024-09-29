package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
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

		slog.ErrorContext(ctx, "can't delete album", slog.Any("err", err))

		return serverError()
	}

	return AlbumDeleted{Album: albumToGraphQL(album)}, nil
}

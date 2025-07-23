package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) DeleteAlbum(ctx context.Context, id string) (DeleteAlbumResult, error) {
	album, err := r.crm.DeleteAlbum(ctx, id)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't delete album", slog.Any("err", err))

		return serverError()
	}

	return AlbumDeleted{Album: albumToGraphQL(album)}, nil
}

package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/albums"
)

func (r *mutationResolver) ChangeAlbumPageOrientation(
	ctx context.Context,
	albumID string,
	orientation PageOrientation,
) (ChangeAlbumPageOrientationResult, error) {
	album, err := r.useCases.ChangeAlbumPageOrientation(ctx, albumID, albums.PageOrientation(orientation))

	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't change album page orientation", slog.Any("err", err))

		return serverError()
	}

	return AlbumPageOrientationChanged{Album: albumToGraphQL(album)}, nil
}

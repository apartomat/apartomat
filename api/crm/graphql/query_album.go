package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
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

		slog.ErrorContext(ctx, "can't resolve project", slog.String("project", id), slog.Any("err", err))

		return serverError()
	}

	return albumToGraphQL(album), nil
}

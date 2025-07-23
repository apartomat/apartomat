package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *queryResolver) Album(ctx context.Context, id string) (AlbumResult, error) {
	album, err := r.crm.GetAlbum(ctx, id)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't resolve project", slog.String("project", id), slog.Any("err", err))

		return serverError()
	}

	return albumToGraphQL(album), nil
}

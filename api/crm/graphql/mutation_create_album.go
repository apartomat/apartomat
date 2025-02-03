package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) CreateAlbum(
	ctx context.Context,
	projectID, name string,
	settings CreateAlbumSettingsInput,
) (CreateAlbumResult, error) {
	album, err := r.useCases.CreateAlbum(
		ctx,
		projectID,
		name,
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		slog.ErrorContext(ctx, "can't create album", slog.Any("err", err))

		return serverError()
	}

	return AlbumCreated{Album: albumToGraphQL(album)}, nil
}

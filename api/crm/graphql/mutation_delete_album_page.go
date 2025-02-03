package graphql

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) DeleteAlbumPage(
	ctx context.Context,
	albumID string,
	pageNumber int,
) (DeleteAlbumPageResult, error) {
	time.Sleep(3 * time.Second)
	page, err := r.useCases.DeleteAlbumPage(ctx, albumID, pageNumber)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't delete album page", slog.Any("err", err))

		return serverError()
	}

	return AlbumPageDeleted{Page: albumPageToGraphQL(*page, pageNumber)}, nil
}

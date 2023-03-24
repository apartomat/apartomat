package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/albums"
	"log"
)

func (r *mutationResolver) ChangeAlbumPageSize(
	ctx context.Context,
	albumID string,
	size PageSize,
) (ChangeAlbumPageSizeResult, error) {
	album, err := r.useCases.ChangeAlbumPageSize(ctx, albumID, albums.PageSize(size))

	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't change album page size: %s", err)

		return ServerError{Message: "can't change album page size"}, nil
	}

	return AlbumPageSizeChanged{Album: albumToGraphQL(album)}, nil
}

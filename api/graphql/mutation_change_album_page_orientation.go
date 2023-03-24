package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/albums"
	"log"
)

func (r *mutationResolver) ChangeAlbumPageOrientation(
	ctx context.Context,
	albumID string,
	orientation PageOrientation,
) (ChangeAlbumPageOrientationResult, error) {
	album, err := r.useCases.ChangeAlbumPageOrientation(ctx, albumID, albums.PageOrientation(orientation))

	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't change album page orientation: %s", err)

		return ServerError{Message: "can't change album page orientation"}, nil
	}

	return AlbumPageOrientationChanged{Album: albumToGraphQL(album)}, nil
}

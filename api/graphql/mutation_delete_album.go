package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) DeleteAlbum(ctx context.Context, id string) (DeleteAlbumResult, error) {
	album, err := r.useCases.DeleteAlbum(ctx, id)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't delete album: %s", err)

		return ServerError{Message: "can't delete album"}, nil
	}

	return AlbumDeleted{Album: albumToGraphQL(album)}, nil
}

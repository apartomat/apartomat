package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) DeleteRoom(ctx context.Context, id string) (DeleteRoomResult, error) {
	room, err := r.useCases.DeleteRoom(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		log.Printf("can't delete room: %s", err)

		return nil, errors.New("server error: can't delete room")
	}

	return RoomDeleted{Room: roomToGraphQL(room)}, nil
}

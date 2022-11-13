package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) AddRoom(ctx context.Context, houseID string, input AddRoomInput) (AddRoomResult, error) {
	room, err := r.useCases.AddRoom(
		ctx,
		houseID,
		input.Name,
		input.Square,
		input.Level,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		log.Printf("can't add room: %s", err)

		return nil, errors.New("server error: can't add room")
	}

	return RoomAdded{Room: roomToGraphQL(room)}, nil
}

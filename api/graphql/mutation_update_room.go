package graphql

import (
	"context"
	"errors"
	apartomat "github.com/apartomat/apartomat/internal"
	"log"
)

func (r *mutationResolver) UpdateRoom(
	ctx context.Context,
	roomID string,
	input UpdateRoomInput,
) (UpdateRoomResult, error) {
	room, err := r.useCases.UpdateRoom(
		ctx,
		roomID,
		input.Name,
		input.Square,
		input.Level,
	)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return Forbidden{}, nil
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return NotFound{}, nil
		}

		log.Printf("can't update room: %s", err)

		return nil, errors.New("server error: can't update room")
	}

	return RoomUpdated{Room: roomToGraphQL(room)}, nil
}

package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
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
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't update room", slog.Any("err", err))

		return nil, errors.New("server error: can't update room")
	}

	return RoomUpdated{Room: roomToGraphQL(room)}, nil
}

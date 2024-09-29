package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
)

func (r *mutationResolver) MoveRoomToPosition(
	ctx context.Context,
	roomID string,
	position int,
) (MoveRoomToPositionResult, error) {
	room, err := r.useCases.MoveRoomToPosition(ctx, roomID, position)
	if err != nil {
		if errors.Is(err, apartomat.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, apartomat.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't move room", slog.Any("err", err))

		return nil, errors.New("server error: can't move room")
	}

	return RoomMovedToPosition{Room: roomToGraphQL(room)}, nil
}

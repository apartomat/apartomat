package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) MoveRoomToPosition(
	ctx context.Context,
	roomID string,
	position int,
) (MoveRoomToPositionResult, error) {
	room, err := r.crm.MoveRoomToPosition(ctx, roomID, position)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't move room", slog.Any("err", err))

		return nil, errors.New("server error: can't move room")
	}

	return RoomMovedToPosition{Room: roomToGraphQL(room)}, nil
}

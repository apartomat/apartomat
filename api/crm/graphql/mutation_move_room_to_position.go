package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
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

		r.logger.Error("can't move room", zap.Error(err))

		return nil, errors.New("server error: can't move room")
	}

	return RoomMovedToPosition{Room: roomToGraphQL(room)}, nil
}

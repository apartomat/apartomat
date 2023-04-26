package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"go.uber.org/zap"
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

		r.logger.Error("can't delete room", zap.Error(err))

		return nil, errors.New("server error: can't delete room")
	}

	return RoomDeleted{Room: roomToGraphQL(room)}, nil
}

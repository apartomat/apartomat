package graphql

import (
	"context"
	"errors"
	"log/slog"

	apartomat "github.com/apartomat/apartomat/internal"
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

		slog.ErrorContext(ctx, "can't delete room", slog.Any("err", err))

		return nil, errors.New("server error: can't delete room")
	}

	return RoomDeleted{Room: roomToGraphQL(room)}, nil
}

package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) DeleteRoom(ctx context.Context, id string) (DeleteRoomResult, error) {
	room, err := r.crm.DeleteRoom(
		ctx,
		id,
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't delete room", slog.Any("err", err))

		return nil, errors.New("server error: can't delete room")
	}

	return RoomDeleted{Room: roomToGraphQL(room)}, nil
}

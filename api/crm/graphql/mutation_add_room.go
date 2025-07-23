package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) AddRoom(ctx context.Context, houseID string, input AddRoomInput) (AddRoomResult, error) {
	room, err := r.crm.AddRoom(
		ctx,
		houseID,
		input.Name,
		input.Square,
		input.Level,
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't add room", slog.Any("err", err))

		return nil, errors.New("server error: can't add room")
	}

	return RoomAdded{Room: roomToGraphQL(room)}, nil
}

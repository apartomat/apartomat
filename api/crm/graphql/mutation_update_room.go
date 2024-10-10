package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
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
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(ctx, "can't update room", slog.Any("err", err))

		return nil, errors.New("server error: can't update room")
	}

	return RoomUpdated{Room: roomToGraphQL(room)}, nil
}

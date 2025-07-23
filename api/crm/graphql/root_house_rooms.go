package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/rooms"
)

func (r *rootResolver) HouseRooms() HouseRoomsResolver {
	return &houseRoomsResolver{r}
}

type houseRoomsResolver struct {
	*rootResolver
}

func (r *houseRoomsResolver) List(ctx context.Context, obj *HouseRooms, limit int, offset int) (HouseRoomsListResult, error) {
	if phouse, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(**House); !ok {
		slog.ErrorContext(ctx, "can't resolve house rooms", slog.String("err", "unknown house"))

		return serverError()
	} else {
		house := *phouse

		items, err := r.crm.GetRooms(
			ctx,
			house.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			slog.ErrorContext(ctx, "can't resolve house rooms", slog.String("house", house.ID), slog.Any("err", err))

			return serverError()
		}

		return HouseRoomsList{Items: roomsToGraphQL(items)}, nil
	}
}

func roomsToGraphQL(rooms []*rooms.Room) []*Room {
	result := make([]*Room, 0, len(rooms))

	for _, item := range rooms {
		result = append(result, roomToGraphQL(item))
	}

	return result
}

func roomToGraphQL(room *rooms.Room) *Room {
	return &Room{
		ID:         room.ID,
		Name:       room.Name,
		Square:     room.Square,
		Level:      room.Level,
		CreatedAt:  room.CreatedAt,
		ModifiedAt: room.ModifiedAt,
	}
}

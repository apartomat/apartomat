package graphql

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/rooms"
	"go.uber.org/zap"
)

func (r *rootResolver) HouseRooms() HouseRoomsResolver {
	return &houseRoomsResolver{r}
}

type houseRoomsResolver struct {
	*rootResolver
}

func (r *houseRoomsResolver) List(ctx context.Context, obj *HouseRooms, limit int, offset int) (HouseRoomsListResult, error) {
	if phouse, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(**House); !ok {
		r.logger.Error("can't resolve house rooms", zap.Error(errors.New("unknown house")))

		return serverError()
	} else {
		house := *phouse

		items, err := r.useCases.GetRooms(
			ctx,
			house.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return forbidden()
			}

			r.logger.Error("can't resolve house rooms", zap.String("house", house.ID), zap.Error(err))

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

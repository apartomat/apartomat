package graphql

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/rooms"
	"log"
)

func (r *rootResolver) HouseRooms() HouseRoomsResolver {
	return &houseRoomsResolver{r}
}

type houseRoomsResolver struct {
	*rootResolver
}

func (r *houseRoomsResolver) List(ctx context.Context, obj *HouseRooms, limit int, offset int) (HouseRoomsListResult, error) {

	gtx := graphql.GetFieldContext(ctx)

	fmt.Printf("context: %#v\n", gtx.Parent.Parent)

	if phouse, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(**House); !ok {
		log.Printf("can't resolve house rooms: %s", errors.New("unknown house"))

		return serverError()
	} else {
		house := *phouse

		rooms, err := r.useCases.GetRooms(
			ctx,
			house.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return Forbidden{}, nil
			}

			log.Printf("can't resolve house (id=%s) rooms: %s", house.ID, err)

			return ServerError{Message: "internal server error"}, nil
		}

		return HouseRoomsList{Items: roomsToGraphQL(rooms)}, nil
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

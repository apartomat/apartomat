package graphql

import (
	"context"
)

func (r *rootResolver) House() HouseResolver {
	return &houseResolver{r}
}

type houseResolver struct {
	*rootResolver
}

func (r *houseResolver) Rooms(ctx context.Context, obj *House) (*HouseRooms, error) {
	return &HouseRooms{}, nil
}

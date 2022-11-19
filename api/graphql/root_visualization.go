package graphql

import (
	"context"

	. "github.com/apartomat/apartomat/internal/store/files"
	"github.com/apartomat/apartomat/internal/store/rooms"
)

func (r *rootResolver) Visualization() VisualizationResolver {
	return &visualizationResolver{r}
}

type visualizationResolver struct {
	*rootResolver
}

func (r *visualizationResolver) File(ctx context.Context, obj *Visualization) (*ProjectFile, error) {
	res, err := r.useCases.Files.List(ctx, IDIn(obj.File.ID), 1, 0)
	if err != nil {
		return nil, err
	}

	return projectFileToGraphQL(res[0]), nil
}

func (r *visualizationResolver) Room(ctx context.Context, obj *Visualization) (*Room, error) {
	if obj.Room == nil {
		return nil, nil
	}

	res, err := r.useCases.Rooms.List(ctx, rooms.IDIn(obj.Room.ID), 1, 0)
	if err != nil {
		return nil, err
	}

	return roomToGraphQL(res[0]), nil
}

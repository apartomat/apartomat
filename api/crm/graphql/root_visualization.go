package graphql

import (
	"context"
	"github.com/apartomat/apartomat/api/crm/graphql/dataloaders"
)

func (r *rootResolver) Visualization() VisualizationResolver {
	return &visualizationResolver{r}
}

type visualizationResolver struct {
	*rootResolver
}

func (r *visualizationResolver) File(ctx context.Context, obj *Visualization) (*File, error) {
	if obj.File == nil {
		return nil, nil
	}

	f, err := dataloaders.FromContext(ctx).Files.Load(ctx, obj.File.ID)
	if err != nil {
		return nil, err
	}

	return fileToGraphQL(f), nil
}

func (r *visualizationResolver) Room(ctx context.Context, obj *Visualization) (*Room, error) {
	if obj.Room == nil {
		return nil, nil
	}

	room, err := dataloaders.FromContext(ctx).Rooms.Load(ctx, obj.Room.ID)
	if err != nil {
		return nil, err
	}

	return roomToGraphQL(room), nil
}

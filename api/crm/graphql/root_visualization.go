package graphql

import (
	"context"
	"log/slog"

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
		slog.ErrorContext(ctx, "can't resolve visualization file", slog.String("file", obj.File.ID), slog.Any("err", err))

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
		slog.ErrorContext(ctx, "can't resolve visualization room", slog.String("room", obj.Room.ID), slog.Any("err", err))

		return nil, err
	}

	return roomToGraphQL(room), nil
}

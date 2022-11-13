package apartomat

import (
	"context"
	"errors"
	"github.com/apartomat/apartomat/internal/dataloader"
)

const dataLoadersCtxKey = "dataloaders"

type DataLoaders struct {
	Users *dataloader.UserLoader
}

func WithDataLoadersCtx(ctx context.Context, dataLoaders *DataLoaders) context.Context {
	return context.WithValue(ctx, dataLoadersCtxKey, dataLoaders)
}

func DataLoadersFromCtx(ctx context.Context) *DataLoaders {
	return ctx.Value(dataLoadersCtxKey).(*DataLoaders)
}

func UserLoaderFromCtx(ctx context.Context) (*dataloader.UserLoader, error) {
	loaders := DataLoadersFromCtx(ctx)
	if loaders == nil {
		return nil, errors.New("there are no data loaders in context")
	}

	if loaders.Users == nil {
		return nil, errors.New("there are no users data loader")
	}

	return loaders.Users, nil
}
